// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package closet

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Token is an auto generated low-level Go binding around an user-defined struct.
type Token struct {
	Name            string
	Cost            *big.Int
	MaxTokens       *big.Int
	MaxPerWallet    *big.Int
	Active          bool
	RevShareAddress common.Address
	RevShareAmount  [2]*big.Int
}

// TokenWithId is an auto generated low-level Go binding around an user-defined struct.
type TokenWithId struct {
	Id    *big.Int
	Token Token
}

// UserToken is an auto generated low-level Go binding around an user-defined struct.
type UserToken struct {
	Id     *big.Int
	Amount *big.Int
	Token  Token
}

// ClosetMetaData contains all meta data concerning the Closet contract.
var ClosetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"BatchTokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"BatchTokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"ChangeERC20Contract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"indexed\":false,\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"name\":\"TokenTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"indexed\":false,\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"name\":\"TokenTypeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"TokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"WalletBanned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"WalletMaxChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"WalletUnbanned\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"name\":\"addNewTokenTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"banWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"changeERC20Contract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"internalType\":\"structTokenWithId[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"name\":\"changeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"existingTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"internalType\":\"structTokenWithId[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"internalType\":\"structTokenWithId[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTokenById\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getTokensByWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"revShareAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"revShareAmount\",\"type\":\"uint256[2]\"}],\"internalType\":\"structToken\",\"name\":\"token\",\"type\":\"tuple\"}],\"internalType\":\"structUserToken[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxTokensPerWalletById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"promoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxTokensForWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setTokensStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"unbanWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletBans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"banned\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b50600054610100900460ff16620000335760005460ff16156200003d565b6200003d620000e2565b620000a55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000c8576000805461ffff19166101011790555b8015620000db576000805461ff00191690555b5062000106565b6000620000fa306200010060201b6200254c1760201c565b15905090565b3b151590565b6080516158c76200013760003960008181610e6501528181610ea5015281816110d8015261111801526158c76000f3fe6080604052600436106102455760003560e01c80638129fc1c11610139578063d3b90b30116100b6578063eed670d71161007a578063eed670d714610729578063f1e54b2314610757578063f242432a14610777578063f2fde38b14610797578063f316534b146107b7578063f52d9e0f146107d757600080fd5b8063d3b90b3014610645578063e046ed0614610672578063e8a3d485146106ab578063e985e9c5146106c0578063ecf9cc241461070957600080fd5b8063a22cb465116100fd578063a22cb465146105b0578063b390c0ab146105d0578063bd85b039146105f0578063bdbed7221461061d578063d351cfdc1461063257600080fd5b80638129fc1c1461051c57806383ca4b6f146105315780638da5cb5b14610551578063938e3d7b146105705780639b642de11461059057600080fd5b80634f1ef286116101c75780636a4aef021161018b5780636a4aef0214610461578063715018a614610481578063736ce85614610496578063785e9e86146104b65780637bdc60d9146104ef57600080fd5b80634f1ef286146103be5780634f558e79146103d157806354fd4d50146104005780635f5817e31461042c57806363d1a9ee1461044157600080fd5b80632a5c792a1161020e5780632a5c792a1461030f5780632eb2c2d6146103315780633659cfe61461035157806338e33e43146103715780634e1273f41461039157600080fd5b8062fdd58e1461024a57806301ffc9a71461027d5780630e89341c146102ad578063138d3f09146102da5780631b2ef1ca146102fc575b600080fd5b34801561025657600080fd5b5061026a61026536600461461e565b6107f7565b6040519081526020015b60405180910390f35b34801561028957600080fd5b5061029d61029836600461465e565b610890565b6040519015158152602001610274565b3480156102b957600080fd5b506102cd6102c836600461467b565b6108e2565b60405161027491906146ec565b3480156102e657600080fd5b506102fa6102f536600461492b565b610976565b005b6102fa61030a366004614a26565b6109e5565b34801561031b57600080fd5b50610324610ba8565b6040516102749190614ad2565b34801561033d57600080fd5b506102fa61034c366004614bb7565b610dc3565b34801561035d57600080fd5b506102fa61036c366004614c60565b610e5a565b34801561037d57600080fd5b506102fa61038c366004614c60565b610f23565b34801561039d57600080fd5b506103b16103ac366004614c7b565b610fa4565b6040516102749190614d80565b6102fa6103cc366004614d93565b6110cd565b3480156103dd57600080fd5b5061029d6103ec36600461467b565b600090815260976020526040902054151590565b34801561040c57600080fd5b506040805180820190915260038152620312e360ec1b60208201526102cd565b34801561043857600080fd5b50610324611183565b34801561044d57600080fd5b506102fa61045c366004614d93565b6112cb565b34801561046d57600080fd5b506102fa61047c366004614c60565b611409565b34801561048d57600080fd5b506102fa6114cb565b3480156104a257600080fd5b506102fa6104b1366004614dd6565b611502565b3480156104c257600080fd5b50610163546104d7906001600160a01b031681565b6040516001600160a01b039091168152602001610274565b3480156104fb57600080fd5b5061050f61050a36600461467b565b611716565b6040516102749190614e27565b34801561052857600080fd5b506102fa611849565b34801561053d57600080fd5b506102fa61054c366004614e3a565b61198f565b34801561055d57600080fd5b5061012d546001600160a01b03166104d7565b34801561057c57600080fd5b506102fa61058b366004614e86565b611a4b565b34801561059c57600080fd5b506102fa6105ab366004614e86565b611a8a565b3480156105bc57600080fd5b506102fa6105cb366004614ec2565b611abe565b3480156105dc57600080fd5b506102fa6105eb366004614a26565b611ac9565b3480156105fc57600080fd5b5061026a61060b36600461467b565b60009081526097602052604090205490565b34801561062957600080fd5b506103b1611b20565b6102fa610640366004614e3a565b611b79565b34801561065157600080fd5b50610665610660366004614c60565b611c07565b6040516102749190614eee565b34801561067e57600080fd5b5061026a61068d366004614f5f565b61016160209081526000928352604080842090915290825290205481565b3480156106b757600080fd5b506102cd611e81565b3480156106cc57600080fd5b5061029d6106db366004614f8b565b6001600160a01b03918216600090815260666020908152604080832093909416825291909152205460ff1690565b34801561071557600080fd5b5061026a61072436600461467b565b611f0b565b34801561073557600080fd5b50610749610744366004614c60565b611f2d565b604051610274929190614fb5565b34801561076357600080fd5b506102fa610772366004614fd0565b611fd7565b34801561078357600080fd5b506102fa610792366004615003565b61215e565b3480156107a357600080fd5b506102fa6107b2366004614c60565b6121e5565b3480156107c357600080fd5b506102fa6107d2366004615067565b61227e565b3480156107e357600080fd5b506102fa6107f2366004615122565b612443565b60006001600160a01b0383166108685760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b5060009081526065602090815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b14806108c157506001600160e01b031982166303a24d0760e21b145b806108dc57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060606780546108f190615195565b80601f016020809104026020016040519081016040528092919081815260200182805461091d90615195565b801561096a5780601f1061093f5761010080835404028352916020019161096a565b820191906000526020600020905b81548152906001019060200180831161094d57829003601f168201915b50505050509050919050565b61012d546001600160a01b031633146109a15760405162461bcd60e51b815260040161085f906151ca565b60005b81518110156109e1576109cf8282815181106109c2576109c26151ff565b6020026020010151612552565b806109d98161522b565b9150506109a4565b5050565b81816109f1828261266d565b61012d546000858152610160602052604090206004015461010090046001600160a01b03908116911614610ae65760008481526101606020526040812060010154610a3d908590615246565b60008681526101606020526040812060058101546006909101549293509091610a669084615265565b610a709190615246565b6000878152610160602052604090206004015461016354919250610aa9916001600160a01b0390811691339161010090910416846126cd565b610adf33610ac061012d546001600160a01b031690565b610aca8486615287565b610163546001600160a01b03169291906126cd565b5050610b1b565b610b1b33610afd61012d546001600160a01b031690565b60008781526101606020526040902060010154610aca908790615246565b610b5e3385856000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061272792505050565b604080518581526020810185905233918101919091527f06437edc6222bc3b01fd946ea7b36242878aef9a533aa8815121bb23dc0a7438906060015b60405180910390a150505050565b61015f546060906000906001600160401b03811115610bc957610bc96146ff565b604051908082528060200260200182016040528015610c0257816020015b610bef614496565b815260200190600190039081610be75790505b50905060005b61015f54811015610dbd57604051806040016040528061015f8381548110610c3257610c326151ff565b90600052602060002001548152602001610160600061015f8581548110610c5b57610c5b6151ff565b906000526020600020015481526020019081526020016000206040518060e0016040529081600082018054610c8f90615195565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbb90615195565b8015610d085780601f10610cdd57610100808354040283529160200191610d08565b820191906000526020600020905b815481529060010190602001808311610ceb57829003601f168201915b50505091835250506001820154602082015260028083015460408084019190915260038401546060840152600484015460ff81161515608085015261010090046001600160a01b031660a084015280518082019182905260c09093019291600585019182845b815481526020019060010190808311610d6e57505050505081525050815250828281518110610d9f57610d9f6151ff565b60200260200101819052508080610db59061522b565b915050610c08565b50919050565b6001600160a01b038516331480610ddf5750610ddf85336106db565b610e465760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b606482015260840161085f565b610e5385858585856127ff565b5050505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415610ea35760405162461bcd60e51b815260040161085f9061529e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610ed5612889565b6001600160a01b031614610efb5760405162461bcd60e51b815260040161085f906152ea565b610f04816128b7565b60408051600080825260208201909252610f20918391906128e2565b50565b61012d546001600160a01b03163314610f4e5760405162461bcd60e51b815260040161085f906151ca565b61016380546001600160a01b0319166001600160a01b0383169081179091556040519081527ff4437fae6d1fbef7b2436233184e29ceee39249efd21c29c42991b93beb52be4906020015b60405180910390a150565b606081518351146110095760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b606482015260840161085f565b600083516001600160401b03811115611024576110246146ff565b60405190808252806020026020018201604052801561104d578160200160208202803683370190505b50905060005b84518110156110c557611098858281518110611071576110716151ff565b602002602001015185838151811061108b5761108b6151ff565b60200260200101516107f7565b8282815181106110aa576110aa6151ff565b60209081029190910101526110be8161522b565b9050611053565b509392505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156111165760405162461bcd60e51b815260040161085f9061529e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611148612889565b6001600160a01b03161461116e5760405162461bcd60e51b815260040161085f906152ea565b611177826128b7565b6109e1828260016128e2565b6060600061118f610ba8565b90506000805b82518110156111e5578281815181106111b0576111b06151ff565b60200260200101516020015160800151156111d357816111cf8161522b565b9250505b806111dd8161522b565b915050611195565b506000816001600160401b03811115611200576112006146ff565b60405190808252806020026020018201604052801561123957816020015b611226614496565b81526020019060019003908161121e5790505b5090506000915060005b83518110156110c55783818151811061125e5761125e6151ff565b60200260200101516020015160800151156112b957838181518110611285576112856151ff565b602002602001015182848151811061129f5761129f6151ff565b602002602001018190525082806112b59061522b565b9350505b806112c38161522b565b915050611243565b61012d546001600160a01b031633146112f65760405162461bcd60e51b815260040161085f906151ca565b604080516020808201909252600090528151908201207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47014156113745760405162461bcd60e51b8152602060048201526016602482015275526561736f6e2063616e6e6f7420626520656d70747960501b604482015260640161085f565b604080518082018252600180825260208083018581526001600160a01b03871660009081526101628352949094208351815460ff1916901515178155935180519394936113c89385019291909101906144b5565b509050507fc644f65bce175b55083b5f7cb20c238f0ee0ead74eeef083e6874ec3db520d7182826040516113fd929190615336565b60405180910390a15050565b61012d546001600160a01b031633146114345760405162461bcd60e51b815260040161085f906151ca565b60408051808201825260008082528251602080820185528282528084019182526001600160a01b03861683526101628152939091208251815460ff1916901515178155905180519293919261148f92600185019201906144b5565b50506040516001600160a01b03831681527fb2ed1ab64220fe50bc8b9058c50791a59bf4e36b5fd723d425c857b0a468fe619150602001610f99565b61012d546001600160a01b031633146114f65760405162461bcd60e51b815260040161085f906151ca565b6115006000612a26565b565b61012d546001600160a01b0316331461152d5760405162461bcd60e51b815260040161085f906151ca565b60005b825181101561171157816101606000858481518110611551576115516151ff565b6020026020010151815260200190815260200160002060040160006101000a81548160ff021916908315150217905550600060405180604001604052808584815181106115a0576115a06151ff565b6020026020010151815260200161016060008786815181106115c4576115c46151ff565b602002602001015181526020019081526020016000206040518060e00160405290816000820180546115f590615195565b80601f016020809104026020016040519081016040528092919081815260200182805461162190615195565b801561166e5780601f106116435761010080835404028352916020019161166e565b820191906000526020600020905b81548152906001019060200180831161165157829003601f168201915b50505091835250506001820154602082015260028083015460408084019190915260038401546060840152600484015460ff81161515608085015261010090046001600160a01b031660a084015280518082019182905260c09093019291600585019182845b8154815260200190600101908083116116d45750505050508152505081525090506116fe81612552565b50806117098161522b565b915050611530565b505050565b61171e614539565b6000828152610160602052604090819020815160e0810190925280548290829061174790615195565b80601f016020809104026020016040519081016040528092919081815260200182805461177390615195565b80156117c05780601f10611795576101008083540402835291602001916117c0565b820191906000526020600020905b8154815290600101906020018083116117a357829003601f168201915b50505091835250506001820154602082015260028083015460408084019190915260038401546060840152600484015460ff81161515608085015261010090046001600160a01b031660a084015280518082019182905260c09093019291600585019182845b815481526020019060010190808311611826575050505050815250509050919050565b600054610100900460ff166118645760005460ff1615611868565b303b155b6118cb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161085f565b600054610100900460ff161580156118ed576000805461ffff19166101011790555b6118f5612a79565b6118fd612ab0565b61193b6040518060400160405280601d81526020017f68747470733a2f2f617777726174732e636f6d2f7b69647d2e6a736f6e000000815250612aef565b60405180606001604052806030815260200161583b60309139805161196991610164916020909101906144b5565b5061016380546001600160a01b03191690558015610f20576000805461ff001916905550565b818180518251146119b25760405162461bcd60e51b815260040161085f9061535a565b60005b8251811015611a0c576119fa8382815181106119d3576119d36151ff565b60200260200101518383815181106119ed576119ed6151ff565b6020026020010151612b02565b80611a048161522b565b9150506119b5565b50611a18338585612b98565b7f3c53c7744097cc8b22f1268b0b5c8b8c8233acf4f6249b7de76d0d5cbf349c19848433604051610b9a93929190615391565b61012d546001600160a01b03163314611a765760405162461bcd60e51b815260040161085f906151ca565b80516109e1906101649060208401906144b5565b61012d546001600160a01b03163314611ab55760405162461bcd60e51b815260040161085f906151ca565b610f2081612aef565b6109e1338383612d29565b8181611ad58282612b02565b611ae0338585612e0a565b604080518581526020810185905233918101919091527f8b70448d0d143d815676e1088155bf2ff4436b50d76b30ff3a1fe2929089270290606001610b9a565b606061015f805480602002602001604051908101604052809291908181526020018280548015611b6f57602002820191906000526020600020905b815481526020019060010190808311611b5b575b5050505050905090565b81818051825114611b9c5760405162461bcd60e51b815260040161085f9061535a565b60005b8251811015611bf657611be4838281518110611bbd57611bbd6151ff565b6020026020010151838381518110611bd757611bd76151ff565b602002602001015161266d565b80611bee8161522b565b915050611b9f565b50611c018484612f0f565b50505050565b61015f546060906000906001600160401b03811115611c2857611c286146ff565b604051908082528060200260200182016040528015611c6157816020015b611c4e614581565b815260200190600190039081611c465790505b5090506000805b61015f54811015611e78576000611c9d8661015f8481548110611c8d57611c8d6151ff565b90600052602060002001546107f7565b1115611e6657604051806060016040528061015f8381548110611cc257611cc26151ff565b90600052602060002001548152602001611cea8761015f8581548110611c8d57611c8d6151ff565b8152602001610160600061015f8581548110611d0857611d086151ff565b906000526020600020015481526020019081526020016000206040518060e0016040529081600082018054611d3c90615195565b80601f0160208091040260200160405190810160405280929190818152602001828054611d6890615195565b8015611db55780601f10611d8a57610100808354040283529160200191611db5565b820191906000526020600020905b815481529060010190602001808311611d9857829003601f168201915b50505091835250506001820154602082015260028083015460408084019190915260038401546060840152600484015460ff81161515608085015261010090046001600160a01b031660a084015280518082019182905260c09093019291600585019182845b815481526020019060010190808311611e1b57505050505081525050815250838381518110611e4c57611e4c6151ff565b60200260200101819052508180611e629061522b565b9250505b80611e708161522b565b915050611c68565b50909392505050565b60606101648054611e9190615195565b80601f0160208091040260200160405190810160405280929190818152602001828054611ebd90615195565b8015611b6f5780601f10611edf57610100808354040283529160200191611b6f565b820191906000526020600020905b815481529060010190602001808311611eed57509395945050505050565b61015f8181548110611f1c57600080fd5b600091825260209091200154905081565b610162602052600090815260409020805460018201805460ff9092169291611f5490615195565b80601f0160208091040260200160405190810160405280929190818152602001828054611f8090615195565b8015611fcd5780601f10611fa257610100808354040283529160200191611fcd565b820191906000526020600020905b815481529060010190602001808311611fb057829003601f168201915b5050505050905082565b61012d546001600160a01b031633146120025760405162461bcd60e51b815260040161085f906151ca565b600082815261016060205260409020805483916120de9161202290615195565b80601f016020809104026020016040519081016040528092919081815260200182805461204e90615195565b801561209b5780601f106120705761010080835404028352916020019161209b565b820191906000526020600020905b81548152906001019060200180831161207e57829003601f168201915b50506040805160208082019092526000905284519401939093207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470149392505050565b156120fb5760405162461bcd60e51b815260040161085f906153cf565b6000838152610161602090815260408083206001600160a01b03881680855290835292819020859055805192835290820185905281018390527f6b30946421f69c81ed9a9f3334ecf2a5ed5b0bcc149d52c2900fa0ce5f93076090606001610b9a565b6001600160a01b03851633148061217a575061217a85336106db565b6121d85760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b606482015260840161085f565b610e53858585858561335f565b61012d546001600160a01b031633146122105760405162461bcd60e51b815260040161085f906151ca565b6001600160a01b0381166122755760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161085f565b610f2081612a26565b61012d546001600160a01b031633146122a95760405162461bcd60e51b815260040161085f906151ca565b60005b81518110156109e15760005b61015f5481101561240e576123b8610160600061015f84815481106122df576122df6151ff565b90600052602060002001548152602001908152602001600020600001805461230690615195565b80601f016020809104026020016040519081016040528092919081815260200182805461233290615195565b801561237f5780601f106123545761010080835404028352916020019161237f565b820191906000526020600020905b81548152906001019060200180831161236257829003601f168201915b5050505050848481518110612396576123966151ff565b6020026020010151600001518051602091820120825192909101919091201490565b156123fc5760405162461bcd60e51b8152602060048201526014602482015273546f6b656e20616c72656164792065786973747360601b604482015260640161085f565b806124068161522b565b9150506122b8565b50612431828281518110612424576124246151ff565b6020026020010151613377565b8061243b8161522b565b9150506122ac565b61012d546001600160a01b0316331461246e5760405162461bcd60e51b815260040161085f906151ca565b828280518251146124915760405162461bcd60e51b815260040161085f9061535a565b60005b82518110156124eb576124d98382815181106124b2576124b26151ff565b60200260200101518383815181106124cc576124cc6151ff565b6020026020010151613498565b806124e38161522b565b915050612494565b506124f68585612f0f565b610e5361250c61012d546001600160a01b031690565b8487876000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610dc392505050565b3b151590565b80516000818152610160602052604090208054612573919061202290615195565b156125905760405162461bcd60e51b815260040161085f906153cf565b6020808301518351600090815261016083526040902081518051929391926125bb92849201906144b5565b5060208201516001820155604082015160028083019190915560608301516003830155608083015160048301805460a08601516001600160a01b031661010002610100600160a81b0319931515939093166001600160a81b03199091161791909117905560c0830151612633916005840191906145a2565b5050825160208401516040517f59a0de9e767f7ea1fa2a7afb6ba676db02d66e9e7781d76a8e83d0cc0787943593506113fd929190615406565b6000828152610160602052604090206004015460ff166126c35760405162461bcd60e51b8152602060048201526011602482015270546f6b656e20697320696e61637469766560781b604482015260640161085f565b6109e18282613498565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611c01908590613684565b6001600160a01b03841661274d5760405162461bcd60e51b815260040161085f9061541f565b3361276d8160008761275e88613756565b61276788613756565b876137a1565b60008481526065602090815260408083206001600160a01b03891684529091528120805485929061279f908490615460565b909155505060408051858152602081018590526001600160a01b0380881692600092918516917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610e53816000878787876138b7565b81518351146128205760405162461bcd60e51b815260040161085f9061535a565b60005b835181101561287b57612869848281518110612841576128416151ff565b602002602001015184838151811061285b5761285b6151ff565b602002602001015187613a13565b806128738161522b565b915050612823565b50610e538585858585613b4f565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b61012d546001600160a01b03163314610f205760405162461bcd60e51b815260040161085f906151ca565b60006128ec612889565b90506128f784613cf4565b6000835111806129045750815b15612915576129138484613d99565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff16610e5357805460ff191660011781556040516001600160a01b038316602482015261299490869060440160408051601f198184030181529190526020810180516001600160e01b0316631b2ce7f360e11b179052613d99565b50805460ff191681556129a5612889565b6001600160a01b0316826001600160a01b031614612a1d5760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201526e75727468657220757067726164657360881b606482015260840161085f565b610e5385613e84565b61012d80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16612aa05760405162461bcd60e51b815260040161085f90615478565b612aa8613ec4565b611500613eeb565b600054610100900460ff16612ad75760405162461bcd60e51b815260040161085f90615478565b612adf613ec4565b612ae7613ec4565b611500613ec4565b80516109e19060679060208401906144b5565b60008281526101606020526040902080548391612b229161202290615195565b15612b3f5760405162461bcd60e51b815260040161085f906153cf565b81612b4a33856107f7565b10156117115760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f74206275726e206d6f7265207468616e206f776e65640000000000604482015260640161085f565b6001600160a01b038316612bbe5760405162461bcd60e51b815260040161085f906154c3565b8051825114612bdf5760405162461bcd60e51b815260040161085f90615506565b6000339050612c02818560008686604051806020016040528060008152506137a1565b60005b8351811015612cca576000848281518110612c2257612c226151ff565b602002602001015190506000848381518110612c4057612c406151ff565b60209081029190910181015160008481526065835260408082206001600160a01b038c168352909352919091205490915081811015612c915760405162461bcd60e51b815260040161085f9061554e565b60009283526065602090815260408085206001600160a01b038b1686529091529092209103905580612cc28161522b565b915050612c05565b5060006001600160a01b0316846001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8686604051612d1b929190615592565b60405180910390a450505050565b816001600160a01b0316836001600160a01b03161415612d9d5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b606482015260840161085f565b6001600160a01b03838116600081815260666020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b038316612e305760405162461bcd60e51b815260040161085f906154c3565b33612e5f81856000612e4187613756565b612e4a87613756565b604051806020016040528060008152506137a1565b60008381526065602090815260408083206001600160a01b038816845290915290205482811015612ea25760405162461bcd60e51b815260040161085f9061554e565b60008481526065602090815260408083206001600160a01b03898116808652918452828520888703905582518981529384018890529092908616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a45050505050565b60008083516001600160401b03811115612f2b57612f2b6146ff565b604051908082528060200260200182016040528015612f54578160200160208202803683370190505b509050600084516001600160401b03811115612f7257612f726146ff565b604051908082528060200260200182016040528015612f9b578160200160208202803683370190505b5090506000805b865181101561323e5761012d546001600160a01b03166001600160a01b03166101606000898481518110612fd857612fd86151ff565b6020026020010151815260200190815260200160002060040160019054906101000a90046001600160a01b03166001600160a01b0316146131ca576000868281518110613027576130276151ff565b602002602001015161016060008a8581518110613046576130466151ff565b602002602001015181526020019081526020016000206001015461306a9190615246565b9050600061016060008a8581518110613085576130856151ff565b602002602001015181526020019081526020016000206005016000600281106130b0576130b06151ff565b015461016060008b86815181106130c9576130c96151ff565b602002602001015181526020019081526020016000206005016001600281106130f4576130f46151ff565b01546131009084615265565b61310a9190615246565b905061016060008a8581518110613123576131236151ff565b6020026020010151815260200190815260200160002060040160019054906101000a90046001600160a01b0316858581518110613162576131626151ff565b60200260200101906001600160a01b031690816001600160a01b03168152505080868581518110613195576131956151ff565b60209081029190910101526131aa8183615287565b6131b49088615460565b9650836131c08161522b565b945050505061322c565b8581815181106131dc576131dc6151ff565b602002602001015161016060008984815181106131fb576131fb6151ff565b602002602001015181526020019081526020016000206001015461321f9190615246565b6132299086615460565b94505b806132368161522b565b915050612fa2565b5060005b818110156132ac5761329a33848381518110613260576132606151ff565b602002602001015186848151811061327a5761327a6151ff565b6020908102919091010151610163546001600160a01b03169291906126cd565b806132a48161522b565b915050613242565b506132d9336132c461012d546001600160a01b031690565b610163546001600160a01b03169190876126cd565b61331c3387876000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613f1b92505050565b7fddc7a7105147d9dc422856b8dac226d6e644f688cf9ea1ea10c67cbc402dc38c86863360405161334f93929190615391565b60405180910390a1505050505050565b61336a838386613a13565b610e538585858585614076565b61015f54600090613389906001615460565b60008181526101606020908152604090912084518051939450859391926133b5928492909101906144b5565b5060208201516001820155604082015160028083019190915560608301516003830155608083015160048301805460a08601516001600160a01b031661010002610100600160a81b0319931515939093166001600160a81b03199091161791909117905560c083015161342d916005840191906145a2565b505061015f80546001810182556000919091527f8c9815a58669fc89297bdc7dd447c098116f98e79093735c0992d0967b696ed901829055506040517fcbd47bdaacef607c448cfd0079c791bcba6a3d01ac2999cdbd4afec8c91c0838906113fd9083908590615406565b600082815261016060205260409020805483916134b89161202290615195565b156134d55760405162461bcd60e51b815260040161085f906153cf565b6134df838361418e565b6134ea838333613a13565b6000838152610160602052604090819020600101546101635491516370a0823160e01b815233600482015290916001600160a01b0316906370a0823190602401602060405180830381865afa158015613547573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061356b91906155b7565b10156135af5760405162461bcd60e51b81526020600482015260136024820152724e6f7420656e6f7567682063757272656e637960681b604482015260640161085f565b600083815261016060205260409081902060010154610163549151636eb1769f60e11b815233600482015230602482015290916001600160a01b03169063dd62ed3e90604401602060405180830381865afa158015613612573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061363691906155b7565b10156117115760405162461bcd60e51b815260206004820152601a60248201527f455243323020616c6c6f77616e6365206e6f7420656e6f756768000000000000604482015260640161085f565b60006136d9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166142619092919063ffffffff16565b80519091501561171157808060200190518101906136f791906155d0565b6117115760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161085f565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110613790576137906151ff565b602090810291909101015292915050565b6001600160a01b0385166138285760005b8351811015613826578281815181106137cd576137cd6151ff565b6020026020010151609760008684815181106137eb576137eb6151ff565b6020026020010151815260200190815260200160002060008282546138109190615460565b9091555061381f90508161522b565b90506137b2565b505b6001600160a01b0384166138af5760005b83518110156138ad57828181518110613854576138546151ff565b602002602001015160976000868481518110613872576138726151ff565b6020026020010151815260200190815260200160002060008282546138979190615287565b909155506138a690508161522b565b9050613839565b505b505050505050565b6001600160a01b0384163b156138af5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906138fb90899089908890889088906004016155ed565b6020604051808303816000875af1925050508015613936575060408051601f3d908101601f1916820190925261393391810190615627565b60015b6139e357613942615644565b806308c379a0141561397c5750613957615660565b80613962575061397e565b8060405162461bcd60e51b815260040161085f91906146ec565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b606482015260840161085f565b6001600160e01b0319811663f23a6e6160e01b146138ad5760405162461bcd60e51b815260040161085f906156e9565b60008381526101606020526040902080548491613a339161202290615195565b15613a505760405162461bcd60e51b815260040161085f906153cf565b6000848152610161602090815260408083206001600160a01b03861684529091529020541580613ab7575082613a8683866107f7565b613a909190615460565b6000858152610161602090815260408083206001600160a01b038716845290915290205410155b8015613b035750600084815261016060205260409020600301541580613b03575082613ae383866107f7565b613aed9190615460565b6000858152610160602052604090206003015410155b611c015760405162461bcd60e51b815260206004820152601d60248201527f4d617820746f6b656e73207265616368656420666f722077616c6c6574000000604482015260640161085f565b8151835114613b705760405162461bcd60e51b815260040161085f90615506565b6001600160a01b038416613b965760405162461bcd60e51b815260040161085f90615731565b33613ba58187878787876137a1565b60005b8451811015613c8e576000858281518110613bc557613bc56151ff565b602002602001015190506000858381518110613be357613be36151ff565b60209081029190910181015160008481526065835260408082206001600160a01b038e168352909352919091205490915081811015613c345760405162461bcd60e51b815260040161085f90615776565b60008381526065602090815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290613c73908490615460565b9250508190555050505080613c879061522b565b9050613ba8565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051613cde929190615592565b60405180910390a46138af81878787878761427a565b803b613d585760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161085f565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b613df85760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161085f565b600080846001600160a01b031684604051613e1391906157c0565b600060405180830381855af49150503d8060008114613e4e576040519150601f19603f3d011682016040523d82523d6000602084013e613e53565b606091505b5091509150613e7b828260405180606001604052806027815260200161586b60279139614335565b95945050505050565b613e8d81613cf4565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff166115005760405162461bcd60e51b815260040161085f90615478565b600054610100900460ff16613f125760405162461bcd60e51b815260040161085f90615478565b61150033612a26565b6001600160a01b038416613f415760405162461bcd60e51b815260040161085f9061541f565b8151835114613f625760405162461bcd60e51b815260040161085f90615506565b33613f72816000878787876137a1565b60005b845181101561400e57838181518110613f9057613f906151ff565b602002602001015160656000878481518110613fae57613fae6151ff565b602002602001015181526020019081526020016000206000886001600160a01b03166001600160a01b031681526020019081526020016000206000828254613ff69190615460565b909155508190506140068161522b565b915050613f75565b50846001600160a01b031660006001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb878760405161405f929190615592565b60405180910390a4610e538160008787878761427a565b6001600160a01b03841661409c5760405162461bcd60e51b815260040161085f90615731565b336140ac81878761275e88613756565b60008481526065602090815260408083206001600160a01b038a168452909152902054838110156140ef5760405162461bcd60e51b815260040161085f90615776565b60008581526065602090815260408083206001600160a01b038b811685529252808320878503905590881682528120805486929061412e908490615460565b909155505060408051868152602081018690526001600160a01b03808916928a821692918616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46138ad8288888888886138b7565b600082815261016060205260409020805483916141ae9161202290615195565b156141cb5760405162461bcd60e51b815260040161085f906153cf565b60008381526101606020526040902060020154158061421557506000838152609760205260409020546141ff908390615460565b6000848152610160602052604090206002015410155b6117115760405162461bcd60e51b815260206004820152601b60248201527f4d617820746f6b656e73207265616368656420666f7220747970650000000000604482015260640161085f565b6060614270848460008561436e565b90505b9392505050565b6001600160a01b0384163b156138af5760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906142be90899089908890889088906004016157dc565b6020604051808303816000875af19250505080156142f9575060408051601f3d908101601f191682019092526142f691810190615627565b60015b61430557613942615644565b6001600160e01b0319811663bc197c8160e01b146138ad5760405162461bcd60e51b815260040161085f906156e9565b60608315614344575081614273565b8251156143545782518084602001fd5b8160405162461bcd60e51b815260040161085f91906146ec565b6060824710156143cf5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161085f565b843b61441d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161085f565b600080866001600160a01b0316858760405161443991906157c0565b60006040518083038185875af1925050503d8060008114614476576040519150601f19603f3d011682016040523d82523d6000602084013e61447b565b606091505b509150915061448b828286614335565b979650505050505050565b6040518060400160405280600081526020016144b0614539565b905290565b8280546144c190615195565b90600052602060002090601f0160209004810192826144e35760008555614529565b82601f106144fc57805160ff1916838001178555614529565b82800160010185558215614529579182015b8281111561452957825182559160200191906001019061450e565b506145359291506145cf565b5090565b6040518060e001604052806060815260200160008152602001600081526020016000815260200160001515815260200160006001600160a01b031681526020016144b06145e4565b604051806060016040528060008152602001600081526020016144b0614539565b8260028101928215614529579160200282018281111561452957825182559160200191906001019061450e565b5b8082111561453557600081556001016145d0565b60405180604001604052806002906020820280368337509192915050565b80356001600160a01b038116811461461957600080fd5b919050565b6000806040838503121561463157600080fd5b61463a83614602565b946020939093013593505050565b6001600160e01b031981168114610f2057600080fd5b60006020828403121561467057600080fd5b813561427381614648565b60006020828403121561468d57600080fd5b5035919050565b60005b838110156146af578181015183820152602001614697565b83811115611c015750506000910152565b600081518084526146d8816020860160208601614694565b601f01601f19169290920160200192915050565b60208152600061427360208301846146c0565b634e487b7160e01b600052604160045260246000fd5b604081018181106001600160401b0382111715614734576147346146ff565b60405250565b601f8201601f191681016001600160401b038111828210171561475f5761475f6146ff565b6040525050565b60405160e081016001600160401b0381118282101715614788576147886146ff565b60405290565b60006001600160401b038211156147a7576147a76146ff565b5060051b60200190565b600082601f8301126147c257600080fd5b81356001600160401b038111156147db576147db6146ff565b6040516147f2601f8301601f19166020018261473a565b81815284602083860101111561480757600080fd5b816020850160208301376000918101602001919091529392505050565b8015158114610f2057600080fd5b803561461981614824565b600082601f83011261484e57600080fd5b60405161485a81614715565b80604084018581111561486c57600080fd5b845b8181101561488657803583526020928301920161486e565b509195945050505050565b600061010082840312156148a457600080fd5b6148ac614766565b905081356001600160401b038111156148c457600080fd5b6148d0848285016147b1565b8252506020820135602082015260408201356040820152606082013560608201526148fd60808301614832565b608082015261490e60a08301614602565b60a08201526149208360c0840161483d565b60c082015292915050565b6000602080838503121561493e57600080fd5b82356001600160401b038082111561495557600080fd5b818501915085601f83011261496957600080fd5b81356149748161478e565b60408051614982838261473a565b83815260059390931b85018601928681019250898411156149a257600080fd5b8686015b84811015614a18578035868111156149be5760008081fd5b8701808c03601f19018413156149d45760008081fd5b83516149df81614715565b89820135815284820135888111156149f75760008081fd5b614a058e8c83860101614891565b828c0152508552509287019287016149a6565b509998505050505050505050565b60008060408385031215614a3957600080fd5b50508035926020909101359150565b60006101008251818552614a5e828601826146c0565b91505060208084015181860152604084015160408601526060840151606086015260808401511515608086015260018060a01b0360a08501511660a086015260c084015160c0860160005b6002811015614ac657825182529183019190830190600101614aa9565b50929695505050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015614b3857888303603f19018552815180518452870151878401879052614b2587850182614a48565b9588019593505090860190600101614af9565b509098975050505050505050565b600082601f830112614b5757600080fd5b81356020614b648261478e565b604051614b71828261473a565b83815260059390931b8501820192828101915086841115614b9157600080fd5b8286015b84811015614bac5780358352918301918301614b95565b509695505050505050565b600080600080600060a08688031215614bcf57600080fd5b614bd886614602565b9450614be660208701614602565b935060408601356001600160401b0380821115614c0257600080fd5b614c0e89838a01614b46565b94506060880135915080821115614c2457600080fd5b614c3089838a01614b46565b93506080880135915080821115614c4657600080fd5b50614c53888289016147b1565b9150509295509295909350565b600060208284031215614c7257600080fd5b61427382614602565b60008060408385031215614c8e57600080fd5b82356001600160401b0380821115614ca557600080fd5b818501915085601f830112614cb957600080fd5b81356020614cc68261478e565b604051614cd3828261473a565b83815260059390931b8501820192828101915089841115614cf357600080fd5b948201945b83861015614d1857614d0986614602565b82529482019490820190614cf8565b96505086013592505080821115614d2e57600080fd5b50614d3b85828601614b46565b9150509250929050565b600081518084526020808501945080840160005b83811015614d7557815187529582019590820190600101614d59565b509495945050505050565b6020815260006142736020830184614d45565b60008060408385031215614da657600080fd5b614daf83614602565b915060208301356001600160401b03811115614dca57600080fd5b614d3b858286016147b1565b60008060408385031215614de957600080fd5b82356001600160401b03811115614dff57600080fd5b614e0b85828601614b46565b9250506020830135614e1c81614824565b809150509250929050565b6020815260006142736020830184614a48565b60008060408385031215614e4d57600080fd5b82356001600160401b0380821115614e6457600080fd5b614e7086838701614b46565b93506020850135915080821115614d2e57600080fd5b600060208284031215614e9857600080fd5b81356001600160401b03811115614eae57600080fd5b614eba848285016147b1565b949350505050565b60008060408385031215614ed557600080fd5b614ede83614602565b91506020830135614e1c81614824565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015614b3857888303603f1901855281518051845287810151888501528601516060878501819052614f4b81860183614a48565b968901969450505090860190600101614f15565b60008060408385031215614f7257600080fd5b82359150614f8260208401614602565b90509250929050565b60008060408385031215614f9e57600080fd5b614fa783614602565b9150614f8260208401614602565b821515815260406020820152600061427060408301846146c0565b600080600060608486031215614fe557600080fd5b614fee84614602565b95602085013595506040909401359392505050565b600080600080600060a0868803121561501b57600080fd5b61502486614602565b945061503260208701614602565b9350604086013592506060860135915060808601356001600160401b0381111561505b57600080fd5b614c53888289016147b1565b6000602080838503121561507a57600080fd5b82356001600160401b038082111561509157600080fd5b818501915085601f8301126150a557600080fd5b81356150b08161478e565b6040516150bd828261473a565b82815260059290921b84018501918581019150888311156150dd57600080fd5b8585015b83811015615115578035858111156150f95760008081fd5b6151078b89838a0101614891565b8452509186019186016150e1565b5098975050505050505050565b60008060006060848603121561513757600080fd5b83356001600160401b038082111561514e57600080fd5b61515a87838801614b46565b9450602086013591508082111561517057600080fd5b5061517d86828701614b46565b92505061518c60408501614602565b90509250925092565b600181811c908216806151a957607f821691505b60208210811415610dbd57634e487b7160e01b600052602260045260246000fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561523f5761523f615215565b5060010190565b600081600019048311821515161561526057615260615215565b500290565b60008261528257634e487b7160e01b600052601260045260246000fd5b500490565b60008282101561529957615299615215565b500390565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6001600160a01b0383168152604060208201819052600090614270908301846146c0565b6020808252601f908201527f4d69736d61746368656420746f6b656e49647320616e6420616d6f756e747300604082015260600190565b6060815260006153a46060830186614d45565b82810360208401526153b68186614d45565b91505060018060a01b0383166040830152949350505050565b60208082526017908201527f546f6b656e20494420646f6573206e6f74206578697374000000000000000000604082015260600190565b8281526040602082015260006142706040830184614a48565b60208082526021908201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736040820152607360f81b606082015260800190565b6000821982111561547357615473615215565b500190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60208082526023908201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b60208082526024908201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604082015263616e636560e01b606082015260800190565b6040815260006155a56040830185614d45565b8281036020840152613e7b8185614d45565b6000602082840312156155c957600080fd5b5051919050565b6000602082840312156155e257600080fd5b815161427381614824565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061448b908301846146c0565b60006020828403121561563957600080fd5b815161427381614648565b600060033d111561565d5760046000803e5060005160e01c5b90565b600060443d101561566e5790565b6040516003193d81016004833e81513d6001600160401b03816024840111818411171561569d57505050505090565b82850191508151818111156156b55750505050505090565b843d87010160208285010111156156cf5750505050505090565b6156de6020828601018761473a565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b600082516157d2818460208701614694565b9190910192915050565b6001600160a01b0386811682528516602082015260a06040820181905260009061580890830186614d45565b828103606084015261581a8186614d45565b9050828103608084015261582e81856146c0565b9897505050505050505056fe68747470733a2f2f617777726174732e636f6d2f636c6f7365742d6f70656e7365612d6d657461646174612e6a736f6e416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202d94ffee18227e7fb02226ef279549b52684f5f0e3ea2646ddf20e883b22401264736f6c634300080b0033",
}

// ClosetABI is the input ABI used to generate the binding from.
// Deprecated: Use ClosetMetaData.ABI instead.
var ClosetABI = ClosetMetaData.ABI

// ClosetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClosetMetaData.Bin instead.
var ClosetBin = ClosetMetaData.Bin

// DeployCloset deploys a new Ethereum contract, binding an instance of Closet to it.
func DeployCloset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Closet, error) {
	parsed, err := ClosetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClosetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Closet{ClosetCaller: ClosetCaller{contract: contract}, ClosetTransactor: ClosetTransactor{contract: contract}, ClosetFilterer: ClosetFilterer{contract: contract}}, nil
}

// Closet is an auto generated Go binding around an Ethereum contract.
type Closet struct {
	ClosetCaller     // Read-only binding to the contract
	ClosetTransactor // Write-only binding to the contract
	ClosetFilterer   // Log filterer for contract events
}

// ClosetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClosetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClosetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClosetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClosetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClosetSession struct {
	Contract     *Closet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClosetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClosetCallerSession struct {
	Contract *ClosetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClosetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClosetTransactorSession struct {
	Contract     *ClosetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClosetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClosetRaw struct {
	Contract *Closet // Generic contract binding to access the raw methods on
}

// ClosetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClosetCallerRaw struct {
	Contract *ClosetCaller // Generic read-only contract binding to access the raw methods on
}

// ClosetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClosetTransactorRaw struct {
	Contract *ClosetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloset creates a new instance of Closet, bound to a specific deployed contract.
func NewCloset(address common.Address, backend bind.ContractBackend) (*Closet, error) {
	contract, err := bindCloset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Closet{ClosetCaller: ClosetCaller{contract: contract}, ClosetTransactor: ClosetTransactor{contract: contract}, ClosetFilterer: ClosetFilterer{contract: contract}}, nil
}

// NewClosetCaller creates a new read-only instance of Closet, bound to a specific deployed contract.
func NewClosetCaller(address common.Address, caller bind.ContractCaller) (*ClosetCaller, error) {
	contract, err := bindCloset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClosetCaller{contract: contract}, nil
}

// NewClosetTransactor creates a new write-only instance of Closet, bound to a specific deployed contract.
func NewClosetTransactor(address common.Address, transactor bind.ContractTransactor) (*ClosetTransactor, error) {
	contract, err := bindCloset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClosetTransactor{contract: contract}, nil
}

// NewClosetFilterer creates a new log filterer instance of Closet, bound to a specific deployed contract.
func NewClosetFilterer(address common.Address, filterer bind.ContractFilterer) (*ClosetFilterer, error) {
	contract, err := bindCloset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClosetFilterer{contract: contract}, nil
}

// bindCloset binds a generic wrapper to an already deployed contract.
func bindCloset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClosetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Closet *ClosetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Closet.Contract.ClosetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Closet *ClosetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Closet.Contract.ClosetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Closet *ClosetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Closet.Contract.ClosetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Closet *ClosetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Closet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Closet *ClosetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Closet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Closet *ClosetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Closet.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Closet *ClosetCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Closet *ClosetSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Closet.Contract.BalanceOf(&_Closet.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Closet *ClosetCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Closet.Contract.BalanceOf(&_Closet.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Closet *ClosetCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Closet *ClosetSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Closet.Contract.BalanceOfBatch(&_Closet.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Closet *ClosetCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Closet.Contract.BalanceOfBatch(&_Closet.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Closet *ClosetCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Closet *ClosetSession) ContractURI() (string, error) {
	return _Closet.Contract.ContractURI(&_Closet.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Closet *ClosetCallerSession) ContractURI() (string, error) {
	return _Closet.Contract.ContractURI(&_Closet.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_Closet *ClosetCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_Closet *ClosetSession) Erc20() (common.Address, error) {
	return _Closet.Contract.Erc20(&_Closet.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_Closet *ClosetCallerSession) Erc20() (common.Address, error) {
	return _Closet.Contract.Erc20(&_Closet.CallOpts)
}

// ExistingTokenIds is a free data retrieval call binding the contract method 0xecf9cc24.
//
// Solidity: function existingTokenIds(uint256 ) view returns(uint256)
func (_Closet *ClosetCaller) ExistingTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "existingTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExistingTokenIds is a free data retrieval call binding the contract method 0xecf9cc24.
//
// Solidity: function existingTokenIds(uint256 ) view returns(uint256)
func (_Closet *ClosetSession) ExistingTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Closet.Contract.ExistingTokenIds(&_Closet.CallOpts, arg0)
}

// ExistingTokenIds is a free data retrieval call binding the contract method 0xecf9cc24.
//
// Solidity: function existingTokenIds(uint256 ) view returns(uint256)
func (_Closet *ClosetCallerSession) ExistingTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Closet.Contract.ExistingTokenIds(&_Closet.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Closet *ClosetCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Closet *ClosetSession) Exists(id *big.Int) (bool, error) {
	return _Closet.Contract.Exists(&_Closet.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Closet *ClosetCallerSession) Exists(id *big.Int) (bool, error) {
	return _Closet.Contract.Exists(&_Closet.CallOpts, id)
}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCaller) GetActiveTokens(opts *bind.CallOpts) ([]TokenWithId, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "getActiveTokens")

	if err != nil {
		return *new([]TokenWithId), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenWithId)).(*[]TokenWithId)

	return out0, err

}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetSession) GetActiveTokens() ([]TokenWithId, error) {
	return _Closet.Contract.GetActiveTokens(&_Closet.CallOpts)
}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCallerSession) GetActiveTokens() ([]TokenWithId, error) {
	return _Closet.Contract.GetActiveTokens(&_Closet.CallOpts)
}

// GetAllTokenIds is a free data retrieval call binding the contract method 0xbdbed722.
//
// Solidity: function getAllTokenIds() view returns(uint256[])
func (_Closet *ClosetCaller) GetAllTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "getAllTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllTokenIds is a free data retrieval call binding the contract method 0xbdbed722.
//
// Solidity: function getAllTokenIds() view returns(uint256[])
func (_Closet *ClosetSession) GetAllTokenIds() ([]*big.Int, error) {
	return _Closet.Contract.GetAllTokenIds(&_Closet.CallOpts)
}

// GetAllTokenIds is a free data retrieval call binding the contract method 0xbdbed722.
//
// Solidity: function getAllTokenIds() view returns(uint256[])
func (_Closet *ClosetCallerSession) GetAllTokenIds() ([]*big.Int, error) {
	return _Closet.Contract.GetAllTokenIds(&_Closet.CallOpts)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCaller) GetAllTokens(opts *bind.CallOpts) ([]TokenWithId, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "getAllTokens")

	if err != nil {
		return *new([]TokenWithId), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenWithId)).(*[]TokenWithId)

	return out0, err

}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetSession) GetAllTokens() ([]TokenWithId, error) {
	return _Closet.Contract.GetAllTokens(&_Closet.CallOpts)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCallerSession) GetAllTokens() ([]TokenWithId, error) {
	return _Closet.Contract.GetAllTokens(&_Closet.CallOpts)
}

// GetTokenById is a free data retrieval call binding the contract method 0x7bdc60d9.
//
// Solidity: function getTokenById(uint256 id) view returns((string,uint256,uint256,uint256,bool,address,uint256[2]))
func (_Closet *ClosetCaller) GetTokenById(opts *bind.CallOpts, id *big.Int) (Token, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "getTokenById", id)

	if err != nil {
		return *new(Token), err
	}

	out0 := *abi.ConvertType(out[0], new(Token)).(*Token)

	return out0, err

}

// GetTokenById is a free data retrieval call binding the contract method 0x7bdc60d9.
//
// Solidity: function getTokenById(uint256 id) view returns((string,uint256,uint256,uint256,bool,address,uint256[2]))
func (_Closet *ClosetSession) GetTokenById(id *big.Int) (Token, error) {
	return _Closet.Contract.GetTokenById(&_Closet.CallOpts, id)
}

// GetTokenById is a free data retrieval call binding the contract method 0x7bdc60d9.
//
// Solidity: function getTokenById(uint256 id) view returns((string,uint256,uint256,uint256,bool,address,uint256[2]))
func (_Closet *ClosetCallerSession) GetTokenById(id *big.Int) (Token, error) {
	return _Closet.Contract.GetTokenById(&_Closet.CallOpts, id)
}

// GetTokensByWallet is a free data retrieval call binding the contract method 0xd3b90b30.
//
// Solidity: function getTokensByWallet(address wallet) view returns((uint256,uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCaller) GetTokensByWallet(opts *bind.CallOpts, wallet common.Address) ([]UserToken, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "getTokensByWallet", wallet)

	if err != nil {
		return *new([]UserToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]UserToken)).(*[]UserToken)

	return out0, err

}

// GetTokensByWallet is a free data retrieval call binding the contract method 0xd3b90b30.
//
// Solidity: function getTokensByWallet(address wallet) view returns((uint256,uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetSession) GetTokensByWallet(wallet common.Address) ([]UserToken, error) {
	return _Closet.Contract.GetTokensByWallet(&_Closet.CallOpts, wallet)
}

// GetTokensByWallet is a free data retrieval call binding the contract method 0xd3b90b30.
//
// Solidity: function getTokensByWallet(address wallet) view returns((uint256,uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[])
func (_Closet *ClosetCallerSession) GetTokensByWallet(wallet common.Address) ([]UserToken, error) {
	return _Closet.Contract.GetTokensByWallet(&_Closet.CallOpts, wallet)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Closet *ClosetCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Closet *ClosetSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Closet.Contract.IsApprovedForAll(&_Closet.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Closet *ClosetCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Closet.Contract.IsApprovedForAll(&_Closet.CallOpts, account, operator)
}

// MaxTokensPerWalletById is a free data retrieval call binding the contract method 0xe046ed06.
//
// Solidity: function maxTokensPerWalletById(uint256 , address ) view returns(uint256)
func (_Closet *ClosetCaller) MaxTokensPerWalletById(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "maxTokensPerWalletById", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTokensPerWalletById is a free data retrieval call binding the contract method 0xe046ed06.
//
// Solidity: function maxTokensPerWalletById(uint256 , address ) view returns(uint256)
func (_Closet *ClosetSession) MaxTokensPerWalletById(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Closet.Contract.MaxTokensPerWalletById(&_Closet.CallOpts, arg0, arg1)
}

// MaxTokensPerWalletById is a free data retrieval call binding the contract method 0xe046ed06.
//
// Solidity: function maxTokensPerWalletById(uint256 , address ) view returns(uint256)
func (_Closet *ClosetCallerSession) MaxTokensPerWalletById(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Closet.Contract.MaxTokensPerWalletById(&_Closet.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Closet *ClosetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Closet *ClosetSession) Owner() (common.Address, error) {
	return _Closet.Contract.Owner(&_Closet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Closet *ClosetCallerSession) Owner() (common.Address, error) {
	return _Closet.Contract.Owner(&_Closet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Closet *ClosetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Closet *ClosetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Closet.Contract.SupportsInterface(&_Closet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Closet *ClosetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Closet.Contract.SupportsInterface(&_Closet.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Closet *ClosetCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Closet *ClosetSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Closet.Contract.TotalSupply(&_Closet.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Closet *ClosetCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Closet.Contract.TotalSupply(&_Closet.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Closet *ClosetCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Closet *ClosetSession) Uri(arg0 *big.Int) (string, error) {
	return _Closet.Contract.Uri(&_Closet.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Closet *ClosetCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Closet.Contract.Uri(&_Closet.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Closet *ClosetCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Closet *ClosetSession) Version() (string, error) {
	return _Closet.Contract.Version(&_Closet.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Closet *ClosetCallerSession) Version() (string, error) {
	return _Closet.Contract.Version(&_Closet.CallOpts)
}

// WalletBans is a free data retrieval call binding the contract method 0xeed670d7.
//
// Solidity: function walletBans(address ) view returns(bool banned, string reason)
func (_Closet *ClosetCaller) WalletBans(opts *bind.CallOpts, arg0 common.Address) (struct {
	Banned bool
	Reason string
}, error) {
	var out []interface{}
	err := _Closet.contract.Call(opts, &out, "walletBans", arg0)

	outstruct := new(struct {
		Banned bool
		Reason string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Banned = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Reason = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// WalletBans is a free data retrieval call binding the contract method 0xeed670d7.
//
// Solidity: function walletBans(address ) view returns(bool banned, string reason)
func (_Closet *ClosetSession) WalletBans(arg0 common.Address) (struct {
	Banned bool
	Reason string
}, error) {
	return _Closet.Contract.WalletBans(&_Closet.CallOpts, arg0)
}

// WalletBans is a free data retrieval call binding the contract method 0xeed670d7.
//
// Solidity: function walletBans(address ) view returns(bool banned, string reason)
func (_Closet *ClosetCallerSession) WalletBans(arg0 common.Address) (struct {
	Banned bool
	Reason string
}, error) {
	return _Closet.Contract.WalletBans(&_Closet.CallOpts, arg0)
}

// AddNewTokenTypes is a paid mutator transaction binding the contract method 0xf316534b.
//
// Solidity: function addNewTokenTypes((string,uint256,uint256,uint256,bool,address,uint256[2])[] tokens) returns()
func (_Closet *ClosetTransactor) AddNewTokenTypes(opts *bind.TransactOpts, tokens []Token) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "addNewTokenTypes", tokens)
}

// AddNewTokenTypes is a paid mutator transaction binding the contract method 0xf316534b.
//
// Solidity: function addNewTokenTypes((string,uint256,uint256,uint256,bool,address,uint256[2])[] tokens) returns()
func (_Closet *ClosetSession) AddNewTokenTypes(tokens []Token) (*types.Transaction, error) {
	return _Closet.Contract.AddNewTokenTypes(&_Closet.TransactOpts, tokens)
}

// AddNewTokenTypes is a paid mutator transaction binding the contract method 0xf316534b.
//
// Solidity: function addNewTokenTypes((string,uint256,uint256,uint256,bool,address,uint256[2])[] tokens) returns()
func (_Closet *ClosetTransactorSession) AddNewTokenTypes(tokens []Token) (*types.Transaction, error) {
	return _Closet.Contract.AddNewTokenTypes(&_Closet.TransactOpts, tokens)
}

// BanWallet is a paid mutator transaction binding the contract method 0x63d1a9ee.
//
// Solidity: function banWallet(address wallet, string reason) returns()
func (_Closet *ClosetTransactor) BanWallet(opts *bind.TransactOpts, wallet common.Address, reason string) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "banWallet", wallet, reason)
}

// BanWallet is a paid mutator transaction binding the contract method 0x63d1a9ee.
//
// Solidity: function banWallet(address wallet, string reason) returns()
func (_Closet *ClosetSession) BanWallet(wallet common.Address, reason string) (*types.Transaction, error) {
	return _Closet.Contract.BanWallet(&_Closet.TransactOpts, wallet, reason)
}

// BanWallet is a paid mutator transaction binding the contract method 0x63d1a9ee.
//
// Solidity: function banWallet(address wallet, string reason) returns()
func (_Closet *ClosetTransactorSession) BanWallet(wallet common.Address, reason string) (*types.Transaction, error) {
	return _Closet.Contract.BanWallet(&_Closet.TransactOpts, wallet, reason)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_Closet *ClosetTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "burn", tokenId, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_Closet *ClosetSession) Burn(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.Burn(&_Closet.TransactOpts, tokenId, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_Closet *ClosetTransactorSession) Burn(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.Burn(&_Closet.TransactOpts, tokenId, amount)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x83ca4b6f.
//
// Solidity: function burnBatch(uint256[] ids, uint256[] amounts) returns()
func (_Closet *ClosetTransactor) BurnBatch(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "burnBatch", ids, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x83ca4b6f.
//
// Solidity: function burnBatch(uint256[] ids, uint256[] amounts) returns()
func (_Closet *ClosetSession) BurnBatch(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.Contract.BurnBatch(&_Closet.TransactOpts, ids, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x83ca4b6f.
//
// Solidity: function burnBatch(uint256[] ids, uint256[] amounts) returns()
func (_Closet *ClosetTransactorSession) BurnBatch(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.Contract.BurnBatch(&_Closet.TransactOpts, ids, amounts)
}

// ChangeERC20Contract is a paid mutator transaction binding the contract method 0x38e33e43.
//
// Solidity: function changeERC20Contract(address erc20Addr) returns()
func (_Closet *ClosetTransactor) ChangeERC20Contract(opts *bind.TransactOpts, erc20Addr common.Address) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "changeERC20Contract", erc20Addr)
}

// ChangeERC20Contract is a paid mutator transaction binding the contract method 0x38e33e43.
//
// Solidity: function changeERC20Contract(address erc20Addr) returns()
func (_Closet *ClosetSession) ChangeERC20Contract(erc20Addr common.Address) (*types.Transaction, error) {
	return _Closet.Contract.ChangeERC20Contract(&_Closet.TransactOpts, erc20Addr)
}

// ChangeERC20Contract is a paid mutator transaction binding the contract method 0x38e33e43.
//
// Solidity: function changeERC20Contract(address erc20Addr) returns()
func (_Closet *ClosetTransactorSession) ChangeERC20Contract(erc20Addr common.Address) (*types.Transaction, error) {
	return _Closet.Contract.ChangeERC20Contract(&_Closet.TransactOpts, erc20Addr)
}

// ChangeTokens is a paid mutator transaction binding the contract method 0x138d3f09.
//
// Solidity: function changeTokens((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[] tokens) returns()
func (_Closet *ClosetTransactor) ChangeTokens(opts *bind.TransactOpts, tokens []TokenWithId) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "changeTokens", tokens)
}

// ChangeTokens is a paid mutator transaction binding the contract method 0x138d3f09.
//
// Solidity: function changeTokens((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[] tokens) returns()
func (_Closet *ClosetSession) ChangeTokens(tokens []TokenWithId) (*types.Transaction, error) {
	return _Closet.Contract.ChangeTokens(&_Closet.TransactOpts, tokens)
}

// ChangeTokens is a paid mutator transaction binding the contract method 0x138d3f09.
//
// Solidity: function changeTokens((uint256,(string,uint256,uint256,uint256,bool,address,uint256[2]))[] tokens) returns()
func (_Closet *ClosetTransactorSession) ChangeTokens(tokens []TokenWithId) (*types.Transaction, error) {
	return _Closet.Contract.ChangeTokens(&_Closet.TransactOpts, tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Closet *ClosetTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Closet *ClosetSession) Initialize() (*types.Transaction, error) {
	return _Closet.Contract.Initialize(&_Closet.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Closet *ClosetTransactorSession) Initialize() (*types.Transaction, error) {
	return _Closet.Contract.Initialize(&_Closet.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Closet *ClosetTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "mint", tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Closet *ClosetSession) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.Mint(&_Closet.TransactOpts, tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Closet *ClosetTransactorSession) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.Mint(&_Closet.TransactOpts, tokenId, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts) payable returns()
func (_Closet *ClosetTransactor) MintBatch(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "mintBatch", ids, amounts)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts) payable returns()
func (_Closet *ClosetSession) MintBatch(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.Contract.MintBatch(&_Closet.TransactOpts, ids, amounts)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts) payable returns()
func (_Closet *ClosetTransactorSession) MintBatch(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Closet.Contract.MintBatch(&_Closet.TransactOpts, ids, amounts)
}

// PromoMint is a paid mutator transaction binding the contract method 0xf52d9e0f.
//
// Solidity: function promoMint(uint256[] ids, uint256[] amounts, address wallet) returns()
func (_Closet *ClosetTransactor) PromoMint(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int, wallet common.Address) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "promoMint", ids, amounts, wallet)
}

// PromoMint is a paid mutator transaction binding the contract method 0xf52d9e0f.
//
// Solidity: function promoMint(uint256[] ids, uint256[] amounts, address wallet) returns()
func (_Closet *ClosetSession) PromoMint(ids []*big.Int, amounts []*big.Int, wallet common.Address) (*types.Transaction, error) {
	return _Closet.Contract.PromoMint(&_Closet.TransactOpts, ids, amounts, wallet)
}

// PromoMint is a paid mutator transaction binding the contract method 0xf52d9e0f.
//
// Solidity: function promoMint(uint256[] ids, uint256[] amounts, address wallet) returns()
func (_Closet *ClosetTransactorSession) PromoMint(ids []*big.Int, amounts []*big.Int, wallet common.Address) (*types.Transaction, error) {
	return _Closet.Contract.PromoMint(&_Closet.TransactOpts, ids, amounts, wallet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Closet *ClosetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Closet *ClosetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Closet.Contract.RenounceOwnership(&_Closet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Closet *ClosetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Closet.Contract.RenounceOwnership(&_Closet.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Closet *ClosetTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Closet *ClosetSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.SafeBatchTransferFrom(&_Closet.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Closet *ClosetTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.SafeBatchTransferFrom(&_Closet.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Closet *ClosetTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Closet *ClosetSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.SafeTransferFrom(&_Closet.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Closet *ClosetTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.SafeTransferFrom(&_Closet.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Closet *ClosetTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Closet *ClosetSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Closet.Contract.SetApprovalForAll(&_Closet.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Closet *ClosetTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Closet.Contract.SetApprovalForAll(&_Closet.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Closet *ClosetTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Closet *ClosetSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Closet.Contract.SetContractURI(&_Closet.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Closet *ClosetTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Closet.Contract.SetContractURI(&_Closet.TransactOpts, newContractURI)
}

// SetMaxTokensForWallet is a paid mutator transaction binding the contract method 0xf1e54b23.
//
// Solidity: function setMaxTokensForWallet(address wallet, uint256 tokenId, uint256 max) returns()
func (_Closet *ClosetTransactor) SetMaxTokensForWallet(opts *bind.TransactOpts, wallet common.Address, tokenId *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "setMaxTokensForWallet", wallet, tokenId, max)
}

// SetMaxTokensForWallet is a paid mutator transaction binding the contract method 0xf1e54b23.
//
// Solidity: function setMaxTokensForWallet(address wallet, uint256 tokenId, uint256 max) returns()
func (_Closet *ClosetSession) SetMaxTokensForWallet(wallet common.Address, tokenId *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.SetMaxTokensForWallet(&_Closet.TransactOpts, wallet, tokenId, max)
}

// SetMaxTokensForWallet is a paid mutator transaction binding the contract method 0xf1e54b23.
//
// Solidity: function setMaxTokensForWallet(address wallet, uint256 tokenId, uint256 max) returns()
func (_Closet *ClosetTransactorSession) SetMaxTokensForWallet(wallet common.Address, tokenId *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Closet.Contract.SetMaxTokensForWallet(&_Closet.TransactOpts, wallet, tokenId, max)
}

// SetTokensStatus is a paid mutator transaction binding the contract method 0x736ce856.
//
// Solidity: function setTokensStatus(uint256[] ids, bool status) returns()
func (_Closet *ClosetTransactor) SetTokensStatus(opts *bind.TransactOpts, ids []*big.Int, status bool) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "setTokensStatus", ids, status)
}

// SetTokensStatus is a paid mutator transaction binding the contract method 0x736ce856.
//
// Solidity: function setTokensStatus(uint256[] ids, bool status) returns()
func (_Closet *ClosetSession) SetTokensStatus(ids []*big.Int, status bool) (*types.Transaction, error) {
	return _Closet.Contract.SetTokensStatus(&_Closet.TransactOpts, ids, status)
}

// SetTokensStatus is a paid mutator transaction binding the contract method 0x736ce856.
//
// Solidity: function setTokensStatus(uint256[] ids, bool status) returns()
func (_Closet *ClosetTransactorSession) SetTokensStatus(ids []*big.Int, status bool) (*types.Transaction, error) {
	return _Closet.Contract.SetTokensStatus(&_Closet.TransactOpts, ids, status)
}

// SetUri is a paid mutator transaction binding the contract method 0x9b642de1.
//
// Solidity: function setUri(string uri) returns()
func (_Closet *ClosetTransactor) SetUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "setUri", uri)
}

// SetUri is a paid mutator transaction binding the contract method 0x9b642de1.
//
// Solidity: function setUri(string uri) returns()
func (_Closet *ClosetSession) SetUri(uri string) (*types.Transaction, error) {
	return _Closet.Contract.SetUri(&_Closet.TransactOpts, uri)
}

// SetUri is a paid mutator transaction binding the contract method 0x9b642de1.
//
// Solidity: function setUri(string uri) returns()
func (_Closet *ClosetTransactorSession) SetUri(uri string) (*types.Transaction, error) {
	return _Closet.Contract.SetUri(&_Closet.TransactOpts, uri)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Closet *ClosetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Closet *ClosetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Closet.Contract.TransferOwnership(&_Closet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Closet *ClosetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Closet.Contract.TransferOwnership(&_Closet.TransactOpts, newOwner)
}

// UnbanWallet is a paid mutator transaction binding the contract method 0x6a4aef02.
//
// Solidity: function unbanWallet(address wallet) returns()
func (_Closet *ClosetTransactor) UnbanWallet(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "unbanWallet", wallet)
}

// UnbanWallet is a paid mutator transaction binding the contract method 0x6a4aef02.
//
// Solidity: function unbanWallet(address wallet) returns()
func (_Closet *ClosetSession) UnbanWallet(wallet common.Address) (*types.Transaction, error) {
	return _Closet.Contract.UnbanWallet(&_Closet.TransactOpts, wallet)
}

// UnbanWallet is a paid mutator transaction binding the contract method 0x6a4aef02.
//
// Solidity: function unbanWallet(address wallet) returns()
func (_Closet *ClosetTransactorSession) UnbanWallet(wallet common.Address) (*types.Transaction, error) {
	return _Closet.Contract.UnbanWallet(&_Closet.TransactOpts, wallet)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Closet *ClosetTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Closet *ClosetSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Closet.Contract.UpgradeTo(&_Closet.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Closet *ClosetTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Closet.Contract.UpgradeTo(&_Closet.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Closet *ClosetTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Closet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Closet *ClosetSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.UpgradeToAndCall(&_Closet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Closet *ClosetTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Closet.Contract.UpgradeToAndCall(&_Closet.TransactOpts, newImplementation, data)
}

// ClosetAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Closet contract.
type ClosetAdminChangedIterator struct {
	Event *ClosetAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetAdminChanged represents a AdminChanged event raised by the Closet contract.
type ClosetAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Closet *ClosetFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ClosetAdminChangedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ClosetAdminChangedIterator{contract: _Closet.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Closet *ClosetFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ClosetAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetAdminChanged)
				if err := _Closet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Closet *ClosetFilterer) ParseAdminChanged(log types.Log) (*ClosetAdminChanged, error) {
	event := new(ClosetAdminChanged)
	if err := _Closet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Closet contract.
type ClosetApprovalForAllIterator struct {
	Event *ClosetApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetApprovalForAll represents a ApprovalForAll event raised by the Closet contract.
type ClosetApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Closet *ClosetFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ClosetApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ClosetApprovalForAllIterator{contract: _Closet.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Closet *ClosetFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ClosetApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetApprovalForAll)
				if err := _Closet.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Closet *ClosetFilterer) ParseApprovalForAll(log types.Log) (*ClosetApprovalForAll, error) {
	event := new(ClosetApprovalForAll)
	if err := _Closet.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetBatchTokensBurnedIterator is returned from FilterBatchTokensBurned and is used to iterate over the raw logs and unpacked data for BatchTokensBurned events raised by the Closet contract.
type ClosetBatchTokensBurnedIterator struct {
	Event *ClosetBatchTokensBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetBatchTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetBatchTokensBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetBatchTokensBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetBatchTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetBatchTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetBatchTokensBurned represents a BatchTokensBurned event raised by the Closet contract.
type ClosetBatchTokensBurned struct {
	TokenIds []*big.Int
	Amounts  []*big.Int
	Wallet   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchTokensBurned is a free log retrieval operation binding the contract event 0x3c53c7744097cc8b22f1268b0b5c8b8c8233acf4f6249b7de76d0d5cbf349c19.
//
// Solidity: event BatchTokensBurned(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) FilterBatchTokensBurned(opts *bind.FilterOpts) (*ClosetBatchTokensBurnedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "BatchTokensBurned")
	if err != nil {
		return nil, err
	}
	return &ClosetBatchTokensBurnedIterator{contract: _Closet.contract, event: "BatchTokensBurned", logs: logs, sub: sub}, nil
}

// WatchBatchTokensBurned is a free log subscription operation binding the contract event 0x3c53c7744097cc8b22f1268b0b5c8b8c8233acf4f6249b7de76d0d5cbf349c19.
//
// Solidity: event BatchTokensBurned(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) WatchBatchTokensBurned(opts *bind.WatchOpts, sink chan<- *ClosetBatchTokensBurned) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "BatchTokensBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetBatchTokensBurned)
				if err := _Closet.contract.UnpackLog(event, "BatchTokensBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTokensBurned is a log parse operation binding the contract event 0x3c53c7744097cc8b22f1268b0b5c8b8c8233acf4f6249b7de76d0d5cbf349c19.
//
// Solidity: event BatchTokensBurned(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) ParseBatchTokensBurned(log types.Log) (*ClosetBatchTokensBurned, error) {
	event := new(ClosetBatchTokensBurned)
	if err := _Closet.contract.UnpackLog(event, "BatchTokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetBatchTokensMintedIterator is returned from FilterBatchTokensMinted and is used to iterate over the raw logs and unpacked data for BatchTokensMinted events raised by the Closet contract.
type ClosetBatchTokensMintedIterator struct {
	Event *ClosetBatchTokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetBatchTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetBatchTokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetBatchTokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetBatchTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetBatchTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetBatchTokensMinted represents a BatchTokensMinted event raised by the Closet contract.
type ClosetBatchTokensMinted struct {
	TokenIds []*big.Int
	Amounts  []*big.Int
	Wallet   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchTokensMinted is a free log retrieval operation binding the contract event 0xddc7a7105147d9dc422856b8dac226d6e644f688cf9ea1ea10c67cbc402dc38c.
//
// Solidity: event BatchTokensMinted(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) FilterBatchTokensMinted(opts *bind.FilterOpts) (*ClosetBatchTokensMintedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "BatchTokensMinted")
	if err != nil {
		return nil, err
	}
	return &ClosetBatchTokensMintedIterator{contract: _Closet.contract, event: "BatchTokensMinted", logs: logs, sub: sub}, nil
}

// WatchBatchTokensMinted is a free log subscription operation binding the contract event 0xddc7a7105147d9dc422856b8dac226d6e644f688cf9ea1ea10c67cbc402dc38c.
//
// Solidity: event BatchTokensMinted(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) WatchBatchTokensMinted(opts *bind.WatchOpts, sink chan<- *ClosetBatchTokensMinted) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "BatchTokensMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetBatchTokensMinted)
				if err := _Closet.contract.UnpackLog(event, "BatchTokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTokensMinted is a log parse operation binding the contract event 0xddc7a7105147d9dc422856b8dac226d6e644f688cf9ea1ea10c67cbc402dc38c.
//
// Solidity: event BatchTokensMinted(uint256[] tokenIds, uint256[] amounts, address wallet)
func (_Closet *ClosetFilterer) ParseBatchTokensMinted(log types.Log) (*ClosetBatchTokensMinted, error) {
	event := new(ClosetBatchTokensMinted)
	if err := _Closet.contract.UnpackLog(event, "BatchTokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Closet contract.
type ClosetBeaconUpgradedIterator struct {
	Event *ClosetBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetBeaconUpgraded represents a BeaconUpgraded event raised by the Closet contract.
type ClosetBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Closet *ClosetFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ClosetBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ClosetBeaconUpgradedIterator{contract: _Closet.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Closet *ClosetFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ClosetBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetBeaconUpgraded)
				if err := _Closet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Closet *ClosetFilterer) ParseBeaconUpgraded(log types.Log) (*ClosetBeaconUpgraded, error) {
	event := new(ClosetBeaconUpgraded)
	if err := _Closet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetChangeERC20ContractIterator is returned from FilterChangeERC20Contract and is used to iterate over the raw logs and unpacked data for ChangeERC20Contract events raised by the Closet contract.
type ClosetChangeERC20ContractIterator struct {
	Event *ClosetChangeERC20Contract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetChangeERC20ContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetChangeERC20Contract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetChangeERC20Contract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetChangeERC20ContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetChangeERC20ContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetChangeERC20Contract represents a ChangeERC20Contract event raised by the Closet contract.
type ClosetChangeERC20Contract struct {
	Erc20Addr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeERC20Contract is a free log retrieval operation binding the contract event 0xf4437fae6d1fbef7b2436233184e29ceee39249efd21c29c42991b93beb52be4.
//
// Solidity: event ChangeERC20Contract(address erc20Addr)
func (_Closet *ClosetFilterer) FilterChangeERC20Contract(opts *bind.FilterOpts) (*ClosetChangeERC20ContractIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "ChangeERC20Contract")
	if err != nil {
		return nil, err
	}
	return &ClosetChangeERC20ContractIterator{contract: _Closet.contract, event: "ChangeERC20Contract", logs: logs, sub: sub}, nil
}

// WatchChangeERC20Contract is a free log subscription operation binding the contract event 0xf4437fae6d1fbef7b2436233184e29ceee39249efd21c29c42991b93beb52be4.
//
// Solidity: event ChangeERC20Contract(address erc20Addr)
func (_Closet *ClosetFilterer) WatchChangeERC20Contract(opts *bind.WatchOpts, sink chan<- *ClosetChangeERC20Contract) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "ChangeERC20Contract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetChangeERC20Contract)
				if err := _Closet.contract.UnpackLog(event, "ChangeERC20Contract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangeERC20Contract is a log parse operation binding the contract event 0xf4437fae6d1fbef7b2436233184e29ceee39249efd21c29c42991b93beb52be4.
//
// Solidity: event ChangeERC20Contract(address erc20Addr)
func (_Closet *ClosetFilterer) ParseChangeERC20Contract(log types.Log) (*ClosetChangeERC20Contract, error) {
	event := new(ClosetChangeERC20Contract)
	if err := _Closet.contract.UnpackLog(event, "ChangeERC20Contract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Closet contract.
type ClosetOwnershipTransferredIterator struct {
	Event *ClosetOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetOwnershipTransferred represents a OwnershipTransferred event raised by the Closet contract.
type ClosetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Closet *ClosetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClosetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClosetOwnershipTransferredIterator{contract: _Closet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Closet *ClosetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClosetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetOwnershipTransferred)
				if err := _Closet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Closet *ClosetFilterer) ParseOwnershipTransferred(log types.Log) (*ClosetOwnershipTransferred, error) {
	event := new(ClosetOwnershipTransferred)
	if err := _Closet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTokenTypeAddedIterator is returned from FilterTokenTypeAdded and is used to iterate over the raw logs and unpacked data for TokenTypeAdded events raised by the Closet contract.
type ClosetTokenTypeAddedIterator struct {
	Event *ClosetTokenTypeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTokenTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTokenTypeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTokenTypeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTokenTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTokenTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTokenTypeAdded represents a TokenTypeAdded event raised by the Closet contract.
type ClosetTokenTypeAdded struct {
	TokenId *big.Int
	Token   Token
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenTypeAdded is a free log retrieval operation binding the contract event 0xcbd47bdaacef607c448cfd0079c791bcba6a3d01ac2999cdbd4afec8c91c0838.
//
// Solidity: event TokenTypeAdded(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) FilterTokenTypeAdded(opts *bind.FilterOpts) (*ClosetTokenTypeAddedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TokenTypeAdded")
	if err != nil {
		return nil, err
	}
	return &ClosetTokenTypeAddedIterator{contract: _Closet.contract, event: "TokenTypeAdded", logs: logs, sub: sub}, nil
}

// WatchTokenTypeAdded is a free log subscription operation binding the contract event 0xcbd47bdaacef607c448cfd0079c791bcba6a3d01ac2999cdbd4afec8c91c0838.
//
// Solidity: event TokenTypeAdded(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) WatchTokenTypeAdded(opts *bind.WatchOpts, sink chan<- *ClosetTokenTypeAdded) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TokenTypeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTokenTypeAdded)
				if err := _Closet.contract.UnpackLog(event, "TokenTypeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenTypeAdded is a log parse operation binding the contract event 0xcbd47bdaacef607c448cfd0079c791bcba6a3d01ac2999cdbd4afec8c91c0838.
//
// Solidity: event TokenTypeAdded(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) ParseTokenTypeAdded(log types.Log) (*ClosetTokenTypeAdded, error) {
	event := new(ClosetTokenTypeAdded)
	if err := _Closet.contract.UnpackLog(event, "TokenTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTokenTypeChangedIterator is returned from FilterTokenTypeChanged and is used to iterate over the raw logs and unpacked data for TokenTypeChanged events raised by the Closet contract.
type ClosetTokenTypeChangedIterator struct {
	Event *ClosetTokenTypeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTokenTypeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTokenTypeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTokenTypeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTokenTypeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTokenTypeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTokenTypeChanged represents a TokenTypeChanged event raised by the Closet contract.
type ClosetTokenTypeChanged struct {
	TokenId *big.Int
	Token   Token
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenTypeChanged is a free log retrieval operation binding the contract event 0x59a0de9e767f7ea1fa2a7afb6ba676db02d66e9e7781d76a8e83d0cc07879435.
//
// Solidity: event TokenTypeChanged(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) FilterTokenTypeChanged(opts *bind.FilterOpts) (*ClosetTokenTypeChangedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TokenTypeChanged")
	if err != nil {
		return nil, err
	}
	return &ClosetTokenTypeChangedIterator{contract: _Closet.contract, event: "TokenTypeChanged", logs: logs, sub: sub}, nil
}

// WatchTokenTypeChanged is a free log subscription operation binding the contract event 0x59a0de9e767f7ea1fa2a7afb6ba676db02d66e9e7781d76a8e83d0cc07879435.
//
// Solidity: event TokenTypeChanged(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) WatchTokenTypeChanged(opts *bind.WatchOpts, sink chan<- *ClosetTokenTypeChanged) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TokenTypeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTokenTypeChanged)
				if err := _Closet.contract.UnpackLog(event, "TokenTypeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenTypeChanged is a log parse operation binding the contract event 0x59a0de9e767f7ea1fa2a7afb6ba676db02d66e9e7781d76a8e83d0cc07879435.
//
// Solidity: event TokenTypeChanged(uint256 tokenId, (string,uint256,uint256,uint256,bool,address,uint256[2]) token)
func (_Closet *ClosetFilterer) ParseTokenTypeChanged(log types.Log) (*ClosetTokenTypeChanged, error) {
	event := new(ClosetTokenTypeChanged)
	if err := _Closet.contract.UnpackLog(event, "TokenTypeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTokensBurnedIterator is returned from FilterTokensBurned and is used to iterate over the raw logs and unpacked data for TokensBurned events raised by the Closet contract.
type ClosetTokensBurnedIterator struct {
	Event *ClosetTokensBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTokensBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTokensBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTokensBurned represents a TokensBurned event raised by the Closet contract.
type ClosetTokensBurned struct {
	TokenId *big.Int
	Amount  *big.Int
	Wallet  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensBurned is a free log retrieval operation binding the contract event 0x8b70448d0d143d815676e1088155bf2ff4436b50d76b30ff3a1fe29290892702.
//
// Solidity: event TokensBurned(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) FilterTokensBurned(opts *bind.FilterOpts) (*ClosetTokensBurnedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TokensBurned")
	if err != nil {
		return nil, err
	}
	return &ClosetTokensBurnedIterator{contract: _Closet.contract, event: "TokensBurned", logs: logs, sub: sub}, nil
}

// WatchTokensBurned is a free log subscription operation binding the contract event 0x8b70448d0d143d815676e1088155bf2ff4436b50d76b30ff3a1fe29290892702.
//
// Solidity: event TokensBurned(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) WatchTokensBurned(opts *bind.WatchOpts, sink chan<- *ClosetTokensBurned) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TokensBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTokensBurned)
				if err := _Closet.contract.UnpackLog(event, "TokensBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensBurned is a log parse operation binding the contract event 0x8b70448d0d143d815676e1088155bf2ff4436b50d76b30ff3a1fe29290892702.
//
// Solidity: event TokensBurned(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) ParseTokensBurned(log types.Log) (*ClosetTokensBurned, error) {
	event := new(ClosetTokensBurned)
	if err := _Closet.contract.UnpackLog(event, "TokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the Closet contract.
type ClosetTokensMintedIterator struct {
	Event *ClosetTokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTokensMinted represents a TokensMinted event raised by the Closet contract.
type ClosetTokensMinted struct {
	TokenId *big.Int
	Amount  *big.Int
	Wallet  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x06437edc6222bc3b01fd946ea7b36242878aef9a533aa8815121bb23dc0a7438.
//
// Solidity: event TokensMinted(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) FilterTokensMinted(opts *bind.FilterOpts) (*ClosetTokensMintedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TokensMinted")
	if err != nil {
		return nil, err
	}
	return &ClosetTokensMintedIterator{contract: _Closet.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x06437edc6222bc3b01fd946ea7b36242878aef9a533aa8815121bb23dc0a7438.
//
// Solidity: event TokensMinted(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *ClosetTokensMinted) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TokensMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTokensMinted)
				if err := _Closet.contract.UnpackLog(event, "TokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMinted is a log parse operation binding the contract event 0x06437edc6222bc3b01fd946ea7b36242878aef9a533aa8815121bb23dc0a7438.
//
// Solidity: event TokensMinted(uint256 tokenId, uint256 amount, address wallet)
func (_Closet *ClosetFilterer) ParseTokensMinted(log types.Log) (*ClosetTokensMinted, error) {
	event := new(ClosetTokensMinted)
	if err := _Closet.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Closet contract.
type ClosetTransferBatchIterator struct {
	Event *ClosetTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTransferBatch represents a TransferBatch event raised by the Closet contract.
type ClosetTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Closet *ClosetFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ClosetTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClosetTransferBatchIterator{contract: _Closet.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Closet *ClosetFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ClosetTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTransferBatch)
				if err := _Closet.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Closet *ClosetFilterer) ParseTransferBatch(log types.Log) (*ClosetTransferBatch, error) {
	event := new(ClosetTransferBatch)
	if err := _Closet.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Closet contract.
type ClosetTransferSingleIterator struct {
	Event *ClosetTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetTransferSingle represents a TransferSingle event raised by the Closet contract.
type ClosetTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Closet *ClosetFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ClosetTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClosetTransferSingleIterator{contract: _Closet.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Closet *ClosetFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ClosetTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetTransferSingle)
				if err := _Closet.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Closet *ClosetFilterer) ParseTransferSingle(log types.Log) (*ClosetTransferSingle, error) {
	event := new(ClosetTransferSingle)
	if err := _Closet.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Closet contract.
type ClosetURIIterator struct {
	Event *ClosetURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetURI represents a URI event raised by the Closet contract.
type ClosetURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Closet *ClosetFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ClosetURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ClosetURIIterator{contract: _Closet.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Closet *ClosetFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ClosetURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetURI)
				if err := _Closet.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Closet *ClosetFilterer) ParseURI(log types.Log) (*ClosetURI, error) {
	event := new(ClosetURI)
	if err := _Closet.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Closet contract.
type ClosetUpgradedIterator struct {
	Event *ClosetUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetUpgraded represents a Upgraded event raised by the Closet contract.
type ClosetUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Closet *ClosetFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ClosetUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Closet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ClosetUpgradedIterator{contract: _Closet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Closet *ClosetFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ClosetUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Closet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetUpgraded)
				if err := _Closet.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Closet *ClosetFilterer) ParseUpgraded(log types.Log) (*ClosetUpgraded, error) {
	event := new(ClosetUpgraded)
	if err := _Closet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetWalletBannedIterator is returned from FilterWalletBanned and is used to iterate over the raw logs and unpacked data for WalletBanned events raised by the Closet contract.
type ClosetWalletBannedIterator struct {
	Event *ClosetWalletBanned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetWalletBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetWalletBanned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetWalletBanned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetWalletBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetWalletBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetWalletBanned represents a WalletBanned event raised by the Closet contract.
type ClosetWalletBanned struct {
	Wallet common.Address
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletBanned is a free log retrieval operation binding the contract event 0xc644f65bce175b55083b5f7cb20c238f0ee0ead74eeef083e6874ec3db520d71.
//
// Solidity: event WalletBanned(address wallet, string reason)
func (_Closet *ClosetFilterer) FilterWalletBanned(opts *bind.FilterOpts) (*ClosetWalletBannedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "WalletBanned")
	if err != nil {
		return nil, err
	}
	return &ClosetWalletBannedIterator{contract: _Closet.contract, event: "WalletBanned", logs: logs, sub: sub}, nil
}

// WatchWalletBanned is a free log subscription operation binding the contract event 0xc644f65bce175b55083b5f7cb20c238f0ee0ead74eeef083e6874ec3db520d71.
//
// Solidity: event WalletBanned(address wallet, string reason)
func (_Closet *ClosetFilterer) WatchWalletBanned(opts *bind.WatchOpts, sink chan<- *ClosetWalletBanned) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "WalletBanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetWalletBanned)
				if err := _Closet.contract.UnpackLog(event, "WalletBanned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWalletBanned is a log parse operation binding the contract event 0xc644f65bce175b55083b5f7cb20c238f0ee0ead74eeef083e6874ec3db520d71.
//
// Solidity: event WalletBanned(address wallet, string reason)
func (_Closet *ClosetFilterer) ParseWalletBanned(log types.Log) (*ClosetWalletBanned, error) {
	event := new(ClosetWalletBanned)
	if err := _Closet.contract.UnpackLog(event, "WalletBanned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetWalletMaxChangedIterator is returned from FilterWalletMaxChanged and is used to iterate over the raw logs and unpacked data for WalletMaxChanged events raised by the Closet contract.
type ClosetWalletMaxChangedIterator struct {
	Event *ClosetWalletMaxChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetWalletMaxChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetWalletMaxChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetWalletMaxChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetWalletMaxChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetWalletMaxChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetWalletMaxChanged represents a WalletMaxChanged event raised by the Closet contract.
type ClosetWalletMaxChanged struct {
	Wallet  common.Address
	TokenId *big.Int
	Max     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWalletMaxChanged is a free log retrieval operation binding the contract event 0x6b30946421f69c81ed9a9f3334ecf2a5ed5b0bcc149d52c2900fa0ce5f930760.
//
// Solidity: event WalletMaxChanged(address wallet, uint256 tokenId, uint256 max)
func (_Closet *ClosetFilterer) FilterWalletMaxChanged(opts *bind.FilterOpts) (*ClosetWalletMaxChangedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "WalletMaxChanged")
	if err != nil {
		return nil, err
	}
	return &ClosetWalletMaxChangedIterator{contract: _Closet.contract, event: "WalletMaxChanged", logs: logs, sub: sub}, nil
}

// WatchWalletMaxChanged is a free log subscription operation binding the contract event 0x6b30946421f69c81ed9a9f3334ecf2a5ed5b0bcc149d52c2900fa0ce5f930760.
//
// Solidity: event WalletMaxChanged(address wallet, uint256 tokenId, uint256 max)
func (_Closet *ClosetFilterer) WatchWalletMaxChanged(opts *bind.WatchOpts, sink chan<- *ClosetWalletMaxChanged) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "WalletMaxChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetWalletMaxChanged)
				if err := _Closet.contract.UnpackLog(event, "WalletMaxChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWalletMaxChanged is a log parse operation binding the contract event 0x6b30946421f69c81ed9a9f3334ecf2a5ed5b0bcc149d52c2900fa0ce5f930760.
//
// Solidity: event WalletMaxChanged(address wallet, uint256 tokenId, uint256 max)
func (_Closet *ClosetFilterer) ParseWalletMaxChanged(log types.Log) (*ClosetWalletMaxChanged, error) {
	event := new(ClosetWalletMaxChanged)
	if err := _Closet.contract.UnpackLog(event, "WalletMaxChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClosetWalletUnbannedIterator is returned from FilterWalletUnbanned and is used to iterate over the raw logs and unpacked data for WalletUnbanned events raised by the Closet contract.
type ClosetWalletUnbannedIterator struct {
	Event *ClosetWalletUnbanned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClosetWalletUnbannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClosetWalletUnbanned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClosetWalletUnbanned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClosetWalletUnbannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClosetWalletUnbannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClosetWalletUnbanned represents a WalletUnbanned event raised by the Closet contract.
type ClosetWalletUnbanned struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletUnbanned is a free log retrieval operation binding the contract event 0xb2ed1ab64220fe50bc8b9058c50791a59bf4e36b5fd723d425c857b0a468fe61.
//
// Solidity: event WalletUnbanned(address wallet)
func (_Closet *ClosetFilterer) FilterWalletUnbanned(opts *bind.FilterOpts) (*ClosetWalletUnbannedIterator, error) {

	logs, sub, err := _Closet.contract.FilterLogs(opts, "WalletUnbanned")
	if err != nil {
		return nil, err
	}
	return &ClosetWalletUnbannedIterator{contract: _Closet.contract, event: "WalletUnbanned", logs: logs, sub: sub}, nil
}

// WatchWalletUnbanned is a free log subscription operation binding the contract event 0xb2ed1ab64220fe50bc8b9058c50791a59bf4e36b5fd723d425c857b0a468fe61.
//
// Solidity: event WalletUnbanned(address wallet)
func (_Closet *ClosetFilterer) WatchWalletUnbanned(opts *bind.WatchOpts, sink chan<- *ClosetWalletUnbanned) (event.Subscription, error) {

	logs, sub, err := _Closet.contract.WatchLogs(opts, "WalletUnbanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClosetWalletUnbanned)
				if err := _Closet.contract.UnpackLog(event, "WalletUnbanned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWalletUnbanned is a log parse operation binding the contract event 0xb2ed1ab64220fe50bc8b9058c50791a59bf4e36b5fd723d425c857b0a468fe61.
//
// Solidity: event WalletUnbanned(address wallet)
func (_Closet *ClosetFilterer) ParseWalletUnbanned(log types.Log) (*ClosetWalletUnbanned, error) {
	event := new(ClosetWalletUnbanned)
	if err := _Closet.contract.UnpackLog(event, "WalletUnbanned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
