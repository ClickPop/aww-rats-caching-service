package blockchain

import (
	"log"
	"math/big"
	"strings"

	"github.com/clickpop/aww-rats-caching-service/closet"
	"github.com/clickpop/aww-rats-caching-service/env"
	"github.com/clickpop/aww-rats-caching-service/rat"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client
var RatABI abi.ABI
var ClosetABI abi.ABI
var RatContract *rat.Rat
var ClosetContract *closet.Closet 
var Opts *bind.CallOpts
var ClosetTokenURI string

func Init() {
	var err error
	
	EthClient, err = ethclient.Dial(env.POLYGON_ENDPOINT)
	if err != nil {
		log.Fatal(err)
	}

	RatABI, err = abi.JSON(strings.NewReader(string(rat.RatABI)))
	if err != nil {
		log.Fatal(err)
	}

	ClosetABI, err = abi.JSON(strings.NewReader(string(closet.ClosetABI)))
	if err != nil {
		log.Fatal(err)
	}

	RatContract, err = rat.NewRat(common.HexToAddress(env.RAT_ADDR), EthClient)
	ClosetContract, err = closet.NewCloset(common.HexToAddress(env.CLOSET_ADDR), EthClient)

	Opts = &bind.CallOpts{From: common.HexToAddress("0x38882F37879aE3d53100F90651Df448960dB4975")}

	ClosetTokenURI, err = ClosetContract.Uri(Opts, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
}