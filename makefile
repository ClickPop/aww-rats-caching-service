installSolc = ""
installProtoc = ""
installGeth = ""

ifeq ($(shell go env GOOS),darwin)
	installSolc = brew update && brew tap ethereum/ethereum && brew install solidity
	installProtoc = brew install protobuf
	installGeth = brew install ethereum
endif

ifeq ($(shell go env GOOS),linux)
	installSolc = sudo add-apt-repository -y ppa:ethereum/ethereum && sudo apt update && sudo apt install -y solc
	installProtoc = sudo apt install -y protobuf-compiler
	installGeth = sudo apt-get install ethereum
endif


all: InstallSolc InstallProtoc GraphQLDeps GraphQLCodegen InstallDeps BuildABIGEN GetContract Compile BuildABI BuildBinary

contract: GetContract Compile BuildABI

build: BuildBinary

docker: BuildBinary DockerBuild

graphql: GraphQLDeps GraphQLCodegen

InstallDeps: 
	go get && pnpm install

InstallProtoc:
	$(installProtoc)

InstallSolc:
	$(installSolc)	

BuildABIGEN: 
	$(installGeth)

GetContract:
	curl -L "https://raw.githubusercontent.com/ClickPop/aww-rats-client/main/smart-contracts/src/contracts/Rat.sol" -o Rat.sol \
	&& curl -L "https://raw.githubusercontent.com/ClickPop/aww-rats-client/main/smart-contracts/src/contracts/Closet.sol" -o Closet.sol

Compile: 
	solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ --optimize --abi --bin ./Rat.sol -o build --overwrite \
	&& solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ --optimize --abi --bin ./Closet.sol -o build --overwrite

BuildABI: 
	abigen --abi=./build/Rat.abi --bin=./build/Rat.bin --pkg=rat --out=./rat/rat.go \
	&& abigen --abi=./build/Closet.abi --bin=./build/Closet.bin --pkg=closet --out=./closet/closet.go

BuildBinary:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o exec/caching-service -a

DockerBuild:
	docker build -t aww-rats-caching-service . --platform linux/amd64

GraphQLDeps:
	go get -u -d github.com/Yamashou/gqlgenc@v0.0.2 \
	&& go install github.com/Yamashou/gqlgenc@v0.0.2

GraphQLCodegen:
	gqlgenc