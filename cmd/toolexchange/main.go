package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"toolexchange"
)

type server struct {
	exchanger *toolexchange.Exchanger
}

func main() {
	s := server{
		exchanger: toolexchange.NewExchanger(),
	}
	router := gin.Default()
	// TODO: Change cors config to only accept selected sources
	router.Use(cors.Default())
	router.POST("/exchange", s.requestToken)
	router.GET("/exchange", s.requestItem)
	panic(router.Run(":4040"))
}

func (s server) requestToken(c *gin.Context) {
	var item toolexchange.Item
	if err := c.ShouldBind(&item); err == nil {
		token := s.exchanger.PutItem(item)
		c.String(http.StatusOK, token)
		return
	} else {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func (s server) requestItem(c *gin.Context) {
	token := c.Query("token")
	item, ok := s.exchanger.GetItem(token)
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	err := json.NewEncoder(c.Writer).Encode(item)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
	return
}
