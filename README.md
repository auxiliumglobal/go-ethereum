## Go Ethereum for Auxilium cryptocurrency

Official Go implementation of the Ethereum protocol, [geth](https://github.com/ethereum/go-ethereum), altered for Auxilium cryptocurrency.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/ethereum/go-ethereum)
[![Go Report Card](https://goreportcard.com/report/github.com/auxiliumglobal/go-ethereum)](https://goreportcard.com/report/github.com/auxiliumglobal/go-ethereum)

Binary archives are published at https://downloads.auxilium.global/node/.

## Auxilium and Auxilium Global

Auxilium is a cryptocurrency for Auxilium Global organisation which is aimed at conservation, renewable resources, helping others and other forms of philantropy. Given these goals, Auxilium takes "green" PoA consensus as a basis pairing it with [Auxilium Interest Distribution platform](https://auxilium.global/code) as a substitue for a reward mechanism.

For more details on Auxilium and Auxilium Global, visit https://auxilium.global.

## Build and run the node

Install Go (version 1.7 or later) and a C compiler. Then move to directory with this repository and run:

    make geth

To start the node use the following command:

    $ ./build/bin/geth --syncmode full

Upon first start node will begin synchronizing from block 1. Due to ChainID fork, initial synchonization with Auxilium network requires some additional steps: once block 2783760 has been downloaded, node should be restarted with the same command

    $ ./build/bin/geth --syncmode full

After that, the node can be used as usual. Number of the last downloaded block can be seen in the logs of `geth` command, via commands passed to `geth --syncmode full console` or with the help of local HTTP-RPC server.

All subsequent starts of Auxilium `geth` do not require any additional restarts but work out of the box.

Please refer to `geth.README.md` for detailed instructions, as well as for description of command line options.

**Note:** for compatibility's sake code of Auxilium node continues to use `"github.com/ethereum/go-ethereum"` import statements. If you are making separate builds for original Ethereum node you might run into mix-up with package imports. In that case, it is advised to use symbolic links, e.g.:

     ln -s -t $GOPATH/src/github.com/ethereum/ $GOPATH/src/github.com/auxiliumglobal/go-ethereum

Here path to Auxilium node code can be arbitrary.

## Auxilium's modifications of the node

Auxilium's node was modified in order to connect only to Auxilium blockchain network configured with PoA Clique consensus, with different Chain ID. Initially Auxilium used Chain ID 8. However, to avoid collisions with other Ethereum-based networks, Auxilium network underwent ChainID fork that took place on June 1st, 2019, block 2783760. Now Auxilium network has Chain ID 28945486.

The `syncmode` option and `console` command operate identically to those in original `geth`; however, Auxilium's node uses port 30308 by default. Likewise, default data directory for Auxilium's node is `~/.auxilium` on Linux and its analogues on Windows and OS X.

## License

Auxilium's node is following licensing terms of Go Ethereum node, namely [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html) and [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). Refer to `geth.README.md` for details.
