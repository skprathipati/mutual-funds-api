package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fund struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func GetFundsHandler(c *gin.Context) {
	funds := []Fund{
		{ID: "1", Name: "QQQ", Value: 700.00},
		{fjfID: "2", Name: "VOO", Value: 770.00},
	}

	c.JSON(http.StatusOK, funds)
}
