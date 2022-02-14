package main

import "fmt"

// Finder shows all methods that helps to find something
type Finder interface {
	FindByName(string) (string, error)
	FindByID(int) (string, error)
}

// Detective is a person that theoretically can find something
type Detective struct {
	Name string
}

// FindByName is a method from interface Finder, that Detective supports ;)
func (d Detective) FindByName(s string) (string, error) {
	return "", nil
}

/* Here is magic!
here we checks that Detective struct, implements all required methods of
interface Finder , otherwise it should trace!
so its kinda simle reminder to dev's that something not implemented!
for example w/o FindByID implementation it will trace with

```
./statichcheck.go:26:5: cannot use (*Detective)(nil) (type *Detective) as type Finder in assignment:
	*Detective does not implement Finder (missing FindByID method)
```
*/
var _ Finder = (*Detective)(nil)

func main() {
	fmt.Println("vim-go")
}
