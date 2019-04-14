package nfts

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token definition
type NFT struct {
	Owner       sdk.AccAddress `json:"owner"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	TokenURI    string         `json:"token_uri"`
}

// NewNFT creates a new NFT
func NewNFT(owner sdk.AccAddress, tokenURI, description, image, name string,
) NFT {
	return NFT{
		Owner:       owner,
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Image:       strings.TrimSpace(image),
		TokenURI:    strings.TrimSpace(tokenURI),
	}
}

// EditMetadata edits metadata of an nft
func (nft NFT) EditMetadata(name, description, image, tokenURI string) NFT {
	nft.Name = name
	nft.Description = description
	nft.Image = image
	nft.TokenURI = tokenURI
	return nft
}

// Denom is a string
type Denom string

// TokenID is a uint64
type TokenID uint64

// Owner of non fungible tokens
type Owner map[Denom][]TokenID

// NewOwner returns a new empty owner
func NewOwner() Owner {
	return map[Denom][]TokenID{}
}

// Collection of non fungible tokens
type Collection map[TokenID]NFT

// NewCollection creates a new NFT Collection
func NewCollection() Collection {
	return make(map[TokenID]NFT)
}

// GetNFT gets a NFT from the collection
func (collection Collection) GetNFT(denom Denom, id TokenID) (nft NFT, err error) {
	nft, ok := collection[id]
	if !ok {
		return nft, fmt.Errorf("collection %s doesn't contain an NFT with TokenID %d", denom, id)
	}
	return
}

// ValidateBasic validates a Collection
func (collection Collection) ValidateBasic() (err error) {
	return
}
