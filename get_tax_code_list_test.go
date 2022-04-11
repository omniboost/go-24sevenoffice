package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTaxCodeList(t *testing.T) {
	req := client.NewGetTaxCodeListRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
