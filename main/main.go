package main

import (
	"crmeb_go/config"
	"crmeb_go/internal/server"
	xconfig "crmeb_go/util/config"
	"flag"
	"fmt"
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
	xconfig.New(&c, *confFile)

	//logger := xlog.NewLogger(c.Log.Path, c.App.Name)

	p.svcContext = server.NewSvcContext(c, logger)

	appCxt := internal.Register(p.svcContext)

	go func() {
		newApp(c, appCxt)
	}()

	return nil
}

//func newApp(c config.Conf, appCxt *internal.AppContent) {
//	// 创建并配置验证器
//	r := gin.New()
//
//	binding.Validator = new(binder.Validator)
//
//	r.Use(otelgin.Middleware(c.App.Name, otelgin.WithPropagators(propagation.TraceContext{})))
//
//	router.Register(r, appCxt)
//
//	err := r.Run(c.Server.Http.Addr)
//}

// Stop 程序停止
func (p *program) Stop() error {
	p.once.Do(func() {
		//defer p.svcContext.RedisClient.Close()
		//defer p.svcContext.DBEngine.Close()
	})
	return nil
}
