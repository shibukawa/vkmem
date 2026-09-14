package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	_ "unsafe"
)
//go:linkname F___syscall_setsockopt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_setsockopt
func F___syscall_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ACLCreateUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLCreateUser
func F_ACLCreateUser(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLFreeUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLFreeUser
func F_ACLFreeUser(m *base.Module, l0 int32)
//go:linkname F_ACLFreeUserAndKillClients github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLFreeUserAndKillClients
func F_ACLFreeUserAndKillClients(m *base.Module, l0 int32)
//go:linkname F_ACLRecomputeCommandBitsFromCommandRulesAllUsers github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLRecomputeCommandBitsFromCommandRulesAllUsers
func F_ACLRecomputeCommandBitsFromCommandRulesAllUsers(m *base.Module)
//go:linkname F_ACLChangeSelectorPerm github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLChangeSelectorPerm
func F_ACLChangeSelectorPerm(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ACLAddAllowedFirstArg github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLAddAllowedFirstArg
func F_ACLAddAllowedFirstArg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_ACLModuleHasCommandRules github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLModuleHasCommandRules
func F_ACLModuleHasCommandRules(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsCatPatternString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsCatPatternString
func F_sdsCatPatternString(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLDescribeSelectorCommandRules github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLDescribeSelectorCommandRules
func F_ACLDescribeSelectorCommandRules(m *base.Module, l0 int32) int32
//go:linkname F_ACLSetUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSetUser
func F_ACLSetUser(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ACLInit github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLInit
func F_ACLInit(m *base.Module)
//go:linkname F_ACLAuthenticateUser github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLAuthenticateUser
func F_ACLAuthenticateUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ACLGetCommandID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLGetCommandID
func F_ACLGetCommandID(m *base.Module, l0 int32) int32
//go:linkname F_ACLSelectorCheckCmd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLSelectorCheckCmd
func F_ACLSelectorCheckCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_ACLCheckAllUserCommandPerm github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLCheckAllUserCommandPerm
func F_ACLCheckAllUserCommandPerm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_ACLCheckAllPerm github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLCheckAllPerm
func F_ACLCheckAllPerm(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLStringSetUser github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLStringSetUser
func F_ACLStringSetUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ACLMergeSelectorArguments github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLMergeSelectorArguments
func F_ACLMergeSelectorArguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ACLCopyUser github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ACLCopyUser
func F_ACLCopyUser(m *base.Module, l0 int32, l1 int32)
//go:linkname F_ACLShouldKillPubsubClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLShouldKillPubsubClient
func F_ACLShouldKillPubsubClient(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ACLLoadUsersAtStartup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ACLLoadUsersAtStartup
func F_ACLLoadUsersAtStartup(m *base.Module)
//go:linkname F_getAclErrorMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getAclErrorMessage
func F_getAclErrorMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_aclAddReplySelectorDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aclAddReplySelectorDescription
func F_aclAddReplySelectorDescription(m *base.Module, l0 int32, l1 int32)
//go:linkname F_aclCatWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aclCatWithFlags
func F_aclCatWithFlags(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32)
//go:linkname F_addReplyCommandCategories github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyCommandCategories
func F_addReplyCommandCategories(m *base.Module, l0 int32, l1 int32)
//go:linkname F_listEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listEmpty
func F_listEmpty(m *base.Module, l0 int32)
//go:linkname F_listAddNodeTail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listAddNodeTail
func F_listAddNodeTail(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listGetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listGetIterator
func F_listGetIterator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_listReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listReleaseIterator
func F_listReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_listNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listNext
func F_listNext(m *base.Module, l0 int32) int32
//go:linkname F_aeDeleteEventLoop github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aeDeleteEventLoop
func F_aeDeleteEventLoop(m *base.Module, l0 int32)
//go:linkname F_aeDeleteFileEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aeDeleteFileEvent
func F_aeDeleteFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_aeApiPoll github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aeApiPoll
func F_aeApiPoll(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aeSetCustomPollProc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aeSetCustomPollProc
func F_aeSetCustomPollProc(m *base.Module, l0 int32, l1 int32)
//go:linkname F_aeSetPollProtect github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aeSetPollProtect
func F_aeSetPollProtect(m *base.Module, l0 int32, l1 int32)
//go:linkname F_anetSetError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetSetError
func F_anetSetError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_anetNonBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetNonBlock
func F_anetNonBlock(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_anetKeepAlive github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetKeepAlive
func F_anetKeepAlive(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_anetResolve github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetResolve
func F_anetResolve(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetTcpGenericConnect github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetTcpGenericConnect
func F_anetTcpGenericConnect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetListen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetListen
func F_anetListen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_anetTcp6Server github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetTcp6Server
func F_anetTcp6Server(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetTcpAccept github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_anetTcpAccept
func F_anetTcpAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_anetGenericAccept github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_anetGenericAccept
func F_anetGenericAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_aofInfoFormat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofInfoFormat
func F_aofInfoFormat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_aofManifestFree github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofManifestFree
func F_aofManifestFree(m *base.Module, l0 int32)
//go:linkname F_aofLoadManifestFromDisk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofLoadManifestFromDisk
func F_aofLoadManifestFromDisk(m *base.Module)
//go:linkname F_aofManifestFreeAndUpdate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aofManifestFreeAndUpdate
func F_aofManifestFreeAndUpdate(m *base.Module, l0 int32)
//go:linkname F_getLastIncrAofName github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getLastIncrAofName
func F_getLastIncrAofName(m *base.Module, l0 int32) int32
//go:linkname F_markRewrittenIncrAofAsHistory github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_markRewrittenIncrAofAsHistory
func F_markRewrittenIncrAofAsHistory(m *base.Module, l0 int32)
//go:linkname F_writeAofManifestFile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_writeAofManifestFile
func F_writeAofManifestFile(m *base.Module, l0 int32) int32
//go:linkname F_aofDelHistoryFiles github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofDelHistoryFiles
func F_aofDelHistoryFiles(m *base.Module) int32
//go:linkname F_rewriteAppendOnlyFileRio github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteAppendOnlyFileRio
func F_rewriteAppendOnlyFileRio(m *base.Module, l0 int32) int32
//go:linkname F_killAppendOnlyChild github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_killAppendOnlyChild
func F_killAppendOnlyChild(m *base.Module)
//go:linkname F_aofRemoveTempFile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_aofRemoveTempFile
func F_aofRemoveTempFile(m *base.Module, l0 int32, l1 int32)
//go:linkname F_flushAppendOnlyFile github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_flushAppendOnlyFile
func F_flushAppendOnlyFile(m *base.Module, l0 int32)
//go:linkname F_startAppendOnly github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_startAppendOnly
func F_startAppendOnly(m *base.Module) int32
//go:linkname F_rewriteAppendOnlyFileBackground github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteAppendOnlyFileBackground
func F_rewriteAppendOnlyFileBackground(m *base.Module) int32
//go:linkname F_loadAppendOnlyFiles github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_loadAppendOnlyFiles
func F_loadAppendOnlyFiles(m *base.Module, l0 int32) int32
//go:linkname F_rioWriteBulkObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkObject
func F_rioWriteBulkObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteSlotToAppendOnlyFileRio github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteSlotToAppendOnlyFileRio
func F_rewriteSlotToAppendOnlyFileRio(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_bioInit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_bioInit
func F_bioInit(m *base.Module)
//go:linkname F_bioExecuteJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bioExecuteJob
func F_bioExecuteJob(m *base.Module, l0 int32)
//go:linkname F_bioDrainWorker github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bioDrainWorker
func F_bioDrainWorker(m *base.Module, l0 int32)
//go:linkname F_getObjectReadOnlyString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getObjectReadOnlyString
func F_getObjectReadOnlyString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_updateStatsOnUnblock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateStatsOnUnblock
func F_updateStatsOnUnblock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_processUnblockedClients github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processUnblockedClients
func F_processUnblockedClients(m *base.Module)
//go:linkname F_unblockClientWaitingData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unblockClientWaitingData
func F_unblockClientWaitingData(m *base.Module, l0 int32)
//go:linkname F_handleClientsBlockedOnKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_handleClientsBlockedOnKeys
func F_handleClientsBlockedOnKeys(m *base.Module)
//go:linkname F_signalKeyAsReady github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_signalKeyAsReady
func F_signalKeyAsReady(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_signalKeyAsReadyLogic github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_signalKeyAsReadyLogic
func F_signalKeyAsReadyLogic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_freeCallReplyInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeCallReplyInternal
func F_freeCallReplyInternal(m *base.Module, l0 int32)
//go:linkname F_callReplyType github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_callReplyType
func F_callReplyType(m *base.Module, l0 int32) int32
//go:linkname F_callReplyGetLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_callReplyGetLongLong
func F_callReplyGetLongLong(m *base.Module, l0 int32) int64
//go:linkname F_callReplyGetSetElement github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_callReplyGetSetElement
func F_callReplyGetSetElement(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_callReplyGetVerbatim github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_callReplyGetVerbatim
func F_callReplyGetVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_callReplyCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_callReplyCreate
func F_callReplyCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_receiveChildInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_receiveChildInfo
func F_receiveChildInfo(m *base.Module)
//go:linkname F_clearCachedClusterSlotsResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clearCachedClusterSlotsResponse
func F_clearCachedClusterSlotsResponse(m *base.Module)
//go:linkname F_generateClusterSlotResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_generateClusterSlotResponse
func F_generateClusterSlotResponse(m *base.Module, l0 int32) int32
//go:linkname F_clusterRedirectBlockedClientIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterRedirectBlockedClientIfNeeded
func F_clusterRedirectBlockedClientIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_getMigratingSlotDest github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getMigratingSlotDest
func F_getMigratingSlotDest(m *base.Module, l0 int32) int32
//go:linkname F_getImportingSlotSource github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getImportingSlotSource
func F_getImportingSlotSource(m *base.Module, l0 int32) int32
//go:linkname F_clusterRemoveNodeFromShard github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterRemoveNodeFromShard
func F_clusterRemoveNodeFromShard(m *base.Module, l0 int32)
//go:linkname F_clusterLoadConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterLoadConfig
func F_clusterLoadConfig(m *base.Module, l0 int32) int32
//go:linkname F_createClusterNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createClusterNode
func F_createClusterNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterLookupNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterLookupNode
func F_clusterLookupNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterSaveConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSaveConfig
func F_clusterSaveConfig(m *base.Module, l0 int32) int32
//go:linkname F_clusterGenNodesDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterGenNodesDescription
func F_clusterGenNodesDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clusterGenNodeDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterGenNodeDescription
func F_clusterGenNodeDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_clusterLockConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterLockConfig
func F_clusterLockConfig(m *base.Module, l0 int32) int32
//go:linkname F_clusterUpdateMyselfFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterUpdateMyselfFlags
func F_clusterUpdateMyselfFlags(m *base.Module)
//go:linkname F_clusterDoBeforeSleep github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterDoBeforeSleep
func F_clusterDoBeforeSleep(m *base.Module, l0 int32)
//go:linkname F_clusterUpdateMyselfAnnouncedPorts github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateMyselfAnnouncedPorts
func F_clusterUpdateMyselfAnnouncedPorts(m *base.Module)
//go:linkname F_updateSdsExtensionField github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateSdsExtensionField
func F_updateSdsExtensionField(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterUpdateMyselfClientIpV4 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterUpdateMyselfClientIpV4
func F_clusterUpdateMyselfClientIpV4(m *base.Module)
//go:linkname F_clusterAutoFailoverOnShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterAutoFailoverOnShutdown
func F_clusterAutoFailoverOnShutdown(m *base.Module)
//go:linkname F_freeClusterNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClusterNode
func F_freeClusterNode(m *base.Module, l0 int32)
//go:linkname F_clusterNodeAddFailureReport github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterNodeAddFailureReport
func F_clusterNodeAddFailureReport(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterNodeCleanupFailureReports github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterNodeCleanupFailureReports
func F_clusterNodeCleanupFailureReports(m *base.Module, l0 int32)
//go:linkname F_markNodeAsFailingIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_markNodeAsFailingIfNeeded
func F_markNodeAsFailingIfNeeded(m *base.Module, l0 int32)
//go:linkname F_clusterBroadcastMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterBroadcastMessage
func F_clusterBroadcastMessage(m *base.Module, l0 int32)
//go:linkname F_verifyGossipSectionNodeIds github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_verifyGossipSectionNodeIds
func F_verifyGossipSectionNodeIds(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_representClusterNodeFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_representClusterNodeFlags
func F_representClusterNodeFlags(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterSetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSetPrimary
func F_clusterSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clusterProcessPacket github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterProcessPacket
func F_clusterProcessPacket(m *base.Module, l0 int32) int32
//go:linkname F_clusterNodeIterNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterNodeIterNext
func F_clusterNodeIterNext(m *base.Module, l0 int32) int32
//go:linkname F_clusterPropagatePublish github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterPropagatePublish
func F_clusterPropagatePublish(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clusterAllReplicasThinkPrimaryIsFail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterAllReplicasThinkPrimaryIsFail
func F_clusterAllReplicasThinkPrimaryIsFail(m *base.Module) int32
//go:linkname F_verifyClusterConfigWithData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_verifyClusterConfigWithData
func F_verifyClusterConfigWithData(m *base.Module) int32
//go:linkname F_getNodeReplicationOffset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getNodeReplicationOffset
func F_getNodeReplicationOffset(m *base.Module, l0 int32) int64
//go:linkname F_clusterBeforeSleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterBeforeSleep
func F_clusterBeforeSleep(m *base.Module)
//go:linkname F_addReplyClusterLinkDescription github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyClusterLinkDescription
func F_addReplyClusterLinkDescription(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getSlotOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getSlotOrReply
func F_getSlotOrReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_clusterCommandShards github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandShards
func F_clusterCommandShards(m *base.Module, l0 int32)
//go:linkname F_clusterCommandSpecial github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterCommandSpecial
func F_clusterCommandSpecial(m *base.Module, l0 int32) int32
//go:linkname F_setSlotImportingStateInDb github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setSlotImportingStateInDb
func F_setSlotImportingStateInDb(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_parseSlotRangesOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_parseSlotRangesOrReply
func F_parseSlotRangesOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fireModuleSlotMigrationEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fireModuleSlotMigrationEvent
func F_fireModuleSlotMigrationEvent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_createSlotImportJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSlotImportJob
func F_createSlotImportJob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_finishSlotMigrationJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_finishSlotMigrationJob
func F_finishSlotMigrationJob(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_sendSyncSlotsMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendSyncSlotsMessage
func F_sendSyncSlotsMessage(m *base.Module, l0 int32, l1 int32)
//go:linkname F_updateSlotMigrationJobState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateSlotMigrationJobState
func F_updateSlotMigrationJobState(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterCommandSyncSlotsPaused github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterCommandSyncSlotsPaused
func F_clusterCommandSyncSlotsPaused(m *base.Module, l0 int32)
//go:linkname F_clusterCommandSyncSlotsFailoverGranted github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandSyncSlotsFailoverGranted
func F_clusterCommandSyncSlotsFailoverGranted(m *base.Module, l0 int32)
//go:linkname F_clusterCommandSyncSlotsFinish github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandSyncSlotsFinish
func F_clusterCommandSyncSlotsFinish(m *base.Module, l0 int32)
//go:linkname F_clusterUpdateSlotImportsOnOwnershipChange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateSlotImportsOnOwnershipChange
func F_clusterUpdateSlotImportsOnOwnershipChange(m *base.Module)
//go:linkname F_proceedWithSlotMigration github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_proceedWithSlotMigration
func F_proceedWithSlotMigration(m *base.Module, l0 int32)
//go:linkname F_clusterCommandSyncSlotsRequestPause github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandSyncSlotsRequestPause
func F_clusterCommandSyncSlotsRequestPause(m *base.Module, l0 int32)
//go:linkname F_backgroundSlotMigrationDoneHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_backgroundSlotMigrationDoneHandler
func F_backgroundSlotMigrationDoneHandler(m *base.Module, l0 int32, l1 int32)
//go:linkname F_clusterSlotMigrationShouldInstallWriteHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSlotMigrationShouldInstallWriteHandler
func F_clusterSlotMigrationShouldInstallWriteHandler(m *base.Module, l0 int32) int32
//go:linkname F_clusterUpdateSlotExportsOnOwnershipChange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterUpdateSlotExportsOnOwnershipChange
func F_clusterUpdateSlotExportsOnOwnershipChange(m *base.Module)
//go:linkname F_clusterGetTotalSlotExportBufferMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterGetTotalSlotExportBufferMemory
func F_clusterGetTotalSlotExportBufferMemory(m *base.Module) int32
//go:linkname F_isImportSlotMigrationJob github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isImportSlotMigrationJob
func F_isImportSlotMigrationJob(m *base.Module, l0 int32) int32
//go:linkname F_clusterCommandSyncSlotsAck github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterCommandSyncSlotsAck
func F_clusterCommandSyncSlotsAck(m *base.Module, l0 int32)
//go:linkname F_clusterSlotStatsAddNetworkBytesOutForUserClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatsAddNetworkBytesOutForUserClient
func F_clusterSlotStatsAddNetworkBytesOutForUserClient(m *base.Module, l0 int32)
//go:linkname F_clusterSlotStatsDecrNetworkBytesOutForReplication github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_clusterSlotStatsDecrNetworkBytesOutForReplication
func F_clusterSlotStatsDecrNetworkBytesOutForReplication(m *base.Module, l0 int64)
//go:linkname F_clusterSlotStatResetAll github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clusterSlotStatResetAll
func F_clusterSlotStatResetAll(m *base.Module)
//go:linkname F_commandlogPushCurrentCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_commandlogPushCurrentCommand
func F_commandlogPushCurrentCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_loadServerConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_loadServerConfig
func F_loadServerConfig(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_performInterfaceSet github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_performInterfaceSet
func F_performInterfaceSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_allowProtectedAction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_allowProtectedAction
func F_allowProtectedAction(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rewriteConfigRewriteLine github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rewriteConfigRewriteLine
func F_rewriteConfigRewriteLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rewriteConfigFormatMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteConfigFormatMemory
func F_rewriteConfigFormatMemory(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_configEnumGetName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_configEnumGetName
func F_configEnumGetName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getConfigDebugInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getConfigDebugInfo
func F_getConfigDebugInfo(m *base.Module) int32
//go:linkname F_rewriteConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteConfig
func F_rewriteConfig(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_setNumericType github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setNumericType
func F_setNumericType(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_addModuleBoolConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addModuleBoolConfig
func F_addModuleBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_addModuleStringConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addModuleStringConfig
func F_addModuleStringConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_connTypeRegister github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connTypeRegister
func F_connTypeRegister(m *base.Module, l0 int32) int32
//go:linkname F_connTypeInitialize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_connTypeInitialize
func F_connTypeInitialize(m *base.Module) int32
//go:linkname F_connectionByType github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connectionByType
func F_connectionByType(m *base.Module, l0 int32) int32
//go:linkname F_connectionTypeTcp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connectionTypeTcp
func F_connectionTypeTcp(m *base.Module) int32
//go:linkname F_connTypeHasPendingData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_connTypeHasPendingData
func F_connTypeHasPendingData(m *base.Module) int32
//go:linkname F_connTypeProcessPendingData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_connTypeProcessPendingData
func F_connTypeProcessPendingData(m *base.Module) int32
//go:linkname F_crc64 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_crc64
func F_crc64(m *base.Module, l0 int64, l1 int32, l2 int64) int64
//go:linkname F_crcspeed64little_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_crcspeed64little_init
func F_crcspeed64little_init(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lookupKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKey
func F_lookupKey(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getKeySlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeySlot
func F_getKeySlot(m *base.Module, l0 int32) int32
//go:linkname F_lookupKeyReadWithFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKeyReadWithFlags
func F_lookupKeyReadWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lookupKeyRead github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lookupKeyRead
func F_lookupKeyRead(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupKeyWriteOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupKeyWriteOrReply
func F_lookupKeyWriteOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dbUpdateObjectWithVolatileItemsTracking github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbUpdateObjectWithVolatileItemsTracking
func F_dbUpdateObjectWithVolatileItemsTracking(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dbAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbAdd
func F_dbAdd(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dbSetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbSetValue
func F_dbSetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_removeExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_removeExpire
func F_removeExpire(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_signalModifiedKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_signalModifiedKey
func F_signalModifiedKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dbGenericDeleteWithDictIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbGenericDeleteWithDictIndex
func F_dbGenericDeleteWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_dbSyncDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbSyncDelete
func F_dbSyncDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dbAsyncDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dbAsyncDelete
func F_dbAsyncDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dbUnshareStringValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dbUnshareStringValue
func F_dbUnshareStringValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emptyData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_emptyData
func F_emptyData(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_selectDb github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_selectDb
func F_selectDb(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getFlushCommandFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getFlushCommandFlags
func F_getFlushCommandFlags(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_parseScanCursorOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_parseScanCursorOrReply
func F_parseScanCursorOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_parseScanOptionsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_parseScanOptionsOrReply
func F_parseScanOptionsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_scanGenericCommandWithOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scanGenericCommandWithOptions
func F_scanGenericCommandWithOptions(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)
//go:linkname F_setExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setExpire
func F_setExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_propagateDeletion github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_propagateDeletion
func F_propagateDeletion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deleteExpiredKeyFromOverwriteAndPropagate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_deleteExpiredKeyFromOverwriteAndPropagate
func F_deleteExpiredKeyFromOverwriteAndPropagate(m *base.Module, l0 int32, l1 int32)
//go:linkname F_propagateFieldsDeletion github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_propagateFieldsDeletion
func F_propagateFieldsDeletion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_getKeysPrepareResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysPrepareResult
func F_getKeysPrepareResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getKeysUsingLegacyRangeSpec github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysUsingLegacyRangeSpec
func F_getKeysUsingLegacyRangeSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getKeysFreeResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getKeysFreeResult
func F_getKeysFreeResult(m *base.Module, l0 int32)
//go:linkname F_genericGetKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericGetKeys
func F_genericGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_xorDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_xorDigest
func F_xorDigest(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_mixStringObjectDigest github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mixStringObjectDigest
func F_mixStringObjectDigest(m *base.Module, l0 int32, l1 int32)
//go:linkname F_logServerInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_logServerInfo
func F_logServerInfo(m *base.Module)
//go:linkname F_logCurrentClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_logCurrentClient
func F_logCurrentClient(m *base.Module, l0 int32, l1 int32)
//go:linkname F_bugReportEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_bugReportEnd
func F_bugReportEnd(m *base.Module, l0 int32, l1 int32)
//go:linkname F_printCrashReport github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_printCrashReport
func F_printCrashReport(m *base.Module)
//go:linkname F__serverAssertPrintClientInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__serverAssertPrintClientInfo
func F__serverAssertPrintClientInfo(m *base.Module, l0 int32)
//go:linkname F_debugPauseProcess github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_debugPauseProcess
func F_debugPauseProcess(m *base.Module)
//go:linkname F_dictCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictCreate
func F_dictCreate(m *base.Module, l0 int32) int32
//go:linkname F_dictResizeWithOptionalCheck github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictResizeWithOptionalCheck
func F_dictResizeWithOptionalCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_dictRehash github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictRehash
func F_dictRehash(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictExpandIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictExpandIfNeeded
func F_dictExpandIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_dictGetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictGetVal
func F_dictGetVal(m *base.Module, l0 int32) int32
//go:linkname F_dictDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictDelete
func F_dictDelete(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictRelease
func F_dictRelease(m *base.Module, l0 int32)
//go:linkname F_dictFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictFind
func F_dictFind(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_dictSetKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictSetKey
func F_dictSetKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_dictInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictInitIterator
func F_dictInitIterator(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dictGetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictGetIterator
func F_dictGetIterator(m *base.Module, l0 int32) int32
//go:linkname F_dictReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dictReleaseIterator
func F_dictReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_dictGetRandomKey github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictGetRandomKey
func F_dictGetRandomKey(m *base.Module, l0 int32) int32
//go:linkname F_dictEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dictEmpty
func F_dictEmpty(m *base.Module, l0 int32, l1 int32)
//go:linkname F_entryConstruct github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_entryConstruct
func F_entryConstruct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_entryFreeValuePtr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_entryFreeValuePtr
func F_entryFreeValuePtr(m *base.Module, l0 int32)
//go:linkname F_entryCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_entryCreate
func F_entryCreate(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_evalExtractShebangFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalExtractShebangFlags
func F_evalExtractShebangFlags(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_evalShaCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalShaCommand
func F_evalShaCommand(m *base.Module, l0 int32)
//go:linkname F_evalRegisterNewScript github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evalRegisterNewScript
func F_evalRegisterNewScript(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_performEvictions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_performEvictions
func F_performEvictions(m *base.Module) int32
//go:linkname F_activeExpireCycleJob github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_activeExpireCycleJob
func F_activeExpireCycleJob(m *base.Module, l0 int32, l1 int32, l2 int64) int64
//go:linkname F_convertExpireArgumentToUnixTime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_convertExpireArgumentToUnixTime
func F_convertExpireArgumentToUnixTime(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32
//go:linkname F_expireGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_expireGenericCommand
func F_expireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_functionsLibCtxClear github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_functionsLibCtxClear
func F_functionsLibCtxClear(m *base.Module, l0 int32, l1 int32)
//go:linkname F_libraryUnlink github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_libraryUnlink
func F_libraryUnlink(m *base.Module, l0 int32, l1 int32)
//go:linkname F_fcallGetCommandFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fcallGetCommandFlags
func F_fcallGetCommandFlags(m *base.Module, l0 int32, l1 int64) int64
//go:linkname F_libraryLink github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_libraryLink
func F_libraryLink(m *base.Module, l0 int32, l1 int32)
//go:linkname F_functionFreeLibMetaData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_functionFreeLibMetaData
func F_functionFreeLibMetaData(m *base.Module, l0 int32)
//go:linkname F_addReplyDoubleDistance github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyDoubleDistance
func F_addReplyDoubleDistance(m *base.Module, l0 int32, l1 float64)
//go:linkname F_georadiusGeneric github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_georadiusGeneric
func F_georadiusGeneric(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_geohashBoundingBox github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_geohashBoundingBox
func F_geohashBoundingBox(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCreate
func F_hashtableCreate(m *base.Module, l0 int32) int32
//go:linkname F_hashtableEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableEmpty
func F_hashtableEmpty(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashtableRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableRelease
func F_hashtableRelease(m *base.Module, l0 int32)
//go:linkname F_hashtableSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableSize
func F_hashtableSize(m *base.Module, l0 int32) int32
//go:linkname F_resize_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_resize_1
func F_resize_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableRehashingInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableRehashingInfo
func F_hashtableRehashingInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rehashStep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rehashStep
func F_rehashStep(m *base.Module, l0 int32)
//go:linkname F_hashtableExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableExpand
func F_hashtableExpand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableTryExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableTryExpand
func F_hashtableTryExpand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableRightsizeIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableRightsizeIfNeeded
func F_hashtableRightsizeIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_hashtableFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableFind
func F_hashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_findBucket_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_findBucket_1
func F_findBucket_1(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_hashtableAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableAdd
func F_hashtableAdd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableFindPositionForInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableFindPositionForInsert
func F_hashtableFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hashtableInsertAtPosition github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableInsertAtPosition
func F_hashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hashtablePop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtablePop
func F_hashtablePop(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hashtableIncrementalFindStep github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableIncrementalFindStep
func F_hashtableIncrementalFindStep(m *base.Module, l0 int32) int32
//go:linkname F_hashtableIncrementalFindGetResult github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableIncrementalFindGetResult
func F_hashtableIncrementalFindGetResult(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableScanDefrag github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableScanDefrag
func F_hashtableScanDefrag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_hashtableRetargetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableRetargetIterator
func F_hashtableRetargetIterator(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashtableCleanupIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCleanupIterator
func F_hashtableCleanupIterator(m *base.Module, l0 int32)
//go:linkname F_hashtableCreateIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableCreateIterator
func F_hashtableCreateIterator(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashtableReleaseIterator
func F_hashtableReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_hashtableNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableNext
func F_hashtableNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashtableRandomEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashtableRandomEntry
func F_hashtableRandomEntry(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hllCount github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hllCount
func F_hllCount(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_hllMerge github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hllMerge
func F_hllMerge(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_isHLLObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isHLLObjectOrReply
func F_isHLLObjectOrReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hllAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hllAdd
func F_hllAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_intsetAdd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_intsetAdd
func F_intsetAdd(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_intsetSearch github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_intsetSearch
func F_intsetSearch(m *base.Module, l0 int32, l1 int64, l2 int32) int32
//go:linkname F_intsetDup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_intsetDup
func F_intsetDup(m *base.Module, l0 int32) int32
//go:linkname F_drainIOThreadsQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_drainIOThreadsQueue
func F_drainIOThreadsQueue(m *base.Module)
//go:linkname F_flushPendingIOResponses github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_flushPendingIOResponses
func F_flushPendingIOResponses(m *base.Module, l0 int32)
//go:linkname F_initIOThreads github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_initIOThreads
func F_initIOThreads(m *base.Module, l0 int32)
//go:linkname F_trySendWriteToIOThreads github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trySendWriteToIOThreads
func F_trySendWriteToIOThreads(m *base.Module, l0 int32) int32
//go:linkname F_tryOffloadFreeArgvToIOThreads github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_tryOffloadFreeArgvToIOThreads
func F_tryOffloadFreeArgvToIOThreads(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sendToMainThread github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendToMainThread
func F_sendToMainThread(m *base.Module, l0 int32, l1 int32)
//go:linkname F_processIOThreadsResponses github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_processIOThreadsResponses
func F_processIOThreadsResponses(m *base.Module) int32
//go:linkname F_kvstoreCreate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreCreate
func F_kvstoreCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreGetNextNonEmptyHashtableIndex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreGetNextNonEmptyHashtableIndex
func F_kvstoreGetNextNonEmptyHashtableIndex(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreHashtableSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableSize
func F_kvstoreHashtableSize(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreGetHashtableIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreGetHashtableIterator
func F_kvstoreGetHashtableIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreHashtableIteratorNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableIteratorNext
func F_kvstoreHashtableIteratorNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_kvstoreHashtableFindRef github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableFindRef
func F_kvstoreHashtableFindRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_kvstoreHashtableInsertAtPosition github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_kvstoreHashtableInsertAtPosition
func F_kvstoreHashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_kvstoreHashtableDelete github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kvstoreHashtableDelete
func F_kvstoreHashtableDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_analyzeLatencyForEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_analyzeLatencyForEvent
func F_analyzeLatencyForEvent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_freeObjAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeObjAsync
func F_freeObjAsync(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_freeTrackingRadixTreeAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeTrackingRadixTreeAsync
func F_freeTrackingRadixTreeAsync(m *base.Module, l0 int32)
//go:linkname F_freeErrorsRadixTreeAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeErrorsRadixTreeAsync
func F_freeErrorsRadixTreeAsync(m *base.Module, l0 int32)
//go:linkname F_freeFunctionsAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeFunctionsAsync
func F_freeFunctionsAsync(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lpNew github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpNew
func F_lpNew(m *base.Module, l0 int32) int32
//go:linkname F_lpFree github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpFree
func F_lpFree(m *base.Module, l0 int32)
//go:linkname F_lpShrinkToFit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpShrinkToFit
func F_lpShrinkToFit(m *base.Module, l0 int32) int32
//go:linkname F_lpPrev github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpPrev
func F_lpPrev(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lpFirst github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpFirst
func F_lpFirst(m *base.Module, l0 int32) int32
//go:linkname F_lpGet github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpGet
func F_lpGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpGetWithSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpGetWithSize
func F_lpGetWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpFind github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpFind
func F_lpFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lpInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpInsert
func F_lpInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_lpInsertString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpInsertString
func F_lpInsertString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_lpInsertInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpInsertInteger
func F_lpInsertInteger(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lpPrepend github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpPrepend
func F_lpPrepend(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpPrependInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpPrependInteger
func F_lpPrependInteger(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_lpAppendInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpAppendInteger
func F_lpAppendInteger(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_lpReplaceInteger github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpReplaceInteger
func F_lpReplaceInteger(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_lpDeleteRange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpDeleteRange
func F_lpDeleteRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpRandomPairsUnique github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpRandomPairsUnique
func F_lpRandomPairsUnique(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_lpNextRandom github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpNextRandom
func F_lpNextRandom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_lpRepr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lpRepr
func F_lpRepr(m *base.Module, l0 int32)
//go:linkname F_lrulfu_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lrulfu_init
func F_lrulfu_init(m *base.Module) int32
//go:linkname F_processClientsCommandsBatch github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_processClientsCommandsBatch
func F_processClientsCommandsBatch(m *base.Module)
//go:linkname F_addCommandToBatch github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addCommandToBatch
func F_addCommandToBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_memtest_progress_start github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memtest_progress_start
func F_memtest_progress_start(m *base.Module, l0 int32, l1 int32)
//go:linkname F_memtest_addressing github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_memtest_addressing
func F_memtest_addressing(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memtest_fill_random github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_memtest_fill_random
func F_memtest_fill_random(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_memtest_fill_value github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memtest_fill_value
func F_memtest_fill_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_memtest_compare github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memtest_compare
func F_memtest_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memtest_alloc_and_test github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_memtest_alloc_and_test
func F_memtest_alloc_and_test(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleEnqueueLoadModule github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleEnqueueLoadModule
func F_moduleEnqueueLoadModule(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_moduleReleaseTempClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleReleaseTempClient
func F_moduleReleaseTempClient(m *base.Module, l0 int32)
//go:linkname F_moduleFreeKeyIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleFreeKeyIterator
func F_moduleFreeKeyIterator(m *base.Module, l0 int32)
//go:linkname F_autoMemoryCollect github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_autoMemoryCollect
func F_autoMemoryCollect(m *base.Module, l0 int32)
//go:linkname F_moduleCreateContext github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleCreateContext
func F_moduleCreateContext(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_moduleGetCommandKeysViaAPI github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleGetCommandKeysViaAPI
func F_moduleGetCommandKeysViaAPI(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleCreateCommandProxy github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleCreateCommandProxy
func F_moduleCreateCommandProxy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_populateArgsStructure github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_populateArgsStructure
func F_populateArgsStructure(m *base.Module, l0 int32) int32
//go:linkname F_categoryFlagsFromString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_categoryFlagsFromString
func F_categoryFlagsFromString(m *base.Module, l0 int32) int64
//go:linkname F_moduleValidateCommandArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleValidateCommandArgs
func F_moduleValidateCommandArgs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleCopyCommandArgs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleCopyCommandArgs
func F_moduleCopyCommandArgs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleGetHandleByName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleGetHandleByName
func F_moduleGetHandleByName(m *base.Module, l0 int32) int32
//go:linkname F_moduleReplyErrorFormatInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleReplyErrorFormatInternal
func F_moduleReplyErrorFormatInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleCreateArgvFromUserFormat github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleCreateArgvFromUserFormat
func F_moduleCreateArgvFromUserFormat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_moduleListIteratorSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleListIteratorSeek
func F_moduleListIteratorSeek(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zsetInitScoreRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetInitScoreRange
func F_zsetInitScoreRange(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_moduleTypeLookupModuleByNameInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleTypeLookupModuleByNameInternal
func F_moduleTypeLookupModuleByNameInternal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moduleTypeLookupModuleByID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleTypeLookupModuleByID
func F_moduleTypeLookupModuleByID(m *base.Module, l0 int64) int32
//go:linkname F_moduleAllDatatypesHandleErrors github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleAllDatatypesHandleErrors
func F_moduleAllDatatypesHandleErrors(m *base.Module) int32
//go:linkname F_moduleAllModulesHandleReplAsyncLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleAllModulesHandleReplAsyncLoad
func F_moduleAllModulesHandleReplAsyncLoad(m *base.Module) int32
//go:linkname F_moduleVerifyAllAllowAtomicSlotMigrationOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleVerifyAllAllowAtomicSlotMigrationOrReply
func F_moduleVerifyAllAllowAtomicSlotMigrationOrReply(m *base.Module, l0 int32) int32
//go:linkname F_moduleLoadString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleLoadString
func F_moduleLoadString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_unblockClientFromModule github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unblockClientFromModule
func F_unblockClientFromModule(m *base.Module, l0 int32)
//go:linkname F_moduleBlockClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleBlockClient
func F_moduleBlockClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32) int32
//go:linkname F_moduleFireServerEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleFireServerEvent
func F_moduleFireServerEvent(m *base.Module, l0 int64, l1 int32, l2 int32)
//go:linkname F_moduleBlockedClientTimedOut github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleBlockedClientTimedOut
func F_moduleBlockedClientTimedOut(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleHandleBlockedClients github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleHandleBlockedClients
func F_moduleHandleBlockedClients(m *base.Module)
//go:linkname F_moduleNotifyKeyspaceEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleNotifyKeyspaceEvent
func F_moduleNotifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_moduleCallClusterReceivers github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleCallClusterReceivers
func F_moduleCallClusterReceivers(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32)
//go:linkname F_moduleNotifyUserChanged github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleNotifyUserChanged
func F_moduleNotifyUserChanged(m *base.Module, l0 int32)
//go:linkname F_modulesCollectInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_modulesCollectInfo
func F_modulesCollectInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleFireCommandResultEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleFireCommandResultEvent
func F_moduleFireCommandResultEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64)
//go:linkname F_moduleFireCommandRejectedEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleFireCommandRejectedEvent
func F_moduleFireCommandRejectedEvent(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleFireCommandACLRejectedEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleFireCommandACLRejectedEvent
func F_moduleFireCommandACLRejectedEvent(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_TerminateModuleForkChild github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_TerminateModuleForkChild
func F_TerminateModuleForkChild(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ModuleForkDoneHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ModuleForkDoneHandler
func F_ModuleForkDoneHandler(m *base.Module, l0 int32, l1 int32)
//go:linkname F_moduleLoad github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_moduleLoad
func F_moduleLoad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleUnregisterCleanup github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleUnregisterCleanup
func F_moduleUnregisterCleanup(m *base.Module, l0 int32)
//go:linkname F_moduleLoadStatic github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleLoadStatic
func F_moduleLoadStatic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_moduleLoadStaticSymbol github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moduleLoadStaticSymbol
func F_moduleLoadStaticSymbol(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setModuleBoolConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setModuleBoolConfig
func F_setModuleBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getModuleBoolConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getModuleBoolConfig
func F_getModuleBoolConfig(m *base.Module, l0 int32) int32
//go:linkname F_getModuleStringConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getModuleStringConfig
func F_getModuleStringConfig(m *base.Module, l0 int32) int32
//go:linkname F_getModuleNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getModuleNumericConfig
func F_getModuleNumericConfig(m *base.Module, l0 int32) int64
//go:linkname F_getModuleUnsignedNumericConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getModuleUnsignedNumericConfig
func F_getModuleUnsignedNumericConfig(m *base.Module, l0 int32) int64
//go:linkname F_queueMultiCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_queueMultiCommand
func F_queueMultiCommand(m *base.Module, l0 int32, l1 int64)
//go:linkname F_getStringObjectSdsUsedMemory github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getStringObjectSdsUsedMemory
func F_getStringObjectSdsUsedMemory(m *base.Module, l0 int32) int32
//go:linkname F_readToQueryBuf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_readToQueryBuf
func F_readToQueryBuf(m *base.Module, l0 int32) int32
//go:linkname F_processInputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processInputBuffer
func F_processInputBuffer(m *base.Module, l0 int32) int32
//go:linkname F_trimCommandQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_trimCommandQueue
func F_trimCommandQueue(m *base.Module, l0 int32)
//go:linkname F_freeClientAsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClientAsync
func F_freeClientAsync(m *base.Module, l0 int32)
//go:linkname F_prepareClientToWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_prepareClientToWrite
func F_prepareClientToWrite(m *base.Module, l0 int32) int32
//go:linkname F_aggregateClientOutputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_aggregateClientOutputBuffer
func F_aggregateClientOutputBuffer(m *base.Module, l0 int32) int32
//go:linkname F_freeClient github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClient
func F_freeClient(m *base.Module, l0 int32) int32
//go:linkname F_isCopyAvoidPreferred github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isCopyAvoidPreferred
func F_isCopyAvoidPreferred(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_upsertPayloadHeader github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_upsertPayloadHeader
func F_upsertPayloadHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32
//go:linkname F_addReplyErrorLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorLength
func F_addReplyErrorLength(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_afterErrorReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_afterErrorReply
func F_afterErrorReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_commitDeferredReplyBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commitDeferredReplyBuffer
func F_commitDeferredReplyBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorObject
func F_addReplyErrorObject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyError
func F_addReplyError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorSds
func F_addReplyErrorSds(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorSdsSafe github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyErrorSdsSafe
func F_addReplyErrorSdsSafe(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyErrorArity github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyErrorArity
func F_addReplyErrorArity(m *base.Module, l0 int32)
//go:linkname F_addReplyErrorExpireTime github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyErrorExpireTime
func F_addReplyErrorExpireTime(m *base.Module, l0 int32)
//go:linkname F_addReplyStatusLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyStatusLength
func F_addReplyStatusLength(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyStatus github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyStatus
func F_addReplyStatus(m *base.Module, l0 int32, l1 int32)
//go:linkname F_setDeferredAggregateLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setDeferredAggregateLen
func F_setDeferredAggregateLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_setDeferredSetLen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setDeferredSetLen
func F_setDeferredSetLen(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_prepareClientForFutureWrites github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_prepareClientForFutureWrites
func F_prepareClientForFutureWrites(m *base.Module, l0 int32) int32
//go:linkname F_addReplyDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyDouble
func F_addReplyDouble(m *base.Module, l0 int32, l1 float64)
//go:linkname F_addReplyBulk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyBulk
func F_addReplyBulk(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyLongLong
func F_addReplyLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F__addReplyLongLongWithPrefix github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__addReplyLongLongWithPrefix
func F__addReplyLongLongWithPrefix(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_addReplyArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyArrayLen
func F_addReplyArrayLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addWritePreparedReplyArrayLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addWritePreparedReplyArrayLen
func F_addWritePreparedReplyArrayLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplySetLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplySetLen
func F_addReplySetLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyNull github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyNull
func F_addReplyNull(m *base.Module, l0 int32)
//go:linkname F_addReplyBulkSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyBulkSds
func F_addReplyBulkSds(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyBulkCString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyBulkCString
func F_addReplyBulkCString(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addWritePreparedReplyBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addWritePreparedReplyBulkLongLong
func F_addWritePreparedReplyBulkLongLong(m *base.Module, l0 int32, l1 int64)
//go:linkname F_addReplyVerbatim github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyVerbatim
func F_addReplyVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_deferredAfterErrorReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_deferredAfterErrorReply
func F_deferredAfterErrorReply(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getClientMemoryUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getClientMemoryUsage
func F_getClientMemoryUsage(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getClientSockname github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getClientSockname
func F_getClientSockname(m *base.Module, l0 int32) int32
//go:linkname F_acceptCommonHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_acceptCommonHandler
func F_acceptCommonHandler(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_clearClientConnectionState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clearClientConnectionState
func F_clearClientConnectionState(m *base.Module, l0 int32)
//go:linkname F_releaseBufReferences github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_releaseBufReferences
func F_releaseBufReferences(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_discardCommandQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_discardCommandQueue
func F_discardCommandQueue(m *base.Module, l0 int32)
//go:linkname F_freeClientOrCloseLater github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClientOrCloseLater
func F_freeClientOrCloseLater(m *base.Module, l0 int32, l1 int32)
//go:linkname F_trimClientQueryBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trimClientQueryBuffer
func F_trimClientQueryBuffer(m *base.Module, l0 int32)
//go:linkname F_freeClientsInAsyncFreeQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeClientsInAsyncFreeQueue
func F_freeClientsInAsyncFreeQueue(m *base.Module) int32
//go:linkname F__writeToClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__writeToClient
func F__writeToClient(m *base.Module, l0 int32) int32
//go:linkname F_postWriteToClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_postWriteToClient
func F_postWriteToClient(m *base.Module, l0 int32) int32
//go:linkname F_writeToReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_writeToReplica
func F_writeToReplica(m *base.Module, l0 int32)
//go:linkname F_setProtocolError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setProtocolError
func F_setProtocolError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_resetClient github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_resetClient
func F_resetClient(m *base.Module, l0 int32)
//go:linkname F_handleClientsWithPendingWrites github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_handleClientsWithPendingWrites
func F_handleClientsWithPendingWrites(m *base.Module) int32
//go:linkname F_freeSharedQueryBuf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_freeSharedQueryBuf
func F_freeSharedQueryBuf(m *base.Module)
//go:linkname F_commandProcessed github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_commandProcessed
func F_commandProcessed(m *base.Module, l0 int32)
//go:linkname F_parseInputBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_parseInputBuffer
func F_parseInputBuffer(m *base.Module, l0 int32)
//go:linkname F_addKeysToIncrFindBatch github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addKeysToIncrFindBatch
func F_addKeysToIncrFindBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_isClientConnIpV6 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isClientConnIpV6
func F_isClientConnIpV6(m *base.Module, l0 int32) int32
//go:linkname F_getClientOutputBufferMemoryUsage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getClientOutputBufferMemoryUsage
func F_getClientOutputBufferMemoryUsage(m *base.Module, l0 int32) int32
//go:linkname F_clientSetName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_clientSetName
func F_clientSetName(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_unpauseActions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unpauseActions
func F_unpauseActions(m *base.Module, l0 int32)
//go:linkname F_securityWarningCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_securityWarningCommand
func F_securityWarningCommand(m *base.Module, l0 int32)
//go:linkname F_rewriteClientCommandVector github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteClientCommandVector
func F_rewriteClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_backupAndUpdateClientArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_backupAndUpdateClientArgv
func F_backupAndUpdateClientArgv(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_rewriteClientCommandArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rewriteClientCommandArgument
func F_rewriteClientCommandArgument(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pauseActions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pauseActions
func F_pauseActions(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_processEventsWhileBlocked github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processEventsWhileBlocked
func F_processEventsWhileBlocked(m *base.Module)
//go:linkname F_evictClients github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_evictClients
func F_evictClients(m *base.Module)
//go:linkname F_keyspaceEventsFlagsToString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_keyspaceEventsFlagsToString
func F_keyspaceEventsFlagsToString(m *base.Module, l0 int32) int32
//go:linkname F_notifyKeyspaceEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_notifyKeyspaceEvent
func F_notifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_createStringObject_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createStringObject_1
func F_createStringObject_1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_objectGetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectGetVal
func F_objectGetVal(m *base.Module, l0 int32) int32
//go:linkname F_objectSetKeyAndExpire github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_objectSetKeyAndExpire
func F_objectSetKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_decrRefCount github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_decrRefCount
func F_decrRefCount(m *base.Module, l0 int32)
//go:linkname F_objectSetVal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_objectSetVal
func F_objectSetVal(m *base.Module, l0 int32, l1 int32)
//go:linkname F_createStringObjectFromLongLongForValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createStringObjectFromLongLongForValue
func F_createStringObjectFromLongLongForValue(m *base.Module, l0 int64) int32
//go:linkname F_createStringObjectFromLongDouble github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createStringObjectFromLongDouble
func F_createStringObjectFromLongDouble(m *base.Module, l0 int64, l1 int64, l2 int32) int32
//go:linkname F_dupStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dupStringObject
func F_dupStringObject(m *base.Module, l0 int32) int32
//go:linkname F_createSetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createSetObject
func F_createSetObject(m *base.Module) int32
//go:linkname F_createIntsetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createIntsetObject
func F_createIntsetObject(m *base.Module) int32
//go:linkname F_createZsetObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createZsetObject
func F_createZsetObject(m *base.Module) int32
//go:linkname F_createStreamObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createStreamObject
func F_createStreamObject(m *base.Module) int32
//go:linkname F_incrRefCount github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_incrRefCount
func F_incrRefCount(m *base.Module, l0 int32)
//go:linkname F_trimStringObjectIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trimStringObjectIfNeeded
func F_trimStringObjectIfNeeded(m *base.Module, l0 int32, l1 int32)
//go:linkname F_tryObjectEncodingEx github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_tryObjectEncodingEx
func F_tryObjectEncodingEx(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getDecodedObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getDecodedObject
func F_getDecodedObject(m *base.Module, l0 int32) int32
//go:linkname F_compareStringObjects github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_compareStringObjects
func F_compareStringObjects(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_collateStringObjects github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_collateStringObjects
func F_collateStringObjects(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_equalStringObjects github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_equalStringObjects
func F_equalStringObjects(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getLongLongFromObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getLongLongFromObject
func F_getLongLongFromObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_getLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getLongFromObjectOrReply
func F_getLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getRangeLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getRangeLongFromObjectOrReply
func F_getRangeLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_getPositiveLongFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getPositiveLongFromObjectOrReply
func F_getPositiveLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_getIntFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getIntFromObjectOrReply
func F_getIntFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_freeMemoryOverheadData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeMemoryOverheadData
func F_freeMemoryOverheadData(m *base.Module, l0 int32)
//go:linkname F_getMemoryOverheadData github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getMemoryOverheadData
func F_getMemoryOverheadData(m *base.Module) int32
//go:linkname F_pqsort github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pqsort
func F_pqsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_addReplyPubsubMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyPubsubMessage
func F_addReplyPubsubMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_addReplyPubsubUnsubscribed github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyPubsubUnsubscribed
func F_addReplyPubsubUnsubscribed(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_addReplyPubsubPatUnsubscribed github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyPubsubPatUnsubscribed
func F_addReplyPubsubPatUnsubscribed(m *base.Module, l0 int32, l1 int32)
//go:linkname F_initClientPubSubData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_initClientPubSubData
func F_initClientPubSubData(m *base.Module, l0 int32)
//go:linkname F_pubsubUnsubscribeAllPatterns github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pubsubUnsubscribeAllPatterns
func F_pubsubUnsubscribeAllPatterns(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pubsubShardUnsubscribeAllChannelsInSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pubsubShardUnsubscribeAllChannelsInSlot
func F_pubsubShardUnsubscribeAllChannelsInSlot(m *base.Module, l0 int32)
//go:linkname F_pubsubSubscribePattern github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pubsubSubscribePattern
func F_pubsubSubscribePattern(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pubsubPublishMessageInternal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pubsubPublishMessageInternal
func F_pubsubPublishMessageInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_channelList github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_channelList
func F_channelList(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_quicklistRelease github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistRelease
func F_quicklistRelease(m *base.Module, l0 int32)
//go:linkname F_quicklistPushHead github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistPushHead
func F_quicklistPushHead(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quicklistPushTail github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistPushTail
func F_quicklistPushTail(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quicklistReplaceEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistReplaceEntry
func F_quicklistReplaceEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F___quicklistCompressNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___quicklistCompressNode
func F___quicklistCompressNode(m *base.Module, l0 int32) int32
//go:linkname F___quicklistCompress github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___quicklistCompress
func F___quicklistCompress(m *base.Module, l0 int32, l1 int32)
//go:linkname F__quicklistInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__quicklistInsert
func F__quicklistInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F__quicklistListpackMerge github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__quicklistListpackMerge
func F__quicklistListpackMerge(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quicklistNext github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistNext
func F_quicklistNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_quicklistReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistReleaseIterator
func F_quicklistReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_quicklistDelRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistDelRange
func F_quicklistDelRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_quicklistGetIteratorAtIdx github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_quicklistGetIteratorAtIdx
func F_quicklistGetIteratorAtIdx(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_quicklistCompare github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_quicklistCompare
func F_quicklistCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_raxNew github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxNew
func F_raxNew(m *base.Module) int32
//go:linkname F_raxAddChild github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxAddChild
func F_raxAddChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxLowWalk github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxLowWalk
func F_raxLowWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_raxInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxInsert
func F_raxInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_raxTryInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxTryInsert
func F_raxTryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_raxFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxFind
func F_raxFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_raxFreeWithCallback github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_raxFreeWithCallback
func F_raxFreeWithCallback(m *base.Module, l0 int32, l1 int32)
//go:linkname F_raxFree github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxFree
func F_raxFree(m *base.Module, l0 int32)
//go:linkname F_raxIteratorNextStep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxIteratorNextStep
func F_raxIteratorNextStep(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_raxIteratorPrevStep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxIteratorPrevStep
func F_raxIteratorPrevStep(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_raxSeek github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_raxSeek
func F_raxSeek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_rdbReportError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbReportError
func F_rdbReportError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_rdbRegisterAuxField github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbRegisterAuxField
func F_rdbRegisterAuxField(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbWriteRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbWriteRaw
func F_rdbWriteRaw(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveType github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveType
func F_rdbSaveType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbLoadMillisecondTime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadMillisecondTime
func F_rdbLoadMillisecondTime(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_rdbSaveLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveLen
func F_rdbSaveLen(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_rdbLoadLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadLen
func F_rdbLoadLen(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_rdbSaveRawString github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveRawString
func F_rdbSaveRawString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbSaveStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveStringObject
func F_rdbSaveStringObject(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbGenericLoadStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbGenericLoadStringObject
func F_rdbGenericLoadStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadStringObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbLoadStringObject
func F_rdbLoadStringObject(m *base.Module, l0 int32) int32
//go:linkname F_rdbLoadDoubleValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadDoubleValue
func F_rdbLoadDoubleValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbSaveBinaryFloatValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveBinaryFloatValue
func F_rdbSaveBinaryFloatValue(m *base.Module, l0 int32, l1 float32) int32
//go:linkname F_rdbGetObjectType github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbGetObjectType
func F_rdbGetObjectType(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbSaveObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveObject
func F_rdbSaveObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbSaveAuxField github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbSaveAuxField
func F_rdbSaveAuxField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_rdbSaveRio github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbSaveRio
func F_rdbSaveRio(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_startSaving github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_startSaving
func F_startSaving(m *base.Module, l0 int32)
//go:linkname F_stopSaving github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_stopSaving
func F_stopSaving(m *base.Module, l0 int32)
//go:linkname F_rdbLoadCheckModuleValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rdbLoadCheckModuleValue
func F_rdbLoadCheckModuleValue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_ziplistPairsConvertAndValidateIntegrity github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ziplistPairsConvertAndValidateIntegrity
func F_ziplistPairsConvertAndValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lpValidateIntegrityAndDups github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lpValidateIntegrityAndDups
func F_lpValidateIntegrityAndDups(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rdbLoadRioWithLoadingCtx github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbLoadRioWithLoadingCtx
func F_rdbLoadRioWithLoadingCtx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_killRDBChild github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_killRDBChild
func F_killRDBChild(m *base.Module)
//go:linkname F_replicationGetReplicaName github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationGetReplicaName
func F_replicationGetReplicaName(m *base.Module, l0 int32) int32
//go:linkname F_freeReplicationBacklog github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeReplicationBacklog
func F_freeReplicationBacklog(m *base.Module)
//go:linkname F_rebaseReplicationBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rebaseReplicationBuffer
func F_rebaseReplicationBuffer(m *base.Module, l0 int64)
//go:linkname F_feedReplicationBufferWithObject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_feedReplicationBufferWithObject
func F_feedReplicationBufferWithObject(m *base.Module, l0 int32)
//go:linkname F_feedReplicationBuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_feedReplicationBuffer
func F_feedReplicationBuffer(m *base.Module, l0 int32, l1 int32)
//go:linkname F_generateSelectCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_generateSelectCommand
func F_generateSelectCommand(m *base.Module, l0 int32) int32
//go:linkname F_replicationFeedMonitors github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationFeedMonitors
func F_replicationFeedMonitors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_initClientReplicationData github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_initClientReplicationData
func F_initClientReplicationData(m *base.Module, l0 int32)
//go:linkname F_replicationHandlePrimaryDisconnection github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationHandlePrimaryDisconnection
func F_replicationHandlePrimaryDisconnection(m *base.Module)
//go:linkname F_connectWithPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_connectWithPrimary
func F_connectWithPrimary(m *base.Module) int32
//go:linkname F_replicaPutOnline github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaPutOnline
func F_replicaPutOnline(m *base.Module, l0 int32) int32
//go:linkname F_closeRepldbfd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_closeRepldbfd
func F_closeRepldbfd(m *base.Module, l0 int32)
//go:linkname F_updateReplicasWaitingBgsave github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_updateReplicasWaitingBgsave
func F_updateReplicasWaitingBgsave(m *base.Module, l0 int32, l1 int32)
//go:linkname F_dualChannelSyncSuccess github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dualChannelSyncSuccess
func F_dualChannelSyncSuccess(m *base.Module)
//go:linkname F_replicationAbortDualChannelSyncTransfer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_replicationAbortDualChannelSyncTransfer
func F_replicationAbortDualChannelSyncTransfer(m *base.Module)
//go:linkname F_sendCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sendCommand
func F_sendCommand(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sendCommandArgv github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendCommandArgv
func F_sendCommandArgv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_dualChannelSyncHandlePsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dualChannelSyncHandlePsync
func F_dualChannelSyncHandlePsync(m *base.Module) int32
//go:linkname F_replicaProcessPsyncReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicaProcessPsyncReply
func F_replicaProcessPsyncReply(m *base.Module, l0 int32) int32
//go:linkname F_dualChannelSetupMainConnForPsync github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dualChannelSetupMainConnForPsync
func F_dualChannelSetupMainConnForPsync(m *base.Module, l0 int32)
//go:linkname F_abortFailover github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_abortFailover
func F_abortFailover(m *base.Module, l0 int32)
//go:linkname F_replicationSetPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationSetPrimary
func F_replicationSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_replicationCachePrimaryUsingMyself github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_replicationCachePrimaryUsingMyself
func F_replicationCachePrimaryUsingMyself(m *base.Module)
//go:linkname F_processClientsWaitingReplicas github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processClientsWaitingReplicas
func F_processClientsWaitingReplicas(m *base.Module)
//go:linkname F_replicationGetReplicaOffset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_replicationGetReplicaOffset
func F_replicationGetReplicaOffset(m *base.Module) int64
//go:linkname F_findReplica github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_findReplica
func F_findReplica(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_replicationStartPendingFork github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_replicationStartPendingFork
func F_replicationStartPendingFork(m *base.Module)
//go:linkname F_rioInitWithFd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rioInitWithFd
func F_rioInitWithFd(m *base.Module, l0 int32, l1 int32)
//go:linkname F_rioWriteBulkCount github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkCount
func F_rioWriteBulkCount(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rioWriteBulkString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rioWriteBulkString
func F_rioWriteBulkString(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_rioWriteBulkLongLong github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioWriteBulkLongLong
func F_rioWriteBulkLongLong(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_rioFdWrite github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rioFdWrite
func F_rioFdWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scriptGetCaller github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptGetCaller
func F_scriptGetCaller(m *base.Module) int32
//go:linkname F_scriptPrepareForRun github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptPrepareForRun
func F_scriptPrepareForRun(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int32
//go:linkname F_scriptResetRun github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptResetRun
func F_scriptResetRun(m *base.Module, l0 int32)
//go:linkname F_scriptCurrFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptCurrFunction
func F_scriptCurrFunction(m *base.Module) int32
//go:linkname F_scriptIsEval github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptIsEval
func F_scriptIsEval(m *base.Module) int32
//go:linkname F_scriptRunDuration github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptRunDuration
func F_scriptRunDuration(m *base.Module) int64
//go:linkname F_scriptAllowsOOM github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptAllowsOOM
func F_scriptAllowsOOM(m *base.Module) int32
//go:linkname F_scriptIsReadOnly github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptIsReadOnly
func F_scriptIsReadOnly(m *base.Module) int32
//go:linkname F_scriptIsWriteDirty github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptIsWriteDirty
func F_scriptIsWriteDirty(m *base.Module) int32
//go:linkname F_scriptSetWriteDirtyFlag github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptSetWriteDirtyFlag
func F_scriptSetWriteDirtyFlag(m *base.Module)
//go:linkname F_scriptAllowsCrossSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptAllowsCrossSlot
func F_scriptAllowsCrossSlot(m *base.Module) int32
//go:linkname F_scriptSetOriginalClientSlot github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptSetOriginalClientSlot
func F_scriptSetOriginalClientSlot(m *base.Module, l0 int32)
//go:linkname F_scriptGetRunningEngineName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptGetRunningEngineName
func F_scriptGetRunningEngineName(m *base.Module) int32
//go:linkname F_scriptClusterSlotStatsInvalidateSlotIfApplicable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptClusterSlotStatsInvalidateSlotIfApplicable
func F_scriptClusterSlotStatsInvalidateSlotIfApplicable(m *base.Module)
//go:linkname F_scriptingEngineCallGetMemoryInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineCallGetMemoryInfo
func F_scriptingEngineCallGetMemoryInfo(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scriptingEngineManagerUnregister github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineManagerUnregister
func F_scriptingEngineManagerUnregister(m *base.Module, l0 int32) int32
//go:linkname F_scriptingEngineCallFreeFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineCallFreeFunction
func F_scriptingEngineCallFreeFunction(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_scriptingEngineCallFunction github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineCallFunction
func F_scriptingEngineCallFunction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_scriptingEngineCallResetEnvFunc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineCallResetEnvFunc
func F_scriptingEngineCallResetEnvFunc(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_scriptingEngineCallDebuggerDisable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineCallDebuggerDisable
func F_scriptingEngineCallDebuggerDisable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scriptingEngineCallDebuggerEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineCallDebuggerEnd
func F_scriptingEngineCallDebuggerEnd(m *base.Module, l0 int32, l1 int32)
//go:linkname F_scriptingEngineDebuggerLogWithMaxLen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineDebuggerLogWithMaxLen
func F_scriptingEngineDebuggerLogWithMaxLen(m *base.Module, l0 int32)
//go:linkname F_scriptingEngineDebuggerKillForkedSessions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scriptingEngineDebuggerKillForkedSessions
func F_scriptingEngineDebuggerKillForkedSessions(m *base.Module)
//go:linkname F_scriptingEngineDebuggerLogRespReplyStr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_scriptingEngineDebuggerLogRespReplyStr
func F_scriptingEngineDebuggerLogRespReplyStr(m *base.Module, l0 int32)
//go:linkname F__sdsnewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__sdsnewlen
func F__sdsnewlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdswrite github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdswrite
func F_sdswrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_sdstrynewlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdstrynewlen
func F_sdstrynewlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsnew github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsnew
func F_sdsnew(m *base.Module, l0 int32) int32
//go:linkname F_sdsfree github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfree
func F_sdsfree(m *base.Module, l0 int32)
//go:linkname F_sdsAllocSize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsAllocSize
func F_sdsAllocSize(m *base.Module, l0 int32) int32
//go:linkname F__sdsMakeRoomFor github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__sdsMakeRoomFor
func F__sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsMakeRoomForNonGreedy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsMakeRoomForNonGreedy
func F_sdsMakeRoomForNonGreedy(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsRemoveFreeSpace github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsRemoveFreeSpace
func F_sdsRemoveFreeSpace(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsResize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsResize
func F_sdsResize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsIncrLen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsIncrLen
func F_sdsIncrLen(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sdsgrowzero github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsgrowzero
func F_sdsgrowzero(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdscatlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatlen
func F_sdscatlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscatsds github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatsds
func F_sdscatsds(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdsfromlonglong github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfromlonglong
func F_sdsfromlonglong(m *base.Module, l0 int64) int32
//go:linkname F_sdscatfmt github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscatfmt
func F_sdscatfmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdscmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdscmp
func F_sdscmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sdssplitlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdssplitlen
func F_sdssplitlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_sdsfreesplitres github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsfreesplitres
func F_sdsfreesplitres(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sdscatrepr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdscatrepr
func F_sdscatrepr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsparsearg github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsparsearg
func F_sdsparsearg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_sdsnsplitargs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sdsnsplitargs
func F_sdsnsplitargs(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sdsjoin github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sdsjoin
func F_sdsjoin(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelCheckConfigFile github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelCheckConfigFile
func F_sentinelCheckConfigFile(m *base.Module)
//go:linkname F_sentinelIsRunning github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelIsRunning
func F_sentinelIsRunning(m *base.Module)
//go:linkname F_sentinelEvent github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelEvent
func F_sentinelEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_createSentinelAddr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSentinelAddr
func F_createSentinelAddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sentinelTryConnectionSharing github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelTryConnectionSharing
func F_sentinelTryConnectionSharing(m *base.Module, l0 int32) int32
//go:linkname F_createSentinelValkeyInstance github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSentinelValkeyInstance
func F_createSentinelValkeyInstance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_removeMatchingSentinelFromPrimary github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_removeMatchingSentinelFromPrimary
func F_removeMatchingSentinelFromPrimary(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sentinelGetPrimaryByName github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelGetPrimaryByName
func F_sentinelGetPrimaryByName(m *base.Module, l0 int32) int32
//go:linkname F_sentinelResetPrimaryAndChangeAddress github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelResetPrimaryAndChangeAddress
func F_sentinelResetPrimaryAndChangeAddress(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_queueSentinelConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_queueSentinelConfig
func F_queueSentinelConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_loadSentinelConfigFromQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_loadSentinelConfigFromQueue
func F_loadSentinelConfigFromQueue(m *base.Module)
//go:linkname F_sentinelReconnectInstance github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelReconnectInstance
func F_sentinelReconnectInstance(m *base.Module, l0 int32)
//go:linkname F_sentinelSendReplicaOf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelSendReplicaOf
func F_sentinelSendReplicaOf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sentinelCheckSubjectivelyDown github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelCheckSubjectivelyDown
func F_sentinelCheckSubjectivelyDown(m *base.Module, l0 int32)
//go:linkname F_sentinelGetLeader github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelGetLeader
func F_sentinelGetLeader(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_sentinelFailoverTo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelFailoverTo
func F_sentinelFailoverTo(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_sentinelStartFailoverIfNeeded github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sentinelStartFailoverIfNeeded
func F_sentinelStartFailoverIfNeeded(m *base.Module, l0 int32) int32
//go:linkname F_sentinelFailoverStateMachine github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sentinelFailoverStateMachine
func F_sentinelFailoverStateMachine(m *base.Module, l0 int32)
//go:linkname F_serverLogRaw github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_serverLogRaw
func F_serverLogRaw(m *base.Module, l0 int32, l1 int32)
//go:linkname F__serverLog github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__serverLog
func F__serverLog(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_commandTimeSnapshot github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commandTimeSnapshot
func F_commandTimeSnapshot(m *base.Module) int64
//go:linkname F_resetChildState github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_resetChildState
func F_resetChildState(m *base.Module)
//go:linkname F_updateClientMemUsageAndBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_updateClientMemUsageAndBucket
func F_updateClientMemUsageAndBucket(m *base.Module, l0 int32) int32
//go:linkname F_prepareForShutdown github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_prepareForShutdown
func F_prepareForShutdown(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_closeListeningSockets github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_closeListeningSockets
func F_closeListeningSockets(m *base.Module, l0 int32)
//go:linkname F_initServerConfig github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_initServerConfig
func F_initServerConfig(m *base.Module)
//go:linkname F_populateCommandStructure github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_populateCommandStructure
func F_populateCommandStructure(m *base.Module, l0 int32) int32
//go:linkname F_setOOMScoreAdj github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setOOMScoreAdj
func F_setOOMScoreAdj(m *base.Module, l0 int32) int32
//go:linkname F_createSocketAcceptHandler github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSocketAcceptHandler
func F_createSocketAcceptHandler(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_initServer github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_initServer
func F_initServer(m *base.Module)
//go:linkname F_initListeners github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_initListeners
func F_initListeners(m *base.Module)
//go:linkname F_populateCommandLegacyRangeSpec github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_populateCommandLegacyRangeSpec
func F_populateCommandLegacyRangeSpec(m *base.Module, l0 int32)
//go:linkname F_catSubCommandFullname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_catSubCommandFullname
func F_catSubCommandFullname(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lookupCommandByCString github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lookupCommandByCString
func F_lookupCommandByCString(m *base.Module, l0 int32) int32
//go:linkname F_lookupCommandOrOriginal github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lookupCommandOrOriginal
func F_lookupCommandOrOriginal(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_alsoPropagate github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_alsoPropagate
func F_alsoPropagate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_forceCommandPropagation github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_forceCommandPropagation
func F_forceCommandPropagation(m *base.Module, l0 int32, l1 int32)
//go:linkname F_updateCommandLatencyHistogram github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_updateCommandLatencyHistogram
func F_updateCommandLatencyHistogram(m *base.Module, l0 int32, l1 int64)
//go:linkname F_postExecutionUnitOperations github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_postExecutionUnitOperations
func F_postExecutionUnitOperations(m *base.Module)
//go:linkname F_rejectCommandSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rejectCommandSds
func F_rejectCommandSds(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_commandCheckArity github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commandCheckArity
func F_commandCheckArity(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getCommandFlags github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_getCommandFlags
func F_getCommandFlags(m *base.Module, l0 int32) int64
//go:linkname F_prepareCommandQueue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_prepareCommandQueue
func F_prepareCommandQueue(m *base.Module, l0 int32)
//go:linkname F_writeCommandsGetDiskErrorMessage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_writeCommandsGetDiskErrorMessage
func F_writeCommandsGetDiskErrorMessage(m *base.Module, l0 int32) int32
//go:linkname F_addReplyFlagsForCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addReplyFlagsForCommand
func F_addReplyFlagsForCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyCommandKeySpecs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyCommandKeySpecs
func F_addReplyCommandKeySpecs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_addReplyCommandSubCommands github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_addReplyCommandSubCommands
func F_addReplyCommandSubCommands(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_generateCommandResponse github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_generateCommandResponse
func F_generateCommandResponse(m *base.Module, l0 int32) int32
//go:linkname F_commandListWithFilter github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_commandListWithFilter
func F_commandListWithFilter(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_createPidFile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_createPidFile
func F_createPidFile(m *base.Module)
//go:linkname F_daemonize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_daemonize
func F_daemonize(m *base.Module)
//go:linkname F_usage github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_usage
func F_usage(m *base.Module)
//go:linkname F_listenerByType github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listenerByType
func F_listenerByType(m *base.Module, l0 int32) int32
//go:linkname F_changeListener github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_changeListener
func F_changeListener(m *base.Module, l0 int32) int32
//go:linkname F_serverFork github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_serverFork
func F_serverFork(m *base.Module, l0 int32) int32
//go:linkname F_sendChildCowInfo github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sendChildCowInfo
func F_sendChildCowInfo(m *base.Module, l0 int32, l1 int32)
//go:linkname F_validateProcTitleTemplate github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_validateProcTitleTemplate
func F_validateProcTitleTemplate(m *base.Module, l0 int32) int32
//go:linkname F_serverIsSupervised github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_serverIsSupervised
func F_serverIsSupervised(m *base.Module, l0 int32) int32
//go:linkname F_SHA1Final github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_SHA1Final
func F_SHA1Final(m *base.Module, l0 int32, l1 int32)
//go:linkname F_sha256_update github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sha256_update
func F_sha256_update(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_siphash_nocase github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_siphash_nocase
func F_siphash_nocase(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F_connNonBlock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connNonBlock
func F_connNonBlock(m *base.Module, l0 int32) int32
//go:linkname F_connSendTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connSendTimeout
func F_connSendTimeout(m *base.Module, l0 int32, l1 int64) int32
//go:linkname F_lookupKeyByPattern github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lookupKeyByPattern
func F_lookupKeyByPattern(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_createSparklineSequence github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_createSparklineSequence
func F_createSparklineSequence(m *base.Module) int32
//go:linkname F_freeSparklineSequence github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeSparklineSequence
func F_freeSparklineSequence(m *base.Module, l0 int32)
//go:linkname F_sparklineRender github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sparklineRender
func F_sparklineRender(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_syncRead github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_syncRead
func F_syncRead(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32
//go:linkname F_hashTypeTrackEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeTrackEntry
func F_hashTypeTrackEntry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeTryConversion github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeTryConversion
func F_hashTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_hashTypeConvert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeConvert
func F_hashTypeConvert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeGetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeGetValue
func F_hashTypeGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_hashTypeTrackUpdateEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeTrackUpdateEntry
func F_hashTypeTrackUpdateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64)
//go:linkname F_hashTypeLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeLength
func F_hashTypeLength(m *base.Module, l0 int32) int32
//go:linkname F_hashTypeUntrackEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeUntrackEntry
func F_hashTypeUntrackEntry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeInitIterator
func F_hashTypeInitIterator(m *base.Module, l0 int32, l1 int32)
//go:linkname F_hashTypeResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeResetIterator
func F_hashTypeResetIterator(m *base.Module, l0 int32)
//go:linkname F_hashTypeCurrentObjectNewSds github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hashTypeCurrentObjectNewSds
func F_hashTypeCurrentObjectNewSds(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_hashTypePersist github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypePersist
func F_hashTypePersist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_addHashIteratorCursorToReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addHashIteratorCursorToReply
func F_addHashIteratorCursorToReply(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_hexpireGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hexpireGenericCommand
func F_hexpireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_httlGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_httlGenericCommand
func F_httlGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32)
//go:linkname F_hashTypeRandomElement github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_hashTypeRandomElement
func F_hashTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_hrandfieldReplyWithListpack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hrandfieldReplyWithListpack
func F_hrandfieldReplyWithListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_listTypeTryConversion github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_listTypeTryConversion
func F_listTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_listTypeLength github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeLength
func F_listTypeLength(m *base.Module, l0 int32) int32
//go:linkname F_listTypeReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeReleaseIterator
func F_listTypeReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_listTypeGetValue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeGetValue
func F_listTypeGetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_listTypeGet github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeGet
func F_listTypeGet(m *base.Module, l0 int32) int32
//go:linkname F_listTypeInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_listTypeInsert
func F_listTypeInsert(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_pushGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pushGenericCommand
func F_pushGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_popGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_popGenericCommand
func F_popGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_mpopGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_mpopGenericCommand
func F_mpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_lmoveGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lmoveGenericCommand
func F_lmoveGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_blmoveGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_blmoveGenericCommand
func F_blmoveGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64)
//go:linkname F_blockingPopGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_blockingPopGenericCommand
func F_blockingPopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_setTypeConvertAndExpand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setTypeConvertAndExpand
func F_setTypeConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_maybeConvertIntset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_maybeConvertIntset
func F_maybeConvertIntset(m *base.Module, l0 int32)
//go:linkname F_setTypeRemoveAux github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeRemoveAux
func F_setTypeRemoveAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32
//go:linkname F_setTypeReleaseIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeReleaseIterator
func F_setTypeReleaseIterator(m *base.Module, l0 int32)
//go:linkname F_setTypeNextObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeNextObject
func F_setTypeNextObject(m *base.Module, l0 int32) int32
//go:linkname F_setTypeRandomElement github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeRandomElement
func F_setTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_setTypePopRandom github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_setTypePopRandom
func F_setTypePopRandom(m *base.Module, l0 int32) int32
//go:linkname F_setTypeConvert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setTypeConvert
func F_setTypeConvert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_srandmemberWithCountCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_srandmemberWithCountCommand
func F_srandmemberWithCountCommand(m *base.Module, l0 int32)
//go:linkname F_sinterGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sinterGenericCommand
func F_sinterGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_streamCreateCG github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamCreateCG
func F_streamCreateCG(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32
//go:linkname F_streamCreateNACK github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamCreateNACK
func F_streamCreateNACK(m *base.Module, l0 int32) int32
//go:linkname F_streamIteratorGetID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamIteratorGetID
func F_streamIteratorGetID(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_streamIteratorStop github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamIteratorStop
func F_streamIteratorStop(m *base.Module, l0 int32)
//go:linkname F_streamTrim github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamTrim
func F_streamTrim(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_streamIteratorGetField github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamIteratorGetField
func F_streamIteratorGetField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_streamDeleteItem github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamDeleteItem
func F_streamDeleteItem(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_streamLastValidID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamLastValidID
func F_streamLastValidID(m *base.Module, l0 int32, l1 int32)
//go:linkname F_streamPropagateGroupID github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamPropagateGroupID
func F_streamPropagateGroupID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_streamFreeNACK github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamFreeNACK
func F_streamFreeNACK(m *base.Module, l0 int32)
//go:linkname F_streamGenericParseIDOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamGenericParseIDOrReply
func F_streamGenericParseIDOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32
//go:linkname F_streamParseAddOrTrimArgsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamParseAddOrTrimArgsOrReply
func F_streamParseAddOrTrimArgsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_xrangeGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_xrangeGenericCommand
func F_xrangeGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_streamCreateConsumer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_streamCreateConsumer
func F_streamCreateConsumer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_streamValidateListpackIntegrity github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_streamValidateListpackIntegrity
func F_streamValidateListpackIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_incrDecrCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_incrDecrCommand
func F_incrDecrCommand(m *base.Module, l0 int32, l1 int64)
//go:linkname F_zslInsert github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zslInsert
func F_zslInsert(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F_zslCreateNode github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zslCreateNode
func F_zslCreateNode(m *base.Module, l0 int32, l1 float64, l2 int32) int32
//go:linkname F_zslInsertNode github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslInsertNode
func F_zslInsertNode(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetFreeLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetFreeLexRange
func F_zsetFreeLexRange(m *base.Module, l0 int32)
//go:linkname F_zsetParseLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetParseLexRange
func F_zsetParseLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zslNthInLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslNthInLexRange
func F_zslNthInLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zzlGetScore github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlGetScore
func F_zzlGetScore(m *base.Module, l0 int32) float64
//go:linkname F_zzlValidateScores github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zzlValidateScores
func F_zzlValidateScores(m *base.Module, l0 int32) int32
//go:linkname F_zzlLexValueGteMin github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlLexValueGteMin
func F_zzlLexValueGteMin(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zzlFirstInLexRange github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlFirstInLexRange
func F_zzlFirstInLexRange(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetConvert github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetConvert
func F_zsetConvert(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zsetScore github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetScore
func F_zsetScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zsetDel github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetDel
func F_zsetDel(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zsetRemoveFromSkiplist github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zsetRemoveFromSkiplist
func F_zsetRemoveFromSkiplist(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zaddGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zaddGenericCommand
func F_zaddGenericCommand(m *base.Module, l0 int32, l1 int32)
//go:linkname F_zzlDeleteRangeByScore github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zzlDeleteRangeByScore
func F_zzlDeleteRangeByScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zslDeleteRangeByScore github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslDeleteRangeByScore
func F_zslDeleteRangeByScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zslDeleteRangeByLex github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zslDeleteRangeByLex
func F_zslDeleteRangeByLex(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zunionInterDiffGenericCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zunionInterDiffGenericCommand
func F_zunionInterDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_zuiInitIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zuiInitIterator
func F_zuiInitIterator(m *base.Module, l0 int32)
//go:linkname F_zuiNext github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zuiNext
func F_zuiNext(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zuiFind github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zuiFind
func F_zuiFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_zuiClearIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zuiClearIterator
func F_zuiClearIterator(m *base.Module, l0 int32)
//go:linkname F_genericZrangebyrankCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericZrangebyrankCommand
func F_genericZrangebyrankCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_genericZrangebyscoreCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericZrangebyscoreCommand
func F_genericZrangebyscoreCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)
//go:linkname F_genericZpopCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_genericZpopCommand
func F_genericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)
//go:linkname F_zrandmemberWithCountCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zrandmemberWithCountCommand
func F_zrandmemberWithCountCommand(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_zsetTypeRandomElement github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zsetTypeRandomElement
func F_zsetTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_addClientToTimeoutTable github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_addClientToTimeoutTable
func F_addClientToTimeoutTable(m *base.Module, l0 int32)
//go:linkname F_handleBlockedClientsTimeout github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_handleBlockedClientsTimeout
func F_handleBlockedClientsTimeout(m *base.Module)
//go:linkname F_getTimeoutFromObjectOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getTimeoutFromObjectOrReply
func F_getTimeoutFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_checkPrefixCollisionsOrReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_checkPrefixCollisionsOrReply
func F_checkPrefixCollisionsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_enableTracking github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_enableTracking
func F_enableTracking(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32)
//go:linkname F_trackingRememberKeys github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_trackingRememberKeys
func F_trackingRememberKeys(m *base.Module, l0 int32, l1 int32)
//go:linkname F_trackingInvalidateKey github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_trackingInvalidateKey
func F_trackingInvalidateKey(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_trackingHandlePendingKeyInvalidations github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trackingHandlePendingKeyInvalidations
func F_trackingHandlePendingKeyInvalidations(m *base.Module)
//go:linkname F_freeTrackingRadixTree github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeTrackingRadixTree
func F_freeTrackingRadixTree(m *base.Module, l0 int32)
//go:linkname F_trackingLimitUsedSlots github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_trackingLimitUsedSlots
func F_trackingLimitUsedSlots(m *base.Module)
//go:linkname F_trackingBuildBroadcastReply github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_trackingBuildBroadcastReply
func F_trackingBuildBroadcastReply(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_stringmatchlen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_stringmatchlen
func F_stringmatchlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_stringmatchlen_impl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_stringmatchlen_impl
func F_stringmatchlen_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
//go:linkname F_memtoull github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memtoull
func F_memtoull(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_ll2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ll2string
func F_ll2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_ull2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ull2string
func F_ull2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32
//go:linkname F_string2ll github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_string2ll
func F_string2ll(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_d2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_d2string
func F_d2string(m *base.Module, l0 int32, l1 int32, l2 float64) int32
//go:linkname F_ld2string github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ld2string
func F_ld2string(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32
//go:linkname F_initializeRandomSeed github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_initializeRandomSeed
func F_initializeRandomSeed(m *base.Module)
//go:linkname F_getRandomHexChars github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getRandomHexChars
func F_getRandomHexChars(m *base.Module, l0 int32, l1 int32)
//go:linkname F_getAbsolutePath github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getAbsolutePath
func F_getAbsolutePath(m *base.Module, l0 int32) int32
//go:linkname F_mstime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mstime
func F_mstime(m *base.Module) int64
//go:linkname F_wangHash64 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_wangHash64
func F_wangHash64(m *base.Module, l0 int64) int64
//go:linkname F_processRESP github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_processRESP
func F_processRESP(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_redis_check_aof_main github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_redis_check_aof_main
func F_redis_check_aof_main(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_rdbCheckError github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_rdbCheckError
func F_rdbCheckError(m *base.Module, l0 int32, l1 int32)
//go:linkname F_redis_check_rdb_main github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_redis_check_rdb_main
func F_redis_check_rdb_main(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_ffc_from_chars_double_options github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ffc_from_chars_double_options
func F_ffc_from_chars_double_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_valkey_strtod_n github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_strtod_n
func F_valkey_strtod_n(m *base.Module, l0 int32, l1 int32, l2 int32) float64
//go:linkname F_ffc_bigint_long_mul github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ffc_bigint_long_mul
func F_ffc_bigint_long_mul(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pvSplit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pvSplit
func F_pvSplit(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_pvInsertAt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pvInsertAt
func F_pvInsertAt(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsetAddEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsetAddEntry
func F_vsetAddEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_insertToBucket_VECTOR github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_insertToBucket_VECTOR
func F_insertToBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_removeFromBucket_VECTOR github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_removeFromBucket_VECTOR
func F_removeFromBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_removeEntryFromRaxBucket github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_removeEntryFromRaxBucket
func F_removeEntryFromRaxBucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_shrinkRaxBucketIfPossible github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_shrinkRaxBucketIfPossible
func F_shrinkRaxBucketIfPossible(m *base.Module, l0 int32, l1 int32)
//go:linkname F_vsetBucketRemoveExpired_HASHTABLE github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vsetBucketRemoveExpired_HASHTABLE
func F_vsetBucketRemoveExpired_HASHTABLE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_vsetResetIterator github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsetResetIterator
func F_vsetResetIterator(m *base.Module, l0 int32)
//go:linkname F_vsetIsEmpty github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vsetIsEmpty
func F_vsetIsEmpty(m *base.Module, l0 int32) int32
//go:linkname F_zipEntry github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_zipEntry
func F_zipEntry(m *base.Module, l0 int32, l1 int32)
//go:linkname F_valkey_malloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_malloc
func F_valkey_malloc(m *base.Module, l0 int32) int32
//go:linkname F_zmalloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zmalloc_usable
func F_zmalloc_usable(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkey_calloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_calloc
func F_valkey_calloc(m *base.Module, l0 int32) int32
//go:linkname F_ztryrealloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ztryrealloc_usable
func F_ztryrealloc_usable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_valkey_realloc github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkey_realloc
func F_valkey_realloc(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_zrealloc_usable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zrealloc_usable
func F_zrealloc_usable(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_valkey_free github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkey_free
func F_valkey_free(m *base.Module, l0 int32)
//go:linkname F_zfree_with_size github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_zfree_with_size
func F_zfree_with_size(m *base.Module, l0 int32, l1 int32)
//go:linkname F_spmcEnqueue github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_spmcEnqueue
func F_spmcEnqueue(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_spscIsFull github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_spscIsFull
func F_spscIsFull(m *base.Module, l0 int32) int32
//go:linkname F_valkeyAsyncConnectWithOptions github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyAsyncConnectWithOptions
func F_valkeyAsyncConnectWithOptions(m *base.Module, l0 int32) int32
//go:linkname F_valkeyGetSubscribeCallback github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkeyGetSubscribeCallback
func F_valkeyGetSubscribeCallback(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_nextArgument github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_nextArgument
func F_nextArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_valkeyContextRegisterFuncs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyContextRegisterFuncs
func F_valkeyContextRegisterFuncs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_valkeyContextWaitReady github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyContextWaitReady
func F_valkeyContextWaitReady(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_moveToNextTask github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_moveToNextTask
func F_moveToNextTask(m *base.Module, l0 int32)
//go:linkname F_freeReplyObject github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_freeReplyObject
func F_freeReplyObject(m *base.Module, l0 int32)
//go:linkname F_valkeyvFormatCommand github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyvFormatCommand
func F_valkeyvFormatCommand(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_valkeySetError github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_valkeySetError
func F_valkeySetError(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_valkeyBufferRead github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyBufferRead
func F_valkeyBufferRead(m *base.Module, l0 int32) int32
//go:linkname F_valkeyAppendCmdLen github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_valkeyAppendCmdLen
func F_valkeyAppendCmdLen(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_hdr_close github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_hdr_close
func F_hdr_close(m *base.Module, l0 int32)
//go:linkname F_fpconv_dtoa github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fpconv_dtoa
func F_fpconv_dtoa(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F_ldbStart github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ldbStart
func F_ldbStart(m *base.Module, l0 int32)
//go:linkname F_ldbEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ldbEnd
func F_ldbEnd(m *base.Module)
//go:linkname F_ldbCatStackValueRec github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ldbCatStackValueRec
func F_ldbCatStackValueRec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_ldbGenerateDebuggerCommandsArray github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_ldbGenerateDebuggerCommandsArray
func F_ldbGenerateDebuggerCommandsArray(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lm_asprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lm_asprintf
func F_lm_asprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_copy_string_from_lua_stack github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_copy_string_from_lua_stack
func F_copy_string_from_lua_stack(m *base.Module, l0 int32) int32
//go:linkname F_luaSetTableProtectionRecursively github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaSetTableProtectionRecursively
func F_luaSetTableProtectionRecursively(m *base.Module, l0 int32)
//go:linkname F_luaSetTableProtectionForBasicTypes github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaSetTableProtectionForBasicTypes
func F_luaSetTableProtectionForBasicTypes(m *base.Module, l0 int32)
//go:linkname F_luaRegisterVersion github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaRegisterVersion
func F_luaRegisterVersion(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaRegisterServerAPI github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaRegisterServerAPI
func F_luaRegisterServerAPI(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaExtractErrorInformation github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaExtractErrorInformation
func F_luaExtractErrorInformation(m *base.Module, l0 int32, l1 int32)
//go:linkname F__serverPanic_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__serverPanic_2
func F__serverPanic_2(m *base.Module, l0 int32)
//go:linkname F_processCollectionElementEnd github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_processCollectionElementEnd
func F_processCollectionElementEnd(m *base.Module, l0 int32)
//go:linkname F_lua_checkstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_checkstack
func F_lua_checkstack(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_replace github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_replace
func F_lua_replace(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_isnumber github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_isnumber
func F_lua_isnumber(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_tointeger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_tointeger
func F_lua_tointeger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_tolstring github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_tolstring
func F_lua_tolstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_objlen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_objlen
func F_lua_objlen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_pushfstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_pushfstring
func F_lua_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_pushcclosure github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_pushcclosure
func F_lua_pushcclosure(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_gettable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_gettable
func F_lua_gettable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_getfield github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_getfield
func F_lua_getfield(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_createtable github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_createtable
func F_lua_createtable(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_settable github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_settable
func F_lua_settable(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_setfield github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_setfield
func F_lua_setfield(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_rawset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_rawset
func F_lua_rawset(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_call github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_call
func F_lua_call(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_lua_dump github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_dump
func F_lua_dump(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_lua_next github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_next
func F_lua_next(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_lua_concat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_concat
func F_lua_concat(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_getinfo github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lua_getinfo
func F_lua_getinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaG_typeerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaG_typeerror
func F_luaG_typeerror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaG_concaterror github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaG_concaterror
func F_luaG_concaterror(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaD_throw github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_throw
func F_luaD_throw(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaD_reallocstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_reallocstack
func F_luaD_reallocstack(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaD_growstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_growstack
func F_luaD_growstack(m *base.Module, l0 int32, l1 int32)
//go:linkname F_growCI github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_growCI
func F_growCI(m *base.Module, l0 int32) int32
//go:linkname F_luaD_poscall github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_poscall
func F_luaD_poscall(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaD_pcall github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaD_pcall
func F_luaD_pcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_luaF_close github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaF_close
func F_luaF_close(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaF_getlocalname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaF_getlocalname
func F_luaF_getlocalname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaC_callGCTM github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaC_callGCTM
func F_luaC_callGCTM(m *base.Module, l0 int32)
//go:linkname F_sweeplist github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sweeplist
func F_sweeplist(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaC_step github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaC_step
func F_luaC_step(m *base.Module, l0 int32)
//go:linkname F_singlestep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_singlestep
func F_singlestep(m *base.Module, l0 int32) int32
//go:linkname F_reallymarkobject github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_reallymarkobject
func F_reallymarkobject(m *base.Module, l0 int32, l1 int32)
//go:linkname F_markmt github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_markmt
func F_markmt(m *base.Module, l0 int32)
//go:linkname F_luaM_growaux_ github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaM_growaux_
func F_luaM_growaux_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_luaM_realloc_ github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaM_realloc_
func F_luaM_realloc_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaO_str2d github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaO_str2d
func F_luaO_str2d(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_save github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_save
func F_save(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaX_syntaxerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaX_syntaxerror
func F_luaX_syntaxerror(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaX_setinput github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaX_setinput
func F_luaX_setinput(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_luaX_next github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaX_next
func F_luaX_next(m *base.Module, l0 int32)
//go:linkname F_luaK_code github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_code
func F_luaK_code(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaK_codeABC github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_codeABC
func F_luaK_codeABC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_patchlistaux github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_patchlistaux
func F_patchlistaux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_luaK_patchtohere github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_patchtohere
func F_luaK_patchtohere(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_reserveregs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaK_reserveregs
func F_luaK_reserveregs(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaK_exp2nextreg github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_exp2nextreg
func F_luaK_exp2nextreg(m *base.Module, l0 int32, l1 int32)
//go:linkname F_discharge2reg github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_discharge2reg
func F_discharge2reg(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_exp2RK github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaK_exp2RK
func F_luaK_exp2RK(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_jumponcond github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_jumponcond
func F_jumponcond(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaK_prefix github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_prefix
func F_luaK_prefix(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_infix github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaK_infix
func F_luaK_infix(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaK_posfix github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaK_posfix
func F_luaK_posfix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_open_func github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_open_func
func F_open_func(m *base.Module, l0 int32, l1 int32)
//go:linkname F_chunk github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_chunk
func F_chunk(m *base.Module, l0 int32)
//go:linkname F_close_func github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_close_func
func F_close_func(m *base.Module, l0 int32)
//go:linkname F_body github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_body
func F_body(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_primaryexp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_primaryexp
func F_primaryexp(m *base.Module, l0 int32, l1 int32)
//go:linkname F_adjust_assign github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_adjust_assign
func F_adjust_assign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_constructor github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_constructor
func F_constructor(m *base.Module, l0 int32, l1 int32)
//go:linkname F_singlevaraux github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_singlevaraux
func F_singlevaraux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaE_newthread github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaE_newthread
func F_luaE_newthread(m *base.Module, l0 int32) int32
//go:linkname F_close_state github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_close_state
func F_close_state(m *base.Module, l0 int32)
//go:linkname F_luaS_newlstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaS_newlstr
func F_luaS_newlstr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_setnodevector github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_setnodevector
func F_setnodevector(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaH_get github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_get
func F_luaH_get(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_new github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_new
func F_luaH_new(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_getnum github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_getnum
func F_luaH_getnum(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_getstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_getstr
func F_luaH_getstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaH_set github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaH_set
func F_luaH_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaH_setstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaH_setstr
func F_luaH_setstr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaT_init github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaT_init
func F_luaT_init(m *base.Module, l0 int32)
//go:linkname F_luaV_lessthan github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaV_lessthan
func F_luaV_lessthan(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_call_binTM github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_call_binTM
func F_call_binTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_luaV_execute github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaV_execute
func F_luaV_execute(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaZ_fill github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaZ_fill
func F_luaZ_fill(m *base.Module, l0 int32) int32
//go:linkname F_luaL_argerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_argerror
func F_luaL_argerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_typerror github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_typerror
func F_luaL_typerror(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_checklstring github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_checklstring
func F_luaL_checklstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_optlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_optlstring
func F_luaL_optlstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_checkstack github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checkstack
func F_luaL_checkstack(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_checktype github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checktype
func F_luaL_checktype(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_checkinteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_checkinteger
func F_luaL_checkinteger(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaL_optinteger github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_optinteger
func F_luaL_optinteger(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_luaL_openlib github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_openlib
func F_luaL_openlib(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_luaL_addlstring github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_addlstring
func F_luaL_addlstring(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_addvalue github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_addvalue
func F_luaL_addvalue(m *base.Module, l0 int32)
//go:linkname F_luaL_unref github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaL_unref
func F_luaL_unref(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_luaL_loadfile github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_loadfile
func F_luaL_loadfile(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_luaL_loadbuffer github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_loadbuffer
func F_luaL_loadbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_luaL_newstate github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_luaL_newstate
func F_luaL_newstate(m *base.Module) int32
//go:linkname F_auxsort github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_auxsort
func F_auxsort(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_fpconv_strtod github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fpconv_strtod
func F_fpconv_strtod(m *base.Module, l0 int32, l1 int32) float64
//go:linkname F_strbuf_resize github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strbuf_resize
func F_strbuf_resize(m *base.Module, l0 int32, l1 int32)
//go:linkname F_lua_cjson_new github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_lua_cjson_new
func F_lua_cjson_new(m *base.Module, l0 int32) int32
//go:linkname F_json_arg_init github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_json_arg_init
func F_json_arg_init(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_json_enum_option github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_json_enum_option
func F_json_enum_option(m *base.Module, l0 int32, l1 int32, l2 int32)
//go:linkname F_json_next_number_token github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_json_next_number_token
func F_json_next_number_token(m *base.Module, l0 int32, l1 int32)
//go:linkname F_optsize github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_optsize
func F_optsize(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_controloptions github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_controloptions
func F_controloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_mp_decode_to_lua_type github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mp_decode_to_lua_type
func F_mp_decode_to_lua_type(m *base.Module, l0 int32, l1 int32)
//go:linkname F_luaopen_create github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_luaopen_create
func F_luaopen_create(m *base.Module, l0 int32) int32
//go:linkname F___memcpy github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___memcpy
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memcpy_bulkmem github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_memcpy_bulkmem
func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___memset github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___memset
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F__emscripten_memset_bulkmem github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_memset_bulkmem
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___errno_location github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___errno_location
func F___errno_location(m *base.Module) int32
//go:linkname F_do_tzset github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_do_tzset
func F_do_tzset(m *base.Module)
//go:linkname F__Exit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__Exit
func F__Exit(m *base.Module, l0 int32)
//go:linkname F_abort github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_abort
func F_abort(m *base.Module)
//go:linkname F_access github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_access
func F_access(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___clock_gettime github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___clock_gettime
func F___clock_gettime(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_close github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_close
func F_close(m *base.Module, l0 int32) int32
//go:linkname F___rem_pio2 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___rem_pio2
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32
//go:linkname F_dirname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_dirname
func F_dirname(m *base.Module, l0 int32) int32
//go:linkname F_memmove github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memmove
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___time github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___time
func F___time(m *base.Module, l0 int32) int64
//go:linkname F_fp_barrier_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fp_barrier_1
func F_fp_barrier_1(m *base.Module, l0 float64) float64
//go:linkname F___math_oflow github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___math_oflow
func F___math_oflow(m *base.Module, l0 int32) float64
//go:linkname F_fabs github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fabs
func F_fabs(m *base.Module, l0 float64) float64
//go:linkname F_fclose github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fclose
func F_fclose(m *base.Module, l0 int32) int32
//go:linkname F_fcntl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fcntl
func F_fcntl(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_fflush github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fflush
func F_fflush(m *base.Module, l0 int32) int32
//go:linkname F___fdopen github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___fdopen
func F___fdopen(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fiprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fiprintf
func F_fiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___overflow github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___overflow
func F___overflow(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_do_putc_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_do_putc_1
func F_do_putc_1(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fputs github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fputs
func F_fputs(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_frexp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_frexp
func F_frexp(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F_fseek github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fseek
func F_fseek(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___fstatat github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___fstatat
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___ftello github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___ftello
func F___ftello(m *base.Module, l0 int32) int64
//go:linkname F_fwrite github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fwrite
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_locking_getc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_locking_getc
func F_locking_getc(m *base.Module, l0 int32) int32
//go:linkname F_getcwd github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getcwd
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___syscall_getpid github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___syscall_getpid
func F___syscall_getpid(m *base.Module) int32
//go:linkname F___syscall_umask github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_umask
func F___syscall_umask(m *base.Module, l0 int32) int32
//go:linkname F_pat_next github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_pat_next
func F_pat_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_fnmatch_internal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_fnmatch_internal
func F_fnmatch_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F___bswap_32_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___bswap_32_1
func F___bswap_32_1(m *base.Module, l0 int32) int32
//go:linkname F___bswap_16_1 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___bswap_16_1
func F___bswap_16_1(m *base.Module, l0 int32) int32
//go:linkname F_inet_ntop github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_inet_ntop
func F_inet_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_inet_pton github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_inet_pton
func F_inet_pton(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_convert_ioctl_struct github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_convert_ioctl_struct
func F_convert_ioctl_struct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_isalnum github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_isalnum
func F_isalnum(m *base.Module, l0 int32) int32
//go:linkname F_iscntrl github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iscntrl
func F_iscntrl(m *base.Module, l0 int32) int32
//go:linkname F_ispunct github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_ispunct
func F_ispunct(m *base.Module, l0 int32) int32
//go:linkname F_isblank github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isblank
func F_isblank(m *base.Module, l0 int32) int32
//go:linkname F_iswprint github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iswprint
func F_iswprint(m *base.Module, l0 int32) int32
//go:linkname F_iswspace github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_iswspace
func F_iswspace(m *base.Module, l0 int32) int32
//go:linkname F_isxdigit github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_isxdigit
func F_isxdigit(m *base.Module, l0 int32) int32
//go:linkname F_kill github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_kill
func F_kill(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F__emscripten_yield github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F__emscripten_yield
func F__emscripten_yield(m *base.Module, l0 float64)
//go:linkname F___lock github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___lock
func F___lock(m *base.Module, l0 int32)
//go:linkname F_log github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_log
func F_log(m *base.Module, l0 float64) float64
//go:linkname F___lseek github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___lseek
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64
//go:linkname F_memchr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_memchr
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_memcmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_memcmp
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_mkdir github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_mkdir
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___mkostemps github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___mkostemps
func F___mkostemps(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___bswap_16_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___bswap_16_2
func F___bswap_16_2(m *base.Module, l0 int32) int32
//go:linkname F_open github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_open
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_pipe github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pipe
func F_pipe(m *base.Module, l0 int32) int32
//go:linkname F_top12_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_top12_2
func F_top12_2(m *base.Module, l0 float64) int32
//go:linkname F_specialcase_2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_specialcase_2
func F_specialcase_2(m *base.Module, l0 float64, l1 int64, l2 int64) float64
//go:linkname F_iprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_iprintf
func F_iprintf(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___get_tp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___get_tp
func F___get_tp(m *base.Module) int32
//go:linkname F___qsort_r github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___qsort_r
func F___qsort_r(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F___srandom github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___srandom
func F___srandom(m *base.Module, l0 int32)
//go:linkname F_lcg64 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lcg64
func F_lcg64(m *base.Module, l0 int64) int64
//go:linkname F_lcg31 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_lcg31
func F_lcg31(m *base.Module, l0 int32) int32
//go:linkname F_rename github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_rename
func F_rename(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_scalbn github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scalbn
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64
//go:linkname F__emscripten_timeout github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F__emscripten_timeout
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64)
//go:linkname F_sigismember github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sigismember
func F_sigismember(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_sin github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sin
func F_sin(m *base.Module, l0 float64) float64
//go:linkname F_sleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sleep
func F_sleep(m *base.Module, l0 int32) int32
//go:linkname F_snprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_snprintf
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F___small_sprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___small_sprintf
func F___small_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_sscanf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_sscanf
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_stat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_stat
func F_stat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strcat
func F_strcat(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strchr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strchr
func F_strchr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strcmp github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strcmp
func F_strcmp(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_strerror github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strerror
func F_strerror(m *base.Module, l0 int32) int32
//go:linkname F___strftime_l github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___strftime_l
func F___strftime_l(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_strncat github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strncat
func F_strncat(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_strncpy github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_strncpy
func F_strncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_threebyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_threebyte_strstr
func F_threebyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_fourbyte_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fourbyte_strstr
func F_fourbyte_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F_twoway_strstr github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_twoway_strstr
func F_twoway_strstr(m *base.Module, l0 int32, l1 int32) int32
//go:linkname F___shgetc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___shgetc
func F___shgetc(m *base.Module, l0 int32) int32
//go:linkname F_fmodl github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_fmodl
func F_fmodl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___floatscan github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___floatscan
func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_scanexp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_scanexp
func F_scanexp(m *base.Module, l0 int32, l1 int32) int64
//go:linkname F_strtoull github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_strtoull
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64
//go:linkname F___syscall_ret github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___syscall_ret
func F___syscall_ret(m *base.Module, l0 int32) int32
//go:linkname F_dprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_dprintf
func F_dprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_tolower github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_tolower
func F_tolower(m *base.Module, l0 int32) int32
//go:linkname F_toupper github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_toupper
func F_toupper(m *base.Module, l0 int32) int32
//go:linkname F_towlower github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_towlower
func F_towlower(m *base.Module, l0 int32) int32
//go:linkname F_unlink github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_unlink
func F_unlink(m *base.Module, l0 int32) int32
//go:linkname F_usleep github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_usleep
func F_usleep(m *base.Module, l0 int32) int32
//go:linkname F___vfprintf_internal github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___vfprintf_internal
func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32
//go:linkname F_pop_arg github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pop_arg
func F_pop_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)
//go:linkname F_pad github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_pad
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)
//go:linkname F_vfprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vfprintf
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_vsnprintf github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_vsnprintf
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
//go:linkname F_vfscanf github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_vfscanf
func F_vfscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F___wasi_fd_is_valid github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___wasi_fd_is_valid
func F___wasi_fd_is_valid(m *base.Module, l0 int32) int32
//go:linkname F_wcrtomb github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_wcrtomb
func F_wcrtomb(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_writev github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_writev
func F_writev(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_emscripten_builtin_malloc github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_emscripten_builtin_malloc
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32
//go:linkname F___ashlti3 github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___ashlti3
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32)
//go:linkname F___wasm_longjmp github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F___wasm_longjmp
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32)
//go:linkname F___multf3 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___multf3
func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)
//go:linkname F___trunctfdf2 github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F___trunctfdf2
func F___trunctfdf2(m *base.Module, l0 int64, l1 int64) float64
//go:linkname F_connect github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_connect
func F_connect(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_getsockname github.com/shibukawa/vkmem/internal/aot/vkaot/p2.F_getsockname
func F_getsockname(m *base.Module, l0 int32, l1 int32, l2 int32) int32
//go:linkname F_recvfrom github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_recvfrom
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
//go:linkname F_sendto github.com/shibukawa/vkmem/internal/aot/vkaot/p0.F_sendto
func F_sendto(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32
