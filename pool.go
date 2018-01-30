package divineformat

import (
	"sync"
)

type (
	bytespool struct {
		sync.Pool
	}
	jsonformatterpool struct {
		sync.Pool
	}
)

var (
	defaultBytesPool         bytespool
	defaultJSONFormatterPool jsonformatterpool
)

func init() {
	defaultBytesPool.New = func() interface{} {
		return make([]byte, 0, 512)
	}
	defaultJSONFormatterPool.New = func() interface{} {
		return &JSONFormatter{
			buffer: defaultBytesPool.Acquire(),
		}
	}
}

func (p *bytespool) Acquire() []byte {
	return p.Get().([]byte)
}
func (p *bytespool) Release(buffer []byte) {
	if p == nil || buffer == nil {
		return
	}
	buffer = buffer[:0]
	p.Put(buffer)
}

func (p *jsonformatterpool) Acquire() *JSONFormatter {
	// formatter := p.Get().(*JSONFormatter)
	// formatter.buffer = defaultBytesPool.Acquire()
	// return formatter
	return p.Get().(*JSONFormatter)
}

func (p *jsonformatterpool) Release(formatter *JSONFormatter) {
	if p == nil || formatter == nil {
		return
	}
	// defaultBytesPool.Release(formatter.buffer)
	formatter.buffer = formatter.buffer[:0]
	p.Put(formatter)
}
