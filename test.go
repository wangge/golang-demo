package main

import (
	"wangge/wg"
)

func main()  {


	controllerDir := "Operation"
	parent := "Order"
	name := "SettlementFlow"
	des := "订单流水"
	service := "Operation"
	tableName := "settlement"
	fields := wg.FieldInfos2("zk_local_toc_main", tableName)

	fucs := []wg.FunTp{
	{
		Name: "FillSettlement",
		Des: "完成划账",
		Param: []wg.PostParam{
			{"settlementNo","string","结算单号",true, "",[]wg.ArrayParams{}},
			{"paymentChannel","string","渠道",true, "",[]wg.ArrayParams{}},
			{"paymentNo","string","流水号",true,"",[]wg.ArrayParams{}},
		},
		PK: "",
		Void: false,
		Class: name,
		Parent: parent,
		Service: service,
	},
		{
			Name: "statisticList",
			Des: "统计",
			Param: []wg.PostParam{
				{"params","array","筛选条件",true, "",[]wg.ArrayParams{
					{"page_no","int","页码",false,"1"},
					{"page_size","int","每页显示条数",false,"10"},
					{"payment_time_start","date","支付开始时间",false,""},
				}},
			},
			PK: "",
			Void: false,
			Class: name,
			Parent: parent,
			Service: service,
		},
		{
			Name: "add",
			Des: "统计",
			Param: []wg.PostParam{
				{"params","array","筛选条件",true, "",fields},
			},
			PK: "",
			Void: false,
			Class: name,
			Parent: parent,
			Service: service,
		},
		{
			Name: "update",
			Des: "统计",
			Param: []wg.PostParam{
				{"params","array","筛选条件",true, "",fields},
			},
			PK: "id",
			Void: false,
			Class: name,
			Parent: parent,
			Service: service,
		},
	}
	for _, fuc := range fucs{
		// controller
		classController := wg.ClassTp{des,"Controller_Base_"+name,"Controller_Base_Base","",""}
		filename := controllerDir+"/Api/Controller/Base/"+name+".php"
		tplClass := "tmp/Partial/classTp.txt"
		// 创建类文件
		wg.CreateClass(filename, tplClass, classController)
		// 添加方法
		tpl := "tmp/Partial/controller_controller.txt"
		wg.AddMethod(filename, tpl, fuc)

		//controller module
		classController = wg.ClassTp{des,name,"ModuleBase","","Module"}
		filename = controllerDir+"/Api/Module/"+name+".php"
		wg.CreateClass(filename, tplClass, classController)
		tpl = "tmp/Partial/controller_module.txt"
		wg.AddMethod(filename, tpl, fuc)

		// controller service
		classController = wg.ClassTp{des,name,"Base",parent+"\\"+name,"Service\\Operation"}
		filename = controllerDir+"/Api/Service/Operation/"+name+".php"
		wg.CreateClass(filename, tplClass, classController)

		// service handel
		classController = wg.ClassTp{des,name,"\\Handle\\Base","","Handle\\"+parent}
		filename = controllerDir+"/Service/Handle/"+parent+"/"+name+".php"
		wg.CreateClass(filename, tplClass, classController)
		tpl = "tmp/Partial/service_handle.txt"
		wg.AddMethod(filename, tpl, fuc)

		// service module
		classController = wg.ClassTp{des,name,"\\Module\\Base","","Module\\"+parent}
		filename = controllerDir+"/Service/Module/"+parent+"/"+name+".php"
		wg.CreateClass(filename, tplClass, classController)
		tpl = "tmp/Partial/service_module.txt"
		wg.AddMethod(filename, tpl, fuc)
	}


}



