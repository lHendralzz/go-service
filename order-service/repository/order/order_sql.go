package order

const (
	InsertIntoOrder = `
		INSERT INTO order
			(user_id,status)
		VALUES
			(?,?)
	`
)
