# astrixminer

Astrixminer is a CPU-based miner for astrixd

## Requirements

Go 1.19 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install astrixd including all dependencies:

```bash
$ git clone https://github.com/astrix-network/astrixd
$ cd astrixd/cmd/astrixminer
$ go install .
```

- Kapaminer should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.
  
## Usage

The full astrixminer configuration options can be seen with:

```bash
$ astrixminer --help
```

But the minimum configuration needed to run it is:
```bash
$ astrixminer --miningaddr=<YOUR_MINING_ADDRESS>
```