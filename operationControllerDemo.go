package main

import (
	"wangge/wg"
)

func main()  {
	controllerDir := "Operation"
	parent := "Refund"
	name := "PaymentOrderRefund"
	des := "退款记录"
	service := "Operation"
	//tableName := "user_gas_card"
	//fields := wg.FieldInfos2("zk_local_toc_main", tableName)
	//fmt.Println(fields)
	fucs := []wg.FunTp{
	//{
	//	Name: "syncBusiness",
	//	Des: "重新推送业务",
	//	Param: []wg.PostParam{
	//		{"id","int","业务id",true,"0",[]wg.ArrayParams{}},
	//	},
	//	PK: "",
	//	Void: false,
	//	Class: name,
	//	Parent: parent,
	//	Service: service,
	//},
	{
		Name: "List",
		Des: "退款记录",
		Param: []wg.PostParam{
			{"searchParam","array","过滤条件",true, "",[]wg.ArrayParams{

				{"partner_id","int","公司id",false,""},
				{"order_sn","string","订单号",false,""},
				{"trade_no","string","支付单号",false,""},
				{"card_no","string","缴费号",false,""},
				{"meter_no","string","表具号",false,""},
				{"source","string","表具号",false,"充值渠道，小程序(huixiang_plus_applet)、终端机(arm_terminal)、APP(huixiang_plus_app)"},
				{"payment_method","string","表具号",false,"支付方式，用于区分不同渠道不同类型的支付,支付宝(Alipay)、微信(Weixin),四川农商(SCRural)"},
				{"payment_time_start","date","支付时间开始",false,""},
				{"payment_time_end","date","支付时间结束",false,""},
				{"refund_time_start","date","退款发起时间开始",false,""},
				{"refund_time_end","date","退款发起时间结束",false,""},
				{"complete_time_start","date","退款完成时间开始",false,""},
				{"complete_time_end","date","退款完成时间结束",false,""},
				{"outer_card_no","string","第三方缴费号",false,""},
				{"nickname","string","付款人",false,""},
				{"mobile","string","付款人手机号",false,""},
				{"category","string","对接系统1-营收,3-金卡",false,""},
			}},
			{"page_no","int","页码",false,"1",[]wg.ArrayParams{}},
			{"page_size","int","每页显示条数",false,"10",[]wg.ArrayParams{}},
		},
		PK: "",
		Void: false,
		Class: name,
		Parent: parent,
		Service: service,
	},
	//{
	//	Name: "businessNoCardList",
	//	Des: "无卡预约",
	//	Param: []wg.PostParam{
	//		{"searchParam","array","过滤条件",true, "",[]wg.ArrayParams{
	//			{"partner_id","int","燃气公司id",false,"0"},
	//			{"business_type_id","int","业务类型id",false,"0"},
	//			{"id","int","申请单号",false,"0"},
	//			{"charge_no","string","缴费号",false,""},
	//			{"mobile","string","申请人手机",false,""},
	//			{"user_mobile","string","申请人账号手机",false,""},
	//			{"create_time_start","date","申请时间开始",false,""},
	//			{"create_time_end","date","申请时间结算",false,""},
	//			{"status","int","业务状态0-待审核，1-审核不通过，2-审核通过，5-业务终止，9-已完成",false,""},
	//			{"sync_status","int","推送状态4001-推送成功，4003-推送失败",false,""},
	//		}},
	//		{"page_no","int","页码",false,"1",[]wg.ArrayParams{}},
	//		{"page_size","int","每页显示条数",false,"10",[]wg.ArrayParams{}},
	//	},
	//	PK: "",
	//	Void: false,
	//	Class: name,
	//	Parent: parent,
	//	Service: service,
	//},
	//{
	//	Name: "update",
	//	Des: "统计",
	//	Param: []wg.PostParam{
	//		{"params","array","筛选条件",true, "",fields},
	//	},
	//	PK: "id",
	//	Void: false,
	//	Class: name,
	//	Parent: parent,
	//	Service: service,
	//},
	}

	for _, fuc := range fucs{
		makeController(controllerDir,des,name,parent,fuc)
	}

}

func makeController(controllerDir,des,name ,parent string, fuc wg.FunTp)  {
	classController := wg.ClassTp{des,"Controller_Base_"+name,"Controller_Base_Base","",""}
	filename := controllerDir+"/Api/Controller/Base/"+name+".php"
	tplClass := "tmp/Partial/classTp.txt"
	// 创建类文件
	wg.CreateClass(filename, tplClass, classController)
	// 添加方法
	tpl := "tmp/Partial/controller_controller.txt"
	wg.AddMethod(filename, tpl, fuc)

	//controller module
	classController = wg.ClassTp{des,name,"\\Module\\ModuleBase","","Module\\"+parent}
	filename = controllerDir+"/Api/Module/"+parent+"/"+name+".php"
	wg.CreateClass(filename, tplClass, classController)
	tpl = "tmp/Partial/controller_module.txt"
	wg.AddMethod(filename, tpl, fuc)

	// controller service
	classController = wg.ClassTp{des,name,"Base",parent+"\\"+name,"Service\\Operation"}
	filename = controllerDir+"/Api/Service/Operation/"+name+".php"
	wg.CreateClass(filename, tplClass, classController)

	// service handel
	classController = wg.ClassTp{des,name,"\\Handler\\Base","","Handler\\"+parent}
	filename = controllerDir+"/Service/Handler/"+parent+"/"+name+".php"
	wg.CreateClass(filename, tplClass, classController)
	tpl = "tmp/Partial/service_handle.txt"
	wg.AddMethod(filename, tpl, fuc)

	// service module
	classController = wg.ClassTp{des,name,"\\Module\\Base","","Module\\"+parent}
	filename = controllerDir+"/Service/Module/"+parent+"/"+name+".php"
	wg.CreateClass(filename, tplClass, classController)
	tpl = "tmp/Partial/service_module.txt"
	wg.AddMethod(filename, tpl, fuc)

	// Doc
	filename = controllerDir+"/Doc/"+parent+"/"+fuc.Name+".txt"
	tpl = "tmp/Partial/doc.txt"
	wg.FileOut(filename,tpl ,fuc)
}



