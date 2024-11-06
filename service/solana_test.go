package services

import (
	"log"
	"os"
	"path"
	"testing"

	nft_proxy "github.com/alphabatem/nft-proxy"
	"github.com/gagliardetto/solana-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(path.Join(nft_proxy.BASE_PATH, ".env"))
	if err != nil {
		log.Print("Error loading .env file")
	}
}

func TestSolanaImageService_FetchMetadata(t *testing.T) {
	pk := solana.MustPublicKeyFromBase58(os.Getenv("METAPLEX_CORE_KEY"))

	svc := SolanaService{}
	svc.Start()

	d, _, err := svc.TokenData(pk)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", d)
}
