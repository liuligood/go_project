package main

import (
	"crmeb_go/config"
	"crmeb_go/internal"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router"
	"crmeb_go/internal/server"
	"crmeb_go/utils/binder"
	iconfig "crmeb_go/utils/iconfig"
	"crmeb_go/utils/izap"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/judwhite/go-svc"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

var confFile = flag.String("file", "", "input file path")

type program struct {
	once       sync.Once
	svcContext *server.SvcContext
}

func main() {
	p := &program{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
		fmt.Println(err)
	}
}

// Init 利于go-svc 帮助管理程序生命周期 程序初始化
func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

// Start 程序启动
func (p *program) Start() error {
	flag.Parse()

	var c config.Conf
	*confFile = "config/local_config.yaml"
	iconfig.New(&c, *confFile)

	// 初始化zap日志库
	izap.NewZap(c.Log)

	p.svcContext = server.NewSvcContext(c)

	appCxt := internal.Register(p.svcContext)

	go func() {
		newApp(c, appCxt)
	}()

	return nil
}

func newApp(c config.Conf, appCxt *internal.AppContent) {
	// 创建并配置验证器
	r := gin.New()

	// CorsByRules 按照配置的规则放行跨域请求
	r.Use(gin.Recovery(), middleware.CorsByRules(c))

	binding.Validator = new(binder.Validator)

	router.Register(r, appCxt)

	appCxt.Svc.Logger.Info("use middleware cors")

	err := r.Run(c.Server.Http.Addr)

	if err != nil {
		panic(err)
	}
}

// Stop 程序停止
func (p *program) Stop() error {
	p.once.Do(func() {
		//defer p.svcContext.RedisClient.Close()
		//defer p.svcContext.DBEngine.Close()
	})
	return nil
}
