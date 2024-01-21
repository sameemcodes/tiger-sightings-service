package tests

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	durable "tigerhall-kittens/cmd/durables"
)

// Initialize MessageQueueVariable
func init() {
	// Initialize MessageQueueVariable
	mq := durable.NewMessageQueue(10)
	durable.MessageQueueVariable = mq
	fmt.Println("MessageQueueVariable Init: ", durable.MessageQueueVariable)
	workerPool := durable.NewWorkerPool(10, durable.MessageQueueVariable)
	workerPool.StartWorkers()
	fmt.Println("MessageQueueVariable: ", durable.MessageQueueVariable)
}

// TestEmailMessageQueue tests the email message queue
// delay with batches of 10 messages
func TestEmailMessageQueue(t *testing.T) {
	testing.Init()

	fmt.Println("MessageQueueVariable TestEmailMessageQueue: ", durable.MessageQueueVariable)

	// Enqueue messages
	for i := 0; i < 100; i++ {
		var message = "message " + strconv.Itoa(i)
		fmt.Println("MessageQueueVariable Enqueue: ", message)
		durable.MessageQueueVariable.Enqueue(message)
		time.Sleep(time.Millisecond)
	}

	defer durable.MessageQueueVariable.Dequeue() // Close the message queue
}
