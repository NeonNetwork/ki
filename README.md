# Ki - Native Dynamic Dashboard

## What is this?

This is a natively compiled, hardware accelerated dashboard built in [Go](https://go.dev/) with [Raylib](https://github.com/raysan5/raylib). Live data is pulled to the Ki Client through the Ki Server which essentially acts as a data pipeline, see [server.go](https://github.com/NeonNetwork/ki/blob/main/pkg/ki/server.go) for example.

## Platform Support

| Windows | Linux | MacOS | Android | iOS | Web |
| :-----: | :-: | :---: | :-: | :---: | :-----: |
|   ✅   | ✅  |  ✅   | ❓  |  ❓   |  ❌  |

I haven't got the time to test Android and iOS. It seems to be possible to run Go on mobile, and WebAssembly is officially supported. Raylib uses Emcripten for WebAssembly target which is not supported by Go's cgo compiler, but raylib-go has an example for android [here](https://github.com/gen2brain/raylib-go/tree/master/examples/others/android/example)

## Installation

Download the repository

```sh
git clone https://github.com/neonnetwork/ki.git
cd ki
```

Install the dependencies

```sh
go get -u ./cmd/*
go mod tidy
```

Start the Ki Server

```sh
go run ./cmd/server
```

Start the Ki Client

```sh
export KI_SETUP=BINANCE # choose data setup
go run ./cmd/ki
```

# Credits

* https://github.com/raysan5/raylib Epic project for cross-platform 2D and 3D graphics
* https://github.com/gen2brain/raylib-go Go bindings for Raylib
