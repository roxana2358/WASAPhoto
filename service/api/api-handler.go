package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users/:userID", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:userID", rt.wrap(rt.setMyUsername))
	rt.router.GET("/users/:userID/stream", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userID/following/:followingID", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userID/following/:followingID", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:userID/ban/:bannedID", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userID/ban/:bannedID", rt.wrap(rt.unbanUser))
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/photos/:photoID", rt.wrap(rt.getImage))
	rt.router.DELETE("/photos/:photoID", rt.wrap(rt.deletePhoto))
	rt.router.POST("/photos/:photoID/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:photoID/comments/:commentID", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/photos/:photoID/likes/:userID", rt.wrap(rt.likePhoto))
	rt.router.PUT("/photos/:photoID/likes/:userID", rt.wrap(rt.unlikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
