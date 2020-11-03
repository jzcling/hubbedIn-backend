// +build integration

package database

import (
	"in-backend/services/profile/configs"
	"in-backend/services/profile/tests"
	"testing"

	pg "github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/require"
)

func TestPgClientIsWorking(t *testing.T) {
	testConfig, err := configs.GetTestConfig()
	require.NoError(t, err)

	setup, err := tests.SetupTestDB(GetPgConnectionOptions(testConfig), "./scripts/migrations/")
	require.NoError(t, err)

	client := NewClient(setup.DB)
	require.NotNil(t, client)
	db := client.GetConnection()

	var num int
	_, err = db.Query(pg.Scan(&num), "SELECT ?", 42)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 42, num)
}
