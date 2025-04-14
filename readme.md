# Gin Transfer Service

## Tecnologie Utilizzate

- **Framework**: Gin (Go)
- **Database**: SQLite
- **Routing**: Definito separatamente in `SetupRoutes.go`
- **Middleware**: Utilizzati per parsing dati e validazioni
- **Transazioni**: Tutte le operazioni critiche sono eseguite in transazione con supporto al rollback

## Database

- **Creazione delle tabelle**  
  Le query per la creazione delle tabelle si trovano nel file:  
  `db/query/create.sql`

- **Inserimento dati iniziali**  
  La tabella `bank_accounts` viene popolata tramite la query:  
  `db/query/insert.sql`

## Routing

Le rotte dell'applicazione sono gestite nel file `SetupRoutes.go`, con una configurazione centralizzata che permette di definire i percorsi e i middleware associati in modo semplice e modulare.

### Route Principale

#### `POST /sendTransfers`

Questa route utilizza due middleware principali:

1. **ParseData**  
   Questo middleware legge **il body** della richiesta e lo converte in un oggetto `TransferData` come segue:

   ```go
   type Transfer struct {
       Employee_id string  `json:"employee_id"`
       Name        string  `json:"name"`
       Iban        string  `json:"iban"`
       Amount      float64 `json:"amount"`
       Note        string  `json:"note"`
       Bic         string  `json:"bic"`
   }

   type TransferData struct {
       Organization_name string     `json:"organization_name"`
       Execution_date    string     `json:"execution_date"`
       Description       string     `json:"description"`
       Transfers         []Transfer `json:"transfers"`
   }
   ```

2. **CheckIfBalanceIsOk**  
   Questo middleware:
   - Verifica che il saldo dell'account mittente sia sufficiente per eseguire i trasferimenti.
   - Controlla che l'organizzazione specificata esista nel database.


## Handler `/handlers/SendTransfers.go`

L'handler associato alla route `/sendTransfers` gestisce l'intera logica del trasferimento e interagisce con il database tramite tre query principali:

1. Recupero ID account mittente  
   - Query: `db/query/getAccountId`

2. Inserimento dei trasferimenti nella tabella `transfers`  
   - Query: `db/query/insertTransfer`

3. Aggiornamento del saldo dell’account mittente  
   - Query: `db/query/updateBalance`

Tutte queste operazioni sono eseguite in una singola transazione, in modo da poter eseguire un rollback in caso di errori e mantenere la consistenza dei dati.

## Avvio del Progetto

1. Clona il repository
2. Il db (db/database.db) è già inizializzato, qualora ci fossero problemi è possibile eseguire nuovamente le seguenti query:
   - `db/query/create.sql`
   - `db/query/insert.sql`
3. Installa le dipendenze Go:
`go mod tidy`
4. Avvia il server
`go run main.go`

## Note

- Il json delle transazioni deve essere passato nel body della richiesta.
- La sezione "amount" di ogni transfer passato nel body deve essere in Euro, in quanto la conversione in Cents viene gestita internamente.
- Il programma si avvia sulla porta 8080, è possibile modificarla nel file `main.go`.
