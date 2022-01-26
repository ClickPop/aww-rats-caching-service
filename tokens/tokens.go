package tokens

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/closet"
	"github.com/ethereum/go-ethereum/common"
)

type OpenseaMeta struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Image       string                 `json:"image,omitempty"`
	Attributes  []OpenseaMetaAttribute `json:"attributes,omitempty"`
}

type OpenseaMetaAttribute struct {
	TraitType   string      `json:"trait_type,omitempty"`
	DisplayType string      `json:"display_type,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	MaxValue    int         `json:"max_value,omitempty"`
}

type RatToken struct {
	Owner common.Address `json:"owner,omitempty"`
	URI   string `json:"uri,omitempty"`
}

type RatTokenWithMeta struct {
	RatToken
	OpenseaMeta
}

type RatTokenWithMetaAndId struct {
	Id *big.Int `json:"id,omitempty"`
	RatTokenWithMeta
}

type ClosetTokenWithMeta struct {
	closet.Token
	OpenseaMeta
}

type ClosetTokenWithMetaAndId struct {
	Id *big.Int `json:"id,omitempty"`
	ClosetTokenWithMeta
}

type ClosetTransfers struct {
	Block     *big.Int              `json:"block,omitempty"`
	Transfers map[common.Hash]ClosetTransfer `json:"transfers,omitempty"`
}

type ClosetTransfer struct {
	From   common.Hash `json:"from,omitempty"`
	To     common.Hash `json:"to,omitempty"`
	Id     *big.Int `json:"id,omitempty"`
	Amount *big.Int `json:"amount,omitempty"`
}

func GetRatMeta(uri string) (OpenseaMeta, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return OpenseaMeta{}, nil
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OpenseaMeta{}, err
	}
	var meta OpenseaMeta
	err = json.Unmarshal(bodyData, &meta)
	if err != nil {
		return OpenseaMeta{}, err
	}
	return meta, nil
}

func GetClosetMeta(id *big.Int) (OpenseaMeta, error) {
	resp, err := http.Get(strings.Replace(blockchain.ClosetTokenURI, "{id}", fmt.Sprintf("%s", id.String()), 1))
		if err != nil {
			return OpenseaMeta{}, err
		}
		bodyData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return OpenseaMeta{}, err
		}
		var meta OpenseaMeta
		err = json.Unmarshal(bodyData, &meta)
		if err != nil {
			return OpenseaMeta{}, err
		}
		return meta, nil
}