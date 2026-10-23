package middlewares

import (
	"log"
	"os"

	"github.com/Shinjon101/go-cache/store"
)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[log]", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("GET %q", key)
	val, err := l.inner.Get(key)
	if err != nil {
		l.logger.Printf("GET %q -> miss (%v)", key, err)
	} else {
		l.logger.Printf("GET %q -> hit", key)
	}
	return val, err
}

func (l *LoggingMiddleware) Set(key string, val string) error {
	l.logger.Printf("SET %q", key)
	return l.inner.Set(key, val)
}

func (l *LoggingMiddleware) Keys() []string {
	l.logger.Printf("KEYS")
	return l.inner.Keys()
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("DELETE %q", key)
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Len() int {
	l.logger.Printf("LEN")
	return l.inner.Len()
}
