package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"wasa-photo/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

/**
* Gets user's stream.
 */
func (rt *_router) getUserStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// check if token is allowed to request userID's stream
	if !checkAuth(token, userID) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbUserStream, err := rt.db.GetUserStream(userID) // list of userPosts
	// check errors
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't extract token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	var userStream []Userpost
	var userPost Userpost
	for i := 0; i < len(dbUserStream); i++ {
		userPost.UserPostFromDatabase(dbUserStream[i])
		userStream = append(userStream, userPost)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userStream)
}
