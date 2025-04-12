SELECT balance_cents >= ? AS has_funds
FROM bank_accounts
WHERE organization_name = ?;