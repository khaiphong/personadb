/* 
  Copyright (c) 2017, 2018 KhaiPhong
  Om: a [kafka/db] for global registry, each legalEnty, and each publicTopic
  OmHub: [kafka/db] publicTopics filtered by locality of the ThankYou Club 
  MuHub: [kafka/db] globally filtered by legalEntity
*/
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
  // Bucket <-> Bucket | Test the serverlessFunction to JSON decode all event data of this Bucket
  // label, tag, serverlessFunction  string
  ContextAtts map[string][]byte

  ListB *Bucket
  ListT *Topic
  // legal permissions
  Grant *Event
}
type Topic struct {
  // Message <- Topic -> Event
  // Test the serverlessFunction to JSON decode all event data of this Node
  // label, tag, serverlessFunction  string
  /*  
    Topics: filtered by place, actionTopic, qualifier -> users, orgnizations, services
    employ, employTeacher, enployNurse,
    eat, eatChinese, eatItalian,
    shop, shopGobal, shopCloth,
   */
  ContextAtts map[string][]byte

  // Producer publishes to a topic. Consumer pulls from a topic
  ListB *Bucket
  ListM *Event
  // Occurance implies both Action and Relationship.
  Occurance *Event
}
type Event struct { 
  // use https://github.com/cloudevents/spec
  // eventType, cloudEventsVersion, source, eventID, schemaURL, contentType  string
  ContextAtts map[string][]byte
  // location, actionTopic, qualifier, SIC string
  // label, tag, rating, value1, value2, serverlessFunction string -> the extensions
  Extensions map[string][]byte

  Object Data
}

// Data ownership of the legal entity while respecting IPR of the service which must
// pass this test to enforce the user legitimate ownership and PersonaAI.
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

