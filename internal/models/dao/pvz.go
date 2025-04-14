package dao

type Pvz struct {
	ID               string `db:"id"`
	RegistrationDate string `db:"registration_date"`
	City             string `db:"city"` // Москва, СПб, Казань
}
