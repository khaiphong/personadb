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
  // length of B, T, P, label, tag, serverlessFunction  string
  ContextAtts map[string][]byte
  B [] Bucket
  T [] Topic
  P [] Event // legal permissions
}

// Registered topic _schema to dynamically change and update the Central Nervous System
type Topic struct {
  /*
    Message <- Topic -> Event
    Test the serverlessFunction to JSON decode all event data of this Node
    length of B, T, O, label, tag, serverlessFunction  string
    Producer publishes to a topic and emit a message. Consumer pulls from a topic

    Topics: filtered by place, actionTopic, qualifier -> users, orgnizations, services
    employ, employTeacher, enployNurse,
    eat, eatChinese, eatItalian,
    shop, shopGobal, shopCloth,
   */
  ContextAtts map[string][]byte
  B [] Bucket
  T [] Topic
  O [] Event // Occurance implies both Action and Relationship.
}

// use goroutines and channels for parallelism and concurrencies
type Event struct { 
  /* 
    use https://github.com/cloudevents/spec
    eventType, cloudEventsVersion, source, eventID, schemaURL, contentType  string
    use "composableEvent" having (fn Event) value as serverless function in Extensions map
    location, actionTopic, qualifier, SIC rating, composableEvent string. Initially, 
    we can use label and tag and put all events in 1 OmHub. Break it out later into topic.
  */
  ContextAtts map[string][]byte
  Extensions map[string][]byte
  Data map[string][]byte
}

/*
  User of can have their own data schema, but write/read PersonaDB using the database REST API
  Data ownership of the legal entity while respecting IPR of the service which must
  pass this test to enforce the user legitimate ownership and PersonaAI. The data is part
  of the event owned by the service and its user of the public occurance.
*/
func (p *Event) dataExtraction() string {
//   dmap := p.Data
   data := "extract Event.Data to json"
   return data
}

/*
  A serverless function f, as result of PersonaAI Normative-Positive Intelligence, to
  change the course of the event processes coming from the interaction of streaming and
  crytographic systems of the Om Central Nervous System. The composableEvent takes the
  passed (Event) to process.
*/
func (p *Event) composableEvent() Event {   
    // process the event extension and inject an event to change its course of action
    var e Event
    return e
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

    newBucket := make([] Bucket, 0, 1)
    newTopic := make([] Topic, 0, 1)
    newPermission := make([] Event, 0, 1)
    newOccurance := make([] Event, 0, 1)

*/

