package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"wasa-photo/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

/**
* Uploads the photo send in the request body.
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
	// get next post id
	postId := rt.db.GetNextPostId()

	// parse request
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't parse request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// create photo file
	file, _, err := r.FormFile("file")
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't form file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// check if directory (photo storage) exists, create if not
	err = os.MkdirAll("tmp/photos/", 0777)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't create dir")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// create file
	fileName := "tmp/photos/" + strconv.FormatInt(int64(postId), 10) + ".png"
	imgFile, err := os.Create(fileName)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't create file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer imgFile.Close()
	// copy photo in file
	_, err = io.Copy(imgFile, file)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't copy file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// proceed uploading the photo
	id, err := rt.db.CreatePhoto(token, postId, time, date, fileName)
	if err != nil {
		// error on our side: log the error and send a 500 to the user
		ctx.Logger.WithError(err).Error("can't upload photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	var identifier ID
	identifier.Id = id
	_ = json.NewEncoder(w).Encode(identifier)
}
