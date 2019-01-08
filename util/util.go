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

package util

import (
	"fmt"
	"net"
)

// ValidateMAC checks if the []byte passed is a Mac Address and if so, checks to see if it is valid.
// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
func ValidateMAC(bytes []byte) string {
	var hw net.HardwareAddr
	hw = bytes
	if len(hw.String()) != 17 {
		return ""
	}
	hw, err := net.ParseMAC(hw.String())
	if err != nil {
		return ""
	}

	return hw.String()
}

//FormatLog returns the message in the pattern that the developer wants.
func FormatLog(host string, oid string, message string) string {
	return fmt.Sprintf("host=%s oid=%s %s", host, oid, message)
}
