package database

import (
	"log"
	"time"

	"github.com/FoldFunc/LAF/structs"
)

func ClaimLostAndFoundDatabase(r structs.ReqestClaimLostAndFound) error {
	log.Println("ClaimLostAndFoundDatabase function called")
	currentDate := time.Now().Format("2006-01-02")
	_, err := DB.Exec(`UPDATE laf SET date_claimed = ?, whofound = ? WHERE name = ?`, currentDate, r.ClaimingName, r.Name)
	if err != nil {
		log.Println("Error while updating the database: ", err)
		return err
	}

	return nil
}
