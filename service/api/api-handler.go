package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users", rt.wrap(rt.getUserId))
	rt.router.GET("/users/:userID", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:userID", rt.wrap(rt.setMyUsername))
	rt.router.GET("/users/:userID/stream", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userID/following/:followingID", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userID/following/:followingID", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:userID/ban/:bannedID", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userID/ban/:bannedID", rt.wrap(rt.unbanUser))
	rt.router.POST("/posts", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/posts/:postID", rt.wrap(rt.getImage))
	rt.router.DELETE("/posts/:postID", rt.wrap(rt.deletePhoto))
	rt.router.POST("/posts/:postID/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/posts/:postID/comments/:commentID", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/posts/:postID/likes/:userID", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/posts/:postID/likes/:userID", rt.wrap(rt.unlikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
