package internals

import (
	"bytes"
	"encoding/json"
	"fintual-cli/internals/models"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

const (
	BaseURL = "https://fintual.cl/api/"
)

// Endpoints
const (
	Banks       = "banks"
	AccessToken = "access_tokens"
	Goals       = "goals"
)

type FintualAPI struct {
	BaseURL string
}

func (fintualPath FintualAPI) GetBanks() ([]models.Bank, error) {
	url := fmt.Sprintf("%s%s", BaseURL, Banks)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bankResponse models.BankResponse

	err = json.Unmarshal(body, &bankResponse)
	if err != nil {
		return nil, err
	}

	banks := bankResponse.Data
	sort.Slice(banks, func(i, j int) bool {
		iInt, iErr := strconv.Atoi(banks[i].Id)
		jInt, jErr := strconv.Atoi(banks[j].Id)
		if iErr != nil {
			panic(iErr)
		}
		if jErr != nil {
			panic(jErr)
		}
		return iInt < jInt
	})

	return banks, nil
}

func (fintualAPI FintualAPI) ObtainToken(email string, password string) (string, error) {
	payload := models.UserPayload{
		User: models.User{
			Email:    email,
			Password: password,
		},
	}
	postBody, _ := json.Marshal(payload)
	responseBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%s%s", BaseURL, AccessToken)

	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error on response, error %d", resp.StatusCode)
	}

	if resp.StatusCode < 400 {
		fmt.Printf("Good response with %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var tokenResponse models.TokenResponse

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Data.Attributes.Token, nil
}

var FintualAPIRepository = FintualAPI{BaseURL: BaseURL}
