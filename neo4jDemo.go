package main

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)
// 保存驱动信息
var neo4jDriver neo4j.Driver
func Init()  error{
	var err error
	// 如果是重新初始化，则县关闭之前的连接
	if neo4jDriver != nil {
		err = neo4jDriver.Close()
		if err != nil {
			return err
		}
	}
	// 连接neo4j
	dbUri := "neo4j://localhost:7687"
	neo4jDriver, err = neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "123456", ""))
	if err != nil {
		panic(err)
	}
	return nil
}
// GetSession 获取session
func GetSession() neo4j.Session {
	return neo4jDriver.NewSession(neo4j.SessionConfig{})
}

type NeoNode struct {
	Name map[string]interface{}
	NextNode []*NeoNode
}

type NN struct {
	Name string
	NNN []*NN
}
func main() {
	err := Init()
	if err != nil {
		fmt.Println("query nearby monitor points failed: %s", err)
	}
	session := GetSession()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		rs, err := tx.Run(
			"match data=(p1:Point)-[:Child*]->(p:Point) where p1.name=$name return data",
			map[string]interface{}{
				"name": "监测点1",
			})
		if err != nil {
			fmt.Println("query nearby monitor points failed: %s", err)
			return nil, err
		}
		var nearbyPoints []map[string]interface{}
		nodeList:=map[int64]map[string]interface{}{}
		var result [][]map[string]interface{}
		i := 0
		relates := make([][]dbtype.Relationship,0)
		// 起点的系统ID
		sNodeId := int64(0)
		for rs.Next() {
			// 输出结果集中的记录
			record := rs.Record()
			//nearbyPoint := record.Values[0].(dbtype.Node)
			//data := nearbyPoint.Props
			//data["relation"] = relation.Type
			//nearbyPoints = append(nearbyPoints, data)
			i +=1
			relates = append(relates, record.Values[0].(dbtype.Path).Relationships)
			//for _, relate := range record.Values[0].(dbtype.Path).Relationships{
			//	fmt.Println(relate.StartId)
			//	fmt.Println(relate.EndId)
			//}
			points := []map[string]interface{}{}
			for _, nd:=range record.Values[0].(dbtype.Path).Nodes{
				if _, ok := nodeList[nd.Id]; !ok {
					nodeList[nd.Id] = map[string]interface{}{
						"name":nd.Props["name"].(string),
						"next":make(map[int64]interface{}),
					}
				}

				points = append(points,nd.Props)

				if pId:= nd.Props["name"].(string); pId == "监测点1"{
					sNodeId = nd.Id //系统分配的节点Id
				}
			}
			paths:= fmt.Sprintf("第%d条路径", i)
			fmt.Println(paths, points)
			result = append(result, points)
		}

		for _,relate := range relates{
			for _, r:= range relate{
				sPoint := nodeList[r.StartId]
				ePoint := nodeList[r.EndId]
				sPoint["next"].(map[int64]interface{})[r.EndId] = ePoint
				nodeList[r.StartId] = sPoint
			}
		}
		fmt.Println(result)
		x,_ := json.Marshal(nodeList[sNodeId])
		fmt.Println(string(x))
		// 处理结果
		return nearbyPoints, nil
	})
	if err != nil {
		fmt.Println("update monitor point graph relation failed: %s", err)
	}
	//fmt.Println(result)
}