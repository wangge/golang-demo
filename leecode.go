package main

import "fmt"

func main()  {
	fmt.Println(reverse(123))
	fmt.Println(reverse(4541))
	fmt.Println(reverse(250))

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
