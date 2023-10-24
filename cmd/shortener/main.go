package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"golang.org/x/net/context"

	"github.com/al-kirpichenko/shortlinks/cmd/shortener/config"
	"github.com/al-kirpichenko/shortlinks/internal/app"
	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/routes"
	"github.com/al-kirpichenko/shortlinks/internal/services/delurls"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	conf := config.NewAppConfig()

	logger.InitLogger()

	newApp := app.NewApp(conf)

	newApp.ConfigureStorage()

	newApp.Channel = make(chan *delurls.Task, 1)

	srv := &http.Server{
		Addr:    conf.Host,
		Handler: routes.Router(newApp),
	}

	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})

	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	//создаем контекст для корректного завершения удаления ссылок
	ctx, cancelFunc := context.WithCancel(context.Background())

	queue := delurls.NewQueue(newApp.Channel)

	for i := 0; i < runtime.NumCPU(); i++ {
		w := delurls.NewWorker(i, queue, delurls.NewDeleter(newApp.Storage))
		go w.Loop(ctx)

	}

	// запускаем горутину обработки пойманных прерываний
	go func() {
		// читаем из канала прерываний
		<-sigint

		cancelFunc()
		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := srv.Shutdown(context.Background()); err != nil {
			// ошибки закрытия Listener
			log.Printf("HTTP server Shutdown: %v", err)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(idleConnsClosed)
	}()
	if conf.EnableHTTPS {
		runHTTPS(srv)
	} else {
		run(srv)
	}

	// ждём завершения процедуры graceful shutdown
	<-idleConnsClosed
	// получили оповещение о завершении
	// здесь можно освобождать ресурсы перед выходом,
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы
	fmt.Println("Server Shutdown gracefully")

}

func run(srv *http.Server) {

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// ошибки запуска или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func runHTTPS(srv *http.Server) {

	if err := srv.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem"); err != http.ErrServerClosed {
		// ошибки запуска или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
