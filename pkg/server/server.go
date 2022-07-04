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

	"github.com/dmitrij/vhosting/internal/group"
	grouphandler "github.com/dmitrij/vhosting/internal/group/handler"
	grouprepo "github.com/dmitrij/vhosting/internal/group/repository"
	groupusecase "github.com/dmitrij/vhosting/internal/group/usecase"
	"github.com/dmitrij/vhosting/internal/info"
	infohandler "github.com/dmitrij/vhosting/internal/info/handler"
	inforepo "github.com/dmitrij/vhosting/internal/info/repository"
	infousecase "github.com/dmitrij/vhosting/internal/info/usecase"
	msg "github.com/dmitrij/vhosting/internal/messages"
	perm "github.com/dmitrij/vhosting/internal/permission"
	permhandler "github.com/dmitrij/vhosting/internal/permission/handler"
	permrepo "github.com/dmitrij/vhosting/internal/permission/repository"
	permusecase "github.com/dmitrij/vhosting/internal/permission/usecase"
	sess "github.com/dmitrij/vhosting/internal/session"
	sessrepo "github.com/dmitrij/vhosting/internal/session/repository"
	sessusecase "github.com/dmitrij/vhosting/internal/session/usecase"
	"github.com/dmitrij/vhosting/internal/video"
	videohandler "github.com/dmitrij/vhosting/internal/video/handler"
	videorepo "github.com/dmitrij/vhosting/internal/video/repository"
	videousecase "github.com/dmitrij/vhosting/internal/video/usecase"
	"github.com/dmitrij/vhosting/pkg/auth"
	authhandler "github.com/dmitrij/vhosting/pkg/auth/handler"
	authrepo "github.com/dmitrij/vhosting/pkg/auth/repository"
	authusecase "github.com/dmitrij/vhosting/pkg/auth/usecase"
	"github.com/dmitrij/vhosting/pkg/config"
	sconfig "github.com/dmitrij/vhosting/pkg/config_stream"
	"github.com/dmitrij/vhosting/pkg/download"
	downloadhandler "github.com/dmitrij/vhosting/pkg/download/handler"
	downloadusecase "github.com/dmitrij/vhosting/pkg/download/usecase"
	"github.com/dmitrij/vhosting/pkg/logger"
	logrepo "github.com/dmitrij/vhosting/pkg/logger/repository"
	logusecase "github.com/dmitrij/vhosting/pkg/logger/usecase"
	"github.com/dmitrij/vhosting/pkg/stream"
	streamhandler "github.com/dmitrij/vhosting/pkg/stream/handler"
	streamrepo "github.com/dmitrij/vhosting/pkg/stream/repository"
	streamusecase "github.com/dmitrij/vhosting/pkg/stream/usecase"
	"github.com/dmitrij/vhosting/pkg/user"
	userhandler "github.com/dmitrij/vhosting/pkg/user/handler"
	userrepo "github.com/dmitrij/vhosting/pkg/user/repository"
	userusecase "github.com/dmitrij/vhosting/pkg/user/usecase"
	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer      *http.Server
	cfg             *config.Config
	scfg            *sconfig.Config
	userUseCase     user.UserUseCase
	authUseCase     auth.AuthUseCase
	sessUseCase     sess.SessUseCase
	logUseCase      logger.LogUseCase
	groupUseCase    group.GroupUseCase
	permUseCase     perm.PermUseCase
	infoUseCase     info.InfoUseCase
	videoUseCase    video.VideoUseCase
	StreamUC        stream.StreamUseCase
	downloadUseCase download.DownloadUseCase
}

func NewApp(cfg *config.Config, scfg *sconfig.Config) *App {
	userRepo := userrepo.NewUserRepository(cfg)
	authRepo := authrepo.NewAuthRepository(cfg)
	sessRepo := sessrepo.NewSessRepository(cfg)
	logRepo := logrepo.NewLogRepository(cfg)
	groupRepo := grouprepo.NewGroupRepository(cfg)
	permRepo := permrepo.NewPermRepository(cfg)
	infoRepo := inforepo.NewInfoRepository(cfg)
	videoRepo := videorepo.NewVideoRepository(cfg)
	streamRepo := streamrepo.NewStreamRepository(cfg)

	return &App{
		cfg:             cfg,
		scfg:            scfg,
		userUseCase:     userusecase.NewUserUseCase(cfg, userRepo),
		authUseCase:     authusecase.NewAuthUseCase(cfg, authRepo),
		sessUseCase:     sessusecase.NewSessUseCase(sessRepo, authRepo),
		logUseCase:      logusecase.NewLogUseCase(logRepo),
		groupUseCase:    groupusecase.NewGroupUseCase(groupRepo),
		permUseCase:     permusecase.NewPermUseCase(permRepo),
		infoUseCase:     infousecase.NewInfoUseCase(infoRepo),
		videoUseCase:    videousecase.NewVideoUseCase(videoRepo),
		StreamUC:        streamusecase.NewStreamUseCase(cfg, scfg, streamRepo),
		downloadUseCase: downloadusecase.NewDownloadUseCase(cfg),
	}
}

func (a *App) Run() error {
	// Debug mode
	if a.cfg.ServerDebugEnable {
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
		streamhandler.RegisterTemplateHTTPEndpoints(router, a.cfg, a.scfg, a.StreamUC,
			a.userUseCase, a.logUseCase, a.authUseCase, a.sessUseCase)
	}

	router.StaticFS("/static", http.Dir("./web/static"))

	// Register routes
	authhandler.RegisterHTTPEndpoints(router, a.cfg, a.authUseCase, a.userUseCase,
		a.sessUseCase, a.logUseCase)
	userhandler.RegisterHTTPEndpoints(router, a.cfg, a.userUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase)
	grouphandler.RegisterHTTPEndpoints(router, a.cfg, a.groupUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)
	permhandler.RegisterHTTPEndpoints(router, a.cfg, a.permUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase, a.groupUseCase)
	infohandler.RegisterHTTPEndpoints(router, a.cfg, a.scfg, a.infoUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)
	videohandler.RegisterHTTPEndpoints(router, a.cfg, a.videoUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)
	streamhandler.RegisterStreamingHTTPEndpoints(router, a.cfg, a.scfg, a.StreamUC,
		a.userUseCase, a.logUseCase, a.authUseCase, a.sessUseCase)
	downloadhandler.RegisterHTTPEndpoints(router, a.cfg, a.downloadUseCase, a.logUseCase,
		a.authUseCase, a.sessUseCase, a.userUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", a.cfg.ServerHost, a.cfg.ServerPort),
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
