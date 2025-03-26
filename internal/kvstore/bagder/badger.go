package bagder

import (
	badger3 "github.com/dgraph-io/badger/v3"
	"log"
	"time"
)

type BadgerStore struct {
	db       *badger3.DB
	stopGC   chan struct{}
	gcTicker *time.Ticker
}

// NewBadgerStore creates a new BadgerStore instance.
func NewBadgerStore(path string, opts ...OptionFunc) (*BadgerStore, error) {
	cfg := &Config{
		BadgerOptions: badger3.DefaultOptions(path),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	db, err := badger3.Open(cfg.BadgerOptions)
	if err != nil {
		return nil, err
	}

	store := &BadgerStore{db: db}

	if cfg.GCInterval > 0 {
		store.startGC(cfg.GCInterval)
	}

	return store, nil
}

// startGC starts a background goroutine to run value log GC.
func (b *BadgerStore) startGC(interval time.Duration) {
	b.stopGC = make(chan struct{})
	b.gcTicker = time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-b.gcTicker.C:
				for {
					if err := b.db.RunValueLogGC(0.5); err != nil {
						break // No more GC work for now
					}
					log.Println("Badger GC cycle completed.")
				}
			case <-b.stopGC:
				log.Println("Stopping background Badger GC...")
				return
			}
		}
	}()
}

// Close closes the BadgerStore.
func (b *BadgerStore) Close() error {
	if b.stopGC != nil {
		close(b.stopGC)
		b.gcTicker.Stop()
	}
	return b.db.Close()
}

// Set stores a key-value pair.
func (b *BadgerStore) Set(key string, value []byte) error {
	return b.db.Update(func(txn *badger3.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

// Get retrieves a key from the store.
func (b *BadgerStore) Get(key string) ([]byte, error) {
	var valCopy []byte
	err := b.db.View(func(txn *badger3.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

// Delete removes a key from the store.
func (b *BadgerStore) Delete(key []byte) error {
	return b.db.Update(func(txn *badger3.Txn) error {
		return txn.Delete(key)
	})
}
