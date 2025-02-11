package structs

import "time"

type Booking struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Contact    string    `json:"contact"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	RoomNumber int       `json:"room_number,omitempty"`
}
