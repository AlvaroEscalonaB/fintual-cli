package internals

import (
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
	Banks = "banks"
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
		iInt, Ierr := strconv.Atoi(banks[i].Id)
		jInt, Jerr := strconv.Atoi(banks[j].Id)
		if Ierr != nil {
			panic(Ierr)
		}
		if Jerr != nil {
			panic(Jerr)
		}
		return iInt < jInt
	})

	return banks, nil
}

var FintualAPIRepository = FintualAPI{BaseURL: BaseURL}
