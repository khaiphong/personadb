/* 
  Copyright (c) 2017, 2018 KhaiPhong
  Om: a [kafka/db] for global registry, each legalEntity, and each publicTopic
  OmHub: [kafka/db] publicTopics filtered by locality of the ThankYou Club 
  MuHub: [kafka/db] globally filtered by legalEntity
  Samadhi-Prajna in Machine Learning and Deep Learning from layered Neural Networks

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License which can be copied at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"

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
  Service can have its own data schema; the write/read PersonaDB use the database REST API.
  Services being parts of Om Central Nervous System must pass this test to enforce the user
  legitimate ownership for its PersonaAI. The data is part of the event owned by the service
  and its user of the public occurance.
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

  For complex situation, look at interface HTTP (Event) Handler, taking struct *Request
  (*Event) and use interface ResponseWriter (ComposeEvent) to dispatch to target places
  with channels.
*/
func (p *Event) composeEvent() Event {   
    // process the event extension and inject an event to change its course of action
    var e Event
    return e
}

/*
    newBucket := make([] Bucket, 0, 1)
    newTopic := make([] Topic, 0, 1)
    newPermission := make([] Event, 0, 1)
    newOccurance := make([] Event, 0, 1)
*/

type Entity struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var entities []Entity

func GetEntities(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(entities)
}

func GetEntity(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range entities {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Entity{})
}

func CreateEntity(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var entity Entity
    _ = json.NewDecoder(r.Body).Decode(&entity)
    entity.ID = params["id"]
    entities = append(entities, entity)
    json.NewEncoder(w).Encode(entities)
}

func DeleteEntity(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range entities {
        if item.ID == params["id"] {
            entities = append(entities[:index], entities[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(entities)
    }
}

func main() {
    entities = append(entities, Entity{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    entities = append(entities, Entity{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    entities = append(entities, Entity{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

    router := mux.NewRouter()
    router.HandleFunc("/entities", GetEntities).Methods("GET")
    router.HandleFunc("/entities/{id}", GetEntity).Methods("GET")
    router.HandleFunc("/entities/{id}", CreateEntity).Methods("POST")
    router.HandleFunc("/entities/{id}", DeleteEntity).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}

