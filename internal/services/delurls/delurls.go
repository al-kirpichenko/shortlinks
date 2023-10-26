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

// Queue - это очередь задач
type Queue struct {
	ch chan *Task
}

// NewQueue конструктор
func NewQueue(channel chan *Task) *Queue {
	return &Queue{
		ch: channel,
	}
}

// Push добавляет задачу в очередь
func (q *Queue) Push(t *Task) {
	q.ch <- t
}

// PopWait - получение задачи из очереди
func (q *Queue) PopWait() *Task {
	// получаем задачу
	return <-q.ch
}

// Worker воркер
type Worker struct {
	id      int
	queue   *Queue
	deleter *Deleter
}

// NewWorker конструктор воркеров
func NewWorker(id int, queue *Queue, resizer *Deleter) *Worker {
	w := Worker{
		id:      id,
		queue:   queue,
		deleter: resizer,
	}
	return &w
}

// Loop -
func (w *Worker) Loop() {

	for {

		t := w.queue.PopWait()

		err := w.deleter.Delete(t.URLs)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker #%d deleted urls\n", w.id)

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
