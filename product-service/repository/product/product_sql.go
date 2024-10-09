package product

const (
	QueryAddStockProduct = `
		UPDATE 
			product
		SET
			stock = stock + ? 
		WHERE id = ?
	`
)
