package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

// biasanya awalan nama func dengan New
// provider func
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("failed create service")
	}

	return &SimpleService{SimpleRepository: simpleRepository}, nil
}
