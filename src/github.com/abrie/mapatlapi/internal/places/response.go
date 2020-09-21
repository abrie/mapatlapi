package places

type Response []Record

type Record struct {
	NAME_OF_PLOI string `json:"NAME_OF_PLOI,omitempty"`
	ADDRESS      string `json:"ADDRESS,omitempty"`
	CITY         string `json:"CITY,omitempty"`
	PHONE        string `json:"PHONE,omitempty"`
	WEBSITE      string `json:"WEBSITE,omitempty"`
	EMAIL        string `json:"EMAIL,omitempty"`
	DISTANCE     string `json:"DISTANCE,omitempty"`
}
