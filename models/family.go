package models

import (
	"crypto/md5"
	// "encoding/json"
	"fmt"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"io"
	"sort"
	"strconv"
	"time"
)

type Member struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Sex   int64  `json:"sex"` // 1男 2女
	Edu   string `json:"edu"`
	Birth string `json:"birth"`
	BF    int64  `json:"bf"` //辈分
	Info  string `json:"info"`
}

var (
	URI = "bolt://neo4j:880821@localhost:7687"
)

func createConnection() bolt.Conn {
	driver := bolt.NewDriver()
	con, err := driver.OpenNeo(URI)
	handleError(err)
	return con
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func prepareSatement(query string, con bolt.Conn) bolt.Stmt {
	st, err := con.PrepareNeo(query)
	handleError(err)
	return st
}

func queryStatement(st bolt.Stmt) bolt.Rows {
	// Even once I get the rows, if I do not consume them and close the
	// rows, Neo will discard and not send the data
	rows, err := st.QueryNeo(nil)
	handleError(err)
	return rows
}

func consumeMetadata(rows bolt.Rows, st bolt.Stmt) (gps []graph.Path, js QT) {
	// Here we loop through the rows until we get the metadata object
	// back, meaning the row stream has been fully consumed

	var err error
	err = nil

	gps = []graph.Path{}

	for err == nil {
		var row []interface{}
		row, _, err = rows.NextNeo()

		fmt.Println(err)
		if err != nil && err != io.EOF {
			panic(err)
		} else if err != io.EOF {
			gps = append(gps, row[0].(graph.Path))

			// fmt.Printf("PATH: %#v\n", row[0].(graph.Path)) // Prints all paths
		}
	}
	st.Close()
	js = computeRel(gps)
	return
}

type Node struct {
	Name string
	BF   int64
	Son  []string
	ID   string
	Rel  string
}

type QT struct {
	Name  string        `json:"name"`
	Child []interface{} `json:"children,omitempty"`
	// Value string
}

func computeRel(gps []graph.Path) QT {
	// type child struct {
	// 	Name string
	// }

	node := map[string]Node{}

	bf := map[string]int64{}

	// json := map[string][]interface{}{}
	// ch := child{}
	// js := rel{}
	var last string
	for i := 0; i < len(gps); i++ {

		g := gps[i]
		nodes := g.Nodes
		for j := 0; j < len(nodes); j++ {

			ID := nodes[j].Properties["id"].(string)

			n := Node{}
			n.Name = nodes[j].Properties["name"].(string)
			n.BF = nodes[j].Properties["bf"].(int64)
			n.ID = ID
			_, ok := node[ID]

			bf[ID] = n.BF

			if j > 0 {
				r := g.Relationships[j-1].Type
				// fmt.Println(r)
				if r == "father" {

					lastNode := node[last]
					// fmt.Println("father", lastNode)
					fmt.Println("befor", lastNode)
					if lastNode.BF < n.BF {
						if ok {
							tmp := lastNode.Son
							tmp = append(tmp, ID)
							lastNode.Son = tmp
							node[last] = lastNode
						} else {
							tmp := lastNode.Son
							tmp = append(tmp, ID)
							lastNode.Son = tmp
							node[last] = lastNode
							node[n.ID] = n
						}
					} else {
						if ok {
							tNode := node[ID]
							tp := tNode.Son
							tp = append(tp, last)
							tNode.Son = tp
							node[ID] = tNode
						} else {
							n.Son = append(n.Son, last)
							node[n.ID] = n
						}
					}
					// if (last == )
					fmt.Println("add father", n.Name, node)

					// node[last] = lastNode
					// fmt.Println(lastNode)
				}
				if r == "son" || r == "daughter" {
					// fmt.Println("son", n)
					lastNode := node[last]
					fmt.Println("befor", lastNode)
					// fmt.Println("father", lastNode)

					if lastNode.BF < n.BF {
						if ok {
							tmp := lastNode.Son
							tmp = append(tmp, ID)
							lastNode.Son = tmp
							node[last] = lastNode
						} else {
							tmp := lastNode.Son
							tmp = append(tmp, ID)
							lastNode.Son = tmp
							node[last] = lastNode
							node[n.ID] = n
						}
					} else {
						if ok {
							tNode := node[ID]
							tp := tNode.Son
							tp = append(tp, last)
							tNode.Son = tp
							node[ID] = tNode
						} else {
							n.Son = append(n.Son, last)
							node[n.ID] = n
						}

					}
					// node[n.ID] = n
					fmt.Println("add son", n.Name, node)

				}
				if r == "wife" || r == "husband" {

				}

			} else {
				if !ok && ID != "" {
					node[n.ID] = n
				}
			}
			last = ID

		}
		// fmt.Println(g.Nodes)

	}
	fmt.Println(node)
	fmt.Println(bf)
	js := D3(node, bf)
	return js

}

func D3(node map[string]Node, bf map[string]int64) QT {

	hack := map[int]string{}
	hackkeys := []int{}
	for key, val := range bf {
		hack[int(val)] = key
		hackkeys = append(hackkeys, int(val))
	}
	sort.Ints(hackkeys)
	// var l
	js := QT{}
	for _, val := range hackkeys {
		fmt.Println(hack[int(val)], val)

		js = createNode(hack[int(val)], node)
		// res, err := json.Marshal(js)
		// fmt.Println(string(res), err)
		break
		// fmt.Println("gogogogo", t)
		// tmp.Child =

	}

	return js

}

func createNode(id string, node map[string]Node) QT {
	var tmp QT
	nodetmp := node[id]
	tmp.Name = nodetmp.Name + "|" + strconv.FormatInt(nodetmp.BF, 10)
	// tmp.Value = nodetmp.ID
	sons := unique(nodetmp.Son)
	fmt.Println("how many son", len(sons))
	if len(sons) > 0 {
		for i := 0; i < len(sons); i++ {
			var tt QT
			tt = createNode(sons[i], node)
			tmp.Child = append(tmp.Child, tt)
		}

	}
	fmt.Println("goooo", tmp)
	return tmp
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func consumeRows(rows bolt.Rows, st bolt.Stmt) []Member {
	// This interface allows you to consume rows one-by-one, as they
	// come off the bolt stream. This is more efficient especially
	// if you're only looking for a particular row/set of rows, as
	// you don't need to load up the entire dataset into memory

	var err error
	err = nil
	ms := []Member{}
	for err == nil {
		// var row map[string]interface{}
		var data []interface{}
		data, _, err = rows.NextNeo()

		// fmt.Println(err)
		if err != nil && err != io.EOF {
			// panic(err)
			fmt.Println(err)
		} else if err != io.EOF {

			m := Member{}
			m.Name = data[0].(graph.Node).Properties["name"].(string)
			if data[0].(graph.Node).Properties["id"] != nil {
				m.ID = data[0].(graph.Node).Properties["id"].(string)
			}
			if data[0].(graph.Node).Properties["birth"] != nil {
				m.Birth = data[0].(graph.Node).Properties["birth"].(string)
			}
			if data[0].(graph.Node).Properties["bf"] != nil {
				m.BF = data[0].(graph.Node).Properties["bf"].(int64)
			}
			if data[0].(graph.Node).Properties["info"] != nil {
				m.Info = data[0].(graph.Node).Properties["info"].(string)
			}
			if data[0].(graph.Node).Properties["sex"] != nil {
				m.Sex = data[0].(graph.Node).Properties["sex"].(int64)
			}
			if data[0].(graph.Node).Properties["edu"] != nil {
				m.Edu = data[0].(graph.Node).Properties["edu"].(string)
			}

			ms = append(ms, m)
			// fmt.Printf("COLUMNS: %v\n", row) // COLUMNS: n.foo,n.bar
			fmt.Printf("FIELDS: %v\n", data[0].(graph.Node).Properties)
			// fmt.Printf("PATH: %#v\n", row[0].(graph.Path)) // Prints all paths
		}
	}
	st.Close()

	return ms
	// data, _, err := rows.NextNeo()
	// handleError(err)

	// // This query only returns 1 row, so once it's done, it will return
	// // the metadata associated with the query completion, along with
	// // io.EOF as the error
	// _, _, err = rows.NextNeo()
	// handleError(err)

	// // fmt.Println(rows, data)
	// fmt.Printf("COLUMNS: %#v\n", rows.Metadata()["fields"].([]interface{})) // COLUMNS: n.foo,n.bar
	// fmt.Printf("FIELDS: %v %f\n", data[0])                                  // FIELDS: 1 2.2

	// st.Close()
}

func (m *Member) ListMember() {
	con := createConnection()
	defer con.Close()
	RelationNode := "MATCH path=(n:Member)-[:SON|:FATHER*]->(m) where n.name='王一' RETURN path"
	st := prepareSatement(RelationNode, con)
	rows := queryStatement(st)
	// fmt.Println(rows)
	consumeMetadata(rows, st)

}

func (m *Member) ListG3Member(id string) ([]graph.Path, QT) {
	con := createConnection()
	defer con.Close()
	RelationNode := "MATCH (p1:Member),(p2),path = shortestpath((p1)-[*1..2]-(p2)) where p1.id='" + id + "' and  p1.id<>p2.id RETURN path"
	// RelationNode := "MATCH path=(n:Member)-[*1..2]->(m) where n.id='" + id + "' RETURN path"
	st := prepareSatement(RelationNode, con)
	rows := queryStatement(st)
	// fmt.Println(rows)
	gps, js := consumeMetadata(rows, st)

	return gps, js
}

func (m *Member) ComputeMember(id1, id2 string) ([]graph.Path, QT) {
	con := createConnection()
	defer con.Close()
	RelationNode := "MATCH (p1:Member),(p2:Member) ,path = shortestpath((p1)-[*]-(p2)) where p1.id='" + id1 + "' and  p2.id='" + id2 + "' RETURN path"
	// RelationNode := "MATCH path=(n:Member)-[*1..2]->(m) where n.id='" + id + "' RETURN path"
	st := prepareSatement(RelationNode, con)
	rows := queryStatement(st)
	// fmt.Println(rows)
	gps, js := consumeMetadata(rows, st)

	return gps, js
}

func (m *Member) GetMember(name string) []Member {
	con := createConnection()
	defer con.Close()
	GetNode := "match(m:Member) where m.name =~'.*" + name + ".*' return m Limit 10"
	// GetNode := "MATCH (m:Member) where m.name='" + name + "' return m"
	st := prepareSatement(GetNode, con)
	rows := queryStatement(st)

	mss := consumeRows(rows, st)
	return mss
}

// son; father; daughter;  mother;
func (m *Member) NewMember(node1 map[string]interface{}, node2 string, ref []string) (nodeNum, refNum int64, err error) {

	con := createConnection()
	defer con.Close()

	node1Id := node1["id"].(string)

	result, err := con.ExecNeo(
		"CREATE (n:Member {id:{id},name: {name}, sex: {sex},edu:{edu},birth:{birth},bf:{bf},info:{info}})",
		node1) // map[string]interface{}{"name": "1", "sex": 1}
	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	var numR int64
	if node2 != "" && len(ref) == 2 {
		rel := Member{}
		numR = rel.AddRelation(node1Id, node2, ref)
		fmt.Println(numR)
	}

	return numResult, numR, err
}

func (m *Member) IdGen() string {

	t := time.Now().UnixNano()

	ts := strconv.FormatInt(t, 10)

	data := []byte(ts)
	has := md5.Sum(data)
	ID := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return ID
}

func ReveseRef(ref string) string {
	refMap := map[string]string{
		"son":      "father",
		"daughter": "father",
	}

	return refMap[ref]
}

func (m *Member) AddRelation(node1Id, node2Id string, ref []string) int64 {

	con := createConnection()
	defer con.Close()
	// _ref := ref
	result, err := con.ExecNeo("MATCH (a:Member),(b:Member) WHERE a.id ='"+node1Id+"' and b.id='"+node2Id+"' CREATE (a)-[:"+ref[0]+"]->(b)", nil)

	result, err = con.ExecNeo("MATCH (a:Member),(b:Member) WHERE a.id = '"+node1Id+"' and b.id='"+node2Id+"' CREATE (b)-[:"+ref[1]+"]->(a)", nil)

	fmt.Println(err)
	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	return numResult
}

func (m *Member) QuickNDirty() {
	driver := bolt.NewDriver()
	conn, _ := driver.OpenNeo("bolt://neo4j:880821@localhost:7687")
	defer conn.Close()

	fmt.Println(conn)
	// Start by creating a node
	result, _ := conn.ExecNeo("CREATE (n:NODE {foo: {foo}, bar: {bar}})", map[string]interface{}{"foo": 1, "bar": 2.2})
	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	// Lets get the node
	data, rowsMetadata, _, _ := conn.QueryNeoAll("MATCH (n:NODE) RETURN n.foo, n.bar", nil)
	fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))    // COLUMNS: n.foo,n.bar
	fmt.Printf("FIELDS: %d %f\n", data[0][0].(int64), data[0][1].(float64)) // FIELDS: 1 2.2

	// oh cool, that worked. lets blast this baby and tell it to run a bunch of statements
	// in neo concurrently with a pipeline
	results, _ := conn.ExecPipeline([]string{
		"MATCH (n:NODE) CREATE (n)-[:REL]->(f:FOO)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(b:BAR)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(z:BAZ)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(f:FOO)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(b:BAR)",
		"MATCH (n:NODE) CREATE (n)-[:REL]->(z:BAZ)",
	}, nil, nil, nil, nil, nil, nil)
	for _, result := range results {
		numResult, _ := result.RowsAffected()
		fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 2 (per each iteration)
	}

	data, _, _, _ = conn.QueryNeoAll("MATCH (n:NODE)-[:REL]->(m) RETURN m", nil)
	for _, row := range data {
		fmt.Printf("NODE: %#v\n", row[0].(graph.Node)) // Prints all nodes
	}

	result, _ = conn.ExecNeo(`MATCH (n) DETACH DELETE n`, nil)
	numResult, _ = result.RowsAffected()
	fmt.Printf("Rows Deleted: %d \n", numResult) // Rows Deleted: 13
}
