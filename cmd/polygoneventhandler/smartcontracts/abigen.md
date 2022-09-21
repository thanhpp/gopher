# Setup

> https://goethereumbook.org/smart-contract-compile/
> https://grpc.io/docs/protoc-installation/

- Install solidity compiler
    ```
    sudo snap install solc --edge
    ```
- Install protoc
    ```
    apt install -y protobuf-compiler
    ```
- Install abigen tool
    ```
    go get -u github.com/ethereum/go-ethereum
    cd $GOPATH/src/github.com/ethereum/go-ethereum/
    make
    make devtools
    ```