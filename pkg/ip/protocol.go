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

import "net"

// Protocol to differentiate between IPv4 and IPv6
type Protocol byte

const (
	ProtocolIPv4 Protocol = iota
	ProtocolIPv6
)

const (
	FamilyIPV4 = "ipv4"
	FamilyIPV6 = "ipv6"
)

func ProtocolByIPNet(ipn net.IPNet) Protocol {
	if ipn.IP.To4() != nil {
		return ProtocolIPv4
	}
	return ProtocolIPv6
}

func ProtocolByIP(ip net.IP) Protocol {
	if ip.To4() != nil {
		return ProtocolIPv4
	}
	return ProtocolIPv6
}
