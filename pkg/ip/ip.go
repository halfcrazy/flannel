// Copyright 2015 flannel authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ip

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
)

// NextIP get next ip
// copy from https://stackoverflow.com/a/49847634
func NextIP(ip net.IP) net.IP {
	// Convert to big.Int and increment
	ipb := new(big.Int).SetBytes([]byte(ip))
	ipb.Add(ipb, big.NewInt(1))

	// Add leading zeros
	b := ipb.Bytes()
	b = append(make([]byte, len(ip)-len(b)), b...)
	return net.IP(b)
}

func NextNIP(ip net.IP, n *big.Int) net.IP {
	// Convert to big.Int and increment
	ipb := new(big.Int).SetBytes([]byte(ip))
	ipb.Add(ipb, n)

	// Add leading zeros
	b := ipb.Bytes()
	b = append(make([]byte, len(ip)-len(b)), b...)
	return net.IP(b)
}

func PreviousNIP(ip net.IP, n *big.Int) net.IP {
	// Convert to big.Int and increment
	ipb := new(big.Int).SetBytes([]byte(ip))
	ipb.Sub(ipb, n)

	// Add leading zeros
	b := ipb.Bytes()
	b = append(make([]byte, len(ip)-len(b)), b...)
	return net.IP(b)
}

func PrefixLenByIPNet(ipn net.IPNet) uint {
	ones, _ := ipn.Mask.Size()
	return uint(ones)
}

func NetEqual(n1, n2 net.IPNet) bool {
	return n1.String() == n2.String()
}

func IPExpand(ip net.IP) string {
	if ip.To4() != nil {
		return ip.String()
	}
	dst := make([]byte, hex.EncodedLen(len(ip)))
	_ = hex.Encode(dst, ip)
	return string(dst[0:4]) + ":" +
		string(dst[4:8]) + ":" +
		string(dst[8:12]) + ":" +
		string(dst[12:16]) + ":" +
		string(dst[16:20]) + ":" +
		string(dst[20:24]) + ":" +
		string(dst[24:28]) + ":" +
		string(dst[28:])
}

func IPNetExpand(ipn net.IPNet) string {
	prefixLen, _ := ipn.Mask.Size()
	return fmt.Sprintf("%s/%d", IPExpand(ipn.IP), prefixLen)
}

func NetEmpty(ipn net.IPNet) bool {
	ones, _ := ipn.Mask.Size()
	return (ipn.IP == nil) && ones == 0
}
