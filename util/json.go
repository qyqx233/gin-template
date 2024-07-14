package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(c *gin.Context, rq any, dump bool) (err error) {
	var bs []byte
	bs, err = io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if dump {
		fmt.Println(string(bs))
	}
	if err = json.Unmarshal(bs, rq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "json error",
		})
		return
	}

	// if err = c.ShouldBindJSON(&rq); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":    500,
	// 		"message": "success",
	// 	})
	// 	return
	// }
	return
}

func RenderSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"success": true,
	})
}

func RenderError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"message": err.Error(),
		"code":    500,
	})
}
