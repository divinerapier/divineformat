package divineformat

import (
	"strconv"
)

type JSONFormatter struct {
	buffer []byte
}

func NewJSONFormatter() Formatter {
	return defaultJSONFormatterPool.Acquire()
}

func (formatter *JSONFormatter) Str(key, value string) Formatter {
	return formatter.appendkey(key).appendstring(value)
}
func (formatter *JSONFormatter) Strs(key string, value ...string) Formatter {
	return formatter.appendkey(key).appendstrings(value...)
}
func (formatter *JSONFormatter) Int(key string, value int64) Formatter {
	return formatter.appendkey(key)
}
func (formatter *JSONFormatter) Ints(key string, values ...int64) Formatter {
	return formatter.appendkey(key).appendints(values...)
}
func (formatter *JSONFormatter) Bool(key string, value bool) Formatter {
	return formatter.appendkey(key).appendbool(value)
}
func (formatter *JSONFormatter) Bools(key string, values ...bool) Formatter {
	return formatter.appendkey(key).appendbools(values...)
}
func (formatter *JSONFormatter) Bytes(key string, value []byte) Formatter {
	return formatter.appendkey(key).appendbytes(value)
}
func (formatter *JSONFormatter) Flush() []byte {
	length := len(formatter.buffer)
	if length == 0 {
		return nil
	}
	if formatter.buffer[length-1] != '}' {
		formatter.buffer = append(formatter.buffer, '}')
	}
	return formatter.buffer
}

func (formatter *JSONFormatter) Release() {
	defaultJSONFormatterPool.Release(formatter)
}

func (formatter *JSONFormatter) appendkey(key string) *JSONFormatter {
	length := len(formatter.buffer)

	// empty buffer
	if length == 0 {
		formatter.buffer = append(formatter.buffer, '{', '"')
	} else {
		formatter.buffer = append(formatter.buffer, ',', '"')
	}

	formatter.buffer = append(formatter.buffer, key...)
	formatter.buffer = append(formatter.buffer, '"', ':')
	return formatter
}

func (formatter *JSONFormatter) appendstring(value string) *JSONFormatter {
	formatter.buffer = append(formatter.buffer, '"')
	formatter.buffer = append(formatter.buffer, value...)
	formatter.buffer = append(formatter.buffer, '"')
	return formatter
}

func (formatter *JSONFormatter) appendstrings(value ...string) *JSONFormatter {
	length := len(value)
	if length == 0 {
		formatter.buffer = append(formatter.buffer, '[', ']')
		return formatter
	}

	formatter.buffer = append(formatter.buffer, '[')
	formatter.appendstring(value[0])

	for i := 1; i < length; i++ {
		formatter.buffer = append(formatter.buffer, ',')
		formatter.appendstring(value[i])
	}
	formatter.buffer = append(formatter.buffer, ']')
	return formatter
}

func (formatter *JSONFormatter) appendint(value int64) *JSONFormatter {
	formatter.buffer = strconv.AppendInt(formatter.buffer, value, 10)
	return formatter
}

func (formatter *JSONFormatter) appendints(values ...int64) *JSONFormatter {
	length := len(values)
	if length == 0 {
		formatter.buffer = append(formatter.buffer, '[', ']')
		return formatter
	}

	formatter.buffer = append(formatter.buffer, '[')
	formatter.appendint(values[0])

	for i := 1; i < length; i++ {
		formatter.buffer = strconv.AppendInt(append(formatter.buffer, ','), values[i], 10)
	}
	formatter.buffer = append(formatter.buffer, ']')
	return formatter
}

func (formatter *JSONFormatter) appendbool(value bool) *JSONFormatter {
	formatter.buffer = strconv.AppendBool(formatter.buffer, value)
	return formatter
}

func (formatter *JSONFormatter) appendbools(values ...bool) *JSONFormatter {
	length := len(values)
	if length == 0 {
		formatter.buffer = append(formatter.buffer, '[', ']')
		return formatter
	}

	formatter.buffer = append(formatter.buffer, '[')
	formatter.appendbool(values[0])

	for i := 1; i < length; i++ {
		formatter.buffer = strconv.AppendBool(append(formatter.buffer, ','), values[i])
	}
	formatter.buffer = append(formatter.buffer, ']')
	return formatter
}

func (formatter *JSONFormatter) appendbytes(value []byte) *JSONFormatter {
	if len(value) == 0 {
		formatter.buffer = append(formatter.buffer, `"nil"`...)
		return formatter
	}
	formatter.buffer = append(formatter.buffer, '"')
	formatter.buffer = append(formatter.buffer, value...)
	formatter.buffer = append(formatter.buffer, '"')
	return formatter
}

func (formatter *JSONFormatter) appendopencurly() *JSONFormatter {

	return formatter
}

func (formatter *JSONFormatter) appendclosecurly() *JSONFormatter {
	return formatter
}
