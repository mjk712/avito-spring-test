package dto

type AddProductRequest struct {
	Type  string `json:"type"`
	PvzID string `json:"pvzId"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	DateTime    string `json:"dateTime"`
	Type        string `json:"type"`
	ReceptionID string `json:"receptionId"`
}
