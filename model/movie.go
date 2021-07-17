package model

import "log"

type Movie struct {
	Model
	MovieNo string `json:"movie_no"`
	Title string `json:"title"`
}

func CreateMovie(data *Movie)  {
	tx := DB.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		log.Fatalf("create fail: %v", err)
	}
	if err := tx.Exec("update fetch_count set count = count + 1 where id = 1").Error; err != nil {
		tx.Rollback()
		log.Fatalf("update fail: %v", err)
	}
	tx.Commit()
}
