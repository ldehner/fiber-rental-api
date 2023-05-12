package models

type SearchProfile struct {
	UserId  string  `json:"UserId"`
	Budget  float32 `json:"Budget"`
	City    string  `json:"City"`
	Country string  `json:"Country"`
	Radius  float32 `json:"Radius"`
	Rooms   uint8   `json:"Rooms"`
	Size    uint16  `json:"Size"`
	Street  string  `json:"Street"`
	Type    uint8   `json:"Type"`
	Zipcode string  `json:"Zipcode"`
}
