package main

import (
	"errors"
	"net/http"

	"github.com/ArenAzibekyan/gin-helper/jsonapi"
	"github.com/ArenAzibekyan/logrus-helper/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.Default()

var Error = jsonapi.Error(jsonapi.NewErrorResponse, http.StatusOK)
var ErrorLong = jsonapi.ErrorLong(jsonapi.NewErrorResponseShort, http.StatusOK)

func recoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorLong(c, log, r, 101, "panic recovered")
		}
	}()
	c.Next()
}

func main() {
	r := gin.New()
	r.Use(recoverMiddleware)

	r.GET("/error", func(c *gin.Context) {
		err := errors.New("error text")
		Error(c, log, 102, err)
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("panic")
	})

	r.Run(":8888")
}
