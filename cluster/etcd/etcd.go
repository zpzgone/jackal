/*
 * Copyright (c) 2020 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package etcd

import (
	v3 "github.com/coreos/etcd/clientv3"
)

func New(cfg Config, allocationID string) (*Candidate, *KV, error) {
	c, err := v3.New(v3.Config{Endpoints: cfg.Endpoints})
	if err != nil {
		return nil, nil, err
	}
	candidate, err := newCandidate(c, allocationID)
	if err != nil {
		return nil, nil, err
	}
	return candidate, &KV{cli: c}, nil
}
