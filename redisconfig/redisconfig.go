package redisconfig

import (
	"os"

	flaggs "github.com/jessevdk/go-flags"
)

// Package cfg handles the application configuration.
//import "flag"

// Configuration variables.
/*
var poolSize int
var totalWorkers int
var keyName string
*/
var (
	Addr          string //The redisHost:Port
	KeyName       string //The Key Name to Read
	MaxIterations int    //The Maximum amount of Iterations per Worker Thread
	PoolSize      int    //The connection pool size
	TotalWorkers  int    //The Total Worker Threads
	// [..]
)

var opts struct {
	Addr          string `short:"a" long:"addr" description:"redishost:port" required:"false" default:":6379"`
	KeyName       string `short:"k" long:"key" description:"Key Name to Read" required:"false"`
	MaxIterations int    `short:"i" long:"iterations" description:"The Maximum Amount of Iterations" required:"false" default:"1"`
	PoolSize      int    `short:"p" long:"pool" description:"Connection Pool Size" required:"true"`
	TotalWorkers  int    `short:"w" long:"workers" description:"Total Workers" required:"true"`
}

//Set the Redis Configuration Varaibles
func Set() {
	_, err := flaggs.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
	}

	Addr = opts.Addr
	KeyName = opts.KeyName
	MaxIterations = opts.MaxIterations
	PoolSize = opts.PoolSize
	TotalWorkers = opts.TotalWorkers
}

// Set configuration variables from os.Args.
/*
func Set() {
	flag.IntVar(&PoolSize, "poolSize", 100, "Connection Pool size")
	flag.IntVar(&TotalWorkers, "totalWorkers", 100, "Total Workers")
	flag.StringVar(&KeyName, "keyName", "", "Key Name to Read")
	// [..]
}
*/
