INSERT INTO transfers (
    counterparty_name,
    counterparty_iban,
    counterparty_bic,
    amount_cents,
    bank_account_id,
    description
) VALUES (?, ?, ?, ?, ?, ?);