### Solution

```
$ go run examples11/main.go
Expression                          |Value              |Nil?   |Type   |Length  |Capacity
a := []int{0, 1, 2, 3, 4, 5, 6, 7}  |[0 1 2 3 4 5 6 7]  |false  |[]int  |8       |8
b := a[2:6]                         |[2 3 4 5]          |false  |[]int  |4       |6
c := a[4:]                          |[4 5 6 7]          |false  |[]int  |4       |4
d := a[:]                           |[0 1 2 3 4 5 6 7]  |false  |[]int  |8       |8
e := c[1:]                          |[5 6 7]            |false  |[]int  |3       |3
```
