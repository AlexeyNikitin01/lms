package pg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type txKey struct{}

// MustCtxWithTx кладем транзакцию в контекст, если она уже есть вылетит паника.
func MustCtxWithTx(ctx context.Context, exec boil.ContextExecutor) context.Context {
	if exec, ok := ctx.Value(txKey{}).(boil.ContextExecutor); ok || exec != nil {
		panic("tx already running")
	}

	return context.WithValue(ctx, txKey{}, exec)
}

// TxFromContext пытается взять транзакцию из контекста, если нет, то берется обычное соединение.
func TxFromContext(ctx context.Context) boil.ContextExecutor {
	if exec, ok := ctx.Value(txKey{}).(boil.ContextExecutor); ok && exec != nil {
		return exec
	}

	return boil.GetContextDB()
}

// ExecTx исполняет query через транзакцию, если была транзакция запущена ранее, то будет использована она
//
// Начатая транзакция прокидывается в контекст, который передается в query.
func ExecTx(ctx context.Context, query func(txCtx context.Context, executor boil.ContextExecutor) error) error {
	// если транзакция уже запущена, то используем ее
	if val := ctx.Value(txKey{}); val != nil {
		if exec, ok := val.(boil.ContextExecutor); ok {
			if err := query(ctx, exec); err != nil {
				return errors.Wrap(err, "on exec query")
			}

			return nil
		}
	}

	// иначе стартуем новую
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "on begin tx")
	}

	newCtx := MustCtxWithTx(ctx, tx)

	if err = query(newCtx, tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrapf(err, "on tx rollback")
		}

		return errors.Wrapf(err, "on exec query")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "on commit tx")
	}

	return nil
}
