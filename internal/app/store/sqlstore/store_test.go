package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseTestURL string
)

func TestMain(m *testing.M) {
	databaseTestURL = os.Getenv("database_test_url")
	if databaseTestURL == "" {
		databaseTestURL = "postgresql://ps5:PsX314159_@localhost:5432/learn_go_crud_test"
	}

	os.Exit(m.Run())
}
