package main

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func main() {
	//vectors()
	matrix()
}

func vectors() {
	//Create 2 vectors
	vectorA := mat.NewVecDense(3, []float64{1, 2, 3})
	vectorB := mat.NewVecDense(3, []float64{4, 5, 6})
	// dot product (1 * 4) + (2 * 5) + (3 * 6) = 32
	dotProduct := mat.Dot(vectorA, vectorB)
	fmt.Println("the dot product is", dotProduct)

	//norm
	//norm is the square root of the sum of all the values at power2
	//for vector A is sqrt(1^2 + 2^2 + 3^2) = sqrt(14) = 3.74
	norm := blas64.Nrm2(vectorA.RawVector())
	fmt.Println("the norm is", norm)
}

func matrix() {
	// matrix
	matA := mat.NewDense(2, 2, []float64{2, 4, 5, 6})
	matB := mat.NewDense(2, 2, []float64{4, 5, 9, 10})

	matAA := mat.NewDense(2, 2, []float64{2, 4, 5, 6})
	matBB := mat.NewDense(2, 2, []float64{4, 5, 9, 10})

	//sum
	matA.Add(matA, matB)
	formatter := mat.Formatted(matA, mat.Prefix(""))
	fmt.Printf("%v\n", formatter)

	matBB.Add(matBB, matAA)
	formatter_b := mat.Formatted(matBB, mat.Prefix(""))
	fmt.Printf("%v\n", formatter_b)

	sales := mat.NewDense(3, 3, []float64{1,3,9,2,4,6,1,5,0})
	fmt.Println(sales.At(2,2))
	fmt.Println(sales.IsEmpty())

	formatterSales := mat.Formatted(sales, mat.Prefix(""))
	fmt.Printf("%v\n",formatterSales)
}
