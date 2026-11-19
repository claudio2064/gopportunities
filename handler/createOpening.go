package handler

import (
	//"net/http"

	"net/http"

	"github.com/claudio2064/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary     Create opening
// @Description  Delete an opening by ID
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "Opening ID"
// @Success      200  {object}   schemas.Opening
// @Router       /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	//logger.Infof("request received: %+v", request

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}

	sendSuccess(ctx, "create_opening", opening)

}
