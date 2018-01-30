package example

import (
	"fmt"

	"github.com/divinerapier/divineformat"
)

func Example() {
	formatter := divineformat.NewJSONFormatter()

	formatter.Str("string0", "value0").
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
		Bools("bools0", 1 == 1, 1 == 0, 1 == 1, 1 == 2)

	fmt.Println(string(formatter.Bytes()) == f())

	// Output: true
}

const tpl = `{"%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":"%v","%v":["%v","%v","%v","%v","%v"],"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":%v,"%v":[%v,%v,%v,%v]}`

func f() string {
	return fmt.Sprintf(tpl, "string0", "value0", "string1", "value1", "string2", "value2", "string3", "value3", "string4", "value4", "strings0", "v0", "v1", "v2", "v3", "v4", "bool0", false, "bool1", true, "bool2", false, "bool3", true, "bool4", false, "bools0", true, false, true, false)
}
