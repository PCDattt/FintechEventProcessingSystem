ALTER TABLE transactions RENAME COLUMN account_id TO from_account_id;

ALTER TABLE transactions ALTER COLUMN from_account_id DROP NOT NULL;