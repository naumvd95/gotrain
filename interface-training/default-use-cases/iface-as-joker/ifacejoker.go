package main

import "fmt"

/* Aaaand 1 more iface feature
   You can use empty interface as joker!
   as you know Go is strict-type lang, but
   interface{} may be a slight exclude of that rule ;)
   BUT: best practice: use this feature use it as little as possible!!!
*/

func main() {

	// hint: better to use struct{} here ;)
	nameAgeDummyStructure := map[string]interface{}{}

	// below code will be compiled good
	nameAgeDummyStructure["Anny"] = 11
	nameAgeDummyStructure["Petr"] = 12.5
	nameAgeDummyStructure["John"] = "one hundred pounds"

	fmt.Printf("Here is our mixed structure %v\n", nameAgeDummyStructure)

	// BUT keep in mind that in such case Anny: 11, where  11 is not int, but interface type!
	// it means you cant +1 it for example, you need to override type first!

	// below will not work: (mismatched types interface {} and int)
	//nameAgeDummyStructure["Anny"] += 1
	// below will work
	nameAgeDummyStructure["Anny"] = nameAgeDummyStructure["Anny"].(int) + 1

	fmt.Printf("Here is our updated structure %v\n", nameAgeDummyStructure)
}
