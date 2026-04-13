package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	Transfer1 := createRandomTransfer(t, account1, account2)
	Transfer2, err := testQueries.GetTransfer(context.Background(), Transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Transfer2)
	require.Equal(t, Transfer1.ID, Transfer2.ID)
	require.Equal(t, Transfer1.FromAccountID, Transfer2.FromAccountID)
	require.Equal(t, Transfer1.ToAccountID, Transfer2.ToAccountID)
	require.Equal(t, Transfer1.Amount, Transfer2.Amount)
	require.WithinDuration(t, Transfer1.CreatedAt, Transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}
	Transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	for _, Transfer := range Transfers {
		require.NotEmpty(t, Transfer)
	}

}
