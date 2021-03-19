package sql

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Departus struct {
	id          int
	id_flight   int
	date        time.Time
	pilote      int
	copilote    int
	aircrew     string
	free_places int
	occupied    int
	ticket_id   int
}

func AddDepartures(id_flight int, date time.Time, pilote int, copilote int, aircrew string, free_places int, occupied int, ticket_id int) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Departus`(`id_fligth`, `date`, `pilote`, `copilote`, `aircrew`, `free_places`, `occupied`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		id_flight, date, pilote, copilote, aircrew, free_places, occupied, ticket_id)

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

	var return_val [][]string
	var tag Departus
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.id_flight, &tag.date, &tag.aircrew, &tag.copilote, &tag.pilote, &tag.free_places,
			&tag.occupied, &tag.ticket_id)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.id_flight), tag.date.Format(time.UnixDate),
			tag.aircrew, strconv.Itoa(tag.copilote), strconv.Itoa(tag.pilote), strconv.Itoa(tag.free_places),
			strconv.Itoa(tag.occupied), strconv.Itoa(tag.ticket_id)}
		return_val = append(return_val, to_inject)
	}

	return return_val

}
