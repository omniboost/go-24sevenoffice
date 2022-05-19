package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	twentyfour "github.com/omniboost/go-24sevenoffice"
)

func TestGetCompanySyncList(t *testing.T) {
	req := client.NewGetCompanySyncListRequest()
	req.RequestBody().SyncSearchParameters.Page = 1
	req.RequestBody().SyncSearchParameters.ChangedAfter = twentyfour.DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
