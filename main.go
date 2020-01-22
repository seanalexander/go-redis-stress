package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/seanalexander/go-redis-stress/redisconfig"
)

var rdb *redis.Client

func init() {

	redisconfig.Set()
	fmt.Println("redisconfig.Addr: ", redisconfig.Addr)
	fmt.Println("redisconfig.KeyName: ", redisconfig.KeyName)
	fmt.Println("redisconfig.PoolSize: ", redisconfig.PoolSize)
	fmt.Println("redisconfig.TotalWorkers: ", redisconfig.TotalWorkers)
	fmt.Println("redisconfig.MaxIterations: ", redisconfig.MaxIterations)

	rdb = redis.NewClient(&redis.Options{
		Addr:         redisconfig.Addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     redisconfig.PoolSize,
		PoolTimeout:  30 * time.Second,
	})
}

func main() {
	fmt.Println("keyName, totalWorkers, poolSize, methodName, elapsed_h, elapsed_n")
	if len(redisconfig.KeyName) > 0 {
		benchmarkRedis(redisconfig.KeyName, redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
	} else {
		benchmarkRedis("4kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
		benchmarkRedis("223504kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
		benchmarkRedis("419235kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
		benchmarkRedis("838470kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
		benchmarkRedis("1257705kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
		benchmarkRedis("1676940kb", redisconfig.TotalWorkers, redisconfig.PoolSize, redisconfig.MaxIterations)
	}
}

func benchmarkRedis(keyName string, totalWorkers int, poolSize int, maxIterations int) {
	const methodName = "benchmarkRedis"
	fmt.Printf("\"%s\", %d, %d, \"%s\"", keyName, totalWorkers, poolSize, methodName)
	//fmt.Println()
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

	for workerID := 1; workerID <= totalWorkers; workerID++ {
		go func(workerID int) {
			defer wg.Done()
			for iTerations := 1; iTerations <= maxIterations; iTerations++ {
				//defer timeTrack(time.Now(), fmt.Sprintf("Worker: %d, Loop: %d", workerID, iTerations))
				//fmt.Println(fmt.Sprintf("Worker: %d, Loop: %d, Time: %s", workerID, n, time.Now()))
				if _, err := getStringValue(keyName); err != nil {
					fmt.Println("getStringValue Error: ", err)
				} else {
					//fmt.Println("getStringValue: ", n)
				}
			}
		}(workerID)
	}
	wg.Wait()

	if _, err := rdb.Get(keyName).Result(); err != nil {
		fmt.Println("getStringValue Error: ", err)
	} else {
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf(", %s, %s, %d\n", name, elapsed, elapsed)
	//fmt.Printf(", %s, %d\n", elapsed, elapsed)
}
