package main

import (
	"fmt"
	"sort"
)

func main()  {
	//fmt.Println(reverse(123))
	//fmt.Println(reverse(4541))
	//fmt.Println(reverse(250))
	s:=[]int{10,20,5}
	d:=[]int{5,5,2}
	fmt.Println(breakfastNumber(s,d,15))

}
func reverse(x int) int {
	var rev int
	rev = 0
	for x != 0{
		if temp := int32(rev); (temp*10)/10 != temp{
			return 0
		}
		rev = rev *10 + x%10
		x = x/10
	}
	return rev
}
func listSum(a []int) int {
	t := 0
	for _,i := range a{
		t +=i
	}
	return t
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)

	l1 := len(staple)
	l2 := len(drinks)
	i:=0
	j:= l2-1
	r:=0
	for i<l1 && j>=0{
		if staple[i]+drinks[j] <=x{
			r+=j+1
			i=i+1
		}else{
			j=j-1
		}
	}
	return r%1000000007
}