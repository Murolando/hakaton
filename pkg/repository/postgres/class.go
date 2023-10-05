package postgres

import (
	"fmt"

	"github.com/Murolando/hakaton_geo/ent"
	"github.com/jmoiron/sqlx"
)

type ClassPostgres struct {
	db *sqlx.DB
}

func NewClassPostgres(db *sqlx.DB) *ClassPostgres {
	return &ClassPostgres{
		db: db,
	}
}

func (r *ClassPostgres) AllClass() ([]*ent.ChildDashClassResponce, error) {

	list := make([]*ent.ChildDashClassResponce, 0)
	query := fmt.Sprintf(`
	SELECT id 
	FROM "%s"
	`, classTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var class ent.ChildDashClassResponce
		if err := rows.Scan(&class.Id); err != nil {
			return nil, err
		}
		list = append(list, &class)
	}
	return list, nil
}
