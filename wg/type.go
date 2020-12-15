package wg
type ArrayParams struct {
	Name, // 名称
	Type, // 类型
	Desc string // 描述
	Required bool // 是否必传
	Default string // 默认值
}
// 列表搜索参数
type PostParam struct {
	Name, // 名称
	Type, // 类型
	Desc string // 描述
	Required bool // 是否必传
	Default string // 默认值
	Array []ArrayParams // type为数组时，填写
}
// 方法参数
type FunTp struct {
	Name  string   // 方法名
	Param []PostParam // 参数
	Des string     // 描述
	PK string      // 主参数
	Void bool      // 是否有返回
	Class string   // 类名
	Parent string
	Service string
}
// 类参数
type ClassTp struct {
	ClassDesc,ClassName,ClassParent,ServiceName,NameSpace string
}
