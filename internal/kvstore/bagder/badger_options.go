package bagder

import (
	badger3 "github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
	"time"
)

type Config struct {
	BadgerOptions badger3.Options
	GCInterval    time.Duration
}

type OptionFunc func(*Config)

func WithCompression(c options.CompressionType) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.Compression = c
	}
}

func WithZSTDLevel(level int) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.ZSTDCompressionLevel = level
	}
}

func WithInMemory() OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.InMemory = true
	}
}

func WithGCInterval(interval time.Duration) OptionFunc {
	return func(cfg *Config) {
		cfg.GCInterval = interval
	}
}

type gcConfig struct {
	enabled  bool
	interval time.Duration
}

func WithAutoGC(interval time.Duration) OptionFunc {
	return func(cfg *Config) {
		cfg.GCInterval = interval
	}
}

func WithDir(dir string) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.Dir = dir
		cfg.BadgerOptions.ValueDir = dir
	}
}

//func WithGCDiscardRatio(ratio float64) OptionFunc {
//	return func(cfg *Config) {
//		cfg.BadgerOptions.ValueDiscardRatio = ratio
//	}
//}

func WithSyncWrites(sync bool) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.SyncWrites = sync
	}
}

func WithEncryptionKey(key []byte, rotation time.Duration) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.EncryptionKey = key
		cfg.BadgerOptions.EncryptionKeyRotationDuration = rotation
	}
}

func WithLogger(logger badger3.Logger) OptionFunc {
	return func(cfg *Config) {
		cfg.BadgerOptions.Logger = logger
	}
}
