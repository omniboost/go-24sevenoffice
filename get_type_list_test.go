package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTypeList(t *testing.T) {
	req := client.NewGetTypeListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
