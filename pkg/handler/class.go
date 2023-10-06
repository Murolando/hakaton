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
func (h *Handler) dashboardClass(c *gin.Context) {
	// get userId from Context
	if _, ok := c.Get("userId"); ok == false {
		newErrorResponse(c, http.StatusBadRequest, "userId not found")
		return
	}
	cont,_ := c.Get("userId")
	userId := cont.(int64)
	class, err := h.service.DashboardClass(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "class", class)
}

// @Summary class
// @Tags class
// @Security JwtKey
// @Description get child class
// @ID get-classes child 
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/class/my-classes [get]
func (h *Handler) myClass(c *gin.Context) {
	// get userId from Context
	if _, ok := c.Get("userId"); ok == false {
		newErrorResponse(c, http.StatusBadRequest, "userId not found")
		return
	}
	cont,_ := c.Get("userId")
	userId := cont.(int64)
	class, err := h.service.MyClass(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "class", class)
}
