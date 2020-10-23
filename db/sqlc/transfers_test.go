package db

import (
	"context"
	"github.com/dbraley/simplebank/util"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestCreateTransfer_InvalidToAccountId(t *testing.T) {
	fromAccount := createRandomAccount(t)
	args := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   0,
		Amount:        util.RandomMoney(),
	}

	// This seems fragile, but is probably better than nothing
	expectedErr := pq.Error{
		Message: "insert or update on table \"transfers\" violates foreign key constraint \"transfers_to_account_id_fkey\"",
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.Error(t, err)
	require.EqualError(t, err, expectedErr.Error())
	require.Empty(t, transfer)
}

func TestCreateTransfer_InvalidFromAccountId(t *testing.T) {
	toAccount := createRandomAccount(t)
	args := CreateTransferParams{
		FromAccountID: 0,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	// This seems fragile, but is probably better than nothing
	expectedErr := pq.Error{
		Message: "insert or update on table \"transfers\" violates foreign key constraint \"transfers_from_account_id_fkey\"",
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.Error(t, err)
	require.EqualError(t, err, expectedErr.Error())
	require.Empty(t, transfer)
}

func TestGetTransfer(t *testing.T) {
	want := createRandomTransfer(t)

	got, err := testQueries.GetTransfer(context.Background(), want.ID)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, want.ID, got.ID)
	require.Equal(t, want.Amount, got.Amount)
	require.Equal(t, want.ToAccountID, got.ToAccountID)
	require.Equal(t, want.FromAccountID, got.FromAccountID)

	require.WithinDuration(t, want.CreatedAt, got.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	args := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	require.Equal(t, toAccount.ID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
