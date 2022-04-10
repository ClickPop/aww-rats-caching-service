package hasura

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/env"
	GraphQL "github.com/clickpop/aww-rats-caching-service/schema/client"
	generated "github.com/clickpop/aww-rats-caching-service/schema/models"
	"github.com/clickpop/aww-rats-caching-service/tokens"
	"github.com/ethereum/go-ethereum/common"
)

type RatInput = generated.RatsInsertInput

type ClosetPieceInput = generated.ClosetPiecesInsertInput

type ClosetTokenInput = generated.ClosetTokensInsertInput

func UpsertRats(rats []tokens.RatTokenWithMetaAndId) {
	parsedRats := make([]*RatInput, 0)
	for _, rat := range rats {
		parsedRats = append(parsedRats, parseRat(rat))
	}
	if len(parsedRats) > 0 {
		graphQLClient := GetGraphQLClient()
		upsertedRats, err := graphQLClient.UpsertRats(context.Background(), parsedRats)
		if err != nil {
			log.Println("error upserting rats", err)
		}
		log.Printf("Upserted %d rats", len(upsertedRats.InsertRats.Returning))
	}
}

func UpsertClosetPieces(pieces []tokens.ClosetTokenWithMetaAndId) {
	parsedPieces := make([]*ClosetPieceInput, 0)
	for _, piece := range pieces {
		parsedPieces = append(parsedPieces, parseClosetPiece(piece))
	}
	if len(parsedPieces) > 0 {
		graphQLClient := GetGraphQLClient()
		upsertedPieces, err := graphQLClient.UpsertClosetPieces(context.Background(), parsedPieces)
		if err != nil {
			log.Println("error upserting pieces", err)
		}
		log.Printf("Upserted %d closet pieces", len(upsertedPieces.InsertClosetPieces.Returning))
	}
}

func UpsertClosetTokens(transfers []string) {
	tokens := handleClosetTransfers(transfers)
	if len(tokens) > 0 {
		graphQLClient := GetGraphQLClient()
		upsertedTokens, err := graphQLClient.UpsertClosetTokens(context.Background(), tokens)
		if err != nil {
			log.Println("error upserting pieces", err)
		}
		log.Printf("Upserted %d closet tokens", len(upsertedTokens.InsertClosetTokens.Returning))
	}
}

func CallHasura(rats []tokens.RatTokenWithMetaAndId, pieces []tokens.ClosetTokenWithMetaAndId, transfers []string) {
	UpsertRats(rats)
	UpsertClosetPieces(pieces)
	UpsertClosetTokens(transfers)
}

func parseRat(rat tokens.RatTokenWithMetaAndId) *RatInput {
	owner := rat.Owner.Hex()
	baseState := 0
	parsedRat := RatInput{ID: rat.Id, Owner: &owner, Name: &rat.Name, Description: &rat.Description, Image: &rat.Image, Rattitude: &baseState, Cunning: &baseState, Cuteness: &baseState}
	for _, v := range rat.OpenseaMeta.Attributes {
		switch v.TraitType {
		case "Accessory":
			val := v.Value.(string)
			parsedRat.Accessory = &val
		case "Background":
			val := v.Value.(string)
			parsedRat.Background = &val
		case "Color":
			val := v.Value.(string)
			parsedRat.Color = &val
		case "Ears":
			val := v.Value.(string)
			parsedRat.Ears = &val
		case "Eyes":
			val := v.Value.(string)
			parsedRat.Eyes = &val
		case "Generation":
			val := v.Value.(string)
			parsedRat.Generation = &val
		case "Glasses":
			val := v.Value.(string)
			parsedRat.Glasses = &val
		case "Hand":
			val := v.Value.(string)
			parsedRat.Hand = &val
		case "Hat":
			val := v.Value.(string)
			parsedRat.Hat = &val
		case "Head":
			val := v.Value.(string)
			parsedRat.Head = &val
		case "Markings":
			val := v.Value.(string)
			parsedRat.Markings = &val
		case "Pet":
			val := v.Value.(string)
			parsedRat.Pet = &val
		case "Shirt":
			val := v.Value.(string)
			parsedRat.Shirt = &val
		case "Snout":
			val := v.Value.(string)
			parsedRat.Snout = &val
		case "Tail":
			val := v.Value.(string)
			parsedRat.Tail = &val
		case "Torso":
			val := v.Value.(string)
			parsedRat.Torso = &val
		case "Type":
			val := v.Value.(string)
			parsedRat.Type = &val
		case "Cunning":
			val := int(v.Value.(float64))
			parsedRat.Cunning = &val
		case "Rattitude":
			val := int(v.Value.(float64))
			parsedRat.Rattitude = &val
		case "Cuteness":
			val := int(v.Value.(float64))
			parsedRat.Cuteness = &val
		}
	}
	return &parsedRat
}

func parseClosetPiece(piece tokens.ClosetTokenWithMetaAndId) *ClosetPieceInput {
	revShareAddress := piece.RevShareAddress.Hex()
	revShareBytes, err := json.Marshal(piece.RevShareAmount)
	if err != nil {
		log.Fatal("cannot marshal closet piece rev share", err)
	}
	revShareParsed := string(revShareBytes)
	parsedPiece := ClosetPieceInput{ID: piece.Id, Name: &piece.OpenseaMeta.Name, Description: &piece.Description, Image: &piece.Image, Cost: piece.Cost, MaxTokens: piece.MaxTokens, MaxPerWallet: piece.MaxPerWallet, RevShareAddress: &revShareAddress, RevShareAmount: &revShareParsed, Active: &piece.Active}
	for _, v := range piece.OpenseaMeta.Attributes {
		switch v.TraitType {
		case "Collection":
			val := v.Value.(string)
			parsedPiece.Collection = &val
		case "Piece Type":
			val := v.Value.(string)
			parsedPiece.PieceType = &val
		case "Sponsor":
			val := v.Value.(string)
			parsedPiece.Sponsor = &val
		case "Sponsor URL":
			val := v.Value.(string)
			parsedPiece.SponsorURL = &val
		case "Cunning":
			val := int(v.Value.(float64))
			parsedPiece.Cunning = &val
		case "Rattitude":
			val := int(v.Value.(float64))
			parsedPiece.Rattitude = &val
		case "Cuteness":
			val := int(v.Value.(float64))
			parsedPiece.Cuteness = &val
		}
	}
	return &parsedPiece
}

func handleClosetTransfers(addresses []string) []*ClosetTokenInput {
	changes := make([]*ClosetTokenInput, 0)
	closet := blockchain.ClosetContract
	opts := blockchain.Opts
	ids, err := closet.GetAllTokenIds(opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addresses {
		address := addr
		if addr == common.BigToHash(big.NewInt(0)).Hex() {
			continue
		}
		log.Println("Handling token changes for address ", common.HexToAddress(addr))
		accountArr := make([]common.Address, len(ids))
		for i := 0; i < len(accountArr); i++ {
			accountArr[i] = common.HexToAddress(addr)
		}
		bals, err := closet.BalanceOfBatch(opts, accountArr, ids)
		if err != nil {
			log.Println("error: ", err)
		}
		for i, bal := range bals {
			token := ClosetTokenInput{TokenID: ids[i], Owner: &address, Amount: bal}
			changes = append(changes, &token)
		}
	}

	return changes
}

func GetGraphQLClient() *GraphQL.Client {
	return GraphQL.NewClient(http.DefaultClient, env.HASURA_ENDPOINT, addHeader)
}

func addHeader(req *http.Request) {
	req.Header.Add("X-Hasura-Admin-Secret", env.HASURA_ADMIN_SECRET)
}
