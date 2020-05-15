# Function receiver
* SYNTAX: func [receiver] funcName(){}
```
// Any variable of type "deck" now has access to the "print" method
// You can think of 'd' as 'this' or 'self' in OOP languages like Java or Python
- EXAMPLE: func (d deck) print() {
    fmt.Println (deck)
}
```

### Why don't we have to import a function from other files?
- Because it's under the same package

# Pointers

- &variable - Gives a memory address of the value this variable is pointing to
- *pointer - Gives the value this memory address is pointing at

### Example and explanation
Assume `person` is a struct
```go
func (pointerToPerson *person) updateName() {
    *pointerToPerson
}
```
*person - This is a type description it means we're working with a poitner to a person
*pointerToPerson - This is an operator, it means we want to manipulate the value the pointer is referencing

### Conversion
- Turn `address` into `value` with `*address`
- Turn `value` into `address` with `&value`
