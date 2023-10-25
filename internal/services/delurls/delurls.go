package delurls

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/middleware/logger"
	"github.com/al-kirpichenko/shortlinks/internal/storage"
)

type Task struct {
	UserID string
	URLs   []string
}

// Queue - это очередь задач
type Queue struct {
	ch chan *Task
}

func NewQueue(channel chan *Task) *Queue {
	return &Queue{
		ch: channel,
	}
}

func (q *Queue) Push(t *Task) {
	// добавляем задачу в очередь
	q.ch <- t
}

func (q *Queue) PopWait() *Task {
	// получаем задачу
	return <-q.ch
}

type Worker struct {
	id      int
	queue   *Queue
	deleter *Deleter
}

func NewWorker(id int, queue *Queue, resizer *Deleter) *Worker {
	w := Worker{
		id:      id,
		queue:   queue,
		deleter: resizer,
	}
	return &w
}

func (w *Worker) Loop(sigint chan os.Signal) {
	var shutdown bool

	for !shutdown {

		t := w.queue.PopWait()

		err := w.deleter.Delete(t.URLs)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker #%d deleted urls\n", w.id)

		if <-sigint; true {
			shutdown = true
			fmt.Printf("workers GRACEFULLY\n")
		}
	}
}

type Deleter struct {
	Storage storage.Storage
}

func NewDeleter(storage storage.Storage) *Deleter {
	return &Deleter{
		Storage: storage,
	}
}

func (d *Deleter) Delete(urls []string) error {
	if err := d.Storage.DelURL(urls); err != nil {
		logger.ZapLogger.Error("don't delete urls", zap.Error(err))
		return err
	}
	return nil
}
