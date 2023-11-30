package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTypeGroupList(t *testing.T) {
	req := client.NewGetTypeGroupListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
