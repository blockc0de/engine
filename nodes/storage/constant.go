package storage

import "fmt"

const (
	KeysLimitPerGraph  = 30
	KeysLimitPerWallet = 100
)

func redisWalletKey(scope string) string {
	return fmt.Sprintf("wallet::keys::{%s}", scope)
}
