/* 
  Copyright (c) 2017, 2018, 2019 KhaiPhong

  Licensed under the Apache License, Version 2.0 (the "License"); you may not 
  use this file except in compliance with the License which can be copied at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

https://github.com/prometheus/tsdb - implement Prometheus
https://www.youtube.com/watch?v=cHXbYLNa0qQ - Dgraph distributed db
https://www.youtube.com/watch?v=E8-e-3fRHBw - Managing Data in Microservices, applying AI.
*/

package persona

import (
   "encoding/json"
   "fmt"
)

type Persona struct {
   LegalId, Password, PhaseForPassword, FirstName, LastName, Salute, KnownAs, Phone, Cell, Email, BirthDate, HomeCommunity, CurrentCommunity, AboutMe, BusinessCard, TokenId, CreatedTime, UpdatedTime, Locales, KeyWords, LastAccessTime  string
   Photo, VoiceForRecognition, FingerPrints []byte
}
func (p Persona) EncodePersona() []byte {
   data, err := json.Marshall(p)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodePersona(data []byte) (Persona, error) {
   var p Persona
   err := json.Unmarshal(data, &p)
   return p, err
}

// the call function build up and provide [] byte key and value
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
// the call function build up and provide []byte key and value
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

type Event struct { 
  /* 
    use goroutines and channels for parallelism and concurrencies.
    In ContextAtts: eventType, eventOwners, eventSource, eventID, schemaURL, contentType string
      label, tag, serverlessFunction  string | dataOwners, composeEvents [] string
      public bool
    In Extensions map, use "composableEvent" as serverless function
      location, actionTopic, qualifier, SIC rating, composableEvent string. 
    We can use label and tag and put all events in 1 OmHub. Break it later into topics/markets.
    Enable producers create registered topics. Detention policy 90 days.
  */
  ContextAtts map[string], Extensions map[string], Data map[string] []byte
}

type EventPro interface {
   Evp() Event
}
func (e *Event) Evp() Event {
   // processing the event and return another event
}



/*
type Eip struct {

}
func (e Eip) EncodeEip() []byte {
   data, err := json.Marshall(e)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeEip(data []byte) (Eip, error) {
   var e Eip
   err := json.Unmarshal(data, &e)
   return e, err
}

type Chat struct {

}
func (c Chat) EncodeService() []byte {
   data, err := json.Marshall(c)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeChat(data []byte) (Chat, error) {
   var c Chat
   err := json.Unmarshal(data, &c)
   return c, err
}

type Service struct {

}
func (s Service) EncodeService() []byte {
   data, err := json.Marshall(s)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeService(data []byte) (Service, error) {
   var s Service
   err := json.Unmarshal(data, &s)
   return s, err
}

type Hr struct {

}
func (h Hr) EncodeHr() []byte {
   data, err := json.Marshall(h)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeHr(data []byte) (Hr, error) {
   var h Hr
   err := json.Unmarshal(data, &h)
   return h, err
}

type Gslp struct {

}
func (g Gslp) EncodeGslp() []byte {
   data, err := json.Marshall(g)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeGslp(data []byte) (Gslp, error) {
   var g Gslp
   err := json.Unmarshal(data, &g)
   return g, err
}

type Link struct {

}
func (l Link) EncodeLink() []byte {
   data, err := json.Marshall(l)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeLink(data []byte) (Link, error) {
   var l Link
   err := json.Unmarshal(data, &l)
   return l, err
}

type Awakening struct {

}
func (a Awakening) EncodeAwakening() []byte {
   data, err := json.Marshall(a)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeAwakening(data []byte) (Awakening, error) {
   var a Awakening
   err := json.Unmarshal(data, &a)
   return a, err
}

type Git struct {

}
func (g Git) EncodeGit() []byte {
   data, err := json.Marshall(g)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeGid(data []byte) (Git, error) {
   var g Git
   err := json.Unmarshal(data, &g)
   return g, err
}

type Iot struct {

}
func (i Iot) EncodeIot() []byte {
   data, err := json.Marshall(i)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeIot(data []byte) (Iot, error) {
   var i Iot
   err := json.Unmarshal(data, &i)
   return i, err
}

type Ai struct {

}
func (a Ai) EncodeAi() []byte {
   data, err := json.Marshall(a)
   if err != nil {
     panic(err)
   }

   return data
}
func DecodeAi(data []byte) (Ai, error) {
   var a Ai
   err := json.Unmarshal(data, &a)
   return a, err
}

*/

