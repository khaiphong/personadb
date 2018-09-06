package storage

import (
    "encoding/json"
    "net/http"

//    "fmt"
//    "github.com/syndtr/goleveldb/leveldb"
//    "log"
)

/*
  Om entity id=0 holds all (1) entities, (2) baskets, (3) topics, (4) events, (5) git.
  Segregation of topics for faster searh. Legal entities: user, organization, service
  Except "entities, buckets, topics, events", the first segment is always entityId,
  the second bucketId, the third topicId, the fourth eventId or B-bucketId. These ids are
  searched from local representations of Om Central Nervous System.
*/
func GetDB(w http.ResponseWriter, r *http.Request) {}

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

/*
    newBucket := make([] Bucket, 0, 1)
    newTopic := make([] Topic, 0, 1)
    newPermission := make([] Event, 0, 1)
    newOccurance := make([] Event, 0, 1)
*/

type Bucket struct {
  ContextAtts map[string][]byte
  T [] Topic
  A [] Event // an act of permissions controlled by the owner
}
type Topic struct {
  /*
    Registered topic _schema to dynamically change and update the Central Nervous System.
    label, tag, serverlessFunction  string | Topic must be in a bucket.
    Producer publishes to a topic and emit a message. Consumer pulls from a topic

    Topics: filtered by place, action, qualifier -> users, orgnizations, services
    employ, teacher, nurse,
    eat, Chinese, Italian,
    shop, gobal, cloth,
    meet, A at TY-AB, xy meeting Calgary,
   */
  ContextAtts map[string][]byte
  O [] Event // Occurance implies both Action and Relationship.
  B [] Bucket
}
type Event struct { 
  /* 
    use goroutines and channels for parallelism and concurrencies.
    In ContextAtts: eventType, eventOwners, eventSource, eventID, schemaURL, contentType string
      public bool
    In Extensions map, use "composableEvent" as serverless function
      location, actionTopic, qualifier, SIC rating, composableEvent string. 
    We can use label and tag and put all events in 1 OmHub. Break it later into topics/markets.
    Enable producers create registered topics. Detention policy 90 days.
  */
  ContextAtts map[string][]byte
  Extensions map[string][]byte
  Data map[string][]byte
}

/*
  Service has its own data schema; the write/read PersonaDB use the database REST API.
  The data is part of the event owned by the service and its user in Settlement layer.

  Entity creates its own tree of buckets, topics, events. Events are in its own database 
  and in the Om for extracting data to the other owner(s), recorded in the Settlement.
*/
func (p *Event) dataExtraction() string {
//   dmap := p.Data
   data := "extract Event.Data to json"
   return data
}

/*
  Serverless function as result of PersonaAI Normative-Positive Intelligence, to change
  the course of the event processes coming from the interaction of streaming and crytographic
  systems of the Om Central Nervous System.

  For complex situation, look at interface HTTP (Event) Handler, taking struct *Request
  (*Event) and use interface ResponseWriter (ComposeEvent) to dispatch to target places
  with channels.
*/
func (p *Event) composeEvent() Event {   
    // process the event extension and inject an event to change its course of action
    var e Event
    return e
}

var buckets []Bucket
func GetEntityBuckets(w http.ResponseWriter, r *http.Request) {}
func GetEntityBucket(w http.ResponseWriter, r *http.Request) {}
func CreateEntityBucket(w http.ResponseWriter, r *http.Request) {}
func DeleteEntityBucket(w http.ResponseWriter, r *http.Request) {}

var topics []Topic
func GetBucketTopics(w http.ResponseWriter, r *http.Request) {}
func GetBucketTopic(w http.ResponseWriter, r *http.Request) {}
func CreateBucketTopic(w http.ResponseWriter, r *http.Request) {}
func DeleteBucketTopic(w http.ResponseWriter, r *http.Request) {}

var events []Event
func GetBucketEvents(w http.ResponseWriter, r *http.Request) {}
func GetBucketEvent(w http.ResponseWriter, r *http.Request) {}
func CreateBucketEvent(w http.ResponseWriter, r *http.Request) {}
func DeleteBucketEvent(w http.ResponseWriter, r *http.Request) {}

func GetTopicEvents(w http.ResponseWriter, r *http.Request) {}
func GetTopicEvent(w http.ResponseWriter, r *http.Request) {}
func CreateTopicEvent(w http.ResponseWriter, r *http.Request) {}
func DeleteTopicEvent(w http.ResponseWriter, r *http.Request) {}

// in later version
func GetTopicBuckets(w http.ResponseWriter, r *http.Request) {}
func GetTopicBucket(w http.ResponseWriter, r *http.Request) {}
func CreateTopicBucket(w http.ResponseWriter, r *http.Request) {}
func DeleteTopicBucket(w http.ResponseWriter, r *http.Request) {}

/*
func storage() {
	db, err := leveldb.OpenFile("/khaiphong/personadb", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()

	//	err = db.Put([]byte("fizz"), []byte("buzz"), nil)
	//	err = db.Put([]byte("fizz2"), []byte("buzz2"), nil)
	//	err = db.Put([]byte("fizz3"), []byte("buzz3"), nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.Seek([]byte("fizz2")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	iter.Release()
	err = iter.Error()
}
*/
