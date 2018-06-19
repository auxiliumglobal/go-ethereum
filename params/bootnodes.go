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

// AuxiliumMainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Auxilium network.
var AuxiliumTestnetBootnodes = []string{
	"enode://e8f4eed1bb6cb5f20c04653d526ae4ad4a6e375b1cb8521c1a2678726c34db6d9d7f4c3151cd5a7a7cb9ca67b40aebc245a5a4e54367ae0742cbe1255a63e70e@149.28.36.115:30308",
}

// AuxiliumTestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Auxilium test network.
var AuxiliumMainnetBootnodes = []string{
	"enode://d47d0f9bd08495d5dbb93d8f431027d662c4173c8bb64398299f0e990724fd6a1350090fb19e55e86f921ce4f01943b84a70dd9b8579d354621c6f4691da9b57@138.68.176.12:30308",
	"enode://b20907893af614595ed72144ac2adb8933e4e6136cdececf36b969df4de991566d4e7c7b6d489c1dd573fc669028d0aafd7a9d9ad33d1031e8f5ee2f32b6ba97@8.12.22.12:30308",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
