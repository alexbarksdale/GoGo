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
