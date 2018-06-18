/* Copyright (c) 2017 KhaiPhong = LevelDB and SimpleDB */
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
//      "context"
//      "github.com/ethereum/go-ethereum/common"
//      "github.com/ethereum/go-ethereum/ethclient"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful api")
}

type Bucket struct {
  Email, Label, Tag string
  ListB *Bucket
  ListN *Node
  Action *Event
  Relation *Event
  // triggers of serverless cloud functions
  Source map[string][]byte
//  TimeStamp time.Now()
}
type Node struct {
  Email, Label, Tag  string
  ListB *Bucket
  Action *Event
  Relation *Event
  // triggers of serverless cloud functions
  Source map[string][]byte
//  TimeStamp time.Now()
}
type Event struct { // use https://github.com/cloudevents/spec
  Email, Label, Tag,
    SourceNode, TypeNode, TargetNode, RdfSource, SourceSIC, TargetSIC string
  Rating, Value1, Value2 string
  Action *Node
  // triggers of serverless cloud functions
  Source map[string][]byte
//  TimeStamp time.Now()
}


func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}

