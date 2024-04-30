USE personal_budget;

DROP TABLE IF EXISTS `transactions`;
DROP TABLE IF EXISTS `goals`;
DROP TABLE IF EXISTS `monthly_budgets`;
DROP TABLE IF EXISTS `proj_income`;
DROP TABLE IF EXISTS `accounts`;
DROP TABLE IF EXISTS `categories`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE users (
    email VARCHAR(100) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);


CREATE TABLE accounts (
    id int AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100),
    name VARCHAR(100),
    type VARCHAR(100),
    bank_name VARCHAR(100),
    starting_balance int,
    balance int,
    FOREIGN KEY (email) REFERENCES users(email)
);

CREATE TABLE categories (
    id int AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100),
    name VARCHAR(100),
    FOREIGN KEY (email) REFERENCES users(email)
);

CREATE TABLE transactions (
    id int AUTO_INCREMENT PRIMARY KEY,
    account_id int,
    date date,
    payee VARCHAR(100),
    amount int,
    memo VARCHAR(100),
    category_id int,
    email VARCHAR(100),
    FOREIGN KEY (email) REFERENCES users(email),
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE goals (
    id int AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100),
    name VARCHAR(100),
    amount int,
    target_date datetime,
    category_id int,
    periodicity VARCHAR(100),
    FOREIGN KEY (email) REFERENCES users(email),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);



CREATE TABLE monthly_budgets(
    id int AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100),
    month int NOT NULL,
    year int NOT NULL,
    category_id int NOT NULL,
    assigned int NOT NULL,
    spent int NOT NULL,
    FOREIGN KEY (email) REFERENCES users(email),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    CONSTRAINT unique_monthly_budget UNIQUE (email, month, year, category_id)
);

CREATE TABLE proj_income (
    id int AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100),
    amount int NOT NULL,
    month int NOT NULL,
    year int NOT NULL,
    date datetime,
    FOREIGN KEY (email) REFERENCES users(email)
);