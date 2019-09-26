package sqlstore_test

import (
	"os"
	"testing"
)

var (
	dataBaseURL string
)

func TestMain(m *testing.M) {

	dataBaseURL = os.Getenv("DATABASE_URL")
	if dataBaseURL == "" {
		dataBaseURL = "postgresql://postgres:Daler1995@localhost:5432/restapi_test"
	}

	os.Exit(m.Run())
}
