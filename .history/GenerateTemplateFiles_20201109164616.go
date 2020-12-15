package main
/**
	根据模板文件生成，运营后台Operation的API和Service文件
    模板文件tmp
 */
import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"wangge/wg"
)
type Perm struct{
	Create bool
	Edit bool
	Delete bool
	List bool
	Export bool
	Online bool
	Offline bool
	Audit bool
	Detail bool
}
type Fuc struct {
	Fname string
	Param []string
	Des string
}
type TData struct {
	Table string
	TableName string
	Parent string
	Name string
	NameLow string
	Fd []wg.Field
	Per Perm
	Fcs []Fuc
}

func main()  {
	workspace := "C:\\workspace\\Operation/"
	workspace = "Operation/"
	tmpPre := "tmp"
	name := "order"
	parent := "Order"
	tableInfo := "商城订单"
	tableName := "payment_order"
	fields := wg.MulTableInfos("zk_local_toc_main", tableName, "order_item")

	pe := Perm{
		Create: false,
		Edit:false,
		Delete:false,
		List:true,
		Export:true,
		Online:false,
		Offline:false,
		Audit: false,
		Detail: true,
	}
	r := TData{
		Table: tableInfo,
		TableName: tableName,
		Parent: parent,
		Name: Capitalize(name),
		NameLow: name,
		Fd: fields,
		Per: pe,
		Fcs: []Fuc{
			// {"statisticList",[]string{},"燃气充值统计"},
			// {"syncToEntService",[]string{"uid","order_id"},"重新推送燃气充值记录"},
		},
	}

	//if r.Name != r.Parent{
	//	r.Name = r.Parent + r.Name
	//}

	// 创建Controller
	// 1创建ApiController目录
	dirName := workspace+"Api/Controller/Base"
	fileName := dirName+"/"+r.Name+".php"
	tp := tmpPre+ "/controller/DemoController.php"
	FileOut(tp, fileName, r)
	fmt.Println("Controller build Success!")

	// 2创建module
	dirName = workspace+"Api/Module/"
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/controller/DemoModule.php"
	FileOut(tp, fileName, r)
	fmt.Println("Module build Success!")


	// 3创建api service
	dirName = workspace+"Api/Service/Operation"
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/controller/DemoService.php"
	FileOut(tp, fileName, r)
	fmt.Println("Service build Success!")


	// 4创建Service
	dirName = workspace+"Service/Handler/"+parent
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/Service/DemoHandle.php"
	FileOut(tp, fileName, r)
	fmt.Println("Handler build Success!")


	// 创建Module
	dirName = workspace+"Service/Module/"+parent
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/Service/DemoModule.php"
	FileOut(tp, fileName, r)
	fmt.Println("Module build Success!")


	// 创建Model
	//dirName = workspace+"Service/Model/Main"
	//fileName = dirName+"/"+r.Name+".php"
	//tp = tmpPre+ "/Service/DemoModel.php"
	//FileOut(tp, fileName, r)
	//fmt.Println("Model build Success!")

	// 生成API DOC
	if pe.Create{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/create.txt"
		tp = tmpPre+ "/Doc/create.txt"
		FileOut(tp, fileName, r)
	}
	if pe.Edit{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/edit.txt"
		tp = tmpPre+ "/Doc/edit.txt"
		FileOut(tp, fileName, r)

	}
	if pe.Detail{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/detail.txt"
		tp = tmpPre+ "/Doc/detail.txt"
		FileOut(tp, fileName, r)
	}

	if pe.List{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/list.txt"
		tp = tmpPre+ "/Doc/list.txt"
		FileOut(tp, fileName, r)
	}
	if pe.Delete{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/delete.txt"
		tp = tmpPre+ "/Doc/delete.txt"
		FileOut(tp, fileName, r)
	}
	if pe.Online{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/online.txt"
		tp = tmpPre+ "/Doc/online.txt"
		FileOut(tp, fileName, r)
	}
	if pe.Offline{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/offline.txt"
		tp = tmpPre+ "/Doc/offline.txt"
		FileOut(tp, fileName, r)
	}
	if pe.Audit{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/audit.txt"
		tp = tmpPre+ "/Doc/audit.txt"
		FileOut(tp, fileName, r)
	}
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func FileOut(tpl,outFile string , r TData)  {
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
}
