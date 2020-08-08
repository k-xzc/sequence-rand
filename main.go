package main

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Simple CLI to shuffle number from range of two given number ")
	fmt.Print("Enter first number : ")
	firstNum := CheckInput()
	fmt.Print("Enter second number : ")
	secondNum := CheckInput()
	minNum , maxNum := CheckMinMaxNum(firstNum,secondNum)
	orderNum := CreateOrderNumber(CheckMinMaxNum(minNum,maxNum))
	shuffleNum := CreateShuffleNumber(orderNum)
	fmt.Println("The order of range of two given number : ")
	fmt.Println(orderNum)
	fmt.Println("The shuffle of range of two given number : ")
	fmt.Println(shuffleNum)
}

func CheckInput()int{
	var i int
	var checker string
	for {
		_, err := fmt.Scan(&checker)
		i, err = strconv.Atoi(checker)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else {
			break
		}
	}
	return i
}

func CheckMinMaxNum(minNum int, maxNum int) (int,int){
	m := minNum
	if m > maxNum {
		minNum = maxNum
		maxNum = m
	}
	return minNum,maxNum
}

func CreateOrderNumber(minNum int,maxNum int)[]int {
	var orderNum []int
	for i:=0 ; i <= maxNum-minNum ; i++{
		orderNum=append(orderNum,i+minNum)
	}
	return orderNum
}

func CreateShuffleNumber(orderNum []int)[]int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var shuffleNum []int
	for _, i := range r.Perm(len(orderNum)){
		shuffleNum=append(shuffleNum,orderNum[i])
	}
	return shuffleNum
}
