package database

import (
	"log"

	"github.com/FoldFunc/LAF/structs"
)

func ShowAllLostAndFoundDatabase() ([]structs.LostAndFound, error) {
	log.Println("ShowAllLostAndFoundDatabase function called")

	rows, err := DB.Query(`SELECT name, whofound, "where", notes, date_added, date_claimed FROM laf`)
	if err != nil {
		log.Println("Error while fetching from database:", err)
		return nil, err
	}
	defer rows.Close()

	var results []structs.LostAndFound

	for rows.Next() {
		var item structs.LostAndFound
		err := rows.Scan(&item.Name, &item.WhoFound, &item.Where, &item.Notes, &item.DateAdded, &item.DateClaimed)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		results = append(results, item)
	}

	return results, nil
}
