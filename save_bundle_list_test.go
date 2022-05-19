package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSaveBundleList(t *testing.T) {
	req := client.NewSaveBundleListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
