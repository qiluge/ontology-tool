package shard

import (
	"encoding/json"
	"github.com/ontio/ontology/common"
	"io/ioutil"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
)

type ChangePeerAttrParam struct {
	ShardId    common.ShardID `json:"shard_id"`
	PeerOwners []string       `json:"peer_owners"`
	Peers      []string       `json:"peers"`
	Amount     []uint64       `json:"amount"`
}

func TestShardChangePeerMaxAuthorization(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardPeerChangeMaxAuth.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangePeerAttrParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.PeerOwners {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardPeerChangeMaxAuth(ctx, param.ShardId, users, param.Peers, param.Amount); err != nil {
		ctx.LogError("TestShardChangePeerMaxAuthorization failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestShardChangePeerProportion(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardPeerChangeProportion.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangePeerAttrParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.PeerOwners {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardPeerChangeProportion(ctx, param.ShardId, users, param.Peers, param.Amount); err != nil {
		ctx.LogError("TestShardChangePeerProportion failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type UserStakeParam struct {
	Path       string         `json:"path"`
	ShardId    common.ShardID `json:"shard_id"`
	PeerPubKey []string       `json:"peer_pub_key"`
	Amount     []uint64       `json:"amount"`
}

func TestShardUserStake(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserStake.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserStakeParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		return false
	}

	if err := ShardUserStake(ctx, user, param.ShardId, param.PeerPubKey, param.Amount); err != nil {
		ctx.LogError("TestShardUserStake failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestShardUserUnfreezeStake(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserUnfreezeStake.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserStakeParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		return false
	}

	if err := ShardUserUnfreezeStake(ctx, user, param.ShardId, param.PeerPubKey, param.Amount); err != nil {
		ctx.LogError("TestShardUserStake failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type UserWithdrawStakeParam struct {
	Path    string         `json:"path"`
	ShardId common.ShardID `json:"shard_id"`
}

func TestShardUserWithdrawStake(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserWithdrawStake.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserWithdrawStakeParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		return false
	}

	if err := ShardUserWithdrawStake(ctx, user, param.ShardId); err != nil {
		ctx.LogError("TestShardUserStake failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

func TestShardUserWithdrawFee(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserWithdrawFee.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserWithdrawStakeParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Path)
	if !ok {
		return false
	}

	if err := ShardUserWithdrawFee(ctx, user, param.ShardId); err != nil {
		ctx.LogError("TestShardUserStake failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type UserWithdrawOngParam struct {
	Wallets []string `json:"wallets"`
}

func TestShardUserWithdrawOng(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardUserWithdrawOng.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &UserWithdrawOngParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal shard init param: %s", err)
		return false
	}
	users := make([]*sdk.Account, 0)
	for _, peerWallet := range param.Wallets {
		user, ok := getAccountByPassword(ctx, peerWallet)
		if !ok {
			ctx.LogError("get account failed")
			return false
		}
		users = append(users, user)
	}
	if err := ShardUserWithdrawOng(ctx, users); err != nil {
		ctx.LogError("TestShardUserWithdrawOng failed: %s", err)
		return false
	}
	waitForBlock(ctx)
	return true
}

type GetShardViewParam struct {
	ShardId common.ShardID `json:"shard_id"`
}

func TestGetShardView(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardGetView.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetShardViewParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	if err := ShardQueryView(ctx, param.ShardId); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

type GetPeerInfoParam struct {
	ShardId common.ShardID `json:"shard_id"`
	View    uint64         `json:"view"`
}

func TestGetShardPeerInfo(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardGetPeerInfo.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetPeerInfoParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}

	if err := ShardQueryPeerInfo(ctx, param.ShardId, param.View); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

type GetUserInfoParam struct {
	ShardId common.ShardID `json:"shard_id"`
	View    uint64         `json:"view"`
	Address string         `json:"address"`
}

func TestGetShardUserInfo(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardGetUserInfo.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &GetUserInfoParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}
	addr, err := common.AddressFromBase58(param.Address)
	if err != nil {
		ctx.LogError("decode addr failed, err: %s", err)
		return false
	}
	if err := ShardQueryUserInfo(ctx, param.ShardId, param.View, addr); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

type ChangeInitPosParam struct {
	Wallet  string         `json:"wallet"`
	ShardId common.ShardID `json:"shard_id"`
	Peer    string         `json:"peer"`
	Amount  uint64         `json:"amount"`
}

func TestAddInitPos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardAddInitPos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangeInitPosParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Wallet)
	if !ok {
		return false
	}
	if err := ShardAddInitPos(ctx, param.ShardId, user, param.Peer, param.Amount); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}

func TestReduceInitPos(ctx *testframework.TestFrameworkContext) bool {
	configFile := "./params/shard_stake/ShardReduceInitPos.json"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		ctx.LogError("read config from %s: %s", configFile, err)
		return false
	}

	param := &ChangeInitPosParam{}
	if err := json.Unmarshal(data, param); err != nil {
		ctx.LogError("unmarshal param: %s", err)
		return false
	}
	user, ok := getAccountByPassword(ctx, param.Wallet)
	if !ok {
		return false
	}
	if err := ShardReduceInitPos(ctx, param.ShardId, user, param.Peer, param.Amount); err != nil {
		ctx.LogError("failed: %s", err)
		return false
	}
	return true
}
