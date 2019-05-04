/* 
  Copyright (c) 2017, 2018, 2019 KhaiPhong

  Licensed under the Apache License, Version 2.0 (the "License"); you may not 
  use this file except in compliance with the License which can be copied at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*/

package main

import (
   "encoding/json"
   "fmt"

   "github.com/dgraph-io/badger"
   "github.com/khaiphong/personadb/persona"
)

type Itemer interface {
   Item() struct
}

func main() {
   opts := badger.DefaultOptions
   opts.Dir = "/apt/data"
   opts.ValueDir = "/apt/data"
   db, err := badger.Open(&opts)
   Handle(err)

   err := db.Update(func(txn *badger.Txn) error {
      if _, err := txn.Get([]byte("pEip")); err == badger.ErrKeyNotFound
         fmt.Println("No existing pEip found or no db")
         // create new database
         db, err := badger.NewKV(&opts)
         if err != nil {
            panic(err)
         }

         // we use key first level from Mu. and its value for sorting the keys
         err = db.Set([]byte("pEip"), []byte("eip"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pChat"), []byte("chat"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pService"), []byte("service"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pHr"), []byte("hr"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pGslp"), []byte("gslp"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pLink"), []byte("link"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pAwakening"), []byte("awakening"))
         if err != nil {
            panic(err)
         }


         // other key first level
         err = db.Set([]byte("pGit"), []byte("git"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pIoT"), []byte("iot"))
         if err != nil {
            panic(err)
         }
         err = db.Set([]byte("pAI"), []byte("ai"))
         if err != nil {
            panic(err)
         }
 
      } // end of if _
   }) // end of db.Update

   var i Itemer
   func (i Itemer) EncodeItemer() []byte {
      data, err := json.Marshall(i)
      if err != nil {
        panic(err)
      }

      return data
   }
   func DecodeItemer(data []byte) (Itemer, error) {
      var i Itemer
      err := json.Unmarshal(data, &i)
      return i, err
   }

   defer db.Close()



}

