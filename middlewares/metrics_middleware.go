package middlewares

import (
	"fmt"
	"time"

	"github.com/Shinjon101/go-cache/store"
)

type MetricsMiddleware struct {
	inner store.Storer

	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int

	getMisses int

	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{
		inner: inner,
	}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()

	val, err := m.inner.Get(key)
	m.getCalls++
	if err != nil {
		m.getMisses++
	}

	m.totalGetLatency += time.Since(start)

	return val, err
}

func (m *MetricsMiddleware) Set(key string, val string) error {
	start := time.Now()

	err := m.inner.Set(key, val)
	m.setCalls++

	m.totalSetLatency += time.Since(start)

	return err
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Delete(key string) {
	m.inner.Delete(key)
	m.deleteCalls++
}

func (m *MetricsMiddleware) Len() int {
	n := m.inner.Len()
	m.lenCalls++
	return n
}

func (m *MetricsMiddleware) Report() {
	fmt.Printf("GET calls: %d (missed: %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("SET calls: %d\n", m.setCalls)
	fmt.Printf("DELETE calls: %d\n", m.deleteCalls)
	fmt.Printf("LEN calls: %d\n", m.lenCalls)

	if m.getCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.getCalls)
		fmt.Printf("Avg Get Latency : %v \n", avg)
	}

	if m.setCalls > 0 {
		avg := m.totalSetLatency / time.Duration(m.setCalls)
		fmt.Printf("Avg Set Latency : %v \n", avg)
	}

}
