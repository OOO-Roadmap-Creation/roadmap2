package repository

import (
	"fmt"
	"github.com/OOO-Roadmap-Creation/roadmap2"
	"github.com/jmoiron/sqlx"
)

type PublishedRoadmapPostgres struct {
	db *sqlx.DB
}

func NewPublishedRoadmapPostgres(db *sqlx.DB) *PublishedRoadmapPostgres {
	return &PublishedRoadmapPostgres{db: db}
}

func (r *PublishedRoadmapPostgres) Create(rm roadmap2.PublishedRoadmap) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (id, \"version\", visible, title, description, dateOfPublish) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", published_roadmapsTable)
	row := tx.QueryRow(createListQuery, rm.Id, rm.Version, rm.Visible, rm.Title, rm.Description, rm.DateOfPublish)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	// TODO SAVE NODES

	return id, tx.Commit()
}

func (r *PublishedRoadmapPostgres) GetAll() ([]roadmap2.PublishedRoadmap, error) {
	var lists []roadmap2.PublishedRoadmap

	query := fmt.Sprintf("SELECT * FROM %s;", published_roadmapsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *PublishedRoadmapPostgres) GetById(rId int) (roadmap2.PublishedRoadmap, error) {
	var rm roadmap2.PublishedRoadmap

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1 ORDER BY version DESC LIMIT 1`,
		published_roadmapsTable)
	err := r.db.Get(&rm, query, rId)

	return rm, err
}
