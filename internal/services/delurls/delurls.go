package delurls

import (
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/net/context"

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

func (w *Worker) Loop(ctx context.Context) {
	for {

		t := w.queue.PopWait()

		err := w.deleter.Delete(t.URLs, t.UserID)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker #%d deleted urls userID #%s\n", w.id, t.UserID)

		if <-ctx.Done(); true {
			fmt.Printf("workers GRACEFULLY\n")
			break
		}
	}
	close(w.queue.ch)
	return
}

type Deleter struct {
	Storage storage.Storage
}

func NewDeleter(storage storage.Storage) *Deleter {
	return &Deleter{
		Storage: storage,
	}
}

func (d *Deleter) Delete(urls []string, userID string) error {
	if err := d.Storage.DelURL(urls, userID); err != nil {
		logger.ZapLogger.Error("don't delete urls", zap.Error(err))
		return err
	}
	return nil
}
