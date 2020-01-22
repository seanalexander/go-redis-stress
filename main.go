package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client
var poolSize int
var totalWorkers int

func init() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "totalWorkers poolSize")
		panic("nope")
	}

	argTotalWorkers := os.Args[1]
	argPoolSize := os.Args[2]

	var err error

	poolSize, err = strconv.Atoi(argPoolSize)
	if err != nil {
		fmt.Println("error: ", err)
		poolSize = 0
	}

	totalWorkers, err = strconv.Atoi(argTotalWorkers)
	if err != nil {
		fmt.Println("error: ", err)
		totalWorkers = 10
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     poolSize,
		PoolTimeout:  30 * time.Second,
	})
}

func main() {
	//exampleClient()
	fmt.Println("keyName, totalWorkers, poolSize, methodName, elapsed_h, elapsed_n")
	benchmarkRedis("419235kb")
	benchmarkRedis("838470kb")
	benchmarkRedis("1257705kb")
	benchmarkRedis("1676940kb")
	//exampleClient_Watch()
}

func benchmarkRedis(keyName string) {
	//const totalWorkers = 2
	const methodName = "benchmarkRedis"
	fmt.Printf("\"%s\", %d, %d, \"%s\"", keyName, totalWorkers, poolSize, methodName)

	defer timeTrack(time.Now(), methodName)

	getStringValue := func(key string) (string, error) {
		n, err := rdb.Get(key).Result()
		if err != nil && err != redis.Nil {
			return n, err
		}
		return n, err
	}

	var wg sync.WaitGroup
	wg.Add(totalWorkers)
	//fmt.Println("totalWorkers: ", totalWorkers)
	for i := 0; i < totalWorkers; i++ {
		go func() {
			defer wg.Done()
			//defer timeTrack(time.Now(), "Inside Loop")

			if _, err := getStringValue(keyName); err != nil {
				fmt.Println("getStringValue Error: ", err)
			} else {
				//fmt.Println("getStringValue: ", n)
			}
		}()
	}
	wg.Wait()

	//n, err := rdb.Get("ALargeStringValue").Result()

	if _, err := rdb.Get(keyName).Result(); err != nil {
		fmt.Println("getStringValue Error: ", err)
	} else {
	}

	// Output: ended with 100 <nil>
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf(", %s, %d\n", elapsed, elapsed)
	//log.Printf("%s took %s", name, elapsed)
}
