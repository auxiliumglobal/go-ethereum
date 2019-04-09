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
	"enode://f3fd9556371841c44e496ec0b25a586f788a26bafc368eb6ab0a2aa34dc3013c30f7d1270ef4c136f5152374560654867f58074466c808bf782b82c3704d4061@149.28.45.212:30308",
	"enode://72a4c51e9c0c7f238a1d9ec00d97f82d8c7c6227d0dd2ac553ebd167e75c41cfa9dc0d3ee4a3dc8f03ea9450d3c699da9bdf0e91e24c5dd40eea505f49e86c87@138.68.176.12:30308",
	"enode://c03c389ec55517a3c49400a89706a744eedb4c3bab09ffedfd305bdbc9a3ffdd9febe75746fe88dbce06a55d02564af54d429c772543a1c720c03518ab2b86c8@8.12.18.68:30308",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
