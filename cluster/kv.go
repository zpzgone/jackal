/*
 * Copyright (c) 2020 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package cluster

import "context"

type KV interface {
	Put(ctx context.Context, key string, value string, ttlInSeconds int64) error
	Get(ctx context.Context, key string) (string, error)

	GetPrefix(ctx context.Context, prefix string) ([]string, error)
}