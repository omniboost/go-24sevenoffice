package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	req := client.NewLoginRequest()
	req.RequestBody().Credential.ApplicationID = client.ApplicationID()
	req.RequestBody().Credential.Username = client.Username()
	req.RequestBody().Credential.Password = client.Password()
	req.RequestBody().Credential.IdentityID = client.IdentityID()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
