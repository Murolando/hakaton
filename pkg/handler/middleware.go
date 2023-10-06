package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty header")
		return
	}
	headerParts := strings.Split(header, " ")

	// if len(headerParts) != 2 {
	// 	newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	// 	return
	// }

	userId, err := h.service.Auth.ParseToken(headerParts[0])
	if err != nil || userId == 0 {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
}
