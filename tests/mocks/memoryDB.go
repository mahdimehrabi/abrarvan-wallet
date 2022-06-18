package mocks

import "time"

type MemoryDB struct {
	MockSetFn func(
		string,
		string,
		time.Duration,
	) error

	MockGetFn func(
		string,
	) (string, error)

	MockDecreaseConsumerCountFn func(code string) (consumerCount int, err error)
}

func (db *MemoryDB) Set(key string, value string, duration time.Duration) error {
	return db.MockSetFn(key, value, duration)
}

func (db *MemoryDB) Get(key string) (string, error) {
	return db.MockGetFn(key)
}

func (db *MemoryDB) DecreaseConsumerCount(code string) (consumerCount int, err error) {
	return db.MockDecreaseConsumerCountFn(code)
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		MockGetFn: func(s string) (string, error) {
			return "", nil
		},
		MockSetFn: func(s string, s2 string, duration time.Duration) error {
			return nil
		},
		MockDecreaseConsumerCountFn: func(code string) (consumerCount int, err error) {
			return 1, nil
		},
	}
}
