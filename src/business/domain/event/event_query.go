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
)
