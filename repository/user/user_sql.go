package user

const (
	SelectUserByUsername = `
		SELECT 
			id, password
		FROM
			user
		WHERE
			username = ?
		LIMIT 1
	`
)
