package main

import "github.com/hucandu/data4life/tokenConsumer/appcontext"
import "github.com/hucandu/data4life/tokenConsumer/services"

func main() {
	appContext := appcontext.Initiate("local")
	tokenService := services.InitTokenService(appContext)
	tokenService.LoadTokens()
}
