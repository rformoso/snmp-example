// MIT License

// Copyright (c) 2018 Robson Garcia Formoso

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rformoso/snmp-example/example"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   %s [command] \n", filepath.Base(os.Args[0]))
		fmt.Printf("     walk      - to run the walk/scan\n")
		fmt.Printf("     get       - to get MIB/Oid from a specific host.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) > 1 || len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	target := flag.Args()[0]

	switch target {
	case "get":
		getersonSnmp := &example.GetersonSnmp{}
		getersonSnmp.Run()
	case "walk":
		bulkWalkSnmp := &example.BulkWalkSnmp{}
		bulkWalkSnmp.Run()
	default:
		flag.Usage()
		os.Exit(1)
	}

}
