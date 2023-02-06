package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"wasa-photo/service/api/reqcontext"
	"wasa-photo/service/database"

	"github.com/julienschmidt/httprouter"
)

/**
* If the user exists, it returns the profile.
 */
func (rt *_router) getUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// token extraction
	_, err := getHeaderToken(r)
	if errors.Is(err, ErrUnauthorized) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't extract token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// username should be in URL
	if !r.URL.Query().Has("username") {
		// username is missing
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := r.URL.Query().Get("username")

	id, err := rt.db.GetUserId(username)
	// check errors
	if errors.Is(err, database.ErrUserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).WithField("username", username).Error("can't get user id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	var identifier ID
	identifier.Id = id
	_ = json.NewEncoder(w).Encode(identifier)
}
