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
* If the user exists, it returns the profile.
 */
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	userID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbUserProfile, err := rt.db.GetUserProfile(token, userID)
	// check errors
	if errors.Is(err, database.ErrUserNotFound) || errors.Is(err, database.ErrUserBanned) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).WithField("userID", userID).Error("can't get user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	var userProfile Userprofile
	userProfile.UserProfileFromDatabase(dbUserProfile)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userProfile)
}
