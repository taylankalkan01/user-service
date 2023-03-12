package user

import "gorm.io/gorm"

type Repos interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Migration() error
}

type repos struct {
	db *gorm.DB
}

var _ Repos = repos{}

func NewRepo(db *gorm.DB) Repos {
	return repos{db: db}
}

func (repo repos) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.First(model).Error
	//SELECT * FROM users WHERE id=1
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repos) Create(model Model) (uint, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (repo repos) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}
