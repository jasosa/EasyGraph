package easygraph

import (
	"flag"
	"log"
	"testing"
)

var serverpath = flag.String("s", ".", "Testing Graphql server path")

func TestMain(m *testing.M) {
	flag.Parse()
	log.Println("Path:", *serverpath)
	m.Run()
}
