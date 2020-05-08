//+build wireinject

package registry

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"

	"github.com/MegaBlackLabel/go-clean-architecture-practice/adapter/controllers"
	repo "github.com/MegaBlackLabel/go-clean-architecture-practice/adapter/gateways/memory"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/infrastructure/handler"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/infrastructure/memory"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/usecase/interactor"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/utils/logs/lggr"
)

func InitialiseApps() handler.UserHandler {
	wire.Build(
		lggr.NewLogger,
		memory.NewUserMemory,
		repo.NewUserRepository,
		interactor.NewUserInteractor,
		controllers.NewUserController,
		handler.NewUserHandler,
	)
	return nil
}
