package main

import "fmt"

func main()  {
	//fmt.Println(add("a","b"))
	arr := []int{1,2,3}
	fmt.Println(solve(arr,3,9))
}
func solve(arr []int,n,m int) int {
	res := 0
	hashMap := make(map[int]int)
	for _,v := range arr{
		hashMap[v] = 1
	}
	for _,v2 := range arr{
		if hashMap[m-v2] == 1{
			res += 1
			hashMap[m-v2] = 2
			hashMap[v2] = 2
		}
	}
	return  res
}




func add(a string, b string) string{
	List := "0123456789abcdefghijklmnopqrstuvwxyz"
	i,j := len(a)-1,len(b)-1
	tmp := 0
	sum := ""
	for i>=0 && j>=0{
		s := getInt(a[i])+getInt(b[j])+tmp
		if s >= 36 {
			tmp = 1
			sum = string(List[s%36])+sum
		} else {
			tmp = 0
			sum = string(List[s])+sum
		}
		i--
		j--
	}
	for i >= 0{
		s := getInt(a[i])+tmp
		if s >= 36{
			tmp = 1
			sum = string(List[s%36])+sum
		} else {
			tmp = 0
			sum = string(List[s])+sum
		}
		i--
	}
	for j >= 0{
		s := getInt(b[j])+tmp
		if s >= 36{
			tmp = 1
			sum = string(List[s%36])+sum
		} else {
			tmp = 0
			sum = string(List[s])+sum
		}
		j--
	}
	if tmp != 0{
		sum = "1"+sum
	}
	return sum
}

func getInt(u uint8) int {
	if u > '0'  && u <= '9' {
		return  int(u-'0')
	}
	return int(u-'a')+10
}
