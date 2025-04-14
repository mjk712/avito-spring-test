package dto

type CreateReceptionRequest struct {
	PvzID string `json:"pvzId"`
}

type ReceptionResponse struct {
	ID       string `json:"id"`
	DateTime string `json:"dateTime"`
	PvzID    string `json:"pvzId"`
	Status   string `json:"status"`
}
