package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hucandu/data4life/tokenConsumer/appcontext"
	"log"
	"os"
	"path/filepath"
	"time"
)
import "github.com/hucandu/data4life/tokenConsumer/models"

type TokenService struct {
	DbOpContext      *models.TokenDataOperations
	TokenFreqHashMap *map[string]models.TokenData
}

func InitTokenService(appCon *appcontext.AppContext) *TokenService {
	return &TokenService{
		DbOpContext:      &models.TokenDataOperations{DB: appCon.DbClient},
		TokenFreqHashMap: &map[string]models.TokenData{},
	}
}

func (tis TokenService) FetchTokenFromDB() (string, error) {
	fmt.Println("Start Fetch Token")
	tokenData, err := tis.DbOpContext.ListTokenData()
	if err != nil {
		return "", err
	}
	response, err := json.Marshal(tokenData)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, response, "", "\t")
	filePath, _ := filepath.Abs("tokenConsumer/output.json")
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(prettyJSON.Bytes()))
	fmt.Println(("tokens fetched at " + filePath))

	return string(prettyJSON.Bytes()), nil
}

func persistTokens(ctx *TokenService, token string) {
	if _, ok := (*ctx.TokenFreqHashMap)[token]; ok {
		(*ctx.TokenFreqHashMap)[token] = models.TokenData{
			Token:     token,
			ID:        (*ctx.TokenFreqHashMap)[token].ID,
			Count:     (*ctx.TokenFreqHashMap)[token].Count + 1,
			CreatedAt: (*ctx.TokenFreqHashMap)[token].CreatedAt,
			UpdatedAt: time.Now(),
		}

	} else {
		(*ctx.TokenFreqHashMap)[token] = models.TokenData{
			Token:     token,
			ID:        uuid.New().String(),
			Count:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

	}
}

func (tis TokenService) LoadTokens() error {
	fmt.Println("Start Load Token")
	filePath, err := filepath.Abs("tokenGenerator/tokens.txt")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		persistTokens(&tis, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	tokenDataList := []models.TokenData{}
	for _, value := range *tis.TokenFreqHashMap {
		tokenDataList = append(tokenDataList, value)
	}
	tis.DbOpContext.BulkInsertTokenData(tokenDataList)
	fmt.Println("Tokens Loaded")
	return err
}
