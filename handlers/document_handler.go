package handlers

import (
	"context"
	"couchdb-go-app/db"
	"couchdb-go-app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-kivik/kivik/v4"
	"net/http"
)

// InsertDocument godoc
// @Summary Create a new document
// @Description Create a new document with the input payload
// @Tags documents
// @Accept json
// @Produce json
// @Param document body models.RequestDoc true "Document to create"
// @Success 201 {object} models.Document
// @Router /documents [post]
func InsertDocument(c *gin.Context) {
	var doc models.RequestDoc

	if err := c.BindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	database := db.GetDB()
	docID := doc.ID

	// Insert the document into CouchDB
	rev, err := database.Put(context.TODO(), docID, doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document: " + err.Error()})
		return
	}

	if docID == "" {
		docID = rev
	}

	responseDoc := models.Documents{
		ID:     docID,
		Rev:    rev,
		Name:   doc.Name,
		Gender: doc.Gender,
		Age:    doc.Age,
		Class:  doc.Class,
		Majors: doc.Majors,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Document created successfully",
		"payload": responseDoc,
	})
}

// GetDocumentByID godoc
// @Summary Get a document by ID
// @Description Get details of a document by its ID
// @Tags documents
// @Produce json
// @Param id path string true "Document ID"
// @Success 200 {object} models.Document
// @Router /documents/{id} [get]
func GetDocumentByID(c *gin.Context) {
	id := c.Param("id")
	database := db.GetDB()

	var doc models.Document
	err := database.Get(context.TODO(), id).ScanDoc(&doc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get document by ID successfully",
		"payload": doc,
	})
}

// UpdateDocument godoc
// @Summary Update a document
// @Description Update details of a document by its ID
// @Tags documents
// @Accept json
// @Produce json
// @Param id path string true "Document ID"
// @Param document body models.RequestUpdateDoc true "Updated document"
// @Success 200 {object} models.Document
// @Router /documents/{id} [put]
func UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var updatedData models.RequestUpdateDoc

	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	database := db.GetDB()

	var existingDoc models.Document
	resp := database.Get(context.TODO(), id)
	if err := resp.ScanDoc(&existingDoc); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found: " + err.Error()})
		return
	}

	updatedDocument := models.Documents{
		ID:     id,
		Rev:    existingDoc.Rev,
		Name:   updatedData.Name,
		Gender: updatedData.Gender,
		Age:    updatedData.Age,
		Class:  updatedData.Class,
		Majors: updatedData.Majors,
	}

	_, err := database.Put(context.TODO(), id, updatedDocument)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update document: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Document updated successfully",
		"payload": updatedDocument,
	})
}

// DeleteDocument godoc
// @Summary Delete a document
// @Description Delete a document by its ID
// @Tags documents
// @Param id path string true "Document ID"
// @Success 200 {string} string "Document deleted successfully"
// @Router /documents/{id} [delete]
func DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	database := db.GetDB()

	var doc models.Document
	row := database.Get(context.TODO(), id)
	if err := row.ScanDoc(&doc); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found: " + err.Error()})
		return
	}

	// Delete the document by passing the _rev
	_, err := database.Delete(context.TODO(), id, doc.Rev)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// GetAllDocuments godoc
// @Summary Get all documents
// @Description Retrieve all documents based on criteria
// @Tags documents
// @Produce json
// @Success 200 {array} models.Document
// @Failure 500 {string} string "Internal Server Error"
// @Failure 404 {string} string "No documents found"
// @Router /documents [get]
func GetAllDocuments(c *gin.Context) {
	database := db.GetDB()
	if database == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	rows := database.AllDocs(context.TODO(), kivik.Params(map[string]interface{}{
		"include_docs": true,
	}))
	defer func() {
		if err := rows.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close rows: " + err.Error()})
		}
	}()

	var results []models.Document
	for rows.Next() {
		var doc models.Document
		if err := rows.ScanDoc(&doc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan document: " + err.Error()})
			return
		}
		results = append(results, doc)
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No documents found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get all documents successfully",
		"payload": results,
	})
}

