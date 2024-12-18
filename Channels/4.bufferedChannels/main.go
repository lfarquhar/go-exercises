/*
We want to be able to send emails in batches. A writing goroutine will write an entire
batch of email messages to a buffered channel, and later, once the channel is full,
a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough

	to store all of the emails it's given. It should then write the emails to the channel in order, and
	finally return the channel.
*/
package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {

	ln := len(emails)

	fmt.Printf("length of array %v", ln)
	fmt.Println("")

	ch := make(chan string, ln)

	for _, email := range emails {
		ch <- email
	}

	return ch
}
