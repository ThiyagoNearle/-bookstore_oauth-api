package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session // previuosly cluster *gocql.ClusterCongig
)

func init() {
	/////  CONNECT TO CASSANDRA CLUSTER
	fmt.Println("started-----------------------------000-------------------------------")

	cluster := gocql.NewCluster("127.0.0.1") // var cluster

	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection created successfully")
}

func GetSession() *gocql.Session {
	fmt.Println("started--------------------------222-----------------------------------")

	return session
}
