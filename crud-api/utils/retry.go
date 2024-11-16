package utils

import (
	"log"
	"time"
)

func Retry(operation func() error) error {
	const maxRetries = 3
	const retryDelay = 100 * time.Millisecond

	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		log.Printf("Retry %d/%d failed: %v\n", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}
	return err
}
