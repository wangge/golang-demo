package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main()  {
	var a,b []int
	for i := 1; i <= 33; i++{
		a = append(a,i)
		if i<17{
			b=append(b,i)
		}
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Printf("Please Enter run times:")
	var t int
	fmt.Scanln(&t)
	f, err := os.OpenFile("lottery.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil{
		return
	}
	defer f.Close()

	for t>0{
		var na = make([]int,33)
		copy(na, a)
		_, so := generateLottery(na, b, r)

		fmt.Println(so)
		t--

		//f.WriteString(strings.Join(so,"") + "\n")
	}
}

func generateLottery(a, b []int, r *rand.Rand)  ([]int, []string){
	var c []int
	var s []string
	n := 33
	for i:=0; i<6; i++{
		m := r.Intn(n)
		n = n-1
		c = append(c, a[m])
		a = resetSlice2(a, m)
	}
	sort.Ints(c)
	for _,i:= range c{
		//if i < 10 {
		//	s = append(s, "0")
		//}
		s = append(s, strconv.Itoa(i))
	}
	bl := b[r.Intn(16)]
	c = append(c,bl)
	s = append(s,strconv.Itoa(bl))
	return c,s
}
func resetSlice2(s []int, i int) []int {
	len := len(s)
	s[i] = s[len-1]
	s = s[:len-1]
	return s
}