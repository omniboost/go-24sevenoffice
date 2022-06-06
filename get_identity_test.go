package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetIdentity(t *testing.T) {
	req := client.NewGetIdentityRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
