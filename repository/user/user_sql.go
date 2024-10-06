package user

const(
	SelectUserByUsernameAndPassword = `
		SELECT 
			isExists(id)
		FROM
			user
		WHERE
			username = ? AND
			password = ?
		LIMIT 1
	`
)