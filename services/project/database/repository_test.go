package database

import (
	"context"
	"in-backend/services/project"
	"in-backend/services/project/configs"
	"in-backend/services/project/models"
	"strings"
	"testing"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx context.Context = context.Background()
	now time.Time       = time.Now()
)

func TestNewRepository(t *testing.T) {
	want := &repository{
		DB: &pg.DB{},
	}

	got := NewRepository(&pg.DB{})

	require.EqualValues(t, want, got)
}

func TestAllCRUD(t *testing.T) {
	testConfig, err := configs.LoadConfig(configs.TestFileName)
	require.NoError(t, err)

	opt := GetPgConnectionOptions(testConfig)

	c, err := setupPGContainer(opt)
	require.NoError(t, err)

	db, err := setupDB(c, opt, "../scripts/migrations/")
	require.NoError(t, err)

	r := NewRepository(db)

	testCreateCandidate(t, r, db)

	cleanDb(db)
	cleanContainer(c)
}

/* --------------- Candidate --------------- */

func testCreateCandidate(t *testing.T, r project.Repository, db *pg.DB) {
	testNoAuthID := &models.Candidate{
		FirstName:     "first",
		LastName:      "last",
		Email:         "first@last.com",
		ContactNumber: "+6591234567",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := *testNoAuthID
	test.AuthID = "authId"

	testDupEmail := test

	// this is required to insert 2 candidates so that one can be used
	// for other tests after the first gets deleted
	test2 := test
	test2.AuthID = "authId2"
	test2.Email = "test@test.com"
	test2.ContactNumber = "+6587654321"

	type args struct {
		ctx   context.Context
		input *models.Candidate
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter candidate is nil")}},
		{"failed not null", args{ctx, testNoAuthID}, expect{nil, errors.New("Failed to insert candidate")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
		{"failed unique", args{ctx, &testDupEmail}, expect{nil, errors.New("Failed to insert candidate")}},
		{"valid2", args{ctx, &test2}, expect{&test2, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCandidate(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
