package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
	"strings"
)

var (
	uptime time.Time
)

const (
	/*V1Path is a router group for the v1 version of storage node*/
	V1Path = "/api/v1"

	/*ToolsPath is the path for tools*/
	ToolsPath = "/tools"
)

const MaxRequestSize = int64(1024*1024*.25) + 1000

type parsableObjectInterface interface {
	// Return the reference of object
	getObjectRef() interface{}
	getObjectAsString() string
}

type StatusRes struct {
	Status string `json:"status" example:"status of the request"`
}

func init() {
}

/*CreateRoutes creates our application's routes*/
func CreateRoutes() {
	uptime = time.Now()

	router := returnEngine()

	// Test frontend is running
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

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func setupV1Paths(v1Router *gin.RouterGroup) {
	v1Router.POST(ToolsPath, CreateToolHandler)
	v1Router.GET(ToolsPath, GetAllToolsHandler)
	v1Router.DELETE(ToolsPath+"/:id", DeleteToolHandler)
	v1Router.PUT(ToolsPath+"/:id", UpdateToolHandler)

	v1Router.Static("/assets", "./assets")
}
