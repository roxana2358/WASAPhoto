package api

import (
	"errors"
	"image/png"
	"net/http"
	"strconv"
	"wasa-photo/service/api/reqcontext"
	"wasa-photo/service/database"

	"github.com/julienschmidt/httprouter"
)

/**
* Gets the image requested.
 */
func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// token extraction; real token not necessary because if we received the photoID user is allowed to request it
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
	postId, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// request the photo
	img, err := rt.db.GetImage(postId)
	if errors.Is(err, database.ErrFileNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't decode image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "image/png")
	_ = png.Encode(w, img)
}
