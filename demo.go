package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//defer func() {
	//	if r:= recover(); r!=nil{
	//		fmt.Println("异常消息", r)
	//	}
	//	fmt.Println("继续执行")
	//}()
	//test()
	//fmt.Println("无异常")
	//ab := &C{}
	//i := aa(ab)
	//ab.run(i)
	//re,_:=regexp.Compile(`^[1-9]\d{4,9}$`)
	//sellerIdStr:= "10001"
	//fmt.Sprintf(sellerIdStr)
	//b:=re.MatchString(sellerIdStr)
	//fmt.Println(b)

	//data := make(map[string]interface{})
	//if reflect.DeepEqual(data, map[string]interface{}{}){
	//	fmt.Println(1)
	//}else {
	//	fmt.Println(2)
	//}
	//data2 := map[string]interface{}{
	//	"ss":1,
	//}
	//if reflect.DeepEqual(data2, map[string]interface{}{}){
	//	fmt.Println(3)
	//}else {
	//	fmt.Println(4)
	//}
	//fmt.Println(len(data))
	//fmt.Println(len(data2))
	//s := []int64{3,4,5,6}
	//fmt.Println( arrJoin(s,","))

	//s := "{\"a\":\"b\",\"c\":2}"
	//s = "11"
	//fmt.Println(json.Valid([]byte(s)))
	//var j []int64
	//fmt.Println(json.Unmarshal([]byte(s),&j))
	//fmt.Println(j)
	//fmt.Println(reflect.TypeOf(j))
	////fmt.Println(len(j))


	s:= "104.042561 30.628152,104.001019 30.598752,104.075348 30.575034,104.098951 30.60747, 104.089853 30.639304 ,104.042561 30.628152"
	//104.089853,30.639304
	coors := getCoordinates(s)
	fmt.Println(coors.Mid.ToString())
}
func test()  {
	fmt.Println("无异常2")
	panic("捕获异常")
}

type aa interface {
	a1()
}
type  A struct {

}

func (a *A) a1()  {
	fmt.Println("a1")
}

func (a *A) run(t aa)  {
	t.a1()
}
type B struct {
	A
}
func (b *B) a1()  {
	fmt.Println("a2")
}

type C struct {
	A
}

func arrJoin(arr []int64, sep string) string{
	strBuf := bytes.Buffer{}
	for i, num := range arr {
		if i > 0 {
			strBuf.WriteString(sep)
		}
		strBuf.WriteString(fmt.Sprintf("%d", num))
	}
	return strBuf.String()
}

// 获取上下左右中经纬度点
func getLocation(str string) (string, string, string, string,string){
	coorArr := strings.Split(str,",")
	var latMax, latMin, lngMax, lngMin float64 = 0,0,0,0
	var top, down,left,right,mid string = "", "", "", "",""
	for _, coor := range coorArr{
		coorS := strings.Split(coor, " ")
		lngstr, latstr := coorS[0], coorS[1]
		lngF,_ := strconv.ParseFloat(lngstr, 64)
		latF,_ := strconv.ParseFloat(latstr, 64)
		if lngF >lngMax{
			lngMax = lngF
			// 右
			right = coor
		}
		if lngMin <=0 || lngF < lngMin{
			lngMin = lngF
			left = coor
		}

		if latF > latMax{
			latMax = latF
			top = coor
		}
		if latMin <= 0 || latF < latMin{
			latMin = latF
			down = coor
		}
	}
	mid = fmt.Sprintf("%f %f", (lngMin+lngMax)/2, (latMin+latMax)/2)
	return top, down, left, right, mid
}

//经纬度
type Coordinate struct {
	Lng float64 // 经度
	Lat float64 // 纬度
}

// 经纬度string
func (c *Coordinate)ToString() string {
	return fmt.Sprintf("%f %f", c.Lng, c.Lat)
}

type Coordinates struct {
	Top *Coordinate
	Down *Coordinate
	Left *Coordinate
	Right *Coordinate
	Mid *Coordinate
}
// 根据多边形四个方位及中间经纬度
func getCoordinates(polygon string)  *Coordinates{
	coorArr := strings.Split(polygon,",")
	var latMax, latMin, lngMax, lngMin float64 = 0,0,0,0
	coordinates := &Coordinates{}
	for _, coor := range coorArr{
		coorS := strings.Split(strings.Trim(coor, " "), " ")
		lngstr, latstr := coorS[0], coorS[1]
		lngF,_ := strconv.ParseFloat(lngstr, 64)
		latF,_ := strconv.ParseFloat(latstr, 64)
		if lngF >lngMax{
			lngMax = lngF
			// 右
			coordinates.Right = &Coordinate{Lat: latF, Lng: lngF}
		}
		if lngMin <=0 || lngF < lngMin{
			// 左
			lngMin = lngF
			coordinates.Left = &Coordinate{Lat: latF, Lng: lngF}
		}

		if latF > latMax{
			//上
			latMax = latF
			fmt.Println(latMax, lngF)
			coordinates.Top = &Coordinate{Lat: latF, Lng: lngF}
		}
		if latMin <= 0 || latF < latMin{
			//下
			latMin = latF
			coordinates.Down = &Coordinate{Lat: latF, Lng: lngF}
		}
	}
	coordinates.Mid = &Coordinate{Lng: (lngMin+lngMax)/2, Lat: (latMin+latMax)/2}
	return coordinates
}