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
        // new PersonaDB Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	txn.Set([]byte("Owner"), persona.PersonaInit())
	if err != nil {
		return err
	}
	fmt.Println("Set Owner of persona data")

	txn.Set([]byte("Eip"), []byte("eip"))
	if err != nil {
		return err
	}
	fmt.Println("Set Eip to eip")
	txn.Set([]byte("Chat"), []byte("chat"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat to chat")
	txn.Set([]byte("Service"), []byte("service"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service to service")
	txn.Set([]byte("Hr"), []byte("hr"))
	if err != nil {
		return err
	}
	fmt.Println("Set Hr to hr")
	txn.Set([]byte("Gslp"), []byte("gslp"))
	if err != nil {
		return err
	}
	fmt.Println("Set Gslp to gslp")
	txn.Set([]byte("Link"), []byte("link"))
	if err != nil {
		return err
	}
	fmt.Println("Set Link to link")
	txn.Set([]byte("Awakening"), []byte("awakening"))
	if err != nil {
		return err
	}
	fmt.Println("Set Awakening to awakening")

	txn.Set([]byte("Git"), []byte("git"))
	if err != nil {
		return err
	}
	fmt.Println("Set Git to git")
	txn.Set([]byte("Iot"), []byte("iot"))
	if err != nil {
		return err
	}
	fmt.Println("Set Iot to iot")
	txn.Set([]byte("Iot"), []byte("iot"))
	if err != nil {
		return err
	}
	fmt.Println("Set Iot to iot")
	txn.Set([]byte("Prometheus"), []byte("prometheus"))
	if err != nil {
		return err
	}
	fmt.Println("Set Prometheus to prometheus")

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
	    return err
	}
      } // end a writable transaction.

      item.Value(func(val []byte) error {
         fmt.Printf("The owner is: %s\n", val)
         return nil
      })

      return nil
   }) // end of Get([]byte("Owner"))

} // end of main()

func iterateKeys(rootKey string) {
   fmt.Println("\nRunning ITERATE")
   opts := badger.DefaultOptions
   opts.Dir = "/tmp/personadb"
   opts.ValueDir = "/tmp/personadb"
   db, err := badger.Open(opts)
   if err != nil {
      log.Fatal(err)
   }
   defer db.Close()

   db.View(func(txn *badger.Txn) error {
      opts := badger.DefaultIteratorOptions
      it := txn.NewIterator(opts)

      for it.Rewind(); it.Valid(); it.Next() {
         k := it.Item().Key
         v := it.Item().Value
         fmt.Println("key=%s, value=%s\n", k, v)
      }
      return nil
   })

   fmt.Println("DB close")
}

