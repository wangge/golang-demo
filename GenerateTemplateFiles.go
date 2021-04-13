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
// 可生成方法
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
// 其他方法
type Fuc struct {
	Fname string
	Param []string
	Des string
	ParamIsArr bool
	NeedPk bool
	PK string
}
// 列表搜索参数
type SearchParmers struct {
	SName string
	SType string
	SMust bool
	SDes string
}
// 模板数据
type TData struct {
	Table string
	TableName string
	Parent string
	Name string
	NameLow string
	Fd []wg.Field
	Per Perm
	Fcs []Fuc
	Search []SearchParmers
}

func main()  {
	workspace := "C:\\workspace\\Operation/"
	workspace = "Operation/"
	tmpPre := "tmp3"
	name := "user"
	parent := "User"
	tableInfo := "用户列表"
	tableName := "user"
	fields := wg.FieldInfos("zk_local_toc_main", tableName)
	
	pe := Perm{
		Create: false,
		Edit:false,
		Delete:false,
		List:true,
		Export:false,
		Online:false,
		Offline:false,
		Audit: false,
		Detail: false,
	}
	searchp := []SearchParmers{
		{"mobile","string",false,"手机号"},
		{"create_time_start","string",false,"注册时间-开始时间"},
		{"create_time_end","string",false,"注册时间-结束时间"},
	}
	var searchkey []string
	for _,s := range searchp{
		searchkey = append(searchkey, s.SName)
	}
	functions := []Fuc{
		//{"FillSettlement",[]string{"settlementNo","paymentChannel","paymentNo"},"完成划账", true, true, "storeId"},
		//{"SettlementChannel",[]string{},"划账渠道", false, false,""},
		//{"getSettlementDetail",[]string{"settlementNo"},"获取结算流水单详情", true, true,"storeId"},
	}
	r := TData{
		Table: tableInfo,
		TableName: tableName,
		Parent: parent,
		Name: Capitalize(name),
		NameLow: name,
		Fd: fields,
		Per: pe,
		Fcs: functions,
		Search:searchp,
	}
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


	r.Parent = "Zeus\\"+r.Parent
	// 3创建api service
	dirName = workspace+"Api/Service/Operation"
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/controller/DemoService.php"
	FileOut(tp, fileName, r)
	fmt.Println("Service build Success!")


	// 4创建Service
	dirName = workspace+"Service/Handler/"+r.Parent
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/Service/DemoHandle.php"
	FileOut(tp, fileName, r)
	fmt.Println("Handler build Success!")


	// 创建Module
	dirName = workspace+"Service/Module/"+r.Parent
	fileName = dirName+"/"+r.Name+".php"
	tp = tmpPre+ "/Service/DemoModule.php"
	FileOut(tp, fileName, r)
	fmt.Println("Module build Success!")




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

	type Data2 struct {
		Name string
		Fuc
	}
	for  _, otp := range functions{
		dirName = "Doc/"
		fileName = dirName+"/"+r.Name+"/"+otp.Fname+".txt"
		tp = tmpPre+ "/Doc/other.txt"
		d2 := Data2{r.Name, otp}
		for _,x:= range d2.Param{
			fmt.Println(x)
		}
		FileOut(tp, fileName, d2)
	}
	// 创建Model
	//for k, model := range map[string] string{"BusinessPayment":"business_payment","Business":"business"}{
	//	dirName = workspace+"Service/Model/Main"
	//	fileName = dirName+"/" + k +".php"
	//	tp = tmpPre+ "/Service/DemoModel.php"
	//	r.Name = k
	//	r.TableName = model
	//	FileOut(tp, fileName, r)
	//}
	//dirName = workspace+"Service/Model/Main"
	//fileName = dirName+"/"+r.Name+".php"
	//tp = tmpPre+ "/Service/DemoModel.php"
	//FileOut(tp, fileName, r)
	//fmt.Println("Model build Success!")
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

/**
生成文件
 */
func FileOut(tpl,outFile string , r interface{})  {
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
