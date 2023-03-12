package user

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
}

type service struct {
	repo Repos
}

var _ Service = service{}

func NewService(repo Repos) service {
	return service{repo: repo}
}

func (s service) Get(id uint) (*Model, error) {
	return s.repo.Get(id)

}

func (s service) Create(model Model) (uint, error) {
	return s.repo.Create(model)

}
