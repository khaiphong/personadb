/* 
  Copyright (c) 2017, 2018 KhaiPhong
  Om: a [kafka/db] for global registry, each legalEntity, and each publicTopic
  OmHub: [kafka/db] publicTopics filtered by locality of the ThankYou Club 
  MuHub: [kafka/db] globally filtered by legalEntity
*/
package main

import (
   "fmt"
    "io/ioutil"
    "net/http"
    "os"

//      "context"
//      "github.com/ethereum/go-ethereum/common"
//      "github.com/ethereum/go-ethereum/ethclient"
)

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
  // location, actionTopic, qualifier, SIC rating, serverlessFunction string. Initially, 
  // we can use label and tag and put all events in 1 OmHub. Break it out later into topic.
  Extensions map[string][]byte

  //Object *Data
}

// Data ownership of the legal entity while respecting IPR of the service which must
// pass this test to enforce the user legitimate ownership and PersonaAI.
func dataExtraction (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Securely Data extraction!")
}

func composableEvent (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Event Injection, Deletion, Massaging!")
}

//func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprintf(w, "This is the RESTful API")
//}

func main() {
    resp, err := http.Get("https://google.com")
    // check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Println(len(body))
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
/*
func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
 	// the server must keep open
	http.ListenAndServe(":8080", router)
}
*/

