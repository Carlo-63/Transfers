package types

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
