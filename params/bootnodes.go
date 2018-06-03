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
var AuxiliumMainnetBootnodes = []string{
	"enode://e8f4eed1bb6cb5f20c04653d526ae4ad4a6e375b1cb8521c1a2678726c34db6d9d7f4c3151cd5a7a7cb9ca67b40aebc245a5a4e54367ae0742cbe1255a63e70e@149.28.36.115:30308",
	"enode://8a0ab923216f65e8499421c4a894c837abceb37cd61a432520b36a0280390a3620d62334e5afc9350c3d994d32634e28cc99715fdde53c0a9462a84c7cd969fb@207.246.122.81:30308",
}

// AuxiliumTestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Auxilium test network.
var AuxiliumTestnetBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
