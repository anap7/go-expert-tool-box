package config

import (
	"fmt"
)

var (
	// String connection with MySQL
	ConnectionString = ""
)

// Load all environments variables
func Load() {
	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"golang",
		"Golang_17396",
		"orders",
	)
}