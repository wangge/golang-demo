package wg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func getStringLen(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	fmt.Print(s[0])
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(a ,b int) int {
	if a>b{
		return a
	}
	return b
}
func CreateClass(fileName string, tp string, cp ClassTp){
	if !isExist(fileName){
		FileOut(fileName, tp, cp)
	}
}
/**
将方法添加到文件中去
 */
func AddMethod (fileName string, tp string, fc FunTp){
	t, err  := template.ParseFiles(tp)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	buff := new(bytes.Buffer)
	t.Execute(buff, fc)
	//fmt.Println(buff.String())

	// 读取文件
	// 读取文件内容
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println(err)
	}
	if strings.Contains(string(body), fc.Name+"("){
		fmt.Println("方法已经添加到文件中")
		return
	}

	//查找}最后出现的位置
	//lastIndex := strings.LastIndex(string(body),"}")
	// 插入内容
	//insertContent := "\n"+buff.String()+"\n"
	// 将模板文件添加到文件
	sSplice := strings.Split(string(body), "}")
	fmt.Println(sSplice)

	sSplice = sSplice[:len(sSplice)-1]
	fmt.Println(sSplice)
	content := strings.Join(sSplice,"}") + "\n"+buff.String()+"\n}"
	fmt.Println(content)
	err2 :=ioutil.WriteFile(fileName, []byte(content), 0777)
	if err != nil{
		fmt.Println(err2)
	}
}
/**
根据模板生成文件
 */
func FileOut(outFile,tpl string , r interface{})  {
	t, err  := template.ParseFiles(tpl)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	dirNames := strings.Split(outFile, "/")
	os.MkdirAll(strings.Join(dirNames[:len(dirNames)-1],"/"), 0777)
	file, err := os.OpenFile(outFile,os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}
	t.Execute(file,r)
	fmt.Println("创建"+outFile+"成功")
}
// 判断资源包是否存在
func isExist(path string)(bool){
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		if os.IsNotExist(err){
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}