package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/utils"
)

var (
	uptime time.Time
)

const (
	/*V1Path is a router group for the v1 version of storage node*/
	V1Path = "/api/v1"

	/*BooksPath is the path for books*/
	BooksPath = "/books"
)

type StatusRes struct {
	Status string `json:"status" example:"status of the request"`
}

func init() {
}

/*CreateRoutes creates our application's routes*/
func CreateRoutes() {
	uptime = time.Now()

	router := returnEngine()

	// Test app is running
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Server is running",
			"uptime":  fmt.Sprintf("%v", time.Now().Sub(uptime)),
		})
	})

	setupV1Paths(returnV1Group(router))

	// Listen and Serve
	err := router.Run(":" + os.Getenv("PORT"))
	utils.PanicOnError(err)
}

func returnEngine() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()

	// TODO:  update to only allow our frontend and localhost
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	return router
}

func returnV1Group(router *gin.Engine) *gin.RouterGroup {
	return router.Group(V1Path)
}

func setupV1Paths(v1Router *gin.RouterGroup) {
	v1Router.POST(BooksPath, CreateBookHandler)
	v1Router.GET(BooksPath, GetAllBooksHandler)
	v1Router.DELETE(BooksPath+"/:id", DeleteBookHandler)
	v1Router.PUT(BooksPath+"/:id", UpdateBookHandler)
}
