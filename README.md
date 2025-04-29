this is a project to learn how to use go-ethereum to develop a dapp
## 1. 安装 go-ethereum
```
go get -u github.com/ethereum/go-ethereum
```
## 2. 安装 ganache
```
npm install -g ganache-cli
```
## 3. 启动 ganache
```
ganache-cli
```
## 4. 编译合约
```
solc --abi --bin ./contracts/HelloWorld.sol
```
## 5. 部署合约
```
go run ./cmd/deploy/main.go
```
## 6. 调用合约
```
go run ./cmd/call/main.go
``` 