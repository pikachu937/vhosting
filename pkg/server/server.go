package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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
	"github.com/mikerumy/vhosting/internal/models"
	perm "github.com/mikerumy/vhosting/internal/permission"
	permhandler "github.com/mikerumy/vhosting/internal/permission/handler"
	permrepo "github.com/mikerumy/vhosting/internal/permission/repository"
	permusecase "github.com/mikerumy/vhosting/internal/permission/usecase"
	sess "github.com/mikerumy/vhosting/internal/session"
	sessrepo "github.com/mikerumy/vhosting/internal/session/repository"
	sessusecase "github.com/mikerumy/vhosting/internal/session/usecase"
	"github.com/mikerumy/vhosting/internal/video"
	videohandler "github.com/mikerumy/vhosting/internal/video/handler"
	videorepo "github.com/mikerumy/vhosting/internal/video/repository"
	videousecase "github.com/mikerumy/vhosting/internal/video/usecase"
	"github.com/mikerumy/vhosting/pkg/auth"
	authhandler "github.com/mikerumy/vhosting/pkg/auth/handler"
	authrepo "github.com/mikerumy/vhosting/pkg/auth/repository"
	authusecase "github.com/mikerumy/vhosting/pkg/auth/usecase"
	"github.com/mikerumy/vhosting/pkg/config"
	"github.com/mikerumy/vhosting/pkg/download"
	downloadhandler "github.com/mikerumy/vhosting/pkg/download/handler"
	downloadusecase "github.com/mikerumy/vhosting/pkg/download/usecase"
	logger "github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/stream"
	streamhandler "github.com/mikerumy/vhosting/pkg/stream/handler"
	streamusecase "github.com/mikerumy/vhosting/pkg/stream/usecase"
	"github.com/mikerumy/vhosting/pkg/user"
	userhandler "github.com/mikerumy/vhosting/pkg/user/handler"
	userrepo "github.com/mikerumy/vhosting/pkg/user/repository"
	userusecase "github.com/mikerumy/vhosting/pkg/user/usecase"
)

type App struct {
	httpServer      *http.Server
	cfg             *config.Config
	scfg            *models.ConfigST
	userUseCase     user.UserUseCase
	authUseCase     auth.AuthUseCase
	sessUseCase     sess.SessUseCase
	logUseCase      lg.LogUseCase
	groupUseCase    group.GroupUseCase
	permUseCase     perm.PermUseCase
	infoUseCase     info.InfoUseCase
	videoUseCase    video.VideoUseCase
	StreamUC        stream.StreamUseCase
	downloadUseCase download.DownloadUseCase
}

func NewApp(cfg *config.Config, scfg *models.ConfigST) *App {
	userRepo := userrepo.NewUserRepository(cfg)
	authRepo := authrepo.NewAuthRepository(cfg)
	sessRepo := sessrepo.NewSessRepository(cfg)
	logRepo := logrepo.NewLogRepository(cfg)
	groupRepo := grouprepo.NewGroupRepository(cfg)
	permRepo := permrepo.NewPermRepository(cfg)
	infoRepo := inforepo.NewInfoRepository(cfg)
	videoRepo := videorepo.NewVideoRepository(cfg)

	return &App{
		cfg:             cfg,
		scfg:            scfg,
		userUseCase:     userusecase.NewUserUseCase(cfg, userRepo),
		authUseCase:     authusecase.NewAuthUseCase(cfg, authRepo),
		sessUseCase:     sessusecase.NewSessUseCase(sessRepo, authRepo),
		logUseCase:      logusecase.NewLogUseCase(logRepo),
		groupUseCase:    groupusecase.NewGroupUseCase(cfg, groupRepo),
		permUseCase:     permusecase.NewPermUseCase(cfg, permRepo),
		infoUseCase:     infousecase.NewInfoUseCase(cfg, infoRepo),
		videoUseCase:    videousecase.NewVideoUseCase(cfg, videoRepo),
		StreamUC:        streamusecase.NewStreamUseCase(scfg),
		downloadUseCase: downloadusecase.NewDownloadUseCase(cfg),
	}
}

func (a *App) Run() error {
	// Debug mode
	if a.cfg.ServerDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init engine
	router := gin.New()

	// Init middleware
	router.Use(CORSMiddleware())

	// Check for web directory exists and register routes
	if _, err := os.Stat("./web"); !os.IsNotExist(err) {
		router.LoadHTMLGlob("./web/templates/*")
		streamhandler.RegisterTemplateHTTPEndpoints(router, a.StreamUC, a.scfg)
	}

	router.StaticFS("/static", http.Dir("./web/static"))

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
	streamhandler.RegisterStreamingHTTPEndpoints(router, a.StreamUC, a.scfg)
	downloadhandler.RegisterHTTPEndpoints(router, a.downloadUseCase, a.logUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + a.cfg.ServerPort,
		Handler:        router,
		ReadTimeout:    time.Duration(a.cfg.ServerReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(a.cfg.ServerWriteTimeoutSeconds) * time.Second,
		MaxHeaderBytes: a.cfg.ServerMaxHeaderBytes,
	}

	// Server start
	var err error
	go func() {
		err = a.httpServer.ListenAndServe()
	}()
	time.Sleep(50 * time.Millisecond)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot start server. Error: %s.", err.Error()))
	}
	a.cfg.ServerIP = getOutboundIP()
	logger.Print(msg.InfoServerStartedSuccessfullyAtLocalAddress(a.cfg.ServerIP, a.cfg.ServerPort))

	// Listening for interrupt signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		logger.Print(msg.InfoRecivedSignal(sig))
		done <- true
	}()
	<-done

	// Server shut down
	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return errors.New(fmt.Sprintf("Cannot shut down the server correctly. Error: %s.", err.Error()))
	}

	logger.Print(msg.InfoServerShutedDownCorrectly())

	return nil
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Print(msg.WarningCannotGetLocalIP(err))
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
