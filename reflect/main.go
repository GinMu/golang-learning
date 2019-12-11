package main

import (
	"fmt"
	"reflect"
)

// T struct
type T struct {
	A int
	B string
}

// User struct
type User struct {
	id   int
	Name string
	Age  int
}

func main() {
	t := T{A: 1, B: "Libre"}
	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeofT.Field(i).Name, f.Type(), f.Interface())
	}

	u := User{1, "Libre", 27}
	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)
	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())
	fmt.Printf("%v\n", vb)
	name := "Hello Libre"
	vc := reflect.ValueOf(name)
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n", vb)
}
