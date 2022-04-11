package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetIdentitiesWithCredential(t *testing.T) {
	req := client.NewGetIdentitiesWithCredentialRequest()
	req.RequestBody().Credential.ApplicationID = client.ApplicationID()
	req.RequestBody().Credential.Username = client.Username()
	req.RequestBody().Credential.Password = client.Password()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
