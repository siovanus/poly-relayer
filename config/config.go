/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/polynetwork/bridge-common/wallet"
)

type ListenerConfig struct {
	ChainId     uint64
	Nodes       []string
	LockProxy   []string
	CCMContract string
	CCDContract string
	Defer       int
}

type PolySubmitterConfig struct {
	ChainId uint64
	Nodes   []string
	Procs   int
	Wallet  *wallet.PolySignerConfig
}

type SubmitterConfig struct {
	ChainId     uint64
	Nodes       []string
	CCMContract string
	CCDContract string
	Wallet      *wallet.Config
}

type WalletConfig struct {
	Nodes    []string
	KeyStore string
	KeyPwd   map[string]string
}

type BusConfig struct {
	Redis *redis.Options
}
