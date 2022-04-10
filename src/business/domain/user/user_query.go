package user

const (
	addUser = `
	INSERT into user (
		name,
		age,
		gender,
		email,
		transaction_id)
	values (?,?,?,?,?)
	`

	getUser = `
	SELECT *
	FROM user
	WHERE email = ?
	`

	getUserTicket = `
	SELECT 
		transaction.code, 
		event.name,
		event.location, 
		event.date, 
		event.organizer 
	FROM user,transaction,event 
	WHERE 
		transaction.action = "accepted" AND
		user.transaction_id = transaction.id AND 
		transaction.event_id = event.id AND
		user.email = ?
	`
)
