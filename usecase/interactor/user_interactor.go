package interactor

import (
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputport "cln-arch/usecase/output/port"
	"cln-arch/usecase/repository"
)

type UserInteractor struct {
	userRepository *repository.UserRepository
	outputPort     *outputport.UserOutputPort
}

func NewUserInteractor(
	userRepository *repository.UserRepository,
	outputPort *outputport.UserOutputPort,
) inputport.UserInputPort {
	return &UserInteractor{
		userRepository: userRepository,
		outputPort:     outputPort,
	}
}

func (it *UserInteractor) Create(user *inputdata.User) {

}
