package watcher

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/closet"
	"github.com/clickpop/aww-rats-caching-service/env"
	"github.com/clickpop/aww-rats-caching-service/hasura"
	"github.com/clickpop/aww-rats-caching-service/tokens"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func Watch() {
	client := blockchain.EthClient
	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening for new blocks")

	for {
		select {
		case err := <-sub.Err():
			if strings.Contains(err.Error(), "429") {
				reg := regexp.MustCompile(`(?m)\d+(?: seconds)`)
				seconds, err := strconv.Atoi(reg.FindString(err.Error()))
				if err != nil {
					log.Fatal("no timeout found", err)
				}
				time.Sleep(time.Duration(seconds))
				Watch()
			}
			log.Fatal("subscription error: ", err)
		case header := <-headers:
			hash := header.Hash()
			block, err := client.BlockByHash(context.Background(), hash)
			if err != nil {
				log.Println("block error: ", err)
				continue
			}
			query := ethereum.FilterQuery{
				BlockHash: &hash,
				Addresses: []common.Address{common.HexToAddress(env.RAT_ADDR), common.HexToAddress(env.CLOSET_ADDR)},
			}

			logs, _ := client.FilterLogs(context.Background(), query)

			closetPieces := make([]tokens.ClosetTokenWithMetaAndId, 0)
			closetTokens := make([]tokens.ClosetTransfer, 0)
			rats := make([]tokens.RatTokenWithMetaAndId, 0)

			for _, l := range logs {
				switch l.Address.Hex() {
				case env.RAT_ADDR:
					switch l.Topics[0].Hex() {
					case blockchain.RatABI.Events["Transfer"].ID.Hex():
						from := l.Topics[1]
						to := l.Topics[2]
						id := l.Topics[3].Big()
						rats = append(rats, tokens.RatTokenWithMetaAndId{Id: id, RatTokenWithMeta: tokens.RatTokenWithMeta{RatToken: tokens.RatToken{Owner: common.BytesToAddress(to.Bytes())}}})
						log.Println("Rat transfer", from, to, id)
					}
				case env.CLOSET_ADDR:
					switch l.Topics[0].Hex() {
					case blockchain.ClosetABI.Events["TokenTypeAdded"].ID.Hex():
						fallthrough
					case blockchain.ClosetABI.Events["TokenTypeChanged"].ID.Hex():
						event := struct {
							TokenId *big.Int
							Token   closet.Token
						}{}
						err = blockchain.ClosetABI.UnpackIntoInterface(&event, "TokenTypeChanged", l.Data)
						if err != nil {
							log.Println("closet TokenTypeAdd/TokenTypeChange unpack error:", err)
						}
						meta, err := tokens.GetClosetMeta(event.TokenId)
						if err != nil {
							log.Println(err)
						}
						token := tokens.ClosetTokenWithMetaAndId{Id: event.TokenId, ClosetTokenWithMeta: tokens.ClosetTokenWithMeta{Token: event.Token, OpenseaMeta: meta}}
						closetPieces = append(closetPieces, token)
						log.Println("Closet piece added/changed", event)
					case blockchain.ClosetABI.Events["TransferBatch"].ID.Hex():
						event := struct {
							Ids    []*big.Int
							Values []*big.Int
						}{}
						err = blockchain.ClosetABI.UnpackIntoInterface(&event, "TransferBatch", l.Data)
						if err != nil {
							log.Println("closet TransferBatch unpack error:", err)
						}
						from := l.Topics[2]
						to := l.Topics[3]
						for i := 0; i < len(event.Ids); i++ {
							id := event.Ids[i]
							amount := event.Values[i]
							transfer := tokens.ClosetTransfer{From: from, To: to, Id: id, Amount: amount}
							closetTokens = append(closetTokens, transfer)
							log.Println("Closet transfer", from, to, amount, event)
						}
					case blockchain.ClosetABI.Events["TransferSingle"].ID.Hex():
						event := struct {
							Id    *big.Int
							Value *big.Int
						}{}
						err = blockchain.ClosetABI.UnpackIntoInterface(&event, "TransferSingle", l.Data)
						if err != nil {
							log.Println("closet TransferSingle unpack error:", err)
						}
						from := l.Topics[2]
						to := l.Topics[3]
						id := event.Id
						amount := event.Value
						transfer := tokens.ClosetTransfer{From: from, To: to, Id: id, Amount: amount}
						closetTokens = append(closetTokens, transfer)
						log.Println("Closet transfer", from, to, amount, event)
					}
				}
			}

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}
				switch *tx.To() {
				case common.HexToAddress(env.RAT_ADDR):
					storeAsset := blockchain.RatABI.Methods["storeAsset"]
					sig := common.Bytes2Hex(tx.Data())[:8]
					if hex.EncodeToString(storeAsset.ID) == sig {
						inputData, err := hex.DecodeString(common.Bytes2Hex(tx.Data())[8:])
						data, err := blockchain.RatABI.Methods["storeAsset"].Inputs.NonIndexed().Unpack(inputData)
						if err != nil {
							log.Println("rat storeAsset unpack error:", err)
						} else {
							id := data[0].(*big.Int)
							uri := data[1].(string)
							owner, err := blockchain.RatContract.OwnerOf(&bind.CallOpts{Context: context.Background()}, id)
							if err != nil {
								log.Println("rat contract error querying owner:", err)
								continue
							}
							meta, err := tokens.GetRatMeta(uri)
							if err != nil {
								log.Println("rat error loading meta:", err)
								continue
							}
							token := tokens.RatTokenWithMetaAndId{Id: id, RatTokenWithMeta: tokens.RatTokenWithMeta{RatToken: tokens.RatToken{Owner: owner, URI: uri}, OpenseaMeta: meta}}
							log.Println("Set Rat URI", id, owner, uri, meta)
							rats = append(rats, token)
						}
					}
				}
			}

			if len(rats) > 0 || len(closetPieces) > 0 || len(closetTokens) > 0 {
				hasura.CallHasura(rats, closetPieces, closetTokens)
			}
		}
	}
}
