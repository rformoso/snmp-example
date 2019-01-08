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
	"time"

	"github.com/rformoso/snmp-example/util"
	"github.com/soniah/gosnmp"
)

type BulkWalk interface {
	Run()
}

type BulkWalkSnmp struct {
}

//BulkWalk is BulkWalk
func (b *BulkWalkSnmp) BulkWalk(target string, oid string, community string) {

	params := &gosnmp.GoSNMP{
		Target:    target,
		Port:      uint16(161),
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   1,
		MaxOids:   gosnmp.MaxOids,
	}

	err := params.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
	}
	defer params.Conn.Close()

	err = params.BulkWalk(oid, func(pdu gosnmp.SnmpPDU) error {
		var buffer bytes.Buffer

		switch pdu.Type {
		case gosnmp.OctetString:
			mac := util.ValidateMAC(pdu.Value.([]byte))
			if mac != "" {
				buffer.WriteString(fmt.Sprintf("host=%s oid=%s mac=%s\n", target, pdu.Name, mac))
				break
			}
			buffer.WriteString(fmt.Sprintf("host=%s oid=%s string=%s\n", target, pdu.Name, string(pdu.Value.([]byte))))

		case gosnmp.Integer:
			buffer.WriteString(fmt.Sprintf("host=%s oid=%s number=%d\n", target, pdu.Name, gosnmp.ToBigInt(pdu.Value)))
		}
		fmt.Printf(buffer.String())

		return nil
	})
	if err != nil {
		fmt.Printf("Host=%s Walk Error: %v\n", target, err)
	}
}

//Run prepares the sample data and sets the Goroutine parameters.
//TODO: Receive sample configuration parameters.
func (b *BulkWalkSnmp) Run() {
	host := "10.29.63.30"
	oid := ".1.3.6.1.2.1.25.3.5.1"
	community := "public"

	b.BulkWalk(host, oid, community)
}
