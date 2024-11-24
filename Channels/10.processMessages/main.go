package main

import (
	"time"
)

func processMessages(messages []string) []string {
	ln := len(messages)
	result := make([]string, ln)
	processedMessagesCh := make(chan string, ln)

	defer close(processedMessagesCh)

	for _, message := range messages {
		go func(message string) {
			processedMessagesCh <- process(message)
		}(message)
	}

	for i := 0; i < ln; i++ {
		result[i] = <-processedMessagesCh
	}

	return result
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

/*
Textio needs to speed up message processing using concurrency.

Assignment
Implement the processMessages function, it takes a slice of messages. It should process each message concurrently
within a goroutine. Use the process function to simulate the benefit of using goroutines. Use a channel to ensure
that all messages are processed and collected correctly then return the slice of processed messages.

messages := []string{"Here are some messages", "Here is another", "and another"}
processedMessages := processMessages(messages)
fmt.Println(processedMessages)
// prints ["Here are some messages-processed", "Here is another-processed", "and another-processed"]
*/
/*
Solution
func processMessages(messages []string) []string {
	if len(messages) == 0 {
		return []string{}
	}

	ch := make(chan string, len(messages))

	for _, msg := range messages {
		go func(m string) {
			ch <- process(m)
		}(msg)
	}

	processedMessages := make([]string, len(messages))
	for i := 0; i < len(messages); i++ {
		processedMessages[i] = <-ch
	}

	return processedMessages
}
*/
