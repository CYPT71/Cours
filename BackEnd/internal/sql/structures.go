package sql

import (
	"time"
)

type Employees struct {
	id              int
	aicrew          int
	ground          int
	social_security int
	name            string
	first_name      string
	adress          string
}

type Pilote struct {
	id       int
	license  time.Time
	among    time.Time
	staff_id int
}

type CabinCrew struct {
	id       int
	fonction string
	among    time.Time
	staff_id int
}

type Passenger struct {
	id         int
	ticket_id  int
	bank       string
	profession string
	name       string
	first_name string
	adress     string
}

type Fligth struct {
	id            int
	id_departures int
	ariaval       time.Time
	id_route      int
	id_device     int
}

type Tickets struct {
	id            int
	expire        time.Time
	price         int
	departures_id int
}

type Departus struct {
	id          int
	id_flight   int
	date        time.Time
	pilote      int
	copilote    int
	aircrew     string
	free_places int
	occupied    int
}

type Route struct {
	id      int
	origin  string
	arrival string
}

type Device struct {
	id       int
	capacity int
	types    string
}
