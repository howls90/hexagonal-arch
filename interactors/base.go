package interactors

import (
	mysql "hexagonal/dataSource/mysql"
	"hexagonal/repositories"
)

var (
	UserRepo   repositories.UserRepo   = mysql.UserMsql{}
	TripRepo   repositories.TripRepo   = mysql.TripMsql{}
	TicketRepo repositories.TicketRepo = mysql.TicketMsql{}
)
