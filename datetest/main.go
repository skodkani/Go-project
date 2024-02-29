package main

import (
	"fmt"
	"time"

	"github.com/wagslane/go-tinytime"
)

func main() {
	// tinytime package uses less memory to store the time.
	// it uses 4 bytes of memory instead of 24 bytes from std time.time()
	// it only supports years from 1970 to 2106.
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)

	fmt.Println(tt)
}
