package main

import (
	"crmeb_go/config"
	"crmeb_go/internal"
	"crmeb_go/internal/cron"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router"
	"crmeb_go/internal/server"
	"crmeb_go/utils/ibinder"
	"crmeb_go/utils/iconfig"
	"crmeb_go/utils/izap"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kardianos/service"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var confFile = flag.String("file", "", "input file path")

// 定义程序结构体
type program struct {
	server     *gin.Engine
	svc        service.Service
	once       sync.Once
	clearFunc  func()
	svcContext *server.SvcContext
}

// Start 方法在服务启动时调用
func (p *program) Start(s service.Service) error {
	// 在一个新的 goroutine 中启动服务
	go p.run()
	return nil
}

// Stop 方法在服务停止时调用
func (p *program) Stop(s service.Service) error {
	// 这里可以添加清理资源的代码
	p.once.Do(func() {
		if p.clearFunc != nil {
			p.clearFunc()
		}
	})
	return nil
}

// 运行 Gin 服务器
func (p *program) run() {
	flag.Parse()

	var c config.Conf
	*confFile = "config/xyj_config.yaml"
	iconfig.New(&c, *confFile)

	// 初始化zap日志库
	izap.NewZap(c.Log)

	p.svcContext = server.NewSvcContext(&c)

	appCxt := internal.Register(p.svcContext)

	// 启动定时任务
	cron.Timer(p.svcContext)

	go func() {
		newApp(c, appCxt)
	}()
}

func main() {
	// 定义服务配置
	svcConfig := &service.Config{
		Name:        "crmeb_go",
		DisplayName: "personal mall project",
		Description: "a fully functional mall project is worth learning",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.svc = s

	// 设置日志
	errs := make(chan error)
	logger, err := s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err := s.Run(); err != nil {
		logger.Error(err)
	}

	// 处理错误
	go func() {
		for {
			select {
			case err := <-errs:
				log.Println("Error:", err)
			}
		}
	}()

	// 等待中断信号以优雅关闭
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down... time:", time.Now())
}

func newApp(c config.Conf, appCxt *internal.AppContent) {
	// 创建并配置验证器
	r := gin.New()

	// CorsByRules 按照配置的规则放行跨域请求
	r.Use(middleware.CorsByRules(c))

	binding.Validator = new(ibinder.Validator)

	router.Register(r, appCxt)

	izap.Log.Info("use middleware cors")

	err := r.Run(c.Server.Http.Addr)

	if err != nil {
		panic(err)
	}
}
