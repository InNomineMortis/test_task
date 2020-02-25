package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"test_task/internal/pkg/models/presenters"
)

/*
 Returning all posts and comments to them
*/
func GetPosts(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := presenters.PostPresenter.SelectPosts(db)
		c.JSON(http.StatusOK, res)
	}
}
