package main

import (
	"fmt"
	"reflect"
)

/* One more caveeat of go ifaces
   by default interface{} is nil
   but if you assign a nil  object(struct) pointer to that iface
   iface will not longer be a nil, because even if both
   values are nil, but iface starts to have a link to
   beloved pointer of an object(struct)!!!
*/

type Man struct {
	Name string
}

type Eater interface{}

func main() {
	var man *Man // by default its nil
	fmt.Printf("[BEFORE] Here is initialization of pointer to the Man struct: %#v, reflect is: %#v\n", man, reflect.TypeOf(man).String())

	var eater Eater // by default its nil
	fmt.Printf("[BEFORE] Here is initialization of Eater interface: %#v\n", eater)
	if eater == nil {
		fmt.Printf("[BEFORE] eater equals to nil == %#v", eater)
	} else {
		fmt.Printf("[BEFORE] eater NOT equals to nil != %#v", eater)
	}

	fmt.Println("\n\nlets try to assign eater to man, is that nil := nil right???\n")
	// lets try to assign eater to man, is that nil := nil right???
	eater = man

	// Lets check what changed in our objects
	fmt.Printf("[AFTER] Here is pointer to the Man struct: %#v, reflect is: %#v\n", man, reflect.TypeOf(man).String())
	// NOW type of eater nil has changed to man-related !!! (*main.Man)(nil)
	fmt.Printf("[AFTER] Here is Eater interface: %#v\n", eater)
	if eater == nil {
		fmt.Printf("[AFTER] eater equals to nil == %#v", eater)
	} else {
		fmt.Printf("[AFTER] eater NOT equals to nil != %#v", eater)
	}

}
