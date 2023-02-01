package api

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"wasa-photo/service/api/reqcontext"
	"wasa-photo/service/database"

	"github.com/julienschmidt/httprouter"
)

/**
* Gets the image requested.
 */
func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// The ID of the photo is a 64-bit unsigned integer
	postId, err := strconv.ParseUint(ps.ByName("postID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// request the filename
	filename, err := rt.db.GetImage(postId)
	if errors.Is(err, database.ErrPostNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// open file
	imgFile, err := os.Open(filename)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't open image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer imgFile.Close()

	// send output to user
	w.Header().Set("Content-Type", "image/*")
	_, err = io.Copy(w, imgFile)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't send image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
