package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	_ "unsafe"
)
//go:linkname F___syscall_setsockopt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_setsockopt
func F___syscall_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F___syscall_prlimit64 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_prlimit64
func F___syscall_prlimit64(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ACLCreateUnlinkedUser github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLCreateUnlinkedUser
func F_ACLCreateUnlinkedUser(m *base.Module) int32
//go:linkname F_ACLCreateUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLCreateUser
func F_ACLCreateUser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLResetFirstArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLResetFirstArgs
func F_ACLResetFirstArgs(m *base.Module, l0 int32)
//go:linkname F_ACLUpdateCommandRules github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLUpdateCommandRules
func F_ACLUpdateCommandRules(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ACLSetSelectorCategory github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSetSelectorCategory
func F_ACLSetSelectorCategory(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsCatPatternString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsCatPatternString
func F_sdsCatPatternString(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLSetUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSetUser
func F_ACLSetUser(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ACLFreeSelector github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLFreeSelector
func F_ACLFreeSelector(m *base.Module, l0 int32)
//go:linkname F_ACLCheckUserCredentials github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLCheckUserCredentials
func F_ACLCheckUserCredentials(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addAuthErrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addAuthErrReply
func F_addAuthErrReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addACLLogEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addACLLogEntry
func F_addACLLogEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_ACLSelectorCheckKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSelectorCheckKey
func F_ACLSelectorCheckKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_ACLSelectorCheckCmd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSelectorCheckCmd
func F_ACLSelectorCheckCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ACLCopyUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLCopyUser
func F_ACLCopyUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getUpcomingChannelList github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getUpcomingChannelList
func F_getUpcomingChannelList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLLoadFromFile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLLoadFromFile
func F_ACLLoadFromFile(m *base.Module, l0 int32) int32
//go:linkname F_ACLUpdateDefaultUserPassword github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ACLUpdateDefaultUserPassword
func F_ACLUpdateDefaultUserPassword(m *base.Module, l0 int32)
//go:linkname F_listCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listCreate
func F_listCreate(m *base.Module) int32
//go:linkname F_listEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listEmpty
func F_listEmpty(m *base.Module, l0 int32)
//go:linkname F_listRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listRelease
func F_listRelease(m *base.Module, l0 int32)
//go:linkname F_listAddNodeHead github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listAddNodeHead
func F_listAddNodeHead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listLinkNodeHead github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listLinkNodeHead
func F_listLinkNodeHead(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listDelNode
func F_listDelNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listUnlinkNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listUnlinkNode
func F_listUnlinkNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listRewind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listRewind
func F_listRewind(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listSearchKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listSearchKey
func F_listSearchKey(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aeCreateEventLoop github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeCreateEventLoop
func F_aeCreateEventLoop(m *base.Module, l0 int32) int32
//go:linkname F_aeCreateFileEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeCreateFileEvent
func F_aeCreateFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_aeWait github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_aeWait
func F_aeWait(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_anetSetError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetSetError
func F_anetSetError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_anetNonBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetNonBlock
func F_anetNonBlock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anetDisableTcpNoDelay github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetDisableTcpNoDelay
func F_anetDisableTcpNoDelay(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anetResolve github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetResolve
func F_anetResolve(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetCreateSocket github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetCreateSocket
func F_anetCreateSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F__anetTcpServer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__anetTcpServer
func F__anetTcpServer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_anetUnixServer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetUnixServer
func F_anetUnixServer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetGenericAccept github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetGenericAccept
func F_anetGenericAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_anetUnixAccept github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_anetUnixAccept
func F_anetUnixAccept(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aofManifestFree github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofManifestFree
func F_aofManifestFree(m *base.Module, l0 int32)
//go:linkname F_getAofManifestAsString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getAofManifestAsString
func F_getAofManifestAsString(m *base.Module, l0 int32) int32
//go:linkname F_aofLoadManifestFromDisk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofLoadManifestFromDisk
func F_aofLoadManifestFromDisk(m *base.Module)
//go:linkname F_writeAofManifestFile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_writeAofManifestFile
func F_writeAofManifestFile(m *base.Module, l0 int32) int32
//go:linkname F_aofDelHistoryFiles github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofDelHistoryFiles
func F_aofDelHistoryFiles(m *base.Module) int32
//go:linkname F_rewriteAppendOnlyFile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteAppendOnlyFile
func F_rewriteAppendOnlyFile(m *base.Module, l0 int32) int32
//go:linkname F_getAppendOnlyFileSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getAppendOnlyFileSize
func F_getAppendOnlyFileSize(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_openNewIncrAofForAppend github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_openNewIncrAofForAppend
func F_openNewIncrAofForAppend(m *base.Module) int32
//go:linkname F_aofRewriteLimited github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofRewriteLimited
func F_aofRewriteLimited(m *base.Module) int32
//go:linkname F_killAppendOnlyChild github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_killAppendOnlyChild
func F_killAppendOnlyChild(m *base.Module)
//go:linkname F_aofRemoveTempFile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofRemoveTempFile
func F_aofRemoveTempFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getBaseAndIncrAppendOnlyFilesSize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getBaseAndIncrAppendOnlyFilesSize
func F_getBaseAndIncrAppendOnlyFilesSize(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_bioCreateLazyFreeJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bioCreateLazyFreeJob
func F_bioCreateLazyFreeJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_allocBioJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_allocBioJob
func F_allocBioJob(m *base.Module, l0 int32) int32
//go:linkname F_bioCreateCloseJob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_bioCreateCloseJob
func F_bioCreateCloseJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_bioCreateFsyncJob github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_bioCreateFsyncJob
func F_bioCreateFsyncJob(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_getBitfieldTypeFromArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getBitfieldTypeFromArgument
func F_getBitfieldTypeFromArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lookupStringForBitCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupStringForBitCommand
func F_lookupStringForBitCommand(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_blockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockClient
func F_blockClient(m *base.Module, l0 int32, l1 int32)
//go:linkname F_unblockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unblockClient
func F_unblockClient(m *base.Module, l0 int32, l1 int32)
//go:linkname F_releaseBlockedEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_releaseBlockedEntry
func F_releaseBlockedEntry(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replyToBlockedClientTimedOut github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replyToBlockedClientTimedOut
func F_replyToBlockedClientTimedOut(m *base.Module, l0 int32)
//go:linkname F_replyToClientsBlockedOnShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replyToClientsBlockedOnShutdown
func F_replyToClientsBlockedOnShutdown(m *base.Module)
//go:linkname F_disconnectOrRedirectAllBlockedClients github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_disconnectOrRedirectAllBlockedClients
func F_disconnectOrRedirectAllBlockedClients(m *base.Module)
//go:linkname F_unblockClientOnError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unblockClientOnError
func F_unblockClientOnError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_blockForKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockForKeys
func F_blockForKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32)
//go:linkname F_signalDeletedKeyAsReady github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_signalDeletedKeyAsReady
func F_signalDeletedKeyAsReady(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_blockClientForReplicaAck github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockClientForReplicaAck
func F_blockClientForReplicaAck(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32)
//go:linkname F_callReplyCreatePromise github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_callReplyCreatePromise
func F_callReplyCreatePromise(m *base.Module, l0 int32) int32
//go:linkname F_callReplyGetString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_callReplyGetString
func F_callReplyGetString(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_callReplyGetArrayElement github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_callReplyGetArrayElement
func F_callReplyGetArrayElement(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_callReplyGetMapElement github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_callReplyGetMapElement
func F_callReplyGetMapElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_callReplyGetAttributeElement github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_callReplyGetAttributeElement
func F_callReplyGetAttributeElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_callReplyCreateError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_callReplyCreateError
func F_callReplyCreateError(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_enableParseExactReplyTypeFlag github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_enableParseExactReplyTypeFlag
func F_enableParseExactReplyTypeFlag(m *base.Module, l0 int32)
//go:linkname F_invokeReplyHandlers github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_invokeReplyHandlers
func F_invokeReplyHandlers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_sendChildInfoGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sendChildInfoGeneric
func F_sendChildInfoGeneric(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32)
//go:linkname F_connTypeOfCluster github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_connTypeOfCluster
func F_connTypeOfCluster(m *base.Module) int32
//go:linkname F_migrateCloseTimedoutSockets github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_migrateCloseTimedoutSockets
func F_migrateCloseTimedoutSockets(m *base.Module)
//go:linkname F_clearCachedClusterSlotsResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clearCachedClusterSlotsResponse
func F_clearCachedClusterSlotsResponse(m *base.Module)
//go:linkname F_clusterSlotByCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSlotByCommand
func F_clusterSlotByCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_clusterRedirectClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterRedirectClient
func F_clusterRedirectClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_addNodeReplyForClusterSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addNodeReplyForClusterSlot
func F_addNodeReplyForClusterSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_humanNodename github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_humanNodename
func F_humanNodename(m *base.Module, l0 int32) int32
//go:linkname F_getImportingSlotSource github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getImportingSlotSource
func F_getImportingSlotSource(m *base.Module, l0 int32) int32
//go:linkname F_clusterRemoveNodeFromShard github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterRemoveNodeFromShard
func F_clusterRemoveNodeFromShard(m *base.Module, l0 int32)
//go:linkname F_clusterAddNodeToShard github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterAddNodeToShard
func F_clusterAddNodeToShard(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterLookupNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterLookupNode
func F_clusterLookupNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterNodeSetSlotBit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterNodeSetSlotBit
func F_clusterNodeSetSlotBit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_setMigratingSlotDest github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setMigratingSlotDest
func F_setMigratingSlotDest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_setImportingSlotSource github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setImportingSlotSource
func F_setImportingSlotSource(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterGenNodesSlotsInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterGenNodesSlotsInfo
func F_clusterGenNodesSlotsInfo(m *base.Module, l0 int32)
//go:linkname F_clusterGenNodeDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterGenNodeDescription
func F_clusterGenNodeDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clusterDoBeforeSleep github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterDoBeforeSleep
func F_clusterDoBeforeSleep(m *base.Module, l0 int32)
//go:linkname F_clusterUpdateMyselfClientIpV6 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterUpdateMyselfClientIpV6
func F_clusterUpdateMyselfClientIpV6(m *base.Module)
//go:linkname F_clusterCloseAllSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterCloseAllSlots
func F_clusterCloseAllSlots(m *base.Module)
//go:linkname F_resetManualFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_resetManualFailover
func F_resetManualFailover(m *base.Module)
//go:linkname F_clusterDelSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterDelSlot
func F_clusterDelSlot(m *base.Module, l0 int32) int32
//go:linkname F_clusterDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterDelNode
func F_clusterDelNode(m *base.Module, l0 int32)
//go:linkname F_createClusterMsgSendBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createClusterMsgSendBlock
func F_createClusterMsgSendBlock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_freeClusterLink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClusterLink
func F_freeClusterLink(m *base.Module, l0 int32)
//go:linkname F_clusterNodeCleanupFailureReports github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterNodeCleanupFailureReports
func F_clusterNodeCleanupFailureReports(m *base.Module, l0 int32)
//go:linkname F_clusterBumpConfigEpochWithoutConsensus github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterBumpConfigEpochWithoutConsensus
func F_clusterBumpConfigEpochWithoutConsensus(m *base.Module) int32
//go:linkname F_clusterNodeIsPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeIsPrimary
func F_clusterNodeIsPrimary(m *base.Module, l0 int32) int32
//go:linkname F_clusterBlacklistAddNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterBlacklistAddNode
func F_clusterBlacklistAddNode(m *base.Module, l0 int32)
//go:linkname F_clusterBlacklistExists github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterBlacklistExists
func F_clusterBlacklistExists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterSetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSetPrimary
func F_clusterSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_delKeysInSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_delKeysInSlot
func F_delKeysInSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_clusterNodeGetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeGetPrimary
func F_clusterNodeGetPrimary(m *base.Module, l0 int32) int32
//go:linkname F_writePingExtensions github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_writePingExtensions
func F_writePingExtensions(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterSendMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSendMessage
func F_clusterSendMessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterBroadcastPong github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterBroadcastPong
func F_clusterBroadcastPong(m *base.Module, l0 int32)
//go:linkname F_clusterPropagatePublish github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterPropagatePublish
func F_clusterPropagatePublish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clusterRequestFailoverAuth github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterRequestFailoverAuth
func F_clusterRequestFailoverAuth(m *base.Module)
//go:linkname F_myselfIsBestRankedReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_myselfIsBestRankedReplica
func F_myselfIsBestRankedReplica(m *base.Module) int32
//go:linkname F_clusterGetReplicaRank github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterGetReplicaRank
func F_clusterGetReplicaRank(m *base.Module) int32
//go:linkname F_clusterGetFailedPrimaryRank github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterGetFailedPrimaryRank
func F_clusterGetFailedPrimaryRank(m *base.Module) int32
//go:linkname F_clusterFailoverReplaceYourPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterFailoverReplaceYourPrimary
func F_clusterFailoverReplaceYourPrimary(m *base.Module)
//go:linkname F_clusterUpdateState github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateState
func F_clusterUpdateState(m *base.Module)
//go:linkname F_verifyClusterConfigWithData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_verifyClusterConfigWithData
func F_verifyClusterConfigWithData(m *base.Module) int32
//go:linkname F_clusterHandleManualFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterHandleManualFailover
func F_clusterHandleManualFailover(m *base.Module)
//go:linkname F_clusterNodeIp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterNodeIp
func F_clusterNodeIp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addReplyClusterLinksDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyClusterLinksDescription
func F_addReplyClusterLinksDescription(m *base.Module, l0 int32)
//go:linkname F_clusterUpdateSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterUpdateSlots
func F_clusterUpdateSlots(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_genClusterInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_genClusterInfoString
func F_genClusterInfoString(m *base.Module, l0 int32) int32
//go:linkname F_getMyClusterNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getMyClusterNode
func F_getMyClusterNode(m *base.Module) int32
//go:linkname F_handleDebugClusterCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_handleDebugClusterCommand
func F_handleDebugClusterCommand(m *base.Module, l0 int32) int32
//go:linkname F_clusterParseSetSlotCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterParseSetSlotCommand
func F_clusterParseSetSlotCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_clusterStartHandshake github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterStartHandshake
func F_clusterStartHandshake(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clusterPromoteSelfToPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterPromoteSelfToPrimary
func F_clusterPromoteSelfToPrimary(m *base.Module)
//go:linkname F_representSlotRangeList github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_representSlotRangeList
func F_representSlotRangeList(m *base.Module, l0 int32) int32
//go:linkname F_fireModuleSlotMigrationEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fireModuleSlotMigrationEvent
func F_fireModuleSlotMigrationEvent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterRDBSaveSlotImports github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterRDBSaveSlotImports
func F_clusterRDBSaveSlotImports(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterRDBLoadSlotImport github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterRDBLoadSlotImport
func F_clusterRDBLoadSlotImport(m *base.Module, l0 int32) int32
//go:linkname F_generateSlotMigrationJobDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_generateSlotMigrationJobDescription
func F_generateSlotMigrationJobDescription(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_generateSyncSlotsEstablishCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_generateSyncSlotsEstablishCommand
func F_generateSyncSlotsEstablishCommand(m *base.Module, l0 int32) int32
//go:linkname F_propagateSyncSlotsFinish github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_propagateSyncSlotsFinish
func F_propagateSyncSlotsFinish(m *base.Module, l0 int32)
//go:linkname F_clusterUpdateSlotImportsOnOwnershipChange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateSlotImportsOnOwnershipChange
func F_clusterUpdateSlotImportsOnOwnershipChange(m *base.Module)
//go:linkname F_clusterCleanSlotImportsBeforeLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterCleanSlotImportsBeforeLoad
func F_clusterCleanSlotImportsBeforeLoad(m *base.Module)
//go:linkname F_createSlotExportJob github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createSlotExportJob
func F_createSlotExportJob(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_proceedWithSlotMigration github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_proceedWithSlotMigration
func F_proceedWithSlotMigration(m *base.Module, l0 int32)
//go:linkname F_clusterCommandCancelSlotMigrations github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandCancelSlotMigrations
func F_clusterCommandCancelSlotMigrations(m *base.Module, l0 int32)
//go:linkname F_clusterIsAnySlotExporting github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterIsAnySlotExporting
func F_clusterIsAnySlotExporting(m *base.Module) int32
//go:linkname F_clusterFeedSlotExportJobs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterFeedSlotExportJobs
func F_clusterFeedSlotExportJobs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_clusterUpdateSlotExportsOnOwnershipChange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateSlotExportsOnOwnershipChange
func F_clusterUpdateSlotExportsOnOwnershipChange(m *base.Module)
//go:linkname F_clusterFailAllSlotExportsWithMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterFailAllSlotExportsWithMessage
func F_clusterFailAllSlotExportsWithMessage(m *base.Module, l0 int32)
//go:linkname F_clusterHandleSlotMigrationErrorResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterHandleSlotMigrationErrorResponse
func F_clusterHandleSlotMigrationErrorResponse(m *base.Module, l0 int32)
//go:linkname F_isImportSlotMigrationJob github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isImportSlotMigrationJob
func F_isImportSlotMigrationJob(m *base.Module, l0 int32) int32
//go:linkname F_clusterCommandGetSlotMigrations github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterCommandGetSlotMigrations
func F_clusterCommandGetSlotMigrations(m *base.Module, l0 int32)
//go:linkname F_clusterCommandSyncSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterCommandSyncSlots
func F_clusterCommandSyncSlots(m *base.Module, l0 int32)
//go:linkname F_clusterSlotStatsAddNetworkBytesOutForSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatsAddNetworkBytesOutForSlot
func F_clusterSlotStatsAddNetworkBytesOutForSlot(m *base.Module, l0 int32, l1 int64)
//go:linkname F_clusterSlotStatsEnabled github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSlotStatsEnabled
func F_clusterSlotStatsEnabled(m *base.Module, l0 int32) int32
//go:linkname F_clusterSlotStatsIncrNetworkBytesOutForReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatsIncrNetworkBytesOutForReplication
func F_clusterSlotStatsIncrNetworkBytesOutForReplication(m *base.Module, l0 int64)
//go:linkname F_clusterSlotStatsAddNetworkBytesOutForShardedPubSubInternalPropagation github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatsAddNetworkBytesOutForShardedPubSubInternalPropagation
func F_clusterSlotStatsAddNetworkBytesOutForShardedPubSubInternalPropagation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterSlotStatReset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatReset
func F_clusterSlotStatReset(m *base.Module, l0 int32)
//go:linkname F_clusterSlotStatsAddCpuDuration github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clusterSlotStatsAddCpuDuration
func F_clusterSlotStatsAddCpuDuration(m *base.Module, l0 int32, l1 int64)
//go:linkname F_commandlogPushCurrentCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_commandlogPushCurrentCommand
func F_commandlogPushCurrentCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_performInterfaceSet github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_performInterfaceSet
func F_performInterfaceSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rewriteConfigFormatMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteConfigFormatMemory
func F_rewriteConfigFormatMemory(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_rewriteConfigStringOption github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteConfigStringOption
func F_rewriteConfigStringOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_rewriteConfigGetContentFromState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rewriteConfigGetContentFromState
func F_rewriteConfigGetContentFromState(m *base.Module, l0 int32) int32
//go:linkname F_rewriteConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteConfig
func F_rewriteConfig(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setConfigBindOption github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setConfigBindOption
func F_setConfigBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_removeConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_removeConfig
func F_removeConfig(m *base.Module, l0 int32)
//go:linkname F_addModuleNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addModuleNumericConfig
func F_addModuleNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64, l7 int64)
//go:linkname F_addModuleUnsignedNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addModuleUnsignedNumericConfig
func F_addModuleUnsignedNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64, l7 int64)
//go:linkname F_crc16 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_crc16
func F_crc16(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_crc64 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_crc64
func F_crc64(m *base.Module, l0 int64, l1 int32, l2 int64) int64
//go:linkname F_getKeySlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeySlot
func F_getKeySlot(m *base.Module, l0 int32) int32
//go:linkname F_expireIfNeededWithDictIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_expireIfNeededWithDictIndex
func F_expireIfNeededWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getKVStoreIndexForKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKVStoreIndexForKey
func F_getKVStoreIndexForKey(m *base.Module, l0 int32) int32
//go:linkname F_lookupKeyRead github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lookupKeyRead
func F_lookupKeyRead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupKeyWriteWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupKeyWriteWithFlags
func F_lookupKeyWriteWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookupKeyWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupKeyWrite
func F_lookupKeyWrite(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupKeyReadOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupKeyReadOrReply
func F_lookupKeyReadOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dbUntrackKeyWithVolatileItems github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbUntrackKeyWithVolatileItems
func F_dbUntrackKeyWithVolatileItems(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dbAddInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbAddInternal
func F_dbAddInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_dbAddRDBLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbAddRDBLoad
func F_dbAddRDBLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dbReplaceValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbReplaceValue
func F_dbReplaceValue(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setKey
func F_setKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_signalModifiedKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_signalModifiedKey
func F_signalModifiedKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dbSyncDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbSyncDelete
func F_dbSyncDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dbDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbDelete
func F_dbDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dbUnshareStringValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbUnshareStringValue
func F_dbUnshareStringValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emptyData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_emptyData
func F_emptyData(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_selectDb github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_selectDb
func F_selectDb(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_flushAllDataAndResetRDB github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_flushAllDataAndResetRDB
func F_flushAllDataAndResetRDB(m *base.Module, l0 int32)
//go:linkname F_parseScanCursorOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_parseScanCursorOrReply
func F_parseScanCursorOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scanGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scanGenericCommand
func F_scanGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_setExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setExpire
func F_setExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_dbFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbFind
func F_dbFind(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dbExpandExpires github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbExpandExpires
func F_dbExpandExpires(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_getKeysPrepareResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysPrepareResult
func F_getKeysPrepareResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getKeysFromCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getKeysFromCommand
func F_getKeysFromCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getKeysFreeResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysFreeResult
func F_getKeysFreeResult(m *base.Module, l0 int32)
//go:linkname F_xorObjectDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_xorObjectDigest
func F_xorObjectDigest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__serverPanic_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverPanic_1
func F__serverPanic_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F__serverAssert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverAssert
func F__serverAssert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_computeDatasetDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_computeDatasetDigest
func F_computeDatasetDigest(m *base.Module, l0 int32)
//go:linkname F__serverAssertWithInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__serverAssertWithInfo
func F__serverAssertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_setupSigSegvHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setupSigSegvHandler
func F_setupSigSegvHandler(m *base.Module)
//go:linkname F_debugPauseProcess github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_debugPauseProcess
func F_debugPauseProcess(m *base.Module)
//go:linkname F_dictExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictExpand
func F_dictExpand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictRehash github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictRehash
func F_dictRehash(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictAdd
func F_dictAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictBucketRehash github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictBucketRehash
func F_dictBucketRehash(m *base.Module, l0 int32, l1 int64)
//go:linkname F_dictGenericDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictGenericDelete
func F_dictGenericDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictUnlink github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictUnlink
func F_dictUnlink(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictFreeUnlinkedEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictFreeUnlinkedEntry
func F_dictFreeUnlinkedEntry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dictClear github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictClear
func F_dictClear(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dictFetchValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictFetchValue
func F_dictFetchValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictMemUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictMemUsage
func F_dictMemUsage(m *base.Module, l0 int32) int32
//go:linkname F_dictInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictInitIterator
func F_dictInitIterator(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dictResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictResetIterator
func F_dictResetIterator(m *base.Module, l0 int32)
//go:linkname F_dictGetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictGetIterator
func F_dictGetIterator(m *base.Module, l0 int32) int32
//go:linkname F_dictGetSafeIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictGetSafeIterator
func F_dictGetSafeIterator(m *base.Module, l0 int32) int32
//go:linkname F_dictNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictNext
func F_dictNext(m *base.Module, l0 int32) int32
//go:linkname F_dictGetRandomKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictGetRandomKey
func F_dictGetRandomKey(m *base.Module, l0 int32) int32
//go:linkname F_dictGetSomeKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dictGetSomeKeys
func F_dictGetSomeKeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictEmpty
func F_dictEmpty(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dictSetResizeEnabled github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictSetResizeEnabled
func F_dictSetResizeEnabled(m *base.Module, l0 int32)
//go:linkname F_entryUpdateAsStringRef github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_entryUpdateAsStringRef
func F_entryUpdateAsStringRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_entryUpdate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_entryUpdate
func F_entryUpdate(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_entryFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_entryFree
func F_entryFree(m *base.Module, l0 int32)
//go:linkname F_entryCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_entryCreate
func F_entryCreate(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_evalInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalInit
func F_evalInit(m *base.Module)
//go:linkname F_freeEvalScripts github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeEvalScripts
func F_freeEvalScripts(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_evalReset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalReset
func F_evalReset(m *base.Module, l0 int32)
//go:linkname F_evalGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_evalGenericCommand
func F_evalGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_evalRegisterNewScript github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalRegisterNewScript
func F_evalRegisterNewScript(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_evictionPoolAlloc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evictionPoolAlloc
func F_evictionPoolAlloc(m *base.Module)
//go:linkname F_startEvictionTimeProc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_startEvictionTimeProc
func F_startEvictionTimeProc(m *base.Module)
//go:linkname F_activeExpireCycleTryExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_activeExpireCycleTryExpire
func F_activeExpireCycleTryExpire(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32
//go:linkname F_activeExpireCycle github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_activeExpireCycle
func F_activeExpireCycle(m *base.Module, l0 int32) int64
//go:linkname F_expireGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_expireGenericCommand
func F_expireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_timestampIsExpired github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_timestampIsExpired
func F_timestampIsExpired(m *base.Module, l0 int64) int32
//go:linkname F_getExpirationPolicyWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getExpirationPolicyWithFlags
func F_getExpirationPolicyWithFlags(m *base.Module, l0 int32) int32
//go:linkname F_functionsLibCtxClear github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_functionsLibCtxClear
func F_functionsLibCtxClear(m *base.Module, l0 int32, l1 int32)
//go:linkname F_functionsLibCtxReleaseCurrent github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_functionsLibCtxReleaseCurrent
func F_functionsLibCtxReleaseCurrent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_functionsLibCtxFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_functionsLibCtxFree
func F_functionsLibCtxFree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_functionsInit github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_functionsInit
func F_functionsInit(m *base.Module) int32
//go:linkname F_fcallCommandGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fcallCommandGeneric
func F_fcallCommandGeneric(m *base.Module, l0 int32, l1 int32)
//go:linkname F_libraryLink github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_libraryLink
func F_libraryLink(m *base.Module, l0 int32, l1 int32)
//go:linkname F_extractUnitOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_extractUnitOrReply
func F_extractUnitOrReply(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_georadiusGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_georadiusGeneric
func F_georadiusGeneric(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_geohashDecodeToLongLatType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashDecodeToLongLatType
func F_geohashDecodeToLongLatType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_geohashDecodeToLongLatWGS84 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashDecodeToLongLatWGS84
func F_geohashDecodeToLongLatWGS84(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_geohashGetDistanceIfInRectangle github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_geohashGetDistanceIfInRectangle
func F_geohashGetDistanceIfInRectangle(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 int32) int32
//go:linkname F_geohashGetDistanceIfInPolygon github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_geohashGetDistanceIfInPolygon
func F_geohashGetDistanceIfInPolygon(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_hashtableSetResizePolicy github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableSetResizePolicy
func F_hashtableSetResizePolicy(m *base.Module, l0 int32)
//go:linkname F_hashtableMemUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableMemUsage
func F_hashtableMemUsage(m *base.Module, l0 int32) int32
//go:linkname F_hashtableResumeAutoShrink github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableResumeAutoShrink
func F_hashtableResumeAutoShrink(m *base.Module, l0 int32)
//go:linkname F_resize_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_resize_1
func F_resize_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findBucketForInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_findBucketForInsert
func F_findBucketForInsert(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_hashtableExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableExpand
func F_hashtableExpand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableFindRef github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableFindRef
func F_hashtableFindRef(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableInsertAtPosition github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableInsertAtPosition
func F_hashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashtablePop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtablePop
func F_hashtablePop(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pruneLastBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pruneLastBucket
func F_pruneLastBucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hashtableDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableDelete
func F_hashtableDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableReplaceReallocatedEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableReplaceReallocatedEntry
func F_hashtableReplaceReallocatedEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableIncrementalFindInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableIncrementalFindInit
func F_hashtableIncrementalFindInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashtableIncrementalFindStep github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableIncrementalFindStep
func F_hashtableIncrementalFindStep(m *base.Module, l0 int32) int32
//go:linkname F_hashtableIncrementalFindGetResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableIncrementalFindGetResult
func F_hashtableIncrementalFindGetResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_compactBucketChain github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_compactBucketChain
func F_compactBucketChain(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashtableReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableReleaseIterator
func F_hashtableReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_hashtableFairRandomEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableFairRandomEntry
func F_hashtableFairRandomEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableSampleEntries github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableSampleEntries
func F_hashtableSampleEntries(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableFreeStats github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableFreeStats
func F_hashtableFreeStats(m *base.Module, l0 int32)
//go:linkname F_hashtableGetStatsMsg github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashtableGetStatsMsg
func F_hashtableGetStatsMsg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_intsetNew github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_intsetNew
func F_intsetNew(m *base.Module) int32
//go:linkname F_intsetRemove github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_intsetRemove
func F_intsetRemove(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_intsetRandom github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_intsetRandom
func F_intsetRandom(m *base.Module, l0 int32) int64
//go:linkname F_flushPendingIOResponses github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_flushPendingIOResponses
func F_flushPendingIOResponses(m *base.Module, l0 int32)
//go:linkname F_trySendWriteToIOThreads github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trySendWriteToIOThreads
func F_trySendWriteToIOThreads(m *base.Module, l0 int32) int32
//go:linkname F_tryOffloadFreeArgvToIOThreads github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_tryOffloadFreeArgvToIOThreads
func F_tryOffloadFreeArgvToIOThreads(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreSize
func F_kvstoreSize(m *base.Module, l0 int32) int64
//go:linkname F_kvstoreExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreExpand
func F_kvstoreExpand(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_kvstoreHashtableExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableExpand
func F_kvstoreHashtableExpand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreGetFairRandomHashtableIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreGetFairRandomHashtableIndex
func F_kvstoreGetFairRandomHashtableIndex(m *base.Module, l0 int32) int32
//go:linkname F_kvstoreIteratorNextHashtable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorNextHashtable
func F_kvstoreIteratorNextHashtable(m *base.Module, l0 int32) int32
//go:linkname F_kvstoreIteratorRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreIteratorRelease
func F_kvstoreIteratorRelease(m *base.Module, l0 int32)
//go:linkname F_kvstoreIteratorInit github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorInit
func F_kvstoreIteratorInit(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreIteratorNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreIteratorNext
func F_kvstoreIteratorNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreTryResizeHashtables github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreTryResizeHashtables
func F_kvstoreTryResizeHashtables(m *base.Module, l0 int32, l1 int32)
//go:linkname F_kvstoreReleaseHashtableIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreReleaseHashtableIterator
func F_kvstoreReleaseHashtableIterator(m *base.Module, l0 int32)
//go:linkname F_kvstoreHashtableFairRandomEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableFairRandomEntry
func F_kvstoreHashtableFairRandomEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreHashtableFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableFind
func F_kvstoreHashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_kvstoreHashtableAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_kvstoreHashtableAdd
func F_kvstoreHashtableAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_cumulativeKeyCountAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_cumulativeKeyCountAdd
func F_cumulativeKeyCountAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_kvstoreHashtableTwoPhasePopFindRef github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableTwoPhasePopFindRef
func F_kvstoreHashtableTwoPhasePopFindRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_kvstoreHashtableTwoPhasePopDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableTwoPhasePopDelete
func F_kvstoreHashtableTwoPhasePopDelete(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_kvstoreHashtablePop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtablePop
func F_kvstoreHashtablePop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_kvstoreHashtableDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableDelete
func F_kvstoreHashtableDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_latencyAddSample github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_latencyAddSample
func F_latencyAddSample(m *base.Module, l0 int32, l1 int64)
//go:linkname F_latencyResetEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_latencyResetEvent
func F_latencyResetEvent(m *base.Module, l0 int32) int32
//go:linkname F_createLatencyReport github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createLatencyReport
func F_createLatencyReport(m *base.Module) int32
//go:linkname F_fillCommandCDF github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fillCommandCDF
func F_fillCommandCDF(m *base.Module, l0 int32, l1 int32)
//go:linkname F_latencySpecificCommandsFillCDF github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_latencySpecificCommandsFillCDF
func F_latencySpecificCommandsFillCDF(m *base.Module, l0 int32)
//go:linkname F_latencyCommandReplyWithSamples github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_latencyCommandReplyWithSamples
func F_latencyCommandReplyWithSamples(m *base.Module, l0 int32, l1 int32)
//go:linkname F_latencyCommandReplyWithLatestEvents github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_latencyCommandReplyWithLatestEvents
func F_latencyCommandReplyWithLatestEvents(m *base.Module, l0 int32)
//go:linkname F_latencyCommandGenSparkeline github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_latencyCommandGenSparkeline
func F_latencyCommandGenSparkeline(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lazyfreeResetStats github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lazyfreeResetStats
func F_lazyfreeResetStats(m *base.Module)
//go:linkname F_freeObjAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeObjAsync
func F_freeObjAsync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_freeFunctionsAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeFunctionsAsync
func F_freeFunctionsAsync(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freeReplicationBacklogRefMemAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeReplicationBacklogRefMemAsync
func F_freeReplicationBacklogRefMemAsync(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freeReplicaKeysWithExpireAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeReplicaKeysWithExpireAsync
func F_freeReplicaKeysWithExpireAsync(m *base.Module, l0 int32)
//go:linkname F_freePendingReplDataBufAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freePendingReplDataBufAsync
func F_freePendingReplDataBufAsync(m *base.Module, l0 int32)
//go:linkname F_lpShrinkToFit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpShrinkToFit
func F_lpShrinkToFit(m *base.Module, l0 int32) int32
//go:linkname F_lpNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpNext
func F_lpNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpPrev github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpPrev
func F_lpPrev(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpLength github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpLength
func F_lpLength(m *base.Module, l0 int32) int32
//go:linkname F_lpGetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpGetValue
func F_lpGetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpFind github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpFind
func F_lpFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lpInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpInsert
func F_lpInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_lpInsertInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpInsertInteger
func F_lpInsertInteger(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lpAppend github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpAppend
func F_lpAppend(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpReplace github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpReplace
func F_lpReplace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpDeleteRangeWithEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpDeleteRangeWithEntry
func F_lpDeleteRangeWithEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpDeleteRange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpDeleteRange
func F_lpDeleteRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpSeek
func F_lpSeek(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpMerge github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpMerge
func F_lpMerge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpValidateIntegrity github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpValidateIntegrity
func F_lpValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpCompare github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpCompare
func F_lpCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpRandomPair github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lpRandomPair
func F_lpRandomPair(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_lpNextRandom github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpNextRandom
func F_lpNextRandom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lwFreeCanvas github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lwFreeCanvas
func F_lwFreeCanvas(m *base.Module, l0 int32)
//go:linkname F_lru_getIdleSecs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lru_getIdleSecs
func F_lru_getIdleSecs(m *base.Module, l0 int32) int32
//go:linkname F_lfu_getFrequency github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lfu_getFrequency
func F_lfu_getFrequency(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lrulfu_updateClockAndPolicy github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lrulfu_updateClockAndPolicy
func F_lrulfu_updateClockAndPolicy(m *base.Module, l0 int64, l1 int32)
//go:linkname F_lzf_decompress github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lzf_decompress
func F_lzf_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freePrefetchCommandsBatch github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freePrefetchCommandsBatch
func F_freePrefetchCommandsBatch(m *base.Module)
//go:linkname F_prefetchCommandsBatchInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_prefetchCommandsBatchInit
func F_prefetchCommandsBatchInit(m *base.Module)
//go:linkname F_addCommandToBatchAndProcessIfFull github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addCommandToBatchAndProcessIfFull
func F_addCommandToBatchAndProcessIfFull(m *base.Module, l0 int32) int32
//go:linkname F_memtest_test github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_memtest_test
func F_memtest_test(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleReleaseTempClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleReleaseTempClient
func F_moduleReleaseTempClient(m *base.Module, l0 int32)
//go:linkname F_moduleDelKeyIfEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleDelKeyIfEmpty
func F_moduleDelKeyIfEmpty(m *base.Module, l0 int32) int32
//go:linkname F_moduleFreeKeyIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleFreeKeyIterator
func F_moduleFreeKeyIterator(m *base.Module, l0 int32)
//go:linkname F_moduleFreeContext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleFreeContext
func F_moduleFreeContext(m *base.Module, l0 int32)
//go:linkname F_VM_FreeCallReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_VM_FreeCallReply
func F_VM_FreeCallReply(m *base.Module, l0 int32)
//go:linkname F_VM_CloseKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_VM_CloseKey
func F_VM_CloseKey(m *base.Module, l0 int32)
//go:linkname F_moduleCallCommandUnblockedHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleCallCommandUnblockedHandler
func F_moduleCallCommandUnblockedHandler(m *base.Module, l0 int32)
//go:linkname F_moduleCreateContext github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleCreateContext
func F_moduleCreateContext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_moduleAllocateContext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleAllocateContext
func F_moduleAllocateContext(m *base.Module) int32
//go:linkname F_moduleScriptingEngineInitContext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleScriptingEngineInitContext
func F_moduleScriptingEngineInitContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleCloseKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleCloseKey
func F_moduleCloseKey(m *base.Module, l0 int32)
//go:linkname F_moduleReplyWithCollection github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleReplyWithCollection
func F_moduleReplyWithCollection(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_modulePopulateClientInfoStructure github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_modulePopulateClientInfoStructure
func F_modulePopulateClientInfoStructure(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zsetInitLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zsetInitLexRange
func F_zsetInitLexRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleCallCommandHelper github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleCallCommandHelper
func F_moduleCallCommandHelper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_moduleTypeLookupModuleByNameIgnoreCase github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleTypeLookupModuleByNameIgnoreCase
func F_moduleTypeLookupModuleByNameIgnoreCase(m *base.Module, l0 int32) int32
//go:linkname F_moduleAllDatatypesHandleErrors github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleAllDatatypesHandleErrors
func F_moduleAllDatatypesHandleErrors(m *base.Module) int32
//go:linkname F_moduleVerifyAllAllowAtomicSlotMigrationOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleVerifyAllAllowAtomicSlotMigrationOrReply
func F_moduleVerifyAllAllowAtomicSlotMigrationOrReply(m *base.Module, l0 int32) int32
//go:linkname F_moduleLogRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleLogRaw
func F_moduleLogRaw(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleUnblockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleUnblockClient
func F_moduleUnblockClient(m *base.Module, l0 int32)
//go:linkname F_VM_UnblockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_VM_UnblockClient
func F_VM_UnblockClient(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleBlockedClientMayTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleBlockedClientMayTimeout
func F_moduleBlockedClientMayTimeout(m *base.Module, l0 int32) int32
//go:linkname F_firePostExecutionUnitJobs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_firePostExecutionUnitJobs
func F_firePostExecutionUnitJobs(m *base.Module)
//go:linkname F_moduleNotifyKeyspaceEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleNotifyKeyspaceEvent
func F_moduleNotifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleGetClusterNodeInfoForClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleGetClusterNodeInfoForClient
func F_moduleGetClusterNodeInfoForClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_moduleUnregisterFilters github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleUnregisterFilters
func F_moduleUnregisterFilters(m *base.Module, l0 int32) int32
//go:linkname F_moduleUnsubscribeAllServerEvents github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleUnsubscribeAllServerEvents
func F_moduleUnsubscribeAllServerEvents(m *base.Module, l0 int32)
//go:linkname F_processModuleLoadingProgressEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processModuleLoadingProgressEvent
func F_processModuleLoadingProgressEvent(m *base.Module, l0 int32)
//go:linkname F_moduleGetFreeEffort github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleGetFreeEffort
func F_moduleGetFreeEffort(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_modulesCron github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_modulesCron
func F_modulesCron(m *base.Module)
//go:linkname F_moduleInitPostOnLoadResolved github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleInitPostOnLoadResolved
func F_moduleInitPostOnLoadResolved(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_addReplyLoadedModules github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyLoadedModules
func F_addReplyLoadedModules(m *base.Module, l0 int32)
//go:linkname F_genModulesInfoStringRenderModuleOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_genModulesInfoStringRenderModuleOptions
func F_genModulesInfoStringRenderModuleOptions(m *base.Module, l0 int32) int32
//go:linkname F_setModuleBoolConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setModuleBoolConfig
func F_setModuleBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setModuleStringConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setModuleStringConfig
func F_setModuleStringConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setModuleEnumConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setModuleEnumConfig
func F_setModuleEnumConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setModuleNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setModuleNumericConfig
func F_setModuleNumericConfig(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_getModuleEnumConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getModuleEnumConfig
func F_getModuleEnumConfig(m *base.Module, l0 int32) int32
//go:linkname F_loadModuleConfigs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_loadModuleConfigs
func F_loadModuleConfigs(m *base.Module, l0 int32) int32
//go:linkname F_moduleConfigApplyConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleConfigApplyConfig
func F_moduleConfigApplyConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_moduleConfigValidityCheck github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_moduleConfigValidityCheck
func F_moduleConfigValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_monotonicInit github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_monotonicInit
func F_monotonicInit(m *base.Module) int32
//go:linkname F_unwatchAllKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unwatchAllKeys
func F_unwatchAllKeys(m *base.Module, l0 int32)
//go:linkname F_resetClientMultiState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resetClientMultiState
func F_resetClientMultiState(m *base.Module, l0 int32)
//go:linkname F_queueMultiCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_queueMultiCommand
func F_queueMultiCommand(m *base.Module, l0 int32, l1 int64)
//go:linkname F_discardTransaction github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_discardTransaction
func F_discardTransaction(m *base.Module, l0 int32)
//go:linkname F_touchWatchedKey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_touchWatchedKey
func F_touchWatchedKey(m *base.Module, l0 int32, l1 int32)
//go:linkname F_createClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createClient
func F_createClient(m *base.Module, l0 int32) int32
//go:linkname F_readToQueryBuf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_readToQueryBuf
func F_readToQueryBuf(m *base.Module, l0 int32) int32
//go:linkname F_handleReadResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_handleReadResult
func F_handleReadResult(m *base.Module, l0 int32) int32
//go:linkname F_processInputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processInputBuffer
func F_processInputBuffer(m *base.Module, l0 int32) int32
//go:linkname F_beforeNextClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_beforeNextClient
func F_beforeNextClient(m *base.Module, l0 int32)
//go:linkname F_installClientWriteHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_installClientWriteHandler
func F_installClientWriteHandler(m *base.Module, l0 int32)
//go:linkname F_writeToClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_writeToClient
func F_writeToClient(m *base.Module, l0 int32) int32
//go:linkname F_freeClientAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClientAsync
func F_freeClientAsync(m *base.Module, l0 int32)
//go:linkname F_prepareClientToWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_prepareClientToWrite
func F_prepareClientToWrite(m *base.Module, l0 int32) int32
//go:linkname F_clientHasPendingReplies github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clientHasPendingReplies
func F_clientHasPendingReplies(m *base.Module, l0 int32) int32
//go:linkname F_createCachedResponseClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createCachedResponseClient
func F_createCachedResponseClient(m *base.Module, l0 int32) int32
//go:linkname F_deleteCachedResponseClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_deleteCachedResponseClient
func F_deleteCachedResponseClient(m *base.Module, l0 int32)
//go:linkname F_freeClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClient
func F_freeClient(m *base.Module, l0 int32) int32
//go:linkname F_closeClientOnOutputBufferLimitReached github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_closeClientOnOutputBufferLimitReached
func F_closeClientOnOutputBufferLimitReached(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__addReplyToBufferOrList github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__addReplyToBufferOrList
func F__addReplyToBufferOrList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_catClientInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_catClientInfoString
func F_catClientInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_addReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReply
func F_addReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplySds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplySds
func F_addReplySds(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyProto github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyProto
func F_addReplyProto(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyErrorLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorLength
func F_addReplyErrorLength(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyErrorObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorObject
func F_addReplyErrorObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyOrErrorObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyOrErrorObject
func F_addReplyOrErrorObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorSdsEx github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorSdsEx
func F_addReplyErrorSdsEx(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyErrorSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorSds
func F_addReplyErrorSds(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorFormat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyErrorFormat
func F_addReplyErrorFormat(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyErrorArity github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorArity
func F_addReplyErrorArity(m *base.Module, l0 int32)
//go:linkname F_addReplyStatusLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyStatusLength
func F_addReplyStatusLength(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyStatusFormat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyStatusFormat
func F_addReplyStatusFormat(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyDeferredLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyDeferredLen
func F_addReplyDeferredLen(m *base.Module, l0 int32) int32
//go:linkname F_setDeferredReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredReply
func F_setDeferredReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_setDeferredAggregateLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setDeferredAggregateLen
func F_setDeferredAggregateLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_setDeferredArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredArrayLen
func F_setDeferredArrayLen(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setDeferredMapLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setDeferredMapLen
func F_setDeferredMapLen(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyDouble
func F_addReplyDouble(m *base.Module, l0 int32, l1 float64)
//go:linkname F_addReplyBigNum github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyBigNum
func F_addReplyBigNum(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyBulkCBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyBulkCBuffer
func F_addReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyHumanLongDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyHumanLongDouble
func F_addReplyHumanLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64)
//go:linkname F_addReplyBulk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyBulk
func F_addReplyBulk(m *base.Module, l0 int32, l1 int32)
//go:linkname F__addReplyLongLongWithPrefix github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__addReplyLongLongWithPrefix
func F__addReplyLongLongWithPrefix(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_addReplyArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyArrayLen
func F_addReplyArrayLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addWritePreparedReplyArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addWritePreparedReplyArrayLen
func F_addWritePreparedReplyArrayLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyMapLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyMapLen
func F_addReplyMapLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplySetLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplySetLen
func F_addReplySetLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyAttributeLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyAttributeLen
func F_addReplyAttributeLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyPushLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyPushLen
func F_addReplyPushLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyNull github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyNull
func F_addReplyNull(m *base.Module, l0 int32)
//go:linkname F_addReplyNullArray github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyNullArray
func F_addReplyNullArray(m *base.Module, l0 int32)
//go:linkname F_addWritePreparedReplyBulkCBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addWritePreparedReplyBulkCBuffer
func F_addWritePreparedReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyBulkSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyBulkSds
func F_addReplyBulkSds(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyBulkLongLong
func F_addReplyBulkLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addWritePreparedReplyBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addWritePreparedReplyBulkLongLong
func F_addWritePreparedReplyBulkLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addReplyVerbatim github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyVerbatim
func F_addReplyVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_addExtendedReplyHelp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addExtendedReplyHelp
func F_addExtendedReplyHelp(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyHelp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyHelp
func F_addReplyHelp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplySubcommandSyntaxError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplySubcommandSyntaxError
func F_addReplySubcommandSyntaxError(m *base.Module, l0 int32)
//go:linkname F_initDeferredReplyBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initDeferredReplyBuffer
func F_initDeferredReplyBuffer(m *base.Module, l0 int32)
//go:linkname F_deferredAfterErrorReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_deferredAfterErrorReply
func F_deferredAfterErrorReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getClientPeerId github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getClientPeerId
func F_getClientPeerId(m *base.Module, l0 int32) int32
//go:linkname F_getClientSockname github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getClientSockname
func F_getClientSockname(m *base.Module, l0 int32) int32
//go:linkname F_copyReplicaOutputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_copyReplicaOutputBuffer
func F_copyReplicaOutputBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freeClientArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClientArgv
func F_freeClientArgv(m *base.Module, l0 int32)
//go:linkname F_freeClientOrCloseLater github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClientOrCloseLater
func F_freeClientOrCloseLater(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initSharedQueryBuf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initSharedQueryBuf
func F_initSharedQueryBuf(m *base.Module)
//go:linkname F_postWriteToReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_postWriteToReplica
func F_postWriteToReplica(m *base.Module, l0 int32)
//go:linkname F_handleParseResults github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_handleParseResults
func F_handleParseResults(m *base.Module, l0 int32) int32
//go:linkname F_protectClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_protectClient
func F_protectClient(m *base.Module, l0 int32)
//go:linkname F_unprotectClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_unprotectClient
func F_unprotectClient(m *base.Module, l0 int32)
//go:linkname F_commandProcessed github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_commandProcessed
func F_commandProcessed(m *base.Module, l0 int32)
//go:linkname F_processCommandAndResetClient github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_processCommandAndResetClient
func F_processCommandAndResetClient(m *base.Module, l0 int32) int32
//go:linkname F_catClientInfoShortString github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_catClientInfoShortString
func F_catClientInfoShortString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getAllClientsInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getAllClientsInfoString
func F_getAllClientsInfoString(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_updatePausedActions github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_updatePausedActions
func F_updatePausedActions(m *base.Module)
//go:linkname F_rewriteClientCommandVector github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteClientCommandVector
func F_rewriteClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replaceClientCommandVector github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replaceClientCommandVector
func F_replaceClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rewriteClientCommandArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteClientCommandArgument
func F_rewriteClientCommandArgument(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_unblockPostponedClients github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_unblockPostponedClients
func F_unblockPostponedClients(m *base.Module)
//go:linkname F_isPausedActionsWithUpdate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_isPausedActionsWithUpdate
func F_isPausedActionsWithUpdate(m *base.Module, l0 int32) int32
//go:linkname F_processEventsWhileBlocked github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processEventsWhileBlocked
func F_processEventsWhileBlocked(m *base.Module)
//go:linkname F_keyspaceEventsFlagsToString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_keyspaceEventsFlagsToString
func F_keyspaceEventsFlagsToString(m *base.Module, l0 int32) int32
//go:linkname F_notifyKeyspaceEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_notifyKeyspaceEvent
func F_notifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_createObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createObject
func F_createObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_createUnembeddedObjectWithKeyAndExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createUnembeddedObjectWithKeyAndExpire
func F_createUnembeddedObjectWithKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_createRawStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createRawStringObject
func F_createRawStringObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_createStringObject_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createStringObject_1
func F_createStringObject_1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_createEmbeddedStringObjectWithKeyAndExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createEmbeddedStringObjectWithKeyAndExpire
func F_createEmbeddedStringObjectWithKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_objectGetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectGetVal
func F_objectGetVal(m *base.Module, l0 int32) int32
//go:linkname F_decrRefCount github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_decrRefCount
func F_decrRefCount(m *base.Module, l0 int32)
//go:linkname F_objectSetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectSetVal
func F_objectSetVal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_objectUnembedVal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectUnembedVal
func F_objectUnembedVal(m *base.Module, l0 int32)
//go:linkname F_tryCreateStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_tryCreateStringObject
func F_tryCreateStringObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_createStringObjectFromLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createStringObjectFromLongLong
func F_createStringObjectFromLongLong(m *base.Module, l0 int64) int32
//go:linkname F_createStringObjectFromLongLongForValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createStringObjectFromLongLongForValue
func F_createStringObjectFromLongLongForValue(m *base.Module, l0 int64) int32
//go:linkname F_createStringObjectFromLongLongWithSds github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createStringObjectFromLongLongWithSds
func F_createStringObjectFromLongLongWithSds(m *base.Module, l0 int64) int32
//go:linkname F_createListListpackObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createListListpackObject
func F_createListListpackObject(m *base.Module) int32
//go:linkname F_createSetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createSetObject
func F_createSetObject(m *base.Module) int32
//go:linkname F_createIntsetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createIntsetObject
func F_createIntsetObject(m *base.Module) int32
//go:linkname F_createSetListpackObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createSetListpackObject
func F_createSetListpackObject(m *base.Module) int32
//go:linkname F_createHashObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createHashObject
func F_createHashObject(m *base.Module) int32
//go:linkname F_createZsetListpackObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createZsetListpackObject
func F_createZsetListpackObject(m *base.Module) int32
//go:linkname F_createModuleObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createModuleObject
func F_createModuleObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_incrRefCount github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_incrRefCount
func F_incrRefCount(m *base.Module, l0 int32)
//go:linkname F_checkType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_checkType
func F_checkType(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tryObjectEncoding github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_tryObjectEncoding
func F_tryObjectEncoding(m *base.Module, l0 int32) int32
//go:linkname F_compareStringObjectsWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_compareStringObjectsWithFlags
func F_compareStringObjectsWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_compareStringObjects github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_compareStringObjects
func F_compareStringObjects(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_stringObjectLen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_stringObjectLen
func F_stringObjectLen(m *base.Module, l0 int32) int32
//go:linkname F_getDoubleFromObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getDoubleFromObject
func F_getDoubleFromObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getLongDoubleFromObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getLongDoubleFromObject
func F_getLongDoubleFromObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getLongLongFromObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getLongLongFromObject
func F_getLongLongFromObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getLongLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getLongLongFromObjectOrReply
func F_getLongLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getRangeLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getRangeLongFromObjectOrReply
func F_getRangeLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_objectComputeSize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectComputeSize
func F_objectComputeSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getMemoryOverheadData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getMemoryOverheadData
func F_getMemoryOverheadData(m *base.Module) int32
//go:linkname F_getMemoryDoctorReport github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getMemoryDoctorReport
func F_getMemoryDoctorReport(m *base.Module) int32
//go:linkname F_objectSetLRUOrLFU github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectSetLRUOrLFU
func F_objectSetLRUOrLFU(m *base.Module, l0 int32, l1 int64, l2 int64) int32
//go:linkname F_addReplyPubsubMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyPubsubMessage
func F_addReplyPubsubMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_freeClientPubSubData github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeClientPubSubData
func F_freeClientPubSubData(m *base.Module, l0 int32)
//go:linkname F_pubsubUnsubscribeChannel github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubUnsubscribeChannel
func F_pubsubUnsubscribeChannel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_pubsubUnsubscribeAllPatterns github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pubsubUnsubscribeAllPatterns
func F_pubsubUnsubscribeAllPatterns(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pubsubUnsubscribeAllChannelsInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubUnsubscribeAllChannelsInternal
func F_pubsubUnsubscribeAllChannelsInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pubsubUnsubscribePattern github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubUnsubscribePattern
func F_pubsubUnsubscribePattern(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pubsubPublishMessageAndPropagateToCluster github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pubsubPublishMessageAndPropagateToCluster
func F_pubsubPublishMessageAndPropagateToCluster(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___quicklistInsertPlainNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___quicklistInsertPlainNode
func F___quicklistInsertPlainNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_quicklistCreateNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistCreateNode
func F_quicklistCreateNode(m *base.Module) int32
//go:linkname F___quicklistInsertNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___quicklistInsertNode
func F___quicklistInsertNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F___quicklistDelNode github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___quicklistDelNode
func F___quicklistDelNode(m *base.Module, l0 int32, l1 int32)
//go:linkname F_quicklistReplaceEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistReplaceEntry
func F_quicklistReplaceEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F___quicklistCompress github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___quicklistCompress
func F___quicklistCompress(m *base.Module, l0 int32, l1 int32)
//go:linkname F__quicklistMergeNodes github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__quicklistMergeNodes
func F__quicklistMergeNodes(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___quicklistDecompressNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___quicklistDecompressNode
func F___quicklistDecompressNode(m *base.Module, l0 int32)
//go:linkname F_quicklistRepr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_quicklistRepr
func F_quicklistRepr(m *base.Module, l0 int32, l1 int32)
//go:linkname F_raxNew github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxNew
func F_raxNew(m *base.Module) int32
//go:linkname F_raxGenericInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxGenericInsert
func F_raxGenericInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_raxRemove github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxRemove
func F_raxRemove(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxLowWalk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxLowWalk
func F_raxLowWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_raxInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxInsert
func F_raxInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_raxRecursiveFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxRecursiveFree
func F_raxRecursiveFree(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_raxFreeWithCallback github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxFreeWithCallback
func F_raxFreeWithCallback(m *base.Module, l0 int32, l1 int32)
//go:linkname F_raxIteratorAddChars github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxIteratorAddChars
func F_raxIteratorAddChars(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_raxSeekGreatest github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxSeekGreatest
func F_raxSeekGreatest(m *base.Module, l0 int32) int32
//go:linkname F_raxNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxNext
func F_raxNext(m *base.Module, l0 int32) int32
//go:linkname F_raxPrev github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxPrev
func F_raxPrev(m *base.Module, l0 int32) int32
//go:linkname F_raxStop github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_raxStop
func F_raxStop(m *base.Module, l0 int32)
//go:linkname F_rdbIsVersionAccepted github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbIsVersionAccepted
func F_rdbIsVersionAccepted(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbReportError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbReportError
func F_rdbReportError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_rdbLoadType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadType
func F_rdbLoadType(m *base.Module, l0 int32) int32
//go:linkname F_rdbLoadTime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadTime
func F_rdbLoadTime(m *base.Module, l0 int32) int64
//go:linkname F_rdbLoadMillisecondTime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadMillisecondTime
func F_rdbLoadMillisecondTime(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_rdbSaveLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveLen
func F_rdbSaveLen(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_rdbLoadLenByRef github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadLenByRef
func F_rdbLoadLenByRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadLen
func F_rdbLoadLen(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_rdbLoadIntegerObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadIntegerObject
func F_rdbLoadIntegerObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbSaveLzfStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveLzfStringObject
func F_rdbSaveLzfStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadLzfStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadLzfStringObject
func F_rdbLoadLzfStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadBinaryFloatValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadBinaryFloatValue
func F_rdbLoadBinaryFloatValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbGetObjectType github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbGetObjectType
func F_rdbGetObjectType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbLoadObjectType github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadObjectType
func F_rdbLoadObjectType(m *base.Module, l0 int32) int32
//go:linkname F_rdbSaveObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveObject
func F_rdbSaveObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbSaveInfoAuxFields github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSaveInfoAuxFields
func F_rdbSaveInfoAuxFields(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveSingleModuleAux github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveSingleModuleAux
func F_rdbSaveSingleModuleAux(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveFunctions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveFunctions
func F_rdbSaveFunctions(m *base.Module, l0 int32) int32
//go:linkname F_rdbSaveDb github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveDb
func F_rdbSaveDb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbSaveRioWithEOFMark github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSaveRioWithEOFMark
func F_rdbSaveRioWithEOFMark(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbSaveInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSaveInternal
func F_rdbSaveInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbSave github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSave
func F_rdbSave(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbSaveBackground github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbSaveBackground
func F_rdbSaveBackground(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbRemoveTempFile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbRemoveTempFile
func F_rdbRemoveTempFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_rdbLoadCheckModuleValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadCheckModuleValue
func F_rdbLoadCheckModuleValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbLoadObject github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadObject
func F_rdbLoadObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int32
//go:linkname F_startLoading github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_startLoading
func F_startLoading(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_stopLoading github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_stopLoading
func F_stopLoading(m *base.Module, l0 int32)
//go:linkname F_rdbFunctionLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbFunctionLoad
func F_rdbFunctionLoad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbLoadRio github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoadRio
func F_rdbLoadRio(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rdbLoad
func F_rdbLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_killRDBChild github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_killRDBChild
func F_killRDBChild(m *base.Module)
//go:linkname F_connTypeOfReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_connTypeOfReplication
func F_connTypeOfReplication(m *base.Module) int32
//go:linkname F_replicationGetReplicaName github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationGetReplicaName
func F_replicationGetReplicaName(m *base.Module, l0 int32) int32
//go:linkname F_createReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createReplicationBacklog
func F_createReplicationBacklog(m *base.Module)
//go:linkname F_incrementalTrimReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_incrementalTrimReplicationBacklog
func F_incrementalTrimReplicationBacklog(m *base.Module, l0 int32)
//go:linkname F_addRdbReplicaToPsyncWait github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addRdbReplicaToPsyncWait
func F_addRdbReplicaToPsyncWait(m *base.Module, l0 int32)
//go:linkname F_backfillRdbReplicasToPsyncWait github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_backfillRdbReplicasToPsyncWait
func F_backfillRdbReplicasToPsyncWait(m *base.Module)
//go:linkname F_prepareReplicasToWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_prepareReplicasToWrite
func F_prepareReplicasToWrite(m *base.Module) int32
//go:linkname F_freeReplicaReferencedReplBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeReplicaReferencedReplBuffer
func F_freeReplicaReferencedReplBuffer(m *base.Module, l0 int32)
//go:linkname F_replicationFeedReplicas github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationFeedReplicas
func F_replicationFeedReplicas(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_replicationFeedMonitors github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationFeedMonitors
func F_replicationFeedMonitors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_replicationSetupReplicaForFullResync github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationSetupReplicaForFullResync
func F_replicationSetupReplicaForFullResync(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_primaryTryPartialResynchronization github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_primaryTryPartialResynchronization
func F_primaryTryPartialResynchronization(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_replicationUnsetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationUnsetPrimary
func F_replicationUnsetPrimary(m *base.Module)
//go:linkname F_changeReplicationId github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_changeReplicationId
func F_changeReplicationId(m *base.Module)
//go:linkname F_replicationHandlePrimaryDisconnection github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationHandlePrimaryDisconnection
func F_replicationHandlePrimaryDisconnection(m *base.Module)
//go:linkname F_replicaStartCommandStream github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicaStartCommandStream
func F_replicaStartCommandStream(m *base.Module, l0 int32)
//go:linkname F_replicaPutOnline github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaPutOnline
func F_replicaPutOnline(m *base.Module, l0 int32) int32
//go:linkname F_replicationSendAck github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationSendAck
func F_replicationSendAck(m *base.Module)
//go:linkname F_replicationSendNewlineToPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_replicationSendNewlineToPrimary
func F_replicationSendNewlineToPrimary(m *base.Module)
//go:linkname F_cleanupTransferResources github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_cleanupTransferResources
func F_cleanupTransferResources(m *base.Module)
//go:linkname F_replicaReceiveRDBFromPrimaryToDisk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaReceiveRDBFromPrimaryToDisk
func F_replicaReceiveRDBFromPrimaryToDisk(m *base.Module, l0 int32, l1 int32)
//go:linkname F_receiveSynchronousResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_receiveSynchronousResponse
func F_receiveSynchronousResponse(m *base.Module, l0 int32) int32
//go:linkname F_sendCurrentOffsetToReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sendCurrentOffsetToReplica
func F_sendCurrentOffsetToReplica(m *base.Module, l0 int32) int32
//go:linkname F_replicationSendAuth github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationSendAuth
func F_replicationSendAuth(m *base.Module, l0 int32) int32
//go:linkname F_streamReplDataBufToDb github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamReplDataBufToDb
func F_streamReplDataBufToDb(m *base.Module, l0 int32) int32
//go:linkname F_replicationSteadyStateInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationSteadyStateInit
func F_replicationSteadyStateInit(m *base.Module)
//go:linkname F_replicaSendPsyncCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaSendPsyncCommand
func F_replicaSendPsyncCommand(m *base.Module, l0 int32) int32
//go:linkname F_replicaProcessPsyncReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaProcessPsyncReply
func F_replicaProcessPsyncReply(m *base.Module, l0 int32) int32
//go:linkname F_syncWithPrimaryHandleConnectingState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_syncWithPrimaryHandleConnectingState
func F_syncWithPrimaryHandleConnectingState(m *base.Module, l0 int32) int32
//go:linkname F_syncWithPrimaryHandleSendHandshakeState github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_syncWithPrimaryHandleSendHandshakeState
func F_syncWithPrimaryHandleSendHandshakeState(m *base.Module, l0 int32) int32
//go:linkname F_syncWithPrimaryHandleReceiveAuthReplyState github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_syncWithPrimaryHandleReceiveAuthReplyState
func F_syncWithPrimaryHandleReceiveAuthReplyState(m *base.Module, l0 int32) int32
//go:linkname F_syncWithPrimaryHandleReceiveIPReplyState github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_syncWithPrimaryHandleReceiveIPReplyState
func F_syncWithPrimaryHandleReceiveIPReplyState(m *base.Module, l0 int32) int32
//go:linkname F_syncWithPrimaryHandleReceiveNodeIDReplyState github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_syncWithPrimaryHandleReceiveNodeIDReplyState
func F_syncWithPrimaryHandleReceiveNodeIDReplyState(m *base.Module, l0 int32) int32
//go:linkname F_replicationSetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationSetPrimary
func F_replicationSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_replicationCron github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationCron
func F_replicationCron(m *base.Module)
//go:linkname F_shouldStartChildReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_shouldStartChildReplication
func F_shouldStartChildReplication(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parseReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseReply
func F_parseReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rioInitWithFd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rioInitWithFd
func F_rioInitWithFd(m *base.Module, l0 int32, l1 int32)
//go:linkname F_rioFreeFd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rioFreeFd
func F_rioFreeFd(m *base.Module, l0 int32)
//go:linkname F_rioWriteBulkString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rioWriteBulkString
func F_rioWriteBulkString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rioWriteBulkDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rioWriteBulkDouble
func F_rioWriteBulkDouble(m *base.Module, l0 int32, l1 float64) int32
//go:linkname F_rioInitWithConnset github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rioInitWithConnset
func F_rioInitWithConnset(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rioFreeConnset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rioFreeConnset
func F_rioFreeConnset(m *base.Module, l0 int32)
//go:linkname F_scriptInterrupt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptInterrupt
func F_scriptInterrupt(m *base.Module, l0 int32) int32
//go:linkname F_scriptKill github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptKill
func F_scriptKill(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scriptingEngineManagerInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineManagerInit
func F_scriptingEngineManagerInit(m *base.Module) int32
//go:linkname F_scriptingEngineCallGetMemoryInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineCallGetMemoryInfo
func F_scriptingEngineCallGetMemoryInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scriptingEngineManagerFind github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineManagerFind
func F_scriptingEngineManagerFind(m *base.Module, l0 int32) int32
//go:linkname F_scriptingEngineManagerForEachEngine github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineManagerForEachEngine
func F_scriptingEngineManagerForEachEngine(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scriptingEngineCallGetFunctionMemoryOverhead github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineCallGetFunctionMemoryOverhead
func F_scriptingEngineCallGetFunctionMemoryOverhead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_scriptingEngineDebuggerEnable github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineDebuggerEnable
func F_scriptingEngineDebuggerEnable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scriptingEngineDebuggerDisable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineDebuggerDisable
func F_scriptingEngineDebuggerDisable(m *base.Module, l0 int32)
//go:linkname F_scriptingEngineDebuggerFlushLogs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineDebuggerFlushLogs
func F_scriptingEngineDebuggerFlushLogs(m *base.Module)
//go:linkname F_scriptingEngineDebuggerStartSession github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineDebuggerStartSession
func F_scriptingEngineDebuggerStartSession(m *base.Module, l0 int32) int32
//go:linkname F_scriptingEngineDebuggerEndSession github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_scriptingEngineDebuggerEndSession
func F_scriptingEngineDebuggerEndSession(m *base.Module, l0 int32)
//go:linkname F_sdsHdrSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsHdrSize
func F_sdsHdrSize(m *base.Module, l0 int32) int32
//go:linkname F_sdsnewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsnewlen
func F_sdsnewlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdstrynewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdstrynewlen
func F_sdstrynewlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsempty github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsempty
func F_sdsempty(m *base.Module) int32
//go:linkname F_sdsdup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsdup
func F_sdsdup(m *base.Module, l0 int32) int32
//go:linkname F_sdsfree github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfree
func F_sdsfree(m *base.Module, l0 int32)
//go:linkname F_sdsAllocPtr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsAllocPtr
func F_sdsAllocPtr(m *base.Module, l0 int32) int32
//go:linkname F_sdsMakeRoomFor github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsMakeRoomFor
func F_sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsResize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsResize
func F_sdsResize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsgrowzero github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsgrowzero
func F_sdsgrowzero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdscat github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscat
func F_sdscat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdscpylen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscpylen
func F_sdscpylen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscpy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdscpy
func F_sdscpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsfromlonglong github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfromlonglong
func F_sdsfromlonglong(m *base.Module, l0 int64) int32
//go:linkname F_sdscatvprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscatvprintf
func F_sdscatvprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscatprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdscatprintf
func F_sdscatprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdssplitlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdssplitlen
func F_sdssplitlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_sdsfreesplitres github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfreesplitres
func F_sdsfreesplitres(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sdscatrepr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdscatrepr
func F_sdscatrepr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsnsplitargs_internal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdsnsplitargs_internal
func F_sdsnsplitargs_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdssplitargs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sdssplitargs
func F_sdssplitargs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsjoin github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsjoin
func F_sdsjoin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelGenerateInitialMonitorEvents github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelGenerateInitialMonitorEvents
func F_sentinelGenerateInitialMonitorEvents(m *base.Module)
//go:linkname F_sentinelEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelEvent
func F_sentinelEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_sentinelScheduleScriptExecution github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelScheduleScriptExecution
func F_sentinelScheduleScriptExecution(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getSentinelValkeyInstanceByAddrAndRunID github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getSentinelValkeyInstanceByAddrAndRunID
func F_getSentinelValkeyInstanceByAddrAndRunID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sentinelDropConnections github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelDropConnections
func F_sentinelDropConnections(m *base.Module) int32
//go:linkname F_sentinelResetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelResetPrimary
func F_sentinelResetPrimary(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelPropagateDownAfterPeriod github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelPropagateDownAfterPeriod
func F_sentinelPropagateDownAfterPeriod(m *base.Module, l0 int32)
//go:linkname F_sentinelProcessHelloMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelProcessHelloMessage
func F_sentinelProcessHelloMessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sentinelKillClients github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelKillClients
func F_sentinelKillClients(m *base.Module, l0 int32) int32
//go:linkname F_sentinelForceHelloUpdateForPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelForceHelloUpdateForPrimary
func F_sentinelForceHelloUpdateForPrimary(m *base.Module, l0 int32) int32
//go:linkname F_sentinelSendReplicaOf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelSendReplicaOf
func F_sentinelSendReplicaOf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sentinelSimFailureCrash github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelSimFailureCrash
func F_sentinelSimFailureCrash(m *base.Module)
//go:linkname F_sentinelValidateArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelValidateArgs
func F_sentinelValidateArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelFlushConfigAndReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelFlushConfigAndReply
func F_sentinelFlushConfigAndReply(m *base.Module, l0 int32)
//go:linkname F_sentinelVoteLeader github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelVoteLeader
func F_sentinelVoteLeader(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32
//go:linkname F_sentinelSelectReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelSelectReplica
func F_sentinelSelectReplica(m *base.Module, l0 int32) int32
//go:linkname F_sentinelStartFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelStartFailover
func F_sentinelStartFailover(m *base.Module, l0 int32)
//go:linkname F_sentinelRoleCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelRoleCommand
func F_sentinelRoleCommand(m *base.Module, l0 int32)
//go:linkname F_sentinelLeaderIncr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelLeaderIncr
func F_sentinelLeaderIncr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sentinelAbortFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelAbortFailover
func F_sentinelAbortFailover(m *base.Module, l0 int32)
//go:linkname F_sentinelFailoverDetectEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelFailoverDetectEnd
func F_sentinelFailoverDetectEnd(m *base.Module, l0 int32)
//go:linkname F_sentinelHandleValkeyInstance github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sentinelHandleValkeyInstance
func F_sentinelHandleValkeyInstance(m *base.Module, l0 int32)
//go:linkname F_sentinelTimer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelTimer
func F_sentinelTimer(m *base.Module)
//go:linkname F_serverLogRawFromHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_serverLogRawFromHandler
func F_serverLogRawFromHandler(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clientsCronResizeQueryBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clientsCronResizeQueryBuffer
func F_clientsCronResizeQueryBuffer(m *base.Module, l0 int32) int32
//go:linkname F_clientsCronResizeOutputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clientsCronResizeOutputBuffer
func F_clientsCronResizeOutputBuffer(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_updateClientMemoryUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_updateClientMemoryUsage
func F_updateClientMemoryUsage(m *base.Module, l0 int32)
//go:linkname F_checkChildrenDone github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_checkChildrenDone
func F_checkChildrenDone(m *base.Module)
//go:linkname F_finishShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_finishShutdown
func F_finishShutdown(m *base.Module) int32
//go:linkname F_bytesToHuman github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_bytesToHuman
func F_bytesToHuman(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_createSharedObjects github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createSharedObjects
func F_createSharedObjects(m *base.Module)
//go:linkname F_restartServer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_restartServer
func F_restartServer(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_setOOMScoreAdj github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setOOMScoreAdj
func F_setOOMScoreAdj(m *base.Module, l0 int32) int32
//go:linkname F_dbHasNoKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_dbHasNoKeys
func F_dbHasNoKeys(m *base.Module, l0 int32) int32
//go:linkname F_createDatabase github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createDatabase
func F_createDatabase(m *base.Module, l0 int32) int32
//go:linkname F_createDatabaseIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_createDatabaseIfNeeded
func F_createDatabaseIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_populateCommandLegacyRangeSpec github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_populateCommandLegacyRangeSpec
func F_populateCommandLegacyRangeSpec(m *base.Module, l0 int32)
//go:linkname F_resetCommandTableStats github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resetCommandTableStats
func F_resetCommandTableStats(m *base.Module, l0 int32)
//go:linkname F_resetErrorTableStats github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resetErrorTableStats
func F_resetErrorTableStats(m *base.Module)
//go:linkname F_lookupCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupCommand
func F_lookupCommand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupCommandBySdsLogic github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lookupCommandBySdsLogic
func F_lookupCommandBySdsLogic(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupCommandByCString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lookupCommandByCString
func F_lookupCommandByCString(m *base.Module, l0 int32) int32
//go:linkname F_alsoPropagate github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_alsoPropagate
func F_alsoPropagate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_forceCommandPropagation github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_forceCommandPropagation
func F_forceCommandPropagation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_updateCommandLatencyHistogram github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_updateCommandLatencyHistogram
func F_updateCommandLatencyHistogram(m *base.Module, l0 int32, l1 int64)
//go:linkname F_afterCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_afterCommand
func F_afterCommand(m *base.Module, l0 int32)
//go:linkname F_commandCheckExistence github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_commandCheckExistence
func F_commandCheckExistence(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_processCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_processCommand
func F_processCommand(m *base.Module, l0 int32) int32
//go:linkname F_incrementErrorCount github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_incrementErrorCount
func F_incrementErrorCount(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyFlagsForKeyArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyFlagsForKeyArgs
func F_addReplyFlagsForKeyArgs(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addReplyCommandInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyCommandInfo
func F_addReplyCommandInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyCommandDocs github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addReplyCommandDocs
func F_addReplyCommandDocs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getKeysSubcommandImpl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysSubcommandImpl
func F_getKeysSubcommandImpl(m *base.Module, l0 int32, l1 int32)
//go:linkname F_invalidateCommandCache github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_invalidateCommandCache
func F_invalidateCommandCache(m *base.Module)
//go:linkname F_shouldFilterFromCommandList github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_shouldFilterFromCommandList
func F_shouldFilterFromCommandList(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fillPercentileDistributionLatencies github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fillPercentileDistributionLatencies
func F_fillPercentileDistributionLatencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_releaseInfoSectionDict github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_releaseInfoSectionDict
func F_releaseInfoSectionDict(m *base.Module, l0 int32)
//go:linkname F_genValkeyInfoString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_genValkeyInfoString
func F_genValkeyInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getVersion github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getVersion
func F_getVersion(m *base.Module) int32
//go:linkname F_serverFork github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_serverFork
func F_serverFork(m *base.Module, l0 int32) int32
//go:linkname F_parseExtendedCommandArgumentsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_parseExtendedCommandArgumentsOrReply
func F_parseExtendedCommandArgumentsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32
//go:linkname F_SHA1Init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_SHA1Init
func F_SHA1Init(m *base.Module, l0 int32)
//go:linkname F_SHA1Update github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_SHA1Update
func F_SHA1Update(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_SHA1Final github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_SHA1Final
func F_SHA1Final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sha256_init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sha256_init
func F_sha256_init(m *base.Module, l0 int32)
//go:linkname F_sha256_final github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sha256_final
func F_sha256_final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_siphash github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_siphash
func F_siphash(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_connBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_connBlock
func F_connBlock(m *base.Module, l0 int32) int32
//go:linkname F_sortCommandGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sortCommandGeneric
func F_sortCommandGeneric(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeHasVolatileFields github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeHasVolatileFields
func F_hashTypeHasVolatileFields(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeTryConversion github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeTryConversion
func F_hashTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hashTypeGetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeGetValue
func F_hashTypeGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_hashTypeExists github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeExists
func F_hashTypeExists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashTypeGetValueObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeGetValueObject
func F_hashTypeGetValueObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashTypeTrackUpdateEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeTrackUpdateEntry
func F_hashTypeTrackUpdateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64)
//go:linkname F_hashTypeSet github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeSet
func F_hashTypeSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32
//go:linkname F_hashTypeLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeLength
func F_hashTypeLength(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeDelete
func F_hashTypeDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashTypeInitVolatileIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeInitVolatileIterator
func F_hashTypeInitVolatileIterator(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeResetIterator
func F_hashTypeResetIterator(m *base.Module, l0 int32)
//go:linkname F_hashTypeNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_hashTypeNext
func F_hashTypeNext(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeCurrentFromListpack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeCurrentFromListpack
func F_hashTypeCurrentFromListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_hashTypeCurrentObjectNewSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeCurrentObjectNewSds
func F_hashTypeCurrentObjectNewSds(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addHashFieldToReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addHashFieldToReply
func F_addHashFieldToReply(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_genericHgetallCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_genericHgetallCommand
func F_genericHgetallCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hexpireGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hexpireGenericCommand
func F_hexpireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_listTypeTryConversionRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeTryConversionRaw
func F_listTypeTryConversionRaw(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)
//go:linkname F_listTypeTryConversionAppend github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeTryConversionAppend
func F_listTypeTryConversionAppend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_listTypePush github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypePush
func F_listTypePush(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_listTypePop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypePop
func F_listTypePop(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listTypeLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeLength
func F_listTypeLength(m *base.Module, l0 int32) int32
//go:linkname F_listTypeInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeInitIterator
func F_listTypeInitIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_listTypeSetIteratorDirection github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeSetIteratorDirection
func F_listTypeSetIteratorDirection(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_listTypeNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeNext
func F_listTypeNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listTypeReplace github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeReplace
func F_listTypeReplace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listTypeDelRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_listTypeDelRange
func F_listTypeDelRange(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addListRangeReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addListRangeReply
func F_addListRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_listElementsRemoved github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listElementsRemoved
func F_listElementsRemoved(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_lmoveGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lmoveGenericCommand
func F_lmoveGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lmpopGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lmpopGenericCommand
func F_lmpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_setTypeCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeCreate
func F_setTypeCreate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setTypeConvertAndExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setTypeConvertAndExpand
func F_setTypeConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setTypeSize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeSize
func F_setTypeSize(m *base.Module, l0 int32) int32
//go:linkname F_setTypeInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeInitIterator
func F_setTypeInitIterator(m *base.Module, l0 int32) int32
//go:linkname F_setTypeNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeNext
func F_setTypeNext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setTypeAddAux github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeAddAux
func F_setTypeAddAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_setTypeIsMemberAux github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setTypeIsMemberAux
func F_setTypeIsMemberAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_sunionDiffGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sunionDiffGenericCommand
func F_sunionDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_lpGetIntegerIfValid github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpGetIntegerIfValid
func F_lpGetIntegerIfValid(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_streamIteratorStop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamIteratorStop
func F_streamIteratorStop(m *base.Module, l0 int32)
//go:linkname F_streamTrim github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamTrim
func F_streamTrim(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_streamEntryExists github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamEntryExists
func F_streamEntryExists(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setDeferredReplyStreamID github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setDeferredReplyStreamID
func F_setDeferredReplyStreamID(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_streamPropagateXCLAIM github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamPropagateXCLAIM
func F_streamPropagateXCLAIM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_streamReplyWithRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_streamReplyWithRange
func F_streamReplyWithRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_streamParseIntervalIDOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamParseIntervalIDOrReply
func F_streamParseIntervalIDOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32
//go:linkname F_streamCreateConsumer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamCreateConsumer
func F_streamCreateConsumer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_setGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_setGenericCommand
func F_setGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_zslCreateNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zslCreateNode
func F_zslCreateNode(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F_zslParseLexRangeItem github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zslParseLexRangeItem
func F_zslParseLexRangeItem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zzlGetScore github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlGetScore
func F_zzlGetScore(m *base.Module, l0 int32) float64
//go:linkname F_zzlNext github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlNext
func F_zzlNext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zzlPrev github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlPrev
func F_zzlPrev(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zzlIsInRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zzlIsInRange
func F_zzlIsInRange(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetLength github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zsetLength
func F_zsetLength(m *base.Module, l0 int32) int32
//go:linkname F_zsetConvertAndExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zsetConvertAndExpand
func F_zsetConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zzlFind github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlFind
func F_zzlFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zsetAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zsetAdd
func F_zsetAdd(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_zsetRemoveFromSkiplist github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zsetRemoveFromSkiplist
func F_zsetRemoveFromSkiplist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zremrangeGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zremrangeGenericCommand
func F_zremrangeGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zslParseRange github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zslParseRange
func F_zslParseRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zunionInterDiffGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zunionInterDiffGenericCommand
func F_zunionInterDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_zrangeGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zrangeGenericCommand
func F_zrangeGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_addZpopInitialReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addZpopInitialReply
func F_addZpopInitialReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_blockingGenericZpopCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_blockingGenericZpopCommand
func F_blockingGenericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_clientsCronHandleTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_clientsCronHandleTimeout
func F_clientsCronHandleTimeout(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_disableTracking github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_disableTracking
func F_disableTracking(m *base.Module, l0 int32)
//go:linkname F_sendTrackingMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sendTrackingMessage
func F_sendTrackingMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_trackingRememberKeyToBroadcast github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trackingRememberKeyToBroadcast
func F_trackingRememberKeyToBroadcast(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_trackingLimitUsedSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trackingLimitUsedSlots
func F_trackingLimitUsedSlots(m *base.Module)
//go:linkname F_stringmatchlen_impl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_stringmatchlen_impl
func F_stringmatchlen_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_prefixmatchlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_prefixmatchlen
func F_prefixmatchlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_memtoull github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memtoull
func F_memtoull(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_ll2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ll2string
func F_ll2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_ull2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ull2string
func F_ull2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_ld2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ld2string
func F_ld2string(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32
//go:linkname F_getRandomBytes github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getRandomBytes
func F_getRandomBytes(m *base.Module, l0 int32, l1 int32)
//go:linkname F_makePath github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_makePath
func F_makePath(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_vsnprintf_async_signal_safe github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsnprintf_async_signal_safe
func F_vsnprintf_async_signal_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ustime github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ustime
func F_ustime(m *base.Module) int64
//go:linkname F_mstime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mstime
func F_mstime(m *base.Module) int64
//go:linkname F_checkSingleAof github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_checkSingleAof
func F_checkSingleAof(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_fileIsManifest github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fileIsManifest
func F_fileIsManifest(m *base.Module, l0 int32) int32
//go:linkname F_freeRdbProfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeRdbProfile
func F_freeRdbProfile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_computeDatasetProfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_computeDatasetProfile
func F_computeDatasetProfile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_rioRead_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rioRead_2
func F_rioRead_2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parseCheckRdbOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_parseCheckRdbOptions
func F_parseCheckRdbOptions(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ffc_digit_comp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ffc_digit_comp
func F_ffc_digit_comp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_vectorInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vectorInit
func F_vectorInit(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vectorGet github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vectorGet
func F_vectorGet(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_vectorPush github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vectorPush
func F_vectorPush(m *base.Module, l0 int32) int32
//go:linkname F_pvInsertAt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pvInsertAt
func F_pvInsertAt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsetRemoveEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vsetRemoveEntry
func F_vsetRemoveEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_removeFromBucket_VECTOR github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_removeFromBucket_VECTOR
func F_removeFromBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_removeFromBucket_HASHTABLE github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_removeFromBucket_HASHTABLE
func F_removeFromBucket_HASHTABLE(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsetRemoveExpired github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vsetRemoveExpired
func F_vsetRemoveExpired(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32
//go:linkname F_freeVsetBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_freeVsetBucket
func F_freeVsetBucket(m *base.Module, l0 int32)
//go:linkname F_vsetResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsetResetIterator
func F_vsetResetIterator(m *base.Module, l0 int32)
//go:linkname F_ziplistGet github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ziplistGet
func F_ziplistGet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ziplistValidateIntegrity github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ziplistValidateIntegrity
func F_ziplistValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_zcalloc_num github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zcalloc_num
func F_zcalloc_num(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zcalloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zcalloc_usable
func F_zcalloc_usable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkey_realloc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkey_realloc
func F_valkey_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zmalloc_usable_size github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zmalloc_usable_size
func F_zmalloc_usable_size(m *base.Module, l0 int32) int32
//go:linkname F_zstrdup github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zstrdup
func F_zstrdup(m *base.Module, l0 int32) int32
//go:linkname F_zmalloc_used_memory github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_zmalloc_used_memory
func F_zmalloc_used_memory(m *base.Module) int32
//go:linkname F_mpscEnqueue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mpscEnqueue
func F_mpscEnqueue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_spscInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_spscInit
func F_spscInit(m *base.Module, l0 int32, l1 int32)
//go:linkname F_valkeyAsyncFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncFree
func F_valkeyAsyncFree(m *base.Module, l0 int32)
//go:linkname F_valkeyAsyncFreeInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncFreeInternal
func F_valkeyAsyncFreeInternal(m *base.Module, l0 int32)
//go:linkname F_valkeyProcessCallbacks github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyProcessCallbacks
func F_valkeyProcessCallbacks(m *base.Module, l0 int32)
//go:linkname F_valkeyAsyncCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyAsyncCommand
func F_valkeyAsyncCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_valkeySetTcpNoDelay github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeySetTcpNoDelay
func F_valkeySetTcpNoDelay(m *base.Module, l0 int32) int32
//go:linkname F_valkeySetBlocking github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeySetBlocking
func F_valkeySetBlocking(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkeySetError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkeySetError
func F_valkeySetError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_valkeyFree github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_valkeyFree
func F_valkeyFree(m *base.Module, l0 int32)
//go:linkname F_valkeyConnectWithOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkeyConnectWithOptions
func F_valkeyConnectWithOptions(m *base.Module, l0 int32) int32
//go:linkname F_hdr_close github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hdr_close
func F_hdr_close(m *base.Module, l0 int32)
//go:linkname F_ldbLog github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ldbLog
func F_ldbLog(m *base.Module, l0 int32)
//go:linkname F_ldbLogSourceLine github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ldbLogSourceLine
func F_ldbLogSourceLine(m *base.Module, l0 int32)
//go:linkname F_ldbRepl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ldbRepl
func F_ldbRepl(m *base.Module, l0 int32) int32
//go:linkname F_initializeLuaState github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_initializeLuaState
func F_initializeLuaState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_list_get_iter github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_list_get_iter
func F_list_get_iter(m *base.Module, l0 int32) int32
//go:linkname F_list_release_iter github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_list_release_iter
func F_list_release_iter(m *base.Module, l0 int32)
//go:linkname F_luaSaveOnRegistry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaSaveOnRegistry
func F_luaSaveOnRegistry(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaGetFromRegistry github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaGetFromRegistry
func F_luaGetFromRegistry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lm_asprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lm_asprintf
func F_lm_asprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaPushErrorBuff github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaPushErrorBuff
func F_luaPushErrorBuff(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaError github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaError
func F_luaError(m *base.Module, l0 int32) int32
//go:linkname F_luaRegisterVersion github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaRegisterVersion
func F_luaRegisterVersion(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaRegisterLogFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaRegisterLogFunction
func F_luaRegisterLogFunction(m *base.Module, l0 int32)
//go:linkname F_luaErrorInformationDiscard github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaErrorInformationDiscard
func F_luaErrorInformationDiscard(m *base.Module, l0 int32)
//go:linkname F_luaExtractErrorInformation github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaExtractErrorInformation
func F_luaExtractErrorInformation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaReplyToServerReply github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaReplyToServerReply
func F_luaReplyToServerReply(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F__serverPanic_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__serverPanic_2
func F__serverPanic_2(m *base.Module, l0 int32)
//go:linkname F_processCollectionElementEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processCollectionElementEnd
func F_processCollectionElementEnd(m *base.Module, l0 int32)
//go:linkname F_lua_checkstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_checkstack
func F_lua_checkstack(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_newthread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_newthread
func F_lua_newthread(m *base.Module, l0 int32) int32
//go:linkname F_lua_isnumber github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_isnumber
func F_lua_isnumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_tonumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_tonumber
func F_lua_tonumber(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_lua_tointeger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_tointeger
func F_lua_tointeger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_objlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_objlen
func F_lua_objlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_pushlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pushlstring
func F_lua_pushlstring(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_pushstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pushstring
func F_lua_pushstring(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_pushfstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_pushfstring
func F_lua_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_pushcclosure github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_pushcclosure
func F_lua_pushcclosure(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_createtable github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_createtable
func F_lua_createtable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_setfield github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_setfield
func F_lua_setfield(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_rawseti github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_rawseti
func F_lua_rawseti(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_setmetatable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_setmetatable
func F_lua_setmetatable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_pcall github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_pcall
func F_lua_pcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lua_gc github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_gc
func F_lua_gc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_error github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_error
func F_lua_error(m *base.Module, l0 int32) int32
//go:linkname F_lua_concat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_concat
func F_lua_concat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_getinfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_getinfo
func F_lua_getinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getobjname github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getobjname
func F_getobjname(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaG_runerror github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaG_runerror
func F_luaG_runerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaG_ordererror github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaG_ordererror
func F_luaG_ordererror(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaD_throw github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_throw
func F_luaD_throw(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaD_growstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_growstack
func F_luaD_growstack(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaD_call github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaD_call
func F_luaD_call(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaF_freeupval github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaF_freeupval
func F_luaF_freeupval(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaF_newproto github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaF_newproto
func F_luaF_newproto(m *base.Module, l0 int32) int32
//go:linkname F_luaF_freeproto github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaF_freeproto
func F_luaF_freeproto(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaF_freeclosure github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaF_freeclosure
func F_luaF_freeclosure(m *base.Module, l0 int32, l1 int32)
//go:linkname F_GCTM github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_GCTM
func F_GCTM(m *base.Module, l0 int32)
//go:linkname F_luaC_freeall github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaC_freeall
func F_luaC_freeall(m *base.Module, l0 int32)
//go:linkname F_propagatemark github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_propagatemark
func F_propagatemark(m *base.Module, l0 int32) int32
//go:linkname F_reallymarkobject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_reallymarkobject
func F_reallymarkobject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_markmt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_markmt
func F_markmt(m *base.Module, l0 int32)
//go:linkname F_luaC_barrierf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaC_barrierf
func F_luaC_barrierf(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaM_toobig github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaM_toobig
func F_luaM_toobig(m *base.Module, l0 int32) int32
//go:linkname F_luaO_pushfstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaO_pushfstring
func F_luaO_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaX_token2str github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_token2str
func F_luaX_token2str(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaX_lexerror github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_lexerror
func F_luaX_lexerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_save github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_save
func F_save(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaX_newstring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaX_newstring
func F_luaX_newstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaX_next github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaX_next
func F_luaX_next(m *base.Module, l0 int32)
//go:linkname F_skip_sep github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_skip_sep
func F_skip_sep(m *base.Module, l0 int32) int32
//go:linkname F_read_numeral github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_read_numeral
func F_read_numeral(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_code github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_code
func F_luaK_code(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaK_codeABC github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_codeABC
func F_luaK_codeABC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_luaK_jump github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_jump
func F_luaK_jump(m *base.Module, l0 int32) int32
//go:linkname F_luaK_codeABx github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_codeABx
func F_luaK_codeABx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaK_ret github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_ret
func F_luaK_ret(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_patchlist github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_patchlist
func F_luaK_patchlist(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_patchtohere github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_patchtohere
func F_luaK_patchtohere(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_checkstack github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_checkstack
func F_luaK_checkstack(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_stringK github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_stringK
func F_luaK_stringK(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addk github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_addk
func F_addk(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaK_setreturns github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_setreturns
func F_luaK_setreturns(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_dischargevars github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_dischargevars
func F_luaK_dischargevars(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_exp2nextreg github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_exp2nextreg
func F_luaK_exp2nextreg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_exp2reg github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_exp2reg
func F_exp2reg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_exp2anyreg github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_exp2anyreg
func F_luaK_exp2anyreg(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaK_storevar github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_storevar
func F_luaK_storevar(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_self github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_self
func F_luaK_self(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_goiftrue github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_goiftrue
func F_luaK_goiftrue(m *base.Module, l0 int32, l1 int32)
//go:linkname F_jumponcond github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_jumponcond
func F_jumponcond(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaK_indexed github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaK_indexed
func F_luaK_indexed(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_setlist github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_setlist
func F_luaK_setlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_close_func github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_close_func
func F_close_func(m *base.Module, l0 int32)
//go:linkname F_subexpr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_subexpr
func F_subexpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_check_match github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_check_match
func F_check_match(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_block github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_block
func F_block(m *base.Module, l0 int32)
//go:linkname F_new_localvar github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_new_localvar
func F_new_localvar(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_forbody github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_forbody
func F_forbody(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_singlevar github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_singlevar
func F_singlevar(m *base.Module, l0 int32, l1 int32)
//go:linkname F_body github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_body
func F_body(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_assignment github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_assignment
func F_assignment(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_adjust_assign github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_adjust_assign
func F_adjust_assign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_luaE_freethread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaE_freethread
func F_luaE_freethread(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_close github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_lua_close
func F_lua_close(m *base.Module, l0 int32)
//go:linkname F_luaS_resize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaS_resize
func F_luaS_resize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaH_next github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_next
func F_luaH_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_resize_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_resize_2
func F_resize_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_newkey github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_newkey
func F_newkey(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_new github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_new
func F_luaH_new(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_free github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_free
func F_luaH_free(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaH_set github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_set
func F_luaH_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaT_gettm github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaT_gettm
func F_luaT_gettm(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaV_tostring github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaV_tostring
func F_luaV_tostring(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_call_orderTM github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_call_orderTM
func F_call_orderTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_error github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_error
func F_luaL_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_where github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_where
func F_luaL_where(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaL_optlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_optlstring
func F_luaL_optlstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_checkstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checkstack
func F_luaL_checkstack(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_checktype github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checktype
func F_luaL_checktype(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_checkany github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_checkany
func F_luaL_checkany(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaL_checknumber github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_checknumber
func F_luaL_checknumber(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_luaL_checkinteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checkinteger
func F_luaL_checkinteger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaL_optinteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_optinteger
func F_luaL_optinteger(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_getmetafield github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_getmetafield
func F_luaL_getmetafield(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_register github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_register
func F_luaL_register(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_findtable github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_findtable
func F_luaL_findtable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_prepbuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_prepbuffer
func F_luaL_prepbuffer(m *base.Module, l0 int32) int32
//go:linkname F_luaL_addlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_addlstring
func F_luaL_addlstring(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_pushresult github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_luaL_pushresult
func F_luaL_pushresult(m *base.Module, l0 int32)
//go:linkname F_luaL_addvalue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_addvalue
func F_luaL_addvalue(m *base.Module, l0 int32)
//go:linkname F_luaL_loadfile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_loadfile
func F_luaL_loadfile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaL_loadbuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_loadbuffer
func F_luaL_loadbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getfunc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getfunc
func F_getfunc(m *base.Module, l0 int32, l1 int32)
//go:linkname F_auxresume github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_auxresume
func F_auxresume(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_str_find_aux github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_str_find_aux
func F_str_find_aux(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_match github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_match
func F_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_push_onecapture github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_push_onecapture
func F_push_onecapture(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_fpconv_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fpconv_init
func F_fpconv_init(m *base.Module)
//go:linkname F_strbuf_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strbuf_init
func F_strbuf_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_die github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_die
func F_die(m *base.Module, l0 int32, l1 int32)
//go:linkname F_strbuf_free github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strbuf_free
func F_strbuf_free(m *base.Module, l0 int32)
//go:linkname F_json_next_token github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_json_next_token
func F_json_next_token(m *base.Module, l0 int32, l1 int32)
//go:linkname F_json_enum_option github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_json_enum_option
func F_json_enum_option(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_append_string github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_json_append_string
func F_json_append_string(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_append_number github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_json_append_number
func F_json_append_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mp_encode_array github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mp_encode_array
func F_mp_encode_array(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_mp_encode_map github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_mp_encode_map
func F_mp_encode_map(m *base.Module, l0 int32, l1 int32, l2 int64)
//go:linkname F_mp_encode_lua_type github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mp_encode_lua_type
func F_mp_encode_lua_type(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mp_decode_to_lua_type github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mp_decode_to_lua_type
func F_mp_decode_to_lua_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_mp_unpack_full github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_mp_unpack_full
func F_mp_unpack_full(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getpwnam_r github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getpwnam_r
func F_getpwnam_r(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F___errno_location github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___errno_location
func F___errno_location(m *base.Module) int32
//go:linkname F_do_tzset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_do_tzset
func F_do_tzset(m *base.Module)
//go:linkname F__exit github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__exit
func F__exit(m *base.Module, l0 int32)
//go:linkname F__Exit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__Exit
func F__Exit(m *base.Module, l0 int32)
//go:linkname F_access github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_access
func F_access(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_asin github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_asin
func F_asin(m *base.Module, l0 float64) float64
//go:linkname F_R_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_R_2
func F_R_2(m *base.Module, l0 float64) float64
//go:linkname F___isspace_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___isspace_1
func F___isspace_1(m *base.Module, l0 int32) int32
//go:linkname F_chdir github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_chdir
func F_chdir(m *base.Module, l0 int32) int32
//go:linkname F_chmod github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_chmod
func F_chmod(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___cos github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___cos
func F___cos(m *base.Module, l0 float64, l1 float64) float64
//go:linkname F___sin github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___sin
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64
//go:linkname F_cos github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_cos
func F_cos(m *base.Module, l0 float64) float64
//go:linkname F_ctime_r github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ctime_r
func F_ctime_r(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dup2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dup2
func F_dup2(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_memmove github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memmove
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___time github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___time
func F___time(m *base.Module, l0 int32) int64
//go:linkname F___gettimeofday github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___gettimeofday
func F___gettimeofday(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___math_xflow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___math_xflow
func F___math_xflow(m *base.Module, l0 int32, l1 float64) float64
//go:linkname F___math_uflow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___math_uflow
func F___math_uflow(m *base.Module, l0 int32) float64
//go:linkname F___math_oflow github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___math_oflow
func F___math_oflow(m *base.Module, l0 int32) float64
//go:linkname F_exp github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_exp
func F_exp(m *base.Module, l0 float64) float64
//go:linkname F_top12_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_top12_1
func F_top12_1(m *base.Module, l0 float64) int32
//go:linkname F_specialcase_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_specialcase_1
func F_specialcase_1(m *base.Module, l0 float64, l1 int64, l2 int64) float64
//go:linkname F_fabs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fabs
func F_fabs(m *base.Module, l0 float64) float64
//go:linkname F___lockfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___lockfile
func F___lockfile(m *base.Module, l0 int32) int32
//go:linkname F___unlockfile github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___unlockfile
func F___unlockfile(m *base.Module, l0 int32)
//go:linkname F_fclose github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fclose
func F_fclose(m *base.Module, l0 int32) int32
//go:linkname F_fcntl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fcntl
func F_fcntl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fflush github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F___toread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___toread
func F___toread(m *base.Module, l0 int32) int32
//go:linkname F___uflow github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___uflow
func F___uflow(m *base.Module, l0 int32) int32
//go:linkname F_fgets github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fgets
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fopen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fopen
func F_fopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fiprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fiprintf
func F_fiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___overflow github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___overflow
func F___overflow(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputc github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fputc
func F_fputc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_locking_putc_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_locking_putc_1
func F_locking_putc_1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fread github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fread
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___fseeko_unlocked github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___fseeko_unlocked
func F___fseeko_unlocked(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_fseek github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fseek
func F_fseek(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fsync github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fsync
func F_fsync(m *base.Module, l0 int32) int32
//go:linkname F___ftello_unlocked github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___ftello_unlocked
func F___ftello_unlocked(m *base.Module, l0 int32) int64
//go:linkname F_ftruncate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_ftruncate
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_fwrite github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_do_getc github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_do_getc
func F_do_getc(m *base.Module, l0 int32) int32
//go:linkname F___syscall_getuid32 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___syscall_getuid32
func F___syscall_getuid32(m *base.Module) int32
//go:linkname F___syscall_shutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___syscall_shutdown
func F___syscall_shutdown(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F___syscall_wait4 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___syscall_wait4
func F___syscall_wait4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getpid github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getpid
func F_getpid(m *base.Module) int32
//go:linkname F_do_glob github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_do_glob
func F_do_glob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_htonl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_htonl
func F_htonl(m *base.Module, l0 int32) int32
//go:linkname F_htons github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_htons
func F_htons(m *base.Module, l0 int32) int32
//go:linkname F___bswap_16_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___bswap_16_1
func F___bswap_16_1(m *base.Module, l0 int32) int32
//go:linkname F_inet_pton github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_inet_pton
func F_inet_pton(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_iswalpha github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_iswalpha
func F_iswalpha(m *base.Module, l0 int32) int32
//go:linkname F_iswalnum github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_iswalnum
func F_iswalnum(m *base.Module, l0 int32) int32
//go:linkname F_iswcntrl github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_iswcntrl
func F_iswcntrl(m *base.Module, l0 int32) int32
//go:linkname F_iswgraph github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_iswgraph
func F_iswgraph(m *base.Module, l0 int32) int32
//go:linkname F_iswlower github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswlower
func F_iswlower(m *base.Module, l0 int32) int32
//go:linkname F_iswprint github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswprint
func F_iswprint(m *base.Module, l0 int32) int32
//go:linkname F_iswpunct github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswpunct
func F_iswpunct(m *base.Module, l0 int32) int32
//go:linkname F_iswupper github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswupper
func F_iswupper(m *base.Module, l0 int32) int32
//go:linkname F_iswxdigit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswxdigit
func F_iswxdigit(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_emscripten_futex_wake github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_futex_wake
func F_emscripten_futex_wake(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pthread_mutex_init github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pthread_mutex_init
func F_pthread_mutex_init(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___pthread_mutex_lock github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___pthread_mutex_lock
func F___pthread_mutex_lock(m *base.Module, l0 int32) int32
//go:linkname F___pthread_mutex_unlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___pthread_mutex_unlock
func F___pthread_mutex_unlock(m *base.Module, l0 int32) int32
//go:linkname F_pthread_mutexattr_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pthread_mutexattr_init
func F_pthread_mutexattr_init(m *base.Module, l0 int32) int32
//go:linkname F_pthread_setcancelstate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pthread_setcancelstate
func F_pthread_setcancelstate(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___unlock github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___unlock
func F___unlock(m *base.Module, l0 int32)
//go:linkname F_emscripten_thread_sleep github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_thread_sleep
func F_emscripten_thread_sleep(m *base.Module, l0 float64)
//go:linkname F_fp_barrier_3 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fp_barrier_3
func F_fp_barrier_3(m *base.Module, l0 float64) float64
//go:linkname F_memchr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___syscall_mmap2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_mmap2
func F___syscall_mmap2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64) int32
//go:linkname F_nanosleep github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_nanosleep
func F_nanosleep(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___bswap_32_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___bswap_32_2
func F___bswap_32_2(m *base.Module, l0 int32) int32
//go:linkname F___bswap_16_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___bswap_16_2
func F___bswap_16_2(m *base.Module, l0 int32) int32
//go:linkname F_perror github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_perror
func F_perror(m *base.Module, l0 int32)
//go:linkname F_poll github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_poll
func F_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fp_barrier_4 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_fp_barrier_4
func F_fp_barrier_4(m *base.Module, l0 float64) float64
//go:linkname F_iprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iprintf
func F_iprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pthread_attr_setstacksize github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_pthread_attr_setstacksize
func F_pthread_attr_setstacksize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_do_putc_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_do_putc_2
func F_do_putc_2(m *base.Module, l0 int32) int32
//go:linkname F_puts github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_puts
func F_puts(m *base.Module, l0 int32) int32
//go:linkname F_sift github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sift
func F_sift(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_trinkle github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trinkle
func F_trinkle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)
//go:linkname F_qsort github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_qsort
func F_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_rand github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_rand
func F_rand(m *base.Module) int32
//go:linkname F_lcg31 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lcg31
func F_lcg31(m *base.Module, l0 int32) int32
//go:linkname F_read github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_read
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rename github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_roundl github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_roundl
func F_roundl(m *base.Module, l0 int32, l1 int64, l2 int64)
//go:linkname F__emscripten_check_timers github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F__emscripten_check_timers
func F__emscripten_check_timers(m *base.Module, l0 float64)
//go:linkname F_setlocale github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setlocale
func F_setlocale(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___sigaction github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___sigaction
func F___sigaction(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sigemptyset github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sigemptyset
func F_sigemptyset(m *base.Module, l0 int32) int32
//go:linkname F_siprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_siprintf
func F_siprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sqrt github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sqrt
func F_sqrt(m *base.Module, l0 float64) float64
//go:linkname F_stat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_stat
func F_stat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcasecmp github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strcasecmp
func F_strcasecmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strcat
func F_strcat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strchrnul github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strchrnul
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strcoll_l github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strcoll_l
func F___strcoll_l(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___stpcpy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___stpcpy
func F___stpcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcpy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strcpy
func F_strcpy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcspn github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strcspn
func F_strcspn(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___strerror_l github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___strerror_l
func F___strerror_l(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strerror github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strerror
func F_strerror(m *base.Module, l0 int32) int32
//go:linkname F___strftime_l github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___strftime_l
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_strlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strlen
func F_strlen(m *base.Module, l0 int32) int32
//go:linkname F_strncat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strncat
func F_strncat(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncmp github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strncmp
func F_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___stpncpy github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___stpncpy
func F___stpncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strnlen github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strnlen
func F_strnlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___memrchr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___memrchr
func F___memrchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_twobyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_twobyte_strstr
func F_twobyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_threebyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_threebyte_strstr
func F_threebyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___floatscan github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___floatscan
func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_strtox_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strtox_1
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_strtod github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strtod
func F_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strtox_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_strtox_2
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F___syscall_ret github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_sysconf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sysconf
func F_sysconf(m *base.Module, l0 int32) int32
//go:linkname F___openlog github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___openlog
func F___openlog(m *base.Module)
//go:linkname F_syslog github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_syslog
func F_syslog(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_casemap github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_casemap
func F_casemap(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_truncate github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_truncate
func F_truncate(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F___vfprintf_internal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___vfprintf_internal
func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_out github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_out
func F_out(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_vsnprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsnprintf
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___small_vsprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___small_vsprintf
func F___small_vsprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___intscan github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___intscan
func F___intscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64
//go:linkname F_vsscanf github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_vsscanf
func F_vsscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_wcslen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_wcslen
func F_wcslen(m *base.Module, l0 int32) int32
//go:linkname F_write github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_write
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_free github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_emscripten_builtin_free
func F_emscripten_builtin_free(m *base.Module, l0 int32)
//go:linkname F_try_realloc_chunk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_try_realloc_chunk
func F_try_realloc_chunk(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_emscripten_builtin_calloc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_emscripten_builtin_calloc
func F_emscripten_builtin_calloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sbrk github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_sbrk
func F_sbrk(m *base.Module, l0 int32) int32
//go:linkname F___addtf3 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___addtf3
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___divtf3 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___divtf3
func F___divtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___fixtfdi github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___fixtfdi
func F___fixtfdi(m *base.Module, l0 int64, l1 int64) int64
//go:linkname F___lshrti3 github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F___lshrti3
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___multf3 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___multf3
func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___trunctfdf2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___trunctfdf2
func F___trunctfdf2(m *base.Module, l0 int64, l1 int64) float64
//go:linkname F_getsockopt github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_getsockopt
func F_getsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_socket github.com/shibukawa/vkmem/internal/aot/vkaot/p1.F_socket
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32
