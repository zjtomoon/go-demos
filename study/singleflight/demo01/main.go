package main

import (
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)

var singleFlightTest singleflight.Group

var errorNotExist = errors.New("redis: key not found")

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// data, err := getData("2000")
			data, err := getDataBySingleFlight("2000")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {

		data, err = getDataFromDB(key)
		if err != nil {
			log.Println(err)
			return "", err
		}

	} else if err != nil {
		return "", err
	}
	return data, nil
}

func getDataBySingleFlight(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		v, err, _ := singleFlightTest.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		data = v.(string)
	} else if err != nil {
		return "", err
	}
	return data, nil
}

func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from database", key)
	return "2000 in db", nil
}
