package main

func main() {
	v := [6]int{0, 1, 2, 3, 4, 5}
	w := v[:]

	w = w[:4]
	w = w[:0]

	// Question: Before you hit run! What do the following two lines result in?
	w = w[1:]
	w = w[1:3]
	// fmt.Printf("%v - len: %d, cap: %d\n", w, len(w), cap(w))
}
