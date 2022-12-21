package main

import (
	"log"
	"os"
)

// Environment - ENV variable utility
//
// # Example
//
//	func init() {
//		fmt.Printf("Initializing (%s) Pagination Module: (%s.%s) %s", provider, service, resource, "\n")
//
//		if verbosity, valid := Environment("VERBOSITY"); valid || !(valid) {
//			Declare("VERBOSITY", verbosity, "DEBUG")
//		}
//
//		if verbosity, valid := Environment("AWS_DEFAULT_REGION"); valid || !(valid) {
//			Declare("AWS_DEFAULT_REGION", verbosity, "us-east-2")
//		}
//	}
func Environment(variable string) (*string, bool) {
	if env, ok := os.LookupEnv(variable); ok {
		return &env, true
	} else {
		return nil, false
	}
}

func Declare(variable string, value *string, defaults string) *string {
	if value != nil && *value != "" {
		if exception := os.Setenv(variable, *value); exception != nil {
			log.Fatalf("Environment-Variable-Specification-Exception: (%s) - %s", variable, exception)
		}
	} else {
		if exception := os.Setenv(variable, defaults); exception != nil {
			log.Fatalf("Environment-Variable-Default-Specification-Exception: (%s) - %s", variable, exception)
		}
	}

	if assignment, validation := Environment(variable); validation {
		return assignment
	}

	return nil
}
