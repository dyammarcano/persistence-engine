package bagder

import (
	"github.com/inovacc/ksuid"
	"log/slog"
	"testing"
)

func TestNewBadgerStore(t *testing.T) {
	store, err := NewBadgerStore("../../../badger", WithLogger(NewSlogLogger(slog.LevelDebug)))
	if err != nil {
		t.Errorf("Error creating BadgerStore: %v", err)
		return
	}
	defer func(store *BadgerStore) {
		_ = store.Close()
	}(store)

	key := ksuid.NewString()

	if err := store.Set(key, []byte("value")); err != nil {
		t.Errorf("Error setting key: %v", err)
		return
	}

	value, err := store.Get(key)
	if err != nil {
		t.Errorf("Error getting key: %v", err)
		return
	}

	if string(value) != "value" {
		t.Errorf("Expected value 'value', got '%s'", value)
		return
	}
}
