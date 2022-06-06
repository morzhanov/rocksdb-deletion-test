package main

// #include "rocksdb/c.h"
import "C"

import (
	"log"
	"os"
	"strconv"

	"github.com/flier/gorocksdb"
)

const (
	counterFileName = "./COUNTER"
)

func main() {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))

	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	opts.SetLevel0FileNumCompactionTrigger(4)

	db, err := gorocksdb.OpenDb(opts, "/home/vagrant/test-data/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()

	// 1. read file COUNTER
	// 2. convert to int
	// 3. if counter > 0 delete previous value from db (key = counter)
	// 4. increment counter and change file
	// 5. write new value to db (key = counter)
	// 6. stop the script
	// 7. perform those steps in the bash loop

	b, err := os.ReadFile(counterFileName)
	if err != nil {
		log.Fatal(err)
	}
	counterStr := string(b)
	log.Println("counter = ", counterStr)
	if counterStr == "" {
		counterStr = "0"
	}

	if counterStr != "0" {
		err = db.Delete(wo, []byte(counterStr))
		if err != nil {
			log.Fatal(err)
		}
	}

	counter, err := strconv.Atoi(counterStr)
	if err != nil {
		log.Fatal(err)
	}
	counter++
	counterStr = strconv.Itoa(counter)
	os.WriteFile(counterFileName, []byte(counterStr), 0644)

	err = db.Put(wo, []byte(counterStr), []byte(counterStr))
	if err != nil {
		log.Fatal(err)
	}
}

//func test() {
//	k := []byte("foo")
//	v := []byte("bar")
//
//	value, err := db.Get(ro, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("value from db by key %s before putting = %s\n", k, value.Data())
//
//	err = db.Delete(wo, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//	value, err = db.Get(ro, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("value from db by key %s after deletion = %s\n", k, value.Data())
//
//	err = db.Put(wo, k, v)
//	if err != nil {
//		log.Fatal(err)
//	}
//	value, err = db.Get(ro, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("value from db by key %s = %s\n", k, value.Data())
//	defer value.Free()
//	err = db.Delete(wo, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	value, err = db.Get(ro, k)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("value from db by key %s after deletion = %s\n", k, value.Data())
//}
