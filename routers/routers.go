package routers

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"hotel_managgement/structs"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

var (
	store     = sessions.NewCookieStore([]byte("super-secret-key"))
	bookings  = make(map[string]structs.Booking)
	rooms     = make(map[int]bool)
	roomCount = 10
	mu        sync.Mutex
)

func init() {
	gob.Register(structs.Booking{})
}

func BookRoom(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "hotel_management_session")

	var req structs.Booking
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	roomNumber := -1
	for i := 1; i <= roomCount; i++ {
		if !rooms[i] {
			roomNumber = i
			rooms[i] = true
			break
		}
	}

	if roomNumber == -1 {
		http.Error(w, "No available rooms", http.StatusConflict)
		return
	}

	req.RoomNumber = roomNumber

	bookings[req.Email] = req

	session.Values[req.Email] = req
	fmt.Println("✅ Session before saving:", session.Values)

	err := session.Save(r, w)
	if err != nil {
		fmt.Println("Error saving session:", err)
	}

	fmt.Println("✅ Session after saving:", session.Values)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func ViewBooking(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "hotel_management_session")

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	fmt.Println("🔍 Current session values:", session.Values)

	bookingData, found := session.Values[email]
	if !found {
		fmt.Println("Booking not found in session for email:", email)
		http.Error(w, `{"error": "Booking not found"}`, http.StatusNotFound)
		return
	}

	booking, ok := bookingData.(structs.Booking)
	if !ok {
		fmt.Println("Data corruption in session for email:", email)
		http.Error(w, `{"error": "Data corruption in session"}`, http.StatusInternalServerError)
		return
	}

	fmt.Println("✅ Booking retrieved from session:", booking)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func ViewAllGuests(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var guests []structs.Booking
	for _, b := range bookings {
		guests = append(guests, b)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guests)
}
func CancelBooking(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "hotel_management_session")
	email := r.URL.Query().Get("email")
	mu.Lock()
	defer mu.Unlock()
	if booking, found := bookings[email]; found {
		delete(bookings, email)
		delete(rooms, booking.RoomNumber)
		delete(session.Values, email)
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Booking not found", http.StatusNotFound)
}
func ModifyBooking(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "hotel_management_session")
	email := r.URL.Query().Get("email")
	var req structs.Booking
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, found := bookings[email]; found {
		req.RoomNumber = bookings[email].RoomNumber
		bookings[email] = req
		session.Values[email] = req
		session.Save(r, w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(req)
		return
	}
	http.Error(w, "Booking not found", http.StatusNotFound)
}
