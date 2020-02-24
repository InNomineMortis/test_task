package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"test_task/internal/pkg/models"
)

/*
 Returning all posts and comments to them
*/
func GetPosts(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
		res := models.SelectPosts(db)
		c.JSON(http.StatusOK, res)
	}
}
