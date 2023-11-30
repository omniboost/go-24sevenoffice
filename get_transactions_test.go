package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	twentyfour "github.com/omniboost/go-24sevenoffice"
)

func TestGetTransactions(t *testing.T) {
	req := client.NewGetTransactionsRequest()

	req.RequestBody().SearchParams.DateStart = twentyfour.Date{time.Date(2022, 9, 13, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().SearchParams.DateEnd = twentyfour.Date{time.Date(2022, 9, 14, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
