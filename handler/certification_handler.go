package handler

import (
	"company-profile-api/config/error_validation"
	"company-profile-api/config/respon"
	"company-profile-api/src/certification"
	"company-profile-api/src/user/account"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type certificate struct {
	certificationService certification.Service
}

func NewCertificate(certificationService certification.Service) *certificate {
	return &certificate{certificationService}
}

func (c *certificate) CreateCertificateHandler(g *gin.Context) {
	var keyCertification certification.CreateCertificateRequest
	err := g.ShouldBind(&keyCertification)
	if err != nil {
		errorMessage := gin.H{"errors": error_validation.ErrorValidation(err)}
		responApi := respon.ResponJson("Failed Create Certificate", http.StatusUnprocessableEntity, errorMessage, []interface{}{})
		g.JSON(http.StatusUnprocessableEntity, responApi)
		return
		return
	} else {
		fileCertificate, err := g.FormFile("fileCertificate")
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			responApi := respon.ResponJson("Failed Create Certificate", http.StatusBadRequest, errorMessage, []interface{}{})
			g.JSON(http.StatusBadRequest, responApi)
			return
		} else {
			account := g.MustGet("current_user_id").(account.Account)
			accountID := account.ID
			path := fmt.Sprintf("assets/image_certificate/%d-%s", accountID, fileCertificate.Filename)

			keyCertification.AccountID = accountID
			keyCertification.FileCertificate = path
			fmt.Println("accountIDaccountIDaccountIDaccountID")
			fmt.Println(accountID)
			_, err := c.certificationService.CreateCertificationServ(keyCertification)
			if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				responApi := respon.ResponJson("Failed Create Certificate", http.StatusBadRequest, errorMessage, []interface{}{})
				g.JSON(http.StatusBadRequest, responApi)
				return
			} else {
				err = g.SaveUploadedFile(fileCertificate, path)
				if err != nil {
					errorMessage := gin.H{"errors": err.Error()}
					responApi := respon.ResponJson("Failed Create Certificate", http.StatusBadRequest, errorMessage, []interface{}{})
					g.JSON(http.StatusBadRequest, responApi)
					return
				}
				responApi := respon.ResponJson("Success Create Certificate", http.StatusOK, []interface{}{}, []interface{}{})
				g.JSON(http.StatusOK, responApi)
			}

		}

	}
}
