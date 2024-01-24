package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type rowData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

var rows = []rowData{
	{ID: "1", Title: "Row One", Value: "this is the value of row one"},
	{ID: "2", Title: "Row two", Value: "this is the value of row two"},
	{ID: "3", Title: "Row Three", Value: "this is the value of row three"},
}
