package nft_proxy

import (
	"os"

	"github.com/gagliardetto/solana-go"
)

const (
	BASE64_PREFIX = ";base64,"
)

var (
	METAPLEX_CORE = solana.MustPublicKeyFromBase58(os.Getenv("METAPLEX_CORE_KEY"))
	TOKEN_2022    = solana.MustPublicKeyFromBase58(os.Getenv("TOKEN_2022_KEY"))
)
