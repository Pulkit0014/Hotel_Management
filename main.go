package main

import (
	"fmt"
	"hotel_managgement/routers"
	"net/http"
)

func main() {
	http.HandleFunc("/book", routers.BookRoom)
	http.HandleFunc("/view", routers.ViewBooking)
	http.HandleFunc("/guests", routers.ViewAllGuests)
	http.HandleFunc("/cancel", routers.CancelBooking)
	http.HandleFunc("/modify", routers.ModifyBooking)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
