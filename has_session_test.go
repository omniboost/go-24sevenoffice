package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHasSession(t *testing.T) {
	sessionID, _ := client.FetchSessionID()
	client.SetSessionID(sessionID)
	req := client.NewHasSessionRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
