package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"wasa-photo/service/api/reqcontext"
	"wasa-photo/service/database"

	"github.com/julienschmidt/httprouter"
)

/**
* If the username is available it updates the old one; otherwise it notifies the user that
* the username is already used.
 */
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// the user ID in the path is a 64-bit unsigned integer
	userId, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
	if err != nil {
		// the value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the user that requested the action is authorized
	if !checkAuth(token, userId) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// read the new username from the request body
	var username Username
	err = json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		// the body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// proceed updating the username
	err = rt.db.UpdateUsername(token, username.Username)
	if errors.Is(err, database.ErrUserNotFound) {
		// the user does not exist, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUsernameNotAvailable) {
		// username not available
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).WithField("username", username).Error("can't update username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
