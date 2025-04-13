package database

import (
	"github.com/FoldFunc/LAF/structs"
	"log"
	"time"
)

func AddLostAndFoundDatabase(queryInfo structs.ReqestAddAndLostHandler) error {
	log.Println("AddLostAndFoundDatabase function called")

	// Get the current date in YYYY-MM-DD format
	currentDate := time.Now().Format("2006-01-02")

	// Insert the data along with the current date
	_, err := DB.Exec(`INSERT INTO laf (name, whofound, "where", notes, date_added, date_claimed) VALUES (?, ?, ?, ?, ?, ?)`,
		queryInfo.Name, queryInfo.WhoFound, queryInfo.Where, queryInfo.Notes, currentDate, " ")
	if err != nil {
		log.Printf("Error while inserting data: %v\n", err) // log the error
		return err
	}
	return nil
}
