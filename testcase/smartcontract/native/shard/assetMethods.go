package shard

import (
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/testframework"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native"
	"github.com/ontio/ontology/smartcontract/service/native/shardasset/oep4"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
	"math/big"
)

func AssetInit(ctx *testframework.TestFrameworkContext, user *sdk.Account) error {
	method := oep4.INIT
	contractAddress := utils.ShardAssetAddress
	txHash, err := ctx.Ont.Native.InvokeNativeContract(ctx.GetGasPrice(), ctx.GetGasLimit(), user, 0,
		contractAddress, method, []interface{}{})
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}

func XShardTransfer(ctx *testframework.TestFrameworkContext, users []*sdk.Account, contractAddress common.Address,
	to []common.Address, amount []uint64, fromShard, toShard common.ShardID, shardUrl, method string) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	for i, user := range users {
		toAddr := to[i]
		num := amount[i]
		param := &oep4.XShardTransferParam{
			From:    user.Address,
			To:      toAddr,
			ToShard: toShard,
			Amount:  new(big.Int).SetUint64(num),
		}
		var txHash common.Uint256
		var err error = nil
		if _, ok := native.Contracts[contractAddress]; ok {
			txHash, err = ctx.Ont.Native.InvokeShardNativeContract(fromShard, ctx.GetGasPrice(),
				ctx.GetGasLimit(), user, 0, contractAddress, method, []interface{}{param})
		} else {
			txHash, err = ctx.Ont.NeoVM.InvokeShardNeoVMContract(fromShard, ctx.GetGasPrice(),
				ctx.GetGasLimit(), user, contractAddress,
				[]interface{}{method, []interface{}{user.Address, toAddr, toShard, num}})
		}
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	}

	return nil
}

func XShardTransferRetry(ctx *testframework.TestFrameworkContext, fromShard common.ShardID, users []*sdk.Account,
	contractAddress common.Address, transferId []uint64, shardUrl, method string) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	for i, user := range users {
		id := transferId[i]
		param := &oep4.XShardTransferRetryParam{
			From:       user.Address,
			TransferId: new(big.Int).SetUint64(id),
		}
		var txHash common.Uint256
		var err error = nil
		if _, ok := native.Contracts[contractAddress]; ok {
			txHash, err = ctx.Ont.Native.InvokeShardNativeContract(fromShard, ctx.GetGasPrice(),
				ctx.GetGasLimit(), user, 0, contractAddress, method, []interface{}{param})
		} else {
			txHash, err = ctx.Ont.NeoVM.InvokeShardNeoVMContract(fromShard, ctx.GetGasPrice(),
				ctx.GetGasLimit(), user, contractAddress, []interface{}{method, []interface{}{user.Address, id}})
		}
		if err != nil {
			return fmt.Errorf("invokeNativeContract error :", err)
		}
		ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	}
	return nil
}

func GetPendingTransfer(ctx *testframework.TestFrameworkContext, addr common.Address, assetId uint64, shardUrl string) error {
	method := oep4.GET_PENDING_TRANSFER
	contractAddress := utils.ShardAssetAddress
	param := &oep4.GetPendingXShardTransferParam{
		Account: addr,
		Asset:   assetId,
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, common.NewShardIDUnchecked(0),
		[]interface{}{param})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("pending transfer is: %s", info)
	return nil
}

func GetTransferDetail(ctx *testframework.TestFrameworkContext, user common.Address, assetId, transferId uint64,
	shardUrl string) error {
	method := oep4.GET_TRANSFER
	contractAddress := utils.ShardAssetAddress
	param := &oep4.GetXShardTransferInfoParam{
		Account:    user,
		Asset:      assetId,
		TransferId: new(big.Int).SetUint64(transferId),
	}
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, common.NewShardIDUnchecked(0),
		[]interface{}{param})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("transfer is: %s", info)
	return nil
}

func GetSupplyInfo(ctx *testframework.TestFrameworkContext, assetId uint64, shardUrl string) error {
	method := oep4.SUPPLY_INFO
	contractAddress := utils.ShardAssetAddress
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.Native.PreExecInvokeShardNativeContract(contractAddress, byte(0), method, common.NewShardIDUnchecked(0),
		[]interface{}{assetId})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToString()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("supply info is: %s", info)
	return nil
}

func GetOep4Balance(ctx *testframework.TestFrameworkContext, user, contract common.Address, shardId common.ShardID,
	shardUrl string) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	value, err := ctx.Ont.NeoVM.PreExecInvokeShardNeoVMContract(shardId, contract,
		[]interface{}{"balanceOf", []interface{}{user}})
	if err != nil {
		return fmt.Errorf("pre-execute err: %s", err)
	}
	info, err := value.Result.ToInteger()
	if err != nil {
		return fmt.Errorf("parse result failed, err: %s", err)
	}
	ctx.LogInfo("balance of %s is: %s", user.ToBase58(), info.String())
	return nil
}

func ChangeMetaData(ctx *testframework.TestFrameworkContext, user *sdk.Account, contract common.Address, shardId common.ShardID,
	shardUrl string, newOwner common.Address, frozen bool, invokedContracts []common.Address) error {
	ctx.Ont.ClientMgr.GetRpcClient().SetAddress(shardUrl)
	txHash, err := ctx.Ont.NeoVM.ChangeMetaData(shardId, ctx.GetGasPrice(), ctx.GetGasLimit(), user, contract,
		newOwner, frozen, invokedContracts)
	if err != nil {
		return fmt.Errorf("invokeNativeContract error :", err)
	}
	ctx.LogInfo("txHash is: %s", txHash.ToHexString())
	return nil
}
