package main

import (
	"company-profile-api/config/middleware"
	"company-profile-api/config/respon"
	"company-profile-api/database/connection_db"
	"company-profile-api/handler"
	"company-profile-api/src/about"
	"company-profile-api/src/certification"
	"company-profile-api/src/experience"
	"company-profile-api/src/user/account"
	"company-profile-api/src/user/profile"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	connectionDB, err := connection_db.ConnectionDB()
	if err != nil {
		log.Fatal("failed connection to database ", err.Error())
	} else {
		// migration.MigrationAll(connectionDB)
		// connectionDB.AutoMigrate(&profile.ProfileImage{})
		// connectionDB.Migrator().DropTable(&account.Account{})
		// connectionDB.Migrator().DropTable(&profile.Profile{})
		// connectionDB.AutoMigrate(&certification.Certificate{})
		// connectionDB.AutoMigrate(&about.About{})
		// connectionDB.AutoMigrate(&experience.Experience{})
		// connectionDB.AutoMigrate(&skils.Skils{})
		// connectionDB.AutoMigrate(&tools.Tools{})
		// proses account
		newRepository := account.NewRepository(connectionDB)
		newService := account.NewService(newRepository)
		newAuthMiddleware := middleware.NewAuthMiddleware()
		newAccountHandler := handler.NewAccountHandler(newService, newAuthMiddleware)
		// profile
		newRepositoryProf := profile.NewRepository(connectionDB)
		newServiceProf := profile.NewService(newRepositoryProf)
		newProfileHandler := handler.NewProfileHandler(newServiceProf)
		// certificate
		newRepositoryCertif := certification.NewRepository(connectionDB)
		newServiceCertif := certification.NewService(newRepositoryCertif)
		newCertificateHandler := handler.NewCertificate(newServiceCertif)
		// about
		newRepositoryAbout := about.NewRepository(connectionDB)
		newServiceAbout := about.NewService(newRepositoryAbout)
		newAboutHandler := handler.NewAboutHandler(newServiceAbout)
		// experience
		newRepositoryExperience := experience.NewRepository(connectionDB)
		newServiceExperience := experience.NewService(newRepositoryExperience)
		newExperienceHandler := handler.NewExperience(newServiceExperience)

		router := gin.Default()
		api := router.Group("api")
		// auth account
		api.POST("auth", newAccountHandler.LoginAccountHandler)
		// account
		api.POST("account", newAccountHandler.CreateAccountHandler)
		// router profile
		api.POST("profile", authMiddleware(newAuthMiddleware, newService), newProfileHandler.CreateOrUpdateProfileHand)
		// certificate
		api.POST("certificate", authMiddleware(newAuthMiddleware, newService), newCertificateHandler.CreateCertificateHandler)
		api.GET("certificate", authMiddleware(newAuthMiddleware, newService), newCertificateHandler.FindAllCertificateHandler)
		api.DELETE("certificate", authMiddleware(newAuthMiddleware, newService), newCertificateHandler.DeleteOneCertificateServ)
		api.PUT("certificate", authMiddleware(newAuthMiddleware, newService))
		// about
		api.POST("about", authMiddleware(newAuthMiddleware, newService), newAboutHandler.CreateAboutHend) // ini bisa create or update
		api.GET("about", authMiddleware(newAuthMiddleware, newService), newAboutHandler.FindIdAccountHend)
		api.DELETE("about", authMiddleware(newAuthMiddleware, newService), newAboutHandler.DeleteAboutHend)
		// experience
		api.POST("experience", authMiddleware(newAuthMiddleware, newService), newExperienceHandler.CreateExperienceHand) // ini bisa create or update
		router.Run()
	}
}

func authMiddleware(auth middleware.AuthMiddleware, service account.Service) gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenHeader := g.GetHeader("Authorization")
		if !strings.Contains(tokenHeader, "Bearer") {
			errorMessage := gin.H{"errors": "invalid token"}
			errorResponJson := respon.ResponJson("failed access token", http.StatusUnauthorized, errorMessage, []interface{}{})
			g.AbortWithStatusJSON(http.StatusUnauthorized, errorResponJson)
			return
		}
		var tokenString = ""
		arrayToken := strings.Split(tokenHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		validationToken, err := auth.ValidationToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			errorResponJson := respon.ResponJson("failed access token", http.StatusUnauthorized, errorMessage, []interface{}{})
			g.AbortWithStatusJSON(http.StatusUnauthorized, errorResponJson)
			return
		} else {
			claim, ok := validationToken.Claims.(jwt.MapClaims)
			if !ok {
				errorMessage := gin.H{"errors": err.Error()}
				errorResponJson := respon.ResponJson("failed access token", http.StatusUnauthorized, errorMessage, []interface{}{})
				g.AbortWithStatusJSON(http.StatusUnauthorized, errorResponJson)
				return
			}
			userID := int(claim["user_id"].(float64))
			findByIDAccount, err := service.FindByIDAccountService(userID)
			if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				errorResponJson := respon.ResponJson("failed access token", http.StatusUnauthorized, errorMessage, []interface{}{})
				g.AbortWithStatusJSON(http.StatusUnauthorized, errorResponJson)
				return
			}
			g.Set("current_user_id", findByIDAccount)
		}
	}
}
