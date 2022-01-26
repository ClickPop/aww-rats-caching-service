installSolc = ""
installProtoc = ""
installGeth = ""

ifeq ($(shell go env GOOS),darwin)
	installSolc = brew update && brew tap ethereum/ethereum && brew install solidity
	installProtoc = brew install protobuf
	installGeth = brew install ethereum
endif

ifeq ($(shell go env GOOS),linux)
	installSolc = add-apt-repository -y ppa:ethereum/ethereum && apt update && apt install -y solc
	installProtoc = apt install -y protobuf-compiler
	installGeth = apt-get install ethereum
endif


all: InstallSolc InstallProtoc InstallDeps BuildABIGEN GetContract Compile BuildABI

contract: GetContract Compile BuildABI

build: BuildBinary

docker: BuildBinary DockerBuild

InstallDeps: 
	go get && npm install

InstallProtoc:
	$(installProtoc)

InstallSolc:
	$(installSolc)	

BuildABIGEN: 
	$(installGeth)

GetContract:
	curl "https://$(shell echo $(ACCESS_TOKEN)@$(CONTRACT_URL))/Rat.sol" -o Rat.sol \
	&& curl "https://$(shell echo $(ACCESS_TOKEN)@$(CONTRACT_URL))/Closet.sol" -o Closet.sol

Compile: 
	solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ --optimize --abi --bin ./Rat.sol -o build --overwrite \
	&& solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ --optimize --abi --bin ./Closet.sol -o build --overwrite

BuildABI: 
	abigen --abi=./build/Rat.abi --bin=./build/Rat.bin --pkg=rat --out=./rat/rat.go \
	&& abigen --abi=./build/Closet.abi --bin=./build/Closet.bin --pkg=closet --out=./closet/closet.go

BuildBinary:
	GOOS=linux \
	GOARCH=amd64 \
	go build -o exec/caching-service

DockerBuild:
	docker build -t aww-rats-caching-service . --platform linux/amd64