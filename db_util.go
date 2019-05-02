package main

import (
	"database/sql"
	"fmt"
	"time"
	config "config"
	_ "github.com/go-sql-driver/mysql"
)

// PersistDSR2016 record in mysql instance.
func PersistDSR2016(entity Entity) bool {
	config := 
	db := getConnectionHandler()
	stmtIns, err := db.Prepare("INSERT INTO dsr_2016 VALUES( ?, ?, ?, ?, ?, ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	t := time.Now()
	fmt.Println(t)
	_, err = stmtIns.Exec(entity.ID, "REINFORCED CONCRETE CEMENT", entity.Description, entity.Unit, entity.Rate, &t, &t)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func getConnectionHandler() (db *sql.DB) {
	dbDriver := "mysql"
	dbSchemaUser := "umesh_app"
	dbPassword := "umesh_app_p"
	dbUsername := "umesh_app"
	db, err := sql.Open(dbDriver, dbSchemaUser+":"+dbPassword+"@/"+dbUsername+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
