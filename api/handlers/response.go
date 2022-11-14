package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorModel struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Error interface{} `json:"error"`
}

// func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
// 	c.JSON(code, response.SuccessModel{
// 		Code:    code,
// 		Message: message,
// 		Data:    data,
// 	})
// }

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	logrus.Error(message)
	c.JSON(code, ErrorModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}
