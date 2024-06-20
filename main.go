package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	_ "tanyaustadz/docs"
	"tanyaustadz/internal/config"
	docs_handler "tanyaustadz/internal/delivery/http/docs"
	static_handler "tanyaustadz/internal/delivery/http/static"
	user_handler "tanyaustadz/internal/delivery/http/user"
	user_repository "tanyaustadz/internal/repository/psql/user"
	"tanyaustadz/pkg/logger"
	"tanyaustadz/pkg/service/jwt"

	"github.com/gorilla/mux"
)

var (
	cfg        = config.Server()
	appLogger  = logger.NewApiLogger()
	db         = config.InitDatabase()
	jwtService = jwt.NewJWTService()
	userRepo   = user_repository.NewUserRepository(db)
)

func main() {
	psqlConn := config.InitDatabase()
	defer func(db *sql.DB) { _ = db.Close() }(psqlConn)

	router := mux.NewRouter()

	initHandler(router)
	http.Handle("/", router)

	appLogger.Info("tanyaustadz Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(cfg.Port, router)
		if err != nil {
			appLogger.Error(err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}

func initHandler(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("tanyaustadz"))
	})
	user_handler.UserHandler(router, jwtService, userRepo)
	docs_handler.DocsHandler(router)
	static_handler.StaticHandler(router)
}
