package divineformat

import (
	"fmt"
	"testing"
)

func J() {
	formatter := NewJSONFormatter().
		Str("string0", "value0").
		Str("string1", "value1").
		Str("string2", "value2").
		Str("string3", "value3").
		Str("string4", "value4").
		Strs("strings0", "v0", "v1", "v2", "v3", "v4").
		Bool("bool0", 1 == 0).
		Bool("bool1", 1 == 1).
		Bool("bool2", 1 == 0).
		Bool("bool3", 1 == 1).
		Bool("bool4", 1 == 0).
		Bools("bools0", 1 == 1, 1 == 0, 1 == 1, 1 == 2).
		Int("int0", 0).
		Int("int1", 1).
		Int("int2", 2).
		Int("int3", 3).
		Int("int4", 4).
		Ints("ints0", 0, 1, 2, 3, 4)

	formatter.Bytes()
	formatter.Release()

	// fmt.Println(formatter.OutputString())
}

const tpl = `{"%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":["%v","%v","%v","%v","%v"],"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":[%v,%v,%v,%v],"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":[%v,%v,%v,%v,%v]}`

func F() string {
	return fmt.Sprintf(tpl, "string0", "value0", "string1", "value1", "string2", "value2", "string3", "value3", "string4", "value4", "strings0", "v0", "v1", "v2", "v3", "v4", "bool0", false, "bool1", true, "bool2", false, "bool3", true, "bool4", false, "bools0", true, false, true, false, "int0", 0, "int1", 1, "int2", 2, "int3", 3, "int4", 4, "ints0", 0, 1, 2, 3, 4)
}

func BenchmarkJSON(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			J()
		}
	})
}

func BenchmarkSTD(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			F()
		}
	})
}
