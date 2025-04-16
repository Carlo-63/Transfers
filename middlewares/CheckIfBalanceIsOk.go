package middlewares

import (
	"database/sql"

	"test/dbUtils"
	"test/types"

	"github.com/gin-gonic/gin"
)

func CheckIfBalanceIsOk() gin.HandlerFunc {
	return func(c *gin.Context) {
		transfersData := c.MustGet("TransfersDataKey").(types.TransferData)

		totalAmountInCents := CalculateTotalAmountInCents(transfersData)
		organizationName := transfersData.Organization_name

		c.Set("TotalAmountInCentsKey", totalAmountInCents)

		// FIXME: qui ti colleghi al DB una volta per richiesta. Dovresti aprire una connessione (o un pool di connessioni) allo startup e riusare sempre quelle.
		db, err := dbUtils.ConnectToDb()
		if err != nil {
			// FIXME: perchè mi ritorni 404? 404 sta per "NotFound". Significa che non è stata trovata una risorsa all'interno del nostro store (accountID inesistente). Qui è più appropriato usare un 500.
			c.JSON(404, gin.H{"error": "Errore nella connessione al database"})
			c.Abort()
			return
		}
		c.Set("DbKey", db)

		query, err := dbUtils.ReadSQLFile("db/query/checkFunds.sql")
		if err != nil {
			// FIXME: stesso a sopra
			c.JSON(404, gin.H{"error": "Errore nella lettura della query per il controllo dei fondi"})
			c.Abort()
			return
		}

		var hasFunds bool
		row := db.QueryRow(query, totalAmountInCents, organizationName)

		err = row.Scan(&hasFunds)
		if err != nil {
			if err == sql.ErrNoRows {
				// NICETOHAVE: usami le constanti fornite dal package net/http. Questi sono "magic strings", un tipico code smell segnalato da Sonar.
				c.JSON(404, gin.H{"error": "Organizzazione non trovata"})
			}
			c.Abort()
			return
		}

		if !hasFunds {
			c.JSON(422, gin.H{"error": "Fondi insufficienti"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CalculateTotalAmountInCents(transfersData types.TransferData) int {
	counter := 0.0

	for _, transfer := range transfersData.Transfers {
		counter += transfer.Amount
	}
	counter *= 100

	return int(counter)
}
