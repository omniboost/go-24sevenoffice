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

	req.RequestBody().SearchParams.DateStart = twentyfour.Date{time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().SearchParams.DateEnd = twentyfour.Date{time.Date(2026, 9, 3, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().SearchParams.TransactionNoStart = 3260902
	req.RequestBody().SearchParams.TransactionNoEnd = 3260902
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
