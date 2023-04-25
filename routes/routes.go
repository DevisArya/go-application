package routes

import (
	"log"
	"os"
	"praktikum/controller"
	m "praktikum/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.Pre(mid.AddTrailingSlash())
	m.LogMiddleware(e)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	Jwt := mid.JWT([]byte(secretKey))

	gUser := e.Group("/users")
	gUser.GET("/", controller.GetUsersController, Jwt)
	gUser.GET("/:id/", controller.GetUserController, Jwt)
	gUser.DELETE("/:id/", controller.DeleteUserController, Jwt)
	gUser.PUT("/:id/", controller.UpdateUserController, Jwt)
	gUser.POST("/login/", controller.LoginUserController)
	gUser.POST("/", controller.CreateUserController)

	gBook := e.Group("/books")
	gBook.GET("/", controller.GetBooksController, Jwt)
	gBook.GET("/:id/", controller.GetBookController, Jwt)
	gBook.POST("/", controller.CreateBookController, Jwt)
	gBook.DELETE("/:id/", controller.DeleteBookController, Jwt)
	gBook.PUT("/:id/", controller.UpdateBookController, Jwt)

	gBlog := e.Group("/blogs")
	gBlog.GET("/", controller.GetBlogsController, Jwt)
	gBlog.GET("/:id/", controller.GetBlogController, Jwt)
	gBlog.POST("/", controller.CreateBlogController, Jwt)
	gBlog.DELETE("/:id/", controller.DeleteBlogController, Jwt)
	gBlog.PUT("/:id/", controller.UpdateBlogController, Jwt)

	return e
}
