package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Obj  interface{} `json:"obj"`
}

func Success(msg string, obj interface{}) *Message {
	return &Message{
		Code: http.StatusOK,
		Msg:  msg,
		Obj:  obj,
	}
}

func Fail(code int, err string) *Message {
	return &Message{
		Code: code,
		Err:  err,
	}
}

func (m *Message) JSON(c *gin.Context) {
	c.JSON(m.Code, m)
}
