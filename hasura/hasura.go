package hasura

import (
	"bytes"
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/env"
	"github.com/clickpop/aww-rats-caching-service/tokens"
	"github.com/ethereum/go-ethereum/common"
)

type RatInput struct {
		Id          *big.Int `json:"id,omitempty"`
		Owner       common.Address   `json:"owner,omitempty"`
		Name        string   `json:"name,omitempty"`
		Image       string   `json:"image,omitempty"`
		Description string   `json:"description,omitempty"`
		Accessory   string   `json:"accessory,omitempty"`
		Background  string   `json:"background,omitempty"`
		Color       string   `json:"color,omitempty"`
		Ears        string   `json:"ears,omitempty"`
		Eyes        string   `json:"eyes,omitempty"`
		Generation  string   `json:"generation,omitempty"`
		Glasses     string   `json:"glasses,omitempty"`
		Hand        string   `json:"hand,omitempty"`
		Hat         string   `json:"hat,omitempty"`
		Head        string   `json:"head,omitempty"`
		Markings    string   `json:"markings,omitempty"`
		Pet         string   `json:"pet,omitempty"`
		Shirt       string   `json:"shirt,omitempty"`
		Snout       string   `json:"snout,omitempty"`
		Tail        string   `json:"tail,omitempty"`
		Torso       string   `json:"torso,omitempty"`
		Cunning     float64      `json:"cunning,omitempty"`
		Cuteness    float64      `json:"cuteness,omitempty"`
		Rattitude   float64      `json:"rattitude,omitempty"`
}

type ClosetPieceInput struct {
		Id              *big.Int `json:"id,omitempty"`
		Name            string   `json:"name,omitempty"`
		Image           string   `json:"image,omitempty"`
		Description     string   `json:"description,omitempty"`
		Cost            *big.Int `json:"cost,omitempty"`
		MaxTokens       *big.Int `json:"max_tokens,omitempty"`
		MaxPerWallet    *big.Int `json:"max_per_wallet,omitempty"`
		Active          bool     `json:"active,omitempty"`
		RevShareAddress string   `json:"rev_share_address,omitempty"`
		RevShareAmount  [2]*big.Int   `json:"rev_share_amount,omitempty"`
		PieceType       string   `json:"piece_type,omitempty"`
		Collection      string   `json:"collection,omitempty"`
		Sponsor         string   `json:"sponsor,omitempty"`
		SponsorURL      string   `json:"sponsor_url,omitempty"`
		Cunning         float64      `json:"cunning,omitempty"`
		Cuteness        float64      `json:"cuteness,omitempty"`
		Rattitude       float64      `json:"rattitude,omitempty"`
}

type ClosetTokenInput struct {
		TokenId *big.Int `json:"token_id,omitempty"`
		Owner   common.Address   `json:"owner,omitempty"`
		Amount  *big.Int `json:"amount,omitempty"`
}

type Body struct {
	Rats         []RatInput         `json:"rats"`
	ClosetPieces []ClosetPieceInput `json:"closet_pieces"`
	ClosetTokens []ClosetTokenInput `json:"closet_tokens"`
}

func CallHasura(rats []tokens.RatTokenWithMetaAndId, pieces []tokens.ClosetTokenWithMetaAndId, transfers []tokens.ClosetTransfer) {
	parsedRats := make([]RatInput, 0)
	for _, rat := range rats {
		parsedRats = append(parsedRats, parseRat(rat))
	}
	parsedPieces := make([]ClosetPieceInput, 0)
	for _, piece := range pieces {
		parsedPieces = append(parsedPieces, parseClosetPiece(piece))
	}

	tokens := handleClosetTransfers(transfers)
	body := Body{Rats: parsedRats, ClosetPieces: parsedPieces, ClosetTokens: tokens}
	bodyData, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}
	bodyReader := bytes.NewReader(bodyData)
	req, err := http.NewRequest(http.MethodPost, env.ENDPOINT_URL, bodyReader)
	if err != nil {
		log.Println("error: ", err)
	}
	req.Header.Add("authorization", env.HASURA_API_KEY)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	res.Body.Close()
}

func parseRat(rat tokens.RatTokenWithMetaAndId) RatInput {
	parsedRat := RatInput{Id: rat.Id, Owner: rat.Owner, Name: rat.Name, Description: rat.Description, Image: rat.Image}
	for _, v := range rat.OpenseaMeta.Attributes {
		switch v.TraitType {
		case "Accessory":
			parsedRat.Accessory = v.Value.(string)
		case "Background":
			parsedRat.Background = v.Value.(string)
		case "Color":
			parsedRat.Color = v.Value.(string)
		case "Ears":
			parsedRat.Ears = v.Value.(string)
		case "Eyes":
			parsedRat.Eyes = v.Value.(string)
		case "Generation":
			parsedRat.Generation = v.Value.(string)
		case "Glasses":
			parsedRat.Glasses = v.Value.(string)
		case "Hand":
			parsedRat.Hand = v.Value.(string)
		case "Hat":
			parsedRat.Hat = v.Value.(string)
		case "Head":
			parsedRat.Head = v.Value.(string)
		case "Markings":
			parsedRat.Markings = v.Value.(string)
		case "Pet":
			parsedRat.Pet = v.Value.(string)
		case "Shirt":
			parsedRat.Shirt = v.Value.(string)
		case "Snout":
			parsedRat.Snout = v.Value.(string)
		case "Tail":
			parsedRat.Tail = v.Value.(string)
		case "Torso":
			parsedRat.Torso = v.Value.(string)
		case "Cunning":
			parsedRat.Cunning = v.Value.(float64)
		case "Rattitude":
			parsedRat.Rattitude = v.Value.(float64)
		case "Cuteness":
			parsedRat.Cuteness = v.Value.(float64)
		}
	}
	return parsedRat
}

func parseClosetPiece(piece tokens.ClosetTokenWithMetaAndId) ClosetPieceInput {
	parsedPiece := ClosetPieceInput{Id: piece.Id, Name: piece.OpenseaMeta.Name, Description: piece.Description, Image: piece.Image, Cost: piece.Cost, MaxTokens: piece.MaxTokens, MaxPerWallet: piece.MaxPerWallet, RevShareAddress: piece.RevShareAddress.Hex(), RevShareAmount: piece.RevShareAmount, Active: piece.Active}
	for _, v := range piece.OpenseaMeta.Attributes {
		switch v.TraitType {
		case "Collection":
			parsedPiece.Collection = v.Value.(string)
		case "Piece Type":
			parsedPiece.PieceType = v.Value.(string)
		case "Sponsor":
			parsedPiece.Sponsor = v.Value.(string)
		case "Sponsor URL":
			parsedPiece.SponsorURL = v.Value.(string)
		case "Cunning":
			parsedPiece.Cunning = v.Value.(float64)
		case "Rattitude":
			parsedPiece.Rattitude = v.Value.(float64)
		case "Cuteness":
			parsedPiece.Cuteness = v.Value.(float64)
		}
	}
	return parsedPiece
}

func handleClosetTransfers(txs []tokens.ClosetTransfer) []ClosetTokenInput {
	changes := make([]ClosetTokenInput, 0)
	addrToIdListMap := make(map[common.Hash][]*big.Int)
	closet := blockchain.ClosetContract
	opts := blockchain.Opts
	for _, tx := range txs {
		id := tx.Id
		from := tx.From
		to := tx.To
		
		if addrToIdListMap[from] == nil {
			addrToIdListMap[from] = make([]*big.Int, 0)
		}
		exists := false
		for _, v := range addrToIdListMap[from] {
			if v.Cmp(id) == 0 {
				exists = true
			}
		}
		if !exists {
			addrToIdListMap[from] = append(addrToIdListMap[from], id)
		}

		if addrToIdListMap[to] == nil {
			addrToIdListMap[to] = make([]*big.Int, 0)
		}
		exists = false
		for _, v := range addrToIdListMap[to] {
			if v.Cmp(id) == 0 {
				exists = true
			}
		}
		if !exists {
			addrToIdListMap[to] = append(addrToIdListMap[to], id)
		}
	}

	for addr, tokens := range addrToIdListMap {
		if (addr == common.BigToHash(big.NewInt(0))) {
			continue
		}
		log.Println("Handling token changes for address ", common.BytesToAddress(addr.Bytes()))
		accountArr := make([]common.Address, len(tokens))
		tokensArr := make([]*big.Int, 0)
		for i := 0; i < len(accountArr); i++ {
			accountArr[i] = common.BytesToAddress(addr.Bytes())
		}
		for _, id := range tokens {
			tokensArr = append(tokensArr, id)
		}
		bals, err := closet.BalanceOfBatch(opts, accountArr, tokensArr)
		if err != nil {
			log.Println("error: ", err)
		}
		for i, bal := range bals {
			token := ClosetTokenInput{TokenId: tokensArr[i], Owner: common.BytesToAddress(addr.Bytes()), Amount: bal}
			changes = append(changes, token)
		}
	}

	return changes
}