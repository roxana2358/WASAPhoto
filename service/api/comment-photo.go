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
* Adds a comment to the photo.
 */
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// The ID of the photo is a 64-bit unsigned integer
	postID, err := strconv.ParseUint(ps.ByName("postID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// read the comment from the request body
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// the body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// comment the photo
	commentID, err := rt.db.CreateComment(token, postID, comment.Comment)
	if errors.Is(err, database.ErrPostNotFound) || errors.Is(err, database.ErrUserBanned) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't create comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	var identifier ID
	identifier.Id = commentID
	_ = json.NewEncoder(w).Encode(identifier)
}
