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

-- Assuming accounts and categories have been populated and have these IDs
-- Adjust account_id and category_id as necessary based on actual IDs

-- Inserting data into `transactions`
INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2024-4-30 02:12:16', 'Utility Company', 10000, 'Electric bill', 1, 'test@test.com'),
(1, '2024-4-04 02:12:16', 'Food Lion', 20000, 'Weekly groceries', 2, 'test@test.com'),
(1, '2023-07-09 02:12:16', 'Cinemark', 50000, 'Movie night', 3, 'test@test.com'),
(1, '2023-07-09 02:12:16', 'Scott Clark Honda', 50000, 'Oil change', 4, 'test@test.com');

-- Inserting data into `goals`
INSERT INTO goals (email, name, amount, target_date, category_id, periodicity) VALUES
('test@test.com', 'Save for Vacation', 200000, '2023-12-25', 1, 'onetime'),
('test@test.com', 'New Car Fund', 1000000, '2024-05-01', 2, 'monthly'),
('test@test.com', 'Gaming PC', 150000, '2023-11-15', 3, 'onetime'),
('test@test.com', 'Registration', 20000, '2024-04-01', 4, 'monthly');

-- Inserting data into `monthly_budgets`
INSERT INTO monthly_budgets (email, month, year, category_id, assigned, spent) VALUES
('test@test.com', 4, 2024, 1, 5000, 10000),
('test@test.com', 4, 2024, 2, 6000, 20000),
('test@test.com', 4, 2024, 3, 7000, 50000),
('test@test.com', 4, 2024, 4, 10000, 50000);

-- Inserting data into `proj_income`
INSERT INTO proj_income (email, amount, month, year, date) VALUES
('test@test.com', 3000, 10, 2022, '2024-01-11 02:12:16'),
('test@test.com', 4000, 11, 2022, '2023-08-24 02:12:16'),
('test@test.com', 5000, 12, 2022, '2023-05-07 02:12:16'),
('test@test.com', 200000, 4, 2024, '2023-05-07 02:12:16');
