package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/middleware"
	"company-profile-api/config/respon"
	"company-profile-api/src/user/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type acountHandler struct {
	service account.Service
	auth    middleware.AuthMiddleware
}

func NewAccountHandler(service account.Service, auth middleware.AuthMiddleware) *acountHandler {
	return &acountHandler{service, auth}
}

func (h *acountHandler) CreateAccountHandler(g *gin.Context) {
	var createAccountRequest account.CreateAccountRequest
	err := g.ShouldBindJSON(&createAccountRequest)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responJson := respon.ResponJson("failed create data account", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responJson)
		return
	} else {
		_, err := h.service.CreateAccountService(createAccountRequest)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responJson := respon.ResponJson("failed create data account", http.StatusBadRequest, errorMessage, []interface{}{})
			g.JSON(http.StatusBadRequest, responJson)
			return
		}
		responJson := respon.ResponJson("succes create data account", http.StatusOK, []interface{}{}, []interface{}{})
		g.JSON(http.StatusOK, responJson)
	}
}

func (h *acountHandler) LoginAccountHandler(g *gin.Context) {
	var createAccountRequest account.CreateAccountRequest
	err := g.ShouldBindJSON(&createAccountRequest)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responJson := respon.ResponJson("failed login access user account", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responJson)
		return
	} else {
		loginAccountService, err := h.service.LoginAccountService(createAccountRequest)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responJson := respon.ResponJson("failed login access user account", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
			g.JSON(http.StatusUnprocessableEntity, responJson)
			return
		} else {
			generateToken, err := h.auth.GenerateToken(loginAccountService.ID)
			if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				responJson := respon.ResponJson("failed login access user account", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
				g.JSON(http.StatusUnprocessableEntity, responJson)
				return
			}
			formatter := account.Formatter(loginAccountService, generateToken)
			responJson := respon.ResponJson("succes login access user account", http.StatusOK, []interface{}{}, formatter)
			g.JSON(http.StatusOK, responJson)
		}
	}
}
