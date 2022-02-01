package historical

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/env"
	"github.com/clickpop/aww-rats-caching-service/hasura"
	"github.com/clickpop/aww-rats-caching-service/tokens"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func Query(shouldSync bool) {
	rats := loadRats()
	pieces := loadClosetPieces()
	transfers := loadClosetTokens()
	if shouldSync {
		hasura.CallHasura(rats, pieces, transfers)
	}
}

func addApiKey(req *http.Request) {
	req.Header.Add("authorization", "test")
}

func loadRats() []tokens.RatTokenWithMetaAndId {
	numRats, err := blockchain.RatContract.NumTokens(blockchain.Opts)
	if err != nil {
		log.Println(err)
	}
	var rats = map[int]tokens.RatTokenWithMeta{}
	ratTokensFile, err := os.OpenFile("./.cache/ratTokens.json", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll("./.cache", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			ratTokensFile, err = os.Create("./.cache/ratTokens.json")

		} else {
			log.Fatal(err)
		}
	}
	cacheFile, err := os.ReadFile("./.cache/ratTokens.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(cacheFile, &rats)
	for i := 100; i < int(numRats.Int64()+100); i++ {
		uri, err := blockchain.RatContract.TokenURI(blockchain.Opts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(i, err)
		}
		owner, err := blockchain.RatContract.OwnerOf(blockchain.Opts, big.NewInt(int64(i)))
		currToken := tokens.RatToken{Owner: owner, URI: uri}
		meta := rats[i].OpenseaMeta
		if rats[i].URI != "" && rats[i].URI != currToken.URI {
			meta = tokens.OpenseaMeta{}
		}
		rats[i] = tokens.RatTokenWithMeta{RatToken: currToken, OpenseaMeta: meta}
		ratJSON, err := json.MarshalIndent(rats, "", "  ")
		ratTokensFile.Truncate(0)
		ratTokensFile.WriteAt(ratJSON, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Loaded token:", i)
	}
	sum := 0
	for _, v := range rats {
		if v.Attributes == nil {
			sum++
		}
	}
	log.Println("Need meta for", sum, "tokens")
	for k, v := range rats {
		if v.Attributes != nil {
			continue
		}
		meta, err := tokens.GetRatMeta(v.URI)
		if err != nil {
			log.Fatal(err)
		}
		rats[k] = tokens.RatTokenWithMeta{RatToken: rats[k].RatToken, OpenseaMeta: meta}
		ratJSON, err := json.MarshalIndent(rats, "", "  ")
		ratTokensFile.Truncate(0)
		ratTokensFile.WriteAt(ratJSON, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Loaded meta for token:", k)
	}

	ratTokensFile.Close()
	ratsWithIds := make([]tokens.RatTokenWithMetaAndId, 0)
	for id, rat := range rats {
		ratsWithIds = append(ratsWithIds, tokens.RatTokenWithMetaAndId{Id: big.NewInt(int64(id)), RatTokenWithMeta: rat})
	}
	return ratsWithIds
}

func loadClosetPieces() []tokens.ClosetTokenWithMetaAndId {
	var closetTokenMap = map[string]tokens.ClosetTokenWithMeta{}
	closetTokensFile, err := os.OpenFile("./.cache/closetTokens.json", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll("./.cache", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			closetTokensFile, err = os.Create("./.cache/closetTokens.json")

		} else {
			log.Fatal(err)
		}
	}
	cacheFile, err := os.ReadFile("./.cache/closetTokens.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(cacheFile, &closetTokenMap)
	if len(closetTokenMap) == 0 {
		closetTokens, err := blockchain.ClosetContract.GetAllTokens(blockchain.Opts)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range closetTokens {
			closetTokenMap[v.Id.String()] = tokens.ClosetTokenWithMeta{Token: v.Token}
		}
	}
	for k, v := range closetTokenMap {
		if v.Attributes != nil {
			continue
		}
		token := v.Token
		id, err := strconv.Atoi(k)
		meta, err := tokens.GetClosetMeta(big.NewInt(int64(id)))
		closetTokenMap[k] = tokens.ClosetTokenWithMeta{Token: token, OpenseaMeta: meta}
		closetJSON, err := json.MarshalIndent(closetTokenMap, "", "  ")
		closetTokensFile.Truncate(0)
		closetTokensFile.WriteAt(closetJSON, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Loaded meta for token:", k)
	}

	closetTokensFile.Close()
	piecesWithIds := make([]tokens.ClosetTokenWithMetaAndId, 0)
	for id, piece := range closetTokenMap {
		newId, _ := strconv.Atoi(id)
		piecesWithIds = append(piecesWithIds, tokens.ClosetTokenWithMetaAndId{Id: big.NewInt(int64(newId)), ClosetTokenWithMeta: piece})
	}
	return piecesWithIds
}

func loadClosetTokens() []tokens.ClosetTransfer {
	lastBlock, err := blockchain.EthClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	bigLastBlock := big.NewInt(int64(lastBlock))
	var closetTransactions tokens.ClosetTransfers
	closetTransactionsFile, err := os.OpenFile("./.cache/closetTransactions.json", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll("./.cache", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			closetTransactionsFile, err = os.Create("./.cache/closetTransactions.json")

		} else {
			log.Fatal(err)
		}
	}
	cacheFile, err := os.ReadFile("./.cache/closetTransactions.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(cacheFile, &closetTransactions)
	if closetTransactions.Block == nil {
		closetTransactions.Block = big.NewInt(22075111)
	}
	if closetTransactions.Transfers == nil {
		closetTransactions.Transfers = make(map[common.Hash]tokens.ClosetTransfer)
	}

	for i := closetTransactions.Block; i.Cmp(bigLastBlock) < 0; i.Add(i, big.NewInt(100)) {
		endBlock := big.NewInt(0).Add(i, big.NewInt(99))
		if endBlock.Cmp(bigLastBlock) > 0 {
			endBlock = bigLastBlock
		}
		log.Printf("Block %s-%s\n", i, endBlock)

		closetTransactions.Block = i
		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		query := ethereum.FilterQuery{
			FromBlock: i,
			ToBlock:   endBlock,
			Addresses: []common.Address{common.HexToAddress(env.CLOSET_ADDR)},
		}

		logs, _ := blockchain.EthClient.FilterLogs(context.Background(), query)

		for _, l := range logs {
			switch l.Address.Hex() {
			case env.CLOSET_ADDR:
				switch l.Topics[0].Hex() {
				case blockchain.ClosetABI.Events["TransferBatch"].ID.Hex():
					event := struct {
						Ids    []*big.Int
						Values []*big.Int
					}{}
					err = blockchain.ClosetABI.UnpackIntoInterface(&event, "TransferBatch", l.Data)
					if err != nil {
						log.Println("Error:", err)
					}
					from := l.Topics[2]
					to := l.Topics[3]
					for i := 0; i < len(event.Ids); i++ {
						id := event.Ids[i]
						amount := event.Values[i]
						transfer := tokens.ClosetTransfer{From: from, To: to, Id: id, Amount: amount}
						closetTransactions.Transfers[l.TxHash] = transfer
					}
					log.Println("Loaded batch transfer transaction", l.TxHash.Hex())
				case blockchain.ClosetABI.Events["TransferSingle"].ID.Hex():
					event := struct {
						Id    *big.Int
						Value *big.Int
					}{}
					err = blockchain.ClosetABI.UnpackIntoInterface(&event, "TransferSingle", l.Data)
					if err != nil {
						log.Println("Error:", err)
					}
					from := l.Topics[2]
					to := l.Topics[3]
					id := event.Id
					amount := event.Value
					transfer := tokens.ClosetTransfer{From: from, To: to, Id: id, Amount: amount}
					closetTransactions.Transfers[l.TxHash] = transfer
					log.Println("Loaded single transfer transaction", l.TxHash.Hex())
				}
			}
		}
		data, err := json.MarshalIndent(closetTransactions, "", "  ")
		if err != nil {
			log.Println(err)
			continue
		}
		closetTransactionsFile.Truncate(0)
		closetTransactionsFile.WriteAt(data, 0)
	}
	transferArr := make([]tokens.ClosetTransfer, 0)
	for _, v := range closetTransactions.Transfers {
		transferArr = append(transferArr, v)
	}
	return transferArr
}
