use personal_budget;;

DROP TRIGGER IF EXISTS account_creation_trigger;;

CREATE TRIGGER account_creation_trigger
AFTER INSERT ON accounts
FOR EACH ROW
BEGIN
    INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email)
    VALUES (NEW.id, NEW.starting_date, 'Starting Balance', NEW.starting_balance, 'Starting Balance', NULL, NEW.email);
END;;