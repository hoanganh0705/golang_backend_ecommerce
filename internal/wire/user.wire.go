//go:build wireinject
// +build wireinject

package wire

import (
	"GolangBackendEcommerce/internal/controller"
	"GolangBackendEcommerce/internal/repo"
	"GolangBackendEcommerce/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
