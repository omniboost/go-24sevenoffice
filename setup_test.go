package twentyfour_test

import (
	"os"
	"testing"

	netsuite "github.com/omniboost/go-24sevenoffice"
)

var (
	client *netsuite.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	applicationID := os.Getenv("APPLICATION_ID")
	identityID := os.Getenv("IDENTITY_ID")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	debug := os.Getenv("DEBUG")

	client = netsuite.NewClient(nil)
	client.SetApplicationID(applicationID)
	client.SetIdentityID(identityID)
	client.SetUsername(username)
	client.SetPassword(password)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	m.Run()
}
