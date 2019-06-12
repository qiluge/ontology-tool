package shard

import "github.com/ontio/ontology-tool/testframework"

func TestShardMgmtContract() {
	testframework.TFramework.RegTestCase("ShardInit", TestShardInit)
	testframework.TFramework.RegTestCase("ShardCreate", TestShardCreate)
	testframework.TFramework.RegTestCase("ShardConfig", TestShardConfig)
	testframework.TFramework.RegTestCase("ShardPeerJoin", TestShardPeerJoin)
	testframework.TFramework.RegTestCase("ShardPeerApply", TestShardPeerApplyJoin)
	testframework.TFramework.RegTestCase("ShardPeerApprove", TestShardPeerApproveJoin)
	testframework.TFramework.RegTestCase("ShardActivate", TestShardActivate)
	testframework.TFramework.RegTestCase("ShardInfoQuery", TestShardInfoQuery)
	testframework.TFramework.RegTestCase("ShardPeerExit", TestShardPeerExit)
	testframework.TFramework.RegTestCase("NotifyParentCommitDpos", TestNotifyRootCommitDpos)
	testframework.TFramework.RegTestCase("NotifyShardRootCommitDpos", TestNotifyShardCommitDpos)
	testframework.TFramework.RegTestCase("ShardRetryCommitDpos", TestShardRetryCommitDpos)
	testframework.TFramework.RegTestCase("GetShardCommitDposInfo", TestGetShardCommitDposInfo)
	testframework.TFramework.RegTestCase("UpdateShardConfig", TestUpdateShardConfig)

	testframework.TFramework.RegTestCase("ShardAssetInit", TestAssetInit)
	testframework.TFramework.RegTestCase("XShardTransferOep4", TestXShardTransferOep4)
	testframework.TFramework.RegTestCase("XShardTransferOng", TestXShardTransferOng)
	testframework.TFramework.RegTestCase("XShardTransferOngRetry", TestXShardTransferOngRetry)
	testframework.TFramework.RegTestCase("XShardTransferOep4Retry", TestXShardTransferOep4Retry)
	testframework.TFramework.RegTestCase("ShardGetPendingTransfer", TestGetPendingTransfer)
	testframework.TFramework.RegTestCase("ShardGetTransferDetail", TestGetTransferDetail)
	testframework.TFramework.RegTestCase("ShardGetSupplyInfo", TestGetSupplyInfo)
	testframework.TFramework.RegTestCase("ShardGetOep4Balance", TestGetOep4Balance)

	testframework.TFramework.RegTestCase("ChangeContractMetaData", TestChangeMetaData)

	testframework.TFramework.RegTestCase("ShardChangePeerMaxAuth", TestShardChangePeerMaxAuthorization)
	testframework.TFramework.RegTestCase("ShardChangePeerProportion", TestShardChangePeerProportion)
	testframework.TFramework.RegTestCase("ShardUserWithdrawOng", TestShardUserWithdrawOng)
	testframework.TFramework.RegTestCase("ShardUserStake", TestShardUserStake)
	testframework.TFramework.RegTestCase("ShardUserUnfreezeStake", TestShardUserUnfreezeStake)
	testframework.TFramework.RegTestCase("ShardUserWithdrawStake", TestShardUserWithdrawStake)
	testframework.TFramework.RegTestCase("ShardUserWithdrawFee", TestShardUserWithdrawFee)
	testframework.TFramework.RegTestCase("ShardAddInitPos", TestAddInitPos)
	testframework.TFramework.RegTestCase("ShardReduceInitPos", TestReduceInitPos)
	testframework.TFramework.RegTestCase("ShardQueryView", TestGetShardView)
	testframework.TFramework.RegTestCase("ShardQueryPeerInfo", TestGetShardPeerInfo)
	testframework.TFramework.RegTestCase("ShardQueryUserInfo", TestGetShardUserInfo)

	testframework.TFramework.RegTestCase("ShardHotelInit", TestShardHotelInit)
	testframework.TFramework.RegTestCase("ShardHotelQuery", TestShardHotelQuery)
	testframework.TFramework.RegTestCase("ShardHotelReserve", TestShardHotelReserve)
	testframework.TFramework.RegTestCase("ShardHotelCheckout", TestShardHotelCheckout)
	testframework.TFramework.RegTestCase("ShardHotelReserve2", TestShardHotelReserve2)
	testframework.TFramework.RegTestCase("ShardHotelCheckout2", TestShardHotelCheckout2)
}
