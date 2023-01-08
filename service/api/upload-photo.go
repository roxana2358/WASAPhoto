package api

import (
	"encoding/json"
	"errors"
	"image/png"
	"net/http"
	"strings"
	"time"
	"wasa-photo/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

/**
* Uploads a photo.
 */
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// get time and date
	parsed_time := strings.Split(time.Now().UTC().String(), " ")
	var date string = parsed_time[0]
	var time string = parsed_time[1]

	// get image
	image, err := png.Decode(r.Body)
	if err != nil {
		// the body was not parseable, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// proceed uploading the photo
	id, err := rt.db.CreatePhoto(token, time, date, image)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.Error("can't upload photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	var identifier ID
	identifier.Id = id
	_ = json.NewEncoder(w).Encode(identifier)
}
