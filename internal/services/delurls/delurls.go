package delurls

import (
	"fmt"
	"sync"

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
	Channel   chan *Task
	deleter   *Deleter
	WaitGroup sync.WaitGroup
}

// NewWorker конструктор воркеров
func NewWorker(deleter *Deleter) *Worker {
	w := Worker{
		deleter: deleter,
	}
	return &w
}

// Run старт воркера
func (w *Worker) Run() {

	w.Channel = make(chan *Task, 1000)
	w.WaitGroup.Add(1)

	go func() {
		w.Loop()
		w.WaitGroup.Done()
	}()
}

func (w *Worker) Stop() {
	close(w.Channel)
	w.WaitGroup.Wait()
}

// PopWait - получение задачи из очереди
func (w *Worker) PopWait() (*Task, bool) {
	// получаем задачу
	task, ok := <-w.Channel
	return task, ok
}

// Loop -
func (w *Worker) Loop() {

	for {

		t, ok := w.PopWait()

		if !ok {
			break
		}

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
