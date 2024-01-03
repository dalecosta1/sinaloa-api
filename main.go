package main

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/dalecosta1/sinaloa-api/api"
	"github.com/dalecosta1/sinaloa-api/controller"
	"github.com/dalecosta1/sinaloa-api/docs" // Swagger generated files
	"github.com/dalecosta1/sinaloa-api/middlewares"
	"github.com/dalecosta1/sinaloa-api/repository"
	"github.com/dalecosta1/sinaloa-api/service"

	"github.com/joho/godotenv"
	
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
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("[ERROR] Error loading .env file:", err)
    }

	// CHeck if args are accepted or get from env
	isArgsEnabled := os.Getenv("ARGS_ENABLED")
	if isArgsEnabled == "true" {
		os.Setenv("PORT", os.Args[1])
		os.Setenv("SWAGGER_ENABLED", os.Args[2])
		os.Setenv("BASIC_AUTH_USER", os.Args[3])
		os.Setenv("BASIC_AUTH_USER", os.Args[4])
		os.Setenv("JWT_SECRET", os.Args[5])
	}

	// We need to setup this env variable from the env variables
	port := os.Getenv("PORT")
	isSwaggerEnabled := os.Getenv("SWAGGER_ENABLED")

	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Sinaloa CLI APIs"
	docs.SwaggerInfo.Description = "APIs to interact with the Sinaloa CLI, executing its commands."
	docs.SwaggerInfo.Version = "0.1.0"
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
