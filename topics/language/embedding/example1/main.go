package main

type Animal struct{}

type Dog struct {
	Animal
}

func main() {
	var a Animal
	a = Dog{} // ./main.go:11:4: cannot use Dog literal (type Dog) as type Animal in assignment
}
