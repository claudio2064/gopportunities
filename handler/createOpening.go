package handler

import (
	//"net/http"

	"net/http"

	"github.com/claudio2064/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary     Create opening
// @Description  Create a nes job opening
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        request body CreateOpeningRequest true "Request body"
// @Success      200  {object}   CreateOpeningResponse
// @Failure 400 (object) ErrorResponse
// @Failure 500 (object) ErrorResponse
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
