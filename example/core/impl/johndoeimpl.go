package impl

import (
	"log"

	"github.com/RajulSahu/lightms/v3/example/core/port"
	"github.com/RajulSahu/lightms/v3/example/core/usecase"
)

func NewJohnDoeImpl(p port.PersistencePort) usecase.JohnDoeUseCase {
	return &JohnDoeImpl{p}
}

type JohnDoeImpl struct {
	persistencePort port.PersistencePort
}

func (t *JohnDoeImpl) Handle(msg string) error {
	log.Printf("Usecase 'JohnDoeUseCase' handling message '%s'.\n", msg)
	return t.persistencePort.Save(msg)
}
