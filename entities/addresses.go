package entities

// Address struct
type Address struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}
