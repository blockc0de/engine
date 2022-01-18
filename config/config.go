package config

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	TIMEOUT               = time.Second * 30
	ETH_RPC_URL           = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	ETH_SOCKET_URL        = "wss://mainnet.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	TELEGRAM_API_ENDPOINT = tgbotapi.APIEndpoint
)
