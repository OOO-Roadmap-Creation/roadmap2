package repository

import "github.com/jmoiron/sqlx"

//	type TodoItem interface {
//		Create(listId int, item roadmap2.TodoItem) (int, error)
//		GetAll(userId, listId int) ([]roadmap2.TodoItem, error)
//		GetById(userId, itemId int) (roadmap2.TodoItem, error)
//		Delete(userId, itemId int) error
//		Update(userId, itemId int, input roadmap2.UpdateItemInput) error
//	}
type Repository struct {
	PR *PublishedRoadmapPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PR: NewPublishedRoadmapPostgres(db),
	}
}
