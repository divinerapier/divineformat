package divineformat

import (
	"fmt"
	"testing"
)

func J0(b *testing.B) {
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

	l := len(formatter.Flush())
	formatter.Release()

	b.SetBytes(int64(l))
}

func J1(b *testing.B) {
	formatter := NewJSONFormatter()
	formatter = formatter.Str("string0", "value0")
	formatter = formatter.Str("string1", "value1")
	formatter = formatter.Str("string2", "value2")
	formatter = formatter.Str("string3", "value3")
	formatter = formatter.Str("string4", "value4")
	formatter = formatter.Strs("strings0", "v0", "v1", "v2", "v3", "v4")
	formatter = formatter.Bool("bool0", 1 == 0)
	formatter = formatter.Bool("bool1", 1 == 1)
	formatter = formatter.Bool("bool2", 1 == 0)
	formatter = formatter.Bool("bool3", 1 == 1)
	formatter = formatter.Bool("bool4", 1 == 0)
	formatter = formatter.Bools("bools0", 1 == 1, 1 == 0, 1 == 1, 1 == 2)
	formatter = formatter.Int("int0", 0)
	formatter = formatter.Int("int1", 1)
	formatter = formatter.Int("int2", 2)
	formatter = formatter.Int("int3", 3)
	formatter = formatter.Int("int4", 4)
	formatter = formatter.Ints("ints0", 0, 1, 2, 3, 4)

	l := len(formatter.Flush())
	formatter.Release()

	b.SetBytes(int64(l))
}

const tpl = `{"%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":["%v","%v","%v","%v","%v"],"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":[%v,%v,%v,%v],"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":[%v,%v,%v,%v,%v]}`

func F0(b *testing.B) string {
	s := fmt.Sprintf(tpl, "string0", "value0", "string1", "value1", "string2", "value2", "string3", "value3", "string4", "value4", "strings0", "v0", "v1", "v2", "v3", "v4", "bool0", false, "bool1", true, "bool2", false, "bool3", true, "bool4", false, "bools0", true, false, true, false, "int0", 0, "int1", 1, "int2", 2, "int3", 3, "int4", 4, "ints0", 0, 1, 2, 3, 4)

	b.SetBytes(int64(len(s)))
	return s
}

func F1(b *testing.B) string {
	s := fmt.Sprintf(`{"%v":"`, "string0")
	s += fmt.Sprintf(`%v"`, "value0")
	s += fmt.Sprintf(`,"%v":"`, "string1")
	s += fmt.Sprintf(`%v","`, "value1")
	s += fmt.Sprintf(`%v`, "string2")
	s += fmt.Sprintf(`":"`)
	s += fmt.Sprintf(`%v","`, "value2")
	s += fmt.Sprintf(`%v`, "string3")
	s += fmt.Sprintf(`":"`)
	s += fmt.Sprintf(`%v","`, "value3")
	s += fmt.Sprintf(`%v":"`, "string4")
	s += fmt.Sprintf(`%v","`, "value4")
	s += fmt.Sprintf(`%v":["`, "strings0")
	s += fmt.Sprintf(`%v","`, "v0")
	s += fmt.Sprintf(`%v","`, "v1")
	s += fmt.Sprintf(`%v","`, "v2")
	s += fmt.Sprintf(`%v","`, "v3")
	s += fmt.Sprintf(`%v"],"`, "v4")
	s += fmt.Sprintf(`%v`, "bool0")
	s += fmt.Sprintf(`":%v,`, false)
	s += fmt.Sprintf(`"%v":`, "bool1")
	s += fmt.Sprintf(`%v,"`, true)
	s += fmt.Sprintf(`%v":`, "bool2")
	s += fmt.Sprintf(`%v,"`, false)
	s += fmt.Sprintf(`%v":`, "bool3")
	s += fmt.Sprintf(`%v,"`, true)
	s += fmt.Sprintf(`%v":`, "bool4")
	s += fmt.Sprintf(`%v,"`, false)
	s += fmt.Sprintf(`%v":[`, "bools0")
	s += fmt.Sprintf(`%v,`, true)
	s += fmt.Sprintf(`%v,`, false)
	s += fmt.Sprintf(`%v,`, true)
	s += fmt.Sprintf(`%v],"`, false)
	s += fmt.Sprintf(`%v":`, "int0")
	s += fmt.Sprintf(`%v,"`, 0)
	s += fmt.Sprintf(`%v":`, "int1")
	s += fmt.Sprintf(`%v,"`, 1)
	s += fmt.Sprintf(`%v":`, "int2")
	s += fmt.Sprintf(`%v,"`, 2)
	s += fmt.Sprintf(`%v":`, "int3")
	s += fmt.Sprintf(`%v,"`, 3)
	s += fmt.Sprintf(`%v":`, "int4")
	s += fmt.Sprintf(`%v,"`, 4)
	s += fmt.Sprintf(`%v":[`, "ints0")
	s += fmt.Sprintf(`%v,`, 0)
	s += fmt.Sprintf(`%v,`, 1)
	s += fmt.Sprintf(`%v,`, 2)
	s += fmt.Sprintf(`%v,`, 3)
	s += fmt.Sprintf(`%v]}`, 4)

	b.SetBytes(int64(len(s)))
	return s
}

func BenchmarkJSON_Chain(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			J0(b)
		}
	})
}

func BenchmarkJSON_ManyTimes(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			J1(b)
		}
	})
}

func BenchmarkSTD_Once(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			F0(b)
		}
	})
}
func BenchmarkSTD_ManyTimes(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			F1(b)
		}
	})
}
