package todo

import (
	v1 "github.com/heggies/todo-server/src/database/v1"
	"github.com/heggies/todo-server/src/entity/v1/todo"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (repo *Repository, err error) {
	db, err := v1.GetInstance()
	if err != nil {
		return
	}

	repo = &Repository{
		db: db,
	}
	return
}

func (repo *Repository) Get() (entities []todo.Todo, err error) {
	query := repo.db
	query = query.Order("updated_at DESC")
	query = query.Find(&entities)
	err = query.Error

	return
}

func (repo *Repository) Create(entity todo.Todo) (todo.Todo, error) {
	query := repo.db.Begin()

	query = query.Create(&entity)
	if err := query.Error; err != nil {
		query.Rollback()
		return entity, err
	}

	err := query.Commit().Error
	return entity, err
}

func (repo *Repository) Update(entity todo.Todo) (todo.Todo, error) {
	query := repo.db.Model(&entity).Begin()
	query = query.Updates(entity)

	if query.Error != nil {
		query.Rollback()
		return entity, query.Error
	}

	err := query.Commit().Error
	if err != nil {
		return entity, err
	}

	if query.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return entity, err
}

func (repo *Repository) Delete(id int) (err error) {
	entity := todo.Todo{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	query := repo.db.Model(&entity).Begin()
	query = query.Delete(&entity)

	if query.Error != nil {
		query.Rollback()
		err = query.Error
		return
	}

	err = query.Commit().Error
	if err != nil {
		return
	}

	if query.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}
