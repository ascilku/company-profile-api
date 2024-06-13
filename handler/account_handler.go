package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/user/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type acountHandler struct {
	service account.Service
}

func NewAccountHandler(service account.Service) *acountHandler {
	return &acountHandler{service}
}

func (s *acountHandler) CreateAccountHandler(g *gin.Context) {
	var createAccountRequest account.CreateAccountRequest
	err := g.ShouldBindJSON(&createAccountRequest)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responJson := respon.ResponJson("failed create data account", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responJson)
		return
	} else {
		createAccount, err := s.service.CreateAccountService(createAccountRequest)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responJson := respon.ResponJson("failed create data account", http.StatusBadRequest, errorMessage, []interface{}{})
			g.JSON(http.StatusBadRequest, responJson)
			return
		}
		formatter := account.Formatter(createAccount)
		responJson := respon.ResponJson("succes create data account", http.StatusOK, []interface{}{}, formatter)
		g.JSON(http.StatusOK, responJson)
	}

}
