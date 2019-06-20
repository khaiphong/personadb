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
   "net/http"

   "github.com/dgraph-io/badger"
   "github.com/khaiphong/personadb/persona"
)

// use simple web server to test docker development and packaging
var valOwner []byte
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("The owner is: %s\n", valOwner)
//    fmt.Printf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

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
      // new PersonaDB Start a writable transaction. There are 7 leveled key
      //Service/blog.khaiphong.io/190923/right-mindfulness.html->content in SSD
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	txn.Set([]byte("Owner"), persona.PersonaInit())
	if err != nil {
		return err
	}
	fmt.Println("Set Owner of persona data")

	txn.Set([]byte("EIP"), []byte("eip"))
	if err != nil {
		return err
	}
	fmt.Println("Set EIP to eip")
	// eip starts with Free Todo leading to Plan, and PersonaAI
	txn.Set([]byte("EIP/Todo"), []byte("eip/todo"))
	if err != nil {
		return err
	}
	fmt.Println("Set EIP/Todo to eip/todo")
	txn.Set([]byte("EIP/PersonaAI"), []byte("eip/personaai"))
	if err != nil {
		return err
	}
	fmt.Println("Set EIP/PersonaAi to eip/personaai")

	txn.Set([]byte("Chat"), []byte("chat"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat to chat")
	// chat starts with Free BodyMind and private Blog o publised yy/mm/dd
	txn.Set([]byte("Chat/BodyMind"), []byte("chat/bodymind"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat/BodyMind to chat/bodymind")
	txn.Set([]byte("Chat/WWW"), []byte("chat/www"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat/WWW to chat/www")
	txn.Set([]byte("Chat/Blog"), []byte("chat/blog"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat/Blog to chat/blog")
/*
to be added in starting a private blog
	txn.Set([]byte("Chat/Blog/2019-09-23"), []byte("chat/blog/2019-09-23"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat/Blog/2019-09-23 to chat/blog/2019-09-23")
	txn.Set([]byte("Chat/Blog/2019-09-23/Preface"), []byte("chat/blog/2019-09-23/preface"))
	if err != nil {
	    return err
	}
	fmt.Println("Set Chat/Blog/2019-09-23/Preface to chat/blog/2019-09-23/preface")
*/

	txn.Set([]byte("Service"), []byte("service"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service to service")
	// service public #Pub, #EmptyTheContent, #Akp namespaces
	txn.Set([]byte("Service/#Pub"), []byte("service/#pub"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service/#Pub to service/#pub")
	txn.Set([]byte("Service/#AKP"), []byte("service/#akp"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service/#AKP to service/#akp")
	txn.Set([]byte("Service/#EmptyTheContent"), []byte("service/#emptythecontent"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service/#EmptyTheContent to service/#emptythecontent")
	// 9 planned clusters
	txn.Set([]byte("Service/#EmptyTheContent/#PrajnaTIP"), []byte("service/#emptythecontent/#prajnatip"))
	if err != nil {
		return err
	}
	fmt.Println("Set Service/#EmptyTheContent/#PrajnaTIP to service/#emptythecontent/#prajnatip")

	txn.Set([]byte("Hr"), []byte("hr"))
	if err != nil {
		return err
	}
	fmt.Println("Set HR to hr")
	// hr starts with #MarketValue corresponding to public Hr/#MarketValue
	txn.Set([]byte("HR/#MarketValue"), []byte("hr/#marketvalue"))
	if err != nil {
		return err
	}
	fmt.Println("Set HR/#MarketValue to hr/#marketvalue")

	txn.Set([]byte("GSLP"), []byte("gslp"))
	if err != nil {
		return err
	}
	fmt.Println("Set GSLP to gslp")
	// gslp starts with #Calgary corresponding to public GSLP/#Calgary
	// coordination of sustainable supply-chain management
	txn.Set([]byte("GSLP/#Ottawa"), []byte("gslp/#ottawa"))
	if err != nil {
		return err
	}
	fmt.Println("Set GSLP/#Ottawa to gslp/#ottawa")
	txn.Set([]byte("GSLP/#Calgary"), []byte("gslp/#calgary"))
	if err != nil {
		return err
	}
	fmt.Println("Set GSLP/#Calgary to gslp/#calgary")
	txn.Set([]byte("GSLP/#SupplyChain"), []byte("gslp/#supplychain"))
	if err != nil {
		return err
	}
	fmt.Println("Set GSLP/#SupplyChain to gslp/#supplychain")

	txn.Set([]byte("Link"), []byte("link"))
	if err != nil {
		return err
	}
	fmt.Println("Set Link to link")
	// link InnerCircle and cach in the order of browser, db for offline
	txn.Set([]byte("Link/InnerCircle"), []byte("link/innercircle"))
	if err != nil {
		return err
	}
	fmt.Println("Set Link/InnerCircle to link/innercircle")
	txn.Set([]byte("Link/Cache"), []byte("link/cache"))
	if err != nil {
		return err
	}
	fmt.Println("Set Link/Cache to link/cache")

	txn.Set([]byte("Awakening"), []byte("awakening"))
	if err != nil {
		return err
	}
	fmt.Println("Set Awakening to awakening")
	// awakening #Prajna corresponding to public Awakening/#Prajna
	txn.Set([]byte("Awakening/#Prajna"), []byte("awakening/#prajna"))
	if err != nil {
		return err
	}
	fmt.Println("Set Awakening/#Prajna to awakening/#prajna")

	txn.Set([]byte("Git"), []byte("git"))
	if err != nil {
		return err
	}
	fmt.Println("Set Git to git")

	txn.Set([]byte("IoT"), []byte("iot"))
	if err != nil {
		return err
	}
	fmt.Println("Set IoT to iot")

	txn.Set([]byte("Dashboard"), []byte("Dashboard"))
	if err != nil {
		return err
	}
	fmt.Println("Set Dashboard to dashboard")

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
	    return err
	}
      } // end a writable transaction.

      item.Value(func(val []byte) error {
         fmt.Printf("The owner is: %s\n", val)
         valOwner = val
         return nil
      })

      return nil
   }) // end of Get([]byte("Owner"))

   http.HandleFunc("/", handler)
   log.Fatal(http.ListenAndServe(":8081", nil))

} // end of main()

func iterateKeys(fromKey string) {
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
	item := it.Item()
   	k := item.Key()
	err := item.Value(func(v []byte) error {
		fmt.Printf("key=%s, value=%s\n", k, v)
		return nil
	})
	if err != nil {
		return err
    	}
      }
      return nil
   })

   fmt.Println("DB close")
}

