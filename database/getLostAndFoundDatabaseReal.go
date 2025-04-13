// GetLostAndFoundDatabase fetches rows from the database based on the provided query information
package database

import (
	"log"

	"github.com/FoldFunc/LAF/structs"
)

func GetLostAndFoundDatabase(queryInfo structs.ReqestGetAndLostHandler) (error, []structs.LostAndFound) {
	log.Println("GetLostAndFoundDatabase function called")

	rows, err := DB.Query(`SELECT name, whoFound, "where", notes, date_added FROM laf WHERE name LIKE ?`, "%"+queryInfo.Name+"%")
	if err != nil {
		log.Println("Error while fetching from database: ", err)
		return err, nil
	}
	defer rows.Close()

	var results []structs.LostAndFound
	for rows.Next() {
		var item structs.LostAndFound
		if err := rows.Scan(&item.Name, &item.WhoFound, &item.Where, &item.Notes, &item.DateAdded); err != nil {
			log.Println("Error scanning row: ", err)
			continue
		}
		results = append(results, item)
	}

	return nil, results
}
