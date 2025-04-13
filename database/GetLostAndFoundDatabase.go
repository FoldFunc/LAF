package database

import (
	"log"

	"github.com/FoldFunc/LAF/structs"
)

func GetLostAndFoundDatabase(queryInfo structs.ReqestGetAndLostHandler) error {
	log.Println("GetLostAndFoundDatabase function called")
	_, err := DB.Exec(`INSERT INTO laf (name, whofound, "where", notes) VALUES (?, ?, ?, ?)`, queryInfo.Name, queryInfo.WhoFound, queryInfo.Where, queryInfo.Notes)
	if err != nil {
		return err
	}
	return nil
}
