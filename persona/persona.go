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
   "fmt"
)

type Persona struct {
   LegalId, Password, PhaseForPassword, FirstName, LastName, Salute, KnownAs, Phone, Cell, Email, BirthDate, HomeCommunity, CurrentCommunity, AboutMe, BusinessCard, TokenId, CreatedTime, UpdatedTime, Locales, KeyWords, LastAccessTime  string
   Photo, VoiceForRecognition, FingerPrints []byte
}
func (i *Persona) Item() struct {
   // Persona initialization
   Persona.LegalId, Persona.Password, Persona.PhaseForPassword, Persona.FirstName, Persona.LastName, Persona.Salute, Persona.KnownAs, Persona.Phone, Persona.Cell, Persona.Email, Persona.BirthDate, Persona.HomeCommunity, Persona.CurrentCommunity, Persona.AboutMe, Persona.BusinessCard, Persona.TokenId, Persona.CreatedTime, Persona.UpdatedTime, Persona.Locales, Persona.KeyWords, Persona.LastAccessTime = ""
   Persona.Photo, Persona.VoiceForRecognition, Persona.FingerPrints = []byte("")

   return
}

// To update item db.Update
// We build up and provide key, value []byte 
func GetItem (key []byte) {

   var item badger.KVItem
   err = kv.Get([]byte(key), &item)
   if err != nil {
      panic(err)
   }

   item, err := decodePersona(item.Value())
   if err != nil {
      panic(err)
   }

   fmt.Println(String(item.Value()))
}
func SetItem (key, value []byte) {
   // assuming badger ia at /app/data directory
   kv, err := badger.Open("/app/data")
   if err != nil {
      panic(err)
   }

   err = kv.Set(key, value)
   if err != nil {
      panic(err)
   }

}
func UpdateItem (func(txn *badger.Txn) error {

}

type Event struct { 
  /* 
    use goroutines and channels for parallelism and concurrencies.
    ContextAtts: eventType, eventOwners, eventSource, eventID, schemaURL, contentType string
      label, tag, serverlessFunction  string
      dataOwners, composeEvents []string
      public bool
    Extensions map, use "composableEvent" as serverless function
      location, actionTopic, qualifier, SIC rating, composableEvent string.

    We can use label and tag and put all events in 1 OmHub. Break it later into topics/markets. Enable producers create registered topics. Detention policy 90 days.
  */
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


type Eip struct {

}
func (i *Eip) Item() struct {
   // Eip initialization

   return
}

type Chat struct {

}
func (i *Chat) Item() struct {
   // Chat initialization

   return
}

type Service struct {

}
func (i *Service) Item() struct {
   // Service initialization

   return
}

type Hr struct {

}
func (i *Hr) Item() struct {
   // Hr initialization

   return
}

type Gslp struct {

}
func (i *Gslp) Item() struct {
   // Gslp initialization

   return
}


type Link struct {

}
func (i *Link) Item() struct {
   // Link initialization

   return
}


type Awakening struct {

}
func (i *Awakening) Item() struct {
   // Awakening initialization

   return
}

type Git struct {

}
func (i *Git) Item() struct {
   // Git initialization

   return
}


type Iot struct {

}
func (i *Iot) Item() struct {
   // Iot initialization

   return
}


type Ai struct {

}
func (i *Ai) Item() struct {
   // Ai initialization

   return
}



