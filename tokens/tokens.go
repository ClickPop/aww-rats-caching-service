package tokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/closet"
	"github.com/clickpop/aww-rats-caching-service/gcloud"
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
	URI   string         `json:"uri,omitempty"`
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
	Block     *big.Int                       `json:"block,omitempty"`
	Transfers map[common.Hash]ClosetTransfer `json:"transfers,omitempty"`
}

type ClosetTransfer struct {
	From   common.Hash `json:"from,omitempty"`
	To     common.Hash `json:"to,omitempty"`
	Id     *big.Int    `json:"id,omitempty"`
	Amount *big.Int    `json:"amount,omitempty"`
}

func GetRatMeta(uri string) (OpenseaMeta, error) {
	url := strings.Replace(uri, "ipfs://", "https://gateway.pinata.cloud/ipfs/", 1)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		if resp.StatusCode == 429 {
			time.Sleep(time.Second * 5)
			return GetRatMeta(uri)
		}
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
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		if resp.StatusCode == 429 {
			time.Sleep(time.Second * 5)
			return GetClosetMeta(id)
		}
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

func getIpfsImage(uri string) ([]byte, error) {
	url := strings.Replace(uri, "ipfs://", "https://gateway.pinata.cloud/ipfs/", 1)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		if resp.StatusCode == 429 {
			time.Sleep(time.Second * 5)
			return getIpfsImage(uri)
		}
		return []byte{}, err
	}
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyData, err
}

func StoreRat(rat RatTokenWithMetaAndId) error {
	if os.Getenv("IGNORE_IMAGES") == "true" {
		return nil
	}
	data, err := getIpfsImage(rat.Image)
	if err != nil {
		return err
	}
	gcloud.StoreFile(fmt.Sprintf("rats/cached-images/%s.png", rat.Id.String()), data)
	return nil
}

func StoreClosetPiece(piece ClosetTokenWithMetaAndId) error {
	if os.Getenv("IGNORE_IMAGES") == "true" {
		return nil
	}
	data, err := getIpfsImage(piece.Image)
	if err != nil {
		return err
	}
	reg := regexp.MustCompile(`(?m)[^/]+\.png$`)
	filename := reg.FindString(piece.Image)
	if filename == "" {
		return errors.New("cannot find filename in image url")
	}
	gcloud.StoreFile(fmt.Sprintf("closet/image-thumbnails/%s", filename), data)
	return nil
}
