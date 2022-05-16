package services

import (
	"bufio"
	"fmt"
	"github.com/hucandu/data4life/tokenConsumer/appcontext"
	"log"
	"os"
	"sync"
)
import "github.com/hucandu/data4life/tokenConsumer/models"

type TokenService struct {
	DbOpContext *models.TokenDataOperations
}

func InitTokenService(appCon *appcontext.AppContext) *TokenService {
	return &TokenService{DbOpContext: &models.TokenDataOperations{DB: appCon.DbClient}}
}

func (tis TokenService) FetchTokenFromDB() ([]models.TokenData, error) {
	tokenData, err := tis.DbOpContext.ListTokenData()
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

func (tis TokenService) LoadTokens() error {
	file, err := os.Open("file/token.text")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var waitgroup sync.WaitGroup
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		waitgroup.Add(1)
		go func(token string, waitgroup *sync.WaitGroup) {
			err := tis.DbOpContext.InsertTokenData(&models.TokenData{
				Token: token,
				Count: 1,
			})
			if err != nil {
				log.Fatal(err)
			}
			waitgroup.Done()
		}(scanner.Text(), &waitgroup)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return err
}
