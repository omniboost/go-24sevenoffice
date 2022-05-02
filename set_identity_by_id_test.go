package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetIdentityByID(t *testing.T) {
	req := client.NewSetIdentityByIDRequest()
	req.RequestBody().IdentityID = "da763c4d-0003-404b-a0b1-0001003b0057"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
