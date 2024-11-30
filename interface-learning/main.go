package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

func main() {
	Greeting("Sri")
	Greeting("")
}

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func WriteGreeting(name string, w Writer) error {
	greeting := Greeting(name)
	_, err := w.Write([]byte(greeting))
	if err != nil {
		return err
	}
	return nil
}

// the methods defined in interface , shoudl be implemented by Obsjects , if not then that Object cannot be used in interface

// type BigInterface interface {
// 	A()
// 	B()
// 	C()
// }

// type Demo struct {
// }

// func (Demo) A() {

// }
// func (Demo) B() {

// }

// func Test(bi BigInterface) {}

// func main() {
// 	var demo Demo
// 	Test(demo)

// }
