package main

import (
	"couchdb-go-app/db"
	"couchdb-go-app/handlers"
	_ "couchdb-go-app/models"
	"github.com/gin-gonic/gin"

	_ "couchdb-go-app/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Student API
// @version 1.0
// @description This is a sample server for managing students.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	db.InitCouchDB()
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/documents", handlers.InsertDocument)
	router.GET("/documents", handlers.GetAllDocuments)
	router.GET("/documents/:id", handlers.GetDocumentByID)
	router.PUT("/documents/:id", handlers.UpdateDocument)
	router.DELETE("/documents/:id", handlers.DeleteDocument)
	router.GET("/documents/filter", handlers.FilterDocuments)

	router.POST("/upload/:stu-id", handlers.UploadFile)
	router.GET("/file/:stu-id/:filename", handlers.GetFile)

	router.Run()
}
