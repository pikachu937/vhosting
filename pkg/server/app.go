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
	"github.com/mikerumy/vhosting/internal/group"
	grouphandler "github.com/mikerumy/vhosting/internal/group/handler"
	grouprepo "github.com/mikerumy/vhosting/internal/group/repository"
	groupusecase "github.com/mikerumy/vhosting/internal/group/usecase"
	"github.com/mikerumy/vhosting/internal/info"
	infohandler "github.com/mikerumy/vhosting/internal/info/handler"
	inforepo "github.com/mikerumy/vhosting/internal/info/repository"
	infousecase "github.com/mikerumy/vhosting/internal/info/usecase"
	lg "github.com/mikerumy/vhosting/internal/logging"
	logrepo "github.com/mikerumy/vhosting/internal/logging/repository"
	logusecase "github.com/mikerumy/vhosting/internal/logging/usecase"
	msg "github.com/mikerumy/vhosting/internal/messages"
	perm "github.com/mikerumy/vhosting/internal/permission"
	permhandler "github.com/mikerumy/vhosting/internal/permission/handler"
	permrepo "github.com/mikerumy/vhosting/internal/permission/repository"
	permusecase "github.com/mikerumy/vhosting/internal/permission/usecase"
	sess "github.com/mikerumy/vhosting/internal/session"
	sessrepo "github.com/mikerumy/vhosting/internal/session/repository"
	sessusecase "github.com/mikerumy/vhosting/internal/session/usecase"
	"github.com/mikerumy/vhosting/internal/user"
	userhandler "github.com/mikerumy/vhosting/internal/user/handler"
	userrepo "github.com/mikerumy/vhosting/internal/user/repository"
	userusecase "github.com/mikerumy/vhosting/internal/user/usecase"
	"github.com/mikerumy/vhosting/internal/video"
	videohandler "github.com/mikerumy/vhosting/internal/video/handler"
	videorepo "github.com/mikerumy/vhosting/internal/video/repository"
	videousecase "github.com/mikerumy/vhosting/internal/video/usecase"
	"github.com/mikerumy/vhosting/pkg/config"
	logger "github.com/mikerumy/vhosting/pkg/logger"
)

type App struct {
	httpServer   *http.Server
	cfg          config.Config
	userUseCase  user.UserUseCase
	authUseCase  auth.AuthUseCase
	sessUseCase  sess.SessUseCase
	logUseCase   lg.LogUseCase
	groupUseCase group.GroupUseCase
	permUseCase  perm.PermUseCase
	infoUseCase  info.InfoUseCase
	videoUseCase video.VideoUseCase
}

func NewApp(cfg config.Config) *App {
	userRepo := userrepo.NewUserRepository(cfg)
	authRepo := authrepo.NewAuthRepository(cfg)
	sessRepo := sessrepo.NewSessRepository(cfg)
	logRepo := logrepo.NewLogRepository(cfg)
	groupRepo := grouprepo.NewGroupRepository(cfg)
	permRepo := permrepo.NewPermRepository(cfg)
	infoRepo := inforepo.NewInfoRepository(cfg)
	videoRepo := videorepo.NewVideoRepository(cfg)

	return &App{
		cfg:          cfg,
		userUseCase:  userusecase.NewUserUseCase(cfg, userRepo),
		authUseCase:  authusecase.NewAuthUseCase(cfg, authRepo),
		sessUseCase:  sessusecase.NewSessUseCase(sessRepo, authRepo),
		logUseCase:   logusecase.NewLogUseCase(logRepo),
		groupUseCase: groupusecase.NewGroupUseCase(cfg, groupRepo),
		permUseCase:  permusecase.NewPermUseCase(cfg, permRepo),
		infoUseCase:  infousecase.NewInfoUseCase(cfg, infoRepo),
		videoUseCase: videousecase.NewVideoUseCase(cfg, videoRepo),
	}
}

func (a *App) Run() error {
	var err error

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
		a.authUseCase, a.sessUseCase)
	grouphandler.RegisterHTTPEndpoints(router, a.groupUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)
	permhandler.RegisterHTTPEndpoints(router, a.permUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase, a.groupUseCase)
	infohandler.RegisterHTTPEndpoints(router, a.infoUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)
	videohandler.RegisterHTTPEndpoints(router, a.videoUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + a.cfg.ServerPort,
		Handler:        router,
		ReadTimeout:    time.Duration(a.cfg.ServerReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(a.cfg.ServerWriteTimeoutSeconds) * time.Second,
		MaxHeaderBytes: a.cfg.ServerMaxHeaderBytes,
	}

	// Server start
	notStarted := false
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

	if err = a.httpServer.Shutdown(ctx); err != nil {
		return errors.New(fmt.Sprintf("Cannot shut down the server correctly. Error: %s.", err.Error()))
	}

	logger.Print(msg.InfoServerWasGracefullyShutDown())

	return nil
}

func getOutboundIP() net.IP {
	var err error
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Print(msg.WarningCannotGetLocalIP(err))
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
