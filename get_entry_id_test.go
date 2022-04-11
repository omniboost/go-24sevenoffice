package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	twentyfour "github.com/omniboost/go-24sevenoffice"
)

func TestGetEntryID(t *testing.T) {
	req := client.NewGetEntryIDRequest()
	req.RequestBody().ArgEntryID.Date = twentyfour.Date{time.Now().AddDate(0, 0, 0)}
	req.RequestBody().ArgEntryID.SortNo = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
