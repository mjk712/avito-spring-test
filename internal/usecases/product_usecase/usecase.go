package product_usecase

type UseCase struct {
	repository repository
}

func New(repository repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}
