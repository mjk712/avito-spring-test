package dto

type CreatePvzRequest struct {
	City string `json:"city"`
}

type PvzResponse struct {
	ID               string `json:"id"`
	RegistrationDate string `json:"registrationDate"`
	City             string `json:"city"`
}

type ProductShort struct {
	ID       string `json:"id"`
	DateTime string `json:"dateTime"`
	Type     string `json:"type"`
}

type ReceptionShort struct {
	ID       string         `json:"id"`
	DateTime string         `json:"dateTime"`
	Status   string         `json:"status"`
	Products []ProductShort `json:"products"`
}

type PVZWithReceptions struct {
	PVZ        PvzResponse      `json:"pvz"`
	Receptions []ReceptionShort `json:"receptions"`
}
