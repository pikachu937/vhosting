package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	authhandler "github.com/mikerumy/vhosting/internal/auth/handler"
	authrepo "github.com/mikerumy/vhosting/internal/auth/repository"
	authusecase "github.com/mikerumy/vhosting/internal/auth/usecase"
	lg "github.com/mikerumy/vhosting/internal/logging"
	logrepo "github.com/mikerumy/vhosting/internal/logging/repository"
	logusecase "github.com/mikerumy/vhosting/internal/logging/usecase"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	sessrepo "github.com/mikerumy/vhosting/internal/session/repository"
	sessusecase "github.com/mikerumy/vhosting/internal/session/usecase"
	"github.com/mikerumy/vhosting/internal/user"
	userhandler "github.com/mikerumy/vhosting/internal/user/handler"
	userrepo "github.com/mikerumy/vhosting/internal/user/repository"
	userusecase "github.com/mikerumy/vhosting/internal/user/usecase"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
	ugrepo "github.com/mikerumy/vhosting/internal/usergroup/repository"
	ugusecase "github.com/mikerumy/vhosting/internal/usergroup/usecase"
	up "github.com/mikerumy/vhosting/internal/userperm"
	uphandler "github.com/mikerumy/vhosting/internal/userperm/handler"
	uprepo "github.com/mikerumy/vhosting/internal/userperm/repository"
	upusecase "github.com/mikerumy/vhosting/internal/userperm/usecase"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	logger "github.com/mikerumy/vhosting/pkg/logger"
)

type App struct {
	httpServer  *http.Server
	cfg         config_tool.Config
	userUseCase user.UserUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	logUseCase  lg.LogUseCase
	ugUseCase   ug.UGUseCase
	upUseCase   up.UPUseCase
}

func NewApp(cfg config_tool.Config) *App {
	userRepo := userrepo.NewUserRepository(cfg)
	authRepo := authrepo.NewAuthRepository(cfg)
	sessRepo := sessrepo.NewSessRepository(cfg)
	logRepo := logrepo.NewLogRepository(cfg)
	ugRepo := ugrepo.NewUGRepository(cfg)
	upRepo := uprepo.NewUPRepository(cfg)

	return &App{
		cfg:         cfg,
		userUseCase: userusecase.NewUserUseCase(cfg, userRepo),
		authUseCase: authusecase.NewAuthUseCase(cfg, authRepo),
		sessUseCase: sessusecase.NewSessUseCase(sessRepo, authRepo),
		logUseCase:  logusecase.NewLogUseCase(logRepo),
		ugUseCase:   ugusecase.NewUGUseCase(ugRepo),
		upUseCase:   upusecase.NewUPUseCase(upRepo),
	}
}

func (a *App) Run() error {
	// Debug mode
	if a.cfg.ServerDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init handlers
	router := gin.New()

	// Register routes
	authhandler.RegisterHTTPEndpoints(router, a.authUseCase, a.userUseCase,
		a.sessUseCase, a.logUseCase)
	userhandler.RegisterHTTPEndpoints(router, a.userUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.ugUseCase)
	uphandler.RegisterHTTPEndpoints(router, a.upUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.ugUseCase, a.userUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + a.cfg.ServerPort,
		Handler:        router,
		ReadTimeout:    time.Duration(a.cfg.ServerReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(a.cfg.ServerWriteTimeoutSeconds) * time.Second,
		MaxHeaderBytes: a.cfg.ServerMaxHeaderBytes,
	}

	var err error
	notStarted := false

	// Server start
	go func() {
		err = a.httpServer.ListenAndServe()
		if err != nil {
			notStarted = true
		}
	}()
	time.Sleep(50 * time.Millisecond)
	if notStarted {
		return errors.New(fmt.Sprintf("Cannot start server. Error: %s.", err.Error()))
	}
	logger.Print(msg.InfoServerWasSuccessfullyStartedAtLocalIP(getOutboundIP().String(), a.cfg.ServerPort))

	// Server shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	time.Sleep(2 * time.Second)

	ctx, shutdown := context.WithTimeout(context.Background(), 1700*time.Millisecond)
	defer shutdown()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return errors.New(fmt.Sprintf("Cannot shut down the server correctly. Error: %s.", err.Error()))
	}

	logger.Print(msg.InfoServerWasGracefullyShutDown())
	return nil
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Print(msg.WarningCannotGetLocalIP(err))
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
