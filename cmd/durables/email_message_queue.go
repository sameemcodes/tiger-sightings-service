package durable

import (
	"fmt"
	"sync"
	"time"
)

var MessageQueueVariable *MessageQueue

type MessageQueue struct {
	messages chan string
	wg       sync.WaitGroup
}

type WorkerPool struct {
	workers int
	queue   *MessageQueue
}

func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		messages: make(chan string, size),
	}
}

func NewWorkerPool(workers int, queue *MessageQueue) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		queue:   queue,
	}
}

func (wp *WorkerPool) StartWorkers() {
	for i := 0; i < wp.workers; i++ {
		wp.queue.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	defer wp.queue.wg.Done()
	for user_id := range wp.queue.messages {
		fmt.Println("Sending email to : ", user_id)
		time.Sleep(1 * time.Second) // Simulate email sending Vendor RateLimit
	}
}

func (mq *MessageQueue) Enqueue(message string) {
	mq.messages <- message
}

func (mq *MessageQueue) Dequeue() {
	close(mq.messages)
	mq.wg.Wait()
}
