package db

import (
	"context"
	"github.com/dbraley/simplebank/util"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestCreateEntry_InvalidAccountId(t *testing.T) {
	args := CreateEntryParams{
		AccountID: 0,
		Amount:    0,
	}

	// This seems fragile, but is probably better than nothing
	expectedErr := pq.Error{
		Message: "insert or update on table \"entries\" violates foreign key constraint \"entries_account_id_fkey\"",
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.Error(t, err)
	require.EqualError(t, err, expectedErr.Error())
	require.Empty(t, entry)
}

func TestGetEntry(t *testing.T) {
	want := createRandomEntry(t)

	got, err := testQueries.GetEntry(context.Background(), want.ID)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, want.ID, got.ID)
	require.Equal(t, want.Amount, got.Amount)
	require.Equal(t, want.AccountID, got.AccountID)

	require.WithinDuration(t, want.CreatedAt, got.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
