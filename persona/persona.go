/* 
  Copyright (c) 2017, 2018, 2019 KhaiPhong

  Licensed under the Apache License, Version 2.0 (the "License"); you may not 
  use this file except in compliance with the License which can be copied at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

https://www.youtube.com/watch?v=szOZ3p-5YIc block chain with badger
https://github.com/prometheus/tsdb - implement Prometheus
https://www.youtube.com/watch?v=cHXbYLNa0qQ - Dgraph distributed db
https://www.youtube.com/watch?v=E8-e-3fRHBw - Managing Data in Microservices, applying AI.
*/

package persona

import (
//   "log"
   "fmt"
   "encoding/json"
)

var LegalId, TokenId, Locales, KeyWords, LastAccessTime string

type persona struct {
	LegalId			string	`json:"legalId"`
	Password		string	`json:"password"`
	PhaseForPassword	string	`json:"phaseForPassword"`
	FirstName		string	`json:"firstName"`
	LastName		string	`json:"lastName"`
	Salute			string	`json:"salute"`
	KnownAs			string	`json:"knownAs"`
	Phone			string	`json:"phone"`
	Cell			string	`json:"cell"`
	Email			string	`json:"email"`
	BirthDate		string	`json:"birthDate"`
	HomeCommunity		string	`json:"homeCommunity"`
	CurrentCommunity	string	`json:"currentCommunity"`
	AboutMe			string	`json:"aboutMe"`
	BusinessCard		string	`json:"businessCard"`
	TokenId			string	`json:"tokenId"`
	CreatedTime		string	`json:"createdTime"`
	UpdatedTime		string	`json:"updatedTime"`
	Locales			string	`json:"locales"`
	KeyWords		string	`json:"keyWords"`
	LastAccessTime		string	`json:"lastAccessTime"`
	Photo			[]byte	`json:"photo"`
	VoiceForRecognition	[]byte	`json:"voiceForRecognition"`
	FingerPrints		[]byte	`json:"fingerPrints"`
}

// generate unique legalId in OmHub and tokenId

func PersonaInit() []byte {
   // struct data are converted to []byte using json.Marshall
   personaD := &persona{
	LegalId:		"123",
	Password:		"pass",
	PhaseForPassword:	"myPass",
	FirstName:		"First",
	LastName:		"Last",
	Salute:			"Dr",
	KnownAs:		"Last First",
	Phone:			"6045748712",
	Cell:			"6041234567",
	Email:			"email.google.com",
	BirthDate:		"201225",
	HomeCommunity:		"Calgary",
	CurrentCommunity:	"Vancouver",
	AboutMe:		"Hi",
	BusinessCard:		"LastFirst",
	TokenId:		"1234",
	CreatedTime:		"",
	UpdatedTime:		"",
	Locales:		"English Vietnamese Freanch",
	KeyWords:		"",
	LastAccessTime:		"190923",
	Photo:			[]byte("Photo"),
	VoiceForRecognition:	[]byte("VoiceForRecognition"),
	FingerPrints:		[]byte("FingerPrints")}

	byteEncode, _ := json.Marshal(personaD)
 	fmt.Println(string(byteEncode))
	return byteEncode
}

func encodeByte(bD *[]byte) []byte {
   byteEncode, _ := json.Marshal(bD)
   fmt.Println(string(byteEncode))
   return byteEncode
}


/*
   fmt.Println("\nRunning ITERATE")
   db.View(func(txn *badger.Txn) error {
      opts := badger.DefaultIteratorOptions
      it := txn.NewIterator(opts)
      defer it.Close()

      for it.Rewind(); it.Valid(); it.Next() {
         k := it.Item().Key
         v := it.Item().Value
         fmt.Println("key=%s, value=%s\n", k, v)
      }
      return nil
   })

   fmt.Println("DB close")
*/

/*
type Event struct { 
   
    use goroutines and channels for parallelism and concurrencies.
    ContextAtts: eventType, eventOwners, eventSource, eventID, schemaURL, contentType string
      label, tag, serverlessFunction  string
      dataOwners, composeEvents []string
      public bool
    Extensions map, use "composableEvent" as serverless function
      location, actionTopic, qualifier, SIC rating, composableEvent string.

    We can use label and tag and put all events in 1 OmHub. Break it later into topics/markets. Enable producers create registered topics. Detention policy 90 days.
  
  ContextAtts map[string], Extensions map[string], Data map[string] []byte
}
func (i *Event) Item() struct {
   // Event initialization
   Event.ContextAtts = make(map[string]) []byte("")
   Event.Extensions = make(map[string]) []byte("")
   Event.Data = Make(map[string]) []byte

   return
}

type EventPro interface {
   Evp() Event
}
func (e *Event) Evp() Event {
   // processing the event and return another event
}




// API structure: available directories for person, depths of API for service
type apiStruct1 struct {
    Owner   string
    Layers []string
}
type apiStruct2 struct {
    Owner   string  `json:"owner"`
    Layers []string `json:"layers"`
}

   // Struct data are converted to []byte using json.Marshall
   // Layers are used to build keys for struct values of the service
   api1D := &apiStruct1{
        Owner:   "KhaiPhong",
        Layers: []string{"git", "iot", "ai"}}
   api1B, _ := json.Marshal(api1D)
   fmt.Println(string(api1B))
   // {"Owner":"KhaiPhong","Layers":["git","iot","ai"]}

   api2D := &apiStruct2{
        Owner:   "KhaiPhong",
        Layers: []string{"git", "iot", "ai"}}
   api2B, _ := json.Marshal(api2D)
   fmt.Println(string(api2B))
   // {"owner":"KhaiPhong","layers":["git","iot","ai"]}

   // the []byte value is converted back via Unmarshall
   byt := []byte(`{
		"num":6.13,
		"strs":["a","b"]	}`)
   var dat map[string]interface{}
   if err := json.Unmarshal(byt, &dat); err != nil {
      panic(err)
   }
   fmt.Println(dat)
   // map[num:6.13 strs:[a b]]

   str := `{
		"owner": "KhaiPhong", 
		"layers": ["git", "iot"]
	   }`
   api := apiStruct2{}
   json.Unmarshal([]byte(str), &api)
   fmt.Println(api)
   // {KhaiPhong [git iot]}
   fmt.Println(api.Layers[0])
   // git

*/

