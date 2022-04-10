package event

const (
	findAll = `
	SELECT *
	FROM
		event`

	getEvent = `
	SELECT *
	FROM
		event
	WHERE id = ?`

	// addTransaction = `
	// INSERT into transaction (
	// 	event_id,
	// 	quantity,
	// 	total_price,
	// 	payment_method,
	// 	action,
	// 	code)
	// values (?,?,?,?,"waiting",?)
	// `
	// addUser = `
	// INSERT into user (
	// 	name,
	// 	age,
	// 	gender,
	// 	email,
	// 	transaction_id)
	// values (?,?,?,?,?)
	// `
)
