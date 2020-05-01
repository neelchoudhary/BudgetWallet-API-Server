BEGIN;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL
);

CREATE TABLE items (
	id SERIAL PRIMARY KEY,
	plaid_item_id TEXT NOT NULL,
	plaid_access_token TEXT NOT NULL,
    user_id INT NOT NULL,
    plaid_institution_id TEXT,
    institution_name TEXT,
    institution_color TEXT,
    institution_logo TEXT,
	error_code TEXT,
	error_dev_msg TEXT,
	error_user_msg TEXT
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    item_id INT NOT NULL,
	plaid_account_id TEXT NOT NULL,
	current_balance FLOAT8,
	available_balance FLOAT8,
	account_name TEXT,
	official_name TEXT,
	account_type TEXT,
	account_subtype TEXT,
    account_mask TEXT,
    selected BOOLEAN
);

CREATE TABLE plaid_categories (
	plaid_category_id TEXT PRIMARY KEY,
	category1 TEXT,
	category2 TEXT,
	category3 TEXT
);

CREATE TABLE categories (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	grouping TEXT NOT NULL
);

CREATE TABLE category_mapping (
	plaid_category_id TEXT PRIMARY KEY,
	category_id INT NOT NULL
);

CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	item_id INT NOT NULL,
	account_id INT NOT NULL,
	category_id INT NOT NULL,
	daily_account_snapshot_id INT DEFAULT -1,
	monthly_account_snapshot_id INT DEFAULT -1,
	plaid_category_id TEXT NOT NULL,
	plaid_transaction_id TEXT NOT NULL,
	name TEXT,
	amount FLOAT8,
  	date TEXT,
    pending BOOLEAN
);

CREATE TABLE daily_account_snapshots (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	item_id INT NOT NULL,
	account_id INT NOT NULL,
  	date TEXT,
  	balance FLOAT8,
  	cash_out FLOAT8,
  	cash_in FLOAT8
);

CREATE TABLE monthly_account_snapshots (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	item_id INT NOT NULL,
	account_id INT NOT NULL,
	date TEXT,
  	balance FLOAT8,
  	cash_out FLOAT8,
  	cash_in FLOAT8
);

COMMIT;