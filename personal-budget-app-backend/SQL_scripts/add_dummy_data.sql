USE personal_budget;

-- Inserting data into `users`
INSERT INTO users (email, first_name, last_name, password) VALUES
('john.doe@example.com', 'John', 'Doe', 'password123'),
('jane.smith@example.com', 'Jane', 'Smith', 'securepassword'),
('alice.jones@example.com', 'Alice', 'Jones', 'mypassword'),
('test@test.com', 'Test', 'User', 'testpassword');

-- Inserting data into `accounts`
INSERT INTO accounts (email, name, type, bank_name, starting_balance, balance) VALUES
('john.doe@example.com', 'John''s Checking', 'Checking', 'Bank A', 1000, 500),
('jane.smith@example.com', 'Jane''s Savings', 'Savings', 'Bank B', 2000, 1500),
('alice.jones@example.com', 'Alice''s Investment', 'Investment', 'Bank C', 3000, 2500),
('test@test.com', 'Test Account', 'Checking', 'Bank D', 500, 500);

-- Inserting data into `categories`
INSERT INTO categories (email, name) VALUES
('john.doe@example.com', 'Utilities'),
('jane.smith@example.com', 'Groceries'),
('alice.jones@example.com', 'Entertainment'),
('test@test.com', 'Auto');

-- Assuming accounts and categories have been populated and have these IDs
-- Adjust account_id and category_id as necessary based on actual IDs

-- Inserting data into `transactions`
INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES
(1, '2023-11-30 02:12:16', 'Utility Company', 10000, 'Electric bill', 1, 'john.doe@example.com'),
(2, '2024-02-04 02:12:16', 'Grocery Store', 20000, 'Weekly groceries', 2, 'jane.smith@example.com'),
(3, '2023-07-09 02:12:16', 'Movie Theater', 5000, 'Movie night', 3, 'alice.jones@example.com'),
(4, '2023-07-09 02:12:16', 'Auto Shop', 5000, 'Oil change', 4, 'test@test.com');

-- Inserting data into `goals`
INSERT INTO goals (email, name, amount, target_date, category_id, periodicity) VALUES
('john.doe@example.com', 'Save for Vacation', 200000, '2023-12-25', 1, 'onetime'),
('jane.smith@example.com', 'New Car Fund', 1000000, '2024-05-01', 2, 'monthly'),
('alice.jones@example.com', 'Gaming PC', 150000, '2023-11-15', 3, 'onetime'),
('test@test.com', 'Registration', 2000, '2024-04-01', 4, 'monthly');

-- Inserting data into `monthly_budgets`
INSERT INTO monthly_budgets (email, month, year, category_id, assigned, spent) VALUES
('john.doe@example.com', 10, 2022, 1, 500, 300),
('jane.smith@example.com', 11, 2022, 2, 600, 400),
('alice.jones@example.com', 12, 2022, 3, 700, 450),
('test@test.com', 4, 2024, 4, 10000, 5000);

-- Inserting data into `proj_income`
INSERT INTO proj_income (email, amount, month, year, date) VALUES
('john.doe@example.com', 3000, 10, 2022, '2024-01-11 02:12:16'),
('jane.smith@example.com', 4000, 11, 2022, '2023-08-24 02:12:16'),
('alice.jones@example.com', 5000, 12, 2022, '2023-05-07 02:12:16'),
('test@test.com', 200000, 4, 2024, '2023-05-07 02:12:16');
