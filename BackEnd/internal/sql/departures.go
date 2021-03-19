package sql

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	id          int       `json:"id"`
	id_flight   int       `json:"id_fligth"`
	date        time.Time `json:"date"`
	pilote      int       `json:"pilote"`
	copilote    int       `json:"copilote"`
	aircrew     string    `json:"aircrew"`
	free_places int       `json:"free_places"`
	occupied    int       `json:"occupied"`
	ticket_id   int       `json:"ticket_id"`
}

func AddDepartures(values string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO Departus (id_flight, 	date_departure) VALUES ", values)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetDepartures(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM Departus "
	if filter != "" {
		query += " WHERE `id` IN (" + filter + ")"
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var departures [][]string
	var tag Tag
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.id_flight, &tag.date, &tag.aircrew, &tag.copilote, &tag.pilote, &tag.free_places,
			&tag.occupied, &tag.ticket_id)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.id_flight), tag.date.Format(time.UnixDate),
			tag.aircrew, strconv.Itoa(tag.copilote), strconv.Itoa(tag.pilote), strconv.Itoa(tag.free_places),
			strconv.Itoa(tag.occupied), strconv.Itoa(tag.ticket_id)}
		departures = append(departures, to_inject)
	}

	return departures

}
