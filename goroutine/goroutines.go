package main

import (
	"fmt"
	"math"
	"math/rand"
)

type matriz struct {
	number int
	values [][]int
}

func main() {
	matrix := squareMatrix(11)
	/* fmt.Println("Matriz:")
	for _, row := range matrix {
		fmt.Println(row)
	} */
	n := make(chan int)
	go getDeterminant(matrix, n)
	y := <-n
	fmt.Println(y)
}

func getDeterminant(m [][]int, n chan int) {
	matrix := matriz{
		number: 1,
		values: m,
	}
	go calculateSubDeterminant(matrix, n)
}

func calculateSubDeterminant(m matriz, resolution chan int) {
	if len(m.values) == 2 {
		var tot int
		tot = (m.values[0][0] * m.values[1][1])
		tot -= (m.values[0][1] * m.values[1][0])
		tot = tot * m.number
		resolution <- tot
	} else {
		var determinants []int
		var subMatrix matriz
		v := m.values[0]
		subDeterminant := make(chan int, len(v))
		for i := range v {
			number, values := reduce(m.values, 1, i+1)
			subMatrix.number = number * m.number
			subMatrix.values = values
			go calculateSubDeterminant(subMatrix, subDeterminant)
		}
		for i := 0; i < len(v); i++ {
			determinants = append(determinants, <-subDeterminant)
		}
		sum := 0
		for _, determinant := range determinants {
			sum += determinant
		}
		resolution <- sum
	}
}

func reduce(m [][]int, i int, j int) (int, [][]int) {
	s := [][]int{}
	e := m[i-1][j-1] * int(math.Pow(-1, float64(i+j)))
	s = append(s, m[:i-1]...)
	s = append(s, m[i:]...)
	for i, v := range s {
		t := []int{}
		t = append(t, v[:j-1]...)
		t = append(t, v[j:]...)
		s[i] = t
	}
	return e, s
}
func newMatrix(i int, j int) [][]int {
	m := [][]int{}
	f := []int{}
	for ii := 0; ii < i; ii++ {
		f = []int{}
		for jj := 0; jj < j; jj++ {
			f = append(f, rand.Intn(30))
		}
		m = append(m, f)
	}
	return m
}

func squareMatrix(n int) [][]int {
	return newMatrix(n, n)
}
