package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case _, ok := <-snapshotTicker:
			if ok {
				takeSnapshot(logChan)
			}
		case _, ok := <-saveAfter:
			if ok {
				saveSnapshot(logChan)
				return
			}
		default:
			waitForData(logChan)
			duration := 500 * time.Millisecond
			time.Sleep(duration)
		}
	}
	/*
		for {
			select {
			case <-snapshotTicker:
				takeSnapshot(logChan)
			case <-saveAfter:
				saveSnapshot(logChan)
				return
			default:
				waitForData(logChan)
				time.Sleep(500 * time.Millisecond)
			}
		}
	*/
}

/*
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.
Complete the saveBackups function.
It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds.
After all, we want to show in the logs that the snapshot service is running.
*/
// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
