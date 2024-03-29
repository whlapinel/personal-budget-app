DROP PROCEDURE IF EXISTS update_account_balance;
DROP PROCEDURE IF EXISTS update_category_balance;

DELIMITER //
CREATE PROCEDURE update_account_balance
(
    IN account_id INT, IN amount INT
) 
BEGIN 
DECLARE new_balance INT DEFAULT 0;
DECLARE old_balance INT DEFAULT 0;
SELECT balance INTO old_balance
FROM accounts
WHERE id = account_id;
SET new_balance = old_balance + amount;
UPDATE accounts
SET balance = new_balance
WHERE id = account_id;
END//
DELIMITER;

DELIMITER //
CREATE PROCEDURE update_category_balance
(
	IN category_id INT,
    IN amount INT
)
BEGIN
DECLARE new_balance INT DEFAULT 0;
DECLARE old_balance INT DEFAULT 0;
SELECT balance INTO old_balance FROM categories WHERE id = category_id;
SET new_balance = old_balance + amount;
UPDATE categories SET balance = new_balance WHERE id = category_id;
END//
DELIMITER;