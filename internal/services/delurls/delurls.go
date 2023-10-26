package delurls

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

// Task - задача
type Task struct {
	UserID string
	URLs   []string
}

// Worker воркер
type Worker struct {
	channel chan *Task
	deleter *Deleter
}

// NewWorker конструктор воркеров
func NewWorker(channel chan *Task, deleter *Deleter) *Worker {
	w := Worker{
		channel: channel,
		deleter: deleter,
	}
	return &w
}

// PopWait - получение задачи из очереди
func (w *Worker) PopWait() *Task {
	// получаем задачу
	return <-w.channel
}

// Loop -
func (w *Worker) Loop() {

	for {

		t := w.PopWait()

		err := w.deleter.Delete(t.URLs)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker deleted urls\n")

	}
}

// Deleter -
type Deleter struct {
	Storage storage.Storage
}

// NewDeleter -
func NewDeleter(storage storage.Storage) *Deleter {
	return &Deleter{
		Storage: storage,
	}
}

// Delete -
func (d *Deleter) Delete(urls []string) error {
	if err := d.Storage.DelURL(urls); err != nil {
		logger.ZapLogger.Error("don't delete urls", zap.Error(err))
		return err
	}
	return nil
}
