package user

func Register() *Service {
	repo := NewInMemoryRepository()
	service := newService(repo)
	return service
}
