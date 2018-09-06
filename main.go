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

package main

import (
//    "encoding/json"
    "log"
//    "net/http"
    "github.com/gorilla/mux"

    "github.com/khaiphong/personadb/storage"
//      "context"
//      "github.com/ethereum/go-ethereum/common"
//      "github.com/ethereum/go-ethereum/ethclient"
)

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
func CreateEntity(w http.ResponseWriter, r *http.Request) { // for both Create and Update
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
   // the background processing emits event for eventually consistent system
    router := mux.NewRouter()
    router.HandleFunc("/", GetDB).Methods("GET")

    router.HandleFunc("/entities", GetEntities).Methods("GET")
   // move directly to entityId and replace it with entity unique name in the database
    router.HandleFunc("/{entityId}", GetEntity).Methods("GET")
    router.HandleFunc("/{entityId}", CreateEntity).Methods("POST")
    router.HandleFunc("/{entityId}", DeleteEntity).Methods("DELETE")

    router.HandleFunc("/{entityId}/buckets", GetEntityBuckets).Methods("GET")
   // move directly to entity name, bucketId in the database
    router.HandleFunc("/{entityId}/{bucketId}", GetEntityBucket).Methods("GET")
    router.HandleFunc("/{entityId}/{bucketId}", CreateEntityBucket).Methods("POST")
    router.HandleFunc("/{entityId}/{bucketId}", DeleteEntityBucket).Methods("DELETE")

    router.HandleFunc("/{entityId}/{bucketId}/topics", GetBucketTopics).Methods("GET")
   // move directly to entity name, bucketId, topicId in the database
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}", GetBucketTopic).Methods("GET")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}", CreateBucketTopic).Methods("POST")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}", DeleteBucketTopic).Methods("DELETE")

    router.HandleFunc("/{entityId}/{bucketId}/events", GetBucketEvents).Methods("GET")
   // move directly to entity name, bucketId, eventId in the database
    router.HandleFunc("/{entityId}/{bucketId}/{eventId}", GetBucketEvent).Methods("GET")
    router.HandleFunc("/{entityId}/{bucketId}/{eventId}", CreateBucketEvent).Methods("POST")
    router.HandleFunc("/{entityId}/{bucketId}/{eventId}", DeleteBucketEvent).Methods("DELETE")

    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/events", GetTopicEvents).Methods("GET")
    // move directly to entity name, bucketId, topicId, eventId in the database
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/{eventId}", GetTopicEvent).Methods("GET")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/{eventId}", CreateTopicEvent).Methods("POST")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/{eventId}", DeleteTopicEvent).Methods("DELETE")

    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/buckets", GetTopicBuckets).Methods("GET")
    // move directly to entity name, bucketId, topicId, bucketId in the database
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/B-{bucketId}", GetTopicBucket).Methods("GET")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/B-{bucketId}", CreateTopicBucket).Methods("POST")
    router.HandleFunc("/{entityId}/{bucketId}/{topicId}/B-{bucketId}", DeleteTopicBucket).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}

