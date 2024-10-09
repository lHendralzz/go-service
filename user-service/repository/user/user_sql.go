package user

const (
	SelectUserByEmail = `
		SELECT 
			id, password
		FROM
			user
		WHERE
			email = ?
		LIMIT 1
	`
)
