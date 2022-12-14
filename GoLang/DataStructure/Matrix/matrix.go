package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UTC().Unix())
	var size int
	fmt.Println(" Enter the size of Square matrix :: ")
	fmt.Scanln(&size)
	m1 := populateMatrix(size)
	m2 := populateMatrix(size)
	fmt.Println(" matrix one :: ")
	printMatrix(m1)
	fmt.Println(" matrix two :: ")
	printMatrix(m2)

	fmt.Println(" matrix addition :: ")
	printMatrix(AddMatrix(m1,m2))
	fmt.Println(" matrix Subtraction :: ")
	printMatrix(SubMatrix(m1,m2))
	fmt.Println(" matrix Multiplication :: ")
	printMatrix(MultiMatrix(m1,m2))
}

func populateMatrix(size int) [][] int{
	fmt.Println(" matrix size :: ",size)
	matrix := make([][] int, size)

	for i:=0;i< size; i++{
		matrix[i] = make([]int,size)
		for j:=0;j<len(matrix[i]);j++{
			//fmt.Println(" matrix indices :: ",i,j)
			matrix[i][j] = rand.Intn(10)
		}
	}
	return matrix
}

func AddMatrix(m1[][] int, m2[][] int)[][]int{
	result := make([][]int, len(m1))
	for i,a := range m1{
		result[i] = make([]int, len(m1[i]))
		for j, _ := range a{
			result[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return result
}


func SubMatrix(m1[][]int,m2 [][] int)[][] int{

	result := make([][]int, len(m2))

	for i,a:= range m2{
		for j,_ := range a{
			result[i] = append(result[i], m1[i][j] - m2[i][j])
		}
	}
	return result
}

func MultiMatrix(m1[][]int, m2[][]int)[][]int{
	result := make([][]int, len(m1))
	//fmt.Println("multi result :: ",result)
	for i:=0;i< len(m1);i++{
		result[i] = make([]int, len(m1[i]))
		//fmt.Println("multi result[i] :: ",result)
		for j:= 0;j<len(m2);j++{
			for k:=0;k<len(m2);k++{
				result[i][j] += (m1[i][k] * m2[j][k])
			}
			
		}
	}
	return result
}

func printMatrix(mat [][]int){

	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			fmt.Printf("%d ",mat[i][j])
		}
		fmt.Println("")
	}
}