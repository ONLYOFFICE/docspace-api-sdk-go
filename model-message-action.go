// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
	"fmt"
)

// MessageAction [1000 - Login success, 1001 - Login success via social account, 1002 - Login fail invalid combination, 1003 - Login fail social account not found, 1004 - Login fail disabled profile, 1005 - Login fail, 1006 - Logout, 1007 - Login success via sms, 1008 - Login fail via sms, 1009 - Login fail ip security, 1010 - Login success via api, 1011 - Login success via social app, 1012 - Login success via api sms, 1013 - Login fail via api, 1014 - Login fail via api sms, 1015 - Login success via SSO, 1016 - Session started, 1017 - Session completed, 1018 - Login fail via SSO, 1019 - Login success via api social account, 1020 - Login fail via api social account, 1021 - Login succes via tfa app, 1022 - Login fail via Tfa app, 1023 - Login fail brute force, 1024 - Login success via api tfa, 1025 - Login fail via api tfa, 1026 - Login fail recaptcha, 1027 - Authorization link activated, 1028 - Login success via OAuth 2.0, 1029 - Login success via login and password, 4000 - User created, 4001 - Guest created, 4002 - User created via invite, 4003 - Guest created via invite, 4004 - User activated, 4005 - Guest activated, 4006 - User updated, 4007 - User updated language, 4008 - User added avatar, 4009 - User deleted avatar, 4010 - User updated avatar thumbnails, 4011 - User linked social account, 4012 - User unlinked social account, 4013 - User sent activation instructions, 4014 - User sent email change instructions, 4015 - User sent password change instructions, 4016 - User sent delete instructions, 4017 - User updated password, 4018 - User deleted, 4019 - Users updated type, 4020 - Users updated status, 4021 - Users sent activation instructions, 4022 - Users deleted, 4023 - Sent invite instructions, 4024 - User imported, 4025 - Guest imported, 4026 - Group created, 4027 - Group updated, 4028 - Group deleted, 4029 - User updated mobile number, 4030 - User data reassigns, 4031 - User data removing, 4032 - User connected tfa app, 4033 - User disconnected tfa app, 4034 - User logout active connections, 4035 - User logout active connection, 4036 - User logout active connections for user, 4037 - Send join invite, 5000 - File created, 5001 - File renamed, 5002 - File updated, 5003 - File created version, 5004 - File deleted version, 5005 - File updated revision comment, 5006 - File locked, 5007 - File unlocked, 5008 - File updated access, 5009 - File downloaded, 5010 - File downloaded as, 5011 - File uploaded, 5012 - File imported, 5013 - File copied, 5014 - File copied with overwriting, 5015 - File moved, 5016 - File moved with overwriting, 5017 - File moved to trash, 5018 - File deleted, 5019 - Folder created, 5020 - Folder renamed, 5021 - Folder updated access, 5022 - Folder copied, 5023 - Folder copied with overwriting, 5024 - Folder moved, 5025 - Folder moved with overwriting, 5026 - Folder moved to trash, 5027 - Folder deleted, 5028 - ThirdParty created, 5029 - ThirdParty updated, 5030 - ThirdParty deleted, 5031 - Documents ThirdParty settings updated, 5032 - Documents overwriting settings updated, 5033 - Documents uploading formats settings updated, 5034 - User file updated, 5035 - File converted, 5036 - File send access link, 5037 - Document service location setting, 5038 - Authorization keys setting, 5039 - Full text search setting, 5040 - Start transfer setting, 5041 - Backup started, 5042 - License key uploaded, 5043 - File change owner, 5044 - File restore version, 5045 - Document send to sign, 5046 - Document sign complete, 5047 - User updated email, 5048 - Documents store forcesave, 5049 - Documents forcesave, 5050 - Start storage encryption, 5051 - Privacy room enable, 5052 - Privacy room disable, 5053 - Start storage decryption, 5054 - File opened for change, 5055 - File marked as favorite, 5056 - File removed from favorite, 5057 - Folder downloaded, 5058 - File removed from list, 5059 - Folder removed from list, 5060 - File external link access updated, 5061 - Trash emptied, 5062 - File revision downloaded, 5063 - File marked as read, 5064 - File readed, 5065 - Folder marked as read, 5066 - Folder updated access for, 5068 - File updated access for, 5069 - Documents external share settings updated, 5070 - Room created, 5071 - Room renamed, 5072 - Room archived, 5073 - Room unarchived, 5074 - Room deleted, 5075 - Room update access for user, 5076 - Tag created, 5077 - Tags deleted, 5078 - Added room tags, 5079 - Deleted room tags, 5080 - Room logo created, 5081 - Room logo deleted, 5082 - Room invitation link updated, 5083 - Documents keep new file name settings updated, 5084 - Room remove user, 5085 - Room create user, 5086 - Room invitation link created, 5087 - Room invitation link deleted, 5088 - Room external link created, 5089 - Room external link updated, 5090 - Room external link deleted, 5091 - File external link created, 5092 - File external link updated, 5093 - File external link deleted, 5094 - Room group added, 5095 - Room update access for group, 5096 - Room group remove, 5097 - Room external link revoked, 5098 - Room external link renamed, 5099 - File uploaded with overwriting, 5100 - Room copied, 5101 - Documents display file extension updated, 5102 - Room color changed, 5103 - Room cover changed, 5104 - Room indexing changed, 5105 - Room deny download changed, 5106 - Room index export saved, 5107 - Folder index changed, 5108 - Folder index reordered, 5109 - Room deny download enabled, 5110 - Room deny download disabled, 5111 - File index changed, 5112 - Room watermark set, 5113 - Room watermark disabled, 5114 - Room index export saved, 5115 - Room indexing disabled, 5116 - Room life time set, 5117 - Room life time disabled, 5118 - Room invite resend, 5119 - File version deleted, 5120 - File custom filter enabled, 5121 - File custom filter disabled, 5122 - Folder external link created, 5123 - Folder external link updated, 5124 - Folder external link deleted, 5125 - Backup completed, 5126 - Backup failed, 5127 - Scheduled backup started, 5128 - Scheduled backup completed, 5129 - Scheduled backup failed, 5130 - Scheduled backup deleted, 5131 - Backup cancelled, 5132 - Restore started, 5133 - Restore cancelled, 5150 - Form started to fill, 5151 - Form partially filled, 5152 - Form completely filled, 5153 - Form stopped, 5154 - AI agent created, 5155 - AI agent renamed, 5156 - AI agent deleted, 5157 - MCP server added to AI agent, 5158 - MCP server deleted from AI agent, 5159 - Room change owner, 5160 - Documents default templates settings updated, 5201 - File saved, user quota exceeded, 5202 - File not saved due to user quota exceeded, 5203 - File saved, room quota exceeded, 5204 - File not saved due to room quota exceeded, 5501 - Ldap enabled, 5502 - Ldap disabled, 5503 - LDAP synchronization completed, 6000 - Language settings updated, 6001 - Time zone settings updated, 6002 - Dns settings updated, 6003 - Trusted mail domain settings updated, 6004 - Password strength settings updated, 6005 - Two factor authentication settings updated, 6006 - Administrator message settings updated, 6007 - Default start page settings updated, 6008 - Products list updated, 6009 - Administrator added, 6010 - Administrator opened full access, 6011 - Administrator deleted, 6012 - Users opened product access, 6013 - Groups opened product access, 6014 - Product access opened, 6015 - Product access restricted, 6016 - Product added administrator, 6017 - Product deleted administrator, 6018 - Greeting settings updated, 6019 - Team template changed, 6020 - Color theme changed, 6021 - Owner sent change owner instructions, 6022 - Owner updated, 6023 - Owner sent portal deactivation instructions, 6024 - Owner sent portal delete instructions, 6025 - Portal deactivated, 6026 - Portal deleted, 6027 - Login history report downloaded, 6028 - Audit trail report downloaded, 6029 - SSO enabled, 6030 - SSO disabled, 6031 - Portal access settings updated, 6032 - Cookie settings updated, 6033 - Mail service settings updated, 6034 - Custom navigation settings updated, 6035 - Audit settings updated, 6036 - Two factor authentication disabled, 6037 - Two factor authentication enabled by sms, 6038 - Two factor authentication enabled by tfa app, 6039 - Portal renamed, 6040 - Quota per room changed, 6041 - Quota per room disabled, 6042 - Quota per user changed, 6043 - Quota per user disabled, 6044 - Quota per portal changed, 6045 - Quota per portal disabled, 6046 - Form submit, 6047 - Form opened for filling, 6048 - Custom quota per room default, 6049 - Custom quota per room changed, 6050 - Custom quota per room disabled, 6051 - Custom quota per user default, 6052 - Custom quota per user changed, 6053 - Custom quota per user disabled, 6054 - DevTools access settings changed, 6055 - Webhook created, 6056 - Webhook updated, 6057 - Webhook deleted, 6058 - Created api key, 6059 - Update api key, 6060 - Deleted User api key, 6061 - Customer wallet topped up, 6062 - Customer operation performed, 6063 - Customer operations report downloaded, 6064 - Customer wallet top up settings updated, 6065 - Customer subscription updated, 6066 - Promotional banners visibility settings changed, 6067 - Customer wallet services settings updated, 6068 - Quota per AI agent changed, 6069 - Quota per AI agent disabled, 6070 - Custom quota per AI agent default, 6071 - Custom quota per AI agent changed, 6072 - Custom quota per AI agent disabled, 6073 - AI provider created, 6074 - AI provider updated, 6075 - AI provider deleted, 6076 - MCP server created, 6077 - MCP server updated, 6078 - MCP server enabled, 6079 - MCP server disabled, 6080 - MCP server deleted, 6081 - WebSearch settings configured, 6082 - WebSearch settings reset, 6083 - Vectorization settings configured, 6084 - Vectorization settings reset, 6085 - Webplugin uploaded, 6086 - Webplugin updated, 6087 - Webplugin deleted, 6088 - Whitelabel settings logo text updated, 6089 - Whitelabel settings logos updated, 6090 - Whitelabel company settings updated, 6091 - Whitelabel additional settings updated, 6092 - Whitelabel mail settings updated, 6093 - Invitation settings updated, 6094 - IP restrictions settings updated, 6095 - Login settings updated, 6096 - AI default provider set, 6097 - AI access enabled, 6098 - AI access disabled, 7000 - Contact admin mail sent, 7001 - Room invite link used, 7002 - User created and added to room, 7003 - Guest created and added to room, 7004 - Contact sales mail sent, 9901 - Create client, 9902 - Update client, 9903 - Regenerate secret, 9904 - Delete client, 9905 - Change client activation, 9906 - Change client visibility, 9907 - Revoke user client, 9908 - Generate authorization code token, 9909 - Generate personal access token, -1 - None]
type MessageAction int32

// List of MessageAction
const (
	MESSAGEACTION_LoginSuccess MessageAction = 1000
	MESSAGEACTION_LoginSuccessViaSocialAccount MessageAction = 1001
	MESSAGEACTION_LoginFailInvalidCombination MessageAction = 1002
	MESSAGEACTION_LoginFailSocialAccountNotFound MessageAction = 1003
	MESSAGEACTION_LoginFailDisabledProfile MessageAction = 1004
	MESSAGEACTION_LoginFail MessageAction = 1005
	MESSAGEACTION_Logout MessageAction = 1006
	MESSAGEACTION_LoginSuccessViaSms MessageAction = 1007
	MESSAGEACTION_LoginFailViaSms MessageAction = 1008
	MESSAGEACTION_LoginFailIpSecurity MessageAction = 1009
	MESSAGEACTION_LoginSuccessViaApi MessageAction = 1010
	MESSAGEACTION_LoginSuccessViaSocialApp MessageAction = 1011
	MESSAGEACTION_LoginSuccessViaApiSms MessageAction = 1012
	MESSAGEACTION_LoginFailViaApi MessageAction = 1013
	MESSAGEACTION_LoginFailViaApiSms MessageAction = 1014
	MESSAGEACTION_LoginSuccessViaSSO MessageAction = 1015
	MESSAGEACTION_SessionStarted MessageAction = 1016
	MESSAGEACTION_SessionCompleted MessageAction = 1017
	MESSAGEACTION_LoginFailViaSSO MessageAction = 1018
	MESSAGEACTION_LoginSuccessViaApiSocialAccount MessageAction = 1019
	MESSAGEACTION_LoginFailViaApiSocialAccount MessageAction = 1020
	MESSAGEACTION_LoginSuccesViaTfaApp MessageAction = 1021
	MESSAGEACTION_LoginFailViaTfaApp MessageAction = 1022
	MESSAGEACTION_LoginFailBruteForce MessageAction = 1023
	MESSAGEACTION_LoginSuccessViaApiTfa MessageAction = 1024
	MESSAGEACTION_LoginFailViaApiTfa MessageAction = 1025
	MESSAGEACTION_LoginFailRecaptcha MessageAction = 1026
	MESSAGEACTION_AuthLinkActivated MessageAction = 1027
	MESSAGEACTION_LoginSuccessViaOAuth MessageAction = 1028
	MESSAGEACTION_LoginSuccessViaPassword MessageAction = 1029
	MESSAGEACTION_UserCreated MessageAction = 4000
	MESSAGEACTION_GuestCreated MessageAction = 4001
	MESSAGEACTION_UserCreatedViaInvite MessageAction = 4002
	MESSAGEACTION_GuestCreatedViaInvite MessageAction = 4003
	MESSAGEACTION_UserActivated MessageAction = 4004
	MESSAGEACTION_GuestActivated MessageAction = 4005
	MESSAGEACTION_UserUpdated MessageAction = 4006
	MESSAGEACTION_UserUpdatedLanguage MessageAction = 4007
	MESSAGEACTION_UserAddedAvatar MessageAction = 4008
	MESSAGEACTION_UserDeletedAvatar MessageAction = 4009
	MESSAGEACTION_UserUpdatedAvatarThumbnails MessageAction = 4010
	MESSAGEACTION_UserLinkedSocialAccount MessageAction = 4011
	MESSAGEACTION_UserUnlinkedSocialAccount MessageAction = 4012
	MESSAGEACTION_UserSentActivationInstructions MessageAction = 4013
	MESSAGEACTION_UserSentEmailChangeInstructions MessageAction = 4014
	MESSAGEACTION_UserSentPasswordChangeInstructions MessageAction = 4015
	MESSAGEACTION_UserSentDeleteInstructions MessageAction = 4016
	MESSAGEACTION_UserUpdatedPassword MessageAction = 4017
	MESSAGEACTION_UserDeleted MessageAction = 4018
	MESSAGEACTION_UsersUpdatedType MessageAction = 4019
	MESSAGEACTION_UsersUpdatedStatus MessageAction = 4020
	MESSAGEACTION_UsersSentActivationInstructions MessageAction = 4021
	MESSAGEACTION_UsersDeleted MessageAction = 4022
	MESSAGEACTION_SentInviteInstructions MessageAction = 4023
	MESSAGEACTION_UserImported MessageAction = 4024
	MESSAGEACTION_GuestImported MessageAction = 4025
	MESSAGEACTION_GroupCreated MessageAction = 4026
	MESSAGEACTION_GroupUpdated MessageAction = 4027
	MESSAGEACTION_GroupDeleted MessageAction = 4028
	MESSAGEACTION_UserUpdatedMobileNumber MessageAction = 4029
	MESSAGEACTION_UserDataReassigns MessageAction = 4030
	MESSAGEACTION_UserDataRemoving MessageAction = 4031
	MESSAGEACTION_UserConnectedTfaApp MessageAction = 4032
	MESSAGEACTION_UserDisconnectedTfaApp MessageAction = 4033
	MESSAGEACTION_UserLogoutActiveConnections MessageAction = 4034
	MESSAGEACTION_UserLogoutActiveConnection MessageAction = 4035
	MESSAGEACTION_UserLogoutActiveConnectionsForUser MessageAction = 4036
	MESSAGEACTION_SendJoinInvite MessageAction = 4037
	MESSAGEACTION_FileCreated MessageAction = 5000
	MESSAGEACTION_FileRenamed MessageAction = 5001
	MESSAGEACTION_FileUpdated MessageAction = 5002
	MESSAGEACTION_FileCreatedVersion MessageAction = 5003
	MESSAGEACTION_FileDeletedVersion MessageAction = 5004
	MESSAGEACTION_FileUpdatedRevisionComment MessageAction = 5005
	MESSAGEACTION_FileLocked MessageAction = 5006
	MESSAGEACTION_FileUnlocked MessageAction = 5007
	MESSAGEACTION_FileUpdatedAccess MessageAction = 5008
	MESSAGEACTION_FileDownloaded MessageAction = 5009
	MESSAGEACTION_FileDownloadedAs MessageAction = 5010
	MESSAGEACTION_FileUploaded MessageAction = 5011
	MESSAGEACTION_FileImported MessageAction = 5012
	MESSAGEACTION_FileCopied MessageAction = 5013
	MESSAGEACTION_FileCopiedWithOverwriting MessageAction = 5014
	MESSAGEACTION_FileMoved MessageAction = 5015
	MESSAGEACTION_FileMovedWithOverwriting MessageAction = 5016
	MESSAGEACTION_FileMovedToTrash MessageAction = 5017
	MESSAGEACTION_FileDeleted MessageAction = 5018
	MESSAGEACTION_FolderCreated MessageAction = 5019
	MESSAGEACTION_FolderRenamed MessageAction = 5020
	MESSAGEACTION_FolderUpdatedAccess MessageAction = 5021
	MESSAGEACTION_FolderCopied MessageAction = 5022
	MESSAGEACTION_FolderCopiedWithOverwriting MessageAction = 5023
	MESSAGEACTION_FolderMoved MessageAction = 5024
	MESSAGEACTION_FolderMovedWithOverwriting MessageAction = 5025
	MESSAGEACTION_FolderMovedToTrash MessageAction = 5026
	MESSAGEACTION_FolderDeleted MessageAction = 5027
	MESSAGEACTION_ThirdPartyCreated MessageAction = 5028
	MESSAGEACTION_ThirdPartyUpdated MessageAction = 5029
	MESSAGEACTION_ThirdPartyDeleted MessageAction = 5030
	MESSAGEACTION_DocumentsThirdPartySettingsUpdated MessageAction = 5031
	MESSAGEACTION_DocumentsOverwritingSettingsUpdated MessageAction = 5032
	MESSAGEACTION_DocumentsUploadingFormatsSettingsUpdated MessageAction = 5033
	MESSAGEACTION_UserFileUpdated MessageAction = 5034
	MESSAGEACTION_FileConverted MessageAction = 5035
	MESSAGEACTION_FileSendAccessLink MessageAction = 5036
	MESSAGEACTION_DocumentServiceLocationSetting MessageAction = 5037
	MESSAGEACTION_AuthorizationKeysSetting MessageAction = 5038
	MESSAGEACTION_FullTextSearchSetting MessageAction = 5039
	MESSAGEACTION_StartTransferSetting MessageAction = 5040
	MESSAGEACTION_BackupStarted MessageAction = 5041
	MESSAGEACTION_LicenseKeyUploaded MessageAction = 5042
	MESSAGEACTION_FileChangeOwner MessageAction = 5043
	MESSAGEACTION_FileRestoreVersion MessageAction = 5044
	MESSAGEACTION_DocumentSendToSign MessageAction = 5045
	MESSAGEACTION_DocumentSignComplete MessageAction = 5046
	MESSAGEACTION_UserUpdatedEmail MessageAction = 5047
	MESSAGEACTION_DocumentsStoreForcesave MessageAction = 5048
	MESSAGEACTION_DocumentsForcesave MessageAction = 5049
	MESSAGEACTION_StartStorageEncryption MessageAction = 5050
	MESSAGEACTION_PrivacyRoomEnable MessageAction = 5051
	MESSAGEACTION_PrivacyRoomDisable MessageAction = 5052
	MESSAGEACTION_StartStorageDecryption MessageAction = 5053
	MESSAGEACTION_FileOpenedForChange MessageAction = 5054
	MESSAGEACTION_FileMarkedAsFavorite MessageAction = 5055
	MESSAGEACTION_FileRemovedFromFavorite MessageAction = 5056
	MESSAGEACTION_FolderDownloaded MessageAction = 5057
	MESSAGEACTION_FileRemovedFromList MessageAction = 5058
	MESSAGEACTION_FolderRemovedFromList MessageAction = 5059
	MESSAGEACTION_FileExternalLinkAccessUpdated MessageAction = 5060
	MESSAGEACTION_TrashEmptied MessageAction = 5061
	MESSAGEACTION_FileRevisionDownloaded MessageAction = 5062
	MESSAGEACTION_FileMarkedAsRead MessageAction = 5063
	MESSAGEACTION_FileReaded MessageAction = 5064
	MESSAGEACTION_FolderMarkedAsRead MessageAction = 5065
	MESSAGEACTION_FolderUpdatedAccessFor MessageAction = 5066
	MESSAGEACTION_FileUpdatedAccessFor MessageAction = 5068
	MESSAGEACTION_DocumentsExternalShareSettingsUpdated MessageAction = 5069
	MESSAGEACTION_RoomCreated MessageAction = 5070
	MESSAGEACTION_RoomRenamed MessageAction = 5071
	MESSAGEACTION_RoomArchived MessageAction = 5072
	MESSAGEACTION_RoomUnarchived MessageAction = 5073
	MESSAGEACTION_RoomDeleted MessageAction = 5074
	MESSAGEACTION_RoomUpdateAccessForUser MessageAction = 5075
	MESSAGEACTION_TagCreated MessageAction = 5076
	MESSAGEACTION_TagsDeleted MessageAction = 5077
	MESSAGEACTION_AddedRoomTags MessageAction = 5078
	MESSAGEACTION_DeletedRoomTags MessageAction = 5079
	MESSAGEACTION_RoomLogoCreated MessageAction = 5080
	MESSAGEACTION_RoomLogoDeleted MessageAction = 5081
	MESSAGEACTION_RoomInvitationLinkUpdated MessageAction = 5082
	MESSAGEACTION_DocumentsKeepNewFileNameSettingsUpdated MessageAction = 5083
	MESSAGEACTION_RoomRemoveUser MessageAction = 5084
	MESSAGEACTION_RoomCreateUser MessageAction = 5085
	MESSAGEACTION_RoomInvitationLinkCreated MessageAction = 5086
	MESSAGEACTION_RoomInvitationLinkDeleted MessageAction = 5087
	MESSAGEACTION_RoomExternalLinkCreated MessageAction = 5088
	MESSAGEACTION_RoomExternalLinkUpdated MessageAction = 5089
	MESSAGEACTION_RoomExternalLinkDeleted MessageAction = 5090
	MESSAGEACTION_FileExternalLinkCreated MessageAction = 5091
	MESSAGEACTION_FileExternalLinkUpdated MessageAction = 5092
	MESSAGEACTION_FileExternalLinkDeleted MessageAction = 5093
	MESSAGEACTION_RoomGroupAdded MessageAction = 5094
	MESSAGEACTION_RoomUpdateAccessForGroup MessageAction = 5095
	MESSAGEACTION_RoomGroupRemove MessageAction = 5096
	MESSAGEACTION_RoomExternalLinkRevoked MessageAction = 5097
	MESSAGEACTION_RoomExternalLinkRenamed MessageAction = 5098
	MESSAGEACTION_FileUploadedWithOverwriting MessageAction = 5099
	MESSAGEACTION_RoomCopied MessageAction = 5100
	MESSAGEACTION_DocumentsDisplayFileExtensionUpdated MessageAction = 5101
	MESSAGEACTION_RoomColorChanged MessageAction = 5102
	MESSAGEACTION_RoomCoverChanged MessageAction = 5103
	MESSAGEACTION_RoomIndexingChanged MessageAction = 5104
	MESSAGEACTION_RoomDenyDownloadChanged MessageAction = 5105
	MESSAGEACTION_RoomIndexExportSaved MessageAction = 5106
	MESSAGEACTION_FolderIndexChanged MessageAction = 5107
	MESSAGEACTION_FolderIndexReordered MessageAction = 5108
	MESSAGEACTION_RoomDenyDownloadEnabled MessageAction = 5109
	MESSAGEACTION_RoomDenyDownloadDisabled MessageAction = 5110
	MESSAGEACTION_FileIndexChanged MessageAction = 5111
	MESSAGEACTION_RoomWatermarkSet MessageAction = 5112
	MESSAGEACTION_RoomWatermarkDisabled MessageAction = 5113
	MESSAGEACTION_RoomIndexingEnabled MessageAction = 5114
	MESSAGEACTION_RoomIndexingDisabled MessageAction = 5115
	MESSAGEACTION_RoomLifeTimeSet MessageAction = 5116
	MESSAGEACTION_RoomLifeTimeDisabled MessageAction = 5117
	MESSAGEACTION_RoomInviteResend MessageAction = 5118
	MESSAGEACTION_FileVersionRemoved MessageAction = 5119
	MESSAGEACTION_FileCustomFilterEnabled MessageAction = 5120
	MESSAGEACTION_FileCustomFilterDisabled MessageAction = 5121
	MESSAGEACTION_FolderExternalLinkCreated MessageAction = 5122
	MESSAGEACTION_FolderExternalLinkUpdated MessageAction = 5123
	MESSAGEACTION_FolderExternalLinkDeleted MessageAction = 5124
	MESSAGEACTION_BackupCompleted MessageAction = 5125
	MESSAGEACTION_BackupFailed MessageAction = 5126
	MESSAGEACTION_ScheduledBackupStarted MessageAction = 5127
	MESSAGEACTION_ScheduledBackupCompleted MessageAction = 5128
	MESSAGEACTION_ScheduledBackupFailed MessageAction = 5129
	MESSAGEACTION_ScheduledBackupDeleted MessageAction = 5130
	MESSAGEACTION_BackupCancelled MessageAction = 5131
	MESSAGEACTION_RestoreStarted MessageAction = 5132
	MESSAGEACTION_RestoreCancelled MessageAction = 5133
	MESSAGEACTION_FormStartedToFill MessageAction = 5150
	MESSAGEACTION_FormPartiallyFilled MessageAction = 5151
	MESSAGEACTION_FormCompletelyFilled MessageAction = 5152
	MESSAGEACTION_FormStopped MessageAction = 5153
	MESSAGEACTION_AgentCreated MessageAction = 5154
	MESSAGEACTION_AgentRenamed MessageAction = 5155
	MESSAGEACTION_AgentDeleted MessageAction = 5156
	MESSAGEACTION_AddedServerToAgent MessageAction = 5157
	MESSAGEACTION_DeletedServerFromAgent MessageAction = 5158
	MESSAGEACTION_RoomChangeOwner MessageAction = 5159
	MESSAGEACTION_DocumentsDefaultTemplatesSettingsUpdated MessageAction = 5160
	MESSAGEACTION_FileSavedButUserQuotaExceeded MessageAction = 5201
	MESSAGEACTION_FileNotSavedDueToUserQuota MessageAction = 5202
	MESSAGEACTION_FileSavedButRoomQuotaExceeded MessageAction = 5203
	MESSAGEACTION_FileNotSavedDueToRoomQuota MessageAction = 5204
	MESSAGEACTION_LdapEnabled MessageAction = 5501
	MESSAGEACTION_LdapDisabled MessageAction = 5502
	MESSAGEACTION_LdapSync MessageAction = 5503
	MESSAGEACTION_LanguageSettingsUpdated MessageAction = 6000
	MESSAGEACTION_TimeZoneSettingsUpdated MessageAction = 6001
	MESSAGEACTION_DnsSettingsUpdated MessageAction = 6002
	MESSAGEACTION_TrustedMailDomainSettingsUpdated MessageAction = 6003
	MESSAGEACTION_PasswordStrengthSettingsUpdated MessageAction = 6004
	MESSAGEACTION_TwoFactorAuthenticationSettingsUpdated MessageAction = 6005
	MESSAGEACTION_AdministratorMessageSettingsUpdated MessageAction = 6006
	MESSAGEACTION_DefaultStartPageSettingsUpdated MessageAction = 6007
	MESSAGEACTION_ProductsListUpdated MessageAction = 6008
	MESSAGEACTION_AdministratorAdded MessageAction = 6009
	MESSAGEACTION_AdministratorOpenedFullAccess MessageAction = 6010
	MESSAGEACTION_AdministratorDeleted MessageAction = 6011
	MESSAGEACTION_UsersOpenedProductAccess MessageAction = 6012
	MESSAGEACTION_GroupsOpenedProductAccess MessageAction = 6013
	MESSAGEACTION_ProductAccessOpened MessageAction = 6014
	MESSAGEACTION_ProductAccessRestricted MessageAction = 6015
	MESSAGEACTION_ProductAddedAdministrator MessageAction = 6016
	MESSAGEACTION_ProductDeletedAdministrator MessageAction = 6017
	MESSAGEACTION_GreetingSettingsUpdated MessageAction = 6018
	MESSAGEACTION_TeamTemplateChanged MessageAction = 6019
	MESSAGEACTION_ColorThemeChanged MessageAction = 6020
	MESSAGEACTION_OwnerSentChangeOwnerInstructions MessageAction = 6021
	MESSAGEACTION_OwnerUpdated MessageAction = 6022
	MESSAGEACTION_OwnerSentPortalDeactivationInstructions MessageAction = 6023
	MESSAGEACTION_OwnerSentPortalDeleteInstructions MessageAction = 6024
	MESSAGEACTION_PortalDeactivated MessageAction = 6025
	MESSAGEACTION_PortalDeleted MessageAction = 6026
	MESSAGEACTION_LoginHistoryReportDownloaded MessageAction = 6027
	MESSAGEACTION_AuditTrailReportDownloaded MessageAction = 6028
	MESSAGEACTION_SSOEnabled MessageAction = 6029
	MESSAGEACTION_SSODisabled MessageAction = 6030
	MESSAGEACTION_PortalAccessSettingsUpdated MessageAction = 6031
	MESSAGEACTION_CookieSettingsUpdated MessageAction = 6032
	MESSAGEACTION_MailServiceSettingsUpdated MessageAction = 6033
	MESSAGEACTION_CustomNavigationSettingsUpdated MessageAction = 6034
	MESSAGEACTION_AuditSettingsUpdated MessageAction = 6035
	MESSAGEACTION_TwoFactorAuthenticationDisabled MessageAction = 6036
	MESSAGEACTION_TwoFactorAuthenticationEnabledBySms MessageAction = 6037
	MESSAGEACTION_TwoFactorAuthenticationEnabledByTfaApp MessageAction = 6038
	MESSAGEACTION_PortalRenamed MessageAction = 6039
	MESSAGEACTION_QuotaPerRoomChanged MessageAction = 6040
	MESSAGEACTION_QuotaPerRoomDisabled MessageAction = 6041
	MESSAGEACTION_QuotaPerUserChanged MessageAction = 6042
	MESSAGEACTION_QuotaPerUserDisabled MessageAction = 6043
	MESSAGEACTION_QuotaPerPortalChanged MessageAction = 6044
	MESSAGEACTION_QuotaPerPortalDisabled MessageAction = 6045
	MESSAGEACTION_FormSubmit MessageAction = 6046
	MESSAGEACTION_FormOpenedForFilling MessageAction = 6047
	MESSAGEACTION_CustomQuotaPerRoomDefault MessageAction = 6048
	MESSAGEACTION_CustomQuotaPerRoomChanged MessageAction = 6049
	MESSAGEACTION_CustomQuotaPerRoomDisabled MessageAction = 6050
	MESSAGEACTION_CustomQuotaPerUserDefault MessageAction = 6051
	MESSAGEACTION_CustomQuotaPerUserChanged MessageAction = 6052
	MESSAGEACTION_CustomQuotaPerUserDisabled MessageAction = 6053
	MESSAGEACTION_DevToolsAccessSettingsChanged MessageAction = 6054
	MESSAGEACTION_WebhookCreated MessageAction = 6055
	MESSAGEACTION_WebhookUpdated MessageAction = 6056
	MESSAGEACTION_WebhookDeleted MessageAction = 6057
	MESSAGEACTION_ApiKeyCreated MessageAction = 6058
	MESSAGEACTION_ApiKeyUpdated MessageAction = 6059
	MESSAGEACTION_ApiKeyDeleted MessageAction = 6060
	MESSAGEACTION_CustomerWalletToppedUp MessageAction = 6061
	MESSAGEACTION_CustomerOperationPerformed MessageAction = 6062
	MESSAGEACTION_CustomerOperationsReportDownloaded MessageAction = 6063
	MESSAGEACTION_CustomerWalletTopUpSettingsUpdated MessageAction = 6064
	MESSAGEACTION_CustomerSubscriptionUpdated MessageAction = 6065
	MESSAGEACTION_BannerSettingsChanged MessageAction = 6066
	MESSAGEACTION_CustomerWalletServicesSettingsUpdated MessageAction = 6067
	MESSAGEACTION_QuotaPerAiAgentChanged MessageAction = 6068
	MESSAGEACTION_QuotaPerAiAgentDisabled MessageAction = 6069
	MESSAGEACTION_CustomQuotaPerAiAgentDefault MessageAction = 6070
	MESSAGEACTION_CustomQuotaPerAiAgentChanged MessageAction = 6071
	MESSAGEACTION_CustomQuotaPerAiAgentDisabled MessageAction = 6072
	MESSAGEACTION_AIProviderCreated MessageAction = 6073
	MESSAGEACTION_AIProviderUpdated MessageAction = 6074
	MESSAGEACTION_AIProviderDeleted MessageAction = 6075
	MESSAGEACTION_ServerCreated MessageAction = 6076
	MESSAGEACTION_ServerUpdated MessageAction = 6077
	MESSAGEACTION_ServerEnabled MessageAction = 6078
	MESSAGEACTION_ServerDisabled MessageAction = 6079
	MESSAGEACTION_ServerDeleted MessageAction = 6080
	MESSAGEACTION_SetWebSearchSettings MessageAction = 6081
	MESSAGEACTION_ResetWebSearchSettings MessageAction = 6082
	MESSAGEACTION_SetVectorizationSettings MessageAction = 6083
	MESSAGEACTION_ResetVectorizationSettings MessageAction = 6084
	MESSAGEACTION_WebpluginUploaded MessageAction = 6085
	MESSAGEACTION_WebpluginUpdated MessageAction = 6086
	MESSAGEACTION_WebpluginDeleted MessageAction = 6087
	MESSAGEACTION_WhiteLabelSettingsLogoTextUpdated MessageAction = 6088
	MESSAGEACTION_WhiteLabelSettingsLogosUpdated MessageAction = 6089
	MESSAGEACTION_WhiteLabelCompanySettingsUpdated MessageAction = 6090
	MESSAGEACTION_WhiteLabelAdditionalSettingsUpdated MessageAction = 6091
	MESSAGEACTION_WhiteLabelMailSettingsUpdated MessageAction = 6092
	MESSAGEACTION_InvitationSettingsUpdated MessageAction = 6093
	MESSAGEACTION_IPRestrictionsSettingsUpdated MessageAction = 6094
	MESSAGEACTION_LoginSettingsUpdated MessageAction = 6095
	MESSAGEACTION_AIDefaultProviderSet MessageAction = 6096
	MESSAGEACTION_AIAccessEnabled MessageAction = 6097
	MESSAGEACTION_AIAccessDisabled MessageAction = 6098
	MESSAGEACTION_ContactAdminMailSent MessageAction = 7000
	MESSAGEACTION_RoomInviteLinkUsed MessageAction = 7001
	MESSAGEACTION_UserCreatedAndAddedToRoom MessageAction = 7002
	MESSAGEACTION_GuestCreatedAndAddedToRoom MessageAction = 7003
	MESSAGEACTION_ContactSalesMailSent MessageAction = 7004
	MESSAGEACTION_CreateClient MessageAction = 9901
	MESSAGEACTION_UpdateClient MessageAction = 9902
	MESSAGEACTION_RegenerateSecret MessageAction = 9903
	MESSAGEACTION_DeleteClient MessageAction = 9904
	MESSAGEACTION_ChangeClientActivation MessageAction = 9905
	MESSAGEACTION_ChangeClientVisibility MessageAction = 9906
	MESSAGEACTION_RevokeUserClient MessageAction = 9907
	MESSAGEACTION_GenerateAuthorizationCodeToken MessageAction = 9908
	MESSAGEACTION_GeneratePersonalAccessToken MessageAction = 9909
	MESSAGEACTION_None MessageAction = -1
)

// All allowed values of MessageAction enum
var AllowedMessageActionEnumValues = []MessageAction{
	1000,
	1001,
	1002,
	1003,
	1004,
	1005,
	1006,
	1007,
	1008,
	1009,
	1010,
	1011,
	1012,
	1013,
	1014,
	1015,
	1016,
	1017,
	1018,
	1019,
	1020,
	1021,
	1022,
	1023,
	1024,
	1025,
	1026,
	1027,
	1028,
	1029,
	4000,
	4001,
	4002,
	4003,
	4004,
	4005,
	4006,
	4007,
	4008,
	4009,
	4010,
	4011,
	4012,
	4013,
	4014,
	4015,
	4016,
	4017,
	4018,
	4019,
	4020,
	4021,
	4022,
	4023,
	4024,
	4025,
	4026,
	4027,
	4028,
	4029,
	4030,
	4031,
	4032,
	4033,
	4034,
	4035,
	4036,
	4037,
	5000,
	5001,
	5002,
	5003,
	5004,
	5005,
	5006,
	5007,
	5008,
	5009,
	5010,
	5011,
	5012,
	5013,
	5014,
	5015,
	5016,
	5017,
	5018,
	5019,
	5020,
	5021,
	5022,
	5023,
	5024,
	5025,
	5026,
	5027,
	5028,
	5029,
	5030,
	5031,
	5032,
	5033,
	5034,
	5035,
	5036,
	5037,
	5038,
	5039,
	5040,
	5041,
	5042,
	5043,
	5044,
	5045,
	5046,
	5047,
	5048,
	5049,
	5050,
	5051,
	5052,
	5053,
	5054,
	5055,
	5056,
	5057,
	5058,
	5059,
	5060,
	5061,
	5062,
	5063,
	5064,
	5065,
	5066,
	5068,
	5069,
	5070,
	5071,
	5072,
	5073,
	5074,
	5075,
	5076,
	5077,
	5078,
	5079,
	5080,
	5081,
	5082,
	5083,
	5084,
	5085,
	5086,
	5087,
	5088,
	5089,
	5090,
	5091,
	5092,
	5093,
	5094,
	5095,
	5096,
	5097,
	5098,
	5099,
	5100,
	5101,
	5102,
	5103,
	5104,
	5105,
	5106,
	5107,
	5108,
	5109,
	5110,
	5111,
	5112,
	5113,
	5114,
	5115,
	5116,
	5117,
	5118,
	5119,
	5120,
	5121,
	5122,
	5123,
	5124,
	5125,
	5126,
	5127,
	5128,
	5129,
	5130,
	5131,
	5132,
	5133,
	5150,
	5151,
	5152,
	5153,
	5154,
	5155,
	5156,
	5157,
	5158,
	5159,
	5160,
	5201,
	5202,
	5203,
	5204,
	5501,
	5502,
	5503,
	6000,
	6001,
	6002,
	6003,
	6004,
	6005,
	6006,
	6007,
	6008,
	6009,
	6010,
	6011,
	6012,
	6013,
	6014,
	6015,
	6016,
	6017,
	6018,
	6019,
	6020,
	6021,
	6022,
	6023,
	6024,
	6025,
	6026,
	6027,
	6028,
	6029,
	6030,
	6031,
	6032,
	6033,
	6034,
	6035,
	6036,
	6037,
	6038,
	6039,
	6040,
	6041,
	6042,
	6043,
	6044,
	6045,
	6046,
	6047,
	6048,
	6049,
	6050,
	6051,
	6052,
	6053,
	6054,
	6055,
	6056,
	6057,
	6058,
	6059,
	6060,
	6061,
	6062,
	6063,
	6064,
	6065,
	6066,
	6067,
	6068,
	6069,
	6070,
	6071,
	6072,
	6073,
	6074,
	6075,
	6076,
	6077,
	6078,
	6079,
	6080,
	6081,
	6082,
	6083,
	6084,
	6085,
	6086,
	6087,
	6088,
	6089,
	6090,
	6091,
	6092,
	6093,
	6094,
	6095,
	6096,
	6097,
	6098,
	7000,
	7001,
	7002,
	7003,
	7004,
	9901,
	9902,
	9903,
	9904,
	9905,
	9906,
	9907,
	9908,
	9909,
	-1,
}

func (v *MessageAction) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageAction(value)
	for _, existing := range AllowedMessageActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageAction", value)
}

// NewMessageActionFromValue returns a pointer to a valid MessageAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageActionFromValue(v int32) (*MessageAction, error) {
	ev := MessageAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageAction: valid values are %v", v, AllowedMessageActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageAction) IsValid() bool {
	for _, existing := range AllowedMessageActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageAction value
func (v MessageAction) Ptr() *MessageAction {
	return &v
}

type NullableMessageAction struct {
	value *MessageAction
	isSet bool
}

func (v NullableMessageAction) Get() *MessageAction {
	return v.value
}

func (v *NullableMessageAction) Set(val *MessageAction) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAction) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAction(val *MessageAction) *NullableMessageAction {
	return &NullableMessageAction{value: val, isSet: true}
}

func (v NullableMessageAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

