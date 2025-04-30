// the go extension on vs code
package main

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool(true)

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	//Type(value)

	g := "Hello @"
	var h []byte
	h = []byte(g)
	fmt.Println(h)

	i := []byte{255, 120,72} // can only take uft-8. means not over 255
	j := string(i)
	fmt.Println(j)

}
