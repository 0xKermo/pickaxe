// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	CreateAmm(ctx context.Context, arg CreateAmmParams) (Amm, error)
	CreatePool(ctx context.Context, arg CreatePoolParams) (Pool, error)
	CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error)
	DeleteAmm(ctx context.Context, ammID int64) error
	DeletePool(ctx context.Context, address string) error
	DeleteToken(ctx context.Context, address string) error
	GetAmmByDEX(ctx context.Context, dexName string) ([]Amm, error)
	GetAmmById(ctx context.Context, ammID int64) (Amm, error)
	GetBaseTokens(ctx context.Context) ([]Token, error)
	GetNativeTokens(ctx context.Context) ([]Token, error)
	GetPoolByAddress(ctx context.Context, address string) (Pool, error)
	GetPoolsByAmm(ctx context.Context, ammID int64) ([]Pool, error)
	GetPoolsByPair(ctx context.Context, arg GetPoolsByPairParams) ([]Pool, error)
	GetPoolsByToken(ctx context.Context, tokenA string) ([]Pool, error)
	GetTokenByAddress(ctx context.Context, address string) (Token, error)
	UpdateBaseNativeStatus(ctx context.Context, arg UpdateBaseNativeStatusParams) (Token, error)
	UpdateTicker(ctx context.Context, arg UpdateTickerParams) (Token, error)
}

var _ Querier = (*Queries)(nil)