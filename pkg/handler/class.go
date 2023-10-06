package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary class
// @Tags class
// @Security JwtKey
// @Description get child dashboard
// @ID get-classes
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/class/dashboard [get]
func (h *Handler) allClass(c *gin.Context) {
	// get userId from Context
	if _, ok := c.Get("userId"); ok == false {
		newErrorResponse(c, http.StatusBadRequest, "userId not found")
		return
	}
	cont,_ := c.Get("userId")
	userId := cont.(int64)
	class, err := h.service.AllClass(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "class", class)
}
