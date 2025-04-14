package dao

type Product struct {
	ID          string `db:"id"`
	DateTime    string `db:"date_time"`
	Type        string `db:"type"` // электроника, одежда, обувь
	ReceptionID string `db:"reception_id"`
}
