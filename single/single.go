package main

import (
	"fmt"
	"math"
	"math/rand"
)

type matrixStruct struct {
	number int
	values [][]int
}

func main() {
	matrix := squareMatrix(11)
	/* fmt.Println("Matriz:")
	for _, row := range matrix {
		fmt.Println(row)
	} */
	n := getDeterminant(matrix)
	fmt.Println(n)
}

func getDeterminant(m [][]int) int {
	matrix := matrixStruct{
		number: 1,
		values: m,
	}
	return calculateSubDeterminant(matrix)
}

func calculateSubDeterminant(m matrixStruct) int {
	if len(m.values) == 2 {
		var tot int
		tot = (m.values[0][0] * m.values[1][1])
		tot -= (m.values[0][1] * m.values[1][0])
		tot = tot * m.number
		return tot
	}

	var determinants []int
	var subMatrix matrixStruct
	v := m.values[0]
	for i := range v {
		number, values := reduce(m.values, 1, i+1)
		subMatrix.number = number * m.number
		subMatrix.values = values
		subDeterminant := calculateSubDeterminant(subMatrix)
		determinants = append(determinants, subDeterminant)
	}
	sum := 0

	for _, determinant := range determinants {
		sum += determinant
	}
	return sum
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
