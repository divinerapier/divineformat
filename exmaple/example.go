package main

import "fmt"

func main() {
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
	fmt.Println(s)
}
