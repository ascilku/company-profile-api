package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/about"
	"company-profile-api/src/user/account"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type aboutHandler struct {
	about about.Service
}

func NewAboutHandler(about about.Service) *aboutHandler {
	return &aboutHandler{about}
}

func (a *aboutHandler) CreateAboutHend(g *gin.Context) {
	var aboutCreateReq about.CreateReq
	err := g.ShouldBindJSON(&aboutCreateReq)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responsAPi := respon.ResponJson(fmt.Sprintf("failed data about"), http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responsAPi)
		return
	}
	account := g.MustGet("current_user_id").(account.Account)
	aboutCreateReq.AccountID = account.ID
	_, message, err := a.about.CreateOrUpdateAboutServ(aboutCreateReq)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responsAPi := respon.ResponJson(fmt.Sprintf("failed %s data about", message), http.StatusBadRequest, errorMessage, []interface{}{})
		g.JSON(http.StatusBadRequest, responsAPi)
		return
	}

	responsAPi := respon.ResponJson(fmt.Sprintf("success %s data about", message), http.StatusOK, []interface{}{}, []interface{}{})
	g.JSON(http.StatusOK, responsAPi)
}

func (a *aboutHandler) FindIdAccountHend(g *gin.Context) {
	account := g.MustGet("current_user_id").(account.Account)
	findAbout, err := a.about.FindAboutServ(account.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responsAPi := respon.ResponJson("failed get data about", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responsAPi)
		return
	}
	formatter := about.Formatter(findAbout)
	responsAPi := respon.ResponJson("success get data about", http.StatusOK, []interface{}{}, formatter)
	g.JSON(http.StatusOK, responsAPi)
	return
}
