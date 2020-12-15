package wg
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB

func init()  {
	db, _ = sql.Open("mysql", "debian-sys-maint:68pIB0llzP77I35h@/zk_local_toc_main?charset=utf8&parseTime=true")
}
func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}
func TableInfo(dbName string) map[string] string {
	sqlStr := `SELECT table_name tableName,TABLE_COMMENT tableDesc
			FROM INFORMATION_SCHEMA.TABLES
			WHERE UPPER(table_type)='BASE TABLE'
			AND LOWER(table_schema) = ?
			ORDER BY table_name asc`
	var result = make(map[string] string)
	rows, err := db.Query(sqlStr, dbName)
	checkErr(err)

	for rows.Next(){
		var tableName, tableDesc string
		err = rows.Scan(&tableName, &tableDesc)
		checkErr(err)
		if len(tableDesc) == 0 {
			tableDesc = tableName
		}
		result[tableName] = tableDesc
	}
	return result
}
type Field struct{
	FieldName string
	FieldDesc string
	DataType  string
	IsNull    string
	length    int
}
func FieldInfos(dbName,tableName string) [] Field{
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns
			WHERE table_schema = ? AND table_name = ?`
	var result [] Field
	rows, err := db.Query(sqlStr,dbName,tableName)
	checkErr(err)
	for rows.Next() {
		var f Field
		err = rows.Scan(&f.FieldName, &f.FieldDesc, &f.DataType, &f.IsNull, &f.length)
		checkErr(err)

		result = append(result, f)
	}
	return result
}
func FieldInfos2(dbName,tableName string) []ArrayParams{
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType
			FROM information_schema.columns
			WHERE table_schema = ? AND table_name = ?`
	var result [] ArrayParams
	rows, err := db.Query(sqlStr,dbName,tableName)
	checkErr(err)
	for rows.Next() {
		var f ArrayParams
		err = rows.Scan(&f.Name, &f.Desc, &f.Type)
		f.Required = false
		f.Default = ""
		checkErr(err)

		result = append(result, f)
	}
	return result
}
func MulTableInfos(dbName string, tables ... string) [] Field{
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns
			WHERE table_schema = ? AND table_name = ?`


	var result [] Field

	for _, tb := range tables{
		rows, err := db.Query(sqlStr,dbName,tb)
		checkErr(err)
	
		for rows.Next() {
			var f Field
			err = rows.Scan(&f.FieldName, &f.FieldDesc, &f.DataType, &f.IsNull, &f.length)
			checkErr(err)
	
			result = append(result, f)
		}
	}
	
	return result
}
func InsertSql(sql string)  {
	_, err := db.Query(sql)
	checkErr(err)
}
func GetDB()* sql.DB {
	return db
}

