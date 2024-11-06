package token_metadata

import (
	"github.com/gagliardetto/solana-go"
	
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
)

type Protocol uint

const (
	PROTOCOL_LEGACY Protocol = iota
	PROTOCOL_TOKEN22_MINT
	PROTOCOL_LIBREPLEX
	PROTOCOL_METAPLEX_CORE
)

type Metadata struct {
	Key             		token_metadata.Key
	UpdateAuthority 		solana.PublicKey
	Mint            		solana.PublicKey
	Data            		Data
	PrimarySaleHappened 	bool
	IsMutable 				bool
	Collection 				*token_metadata.Collection	`bin:"optional"`
	Protocol 				Protocol					`bin:"-"`
}

type Data struct {
	Name 					string
	Symbol					string
	Uri 					string
	SellerFeeBasisPoints 	uint16
	Creators 				*[]token_metadata.Creator `bin:"optional"`
}
