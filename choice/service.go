package story

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	s := new(Service)
	s.db = db

	return s
}

func (s Service) Create(name, text string, sceneId, nextScene int) {
	_, err := s.db.Exec("INSERT INTO choices (name, text, scene_id, next_scene) "+
		"VALUES($1, $2, $3, $4)", name, text, sceneId, nextScene)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) GetById(id int) gin.H {
	var name, text string
	var sceneId, nextScene int
	err := s.db.QueryRow("SELECT name, text, scene_id, next_scene FROM choices WHERE id = $1", id).
		Scan(&name, &text, &sceneId, &nextScene)
	if err != nil {
		log.Fatal(err)
	}

	return gin.H{
		"name":       name,
		"text":       text,
		"scene_id":   sceneId,
		"next_scene": nextScene,
	}
}

func (s Service) GetAll() []gin.H {
	rows, err := s.db.Query("SELECT name, text, scene_id, next_scene FROM choices")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var name, text string
		var sceneId, nextScene int
		err = rows.Scan(&name, &text, &sceneId, &nextScene)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, gin.H{
			"name":       name,
			"text":       text,
			"scene_id":   sceneId,
			"next_scene": nextScene,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return all
}

func (s Service) Update(name, text string, sceneId, nextScene int) {
	_, err := s.db.Exec("UPDATE choices SET title = $2, text = $3, story_id = $4, next_scene = $5 "+
		"WHERE id = $1", id, name, text, sceneId, nextScene)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Delete(id int) {
	_, err := s.db.Exec("DELETE FROM choices WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}