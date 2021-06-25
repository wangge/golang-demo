package main

import (
	"encoding/json"
	"fmt"
)

type Test3 struct {
	ID int64
	TestID testID
}

type Test4 struct {
	ID int64 `json:"id"`
}

type testID int64
type Test5 struct {
	*Test4
	ID int64 `json:"id"`
	Name string `json:"name"`
}
func main()  {
	test3List := make([]*Test3, 0)
	test4 := &Test4{
		ID:  int64(4),
	}
	id := testID(int64(12))
	test3 := &Test3{
		ID: test4.ID,
		TestID: id,
	}

	test5 := &Test5{
		Test4:test4,
		Name: "hahahah",
		ID:test4.ID,
	}
	p5,_ := json.Marshal(test5)
	fmt.Println(string(p5))
	test4.ID = 5
	p52,_ := json.Marshal(test5)
	fmt.Println(string(p52))
	for i:=0; i<3;i++{
		test3List = append(test3List, &Test3{
			ID: test4.ID,
		})
	}

	p, _:= json.Marshal(test3List)
	print(string(p))
	fmt.Printf("%v", test3)
	fmt.Printf("%v", test5.ID)
	test4.ID = 5
	id = testID(int64(13))
	p2, _:= json.Marshal(test3List)
	print(string(p2))
	fmt.Printf("%v", test3)
	fmt.Printf("%v", test5.ID)
}
