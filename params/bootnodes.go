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
	"enode://de6bb59fb71c2f1c382cd9fa565b306dc914fe8870471377b64aee4741f9aa04cf762e1e21271fa1fca5acd54b2a46d867551c10c5cf966edb9f7ae1ad8b430d@8.12.18.68:30308",
	"enode://88c44959468a2479704be090e22cf7d47e536c4a1887853825c8cb4104c573892a26d94c3be33f4ee95185bf126a5b44f0f21116a8b2a6282587534c0161498e@149.28.45.212:30308",
	"enode://79bcbbd5c1d0ca6e2ab7f385b7e8dab8b1188438dcb89ac6b09ed11797460367d09d50362bc56e9330e2e1f005e4374e36bfecd8a62b53192d5e5ec8212e2f3c@138.68.176.12:30308",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
