package user

const (
	SelectUserByUsernameAndPassword = `
		SELECT 
			id
		FROM
			users
		WHERE
			username = ? AND
			password = ?
		LIMIT 1
	`
)
