package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main()  {
	var a,b []int
	for i := 1; i <= 35; i++{
		a = append(a,i)
		if i<12{
			b=append(b,i)
		}
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Printf("Please Enter run times:")
	var t int
	fmt.Scanln(&t)

	for t>0{
		var na = make([]int,35)
		var nb = make([]int,12)
		copy(na, a)
		copy(nb, a)
		_, so := generateLotteryLT(na, nb, r)
		so2 := append([]string{"21200"}, so...)
		fmt.Println(strings.Join(so2, "\t"))
		t--
	}
}

func generateLotteryLT(a, b []int, r *rand.Rand)  ([]int, []string){
	var c []int
	var s []string
	n := 35
	for i:=0; i<5; i++{
		m := r.Intn(n)
		n = n-1
		c = append(c, a[m])
		a = resetSliceLT(a, m)
	}
	sort.Ints(c)
	for _,i:= range c{
		if i <10 {
			s = append(s, "0"+strconv.Itoa(i))
		}else{
			s = append(s, strconv.Itoa(i))
		}

	}
	nn :=12
	for i:=0;i<2;i++{
		rr := r.Intn(nn)
		bl := b[rr]
		b = resetSliceLT(b,rr)
		c = append(c,bl)
		if bl < 10 {
			s = append(s, "0"+strconv.Itoa(bl))
		}else{
			s = append(s, strconv.Itoa(bl))
		}
		nn--
	}


	return c,s
}
func resetSliceLT(s []int, i int) []int {
	len := len(s)
	s[i] = s[len-1]
	s = s[:len-1]
	return s
}
