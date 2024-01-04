package main

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/dalecosta1/sinaloa-api/api"
	"github.com/dalecosta1/sinaloa-api/controller"
	"github.com/dalecosta1/sinaloa-api/docs" // Swagger generated files
	"github.com/dalecosta1/sinaloa-api/middlewares"
	"github.com/dalecosta1/sinaloa-api/repository"
	"github.com/dalecosta1/sinaloa-api/service"
	"github.com/dalecosta1/sinaloa-api/helpers"
	
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// Load the config
	helpers.LoadConfig()

	// Print the version
	fmt.Println("\nSinaloa CLI API v" + helpers.AppConfig.VERSION + "\n")

	// We need to setup this env variable from the env variables
	port := helpers.AppConfig.PORT // Use helpers.AppConfig
	isSwaggerEnabled := helpers.AppConfig.SWAGGER_ENABLED // Use helpers.AppConfig

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Sinaloa CLI APIs"
	docs.SwaggerInfo.Description = "APIs to interact with the Sinaloa CLI, executing its commands."
	docs.SwaggerInfo.Version = helpers.AppConfig.VERSION // Use helpers.AppConfig
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer videoRepository.CloseDB()

	server := gin.Default()

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	if isSwaggerEnabled == "true" {
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server.Run(":" + port)
}
