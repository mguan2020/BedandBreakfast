package repository

import (
	"time"

	"github.com/tsawler/bookings-app/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	GetRoomByID(id int) (models.Room, error)
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	GetReservationByID(id int) (models.Reservation, error)
}
