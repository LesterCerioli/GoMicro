package memory

import (
	"github.com/chinahtl/go-micro/v3/debug/log"
)

type logStream struct {
	stream <-chan log.Record
	stop   chan bool
}

func (l *logStream) Chan() <-chan log.Record {
	return l.stream
}

func (l *logStream) Stop() error {
	select {
	case <-l.stop:
		return nil
	default:
		close(l.stop)
	}
	return nil
}
