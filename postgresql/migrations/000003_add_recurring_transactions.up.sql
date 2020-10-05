BEGIN;

CREATE TABLE recurring_transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name TEXT NOT NULL,
    category_id INT NOT NULL,
    recurring_count INT NOT NULL,
    recurring_score INT NOT NULL,
    is_recurring TEXT NOT NULL,
    recurring_plaid_ids TEXT[] NOT NULL
);

COMMIT;