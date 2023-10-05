package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary class
// @Tags class
// @Description get lists class
// @ID get-classes
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/class/ [get]
func (h *Handler) allClass(c *gin.Context) {
	class, err := h.service.AllClass()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "class", class)
}
