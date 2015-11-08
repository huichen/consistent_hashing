// Copyright (C) 2012 Numerotron Inc.
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/huichen/consistent"
	"log"
	"math/rand"
)

func main() {
	c := consistent.New()
	c.Add("127.9.9.3:3322")
	c.Add("127.9.9.3:3323")
	c.Add("127.9.9.3:3324")
	c.Add("127.9.9.3:3325")
	c.Add("127.9.9.3:3326")
	c.Add("127.9.9.3:3327")
	c.Add("127.9.9.3:3328")
	c.Add("127.9.9.3:3329")
	servers := make(map[string]int32)
	for i := 0; i < 100000; i++ {
		n := rand.Intn(1000000000)
		key := fmt.Sprintf("%d", n)
		server, _ := c.Get(key)
		if v, ok := servers[server]; ok {
			servers[server] = v + 1
		} else {
			servers[server] += 0
		}
	}

	log.Printf("%v", servers)
}
