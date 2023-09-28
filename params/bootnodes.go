// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
    "enode://f938cc5bb433d37c68078980811eb521786f520063be5d1ce4684824c2241e5bdd666553355d94743641ee96faee2b540dfa94138c5f0c07b66add3f3fdc9710@146.59.126.215:44303",
    "enode://171651b8785a0188af1df7064c1586d3bcd05195b9cf73e3ec41362d2216ac299a356df204b57c2579c68b130f4187fb4c06fc5538187e565bd55cba61dd13b6@51.83.185.32:44303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
    "",
}

var V5Bootnodes = []string{
}

const dnsPrefix = "enrtree://AKNLJUOC5L5NOXNFIANWHEJEOL4MPBKYJ7T2WWQDB3HEIKTNP6XPS@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case TestnetGenesisHash:
		net = "testnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + ".disco." + net + ".kopalniekrypto.pl"
}
