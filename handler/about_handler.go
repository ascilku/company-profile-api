package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/about"
	"company-profile-api/src/user/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type aboutHandler struct {
	about about.Service
}

func NewAboutHandler(about about.Service) *aboutHandler {
	return &aboutHandler{about}
}

func (a *aboutHandler) CreateAboutServ(g *gin.Context) {
	var aboutCreateReq about.CreateReq
	err := g.ShouldBindJSON(&aboutCreateReq)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responsAPi := respon.ResponJson("failed create data about", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responsAPi)
		return
	}
	account := g.MustGet("current_user_id").(account.Account)
	aboutCreateReq.AccountID = account.ID
	_, err = a.about.CreateAboutServ(aboutCreateReq)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responsAPi := respon.ResponJson("failed create data about", http.StatusBadRequest, errorMessage, []interface{}{})
		g.JSON(http.StatusBadRequest, responsAPi)
		return
	}

	responsAPi := respon.ResponJson("success create data about", http.StatusOK, []interface{}{}, []interface{}{})
	g.JSON(http.StatusOK, responsAPi)
}
