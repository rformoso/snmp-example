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

package example

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/rformoso/snmp-example/util"
	"github.com/soniah/gosnmp"
)

type GetSnmp interface {
	Run()
}

type GetersonSnmp struct {
}

//Get is an example of how to get the value of a given OID.
//Using the default SNMP connection parameters of Gosnmp
func (g *GetersonSnmp) Get(host string, oids []string, community string, wg *sync.WaitGroup, resp chan<- string) {
	defer wg.Done()

	//Gosnmp provides a standard structure containing commonly used values.
	//However, to use Goroutines a GoSNMP object must be created for each process, otherwise a memory violation occurs.
	params := &gosnmp.GoSNMP{
		Target:    host,
		Port:      uint16(161),
		Community: "public",
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(500) * time.Millisecond,
		Retries:   1,
		MaxOids:   gosnmp.MaxOids,
	}

	err := params.Connect()
	if err != nil {
		return
	}
	defer params.Conn.Close()

	result, err := params.Get(oids)
	if err != nil {
		return
	}

	var buffer bytes.Buffer
	for _, variable := range result.Variables {
		switch variable.Type {
		case gosnmp.OctetString:
			if len(variable.Value.([]byte)) != 0 {
				mac := util.ValidateMAC(variable.Value.([]byte))
				if mac != "" {
					buffer.WriteString(util.FormatLog(host, variable.Name, fmt.Sprintf("MAC address=%s\n", mac)))
					break
				}
				buffer.WriteString(util.FormatLog(host, variable.Name, fmt.Sprintf("string=%x\n", string(variable.Value.([]byte)))))
			}
		case gosnmp.Integer:
			buffer.WriteString(util.FormatLog(host, variable.Name, fmt.Sprintf("number=%d\n", gosnmp.ToBigInt(variable.Value))))
		}
	}
	resp <- buffer.String()
}

//Run prepares the sample data and sets the Goroutine parameters.
//TODO: Receive sample configuration parameters.
func (g *GetersonSnmp) Run() {
	resp := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i < 254; i++ {
		wg.Add(1)
		host := fmt.Sprintf("10.29.63.%d", i)
		oids := []string{
			".1.3.6.1.2.1.2.2.1.6.1",
			".1.3.6.1.2.1.2.2.1.6.2",
			".1.3.6.1.2.1.2.2.1.6.3",
			".1.3.6.1.2.1.1.1.0",
		}

		go g.Get(host, oids, "public", &wg, resp)
	}

	go func() {
		for response := range resp {
			fmt.Printf("%s", response)
		}

	}()

	wg.Wait()
	close(resp)
}
