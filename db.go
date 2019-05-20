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
   "log"
   "fmt"

   "github.com/dgraph-io/badger"
   "github.com/khaiphong/personadb/persona"
)


func main() {
   opts := badger.DefaultOptions
   opts.Dir = "/tmp/personadb"
   opts.ValueDir = "/tmp/personadb"
   db, err := badger.Open(opts)
   if err != nil {
      log.Fatal(err)
   }

   defer db.Close()

   db.View(func(txn *badger.Txn) error {
      item, err := txn.Get([]byte("Owner"))
      if err != nil {
        // new PersonaDB
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Owner"), persona.PersonaInit())
		if err != nil {
			return err
		}
		fmt.Println("Set Owner of persona data")
		return nil
	})

	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Eip"), []byte("eip"))
		if err != nil {
			return err
		}
		fmt.Println("Set Eip to eip")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Chat"), []byte("chat"))
		if err != nil {
			return err
		}
		fmt.Println("Set Chat to chat")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Service"), []byte("service"))
		if err != nil {
			return err
		}
		fmt.Println("Set Service to service")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Hr"), []byte("hr"))
		if err != nil {
			return err
		}
		fmt.Println("Set Hr to hr")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Gslp"), []byte("gslp"))
		if err != nil {
			return err
		}
		fmt.Println("Set Gslp to gslp")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Link"), []byte("link"))
		if err != nil {
			return err
		}
		fmt.Println("Set Link to link")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Awakening"), []byte("awakening"))
		if err != nil {
			return err
		}
		fmt.Println("Set Awakening to awakening")
		return nil
	})

	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Git"), []byte("git"))
		if err != nil {
			return err
		}
		fmt.Println("Set Git to git")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Iot"), []byte("iot"))
		if err != nil {
			return err
		}
		fmt.Println("Set Iot to iot")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Ai"), []byte("ai"))
		if err != nil {
			return err
		}
		fmt.Println("Set Ai to ai")
		return nil
	})
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("Prometheus"), []byte("prometheus"))
		if err != nil {
			return err
		}
		fmt.Println("Set Prometheus to prometheus")
		return nil
	})

	return err
      }

      item.Value(func(val []byte) error {
         fmt.Printf("The owner is: %s\n", val)
         return nil
      })

      return nil
   }) // end of Get([]byte("Owner"))

} // end of main()



