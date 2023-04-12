// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: pools.sql

package db

import (
	"context"
	"database/sql"
)

const createPool = `-- name: CreatePool :one
INSERT INTO pools_v2 (
  address,
  amm_id,
  token_a,
  token_b,
  fee
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data
`

type CreatePoolParams struct {
	Address string `json:"address"`
	AmmID   int64  `json:"amm_id"`
	TokenA  string `json:"token_a"`
	TokenB  string `json:"token_b"`
	Fee     string `json:"fee"`
}

func (q *Queries) CreatePool(ctx context.Context, arg CreatePoolParams) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, createPool,
		arg.Address,
		arg.AmmID,
		arg.TokenA,
		arg.TokenB,
		arg.Fee,
	)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}

const deletePool = `-- name: DeletePool :exec
DELETE FROM pools_v2
WHERE address = $1
`

func (q *Queries) DeletePool(ctx context.Context, address string) error {
	_, err := q.db.ExecContext(ctx, deletePool, address)
	return err
}

const getAllPools = `-- name: GetAllPools :many
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
ORDER BY address
`

func (q *Queries) GetAllPools(ctx context.Context) ([]PoolsV2, error) {
	rows, err := q.db.QueryContext(ctx, getAllPools)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PoolsV2{}
	for rows.Next() {
		var i PoolsV2
		if err := rows.Scan(
			&i.PoolID,
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.Fee,
			&i.TotalValue,
			&i.LastUpdated,
			&i.ExtraData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPoolByAddress = `-- name: GetPoolByAddress :one
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
WHERE address = $1 LIMIT 1
`

func (q *Queries) GetPoolByAddress(ctx context.Context, address string) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, getPoolByAddress, address)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}

const getPoolByAddressExtra = `-- name: GetPoolByAddressExtra :one
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
WHERE address = $1 AND extra_data=$2 LIMIT 1
`

type GetPoolByAddressExtraParams struct {
	Address   string         `json:"address"`
	ExtraData sql.NullString `json:"extra_data"`
}

func (q *Queries) GetPoolByAddressExtra(ctx context.Context, arg GetPoolByAddressExtraParams) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, getPoolByAddressExtra, arg.Address, arg.ExtraData)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}

const getPoolsByAmm = `-- name: GetPoolsByAmm :many
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
WHERE amm_id = $1
ORDER BY address
`

func (q *Queries) GetPoolsByAmm(ctx context.Context, ammID int64) ([]PoolsV2, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByAmm, ammID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PoolsV2{}
	for rows.Next() {
		var i PoolsV2
		if err := rows.Scan(
			&i.PoolID,
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.Fee,
			&i.TotalValue,
			&i.LastUpdated,
			&i.ExtraData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPoolsByPair = `-- name: GetPoolsByPair :many
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
WHERE token_a = $1 AND token_b = $2
ORDER BY amm_id
`

type GetPoolsByPairParams struct {
	TokenA string `json:"token_a"`
	TokenB string `json:"token_b"`
}

func (q *Queries) GetPoolsByPair(ctx context.Context, arg GetPoolsByPairParams) ([]PoolsV2, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByPair, arg.TokenA, arg.TokenB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PoolsV2{}
	for rows.Next() {
		var i PoolsV2
		if err := rows.Scan(
			&i.PoolID,
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.Fee,
			&i.TotalValue,
			&i.LastUpdated,
			&i.ExtraData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPoolsByToken = `-- name: GetPoolsByToken :many
SELECT pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data FROM pools_v2
WHERE token_a = $1 OR token_b = $1
ORDER BY amm_id
`

func (q *Queries) GetPoolsByToken(ctx context.Context, tokenA string) ([]PoolsV2, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByToken, tokenA)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PoolsV2{}
	for rows.Next() {
		var i PoolsV2
		if err := rows.Scan(
			&i.PoolID,
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.Fee,
			&i.TotalValue,
			&i.LastUpdated,
			&i.ExtraData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePoolExtraData = `-- name: UpdatePoolExtraData :one
UPDATE pools_v2
SET extra_data = $2
WHERE pool_id = $1
RETURNING pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data
`

type UpdatePoolExtraDataParams struct {
	PoolID    int64          `json:"pool_id"`
	ExtraData sql.NullString `json:"extra_data"`
}

func (q *Queries) UpdatePoolExtraData(ctx context.Context, arg UpdatePoolExtraDataParams) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, updatePoolExtraData, arg.PoolID, arg.ExtraData)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}

const updatePoolReserves = `-- name: UpdatePoolReserves :one
UPDATE pools_v2
SET reserve_a = $1, reserve_b = $2, last_updated = NOW()
WHERE pool_id = $3
RETURNING pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data
`

type UpdatePoolReservesParams struct {
	ReserveA string `json:"reserve_a"`
	ReserveB string `json:"reserve_b"`
	PoolID   int64  `json:"pool_id"`
}

func (q *Queries) UpdatePoolReserves(ctx context.Context, arg UpdatePoolReservesParams) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, updatePoolReserves, arg.ReserveA, arg.ReserveB, arg.PoolID)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}

const updatePoolTV = `-- name: UpdatePoolTV :one
UPDATE pools_v2
SET total_value = $1
WHERE pool_id = $2
RETURNING pool_id, address, amm_id, token_a, token_b, reserve_a, reserve_b, fee, total_value, last_updated, extra_data
`

type UpdatePoolTVParams struct {
	TotalValue string `json:"total_value"`
	PoolID     int64  `json:"pool_id"`
}

func (q *Queries) UpdatePoolTV(ctx context.Context, arg UpdatePoolTVParams) (PoolsV2, error) {
	row := q.db.QueryRowContext(ctx, updatePoolTV, arg.TotalValue, arg.PoolID)
	var i PoolsV2
	err := row.Scan(
		&i.PoolID,
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.Fee,
		&i.TotalValue,
		&i.LastUpdated,
		&i.ExtraData,
	)
	return i, err
}
