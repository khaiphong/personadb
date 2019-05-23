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
https://www.gameplan.global/ we go beyond task managements, project planning, and cost optimization with visual white board, generation of Gantt chart, and integration of SIS to dynamically change oneself and the environment toward What Count.
*/

package persona

import (
//   "log"
   "fmt"
   "encoding/json"
)

var LegalId, TokenId, Locales, KeyWords, LastAccessTime string

type persona struct {
	FirstName		string	`json:"firstName"`
	LastName		string	`json:"lastName"`
	Salute			string	`json:"salute"`
	KnownAs			string	`json:"knownAs"`
	Phone			string	`json:"phone"`
	Cell			string	`json:"cell"`
	BirthDate		string	`json:"birthDate"`
	HomeCommunity		string	`json:"homeCommunity"`
	CurrentCommunity	string	`json:"currentCommunity"`
	AboutMe			string	`json:"aboutMe"`
	BusinessCard		string	`json:"businessCard"`
	CreatedTime		string	`json:"createdTime"`
	UpdatedTime		string	`json:"updatedTime"`
	Locales			string	`json:"locales"`
	KeyWords		string	`json:"keyWords"`
	LastAccessTime		string	`json:"lastAccessTime"`
	Photo			[]byte	`json:"photo"`

	voiceForRecognition	[]byte	`json:"voiceForRecognition"`
	fingerPrints		[]byte	`json:"fingerPrints"`
	legalId			string	`json:"legalId"`
	password		string	`json:"password"`
	phaseForPassword	string	`json:"phaseForPassword"`
	tokenId			string	`json:"tokenId"`
	email			string	`json:"email"`
}

// generate unique legalId in OmHub and tokenId

func PersonaInit() []byte {
   // struct data are converted to []byte using json.Marshall
   personaD := &persona{
	FirstName:		"First",
	LastName:		"Last",
	Salute:			"Dr",
	KnownAs:		"Last First",
	Phone:			"6045748712",
	Cell:			"6041234567",
	BirthDate:		"201225",
	HomeCommunity:		"Calgary",
	CurrentCommunity:	"Edmonton",
	AboutMe:		"Hi",
	BusinessCard:		"LastFirst",
	CreatedTime:		"",
	UpdatedTime:		"",
	Locales:		"English Vietnamese French",
	KeyWords:		"",
	LastAccessTime:		"190923",
	Photo:			[]byte("Photo"),

	voiceForRecognition:	[]byte("VoiceForRecognition"),
	fingerPrints:		[]byte("FingerPrints"),
	legalId:		"newID123",
	password:		"pass",
	phaseForPassword:	"myPass",
	tokenId:		"34567",
	email:			"email.google.com"}

	byteEncode, _ := json.Marshal(personaD)
 	fmt.Println(string(byteEncode))
	return byteEncode
}

func encodeByte(bD *[]byte) []byte {
   byteEncode, _ := json.Marshal(bD)
   fmt.Println(string(byteEncode))
   return byteEncode
}

// goroutine and channel for parallelism and concurrencies.
// Producers create registered topics. Detention policy 90 days.
type Event struct {
	// ContextAtts
	EventType		string	`json:"eventType"`
	Owners			string	`json:"owners"`
	Source			string	`json:"source"`
	ID			string	`json:"iD"`
	SchemaURL		string	`json:"schemaURL"`
	ContentType		string	`json:"contentType"`
	// Extensions: label and Tag for all events, then
	Label			string	`json:"label"`
	Tag			string	`json:"tag"`
	DataOwners		string	`json:"dataOwners"`
	// break into topics/markets. ServerlessFunction process ComposeEvents
	ServerlessFunction	string	`json:"serverlessFunction"`
	ComposeEvents		string	`json:"composeEvents"`
	Public			bool	`json:"public"`
	// Data
	Location		string	`json:"location"`
	ActionTopic		string	`json:"actionTopic"`
	Qualifier		string	`json:"qualifier"`
	SIC			string	`json:"sIC"`
	Rating			string	`json:"rating"`

	ContextAtts	map[string]	[]byte	`json:"contextAtts"`
	Extensions	map[string]	[]byte	`json:"extensions"`
	Data		map[string]	[]byte	`json:"data"`
}

func EventInit() []byte {
   // struct data are converted to []byte using json.Marshall
   eventD := &Event{
	EventType:		"123",
	Owners:			"abc, def",
	Source:			"abc",
	ID:			"456",
	SchemaURL:		"https://service.mydomain.io",
	ContentType:		"Chat",

	Label:			"Private",
	Tag:			"Business",
	DataOwners:		"abc, def",
	ServerlessFunction:	"CreditChecking",
	ComposeEvents:		"Calgary, Edmonton",
	Public:			true,

	Location:		"Calgary",
	ActionTopic:		"ApprovalPending",
	Qualifier:		"1234",
	SIC:			"11169",
	Rating:			"",

	ContextAtts:		make(map[string] []byte),
	Extensions:		make(map[string] []byte),
	Data:			make(map[string] []byte) }

	byteEncode, _ := json.Marshal(eventD)
 	fmt.Println(string(byteEncode))
	return byteEncode
}

// different function process Evp() Event
type EventPro interface {
   Evp() Event
}
/*
func (e *Event) Evp() Event {
   // processing the event and return another event
   return
}
*/
