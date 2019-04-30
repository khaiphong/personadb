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
//   "fmt"

   "github.com/dgraph-io/badger"
   "github.com/khaiphong/personadb/persona"
)

func main() {
   opts := badger.DefaultOptions
   opts.Dir = "/apt/data"
   opts.ValueDir = "/apt/data"
   kv, err := badger.NewKV(&opts)
   if err != nil {
      panic(err)
   }

   defer kv.Close()

   // we use key first level from Mu. and its value for sorting the keys
   err = kv.Set([]byte("pEip"), []byte("eip"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pChat"), []byte("chat"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pService"), []byte("service"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pHr"), []byte("hr"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pGslp"), []byte("gslp"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pLink"), []byte("link"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pAwakening"), []byte("awakening"))
   if err != nil {
      panic(err)
   }

   // other key first level
   err = kv.Set([]byte("pGit"), []byte("git"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pIoT"), []byte("iot"))
   if err != nil {
      panic(err)
   }
   err = kv.Set([]byte("pAI"), []byte("ai"))
   if err != nil {
      panic(err)
   }
 


}

