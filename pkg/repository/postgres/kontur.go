package postgres

import (
	"fmt"
	"math/rand"

	"github.com/Murolando/hakaton_geo/ent"
	"github.com/jmoiron/sqlx"
)

type KonturPostgres struct {
	db *sqlx.DB
}

func NewKonturPostgres(db *sqlx.DB) *KonturPostgres {
	return &KonturPostgres{
		db: db,
	}
}

func (r *KonturPostgres) StartKonturGame(n int) ([]*ent.KonturResponse, error) {
	var list []*ent.KonturResponse
	for i := 0; i < n; i++ {
		query := fmt.Sprintf(`
		SELECT * 
		FROM "%s"
		ORDER BY RANDOM() LIMIT 4
		`, konturTable)
		rows, err := r.db.Query(query)
		if err != nil {
			return nil, err
		}
		var kontur ent.KonturResponse
		if rows.Next() {
			var name string
			if err := rows.Scan(&kontur.Id, &name, &kontur.ImageSrc); err != nil {
				return nil, err
			}
			kontur.Correct = name
			kontur.Name[0] = name
		}
		i := 1
		for rows.Next() {
			var name string
			var id string
			var img_src string
			if err := rows.Scan(&id, &name, &img_src); err != nil {
				return nil, err
			}
			fmt.Println(id, img_src)
			kontur.Name[i] = name
			i++
		}

		d := len(kontur.Name)
		for i := d - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			kontur.Name[i], kontur.Name[j] = kontur.Name[j], kontur.Name[i]
		}
		list = append(list, &kontur)
	}

	return list, nil
}

func (r *KonturPostgres) ProcessKonturGame(params *ent.ProcessRequest, userId int) (*ent.ProcessResponse, error) {
	winRate := 0
	for _, v := range params.Answers {
		if v {
			winRate += 1
		}
	}
	query := fmt.Sprintf(`
		SELECT * 
		FROM "%s"
		WHERE user_id = "%s" AND lesson_id = "%s"
		`, konturTable, userId, params.LessonId)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)
	return nil, nil

}
