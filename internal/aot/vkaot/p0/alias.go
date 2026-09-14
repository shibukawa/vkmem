package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	_ "unsafe"
)
//go:linkname F_ACLAddCommandCategory github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLAddCommandCategory
func F_ACLAddCommandCategory(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_ACLCreateUnlinkedUser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLCreateUnlinkedUser
func F_ACLCreateUnlinkedUser(m *base.Module) int32
//go:linkname F_ACLFreeUserAndKillClients github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLFreeUserAndKillClients
func F_ACLFreeUserAndKillClients(m *base.Module, l0 int32)
//go:linkname F_ACLSetSelector github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLSetSelector
func F_ACLSetSelector(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ACLDescribeUser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLDescribeUser
func F_ACLDescribeUser(m *base.Module, l0 int32) int32
//go:linkname F_ACLHashPassword github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLHashPassword
func F_ACLHashPassword(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aclCreateSelectorFromOpSet github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aclCreateSelectorFromOpSet
func F_aclCreateSelectorFromOpSet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLUserGetRootSelector github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLUserGetRootSelector
func F_ACLUserGetRootSelector(m *base.Module, l0 int32) int32
//go:linkname F_addACLLogEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addACLLogEntry
func F_addACLLogEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ACLCheckAllPerm github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLCheckAllPerm
func F_ACLCheckAllPerm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLStringSetUser github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLStringSetUser
func F_ACLStringSetUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ACLUpdateDefaultUserPassword github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLUpdateDefaultUserPassword
func F_ACLUpdateDefaultUserPassword(m *base.Module, l0 int32)
//go:linkname F_ACLSetSelectorCommandBitsForCategory github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLSetSelectorCommandBitsForCategory
func F_ACLSetSelectorCommandBitsForCategory(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_listCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listCreate
func F_listCreate(m *base.Module) int32
//go:linkname F_listRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listRelease
func F_listRelease(m *base.Module, l0 int32)
//go:linkname F_listAddNodeHead github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listAddNodeHead
func F_listAddNodeHead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listAddNodeTail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listAddNodeTail
func F_listAddNodeTail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listDelNode
func F_listDelNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listUnlinkNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listUnlinkNode
func F_listUnlinkNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listRewind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listRewind
func F_listRewind(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listNext
func F_listNext(m *base.Module, l0 int32) int32
//go:linkname F_listDup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listDup
func F_listDup(m *base.Module, l0 int32) int32
//go:linkname F_listSearchKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listSearchKey
func F_listSearchKey(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aeCreateFileEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeCreateFileEvent
func F_aeCreateFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_aeDeleteFileEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aeDeleteFileEvent
func F_aeDeleteFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_aeCreateTimeEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aeCreateTimeEvent
func F_aeCreateTimeEvent(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int64
//go:linkname F_aeProcessEvents github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeProcessEvents
func F_aeProcessEvents(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aeWait github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeWait
func F_aeWait(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_anetRecvTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_anetRecvTimeout
func F_anetRecvTimeout(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_anetTcpNonBlockConnect github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_anetTcpNonBlockConnect
func F_anetTcpNonBlockConnect(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_anetListen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetListen
func F_anetListen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_anetTcpAccept github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetTcpAccept
func F_anetTcpAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetFdToString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_anetFdToString
func F_anetFdToString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetPipe github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetPipe
func F_anetPipe(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getAofManifestAsString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getAofManifestAsString
func F_getAofManifestAsString(m *base.Module, l0 int32) int32
//go:linkname F_aofLoadManifestFromFile github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aofLoadManifestFromFile
func F_aofLoadManifestFromFile(m *base.Module, l0 int32) int32
//go:linkname F_aofManifestDup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aofManifestDup
func F_aofManifestDup(m *base.Module, l0 int32) int32
//go:linkname F_getAppendOnlyFileSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getAppendOnlyFileSize
func F_getAppendOnlyFileSize(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_stopAppendOnly github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_stopAppendOnly
func F_stopAppendOnly(m *base.Module)
//go:linkname F_rewriteAppendOnlyFileBackground github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteAppendOnlyFileBackground
func F_rewriteAppendOnlyFileBackground(m *base.Module) int32
//go:linkname F_restartAOFWithSyncRdb github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_restartAOFWithSyncRdb
func F_restartAOFWithSyncRdb(m *base.Module) int32
//go:linkname F_genAofTimestampAnnotationIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genAofTimestampAnnotationIfNeeded
func F_genAofTimestampAnnotationIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_rioWriteBulkObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkObject
func F_rioWriteBulkObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteListObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteListObject
func F_rewriteListObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rewriteSetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteSetObject
func F_rewriteSetObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rewriteSortedSetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteSortedSetObject
func F_rewriteSortedSetObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rewriteHashObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteHashObject
func F_rewriteHashObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_bioExecuteJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bioExecuteJob
func F_bioExecuteJob(m *base.Module, l0 int32)
//go:linkname F_bioCreateLazyFreeJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bioCreateLazyFreeJob
func F_bioCreateLazyFreeJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_allocBioJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_allocBioJob
func F_allocBioJob(m *base.Module, l0 int32) int32
//go:linkname F_bioCreateCloseJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bioCreateCloseJob
func F_bioCreateCloseJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bioCreateCloseAofJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bioCreateCloseAofJob
func F_bioCreateCloseAofJob(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_bioCreateSaveRDBToDiskJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bioCreateSaveRDBToDiskJob
func F_bioCreateSaveRDBToDiskJob(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bioDrainWorker github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bioDrainWorker
func F_bioDrainWorker(m *base.Module, l0 int32)
//go:linkname F_initClientBlockingState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initClientBlockingState
func F_initClientBlockingState(m *base.Module, l0 int32)
//go:linkname F_blockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockClient
func F_blockClient(m *base.Module, l0 int32, l1 int32)
//go:linkname F_updateStatsOnUnblock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateStatsOnUnblock
func F_updateStatsOnUnblock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_unblockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unblockClient
func F_unblockClient(m *base.Module, l0 int32, l1 int32)
//go:linkname F_blockForKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockForKeys
func F_blockForKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32)
//go:linkname F_signalDeletedKeyAsReady github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_signalDeletedKeyAsReady
func F_signalDeletedKeyAsReady(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_blockClientShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_blockClientShutdown
func F_blockClientShutdown(m *base.Module, l0 int32)
//go:linkname F_unblockClientOnTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unblockClientOnTimeout
func F_unblockClientOnTimeout(m *base.Module, l0 int32)
//go:linkname F_callReplyGetLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_callReplyGetLongLong
func F_callReplyGetLongLong(m *base.Module, l0 int32) int64
//go:linkname F_callReplyGetBigNumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_callReplyGetBigNumber
func F_callReplyGetBigNumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_connTypeOfCluster github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_connTypeOfCluster
func F_connTypeOfCluster(m *base.Module) int32
//go:linkname F_createDumpPayload github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createDumpPayload
func F_createDumpPayload(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_migrateGetSocket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_migrateGetSocket
func F_migrateGetSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_migrateCloseSocket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_migrateCloseSocket
func F_migrateCloseSocket(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterRedirectBlockedClientIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterRedirectBlockedClientIfNeeded
func F_clusterRedirectBlockedClientIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_humanNodename github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_humanNodename
func F_humanNodename(m *base.Module, l0 int32) int32
//go:linkname F_clusterAddNodeToShard github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterAddNodeToShard
func F_clusterAddNodeToShard(m *base.Module, l0 int32, l1 int32)
//go:linkname F_createClusterNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createClusterNode
func F_createClusterNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterGetNodesInMyShard github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterGetNodesInMyShard
func F_clusterGetNodesInMyShard(m *base.Module, l0 int32) int32
//go:linkname F_clusterNodeAddReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeAddReplica
func F_clusterNodeAddReplica(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterSaveConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSaveConfig
func F_clusterSaveConfig(m *base.Module, l0 int32) int32
//go:linkname F_updateSdsExtensionField github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateSdsExtensionField
func F_updateSdsExtensionField(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterUpdateMyselfAvailabilityZone github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterUpdateMyselfAvailabilityZone
func F_clusterUpdateMyselfAvailabilityZone(m *base.Module)
//go:linkname F_clusterDelSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterDelSlot
func F_clusterDelSlot(m *base.Module, l0 int32) int32
//go:linkname F_clusterDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterDelNode
func F_clusterDelNode(m *base.Module, l0 int32)
//go:linkname F_clusterSetNodeAsPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSetNodeAsPrimary
func F_clusterSetNodeAsPrimary(m *base.Module, l0 int32)
//go:linkname F_createClusterMsgSendBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createClusterMsgSendBlock
func F_createClusterMsgSendBlock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_freeClusterLink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClusterLink
func F_freeClusterLink(m *base.Module, l0 int32)
//go:linkname F_setClusterNodeToInboundClusterLink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setClusterNodeToInboundClusterLink
func F_setClusterNodeToInboundClusterLink(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterHandleConfigEpochCollision github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterHandleConfigEpochCollision
func F_clusterHandleConfigEpochCollision(m *base.Module, l0 int32)
//go:linkname F_clusterNodeIsPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeIsPrimary
func F_clusterNodeIsPrimary(m *base.Module, l0 int32) int32
//go:linkname F_markNodeAsFailing github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_markNodeAsFailing
func F_markNodeAsFailing(m *base.Module, l0 int32)
//go:linkname F_clusterBroadcastMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterBroadcastMessage
func F_clusterBroadcastMessage(m *base.Module, l0 int32)
//go:linkname F_clusterSendFail github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendFail
func F_clusterSendFail(m *base.Module, l0 int32)
//go:linkname F_clearNodeFailureIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clearNodeFailureIfNeeded
func F_clearNodeFailureIfNeeded(m *base.Module, l0 int32)
//go:linkname F_clusterProcessGossipSection github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterProcessGossipSection
func F_clusterProcessGossipSection(m *base.Module, l0 int32, l1 int32)
//go:linkname F_representClusterNodeFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_representClusterNodeFlags
func F_representClusterNodeFlags(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_nodeUpdateAddressIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_nodeUpdateAddressIfNeeded
func F_nodeUpdateAddressIfNeeded(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clusterUpdateSlotsConfigWith github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterUpdateSlotsConfigWith
func F_clusterUpdateSlotsConfigWith(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_delKeysInSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_delKeysInSlot
func F_delKeysInSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_updateShardId github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateShardId
func F_updateShardId(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterNodeGetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeGetPrimary
func F_clusterNodeGetPrimary(m *base.Module, l0 int32) int32
//go:linkname F_clusterIsValidPacket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterIsValidPacket
func F_clusterIsValidPacket(m *base.Module, l0 int32) int32
//go:linkname F_clusterSendPing github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSendPing
func F_clusterSendPing(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterProcessModulePacket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterProcessModulePacket
func F_clusterProcessModulePacket(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterMoveNodeSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterMoveNodeSlots
func F_clusterMoveNodeSlots(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_clusterSendUpdate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendUpdate
func F_clusterSendUpdate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterSendMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendMessage
func F_clusterSendMessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterSendFailoverAuth github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendFailoverAuth
func F_clusterSendFailoverAuth(m *base.Module, l0 int32)
//go:linkname F_clusterCreatePublishMsgBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterCreatePublishMsgBlock
func F_clusterCreatePublishMsgBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_clusterSendModule github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendModule
func F_clusterSendModule(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32)
//go:linkname F_clusterNodeIterNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterNodeIterNext
func F_clusterNodeIterNext(m *base.Module, l0 int32) int32
//go:linkname F_clusterAllReplicasThinkPrimaryIsFail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterAllReplicasThinkPrimaryIsFail
func F_clusterAllReplicasThinkPrimaryIsFail(m *base.Module) int32
//go:linkname F_clusterNodeIp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeIp
func F_clusterNodeIp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addNodeDetailsToShardReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addNodeDetailsToShardReply
func F_addNodeDetailsToShardReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_genClusterInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_genClusterInfoString
func F_genClusterInfoString(m *base.Module, l0 int32) int32
//go:linkname F_getMyClusterNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getMyClusterNode
func F_getMyClusterNode(m *base.Module) int32
//go:linkname F_finishSlotMigrationJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_finishSlotMigrationJob
func F_finishSlotMigrationJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_updateSlotMigrationJobState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateSlotMigrationJobState
func F_updateSlotMigrationJobState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_performSlotImportJobFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_performSlotImportJobFailover
func F_performSlotImportJobFailover(m *base.Module, l0 int32)
//go:linkname F_clusterCleanSlotImportsBeforeLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterCleanSlotImportsBeforeLoad
func F_clusterCleanSlotImportsBeforeLoad(m *base.Module)
//go:linkname F_slotMigrationJobSendAuth github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_slotMigrationJobSendAuth
func F_slotMigrationJobSendAuth(m *base.Module, l0 int32)
//go:linkname F_slotExportJobBeginSnapshotToTargetSocket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_slotExportJobBeginSnapshotToTargetSocket
func F_slotExportJobBeginSnapshotToTargetSocket(m *base.Module, l0 int32) int32
//go:linkname F_slotExportTryDoPause github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_slotExportTryDoPause
func F_slotExportTryDoPause(m *base.Module, l0 int32) int32
//go:linkname F_checkSlotExportOwnership github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_checkSlotExportOwnership
func F_checkSlotExportOwnership(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterIsAnySlotExporting github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterIsAnySlotExporting
func F_clusterIsAnySlotExporting(m *base.Module) int32
//go:linkname F_slotExportBeginStreaming github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_slotExportBeginStreaming
func F_slotExportBeginStreaming(m *base.Module, l0 int32)
//go:linkname F_killSlotMigrationChild github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_killSlotMigrationChild
func F_killSlotMigrationChild(m *base.Module)
//go:linkname F_clusterGetTotalSlotExportBufferMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterGetTotalSlotExportBufferMemory
func F_clusterGetTotalSlotExportBufferMemory(m *base.Module) int32
//go:linkname F_clusterSlotStatsUpdateNetworkBytesOutForReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSlotStatsUpdateNetworkBytesOutForReplication
func F_clusterSlotStatsUpdateNetworkBytesOutForReplication(m *base.Module, l0 int64)
//go:linkname F_appendServerSaveParams github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_appendServerSaveParams
func F_appendServerSaveParams(m *base.Module, l0 int64, l1 int32)
//go:linkname F_resetServerSaveParams github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resetServerSaveParams
func F_resetServerSaveParams(m *base.Module)
//go:linkname F_loadServerConfigFromString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_loadServerConfigFromString
func F_loadServerConfigFromString(m *base.Module, l0 int32)
//go:linkname F_performModuleConfigSetFromName github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_performModuleConfigSetFromName
func F_performModuleConfigSetFromName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performModuleConfigSetDefaultFromName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_performModuleConfigSetDefaultFromName
func F_performModuleConfigSetDefaultFromName(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_restoreBackupConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_restoreBackupConfig
func F_restoreBackupConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_rewriteConfigReadOldFile github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteConfigReadOldFile
func F_rewriteConfigReadOldFile(m *base.Module, l0 int32) int32
//go:linkname F_rewriteConfigRewriteLine github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteConfigRewriteLine
func F_rewriteConfigRewriteLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rewriteConfigSdsOption github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteConfigSdsOption
func F_rewriteConfigSdsOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_rewriteConfigGetContentFromState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteConfigGetContentFromState
func F_rewriteConfigGetContentFromState(m *base.Module, l0 int32) int32
//go:linkname F_rewriteConfigOverwriteFile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteConfigOverwriteFile
func F_rewriteConfigOverwriteFile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteConfigBindOption github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteConfigBindOption
func F_rewriteConfigBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_initConfigValues github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initConfigValues
func F_initConfigValues(m *base.Module)
//go:linkname F_addModuleEnumConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addModuleEnumConfig
func F_addModuleEnumConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_connectionByType github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connectionByType
func F_connectionByType(m *base.Module, l0 int32) int32
//go:linkname F_connectionTypeTcp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connectionTypeTcp
func F_connectionTypeTcp(m *base.Module) int32
//go:linkname F_connectionTypeTls github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_connectionTypeTls
func F_connectionTypeTls(m *base.Module) int32
//go:linkname F_getListensInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getListensInfoString
func F_getListensInfoString(m *base.Module, l0 int32) int32
//go:linkname F_crc16 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_crc16
func F_crc16(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKey
func F_lookupKey(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_deleteExpiredKeyAndPropagateWithDictIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_deleteExpiredKeyAndPropagateWithDictIndex
func F_deleteExpiredKeyAndPropagateWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lookupKeyReadWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKeyReadWithFlags
func F_lookupKeyReadWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookupKeyWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupKeyWrite
func F_lookupKeyWrite(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupKeyReadOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupKeyReadOrReply
func F_lookupKeyReadOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookupKeyWriteOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKeyWriteOrReply
func F_lookupKeyWriteOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dbUpdateObjectWithVolatileItemsTracking github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbUpdateObjectWithVolatileItemsTracking
func F_dbUpdateObjectWithVolatileItemsTracking(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dbAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbAdd
func F_dbAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dbAddInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbAddInternal
func F_dbAddInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_dbReplaceValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbReplaceValue
func F_dbReplaceValue(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setKey
func F_setKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_dbRandomKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbRandomKey
func F_dbRandomKey(m *base.Module, l0 int32) int32
//go:linkname F_dbGenericDeleteWithDictIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbGenericDeleteWithDictIndex
func F_dbGenericDeleteWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_dbGenericDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbGenericDelete
func F_dbGenericDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dbDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbDelete
func F_dbDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_discardTempDb github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_discardTempDb
func F_discardTempDb(m *base.Module, l0 int32)
//go:linkname F_flushAllDataAndResetRDB github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_flushAllDataAndResetRDB
func F_flushAllDataAndResetRDB(m *base.Module, l0 int32)
//go:linkname F_delGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_delGenericCommand
func F_delGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_renameGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_renameGenericCommand
func F_renameGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scanDatabaseForReadyKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scanDatabaseForReadyKeys
func F_scanDatabaseForReadyKeys(m *base.Module, l0 int32)
//go:linkname F_dbFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbFind
func F_dbFind(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_propagateDeletion github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_propagateDeletion
func F_propagateDeletion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deleteExpiredKeyFromOverwriteAndPropagate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_deleteExpiredKeyFromOverwriteAndPropagate
func F_deleteExpiredKeyFromOverwriteAndPropagate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_keyIsExpired github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_keyIsExpired
func F_keyIsExpired(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getKeysFromCommandWithSpecs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getKeysFromCommandWithSpecs
func F_getKeysFromCommandWithSpecs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_genericGetKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericGetKeys
func F_genericGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_xorDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_xorDigest
func F_xorDigest(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_xorObjectDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_xorObjectDigest
func F_xorObjectDigest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__serverPanic_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverPanic_1
func F__serverPanic_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__serverAssert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverAssert
func F__serverAssert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bugReportEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bugReportEnd
func F_bugReportEnd(m *base.Module, l0 int32, l1 int32)
//go:linkname F_printCrashReport github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_printCrashReport
func F_printCrashReport(m *base.Module)
//go:linkname F__serverAssertWithInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverAssertWithInfo
func F__serverAssertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_getArgvReprString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getArgvReprString
func F_getArgvReprString(m *base.Module, l0 int32) int32
//go:linkname F_debugDelay github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_debugDelay
func F_debugDelay(m *base.Module, l0 int32)
//go:linkname F_dictCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictCreate
func F_dictCreate(m *base.Module, l0 int32) int32
//go:linkname F_dictResizeWithOptionalCheck github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictResizeWithOptionalCheck
func F_dictResizeWithOptionalCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictAdd
func F_dictAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictAddRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictAddRaw
func F_dictAddRaw(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictBucketRehash github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictBucketRehash
func F_dictBucketRehash(m *base.Module, l0 int32, l1 int64)
//go:linkname F_dictGetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictGetVal
func F_dictGetVal(m *base.Module, l0 int32) int32
//go:linkname F_dictAddOrFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictAddOrFind
func F_dictAddOrFind(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictDelete
func F_dictDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictShrinkIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictShrinkIfNeeded
func F_dictShrinkIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_dictFreeUnlinkedEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictFreeUnlinkedEntry
func F_dictFreeUnlinkedEntry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dictRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictRelease
func F_dictRelease(m *base.Module, l0 int32)
//go:linkname F_dictFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictFind
func F_dictFind(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictFetchValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictFetchValue
func F_dictFetchValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictMemUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictMemUsage
func F_dictMemUsage(m *base.Module, l0 int32) int32
//go:linkname F_dictResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictResetIterator
func F_dictResetIterator(m *base.Module, l0 int32)
//go:linkname F_dictGetSafeIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictGetSafeIterator
func F_dictGetSafeIterator(m *base.Module, l0 int32) int32
//go:linkname F_dictNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictNext
func F_dictNext(m *base.Module, l0 int32) int32
//go:linkname F_dictReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictReleaseIterator
func F_dictReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_entrySetExpiry github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_entrySetExpiry
func F_entrySetExpiry(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_entryFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_entryFree
func F_entryFree(m *base.Module, l0 int32)
//go:linkname F_evalRemoveScriptsFromEngine github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_evalRemoveScriptsFromEngine
func F_evalRemoveScriptsFromEngine(m *base.Module, l0 int32)
//go:linkname F_evalGetCommandFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_evalGetCommandFlags
func F_evalGetCommandFlags(m *base.Module, l0 int32, l1 int64) int64
//go:linkname F_evalGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_evalGenericCommand
func F_evalGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_evalMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_evalMemory
func F_evalMemory(m *base.Module) int32
//go:linkname F_evictionPoolPopulate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_evictionPoolPopulate
func F_evictionPoolPopulate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rememberReplicaKeyWithExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rememberReplicaKeyWithExpire
func F_rememberReplicaKeyWithExpire(m *base.Module, l0 int32, l1 int32)
//go:linkname F_flushReplicaKeysWithExpireList github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_flushReplicaKeysWithExpireList
func F_flushReplicaKeysWithExpireList(m *base.Module, l0 int32)
//go:linkname F_parseExtendedExpireArgumentsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_parseExtendedExpireArgumentsOrReply
func F_parseExtendedExpireArgumentsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_functionsLibCtxFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_functionsLibCtxFree
func F_functionsLibCtxFree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_functionReset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_functionReset
func F_functionReset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_functionsLibCtxSwapWithCurrent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_functionsLibCtxSwapWithCurrent
func F_functionsLibCtxSwapWithCurrent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_functionsRemoveLibFromEngine github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_functionsRemoveLibFromEngine
func F_functionsRemoveLibFromEngine(m *base.Module, l0 int32)
//go:linkname F_libraryUnlink github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_libraryUnlink
func F_libraryUnlink(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fcallGetCommandFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fcallGetCommandFlags
func F_fcallGetCommandFlags(m *base.Module, l0 int32, l1 int64) int64
//go:linkname F_functionsCreateWithLibraryCtx github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_functionsCreateWithLibraryCtx
func F_functionsCreateWithLibraryCtx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_functionsMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_functionsMemory
func F_functionsMemory(m *base.Module) int32
//go:linkname F_geoArrayCleanup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geoArrayCleanup
func F_geoArrayCleanup(m *base.Module, l0 int32)
//go:linkname F_extractLongLatOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_extractLongLatOrReply
func F_extractLongLatOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_longLatFromMemberOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_longLatFromMemberOrReply
func F_longLatFromMemberOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_extractUnitOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_extractUnitOrReply
func F_extractUnitOrReply(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_extractBoxOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_extractBoxOrReply
func F_extractBoxOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_membersOfAllNeighbors github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_membersOfAllNeighbors
func F_membersOfAllNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_geohashEncodeType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashEncodeType
func F_geohashEncodeType(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32) int32
//go:linkname F_geohashDecodeToLongLatType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashDecodeToLongLatType
func F_geohashDecodeToLongLatType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_geohashCalculateAreasByShapeWGS84 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashCalculateAreasByShapeWGS84
func F_geohashCalculateAreasByShapeWGS84(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashtableCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCreate
func F_hashtableCreate(m *base.Module, l0 int32) int32
//go:linkname F_hashtableRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableRelease
func F_hashtableRelease(m *base.Module, l0 int32)
//go:linkname F_hashtableSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableSize
func F_hashtableSize(m *base.Module, l0 int32) int32
//go:linkname F_hashtableMemUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableMemUsage
func F_hashtableMemUsage(m *base.Module, l0 int32) int32
//go:linkname F_hashtableResumeAutoShrink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableResumeAutoShrink
func F_hashtableResumeAutoShrink(m *base.Module, l0 int32)
//go:linkname F_rehashStep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rehashStep
func F_rehashStep(m *base.Module, l0 int32)
//go:linkname F_rehashingCompleted github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rehashingCompleted
func F_rehashingCompleted(m *base.Module, l0 int32)
//go:linkname F_hashtableFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableFind
func F_hashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findBucket_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_findBucket_1
func F_findBucket_1(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hashtableAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableAdd
func F_hashtableAdd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableFindPositionForInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableFindPositionForInsert
func F_hashtableFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fillBucketHole github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fillBucketHole
func F_fillBucketHole(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moveEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moveEntry
func F_moveEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hashtableDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableDelete
func F_hashtableDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableReplaceReallocatedEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableReplaceReallocatedEntry
func F_hashtableReplaceReallocatedEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableTwoPhasePopFindRef github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableTwoPhasePopFindRef
func F_hashtableTwoPhasePopFindRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableScanDefrag github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableScanDefrag
func F_hashtableScanDefrag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_hashtableCleanupIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCleanupIterator
func F_hashtableCleanupIterator(m *base.Module, l0 int32)
//go:linkname F_hashtableCreateIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCreateIterator
func F_hashtableCreateIterator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableNext
func F_hashtableNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableFairRandomEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableFairRandomEntry
func F_hashtableFairRandomEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hllSparseToDense github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hllSparseToDense
func F_hllSparseToDense(m *base.Module, l0 int32) int32
//go:linkname F_hllAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hllAdd
func F_hllAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hllSparseSet github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hllSparseSet
func F_hllSparseSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_intsetNew github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_intsetNew
func F_intsetNew(m *base.Module) int32
//go:linkname F_intsetFree github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_intsetFree
func F_intsetFree(m *base.Module, l0 int32)
//go:linkname F_IOThreadsAfterSleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_IOThreadsAfterSleep
func F_IOThreadsAfterSleep(m *base.Module, l0 int32)
//go:linkname F_sendToMainThread github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendToMainThread
func F_sendToMainThread(m *base.Module, l0 int32, l1 int32)
//go:linkname F_kvstoreEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreEmpty
func F_kvstoreEmpty(m *base.Module, l0 int32, l1 int32)
//go:linkname F_kvstoreRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreRelease
func F_kvstoreRelease(m *base.Module, l0 int32)
//go:linkname F_kvstoreSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreSize
func F_kvstoreSize(m *base.Module, l0 int32) int64
//go:linkname F_kvstoreScan github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreScan
func F_kvstoreScan(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64
//go:linkname F_kvstoreExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreExpand
func F_kvstoreExpand(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_kvstoreGetFairRandomHashtableIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreGetFairRandomHashtableIndex
func F_kvstoreGetFairRandomHashtableIndex(m *base.Module, l0 int32) int32
//go:linkname F_kvstoreIteratorInit github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorInit
func F_kvstoreIteratorInit(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreIteratorGetCurrentHashtableIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorGetCurrentHashtableIndex
func F_kvstoreIteratorGetCurrentHashtableIndex(m *base.Module, l0 int32) int32
//go:linkname F_kvstoreIteratorNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorNext
func F_kvstoreIteratorNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreHashtableSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableSize
func F_kvstoreHashtableSize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreGetHashtableIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreGetHashtableIterator
func F_kvstoreGetHashtableIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreReleaseHashtableIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreReleaseHashtableIterator
func F_kvstoreReleaseHashtableIterator(m *base.Module, l0 int32)
//go:linkname F_kvstoreHashtableIteratorNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableIteratorNext
func F_kvstoreHashtableIteratorNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreHashtableRandomEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableRandomEntry
func F_kvstoreHashtableRandomEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreHashtableFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableFind
func F_kvstoreHashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_kvstoreHashtableFindRef github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableFindRef
func F_kvstoreHashtableFindRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreHashtableAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableAdd
func F_kvstoreHashtableAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cumulativeKeyCountAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_cumulativeKeyCountAdd
func F_cumulativeKeyCountAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_kvstoreHashtableFindPositionForInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableFindPositionForInsert
func F_kvstoreHashtableFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_kvstoreHashtableInsertAtPosition github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableInsertAtPosition
func F_kvstoreHashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_latencyAddSample github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_latencyAddSample
func F_latencyAddSample(m *base.Module, l0 int32, l1 int64)
//go:linkname F_fillCommandCDF github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fillCommandCDF
func F_fillCommandCDF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lazyfreeGetFreeEffort github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lazyfreeGetFreeEffort
func F_lazyfreeGetFreeEffort(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emptyDbAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emptyDbAsync
func F_emptyDbAsync(m *base.Module, l0 int32)
//go:linkname F_freeEvalScriptsAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeEvalScriptsAsync
func F_freeEvalScriptsAsync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lpNew github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpNew
func F_lpNew(m *base.Module, l0 int32) int32
//go:linkname F_lpNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpNext
func F_lpNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpFirst github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpFirst
func F_lpFirst(m *base.Module, l0 int32) int32
//go:linkname F_lpLength github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpLength
func F_lpLength(m *base.Module, l0 int32) int32
//go:linkname F_lpGet github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpGet
func F_lpGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpGetWithSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpGetWithSize
func F_lpGetWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpGetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpGetValue
func F_lpGetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpInsertString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpInsertString
func F_lpInsertString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_lpPrepend github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpPrepend
func F_lpPrepend(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpAppend github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpAppend
func F_lpAppend(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpAppendInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpAppendInteger
func F_lpAppendInteger(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_lpReplace github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpReplace
func F_lpReplace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpReplaceInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpReplaceInteger
func F_lpReplaceInteger(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_lpDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpDelete
func F_lpDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpDeleteRangeWithEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpDeleteRangeWithEntry
func F_lpDeleteRangeWithEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpSeek
func F_lpSeek(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpDup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpDup
func F_lpDup(m *base.Module, l0 int32) int32
//go:linkname F_lpValidateIntegrity github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpValidateIntegrity
func F_lpValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpRandomPairs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpRandomPairs
func F_lpRandomPairs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_lpRandomPairsUnique github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpRandomPairsUnique
func F_lpRandomPairsUnique(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lolwutUnstableCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lolwutUnstableCommand
func F_lolwutUnstableCommand(m *base.Module, l0 int32)
//go:linkname F_lwCreateCanvas github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lwCreateCanvas
func F_lwCreateCanvas(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lwDrawPixel github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lwDrawPixel
func F_lwDrawPixel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_lwGetPixel github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lwGetPixel
func F_lwGetPixel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lolwut5Command github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lolwut5Command
func F_lolwut5Command(m *base.Module, l0 int32)
//go:linkname F_generateSkyscraper github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_generateSkyscraper
func F_generateSkyscraper(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lolwut9Command github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lolwut9Command
func F_lolwut9Command(m *base.Module, l0 int32)
//go:linkname F_lfu_getFrequency github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lfu_getFrequency
func F_lfu_getFrequency(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lrulfu_updateClockAndPolicy github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lrulfu_updateClockAndPolicy
func F_lrulfu_updateClockAndPolicy(m *base.Module, l0 int64, l1 int32)
//go:linkname F_lzf_decompress github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lzf_decompress
func F_lzf_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freeClientModuleData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeClientModuleData
func F_freeClientModuleData(m *base.Module, l0 int32)
//go:linkname F_moduleLoadQueueEntryToLoadmoduleOptionStr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleLoadQueueEntryToLoadmoduleOptionStr
func F_moduleLoadQueueEntryToLoadmoduleOptionStr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleDelKeyIfEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleDelKeyIfEmpty
func F_moduleDelKeyIfEmpty(m *base.Module, l0 int32) int32
//go:linkname F_moduleFreeContext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleFreeContext
func F_moduleFreeContext(m *base.Module, l0 int32)
//go:linkname F_moduleCallCommandUnblockedHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleCallCommandUnblockedHandler
func F_moduleCallCommandUnblockedHandler(m *base.Module, l0 int32)
//go:linkname F_moduleScriptingEngineInitContext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleScriptingEngineInitContext
func F_moduleScriptingEngineInitContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleListIteratorSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleListIteratorSeek
func F_moduleListIteratorSeek(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_VM_ListPush github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_VM_ListPush
func F_VM_ListPush(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_VM_CallArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_VM_CallArgv
func F_VM_CallArgv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_moduleFromCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleFromCommand
func F_moduleFromCommand(m *base.Module, l0 int32) int32
//go:linkname F_moduleTypeDupOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleTypeDupOrReply
func F_moduleTypeDupOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_moduleAllModulesHandleReplAsyncLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleAllModulesHandleReplAsyncLoad
func F_moduleAllModulesHandleReplAsyncLoad(m *base.Module) int32
//go:linkname F_VM_SaveStringBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_VM_SaveStringBuffer
func F_VM_SaveStringBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_moduleLoadString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleLoadString
func F_moduleLoadString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_moduleFireServerEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleFireServerEvent
func F_moduleFireServerEvent(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_moduleCallClusterReceivers github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleCallClusterReceivers
func F_moduleCallClusterReceivers(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32)
//go:linkname F_moduleNotifyUserChanged github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleNotifyUserChanged
func F_moduleNotifyUserChanged(m *base.Module, l0 int32)
//go:linkname F_authenticateClientWithUser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_authenticateClientWithUser
func F_authenticateClientWithUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_modulesCollectInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_modulesCollectInfo
func F_modulesCollectInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleFireCommandRejectedEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleFireCommandRejectedEvent
func F_moduleFireCommandRejectedEvent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleNotifyKeyUnlink github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleNotifyKeyUnlink
func F_moduleNotifyKeyUnlink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_VM_GetCommandKeysWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_VM_GetCommandKeysWithFlags
func F_VM_GetCommandKeysWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_moduleInitPostOnLoadResolved github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleInitPostOnLoadResolved
func F_moduleInitPostOnLoadResolved(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_parseLoadexArguments github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseLoadexArguments
func F_parseLoadexArguments(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleUnloadInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleUnloadInternal
func F_moduleUnloadInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addReplyLoadedModules github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyLoadedModules
func F_addReplyLoadedModules(m *base.Module, l0 int32)
//go:linkname F_genModulesInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genModulesInfoString
func F_genModulesInfoString(m *base.Module, l0 int32) int32
//go:linkname F_setModuleUnsignedNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setModuleUnsignedNumericConfig
func F_setModuleUnsignedNumericConfig(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_getModuleBoolConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getModuleBoolConfig
func F_getModuleBoolConfig(m *base.Module, l0 int32) int32
//go:linkname F_getModuleEnumConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getModuleEnumConfig
func F_getModuleEnumConfig(m *base.Module, l0 int32) int32
//go:linkname F_getModuleUnsignedNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getModuleUnsignedNumericConfig
func F_getModuleUnsignedNumericConfig(m *base.Module, l0 int32) int64
//go:linkname F_addModuleConfigApply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addModuleConfigApply
func F_addModuleConfigApply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleConfigValidityCheck github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleConfigValidityCheck
func F_moduleConfigValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freeClientMultiState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeClientMultiState
func F_freeClientMultiState(m *base.Module, l0 int32)
//go:linkname F_unwatchAllKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unwatchAllKeys
func F_unwatchAllKeys(m *base.Module, l0 int32)
//go:linkname F_resetClientMultiState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resetClientMultiState
func F_resetClientMultiState(m *base.Module, l0 int32)
//go:linkname F_flagTransaction github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_flagTransaction
func F_flagTransaction(m *base.Module, l0 int32)
//go:linkname F_execCommandAbort github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_execCommandAbort
func F_execCommandAbort(m *base.Module, l0 int32, l1 int32)
//go:linkname F_touchWatchedKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_touchWatchedKey
func F_touchWatchedKey(m *base.Module, l0 int32, l1 int32)
//go:linkname F_multiStateMemOverhead github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_multiStateMemOverhead
func F_multiStateMemOverhead(m *base.Module, l0 int32) int32
//go:linkname F_mutexQueueCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mutexQueueCreate
func F_mutexQueueCreate(m *base.Module) int32
//go:linkname F_getStringObjectLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getStringObjectLen
func F_getStringObjectLen(m *base.Module, l0 int32) int32
//go:linkname F_sdslen_7 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdslen_7
func F_sdslen_7(m *base.Module, l0 int32) int32
//go:linkname F_linkClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_linkClient
func F_linkClient(m *base.Module, l0 int32)
//go:linkname F_createClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createClient
func F_createClient(m *base.Module, l0 int32) int32
//go:linkname F_beforeNextClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_beforeNextClient
func F_beforeNextClient(m *base.Module, l0 int32)
//go:linkname F_clientHasPendingReplies github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clientHasPendingReplies
func F_clientHasPendingReplies(m *base.Module, l0 int32) int32
//go:linkname F_releaseReplyReferences github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_releaseReplyReferences
func F_releaseReplyReferences(m *base.Module, l0 int32)
//go:linkname F_unlinkClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unlinkClient
func F_unlinkClient(m *base.Module, l0 int32)
//go:linkname F__addReplyPayloadToList github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__addReplyPayloadToList
func F__addReplyPayloadToList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__addReplyToBufferOrList github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__addReplyToBufferOrList
func F__addReplyToBufferOrList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_catClientInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_catClientInfoString
func F_catClientInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_addReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReply
func F_addReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_afterErrorReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_afterErrorReply
func F_afterErrorReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_commitDeferredReplyBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commitDeferredReplyBuffer
func F_commitDeferredReplyBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyOrErrorObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyOrErrorObject
func F_addReplyOrErrorObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyError
func F_addReplyError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorFormatInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyErrorFormatInternal
func F_addReplyErrorFormatInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_addReplyErrorFormat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyErrorFormat
func F_addReplyErrorFormat(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyErrorExpireTime github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyErrorExpireTime
func F_addReplyErrorExpireTime(m *base.Module, l0 int32)
//go:linkname F_addReplyStatus github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyStatus
func F_addReplyStatus(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyDeferredLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyDeferredLen
func F_addReplyDeferredLen(m *base.Module, l0 int32) int32
//go:linkname F_setDeferredReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredReply
func F_setDeferredReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_setDeferredArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredArrayLen
func F_setDeferredArrayLen(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setDeferredMapLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredMapLen
func F_setDeferredMapLen(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyBulkCBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyBulkCBuffer
func F_addReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyHumanLongDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyHumanLongDouble
func F_addReplyHumanLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64)
//go:linkname F_addReplyBulkLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyBulkLen
func F_addReplyBulkLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyLongLong
func F_addReplyLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addReplyMapLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyMapLen
func F_addReplyMapLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyPushLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyPushLen
func F_addReplyPushLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyNullArray github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyNullArray
func F_addReplyNullArray(m *base.Module, l0 int32)
//go:linkname F_addWritePreparedReplyBulkCBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addWritePreparedReplyBulkCBuffer
func F_addWritePreparedReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyBulkCString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyBulkCString
func F_addReplyBulkCString(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyBulkLongLong
func F_addReplyBulkLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addExtendedReplyHelp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addExtendedReplyHelp
func F_addExtendedReplyHelp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyHelp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyHelp
func F_addReplyHelp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplySubcommandSyntaxError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplySubcommandSyntaxError
func F_addReplySubcommandSyntaxError(m *base.Module, l0 int32)
//go:linkname F_initDeferredReplyBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initDeferredReplyBuffer
func F_initDeferredReplyBuffer(m *base.Module, l0 int32)
//go:linkname F_getClientPeerId github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getClientPeerId
func F_getClientPeerId(m *base.Module, l0 int32) int32
//go:linkname F_freeClientOriginalArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClientOriginalArgv
func F_freeClientOriginalArgv(m *base.Module, l0 int32)
//go:linkname F_disconnectReplicas github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_disconnectReplicas
func F_disconnectReplicas(m *base.Module)
//go:linkname F_clearClientConnectionState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clearClientConnectionState
func F_clearClientConnectionState(m *base.Module, l0 int32)
//go:linkname F_initSharedQueryBuf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initSharedQueryBuf
func F_initSharedQueryBuf(m *base.Module)
//go:linkname F__writeToClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__writeToClient
func F__writeToClient(m *base.Module, l0 int32) int32
//go:linkname F_handleQbLimitReached github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_handleQbLimitReached
func F_handleQbLimitReached(m *base.Module, l0 int32)
//go:linkname F_handleParseResults github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_handleParseResults
func F_handleParseResults(m *base.Module, l0 int32) int32
//go:linkname F_resetClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_resetClient
func F_resetClient(m *base.Module, l0 int32)
//go:linkname F_protectClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_protectClient
func F_protectClient(m *base.Module, l0 int32)
//go:linkname F_parseInlineBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_parseInlineBuffer
func F_parseInlineBuffer(m *base.Module, l0 int32)
//go:linkname F_parseMultibulkBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseMultibulkBuffer
func F_parseMultibulkBuffer(m *base.Module, l0 int32)
//go:linkname F_processPendingCommandAndInputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_processPendingCommandAndInputBuffer
func F_processPendingCommandAndInputBuffer(m *base.Module, l0 int32) int32
//go:linkname F_prefetchCommandQueueKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_prefetchCommandQueueKeys
func F_prefetchCommandQueueKeys(m *base.Module, l0 int32)
//go:linkname F_catClientInfoShortString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_catClientInfoShortString
func F_catClientInfoShortString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clientSetName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clientSetName
func F_clientSetName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parseClientFiltersOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseClientFiltersOrReply
func F_parseClientFiltersOrReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_freeClientFilter github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeClientFilter
func F_freeClientFilter(m *base.Module, l0 int32)
//go:linkname F_clientMatchesFlagFilter github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clientMatchesFlagFilter
func F_clientMatchesFlagFilter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clientMatchesIpFilter github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clientMatchesIpFilter
func F_clientMatchesIpFilter(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_updatePausedActions github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_updatePausedActions
func F_updatePausedActions(m *base.Module)
//go:linkname F_unpauseActions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unpauseActions
func F_unpauseActions(m *base.Module, l0 int32)
//go:linkname F_redactClientCommandArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_redactClientCommandArgument
func F_redactClientCommandArgument(m *base.Module, l0 int32, l1 int32)
//go:linkname F_replaceClientCommandVector github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replaceClientCommandVector
func F_replaceClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_flushReplicasOutputBuffers github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_flushReplicasOutputBuffers
func F_flushReplicasOutputBuffers(m *base.Module)
//go:linkname F_pauseActions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pauseActions
func F_pauseActions(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_isPausedActionsWithUpdate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_isPausedActionsWithUpdate
func F_isPausedActionsWithUpdate(m *base.Module, l0 int32) int32
//go:linkname F_ioThreadReadQueryFromClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ioThreadReadQueryFromClient
func F_ioThreadReadQueryFromClient(m *base.Module, l0 int32)
//go:linkname F_createObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createObject
func F_createObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_makeObjectShared github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_makeObjectShared
func F_makeObjectShared(m *base.Module, l0 int32) int32
//go:linkname F_createRawStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createRawStringObject
func F_createRawStringObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_createEmbeddedStringObjectWithKeyAndExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createEmbeddedStringObjectWithKeyAndExpire
func F_createEmbeddedStringObjectWithKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_createStringObjectFromSds github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createStringObjectFromSds
func F_createStringObjectFromSds(m *base.Module, l0 int32) int32
//go:linkname F_objectSetExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_objectSetExpire
func F_objectSetExpire(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_objectSetKeyAndExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_objectSetKeyAndExpire
func F_objectSetKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_freeSetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeSetObject
func F_freeSetObject(m *base.Module, l0 int32)
//go:linkname F_freeZsetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeZsetObject
func F_freeZsetObject(m *base.Module, l0 int32)
//go:linkname F_freeModuleObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeModuleObject
func F_freeModuleObject(m *base.Module, l0 int32)
//go:linkname F_freeStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeStringObject
func F_freeStringObject(m *base.Module, l0 int32)
//go:linkname F_createStringObjectFromLongLongWithOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createStringObjectFromLongLongWithOptions
func F_createStringObjectFromLongLongWithOptions(m *base.Module, l0 int64, l1 int32) int32
//go:linkname F_createStringObjectFromLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createStringObjectFromLongLong
func F_createStringObjectFromLongLong(m *base.Module, l0 int64) int32
//go:linkname F_createStringObjectFromLongLongWithSds github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createStringObjectFromLongLongWithSds
func F_createStringObjectFromLongLongWithSds(m *base.Module, l0 int64) int32
//go:linkname F_createStringObjectFromLongDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createStringObjectFromLongDouble
func F_createStringObjectFromLongDouble(m *base.Module, l0 int64, l1 int64, l2 int32) int32
//go:linkname F_dupStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dupStringObject
func F_dupStringObject(m *base.Module, l0 int32) int32
//go:linkname F_createSetListpackObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createSetListpackObject
func F_createSetListpackObject(m *base.Module) int32
//go:linkname F_createHashObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createHashObject
func F_createHashObject(m *base.Module) int32
//go:linkname F_createZsetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createZsetObject
func F_createZsetObject(m *base.Module) int32
//go:linkname F_createStreamObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createStreamObject
func F_createStreamObject(m *base.Module) int32
//go:linkname F_checkType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_checkType
func F_checkType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tryObjectEncoding github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_tryObjectEncoding
func F_tryObjectEncoding(m *base.Module, l0 int32) int32
//go:linkname F_getDecodedObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getDecodedObject
func F_getDecodedObject(m *base.Module, l0 int32) int32
//go:linkname F_compareStringObjectsWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_compareStringObjectsWithFlags
func F_compareStringObjectsWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_stringObjectLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_stringObjectLen
func F_stringObjectLen(m *base.Module, l0 int32) int32
//go:linkname F_getDoubleFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getDoubleFromObjectOrReply
func F_getDoubleFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getLongDoubleFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getLongDoubleFromObjectOrReply
func F_getLongDoubleFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getLongLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getLongLongFromObjectOrReply
func F_getLongLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getLongFromObjectOrReply
func F_getLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getPositiveLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getPositiveLongFromObjectOrReply
func F_getPositiveLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getIntFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getIntFromObjectOrReply
func F_getIntFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freeMemoryOverheadData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeMemoryOverheadData
func F_freeMemoryOverheadData(m *base.Module, l0 int32)
//go:linkname F_addReplyPubsubPatUnsubscribed github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyPubsubPatUnsubscribed
func F_addReplyPubsubPatUnsubscribed(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freeClientPubSubData github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClientPubSubData
func F_freeClientPubSubData(m *base.Module, l0 int32)
//go:linkname F_pubsubUnsubscribePattern github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubUnsubscribePattern
func F_pubsubUnsubscribePattern(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pubsubSubscribeChannel github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubSubscribeChannel
func F_pubsubSubscribeChannel(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pubsubShardUnsubscribeAllChannelsInSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pubsubShardUnsubscribeAllChannelsInSlot
func F_pubsubShardUnsubscribeAllChannelsInSlot(m *base.Module, l0 int32)
//go:linkname F_pubsubPublishMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubPublishMessage
func F_pubsubPublishMessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___quicklistInsertNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___quicklistInsertNode
func F___quicklistInsertNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_quicklistPushTail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistPushTail
func F_quicklistPushTail(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___quicklistDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___quicklistDelNode
func F___quicklistDelNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_quicklistDelIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistDelIndex
func F_quicklistDelIndex(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F___quicklistCompressNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___quicklistCompressNode
func F___quicklistCompressNode(m *base.Module, l0 int32) int32
//go:linkname F__quicklistSplitNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__quicklistSplitNode
func F__quicklistSplitNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___quicklistCreateNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___quicklistCreateNode
func F___quicklistCreateNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__quicklistMergeNodes github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__quicklistMergeNodes
func F__quicklistMergeNodes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__quicklistInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__quicklistInsert
func F__quicklistInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_quicklistInsertAfter github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_quicklistInsertAfter
func F_quicklistInsertAfter(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_quicklistDup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistDup
func F_quicklistDup(m *base.Module, l0 int32) int32
//go:linkname F_quicklistPopCustom github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistPopCustom
func F_quicklistPopCustom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_raxGenericInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxGenericInsert
func F_raxGenericInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_raxRemove github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxRemove
func F_raxRemove(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxTryInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxTryInsert
func F_raxTryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_raxFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxFind
func F_raxFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxRecursiveFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxRecursiveFree
func F_raxRecursiveFree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_raxFree github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxFree
func F_raxFree(m *base.Module, l0 int32)
//go:linkname F_raxSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxSeek
func F_raxSeek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxNext
func F_raxNext(m *base.Module, l0 int32) int32
//go:linkname F_raxRandomWalk github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxRandomWalk
func F_raxRandomWalk(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_raxStop github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxStop
func F_raxStop(m *base.Module, l0 int32)
//go:linkname F_raxSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxSize
func F_raxSize(m *base.Module, l0 int32) int64
//go:linkname F_rdbWriteRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbWriteRaw
func F_rdbWriteRaw(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadLenByRef github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadLenByRef
func F_rdbLoadLenByRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveLzfBlob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveLzfBlob
func F_rdbSaveLzfBlob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbSaveRawString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveRawString
func F_rdbSaveRawString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveStringObject
func F_rdbSaveStringObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbGenericLoadStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbGenericLoadStringObject
func F_rdbGenericLoadStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadBinaryDoubleValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadBinaryDoubleValue
func F_rdbLoadBinaryDoubleValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbSave github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSave
func F_rdbSave(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbSaveBackground github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSaveBackground
func F_rdbSaveBackground(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_stopLoading github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_stopLoading
func F_stopLoading(m *base.Module, l0 int32)
//go:linkname F_rdbLoadRioWithLoadingCtxScopedRdb github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadRioWithLoadingCtxScopedRdb
func F_rdbLoadRioWithLoadingCtxScopedRdb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoad
func F_rdbLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_serverBuildIdString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_serverBuildIdString
func F_serverBuildIdString(m *base.Module) int32
//go:linkname F_bg_unlink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bg_unlink
func F_bg_unlink(m *base.Module, l0 int32) int32
//go:linkname F_createReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createReplicationBacklog
func F_createReplicationBacklog(m *base.Module)
//go:linkname F_incrementalTrimReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_incrementalTrimReplicationBacklog
func F_incrementalTrimReplicationBacklog(m *base.Module, l0 int32)
//go:linkname F_freeReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeReplicationBacklog
func F_freeReplicationBacklog(m *base.Module)
//go:linkname F_removeReplicaFromPsyncWait github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_removeReplicaFromPsyncWait
func F_removeReplicaFromPsyncWait(m *base.Module, l0 int32)
//go:linkname F_feedReplicationBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_feedReplicationBuffer
func F_feedReplicationBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_replicationFeedReplicas github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationFeedReplicas
func F_replicationFeedReplicas(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyReplicationBacklog
func F_addReplyReplicationBacklog(m *base.Module, l0 int32, l1 int64) int64
//go:linkname F_startBgsaveForReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_startBgsaveForReplication
func F_startBgsaveForReplication(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_replicationUnsetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationUnsetPrimary
func F_replicationUnsetPrimary(m *base.Module)
//go:linkname F_initClientReplicationData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_initClientReplicationData
func F_initClientReplicationData(m *base.Module, l0 int32)
//go:linkname F_cancelReplicationHandshake github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_cancelReplicationHandshake
func F_cancelReplicationHandshake(m *base.Module, l0 int32) int32
//go:linkname F_restartAOFAfterSYNC github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_restartAOFAfterSYNC
func F_restartAOFAfterSYNC(m *base.Module)
//go:linkname F_freeClientReplicationData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeClientReplicationData
func F_freeClientReplicationData(m *base.Module, l0 int32)
//go:linkname F_removeRDBUsedToSyncReplicas github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_removeRDBUsedToSyncReplicas
func F_removeRDBUsedToSyncReplicas(m *base.Module)
//go:linkname F_replicationCreatePrimaryClientWithHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_replicationCreatePrimaryClientWithHandler
func F_replicationCreatePrimaryClientWithHandler(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replicationAttachToNewPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationAttachToNewPrimary
func F_replicationAttachToNewPrimary(m *base.Module)
//go:linkname F_tryReadBulkPayloadMetadata github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_tryReadBulkPayloadMetadata
func F_tryReadBulkPayloadMetadata(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_dualChannelSyncHandleRdbLoadCompletion github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dualChannelSyncHandleRdbLoadCompletion
func F_dualChannelSyncHandleRdbLoadCompletion(m *base.Module)
//go:linkname F_dualChannelSyncSuccess github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dualChannelSyncSuccess
func F_dualChannelSyncSuccess(m *base.Module)
//go:linkname F_sendCommandArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendCommandArgv
func F_sendCommandArgv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dualChannelReplMainConnRecvPsyncReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dualChannelReplMainConnRecvPsyncReply
func F_dualChannelReplMainConnRecvPsyncReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replicationCachePrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationCachePrimary
func F_replicationCachePrimary(m *base.Module, l0 int32)
//go:linkname F_replicationGetReplicaOffset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_replicationGetReplicaOffset
func F_replicationGetReplicaOffset(m *base.Module) int64
//go:linkname F_updateFailoverStatus github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_updateFailoverStatus
func F_updateFailoverStatus(m *base.Module)
//go:linkname F_shouldStartChildReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_shouldStartChildReplication
func F_shouldStartChildReplication(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parseReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseReply
func F_parseReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rioWriteBulkCount github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkCount
func F_rioWriteBulkCount(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rioWriteBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkLongLong
func F_rioWriteBulkLongLong(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_rioConnsetWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioConnsetWrite
func F_rioConnsetWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scriptIsTimedout github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptIsTimedout
func F_scriptIsTimedout(m *base.Module) int32
//go:linkname F_scriptGetCaller github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptGetCaller
func F_scriptGetCaller(m *base.Module) int32
//go:linkname F_scriptIsEval github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptIsEval
func F_scriptIsEval(m *base.Module) int32
//go:linkname F_scriptingEngineManagerGetTotalMemoryOverhead github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineManagerGetTotalMemoryOverhead
func F_scriptingEngineManagerGetTotalMemoryOverhead(m *base.Module) int32
//go:linkname F_scriptingEngineManagerGetMemoryUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineManagerGetMemoryUsage
func F_scriptingEngineManagerGetMemoryUsage(m *base.Module) int32
//go:linkname F_scriptingEngineManagerRegister github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineManagerRegister
func F_scriptingEngineManagerRegister(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_scriptingEngineManagerFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineManagerFind
func F_scriptingEngineManagerFind(m *base.Module, l0 int32) int32
//go:linkname F_scriptingEngineManagerForEachEngine github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineManagerForEachEngine
func F_scriptingEngineManagerForEachEngine(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scriptingEngineCallCompileCode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineCallCompileCode
func F_scriptingEngineCallCompileCode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_scriptingEngineCallFreeFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineCallFreeFunction
func F_scriptingEngineCallFreeFunction(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scriptingEngineCallDebuggerEnable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineCallDebuggerEnable
func F_scriptingEngineCallDebuggerEnable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_scriptingEngineCallDebuggerStart github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineCallDebuggerStart
func F_scriptingEngineCallDebuggerStart(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scriptingEngineDebuggerFlushLogs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineDebuggerFlushLogs
func F_scriptingEngineDebuggerFlushLogs(m *base.Module)
//go:linkname F_sdsHdrSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsHdrSize
func F_sdsHdrSize(m *base.Module, l0 int32) int32
//go:linkname F__sdsnewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__sdsnewlen
func F__sdsnewlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdswrite github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdswrite
func F_sdswrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_sdsnewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsnewlen
func F_sdsnewlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsempty github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsempty
func F_sdsempty(m *base.Module) int32
//go:linkname F_sdsnew github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsnew
func F_sdsnew(m *base.Module, l0 int32) int32
//go:linkname F_sdsdup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsdup
func F_sdsdup(m *base.Module, l0 int32) int32
//go:linkname F_sdsAllocPtr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsAllocPtr
func F_sdsAllocPtr(m *base.Module, l0 int32) int32
//go:linkname F_sdsAllocSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsAllocSize
func F_sdsAllocSize(m *base.Module, l0 int32) int32
//go:linkname F__sdsMakeRoomFor github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__sdsMakeRoomFor
func F__sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsMakeRoomFor github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsMakeRoomFor
func F_sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsRemoveFreeSpace github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsRemoveFreeSpace
func F_sdsRemoveFreeSpace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsIncrLen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsIncrLen
func F_sdsIncrLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sdscatlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatlen
func F_sdscatlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscat
func F_sdscat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdscatsds github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatsds
func F_sdscatsds(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdscpylen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscpylen
func F_sdscpylen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscatvprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscatvprintf
func F_sdscatvprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscatprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscatprintf
func F_sdscatprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscatfmt github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatfmt
func F_sdscatfmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscmp
func F_sdscmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdssplitargs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdssplitargs
func F_sdssplitargs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsnsplitargs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsnsplitargs
func F_sdsnsplitargs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_createSentinelAddr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSentinelAddr
func F_createSentinelAddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelScheduleScriptExecution github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelScheduleScriptExecution
func F_sentinelScheduleScriptExecution(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelRunPendingScripts github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelRunPendingScripts
func F_sentinelRunPendingScripts(m *base.Module)
//go:linkname F_sentinelCollectTerminatedScripts github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelCollectTerminatedScripts
func F_sentinelCollectTerminatedScripts(m *base.Module)
//go:linkname F_sentinelKillTimedoutScripts github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelKillTimedoutScripts
func F_sentinelKillTimedoutScripts(m *base.Module)
//go:linkname F_sentinelPendingScriptsCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelPendingScriptsCommand
func F_sentinelPendingScriptsCommand(m *base.Module, l0 int32)
//go:linkname F_instanceLinkCloseConnection github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_instanceLinkCloseConnection
func F_instanceLinkCloseConnection(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getSentinelValkeyInstanceByAddrAndRunID github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getSentinelValkeyInstanceByAddrAndRunID
func F_getSentinelValkeyInstanceByAddrAndRunID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_createSentinelValkeyInstance github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSentinelValkeyInstance
func F_createSentinelValkeyInstance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sentinelGetPrimaryByName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelGetPrimaryByName
func F_sentinelGetPrimaryByName(m *base.Module, l0 int32) int32
//go:linkname F_sentinelInstanceMapCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelInstanceMapCommand
func F_sentinelInstanceMapCommand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sentinelHandleConfiguration github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelHandleConfiguration
func F_sentinelHandleConfiguration(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteConfigSentinelOption github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteConfigSentinelOption
func F_rewriteConfigSentinelOption(m *base.Module, l0 int32)
//go:linkname F_sentinelSendAuthIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelSendAuthIfNeeded
func F_sentinelSendAuthIfNeeded(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelSendPing github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelSendPing
func F_sentinelSendPing(m *base.Module, l0 int32) int32
//go:linkname F_sentinelRefreshInstanceInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelRefreshInstanceInfo
func F_sentinelRefreshInstanceInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelConfigSetCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelConfigSetCommand
func F_sentinelConfigSetCommand(m *base.Module, l0 int32)
//go:linkname F_sentinelValidateArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelValidateArgs
func F_sentinelValidateArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelConfigGetCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelConfigGetCommand
func F_sentinelConfigGetCommand(m *base.Module, l0 int32)
//go:linkname F_addReplySentinelValkeyInstance github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplySentinelValkeyInstance
func F_addReplySentinelValkeyInstance(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelSetDebugConfigParameters github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelSetDebugConfigParameters
func F_sentinelSetDebugConfigParameters(m *base.Module, l0 int32)
//go:linkname F_addReplySentinelDebugInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplySentinelDebugInfo
func F_addReplySentinelDebugInfo(m *base.Module, l0 int32)
//go:linkname F_addReplyDictOfValkeyInstances github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyDictOfValkeyInstances
func F_addReplyDictOfValkeyInstances(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelAskPrimaryStateToOtherSentinels github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelAskPrimaryStateToOtherSentinels
func F_sentinelAskPrimaryStateToOtherSentinels(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelSetCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelSetCommand
func F_sentinelSetCommand(m *base.Module, l0 int32)
//go:linkname F_sentinelFailoverWaitStart github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelFailoverWaitStart
func F_sentinelFailoverWaitStart(m *base.Module, l0 int32)
//go:linkname F_sentinelAbortFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelAbortFailover
func F_sentinelAbortFailover(m *base.Module, l0 int32)
//go:linkname F_sentinelFailoverSelectReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelFailoverSelectReplica
func F_sentinelFailoverSelectReplica(m *base.Module, l0 int32)
//go:linkname F_sentinelFailoverSendFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelFailoverSendFailover
func F_sentinelFailoverSendFailover(m *base.Module, l0 int32)
//go:linkname F_sentinelFailoverReconfNextReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelFailoverReconfNextReplica
func F_sentinelFailoverReconfNextReplica(m *base.Module, l0 int32)
//go:linkname F_sentinelHandleDictOfValkeyInstances github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelHandleDictOfValkeyInstances
func F_sentinelHandleDictOfValkeyInstances(m *base.Module, l0 int32)
//go:linkname F_serverLogRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_serverLogRaw
func F_serverLogRaw(m *base.Module, l0 int32, l1 int32)
//go:linkname F__serverLog github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__serverLog
func F__serverLog(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_commandTimeSnapshot github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commandTimeSnapshot
func F_commandTimeSnapshot(m *base.Module) int64
//go:linkname F_resetChildState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_resetChildState
func F_resetChildState(m *base.Module)
//go:linkname F_removeClientFromMemUsageBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_removeClientFromMemUsageBucket
func F_removeClientFromMemUsageBucket(m *base.Module, l0 int32, l1 int32)
//go:linkname F_updateClientMemUsageAndBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateClientMemUsageAndBucket
func F_updateClientMemUsageAndBucket(m *base.Module, l0 int32) int32
//go:linkname F_cronUpdateMemoryStats github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_cronUpdateMemoryStats
func F_cronUpdateMemoryStats(m *base.Module)
//go:linkname F_prepareForShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_prepareForShutdown
func F_prepareForShutdown(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_closeListeningSockets github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_closeListeningSockets
func F_closeListeningSockets(m *base.Module, l0 int32)
//go:linkname F_createSharedObjectsForCompat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createSharedObjectsForCompat
func F_createSharedObjectsForCompat(m *base.Module, l0 int32)
//go:linkname F_freeServerClientMemUsageBuckets github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeServerClientMemUsageBuckets
func F_freeServerClientMemUsageBuckets(m *base.Module)
//go:linkname F_populateCommandTable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_populateCommandTable
func F_populateCommandTable(m *base.Module)
//go:linkname F_listenToPort github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listenToPort
func F_listenToPort(m *base.Module, l0 int32) int32
//go:linkname F_dbHasNoKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbHasNoKeys
func F_dbHasNoKeys(m *base.Module, l0 int32) int32
//go:linkname F_createDatabase github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createDatabase
func F_createDatabase(m *base.Module, l0 int32) int32
//go:linkname F_createDatabaseIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createDatabaseIfNeeded
func F_createDatabaseIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_lookupCommandBySdsLogic github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupCommandBySdsLogic
func F_lookupCommandBySdsLogic(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupCommandBySds github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupCommandBySds
func F_lookupCommandBySds(m *base.Module, l0 int32) int32
//go:linkname F_lookupCommandOrOriginal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupCommandOrOriginal
func F_lookupCommandOrOriginal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_postExecutionUnitOperations github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_postExecutionUnitOperations
func F_postExecutionUnitOperations(m *base.Module)
//go:linkname F_call github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_call
func F_call(m *base.Module, l0 int32, l1 int32)
//go:linkname F_prepareCommandGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_prepareCommandGeneric
func F_prepareCommandGeneric(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_processCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_processCommand
func F_processCommand(m *base.Module, l0 int32) int32
//go:linkname F_abortShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_abortShutdown
func F_abortShutdown(m *base.Module) int32
//go:linkname F_getSafeInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getSafeInfoString
func F_getSafeInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_genValkeyInfoStringCommandStats github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_genValkeyInfoStringCommandStats
func F_genValkeyInfoStringCommandStats(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_genValkeyInfoStringLatencyStats github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genValkeyInfoStringLatencyStats
func F_genValkeyInfoStringLatencyStats(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addInfoSectionsToDict github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addInfoSectionsToDict
func F_addInfoSectionsToDict(m *base.Module, l0 int32, l1 int32)
//go:linkname F_releaseInfoSectionDict github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_releaseInfoSectionDict
func F_releaseInfoSectionDict(m *base.Module, l0 int32)
//go:linkname F_genInfoSectionDict github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genInfoSectionDict
func F_genInfoSectionDict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getVersion github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getVersion
func F_getVersion(m *base.Module) int32
//go:linkname F_listenerByType github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listenerByType
func F_listenerByType(m *base.Module, l0 int32) int32
//go:linkname F_changeListener github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_changeListener
func F_changeListener(m *base.Module, l0 int32) int32
//go:linkname F_sendChildInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sendChildInfo
func F_sendChildInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_parseExtendedCommandArgumentsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseExtendedCommandArgumentsOrReply
func F_parseExtendedCommandArgumentsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_SHA1Init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_SHA1Init
func F_SHA1Init(m *base.Module, l0 int32)
//go:linkname F_SHA1Update github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_SHA1Update
func F_SHA1Update(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sha256_init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sha256_init
func F_sha256_init(m *base.Module, l0 int32)
//go:linkname F_sha256_update github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sha256_update
func F_sha256_update(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sha256_final github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sha256_final
func F_sha256_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_siphash github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_siphash
func F_siphash(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_siphash_nocase github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_siphash_nocase
func F_siphash_nocase(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_connBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_connBlock
func F_connBlock(m *base.Module, l0 int32) int32
//go:linkname F_connNonBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connNonBlock
func F_connNonBlock(m *base.Module, l0 int32) int32
//go:linkname F_connKeepAlive github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connKeepAlive
func F_connKeepAlive(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_connSendTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connSendTimeout
func F_connSendTimeout(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_RedisRegisterConnectionTypeSocket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_RedisRegisterConnectionTypeSocket
func F_RedisRegisterConnectionTypeSocket(m *base.Module) int32
//go:linkname F_sortCommandGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sortCommandGeneric
func F_sortCommandGeneric(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sparklineRenderRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sparklineRenderRange
func F_sparklineRenderRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_syncRead github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_syncRead
func F_syncRead(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_hashTypeHasVolatileFields github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeHasVolatileFields
func F_hashTypeHasVolatileFields(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeGetOrcreateVolatileSet github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeGetOrcreateVolatileSet
func F_hashTypeGetOrcreateVolatileSet(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeConvert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeConvert
func F_hashTypeConvert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeGetFromListpack github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeGetFromListpack
func F_hashTypeGetFromListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hashTypeSet github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeSet
func F_hashTypeSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32
//go:linkname F_hashTypeDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeDelete
func F_hashTypeDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashTypeCurrentFromHashTable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeCurrentFromHashTable
func F_hashTypeCurrentFromHashTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashTypeDup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeDup
func F_hashTypeDup(m *base.Module, l0 int32) int32
//go:linkname F_addHashFieldToReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addHashFieldToReply
func F_addHashFieldToReply(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashTypePersist github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypePersist
func F_hashTypePersist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_genericHgetallCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_genericHgetallCommand
func F_genericHgetallCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listTypeTryConversion github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listTypeTryConversion
func F_listTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_listTypeTryConversionRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeTryConversionRaw
func F_listTypeTryConversionRaw(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_listTypeTryConversionAppend github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeTryConversionAppend
func F_listTypeTryConversionAppend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_listTypePush github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypePush
func F_listTypePush(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_listTypeInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeInitIterator
func F_listTypeInitIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_listTypeNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeNext
func F_listTypeNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listTypeEqual github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeEqual
func F_listTypeEqual(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addListListpackRangeReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addListListpackRangeReply
func F_addListListpackRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_addListQuicklistRangeReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addListQuicklistRangeReply
func F_addListQuicklistRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_popGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_popGenericCommand
func F_popGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lmoveHandlePush github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lmoveHandlePush
func F_lmoveHandlePush(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_lmpopGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lmpopGenericCommand
func F_lmpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setTypeSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeSize
func F_setTypeSize(m *base.Module, l0 int32) int32
//go:linkname F_setTypeInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeInitIterator
func F_setTypeInitIterator(m *base.Module, l0 int32) int32
//go:linkname F_setTypeNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeNext
func F_setTypeNext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setTypeAddAux github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeAddAux
func F_setTypeAddAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_setTypeRemoveAux github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeRemoveAux
func F_setTypeRemoveAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_setTypeReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeReleaseIterator
func F_setTypeReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_setTypeRandomElement github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeRandomElement
func F_setTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setTypeDup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeDup
func F_setTypeDup(m *base.Module, l0 int32) int32
//go:linkname F_spopWithCountCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_spopWithCountCommand
func F_spopWithCountCommand(m *base.Module, l0 int32)
//go:linkname F_sunionDiffGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sunionDiffGenericCommand
func F_sunionDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_sinterGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sinterGenericCommand
func F_sinterGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_freeStream github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeStream
func F_freeStream(m *base.Module, l0 int32)
//go:linkname F_streamCreateCG github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamCreateCG
func F_streamCreateCG(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32
//go:linkname F_lpGetEdgeStreamID github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpGetEdgeStreamID
func F_lpGetEdgeStreamID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_streamIteratorGetID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamIteratorGetID
func F_streamIteratorGetID(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_streamIteratorStart github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamIteratorStart
func F_streamIteratorStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_streamAppendItem github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamAppendItem
func F_streamAppendItem(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_streamTrimByID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamTrimByID
func F_streamTrimByID(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_streamReplyWithRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamReplyWithRange
func F_streamReplyWithRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_streamGenericParseIDOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamGenericParseIDOrReply
func F_streamGenericParseIDOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32
//go:linkname F_streamRewriteTrimArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamRewriteTrimArgument
func F_streamRewriteTrimArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_streamDelConsumer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamDelConsumer
func F_streamDelConsumer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_setGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setGenericCommand
func F_setGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_msetGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_msetGenericCommand
func F_msetGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_incrDecrCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_incrDecrCommand
func F_incrDecrCommand(m *base.Module, l0 int32, l1 int64)
//go:linkname F_zslInsertNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslInsertNode
func F_zslInsertNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetFreeLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetFreeLexRange
func F_zsetFreeLexRange(m *base.Module, l0 int32)
//go:linkname F_zsetParseLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetParseLexRange
func F_zsetParseLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zslNthInLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslNthInLexRange
func F_zslNthInLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zzlNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlNext
func F_zzlNext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zzlPrev github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlPrev
func F_zzlPrev(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zzlFirstInRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zzlFirstInRange
func F_zzlFirstInRange(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zzlLexValueLteMax github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlLexValueLteMax
func F_zzlLexValueLteMax(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zzlIsInLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlIsInLexRange
func F_zzlIsInLexRange(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zzlLastInLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlLastInLexRange
func F_zzlLastInLexRange(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetLength github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zsetLength
func F_zsetLength(m *base.Module, l0 int32) int32
//go:linkname F_zsetConvertAndExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zsetConvertAndExpand
func F_zsetConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zsetConvertToListpackIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetConvertToListpackIfNeeded
func F_zsetConvertToListpackIfNeeded(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zsetScore github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetScore
func F_zsetScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zzlInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlInsert
func F_zzlInsert(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_zslUpdateScore github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zslUpdateScore
func F_zslUpdateScore(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_zsetDup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetDup
func F_zsetDup(m *base.Module, l0 int32) int32
//go:linkname F_zaddCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zaddCommand
func F_zaddCommand(m *base.Module, l0 int32)
//go:linkname F_zaddGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zaddGenericCommand
func F_zaddGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zremrangeGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zremrangeGenericCommand
func F_zremrangeGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zuiLength github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zuiLength
func F_zuiLength(m *base.Module, l0 int32) int32
//go:linkname F_zuiFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zuiFind
func F_zuiFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zdiff github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zdiff
func F_zdiff(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_zrangeGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zrangeGenericCommand
func F_zrangeGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_zrankGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zrankGenericCommand
func F_zrankGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_genericZpopCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericZpopCommand
func F_genericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_zmpopGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zmpopGenericCommand
func F_zmpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_blockingGenericZpopCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockingGenericZpopCommand
func F_blockingGenericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_zrandmemberReplyWithListpack github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zrandmemberReplyWithListpack
func F_zrandmemberReplyWithListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_zsetTypeRandomElement github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetTypeRandomElement
func F_zsetTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_getTimeoutFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getTimeoutFromObjectOrReply
func F_getTimeoutFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_tlsResetCertInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_tlsResetCertInfo
func F_tlsResetCertInfo(m *base.Module)
//go:linkname F_sendTrackingMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sendTrackingMessage
func F_sendTrackingMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_trackingInvalidateKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_trackingInvalidateKey
func F_trackingInvalidateKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_trackingInvalidateKeysOnFlush github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_trackingInvalidateKeysOnFlush
func F_trackingInvalidateKeysOnFlush(m *base.Module, l0 int32)
//go:linkname F_RedisRegisterConnectionTypeUnix github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_RedisRegisterConnectionTypeUnix
func F_RedisRegisterConnectionTypeUnix(m *base.Module) int32
//go:linkname F_stringmatchlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_stringmatchlen
func F_stringmatchlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_string2ll github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_string2ll
func F_string2ll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_string2ld github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_string2ld
func F_string2ld(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fixedpoint_d2string github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fixedpoint_d2string
func F_fixedpoint_d2string(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32
//go:linkname F_getRandomHexChars github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getRandomHexChars
func F_getRandomHexChars(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makePath github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_makePath
func F_makePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ustime github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ustime
func F_ustime(m *base.Module) int64
//go:linkname F_wangHash64 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_wangHash64
func F_wangHash64(m *base.Module, l0 int64) int64
//go:linkname F_rdbCheckError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbCheckError
func F_rdbCheckError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_redis_check_rdb_main github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_redis_check_rdb_main
func F_redis_check_rdb_main(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ffc_from_chars_double_options github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ffc_from_chars_double_options
func F_ffc_from_chars_double_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_valkey_strtod_n github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_strtod_n
func F_valkey_strtod_n(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_vectorPush github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vectorPush
func F_vectorPush(m *base.Module, l0 int32) int32
//go:linkname F_insertToBucket_RAX github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_insertToBucket_RAX
func F_insertToBucket_RAX(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_insertToBucket_VECTOR github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_insertToBucket_VECTOR
func F_insertToBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsetUpdateEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vsetUpdateEntry
func F_vsetUpdateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) int32
//go:linkname F_vsetMemUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vsetMemUsage
func F_vsetMemUsage(m *base.Module, l0 int32) int32
//go:linkname F_vsetRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vsetRelease
func F_vsetRelease(m *base.Module, l0 int32)
//go:linkname F_vsetIsEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vsetIsEmpty
func F_vsetIsEmpty(m *base.Module, l0 int32) int32
//go:linkname F_ziplistGet github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ziplistGet
func F_ziplistGet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_valkey_malloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_malloc
func F_valkey_malloc(m *base.Module, l0 int32) int32
//go:linkname F_zmalloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zmalloc_usable
func F_zmalloc_usable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkey_calloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_calloc
func F_valkey_calloc(m *base.Module, l0 int32) int32
//go:linkname F_ztryrealloc_usable_internal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ztryrealloc_usable_internal
func F_ztryrealloc_usable_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zrealloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zrealloc_usable
func F_zrealloc_usable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zmalloc_usable_size github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zmalloc_usable_size
func F_zmalloc_usable_size(m *base.Module, l0 int32) int32
//go:linkname F_valkey_free github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_free
func F_valkey_free(m *base.Module, l0 int32)
//go:linkname F_zfree_with_size github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zfree_with_size
func F_zfree_with_size(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zstrdup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zstrdup
func F_zstrdup(m *base.Module, l0 int32) int32
//go:linkname F_zmalloc_used_memory github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zmalloc_used_memory
func F_zmalloc_used_memory(m *base.Module) int32
//go:linkname F_spmcEnqueue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_spmcEnqueue
func F_spmcEnqueue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkeyAsyncConnectBind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncConnectBind
func F_valkeyAsyncConnectBind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_valkeyAsyncSetConnectCallback github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncSetConnectCallback
func F_valkeyAsyncSetConnectCallback(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkeyAsyncFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncFree
func F_valkeyAsyncFree(m *base.Module, l0 int32)
//go:linkname F_valkeyAsyncFreeInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncFreeInternal
func F_valkeyAsyncFreeInternal(m *base.Module, l0 int32)
//go:linkname F_valkeyAsyncHandleConnectFailure github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyAsyncHandleConnectFailure
func F_valkeyAsyncHandleConnectFailure(m *base.Module, l0 int32)
//go:linkname F_valkeyAsyncCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncCommand
func F_valkeyAsyncCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_valkeyContextRegisterFuncs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyContextRegisterFuncs
func F_valkeyContextRegisterFuncs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkeyContextSetFuncs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyContextSetFuncs
func F_valkeyContextSetFuncs(m *base.Module, l0 int32)
//go:linkname F_valkeySetTcpNoDelay github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeySetTcpNoDelay
func F_valkeySetTcpNoDelay(m *base.Module, l0 int32) int32
//go:linkname F_valkeyCheckSocketError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyCheckSocketError
func F_valkeyCheckSocketError(m *base.Module, l0 int32) int32
//go:linkname F_freeReplyObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeReplyObject
func F_freeReplyObject(m *base.Module, l0 int32)
//go:linkname F_valkeyFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyFree
func F_valkeyFree(m *base.Module, l0 int32)
//go:linkname F_valkeyBufferWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyBufferWrite
func F_valkeyBufferWrite(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hdr_init github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hdr_init
func F_hdr_init(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_fpconv_dtoa github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fpconv_dtoa
func F_fpconv_dtoa(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F_ldbLog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ldbLog
func F_ldbLog(m *base.Module, l0 int32)
//go:linkname F_ldbLogCString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ldbLogCString
func F_ldbLogCString(m *base.Module, l0 int32)
//go:linkname F_isLuaInsecureAPIEnabled github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_isLuaInsecureAPIEnabled
func F_isLuaInsecureAPIEnabled(m *base.Module, l0 int32) int32
//go:linkname F_initializeLuaState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initializeLuaState
func F_initializeLuaState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaFunctionLibraryCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaFunctionLibraryCreate
func F_luaFunctionLibraryCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_list_add github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_list_add
func F_list_add(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaGetFromRegistry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaGetFromRegistry
func F_luaGetFromRegistry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaPushError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaPushError
func F_luaPushError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaPushErrorBuff github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaPushErrorBuff
func F_luaPushErrorBuff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaError
func F_luaError(m *base.Module, l0 int32) int32
//go:linkname F_luaCallFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaCallFunction
func F_luaCallFunction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)
//go:linkname F_luaMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaMemory
func F_luaMemory(m *base.Module, l0 int32) int32
//go:linkname F_lua_newthread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_newthread
func F_lua_newthread(m *base.Module, l0 int32) int32
//go:linkname F_lua_lessthan github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_lessthan
func F_lua_lessthan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_tonumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_tonumber
func F_lua_tonumber(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_lua_tolstring github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_tolstring
func F_lua_tolstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_pushlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pushlstring
func F_lua_pushlstring(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_pushstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pushstring
func F_lua_pushstring(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_gettable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_gettable
func F_lua_gettable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_getfield github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_getfield
func F_lua_getfield(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_settable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_settable
func F_lua_settable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_rawset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_rawset
func F_lua_rawset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_setmetatable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_setmetatable
func F_lua_setmetatable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_call github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_call
func F_lua_call(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_pcall github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pcall
func F_lua_pcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lua_gc github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_gc
func F_lua_gc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_error github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_error
func F_lua_error(m *base.Module, l0 int32) int32
//go:linkname F_lua_next github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_next
func F_lua_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getobjname github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getobjname
func F_getobjname(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaG_typeerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaG_typeerror
func F_luaG_typeerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaG_runerror github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaG_runerror
func F_luaG_runerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaD_rawrunprotected github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_rawrunprotected
func F_luaD_rawrunprotected(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaD_callhook github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_callhook
func F_luaD_callhook(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaD_precall github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_precall
func F_luaD_precall(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaD_call github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_call
func F_luaD_call(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_resume github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_resume
func F_lua_resume(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_yield github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_yield
func F_lua_yield(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaD_protectedparser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_protectedparser
func F_luaD_protectedparser(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_DumpFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_DumpFunction
func F_DumpFunction(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaF_newupval github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaF_newupval
func F_luaF_newupval(m *base.Module, l0 int32) int32
//go:linkname F_luaF_close github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaF_close
func F_luaF_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaF_getlocalname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaF_getlocalname
func F_luaF_getlocalname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaC_step github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaC_step
func F_luaC_step(m *base.Module, l0 int32)
//go:linkname F_luaM_growaux_ github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaM_growaux_
func F_luaM_growaux_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_luaM_realloc_ github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaM_realloc_
func F_luaM_realloc_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaM_toobig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaM_toobig
func F_luaM_toobig(m *base.Module, l0 int32) int32
//go:linkname F_luaO_pushvfstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaO_pushvfstring
func F_luaO_pushvfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaO_pushfstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaO_pushfstring
func F_luaO_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaX_token2str github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_token2str
func F_luaX_token2str(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaX_lexerror github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_lexerror
func F_luaX_lexerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaX_syntaxerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaX_syntaxerror
func F_luaX_syntaxerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaX_newstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_newstring
func F_luaX_newstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_llex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_llex
func F_llex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaK_jump github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_jump
func F_luaK_jump(m *base.Module, l0 int32) int32
//go:linkname F_patchlistaux github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_patchlistaux
func F_patchlistaux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_luaK_reserveregs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaK_reserveregs
func F_luaK_reserveregs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_dischargevars github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_dischargevars
func F_luaK_dischargevars(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exp2reg github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_exp2reg
func F_exp2reg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_discharge2reg github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_discharge2reg
func F_discharge2reg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_exp2RK github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaK_exp2RK
func F_luaK_exp2RK(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_condjump github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_condjump
func F_condjump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_luaY_parser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaY_parser
func F_luaY_parser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_open_func github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_open_func
func F_open_func(m *base.Module, l0 int32, l1 int32)
//go:linkname F_chunk github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_chunk
func F_chunk(m *base.Module, l0 int32)
//go:linkname F_parlist github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_parlist
func F_parlist(m *base.Module, l0 int32)
//go:linkname F_pushclosure github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pushclosure
func F_pushclosure(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_newstate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_newstate
func F_lua_newstate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_close github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_close
func F_lua_close(m *base.Module, l0 int32)
//go:linkname F_luaS_newlstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaS_newlstr
func F_luaS_newlstr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_resizearray github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_resizearray
func F_luaH_resizearray(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setnodevector github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setnodevector
func F_setnodevector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_newkey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_newkey
func F_newkey(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_get github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_get
func F_luaH_get(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_getnum github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_getnum
func F_luaH_getnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_getstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_getstr
func F_luaH_getstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_setnum github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaH_setnum
func F_luaH_setnum(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaV_tonumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaV_tonumber
func F_luaV_tonumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaV_tostring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaV_tostring
func F_luaV_tostring(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaV_gettable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaV_gettable
func F_luaV_gettable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_luaV_settable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaV_settable
func F_luaV_settable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_luaV_lessthan github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaV_lessthan
func F_luaV_lessthan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_call_orderTM github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_call_orderTM
func F_call_orderTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaV_equalval github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaV_equalval
func F_luaV_equalval(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaV_concat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaV_concat
func F_luaV_concat(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_call_binTM github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_call_binTM
func F_call_binTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_Arith github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_Arith
func F_Arith(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_luaZ_lookahead github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaZ_lookahead
func F_luaZ_lookahead(m *base.Module, l0 int32) int32
//go:linkname F_luaL_argerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_argerror
func F_luaL_argerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_error github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_error
func F_luaL_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_typerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_typerror
func F_luaL_typerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_checkoption github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_checkoption
func F_luaL_checkoption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_checklstring github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_checklstring
func F_luaL_checklstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_checkany github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_checkany
func F_luaL_checkany(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaL_checknumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_checknumber
func F_luaL_checknumber(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_luaL_getmetafield github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_getmetafield
func F_luaL_getmetafield(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_register github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_register
func F_luaL_register(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_prepbuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_prepbuffer
func F_luaL_prepbuffer(m *base.Module, l0 int32) int32
//go:linkname F_luaL_pushresult github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_pushresult
func F_luaL_pushresult(m *base.Module, l0 int32)
//go:linkname F_luaL_ref github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_ref
func F_luaL_ref(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaL_unref github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_unref
func F_luaL_unref(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_gethooktable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_gethooktable
func F_gethooktable(m *base.Module, l0 int32)
//go:linkname F_auxupvalue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_auxupvalue
func F_auxupvalue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_match github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_match
func F_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_push_onecapture github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_push_onecapture
func F_push_onecapture(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fpconv_g_fmt github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fpconv_g_fmt
func F_fpconv_g_fmt(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F_die github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_die
func F_die(m *base.Module, l0 int32, l1 int32)
//go:linkname F_strbuf_free github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strbuf_free
func F_strbuf_free(m *base.Module, l0 int32)
//go:linkname F_strbuf_resize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strbuf_resize
func F_strbuf_resize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_json_next_token github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_json_next_token
func F_json_next_token(m *base.Module, l0 int32, l1 int32)
//go:linkname F_json_process_value github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_json_process_value
func F_json_process_value(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mp_encode_bytes github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_mp_encode_bytes
func F_mp_encode_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mp_encode_double github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_encode_double
func F_mp_encode_double(m *base.Module, l0 int32, l1 int32, l2 float64)
//go:linkname F_mp_encode_lua_bool github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_encode_lua_bool
func F_mp_encode_lua_bool(m *base.Module, l0 int32, l1 int32)
//go:linkname F_mp_encode_lua_table_as_array github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_encode_lua_table_as_array
func F_mp_encode_lua_table_as_array(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mp_encode_lua_table_as_map github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_encode_lua_table_as_map
func F_mp_encode_lua_table_as_map(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mp_decode_to_lua_array github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_decode_to_lua_array
func F_mp_decode_to_lua_array(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mp_decode_to_lua_hash github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mp_decode_to_lua_hash
func F_mp_decode_to_lua_hash(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F___memcpy github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fpclassify github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___fpclassify
func F___fpclassify(m *base.Module, l0 float64) int32
//go:linkname F_abort github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_abort
func F_abort(m *base.Module)
//go:linkname F_R_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_R_1
func F_R_1(m *base.Module, l0 float64) float64
//go:linkname F_R_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_R_2
func F_R_2(m *base.Module, l0 float64) float64
//go:linkname F_atan github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_atan
func F_atan(m *base.Module, l0 float64) float64
//go:linkname F___DOUBLE_BITS_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___DOUBLE_BITS_2
func F___DOUBLE_BITS_2(m *base.Module, l0 float64) int64
//go:linkname F___isspace_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___isspace_1
func F___isspace_1(m *base.Module, l0 int32) int32
//go:linkname F_close github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___cos github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___rem_pio2 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F___sin github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F___asctime_r github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___asctime_r
func F___asctime_r(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dirname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dirname
func F_dirname(m *base.Module, l0 int32) int32
//go:linkname F___gettimeofday github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___gettimeofday
func F___gettimeofday(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___math_xflow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___math_xflow
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64
//go:linkname F___math_uflow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___math_uflow
func F___math_uflow(m *base.Module, l0 int32) float64
//go:linkname F_top12_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_top12_1
func F_top12_1(m *base.Module, l0 float64) int32
//go:linkname F_expm1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_expm1
func F_expm1(m *base.Module, l0 float64) float64
//go:linkname F___lockfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___lockfile
func F___lockfile(m *base.Module, l0 int32) int32
//go:linkname F___unlockfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___unlockfile
func F___unlockfile(m *base.Module, l0 int32)
//go:linkname F___toread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___toread
func F___toread(m *base.Module, l0 int32) int32
//go:linkname F_fgets github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fopen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputc github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fputc
func F_fputc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fputs
func F_fputs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___dup3 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___dup3
func F___dup3(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fseeko github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___fseeko
func F___fseeko(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F___fstatat github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fsync github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F_ftruncate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ftruncate
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F___fwritex github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___fwritex
func F___fwritex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getc
func F_getc(m *base.Module, l0 int32) int32
//go:linkname F_getcwd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getcwd
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getenv github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getenv
func F_getenv(m *base.Module, l0 int32) int32
//go:linkname F___syscall_uname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___syscall_uname
func F___syscall_uname(m *base.Module, l0 int32) int32
//go:linkname F___syscall_getpid github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___syscall_getpid
func F___syscall_getpid(m *base.Module) int32
//go:linkname F___syscall_shutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___syscall_shutdown
func F___syscall_shutdown(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F___syscall_wait4 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___syscall_wait4
func F___syscall_wait4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getpid github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getpid
func F_getpid(m *base.Module) int32
//go:linkname F_match_bracket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_match_bracket
func F_match_bracket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_glob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_glob
func F_glob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___bswap_32_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___bswap_32_1
func F___bswap_32_1(m *base.Module, l0 int32) int32
//go:linkname F_inet_ntop github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_inet_ntop
func F_inet_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_isalnum github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_pthread_mutex_init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pthread_mutex_init
func F_pthread_mutex_init(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___pthread_mutex_unlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___pthread_mutex_unlock
func F___pthread_mutex_unlock(m *base.Module, l0 int32) int32
//go:linkname F_pthread_mutexattr_settype github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pthread_mutexattr_settype
func F_pthread_mutexattr_settype(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___lock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___lock
func F___lock(m *base.Module, l0 int32)
//go:linkname F___unlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___unlock
func F___unlock(m *base.Module, l0 int32)
//go:linkname F___acquire_ptc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___acquire_ptc
func F___acquire_ptc(m *base.Module)
//go:linkname F_memcmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___randname github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___randname
func F___randname(m *base.Module, l0 int32) int32
//go:linkname F___clock_nanosleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___clock_nanosleep
func F___clock_nanosleep(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___bswap_32_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___bswap_32_2
func F___bswap_32_2(m *base.Module, l0 int32) int32
//go:linkname F_ntohs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ntohs
func F_ntohs(m *base.Module, l0 int32) int32
//go:linkname F___ofl_lock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___ofl_lock
func F___ofl_lock(m *base.Module) int32
//go:linkname F___ofl_unlock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___ofl_unlock
func F___ofl_unlock(m *base.Module)
//go:linkname F_open github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_poll github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_poll
func F_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pow
func F_pow(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F_pthread_attr_getstacksize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pthread_attr_getstacksize
func F_pthread_attr_getstacksize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pthread_attr_init github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pthread_attr_init
func F_pthread_attr_init(m *base.Module, l0 int32) int32
//go:linkname F_pthread_attr_setstacksize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pthread_attr_setstacksize
func F_pthread_attr_setstacksize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dummy_4 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dummy_4
func F_dummy_4(m *base.Module, l0 int32)
//go:linkname F___get_tp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___get_tp
func F___get_tp(m *base.Module) int32
//go:linkname F_putchar github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_putchar
func F_putchar(m *base.Module, l0 int32) int32
//go:linkname F_locking_putc_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_locking_putc_2
func F_locking_putc_2(m *base.Module, l0 int32) int32
//go:linkname F_puts github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_puts
func F_puts(m *base.Module, l0 int32) int32
//go:linkname F___builtin_ctz github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___builtin_ctz
func F___builtin_ctz(m *base.Module, l0 int32) int32
//go:linkname F_qsort github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_qsort
func F_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_raise github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raise
func F_raise(m *base.Module, l0 int32) int32
//go:linkname F_rand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rand
func F_rand(m *base.Module) int32
//go:linkname F_read github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_select_ github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_select_
func F_select_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_setitimer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setitimer
func F_setitimer(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sigemptyset github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32) int32
//go:linkname F_snprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sqrt github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sqrt
func F_sqrt(m *base.Module, l0 float64) float64
//go:linkname F_sscanf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strcasecmp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strcasecmp
func F_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strchr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strcoll_l github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strcoll_l
func F___strcoll_l(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strcspn github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strcspn
func F_strcspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strerror_l github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strerror_l
func F___strerror_l(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___month_to_secs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___month_to_secs
func F___month_to_secs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_is_leap github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_is_leap
func F_is_leap(m *base.Module, l0 int32) int32
//go:linkname F_strlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F___stpncpy github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___stpncpy
func F___stpncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memrchr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___memrchr
func F___memrchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_twobyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_twobyte_strstr
func F_twobyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fourbyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fourbyte_strstr
func F_fourbyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_twoway_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_twoway_strstr
func F_twoway_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___shgetc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___shgetc
func F___shgetc(m *base.Module, l0 int32) int32
//go:linkname F_hexfloat github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hexfloat
func F_hexfloat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_decfloat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_decfloat
func F_decfloat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_strtod github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtoull github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strtoull
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_strtox_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F__vsyslog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__vsyslog
func F__vsyslog(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F___tan github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___tan
func F___tan(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_tolower github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_tolower
func F_tolower(m *base.Module, l0 int32) int32
//go:linkname F_towupper github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_towupper
func F_towupper(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_dummy_5 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dummy_5
func F_dummy_5(m *base.Module, l0 int32, l1 int32)
//go:linkname F_usleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_usleep
func F_usleep(m *base.Module, l0 int32) int32
//go:linkname F_vdprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vdprintf
func F_vdprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_printf_core github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_printf_core
func F_printf_core(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_vfprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vfprintf
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vfiprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vfiprintf
func F_vfiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___small_vsnprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___small_vsnprintf
func F___small_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___wasi_fd_is_valid github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___wasi_fd_is_valid
func F___wasi_fd_is_valid(m *base.Module, l0 int32) int32
//go:linkname F_write github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F___addtf3 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___addtf3
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___ashlti3 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___ashlti3
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___wasm_longjmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_accept github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_accept
func F_accept(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_connect github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connect
func F_connect(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getsockopt github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getsockopt
func F_getsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_socket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_socket
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
