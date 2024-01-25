package main

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type row struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

var rows = []row{
	{ID: 1, Title: "Row One", Value: "this is the value of row one"},
	{ID: 2, Title: "Row two", Value: "this is the value of row two"},
	{ID: 3, Title: "Row Three", Value: "this is the value of row three"},
}

func getRows(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rows)
}

// func rowById(c *gin.Context) {
// 	id := c.Param("id")
// 	row, err := getRowById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, row)
// }

func getRowById(id int) (*row, error) {
	for i, r := range rows {
		if r.ID == id {
			return &rows[i], nil
		}
	}

	return nil, errors.New("row not found")
}

func deleteElement(index int) []row {
	return append(rows[:index], rows[index+1:]...)
}

func deleteRow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid row id"})
		return
	}

	deleteElement(id)

	c.IndentedJSON(http.StatusOK, rows)
}

func createRow(c *gin.Context) {
	var newRow row

	if err := c.BindJSON(&newRow); err != nil {
		return
	}

	newRow.ID = len(rows) + 1

	rows = append(rows, newRow)
	c.IndentedJSON(http.StatusCreated, newRow)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.GET("/rows", getRows)
	router.POST("/rows", createRow)
	router.DELETE("/rows/:id", deleteRow)
	router.Run("localhost:8080")
}
