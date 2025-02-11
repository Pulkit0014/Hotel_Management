# Hotel Room Booking System

## Overview
This is a simple **Hotel Room Booking System** built using **Golang** with **Gorilla Sessions** for session management. It allows users to book rooms, view bookings, modify/cancel bookings, and list all guests currently staying in the hotel.

## Features
- **Book a Room:** Users can book a room by providing their details.
- **View Booking:** Retrieve booking details using the guest's email.
- **View All Guests:** List all current guests with their room numbers.
- **Cancel Booking:** Cancel an existing booking.
- **Modify Booking:** Change check-in/check-out dates.
- **Session-Based Storage:** Uses `gorilla/sessions` to store booking data in session.

## Technologies Used
- **Golang**
- **Gorilla Mux** (Router)
- **Gorilla Sessions** (Session management)

---

## Installation & Setup
### **1. Prerequisites**
- Install Go (latest version)
- Install dependencies using `go mod`.

### **2. Clone the Repository**
```sh
 git clone https://github.com/yourusername/hotel-booking-system.git
 cd hotel-booking-system
```

### **3. Install Dependencies**
```sh
go mod tidy
```

### **4. Run the Server**
```sh
go run main.go
```
The server will start at `http://localhost:8080`

---

## API Endpoints

### **1️⃣ Book a Room**
**Endpoint:** `POST /book`

#### **Request JSON:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "contact": "+1234567890",
  "check_in": "2025-02-15T14:00:00Z",
  "check_out": "2025-02-20T11:00:00Z"
}
```

#### **Response JSON:**
```json
{
  "room_number": 101,
  "name": "John Doe",
  "email": "johndoe@example.com",
  "check_in": "2025-02-15T14:00:00Z",
  "check_out": "2025-02-20T11:00:00Z"
}
```

---

### **2️⃣ View Booking Details**
**Endpoint:** `GET /viewBooking?email=johndoe@example.com`

#### **Response JSON:**
```json
{
  "room_number": 101,
  "name": "John Doe",
  "email": "johndoe@example.com",
  "check_in": "2025-02-15T14:00:00Z",
  "check_out": "2025-02-20T11:00:00Z"
}
```

---

### **3️⃣ View All Guests**
**Endpoint:** `GET /guests`

#### **Response JSON:**
```json
[
  {"room_number": 101, "name": "John Doe"},
  {"room_number": 102, "name": "Jane Smith"}
]
```

---

### **4️⃣ Cancel Booking**
**Endpoint:** `DELETE /cancel`

#### **Request JSON:**
```json
{
  "email": "johndoe@example.com"
}
```

#### **Response JSON:**
```json
{
  "message": "Booking cancelled successfully"
}
```

---

### **5️⃣ Modify Booking**
**Endpoint:** `PUT /modify`

#### **Request JSON:**
```json
{
  "email": "johndoe@example.com",
  "check_in": "2025-02-16T14:00:00Z",
  "check_out": "2025-02-22T11:00:00Z"
}
```

#### **Response JSON:**
```json
{
  "message": "Booking updated successfully"
}
```

---

## Session Management
- The system stores booking details in a **session** using `gorilla/sessions`.
- **Fix for session storage:** Register custom structs with `gob`.

### **Registering `Booking` Struct**
Modify `main.go`:
```go
import (
    "encoding/gob"
    "your_project/structs"
)

func init() {
    gob.Register(structs.Booking{})
}
```

---

## Troubleshooting
### **🔹 "Booking not found" Issue**
- Ensure you pass the correct `email` in the request.
- Check if session storage is working correctly.
- Restart the server after making changes.

### **🔹 "securecookie: error - gob: type not registered" Issue**
- Ensure `gob.Register(structs.Booking{})` is called **before** using sessions.
- Restart the server after adding the `gob.Register()` fix.

---

