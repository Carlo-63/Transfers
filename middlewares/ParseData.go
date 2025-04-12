package middlewares

import (
	"encoding/json"
	"test/types"

	"github.com/gin-gonic/gin"
)

func ParseData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var transfersData types.TransferData

		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&transfersData)
		if err != nil {
			c.JSON(404, gin.H{"error": "Errore nel parsing dei dati"})
			c.Abort()
			return
		}

		if decoder.More() {
			c.JSON(404, gin.H{"error": "Il JSON inviato contiene più campi di quelli richiesti"})
			c.Abort()
			return
		}

		c.Set("TransfersDataKey", transfersData)
		c.Next()
	}
}
