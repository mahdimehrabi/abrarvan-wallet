package interfaces

import "time"

type MemoryDB interface {
	//store in memory db
	Set(key string, value string, expiration time.Duration) error

	//get from memory db
	Get(key string) (string, error)
}
