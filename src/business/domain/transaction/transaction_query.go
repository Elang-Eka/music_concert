package transaction

const (
	addTransaction = `
	INSERT into transaction (
		event_id,
		quantity,
		total_price,
		payment_method,
		action,
		code) 
	values (?,?,?,?,"waiting",?)`

	updateTransaction = `
	UPDATE transaction SET action = ?
	WHERE event_id =? AND total_price = ? AND payment_method = ? AND action = "waiting" AND code = ?`

	getTransaction = `
	SELECT *
	FROM 
		transaction
	WHERE 
		code = ?`
)
