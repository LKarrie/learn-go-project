package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all func to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context,arg TransferTxParams) (TransferTxResult,error)
}

// SQLStore provides all func to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:db,
	}
}

// execTx executes a func within a db transaction
func (store *SQLStore) execTx(ctx context.Context,fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			 return fmt.Errorf("tx err: %v,rb err: %v",err,rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

// var txKey = struct{}{}

// TransferTx performs a money transfer from one account to other
// It create a transfer record, and account entries, and update accounts balance within a single database transaction
func (Store *SQLStore) TransferTx(ctx context.Context,arg TransferTxParams) (TransferTxResult,error) {
	var result TransferTxResult

	err := Store.execTx(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey)

		// fmt.Println(txName,"create transfer")
		result.Transfer,err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		// fmt.Println(txName,"create entry 1")
		result.FromEntry,err =q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		// fmt.Println(txName,"create entry 2")
		result.ToEntry,err =q.CreateEntry(ctx,CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		// update balance
		// fmt.Println(txName,"get account 1")
		// account1,err := q.GetAccountForUpdate(ctx,arg.FromAccountID)
		// if err != nil {
		// 	return err
		// }

		// fmt.Println(txName,"update account 1")
		// result.FromAccount, err = q.UpdateAccount(ctx,UpdateAccountParams{
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount,result.ToAccount,_ = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount,result.FromAccount,_ = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result,err
}


func addMoney (
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account,account2 Account,err error) {
	account1,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID1,
		Amount: amount1,
	})
	if err != nil {
		   return
	}
	account2,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID2,
		Amount: amount2,
	})
	return
}