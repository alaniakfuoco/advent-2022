package day20

import (
	"advent/util"
	"fmt"
	"strconv"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	file := NewLinkedList()
	moveOrder := []int{}
	for _, line := range data {
		val, _ := strconv.Atoi(line)
		file.Add(val)
		moveOrder = append(moveOrder, val)
	}

	//file.Print()
	for _, v := range moveOrder {
		//fmt.Println("mixing:", v)
		file.Mix(v, 0)
		//file.Print()
	}
	//file.Print()

	i, _ := file.ScanFor(0, 1)
	//fmt.Println(i, n)
	fmt.Println(file.Get(1000 + i).value)
	fmt.Println(file.Get(2000 + i).value)
	fmt.Println(file.Get(3000 + i).value)
}

func Solve2(path string) {
	data := util.GetFileLines(path)
	decryptionkey := 811589153
	file := NewLinkedList()
	moveOrder := []int{}
	for _, line := range data {
		val, _ := strconv.Atoi(line)
		dval := val * decryptionkey
		file.Add(dval)
		moveOrder = append(moveOrder, dval)
	}

	//file.Print()
	fmt.Println("starting", file.size)
	for i := 0; i < 10; i++ {
		for i := range moveOrder {
			//fmt.Println("mixing:", v, i)
			file.ModMix(i)
		}
		//file.Print()
		//fmt.Println("done iteration:", i)
	}

	i, _ := file.ScanFor(0, 10)
	a := file.Get(1000 + i).value
	b := file.Get(2000 + i).value
	c := file.Get(3000 + i).value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a + b + c)
}
