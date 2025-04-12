package handlers

import (
	"database/sql"
	"test/dbUtils"
	"test/types"

	"github.com/gin-gonic/gin"
)

func SendHandlers(c *gin.Context) {
	transfersData := c.MustGet("TransfersDataKey").(types.TransferData)
	db := c.MustGet("DbKey").(*sql.DB)
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nell'inizializzazione della transazione"})
		c.Abort()
		return
	}
	defer tx.Rollback()

	id, ok := getAccountId(c, transfersData, tx)
	if !ok {
		return
	}

	ok = insertTransfers(c, transfersData, tx, id)
	if !ok {
		return
	}

	ok = updateBalance(c, tx, id)
	if !ok {
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nel commit della transazione"})
		return
	}

	c.JSON(201, gin.H{"success": "Operazione completata con successo"})
}

func getAccountId(c *gin.Context, transfersData types.TransferData, tx *sql.Tx) (int, bool) {
	query, err := dbUtils.ReadSQLFile("db/query/getAccountId.sql")
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nella lettura della query per il recupero dell'id dell'account dell'org"})
		c.Abort()
		return 0, false
	}

	var id int
	row := tx.QueryRow(query, transfersData.Organization_name)

	err = row.Scan(&id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nell'esecuzione della query per il recupero dell'id dell'account dell'org"})
		c.Abort()
		return 0, false
	}

	return id, true
}

func insertTransfers(c *gin.Context, transfersData types.TransferData, tx *sql.Tx, id int) bool {
	query, err := dbUtils.ReadSQLFile("db/query/insertTransfer.sql")
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nella lettura della query per l'inserimento dei transfer"})
		c.Abort()
		return false
	}

	for _, transfer := range transfersData.Transfers {
		_, err := tx.Exec(query, transfer.Name, transfer.Iban, transfer.Bic, transfer.Amount*100, id, transfer.Note)
		if err != nil {
			c.JSON(404, gin.H{"error": "Errore nell'esecuzione della query per l'inserimento dei transfer"})
			c.Abort()
			return false
		}
	}

	return true
}

func updateBalance(c *gin.Context, tx *sql.Tx, id int) bool {
	query, err := dbUtils.ReadSQLFile("db/query/updateBalance.sql")
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nella lettura della query per l'aggiornamento del saldo"})
		c.Abort()
		return false
	}

	totalAmountInCents := c.MustGet("TotalAmountInCentsKey")

	_, err = tx.Exec(query, totalAmountInCents, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Errore nell'esecuzione della query per l'aggiornamento del saldo"})
		c.Abort()
		return false
	}

	return true
}
