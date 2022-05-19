package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSaveInvoices(t *testing.T) {
	req := client.NewSaveInvoicesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
