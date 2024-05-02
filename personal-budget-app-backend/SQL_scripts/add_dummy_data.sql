USE personal_budget;

-- Inserting data into `users`
INSERT INTO users (email, first_name, last_name, password) VALUES
('test@test.com', 'Test', 'User', '$2b$10$gTqeWbIpmnIQaososAePU.ZK1DBeAbIJyrR.z4uQn6FHQ123u2Ru.');

-- Inserting data into `accounts`
INSERT INTO accounts (email, name, type, bank_name, starting_balance, balance) VALUES
('test@test.com', 'Test Account', 'Checking', 'Bank D', 500, 500);

-- Inserting data into `categories`
INSERT INTO categories (email, name) VALUES
('test@test.com', 'Utilities'),
('test@test.com', 'Groceries'),
('test@test.com', 'Entertainment'),
('test@test.com', 'Auto');
('test@test.com', 'Vacation');

-- Inserting data into `goals`
INSERT INTO goals (email, name, amount, target_date, category_id, periodicity) VALUES
('test@test.com', 'Save for Vacation', 200000, '2024-12-25', 1, 'onetime'),
('test@test.com', 'New Car Fund', 1000000, '2024-05-01', 2, 'monthly'),
('test@test.com', 'New Car Fund', 1000000, '2024-05-01', 2, 'monthly'),
('test@test.com', 'Gaming PC', 150000, '2024-11-15', 3, 'onetime'),
('test@test.com', 'Registration', 20000, '2024-04-01', 4, 'onetime');

-- Inserting data into `monthly_budgets`: this needs to be created before transactions or the proc 
INSERT INTO monthly_budgets (email, month, year, category_id, assigned, spent) VALUES
('test@test.com', 4, 2024, 1, 500000, 0),
('test@test.com', 4, 2024, 2, 60000, 0),
('test@test.com', 4, 2024, 3, 7000, 0),
('test@test.com', 4, 2024, 4, 10000, 0);

-- Inserting data into `transactions`
INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2024-4-30 02:12:16', 'Utility Company', -10000, 'Electric bill', 1, 'test@test.com');
CALL update_account_balance(1, 4, 2024, -10000);

INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2024-4-04 02:12:16', 'Food Lion', -20000, 'Weekly groceries', 2, 'test@test.com');
CALL update_account_balance(1, 4, 2024, -20000);
CALL update_monthly_budget_spent(1, 4, 2024, -20000);

INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2023-07-09 02:12:16', 'Cinemark', -50000, 'Movie night', 3, 'test@test.com');

INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2023-07-09 02:12:16', 'Scott Clark Honda', -50000, 'Oil change', 4, 'test@test.com');

-- Inserting data into `proj_income`
INSERT INTO proj_income (email, amount, month, year, date) VALUES
('test@test.com', 3000, 10, 2022, '2024-01-11 02:12:16'),
('test@test.com', 4000, 11, 2022, '2023-08-24 02:12:16'),
('test@test.com', 5000, 12, 2022, '2023-05-07 02:12:16'),
('test@test.com', 200000, 4, 2024, '2023-05-07 02:12:16');
