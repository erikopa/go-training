package routes

type Route struct {
	ID string `json:"id,omitempty"`
	Modality int `json:"modalityId,omitempty"`
	Threshold int `json:"threshold,omitempty"`
	CreateAt string `json:"create_at,omitempty"`
	UpdateAt string `json: "update_at,omitempty"`
}

type Routes []Route
