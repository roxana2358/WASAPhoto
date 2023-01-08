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
* Removes comment from photo.
 */
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	postID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// The ID of the comment is a 64-bit unsigned integer
	commentID, err := strconv.ParseUint(ps.ByName("commentID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteComment(token, postID, commentID)
	if errors.Is(err, database.ErrCommentNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).WithField("commentID", commentID).Error("can't delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
