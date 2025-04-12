UPDATE bank_accounts
SET balance_cents = balance_cents - ?
WHERE id = ?;
