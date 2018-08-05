/* Copyright (c) 2017, 2018 KhaiPhong */
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
  // Test the serverlessFunction to JSON decode all event data of this Bucket
  // label, tag, serverlessFunction  string
  ContextAtts map[string][]byte

  ListB *Bucket
  ListN *Node
  // Occurance implies both Action and Relationship
  Occurance *Event
}
type Node struct {
  // Test the serverlessFunction to JSON decode all event data of this Node
  // label, tag, serverlessFunction  string
  ContextAtts map[string][]byte

  ListB *Bucket
  // Occurance implies both Action and Relationship
  Occurance *Event
}
type Event struct { // use https://github.com/cloudevents/spec
  TimeStamp eventTime
  // eventType, cloudEventsVersion, source, eventID, schemaURL, contentType  string
  ContextAtts map[string][]byte
  // sourceNode, typeNode, targetNode, rdfSource, sourceSIC, targetSIC string
  // label, tag, rating, value1, value2, serverlessFunction string -> the extensions
  Extensions map[string][]byte

  Object Data
}

func dataExtraction (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Securely Data extraction!")
}

func composableEvent (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Event Injection, Deletion, Massaging!")
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

