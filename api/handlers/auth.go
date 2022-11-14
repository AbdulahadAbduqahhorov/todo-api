package handlers

import (
	"net/http"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var form models.User
	if err := c.ShouldBindJSON(&form); err != nil {
		h.handleErrorResponse(c,http.StatusBadRequest,"parse error", err)
		return
	}
}
func (h *Handler) signIn(c *gin.Context) {

}
