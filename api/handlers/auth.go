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
	result,err:=h.services.Authorization.CreateUser(form)
	if err!=nil{
		h.handleErrorResponse(c,http.StatusInternalServerError,"server error",err)
		return
	}

	c.JSON(http.StatusCreated,map[string]interface{}{
		"id":result,
	})


}
func (h *Handler) signIn(c *gin.Context) {
	var u models.SignInInput
	if err := c.ShouldBindJSON(&u); err != nil {
		h.handleErrorResponse(c,http.StatusBadRequest,"parse error", err)
		return
	}
	token,err:=h.services.Authorization.GenerateToken(u)
	if err!=nil{
		h.handleErrorResponse(c,http.StatusInternalServerError,"server error",err)
		return
		
		
	}

	c.JSON(http.StatusCreated,map[string]interface{}{
		"token":token,
	})
}
