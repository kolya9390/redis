package run
/*
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"studentgit.kata.academy/Nikolai/historysearch/internal/config"
	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/responder"
	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/server"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules/geoservis/repository"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules/geoservis/servis"
	"studentgit.kata.academy/Nikolai/historysearch/internal/router"
)

type App struct {
	conf     config.AppConf
	Sig      chan os.Signal
	srv      server.Server
	logger   *zap.Logger
//	Storages *storages.Storages
//	Servises *modules.Services

}

// Application - интерфейс приложения
type Application interface {
	Runner
	InitApper
}

// Runner - интерфейс запуска приложения
type Runner interface {
	Run() error
}

// InitApper - интерфейс инициализации приложения
type InitApper interface {
	InitApp(options ...interface{}) Runner
}


func NewApp() *App {

	return &App{Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() error {

	// создаем контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	errGroup, ctx := errgroup.WithContext(ctx)

	// запускаем горутину для graceful shutdown
	// при получении сигнала SIGINT
	// вызываем cancel для контекста
	errGroup.Go(func() error {
		sigInt := <-a.Sig
		a.logger.Info("signal interrupt recieved", zap.Stringer("os_signal", sigInt))
		cancel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && err != http.ErrServerClosed {
			a.logger.Error("app: server error", zap.Error(err))
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf(err.Error())
	}

	return fmt.Errorf("NOT ERROR")

}

func (a *App)SetConfig() *config.AppConf{

	// init config
	
	env, err := godotenv.Read("web/.env")

	if err != nil {
		log.Println(err)
	}
//"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	configDB := config.DB{
		Host: env["DB_HOST"],
		Port: env["DB_PORT"],
		User: env["DB_USER"],
		Password: env["DB_PASSWORD"],
		Name: env["DB_NAME"],
		Timeout: 5,
		Driver: env["DB_DRIVER"],

	}

 configSrver := config.Server{
	Port: env["SERVER_PORT"],
	ShutdownTimeout: 3,

 }


 return &config.AppConf{
	Server: configSrver,
	DB: configDB,
 }
}

func (a *App) InitApp() Runner {
	
	// инициализация scanner
	//	tableScanner := scanner.NewTableScanner()
	// регистрация таблиц

	// инициализация базы данных sql и его адаптера
	//	dbx, sqlAdapter, err := db.NewSqlDB(*configDB, tableScanner)

	// инициализация мигратора
	//	migrator := migrate.NewMigrator(dbx, *configDB, tableScanner)

	
	// инициализация storage /инициализация хранилищ
	//	newStorages := storage.NewUserStorage(sqlAdapter)

	// инициализация services
	//	services := service.NewUserService(newStorages)

	// инициализация controllers
	
	controllers := modules.NewControllers(&servis.DadataServiceImpl{
		AuthorizationDADATA: a.conf.AuthorizationDADATA},responder.NewResponder(&zap.Logger{}),
	repository.NewGeoRepository(&sqlx.DB{}))

	
	// инициализация роутера


	r := router.NewApiRouter(*controllers)


	// конфигурация сервера

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler:      r, // Здесь должен быть ваш обработчик запросов
	}

	// инициализация сервера

	a.srv = server.NewHttpServer(a.conf.Server, srv, a.logger)

	// возвращаем приложение
	return a


}
*/