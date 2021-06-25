package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
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

	//s:= "104.042561 30.628152,104.001019 30.598752,104.075348 30.575034,104.098951 30.60747, 104.089853 30.639304 ,104.042561 30.628152"
	////104.089853,30.639304
	//coors := getCoordinates(s)
	//fmt.Println(coors.Mid.ToString())

	//s := []int64{1,2,3,4,5,6,7,8}
	//s1 :=SlinceDeleteInt64(3, s)
	//fmt.Println(s)
	//fmt.Println(s1)
	//fmt.Println(s[:7])

	//fmt.Println(fmt.Sprintf("%%%s%%", "dsdss"))
	//fmt.Println(convertToBin(24))
	//fmt.Println(calculateBinPosition(24))

	//a1:= &A{}
	//a2:= &A{}
	//a3:= A{}
	//a4:= A{}
	//	fmt.Println(&a1)
	//	fmt.Println(&a2)
	//	fmt.Println(a3)
	//	fmt.Println(a4)

	//ids := make([]int64, 0)
	////ids = append(ids,int64(12))
	//for _, id:= range ids{
	//	fmt.Println("循环")
	//	fmt.Println(id)
	//}
	//fmt.Println("11")

	//s := "[ [ 104.05324406212371, 30.538756829196114 ], [ 104.053158078680312, 30.540231211726617 ], [ 104.053603629250674, 30.541234317639269 ], [ 104.054158613294476, 30.541880608577586 ], [ 104.055252948028723, 30.541988323315795 ], [ 104.056409816176384, 30.541799822445515 ], [ 104.056495799619768, 30.540877509340898 ], [ 104.056527066326481, 30.539786879621779 ], [ 104.056370732793027, 30.538615448873276 ], [ 104.054846480841732, 30.538332687610112 ], [ 104.053587995897345, 30.538238433672809 ], [ 104.05324406212371, 30.538756829196114 ] ]"
	//data := make([][]float64, 0)
	//err := json.Unmarshal([]byte(s), &data)
	//fmt.Println(err)
	//fmt.Println(data)

	//s := `{"rows":[{"store_id":"1","sku_no":"10048","deal_hash_id":"76t1h20004","product_id":"76","sku_name":"","sale_price":"127.00","inventory_num":"19300","product_name":"","is_onsale":"1","store_name":"","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/23\/1\/main_image\/20200312\/1_00035048720345034.png","specs":[]},{"store_id":"1","sku_no":"10195","deal_hash_id":"106t1h20021","product_id":"106","sku_name":"","sale_price":"130.00","inventory_num":"43","product_name":"","is_onsale":"1","store_name":"","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/23\/1\/main_image\/20200315\/1_00037563403077284.jpg","specs":[]},{"store_id":"1","sku_no":"10233","deal_hash_id":"108t1h20022","product_id":"108","sku_name":"\u9ed8\u8ba4\uff0c\u4e0d\u542f\u7528\u591a\u89c4\u683c","sale_price":"20.00","inventory_num":"85","product_name":"\u4f18\u9009\u70ed\u6c34\u5668","is_onsale":"1","store_name":"\u4e2d\u79d1\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/23\/1\/main_image\/20200315\/1_00037565763159303.jpg","specs":[]},{"store_id":"2","sku_no":"10262","deal_hash_id":"118t1h20028","product_id":"118","sku_name":"2","sale_price":"10.00","inventory_num":"500","product_name":"\u5546\u62372 \u5546\u54c11","is_onsale":"1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200327\/2_00048131705749955.jpg","specs":[{"sku_no":"10262","spec_name":"\u89c4\u683c","spec_value":"2"}]},{"store_id":"2","sku_no":"10263","deal_hash_id":"118t1h20028","product_id":"118","sku_name":"1","sale_price":"15.00","inventory_num":"95","product_name":"\u5546\u62372 \u5546\u54c11","is_onsale":"1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200327\/2_00048131675170965.jpg","specs":[{"sku_no":"10263","spec_name":"\u89c4\u683c","spec_value":"1"}]},{"store_id":"2","sku_no":"10207","deal_hash_id":"127t1h20033","product_id":"127","sku_name":"1\u9ed1\u8272","sale_price":"5.00","inventory_num":"74","product_name":"\u5546\u62372 \u5546\u54c121","is_onsale":"1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200327\/2_00048132658390538.jpg","specs":[{"sku_no":"10207","spec_name":"\u89c4\u683c","spec_value":"1"},{"sku_no":"10207","spec_name":"\u989c\u8272","spec_value":"\u9ed1\u8272"}]},{"store_id":"2","sku_no":"10270","deal_hash_id":"127t1h20033","product_id":"127","sku_name":"1\u7ea2\u8272","sale_price":"5.00","inventory_num":"187","product_name":"\u5546\u62372 \u5546\u54c121","is_onsale":"1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200327\/2_00048132465898653.png","specs":[{"sku_no":"10270","spec_name":"\u89c4\u683c","spec_value":"1"},{"sku_no":"10270","spec_name":"\u989c\u8272","spec_value":"\u7ea2\u8272"}]},{"store_id":"3","sku_no":"10269","deal_hash_id":"147t1h20047","product_id":"147","sku_name":"\u9ed8\u8ba4\uff0c\u4e0d\u542f\u7528\u591a\u89c4\u683c","sale_price":"0.01","inventory_num":"0","product_name":"\u8bd5\u4e0b\u5355\u6d41\u7a0b","is_onsale":"1","store_name":"","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/11\/3\/main_image\/20200324\/3_00045421377715215.png","specs":[]},{"store_id":"2","sku_no":"10271","deal_hash_id":"148t1h20048","product_id":"148","sku_name":"3\u53f7","sale_price":"4.00","inventory_num":"198","product_name":"\u5546\u62372 \u5546\u54c13 \u90ae\u8d395","is_onsale""1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200329\/2_00049790405957365.jpg","specs":[{"sku_no":"10271","spec_name":"\u89c4\u683c\u578b\u53f7","spec_value":"3\u53f7"}]},{"store_id":"2","sku_no":"10272","deal_hash_id":"148t1h20048","product_id":"148","sku_name":"2\u53f7","sale_price":"3.00","inventory_num":"181","product_name":"\u5546\u62372 \u5546\u54c13 \u90ae\u8d395","is_onsale":"1","store_name":"\u6c47\u4eab\u4e50\u65d7\u8230\u5e97","image":"https:\/\/appletshop.oss-cn-beijing.aliyuncs.com\/ftp\/product\/6\/2\/sku_image\/20200329\/2_00049790381587357.jpg","specs":[{"sku_no":"10272","spec_name":"\u89c4\u683c\u578b\u53f7","spec_value":"2\u53f7"}]}],"rowCount":60,"rowsPerPage":10,"pageIndex":0,"pageNumber":1,"pageCount":6}`
	//fmt.Println(s)
	//UniCode2Ch(s)
	//rs := make(map[string]interface{}, 0)
	//err := json.Unmarshal([]byte(s), &rs)
	//fmt.Println(err.Error())
	//fmt.Println(rs)

	//s := `{"appCode":"CMW","resStatus":504,"resMsg":"\u9ed8\u8ba4\uff0c\u4e0d\u542f\u7528\u591a\u89c4\u683c","data":null}`
	//b := []byte(s)
	//fmt.Println(string(b))
	//// {"appCode":"CMW","resStatus":504,"resMsg":"\u57df\u540d\u670d\u540d\u670d\u52a1\u7c7b\u578b\u6709\u8bef\uff01","data":null}
	//m := make(map[string]interface{})
	//err :=json.Unmarshal(b,&m)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(m)

	//s:="2021-05-27 00:00:01"
	//d:="2021-05-28 00:00:01"
	////d:="2021-05-28 23:59:01"
	//print(SubDays(s,d))

	//s := "510203"
	//fmt.Println(s[:2])
	//fmt.Println(s[2:4])
	//fmt.Println(s[4:])
	t := "2021-06-10"
	fmt.Println(Unix(t))
	fmt.Println(UnixTime(t, "2006-01-02"))
}

// Unix 将日期时间格式转换为unix时间戳
func Unix(t string) int64 {
	format := GetFormat(t)
	if t == "" || format == "" {
		return 0
	}
	return UnixTime(t, format)
}

// GetFormat 获取日期时间戳格式
func GetFormat(t string) string {
	// 默认24小时制
	if m, e := regexp.MatchString(`\d{4}\-\d{1,2}\-\d{1,2}\s\d{1,2}:\d{1,2}:\d{1,2}`, t); e == nil && m {
		return "2006-01-02 03:04:05"
	}
	if m, e := regexp.MatchString(`\d{4}\-\d{1,2}\-\d{1,2}`, t); e == nil && m {
		return "2006-01-02"
	}
	if m, e := regexp.MatchString(`\d{1,2}\/\d{1,2}\/\d{4}`, t); e == nil && m {
		return "01/02/2006"
	}
	return ""
}
// UnixTime 将日期时间格式转换为unix时间戳
func UnixTime(dt, format string) int64 {
	if dt == "" {
		return 0
	}
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	// t, err := time.Parse(format, dt)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dt, GetLocation())
	if err != nil {
		return 0
	}
	return t.Unix()
}
func GetLocation() *time.Location {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return location
}
func SubDays(beginTime ,endTime string) int {
	timeFormat := "2006-01-02 15:04:05"        // 默认转换日期格式
	end, err := time.Parse(timeFormat, endTime) // string -> time
	if err != nil {
		fmt.Println("TIME:",beginTime, endTime)
		fmt.Println("time.Parse error", err.Error())
		return 0
	}
	start, err := time.Parse(timeFormat, beginTime) // string -> time
	if err != nil {
		return 0
	}
	return int(math.Ceil(end.Sub(start).Hours() / float64(24)))
}

func UniCode2Ch(textUnquoted string) {
	sUnicodev := strings.Split(textUnquoted, "\\u")

	var context string
	for _, v := range sUnicodev {

		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			//fmt.Println(err.Error())
		}
		context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)
}
func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func test() {
	fmt.Println("无异常2")
	panic("捕获异常")
}

type aa interface {
	a1()
}
type A struct {
}

func (a *A) a1() {
	fmt.Println("a1")
}

func (a *A) run(t aa) {
	t.a1()
}

type B struct {
	A
}

func (b *B) a1() {
	fmt.Println("a2")
}

type C struct {
	A
}

func arrJoin(arr []int64, sep string) string {
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
func getLocation(str string) (string, string, string, string, string) {
	coorArr := strings.Split(str, ",")
	var latMax, latMin, lngMax, lngMin float64 = 0, 0, 0, 0
	var top, down, left, right, mid string = "", "", "", "", ""
	for _, coor := range coorArr {
		coorS := strings.Split(coor, " ")
		lngstr, latstr := coorS[0], coorS[1]
		lngF, _ := strconv.ParseFloat(lngstr, 64)
		latF, _ := strconv.ParseFloat(latstr, 64)
		if lngF > lngMax {
			lngMax = lngF
			// 右
			right = coor
		}
		if lngMin <= 0 || lngF < lngMin {
			lngMin = lngF
			left = coor
		}

		if latF > latMax {
			latMax = latF
			top = coor
		}
		if latMin <= 0 || latF < latMin {
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
func (c *Coordinate) ToString() string {
	return fmt.Sprintf("%f %f", c.Lng, c.Lat)
}

type Coordinates struct {
	Top   *Coordinate
	Down  *Coordinate
	Left  *Coordinate
	Right *Coordinate
	Mid   *Coordinate
}

// 根据多边形四个方位及中间经纬度
func getCoordinates(polygon string) *Coordinates {
	coorArr := strings.Split(polygon, ",")
	var latMax, latMin, lngMax, lngMin float64 = 0, 0, 0, 0
	coordinates := &Coordinates{}
	for _, coor := range coorArr {
		coorS := strings.Split(strings.Trim(coor, " "), " ")
		lngstr, latstr := coorS[0], coorS[1]
		lngF, _ := strconv.ParseFloat(lngstr, 64)
		latF, _ := strconv.ParseFloat(latstr, 64)
		if lngF > lngMax {
			lngMax = lngF
			// 右
			coordinates.Right = &Coordinate{Lat: latF, Lng: lngF}
		}
		if lngMin <= 0 || lngF < lngMin {
			// 左
			lngMin = lngF
			coordinates.Left = &Coordinate{Lat: latF, Lng: lngF}
		}

		if latF > latMax {
			//上
			latMax = latF
			fmt.Println(latMax, lngF)
			coordinates.Top = &Coordinate{Lat: latF, Lng: lngF}
		}
		if latMin <= 0 || latF < latMin {
			//下
			latMin = latF
			coordinates.Down = &Coordinate{Lat: latF, Lng: lngF}
		}
	}
	coordinates.Mid = &Coordinate{Lng: (lngMin + lngMax) / 2, Lat: (latMin + latMax) / 2}
	return coordinates
}

func SlinceDeleteInt64(needle int64, haystack []int64) []int64 {
	length := len(haystack)
	for i := 0; i < length; i++ {
		if haystack[i] == needle {
			// 删除第i的元素
			// 将最后一个元素，赋值给第i的位置
			//haystack[i] = haystack[length-1]
			//haystack = haystack[:length-1]
			haystack = append(haystack[:i], haystack[i+1:]...)
			break
		}
	}
	// 然后删掉最后一个元素
	return haystack
}

func intToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

// 根据整数，计算二进制每一位是1，还是0
func calculateBinPosition(i int) []byte {
	var t []byte
	for ; i > 0; i /= 2 {
		lsb := i % 2
		t = append(t, byte(lsb))
	}
	return t
}

func disableEscapeHtml(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return "", err
	}
	return bf.String(), nil
}
