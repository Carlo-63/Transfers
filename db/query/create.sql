-- Creazione della tabella bank_accounts
CREATE TABLE bank_accounts (
    id INTEGER PRIMARY KEY,
    organization_name TEXT,
    balance_cents INTEGER,
    iban TEXT,
    bic TEXT
);

-- Creazione della tabella transfers
CREATE TABLE transfers (
    id INTEGER PRIMARY KEY,
    counterparty_name TEXT,
    counterparty_iban TEXT,
    counterparty_bic TEXT,
    amount_cents INTEGER,
    bank_account_id INTEGER,
    description TEXT,
    FOREIGN KEY (bank_account_id) REFERENCES bank_accounts(id)
);