package dao

type Reception struct {
	ID       string `db:"id"`
	DateTime string `db:"date_time"`
	PvzID    string `db:"pvz_id"`
	Status   string `db:"status"` // in_progress или close
}
