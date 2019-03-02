package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/linehk/go-douban-movies/model"
)

func Movies(c *gin.Context) {
	start := c.DefaultQuery("s", "0")
	s, err := strconv.Atoi(start)
	if err != nil || s < 0 || s >= 250 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "must: 0 <= s <= 250")
		return
	}
	end := c.DefaultQuery("e", "250")
	e, err := strconv.Atoi(end)
	if err != nil || e < 0 || e > 250 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "must: 0 <= e <= 250")
		return
	}
	c.IndentedJSON(http.StatusOK, model.Movies[s:e])
}
