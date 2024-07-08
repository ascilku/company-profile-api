package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/experience"
	"company-profile-api/src/user/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type experienceHandler struct {
	experienceService experience.Service
}

func NewExperience(experienceService experience.Service) *experienceHandler {
	return &experienceHandler{experienceService}
}

func (h *experienceHandler) CreateExperienceHand(g *gin.Context) {
	var experience experience.CreateExperienceRequest
	err := g.ShouldBindJSON(&experience)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responsApi := respon.ResponJson("failed create experience", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responsApi)
		return
	}
	account := g.MustGet("current_user_id").(account.Account)
	experience.AccountID = account.ID
	createExperience, err := h.experienceService.CreateExperienceServ(experience)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responsApi := respon.ResponJson("failed create experience", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responsApi)
		return
	}
	responMessage := gin.H{"is_available": createExperience}
	responsApi := respon.ResponJson("success create experience", http.StatusOK, []interface{}{}, responMessage)
	g.JSON(http.StatusOK, responsApi)
}
