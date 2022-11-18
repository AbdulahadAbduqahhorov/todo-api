package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx="userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.handleErrorResponse(c, http.StatusUnauthorized, "empty auth header", errors.New("empty auth header"))
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.handleErrorResponse(c, http.StatusUnauthorized, "invalid auth header", errors.New("invalid auth header"))
		return
	}

	user_id,err:=h.services.Authorization.ParseToken(headerParts[1])
	if err!=nil{
		h.handleErrorResponse(c, http.StatusUnauthorized, err.Error(), err)
		return
	}
	c.Set(userCtx,user_id)
}
