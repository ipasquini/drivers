package data

type Driver struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ID        int     `json:"driver_id"`
}
