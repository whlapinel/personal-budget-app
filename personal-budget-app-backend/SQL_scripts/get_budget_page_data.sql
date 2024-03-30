
-- Get the budget page data for the user

	SELECT c.id, c.name,
    SUM(COALESCE(mb.assigned, 0)) AS assigned,
    SUM(COALESCE(mb.spent, 0)) AS spent,
    SUM(COALESCE(CASE WHEN mb.assigned - mb.spent > 0 THEN mb.assigned - mb.spent ELSE 0 END, 0)) AS available,
    SUM(COALESCE(g.amount, 0)) AS total_goal_amount
	FROM categories c
	LEFT JOIN monthly_budgets mb ON c.id = mb.category_id 
    AND (mb.year < ? OR (mb.year = ? AND mb.month < ?))
	LEFT JOIN goals g ON c.id = g.category_id
	WHERE c.email = ?
	GROUP BY c.id, c.name;
