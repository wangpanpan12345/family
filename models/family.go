package models

import(
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"fmt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"time"
	"strconv"
	"crypto/md5"
)

type Member struct{
	Name string `json:"name"`
	Sex	 int	`json:"sex"` // 1男 2女
	Edu	 string `json:"edu"`
	Birth string	`json:"birth"`
	BF  int `json:"bf"`
	Info string `json:"info"`
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

// son; father; daughter;  mother;
func (m *Member) NewMember(node1 map[string]interface{}, node2,ref string) int64 {

	con := createConnection()
	defer con.Close()

	node1Id := node1["id"].(string)

	result, _ := con.ExecNeo(
		"CREATE (n:Member {name: {name}, sex: {sex},edu:{edu},birth:{birth},bf:{bf},info:{info}})",
	node1) // map[string]interface{}{"name": "1", "sex": 1}
	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	if node2 != "" && ref != ""{
		rel := Member{}
		numR := rel.AddRelation(node1Id,node2,ref)
		fmt.Println(numR)
	}

	return numResult
}

func IdGen() string{

	t := time.Now().UnixNano()

	ts:=strconv.FormatInt(t,10)

	data := []byte(ts)
	has := md5.Sum(data)
	ID := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return ID
}

func ReveseRef(ref string) string{
	refMap := map[string]string{
		"son":"father",
		"daughter":"father",
	}

	return refMap[ref]
}

func (m *Member) AddRelation(node1Id,node2Id,ref string) int64 {


	con := createConnection()
	defer con.Close()
	_ref := ref
	result, _ := con.ExecNeo("MATCH (a:Member),(b:Member) WHERE a.id = "+node1Id+"and b.id="+node2Id+" CREATE (a)-[:"+ref+"]->(b)",nil)
	result, _ = con.ExecNeo("MATCH (a:Member),(b:Member) WHERE a.id = "+node1Id+"and b.id="+node2Id+" CREATE (b)-[:"+_ref+"]->(a)",nil)

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
	fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
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