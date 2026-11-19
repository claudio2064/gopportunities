package handler

import (
	"net/http"

	"github.com/claudio2064/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary      List opening
// @Description  Delete an opening by ID
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "Opening ID"
// @Success      200  {object}   schemas.Opening
// @Router       /opening [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
