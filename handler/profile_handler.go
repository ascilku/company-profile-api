package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/user/account"
	"company-profile-api/src/user/profile"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type profileHandler struct {
	service profile.Service
}

func NewProfileHandler(service profile.Service) *profileHandler {
	return &profileHandler{service}
}

func (p *profileHandler) CreateOrUpdateProfileHand(g *gin.Context) {
	var keyCreateProfile profile.CreateProfileReq
	err := g.ShouldBind(&keyCreateProfile)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		errorResponJson := respon.ResponJson("failed create profileee", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, errorResponJson)
		return
	} else {
		file, err := g.FormFile("image")
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			errorResponJson := respon.ResponJson("failed create profileee", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
			g.JSON(http.StatusUnprocessableEntity, errorResponJson)
			return
		} else {
			account := g.MustGet("current_user_id").(account.Account)
			accountID := account.ID
			path := fmt.Sprintf("assets/image_profile/%d-%s", accountID, file.Filename)
			err := g.SaveUploadedFile(file, path)
			if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				errorResponJson := respon.ResponJson("failed create profileee", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
				g.JSON(http.StatusUnprocessableEntity, errorResponJson)
				return
			} else {

				keyCreateProfile.AccountID = accountID
				keyCreateProfile.ProfileImages = path
				_, message, err := p.service.CreateOrUpdateProfileServ(keyCreateProfile)
				if err != nil {
					errorMessage := gin.H{"errors": err.Error()}
					errorResponJson := respon.ResponJson("failed create profile", http.StatusBadRequest, errorMessage, []interface{}{})
					g.JSON(http.StatusBadRequest, errorResponJson)
					return
				}
				errorResponJson := respon.ResponJson(fmt.Sprintf("success %s profile", message), http.StatusOK, []interface{}{}, []interface{}{})
				g.JSON(http.StatusOK, errorResponJson)
			}
		}
	}
}
