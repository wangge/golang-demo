package main

import (
	"errors"
	"fmt"
)

var errs = make(chan error)

func main()  {


	go func() {
		errs <- errors.New("try..........")
	}()
	go func() {
		e:= <-errs
		fmt.Println("catch", e.Error())
	}()

	for{

	}
}
