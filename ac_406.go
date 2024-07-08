// Type checking
package main

import (
	"fmt"
)


func checkType(v interface{}) (string, error) {
	switch t := v.(type) {
	case int:
		return fmt.Sprintf("%v is int", t), nil
	case float32, float64:
		return fmt.Sprintf("%v is float", t), nil
	case string:
		return fmt.Sprintf("%v is string", t), nil
	case bool:
		return fmt.Sprintf("%v is bool", t), nil
	default:
		return fmt.Sprintf("%v is unknown", t), nil
	}
}

func main() {
	/* 
	if len(os.Args) < 2 {
		fmt.Println("must enter a variable!")
		os.Exit(1)
	} 
	vartype := os.Args[1]
	*/
	inttype := 5
	res, _ := checkType(inttype)
	fmt.Println(res)

	floattype := 3.14
	res, _ = checkType(floattype)
	fmt.Println(res)

	strtype := "hello"
	res, _ = checkType(strtype)
	fmt.Println(res)

	booltype := true
	res, _ = checkType(booltype)
	fmt.Println(res)

	structtype := struct {
		x int
		y float64
	} {}
	res, _ = checkType(structtype)
	fmt.Println(res)
}
