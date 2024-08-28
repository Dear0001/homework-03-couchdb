package handlers

import (
	"context"
	"couchdb-go-app/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kivik/kivik/v4"
	"github.com/go-kivik/kivik/v4/cmd/kivik/errors"
	_ "io"
	"net/http"
)

// UploadFile godoc
// @Summary Upload a file and attach it to a specific student's document
// @Description Upload a file and store it as an attachment in a student's CouchDB document
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param stu-id path string true "Student ID (Document ID)"
// @Param file formData file true "File to upload"
// @Success 200 {string} string "File uploaded successfully"
// @Failure 400 {string} string "Failed to upload file"
// @Router /upload/{stu-id} [post]
func UploadFile(c *gin.Context) {
	docID := c.Param("stu-id")
	fmt.Println("UploadFile handler called")
	fmt.Println("Student ID:", docID)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()

	doc := make(map[string]interface{})
	err = db.Get(context.TODO(), docID).ScanDoc(&doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get document."})
		return
	}
	rev, _ := doc["_rev"].(string)
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer openedFile.Close()

	_, err = db.PutAttachment(context.TODO(), docID, &kivik.Attachment{
		Filename:    file.Filename,
		Content:     openedFile,
		ContentType: file.Header.Get("Content-Type"),
	}, kivik.Rev(rev))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "File uploaded successfully"})
}

// GetFile godoc
// @Summary Get a file (attachment) by filename and student ID
// @Description Retrieve a file from CouchDB for a specific student document
// @Tags file
// @Produce octet-stream
// @Param stu-id path string true "Student ID (Document ID)"
// @Param filename path string true "Filename to retrieve"
// @Success 200 {file} file "File retrieved successfully"
// @Failure 404 {string} string "File not found"
// @Failure 500 {string} string "Internal server error"
// @Router /file/{stu-id}/{filename} [get]
func GetFile(c *gin.Context) {
	docID := c.Param("stu-id")
	filename := c.Param("filename")

	db := db.GetDB()

	// Retrieve the document to get the revision
	doc := make(map[string]interface{})
	err := db.Get(context.TODO(), docID).ScanDoc(&doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get document."})
		return
	}

	rev, _ := doc["_rev"].(string)

	// Retrieve the attachment
	attachment, err := db.GetAttachment(context.TODO(), docID, filename, kivik.Rev(rev))
	if err != nil {
		if errors.Is(err, kivik.ErrDatabaseClosed) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file"})
		return
	}
	// Set headers and serve the file
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", attachment.ContentType)
	c.DataFromReader(http.StatusOK, attachment.Size, attachment.ContentType, attachment.Content, nil)
}
