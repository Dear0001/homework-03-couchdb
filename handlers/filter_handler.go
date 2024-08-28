package handlers

import (
	"context"
	"couchdb-go-app/db"
	"couchdb-go-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FilterDocuments godoc
// @Summary Filter documents by an age range and class
// @Description Retrieve documents where age is between n1 and n2 and class matches the user-provided input
// @Tags documents
// @Produce json
// @Param min_age query int true "Minimum Age to filter by"
// @Param max_age query int true "Maximum Age to filter by"
// @Param class query string true "Class to filter by"
// @Success 200 {array} models.Document
// @Router /documents/filter [get]
func FilterDocuments(c *gin.Context) {
	database := db.GetDB()

	// Parse query parameters
	minAge, err := strconv.Atoi(c.Query("min_age"))
	maxAge, err2 := strconv.Atoi(c.Query("max_age"))
	class := c.Query("class")

	if err != nil || err2 != nil || class == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Build and execute the Mango query
	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"$and": []map[string]interface{}{
				{"age": map[string]interface{}{"$gt": minAge, "$lt": maxAge}},
				{"class": class},
			},
		},
	}

	rows := database.Find(context.TODO(), query)
	defer rows.Close()

	var results []models.Document
	for rows.Next() {
		var doc models.Document
		if err := rows.ScanDoc(&doc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan document"})
			return
		}
		results = append(results, doc)
	}

	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over documents"})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No matching documents found",
			 "payload": []models.Document{},
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filtered documents retrieved successfully",
	 "payload": results,
	})
}
