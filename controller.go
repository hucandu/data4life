package data4life

import (
	"fmt"
	"github.com/hucandu/data4life/tokenConsumer/appcontext"
	"github.com/hucandu/data4life/tokenConsumer/services"
	"github.com/hucandu/data4life/tokenGenerator"
	"github.com/manifoldco/promptui"
	"log"
)

func Controller(env string) {
	appContext := appcontext.Initiate(env)
	tokenService := services.InitTokenService(appContext)
	for true {
		prompt := promptui.Select{
			Label: "Choose an option",
			Items: []string{"Generate Tokens", "Load Tokens Into DB", "Fetch Tokens From DB"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case "Generate Tokens":
			tokenGenerator.GenerateToken()
		case "Load Tokens Into DB":
			if err := tokenService.LoadTokens(); err != nil {
				log.Fatal(err)
			}
		case "Fetch Tokens From DB":
			if _, err := tokenService.FetchTokenFromDB(); err != nil {
				log.Fatal(err)
			}
		}

	}

}
