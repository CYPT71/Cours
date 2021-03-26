package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Employees struct {
	Id             int    `json:"id"`
	Salary         int    `json:"salary"`
	SocialSecurity int    `json:"social_security"`
	Name           string `json:"name"`
	FirstName      string `json:"first_name"`
	Address        string `json:"address"`
}

func AddEmployees(Aicrew int, Ground int, Social_security int, Name string, First_name string, Address string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `employees`(`aircrew`, `ground`, `social_security`, `name`, `first_name`, `address`) VALUES (?, ?, ?, ?, ?, ?)",
		Aicrew, Ground, Social_security, Name, First_name, Address)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetEmployees(selector string, filter string) []map[string]interface{} {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "*"
	}
	query += " FROM `employees` "
	if filter != "" {
		query += "WHERE " + filter
	}

	query += ";"

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var return_val []map[string]interface{}

	for results.Next() {
		var tag Employees

		err = results.Scan(&tag.Id, &tag.Salary, &tag.SocialSecurity, &tag.Name, &tag.FirstName, &tag.Address)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"id":              tag.Id,
			"name":            tag.Name,
			"first name":      tag.FirstName,
			"social security": tag.SocialSecurity,
		})
		//log.Print(tag.Id, tag.Salary, tag.SocialSecurity, tag.Name, tag.FirstName, tag.Address)
	}

	return return_val

}

func UpdateEmployees(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	query := "UPDATE `employees` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteEmployees(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	query := "DELETE FROM `employees` WHERE " + condition

	db.Query(query)

}
