package http

import "go-boilerplate/src/bookings/domain/bookings"

func TransformToHTTPBooking(bookingObj *bookings.Booking) Booking {
	return Booking{
		Id:       bookingObj.ID,
		BookId:   bookingObj.BookID,
		DateFrom: bookingObj.DateFrom,
		DateTo:   bookingObj.DateTo,
	}
}

func TransformToHTTPBookings(bookingsObj []bookings.Booking) []Booking {
	var bookings []Booking = make([]Booking, 0)
	for _, booking := range bookingsObj {
		bookings = append(bookings, TransformToHTTPBooking(&booking))
	}
	return bookings
}
