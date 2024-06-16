package main

import (
	"company-profile-api/config/middleware"
	"company-profile-api/config/respon"
	"company-profile-api/database/connection_db"
	"company-profile-api/handler"
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
		// proses account
		newRepository := account.NewRepository(connectionDB)
		newService := account.NewService(newRepository)
		newAuthMiddleware := middleware.NewAuthMiddleware()
		newAccountHandler := handler.NewAccountHandler(newService, newAuthMiddleware)
		// profile
		newRepositoryProf := profile.NewRepository(connectionDB)
		newServiceProf := profile.NewService(newRepositoryProf)
		newProfileHandler := handler.NewProfileHandler(newServiceProf)

		router := gin.Default()
		api := router.Group("api")
		api.POST("account", newAccountHandler.CreateAccountHandler)
		api.POST("auth", newAccountHandler.LoginAccountHandler)
		// router profile
		api.POST("profile", authMiddleware(newAuthMiddleware, newService), newProfileHandler.CreateProfileHand)
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
