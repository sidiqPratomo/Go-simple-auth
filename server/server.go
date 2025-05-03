package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appvalidator"
	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
	"github.com/sidiqPratomo/DJKI-Pengaduan/handler"
	"github.com/sidiqPratomo/DJKI-Pengaduan/middleware"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
	"github.com/sirupsen/logrus"
)

type routerOpts struct {
	Authentication     *handler.AuthenticationHandler
	User 			   *handler.UserHandler
}

type utilOpts struct {
	JwtHelper util.TokenAuthentication
}

func newRouter(h routerOpts, u utilOpts, config *config.Config, log *logrus.Logger) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()

	router.ContextWithFallback = true

	appvalidator.AppValidator()

	router.Use(
		middleware.Logger(log),
		middleware.RequestIdHandlerMiddleware,
		middleware.ErrorHandlerMiddleware,
		gin.Recovery(),
	)

	authMiddleware := middleware.AuthMiddleware(u.JwtHelper, config)

	// userAuthorizationMiddleware := middleware.UserAuthorizationMiddleware
	// doctorAuthorizationMiddleware := middleware.DoctorAuthorizationMiddleware
	// pharmacyManagerAuthorizationMiddleware := middleware.PharmacyManagerAuthorizationMiddleware
	// adminAuthorizationMiddleware := middleware.AdminAuthorizationMiddleware

	corsRouting(router, corsConfig)
	// authenticationRouting(router, h.Authentication)

	api := router.Group("/api/v1")
	{
		authenticationRouting(api, h.Authentication, authMiddleware)
		userRouting(api, h.User, authMiddleware)
	}

	return router
}

func corsRouting(router *gin.Engine, configCors cors.Config) {
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"POST", "GET", "PUT", "PATCH", "DELETE"}
	configCors.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Accept", "User-Agent", "Cache-Control"}
	configCors.ExposeHeaders = []string{"Content-Length"}
	configCors.AllowCredentials = true
	router.Use(cors.New(configCors))
}

func authenticationRouting(router *gin.RouterGroup, handler *handler.AuthenticationHandler, authMiddleware gin.HandlerFunc) {
	authRouter:= router.Group("/auth")

	authRouter.POST("/register-user", handler.RegisterUser)
	authRouter.POST("/verify-registration", handler.VerifyRegisterUser)
	authRouter.POST("/signin", handler.Login)
	authRouter.POST("/verify-otp", handler.VerifyLoginUser)
	authRouter.POST("/refresh", authMiddleware,  handler.RefreshToken)
}

func userRouting(router *gin.RouterGroup, handler *handler.UserHandler, authMiddleware gin.HandlerFunc) {
	authRouter:= router.Group("/users")

	authRouter.GET("/",authMiddleware, handler.IndexUser)
	authRouter.GET("/:id",authMiddleware, handler.ReadUser)
	authRouter.PUT("/:id",authMiddleware, handler.UpdateUser)
}