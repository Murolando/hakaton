package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Murolando/hakaton_geo/ent"
	"github.com/gin-gonic/gin"
)

// @Summary kontur-start
// @Tags kontur
// @Description start contur game
// @ID kontur-start
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure default {object} map[string]interface{}
// @Router /api/class/ [get]
func (h *Handler) startKonturGame(c *gin.Context) {

	count_quiz, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if count_quiz == 0 {
		count_quiz = 1
	}
	kontur, err := h.service.StartKonturGame(count_quiz)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "kontur", kontur)
}

func (h *Handler) processKonturGame(c *gin.Context) {
	var input ent.ProcessRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cont, _ := c.Get("userId")
	userId := cont.(int64)
	fmt.Println(userId)
	kontur, err := h.service.ProcessKonturGame(&input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "contur-status", kontur)

}
