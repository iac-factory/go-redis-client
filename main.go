package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/go-redis/redis/v9"
)

type Client struct {
	CTX context.Context
	API *redis.Client

	Connection func(ctx context.Context, instance *redis.Conn) error

	Credentials Credentials
}

var Instance *Client = nil

type Command struct {
	String    string
	Arguments []interface{}
	Response  string
}

func (instance Client) evaluate(cmd *redis.StatusCmd, throw bool) Command {
	response, exception := cmd.Result()
	if exception != nil && throw {
		panic(exception)
	}

	return Command{
		String:    cmd.String(),
		Arguments: cmd.Args(),
		Response:  response,
	}
}

func (instance Client) String(key string, value string) Command {
	command := instance.API.Set(instance.CTX, key, value, 0)

	return instance.evaluate(command, true)
}

// Flush - Deletes all keys in all database
func (instance Client) Flush() Command {
	command := instance.API.FlushAll(instance.CTX)

	return instance.evaluate(command, true)
}

// Clear - Deletes the keys in a database
//
// - See Flush for flushing all keys in all databases
func (instance Client) Clear() Command {
	command := instance.API.FlushDB(instance.CTX)

	return instance.evaluate(command, true)
}

func (instance Client) Shutdown() Command {
	command := instance.API.Shutdown(instance.CTX)

	return instance.evaluate(command, false)
}

func API() *Client {
	if Instance == nil {
		panic("Fatal Redis-Client Package Initialization Failure")
	}

	return Instance
}

func init() {
	log.Printf("Initializing Redis-Client Module %s\n", "...")

	ACL.Hydrate()

	var client = redis.NewClient(&redis.Options{
		Addr:                  ":6379",
		Network:               "tcp",
		Username:              *ACL.Username,
		Password:              *ACL.Password,
		ContextTimeoutEnabled: true,
		DB:                    0,
		DialTimeout:           3 * time.Second, // no time unit = seconds
		ConnMaxIdleTime:       30 * time.Second,
		ReadTimeout:           6 * time.Second,
		MaxRetries:            3,
		PoolFIFO:              false,
		PoolSize:              runtime.NumCPU() * 5,
	})

	Instance = &Client{API: client, CTX: context.Background(), Credentials: *ACL}
}

func main() {
	client := API()

	cmd := client.String("J-Key", "J-Value")

	log.Println(cmd)

	client.Clear()
	client.Flush()
	client.Shutdown()
}
