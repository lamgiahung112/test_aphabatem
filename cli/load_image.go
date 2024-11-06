package main

import (
	"context"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	client := rpc.New(os.Getenv("RPC_URL"))
	resp, err := client.GetAccountInfo(context.TODO(), solana.MustPublicKeyFromBase58(os.Getenv("ACCOUNT_INFO_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", resp.Value.Data.GetBinary())

	pResp, err := client.GetProgramAccountsWithOpts(context.TODO(), solana.MustPublicKeyFromBase58(os.Getenv("PROGRAM_ACCOUNT_KEY")), &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: 1 + 32 + 33 + 4 + 4 + 1},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range pResp {
		log.Printf("%s - Size: %v", r.Pubkey, len(r.Account.Data.GetBinary()))
	}
}
