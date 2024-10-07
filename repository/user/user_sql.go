package user

const (
	SelectUserByUsernameAndPassword = `
		SELECT 
			id
		FROM
			user
		WHERE
			username = ? AND
			password = ?
		LIMIT 1
	`
)
