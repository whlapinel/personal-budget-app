DROP PROCEDURE IF EXISTS update_account_balance;;
DROP PROCEDURE IF EXISTS update_monthly_budget_spent;;

-- should be used with any transaction insert, update, or delete
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
END;;


-- should be used when transaction has a category_id
CREATE PROCEDURE update_monthly_budget_spent
(
	IN category_id INT,
    IN month INT,
    IN year INT,
    IN amount INT
)
BEGIN
DECLARE new_spent INT DEFAULT 0;
DECLARE old_spent INT DEFAULT 0;
SELECT spent INTO old_spent FROM monthly_budgets m WHERE id = category_id AND m.month = month AND m.year = year;
SET new_spent = old_spent + (amount * -1);
UPDATE monthly_budgets m SET spent = new_spent WHERE id = category_id AND m.month = month AND m.year = year;
END;;


