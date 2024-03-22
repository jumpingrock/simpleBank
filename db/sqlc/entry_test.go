package db

import (
	"context"
	"github.com/simpleBank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	randAmt := util.RandomMoney()

	arg := CreateEntryParams{
		Amount:    randAmt,
		AccountID: account.ID,
	}

	entryCreate, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entryCreate)

	require.Equal(t, arg.AccountID, entryCreate.AccountID)
	require.Equal(t, arg.Amount, entryCreate.Amount)

	require.NotZero(t, entryCreate.ID)
	require.NotZero(t, entryCreate.CreatedAt)

	return entryCreate
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	randAmt := util.RandomMoney()

	arg := CreateEntryParams{
		Amount:    randAmt,
		AccountID: account.ID,
	}

	entryCreate, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entryCreate)

	require.Equal(t, arg.AccountID, entryCreate.AccountID)
	require.Equal(t, arg.Amount, entryCreate.Amount)

	require.NotZero(t, entryCreate.ID)
	require.NotZero(t, entryCreate.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	entry := CreateEntry(t)

	getEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)

	require.Equal(t, entry.AccountID, getEntry.AccountID)
	require.Equal(t, entry.Amount, getEntry.Amount)
	require.Equal(t, entry.ID, getEntry.ID)

	require.NotZero(t, getEntry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	randAmt := util.RandomMoney()

	arg := CreateEntryParams{
		Amount:    randAmt,
		AccountID: account.ID,
	}

	entryCreate, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryCreate)

	arg.Amount = util.RandomMoney()

	entryCreate2, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryCreate)
	require.Equal(t, entryCreate.AccountID, entryCreate2.AccountID)

	listEntriesArg := ListEntriesParams{
		AccountID: arg.AccountID,
		Limit:     10,
		Offset:    0,
	}

	getEntries, err := testQueries.ListEntries(context.Background(), listEntriesArg)
	require.NoError(t, err)
	require.NotEmpty(t, getEntries)
	require.Len(t, getEntries, 2)
}
