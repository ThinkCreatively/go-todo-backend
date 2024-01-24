package main

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type row struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

var rows = []row{
	{ID: "1", Title: "Row One", Value: "this is the value of row one"},
	{ID: "2", Title: "Row two", Value: "this is the value of row two"},
	{ID: "3", Title: "Row Three", Value: "this is the value of row three"},
}

func getRows(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rows)
}

func rowById(c *gin.Context) {
	id := c.Param("id")
	row, err := getRowById(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, row)
}

func getRowById(id string) (*row, error) {
	for i, r := range rows {
		if r.ID == id {
			return &rows[i], nil
		}
	}

	return nil, errors.New("row not found")
}

func createRow(c *gin.Context) {
	var newRow row

	if err := c.BindJSON(&newRow); err != nil {
		return
	}

	newRow.ID = strconv.Itoa(len(rows) + 1)

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
	// router.DELETE("/rows", deleteRow)
	router.Run("localhost:8080")
}
