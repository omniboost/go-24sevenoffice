package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAccountList(t *testing.T) {
	req1 := client.NewSetIdentityByIDRequest()
	req1.RequestBody().IdentityID = client.IdentityID()
	req1.Do()

	req := client.NewGetAccountListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
