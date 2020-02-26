package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/internal/pkg/models/presenters"
)

func SetupPosts(e *gin.RouterGroup) {
	posts := e.Group("posts")
	posts.GET("", GetPosts)
}

/*
 Returning all posts and comments to them
*/
func GetPosts(c *gin.Context) {
	res := presenters.PostPresenter.SelectPosts()
	c.JSON(http.StatusOK, res)
}
