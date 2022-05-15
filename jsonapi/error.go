package jsonapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response interface {
	SetContent(code int, mess string)
}

func ErrorString(newResp func() Response, httpCode int) func(c *gin.Context, log *logrus.Entry, code int, mess string) {
	return func(c *gin.Context, log *logrus.Entry, code int, mess string) {
		log.WithField("code", code).Error(mess)

		resp := newResp()
		resp.SetContent(code, mess)

		c.AbortWithStatusJSON(httpCode, resp)
	}
}

func Error(newResp func() Response, httpCode int) func(c *gin.Context, log *logrus.Entry, code int, err error) {
	errorString := ErrorString(newResp, httpCode)

	return func(c *gin.Context, log *logrus.Entry, code int, err error) {
		errorString(c, log, code, err.Error())
	}
}

func ErrorLong(newResp func() Response, httpCode int) func(c *gin.Context, log *logrus.Entry, err interface{}, code int, mess string) {
	errorString := ErrorString(newResp, httpCode)

	return func(c *gin.Context, log *logrus.Entry, err interface{}, code int, mess string) {
		if err != nil {
			log = log.WithField("err", err)
		}
		errorString(c, log, code, mess)
	}
}
