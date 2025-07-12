ALTER TABLE transactions RENAME COLUMN from_account_id TO account_id;

ALTER TABLE transactions ALTER COLUMN account_id SET NOT NULL;