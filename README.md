# Go experiment: Matrix determinant
single/single.go find the determinant of a square matrix in a single thread, giving the number of row and columns
goroutine/goroutines.go do the same thing but with goroutine instead

This code was made with completely experimental intentions and is not intended to be used in production at any level.

To prove it, execute the folowing instructions:
```
go run single/single.go
go run goroutine/goroutines.go
```

In GNU/Linux systems you can see the execution time with 'time' command:
```
time go run single/single.go
time go run goroutine/goroutines.go
```