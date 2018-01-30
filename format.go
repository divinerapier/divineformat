package divineformat

type Formatter interface {
	Str(key, value string) Formatter
	Strs(key string, values ...string) Formatter

	Int(key string, value int64) Formatter
	Ints(key string, values ...int64) Formatter
	Bool(key string, value bool) Formatter
	Bools(key string, values ...bool) Formatter

	Bytes() []byte

	Release()
}

func NewFormatter(t string) Formatter {
	switch t {
	case "json":
		return NewJSONFormatter()
	}
	return NewJSONFormatter()
}
