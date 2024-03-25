package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomRequest struct {
	BoolField    bool    `json:"boolField"`
	IntField     int     `json:"intField"`
	StringField  string  `json:"stringField"`
	RuneField    rune    `json:"runeField"`
	PointerField *string `json:"pointerField"`
}

func Task3() {
	r := gin.Default()

	r.POST("/custom", func(c *gin.Context) {
		var req CustomRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		boolFieldValue := req.BoolField
		intFieldValue := req.IntField
		stringFieldValue := req.StringField
		runeFieldValue := req.RuneField
		pointerFieldValue := req.PointerField

		c.JSON(http.StatusOK, gin.H{
			"boolField":    boolFieldValue,
			"intField":     intFieldValue,
			"stringField":  stringFieldValue,
			"runeField":    runeFieldValue,
			"pointerField": pointerFieldValue,
		})
	})

	r.Run(":8080")
}
