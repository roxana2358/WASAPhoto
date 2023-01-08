package api

import (
	"errors"
	"net/http"
	"strconv"
	"wasa-photo/service/api/reqcontext"
	"wasa-photo/service/database"

	"github.com/julienschmidt/httprouter"
)

/**
* Removes user from banned list.
 */
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// token extraction
	token, err := getHeaderToken(r)
	if errors.Is(err, ErrUnauthorized) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't extract token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// The user ID in the path is a 64-bit unsigned integer
	userId, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the user that requested the action is authorized
	if !checkAuth(token, userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// The banned ID in the path is a 64-bit unsigned integer - user to be unbanned
	unbanId, err := strconv.ParseUint(ps.ByName("bannedID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// proceed unfollowing the user
	err = rt.db.DeleteFollow(token, unbanId)
	if errors.Is(err, database.ErrBanNotFound) {
		// the "unban" user does not exist or token wasn't following the user
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't unban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
