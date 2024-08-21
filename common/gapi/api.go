package gapi

import (
	"github.com/bufbuild/protovalidate-go"
	accountsUsecase "github.com/darwishdev/devkit-api-base/app/accounts/usecase"

	propertiesUsecase "github.com/darwishdev/devkit-api-base/app/properties/usecase"
	ratesUsecase "github.com/darwishdev/devkit-api-base/app/rates/usecase"
	reservationsUsecase "github.com/darwishdev/devkit-api-base/app/reservations/usecase"

	publicUsecase "github.com/darwishdev/devkit-api-base/app/public/usecase"
	"github.com/darwishdev/devkit-api-base/common/auth"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	"github.com/darwishdev/devkit-api-base/common/pb/abc/v1/abcv1connect"
	"github.com/darwishdev/devkit-api-base/common/redisclient"
	"github.com/darwishdev/devkit-api-base/common/resend"
	"github.com/darwishdev/devkit-api-base/common/supaclient"

	"github.com/darwishdev/devkit-api-base/config"
)

// Server serves gRPC requests for our banking usecase.
type Api struct {
	abcv1connect.UnimplementedAbcServiceHandler
	config              config.Config
	tokenMaker          auth.Maker
	accountsUsecase     accountsUsecase.AccountsUsecaseInterface
	propertiesUsecase   propertiesUsecase.PropertiesUsecaseInterface
	ratesUsecase        ratesUsecase.RatesUsecaseInterface
	redisClient         redisclient.RedisClientInterface
	publicUsecase       publicUsecase.PublicUsecaseInterface
	supaClient          supaclient.SupabaseServiceInterface
	store               db.Store
	reservationsUsecase reservationsUsecase.ReservationsUsecaseInterface
}

// NewServer creates a new gRPC server.
func NewApi(config config.Config, store db.Store, redisClient redisclient.RedisClientInterface) abcv1connect.AbcServiceHandler {
	validator, err := protovalidate.New()

	if err != nil {
		panic("cann't create validator in gapi/api.go")
	}
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic("cann't create paset maker in gapi/api.go")
	}

	supaClient, err := supaclient.NewSupabaseService(config.SupaUrl, config.SupaKey)
	if err != nil {
		panic("cann't create supabase client in gapi/api.go")
	}
	resendClient, err := resend.NewResendService(config.ResendApiKey, config.ClientBaseUrl)
	if err != nil {
		panic("cann't create resendbase client in gapi/api.go")
	}
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, validator, tokenMaker, config.AccessTokenDuration, redisClient, supaClient, config.SupaKey, resendClient)
	propertiesUsecase := propertiesUsecase.NewPropertiesUsecase(store, validator)
	ratesUsecase := ratesUsecase.NewRatesUsecase(store, validator)
	reservationsUsecase := reservationsUsecase.NewReservationsUsecase(store, validator, resendClient)
	publicUsecase := publicUsecase.NewPublicUsecase(store, validator, supaClient)
	return &Api{
		config:              config,
		tokenMaker:          tokenMaker,
		store:               store,
		supaClient:          supaClient,
		accountsUsecase:     accountsUsecase,
		redisClient:         redisClient,
		propertiesUsecase:   propertiesUsecase,
		publicUsecase:       publicUsecase,
		ratesUsecase:        ratesUsecase,
		reservationsUsecase: reservationsUsecase,
	}
}
