package domain

import (
	"airflight/internal/sql_request"
	"log"
)

type counter struct {
	PassengerId     int
	count_passenger int
}

func countPassenger() map[string]int {
	passenger := sql_request.GetPassenger("", "")
	// var counterPassenger []counter
	count := make(map[string]int)
	for i := range passenger {
		passenger := passenger[i]["Name"].(string) + " " + passenger[i]["First Name"].(string)
		log.Print(passenger)
		count[passenger] += 1
	}

	return count
}

func RegularPassenger() {
	countPassenger()
}
