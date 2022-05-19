package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSaveCompanies(t *testing.T) {
	req := client.NewSaveCompaniesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
