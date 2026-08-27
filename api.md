# Cadenya Go SDK reference

Every call takes a `context.Context` and returns an error; see README.md for usage patterns.

## accounts

Retrieves the current account for the token accessing the API

```go
client.Accounts().Retrieve(ctx context.Context, opts ...RequestOption) (*Account, error)
```
Rotates the challenge token for the account

```go
client.Accounts().RotateChallengeToken(ctx context.Context, opts ...RequestOption) (*RotateChallengeTokenResponse, error)
```
Rotates the webhook signing key for the account

```go
client.Accounts().RotateWebhookSigningKey(ctx context.Context, opts ...RequestOption) (*RotateWebhookEventsHMACSecretResponse, error)
```

## api_keys

Get the global API key

```go
client.APIKeys().RetrieveGlobal(ctx context.Context, opts ...RequestOption) (*APIKey, error)
```
Disable the global API key

```go
client.APIKeys().DisableGlobal(ctx context.Context, opts ...RequestOption) (*APIKey, error)
```
Enable the global API key

```go
client.APIKeys().EnableGlobal(ctx context.Context, opts ...RequestOption) (*APIKey, error)
```
Rotate the global API key

```go
client.APIKeys().RotateGlobal(ctx context.Context, opts ...RequestOption) (*APIKey, error)
```
List API keys

```go
client.APIKeys().List(ctx context.Context, params *APIKeyListParams, opts ...RequestOption) (*Page[APIKey], error)
```
Create a new API key

```go
client.APIKeys().Create(ctx context.Context, params *APIKeyCreateParams, opts ...RequestOption) (*APIKey, error)
```
Get an API key by ID

```go
client.APIKeys().Retrieve(ctx context.Context, id string, params *APIKeyRetrieveParams, opts ...RequestOption) (*APIKey, error)
```
Delete an API key

```go
client.APIKeys().Delete(ctx context.Context, id string, params *APIKeyDeleteParams, opts ...RequestOption) error
```
Update an API key

```go
client.APIKeys().Update(ctx context.Context, id string, params *APIKeyUpdateParams, opts ...RequestOption) (*APIKey, error)
```
Disable an API key

```go
client.APIKeys().Disable(ctx context.Context, id string, params *APIKeyDisableParams, opts ...RequestOption) (*APIKey, error)
```
Enable an API key

```go
client.APIKeys().Enable(ctx context.Context, id string, params *APIKeyEnableParams, opts ...RequestOption) (*APIKey, error)
```
Rotate an API key

```go
client.APIKeys().Rotate(ctx context.Context, id string, params *APIKeyRotateParams, opts ...RequestOption) (*APIKey, error)
```

## workspace_admin

Search account profiles

```go
client.WorkspaceAdmin().ListProfiles(ctx context.Context, params *WorkspaceAdminListProfilesParams, opts ...RequestOption) (*Page[Profile], error)
```
List all workspaces in the account

```go
client.WorkspaceAdmin().ListAccount(ctx context.Context, params *WorkspaceAdminListAccountParams, opts ...RequestOption) (*Page[Workspace], error)
```
Create a workspace

```go
client.WorkspaceAdmin().Create(ctx context.Context, params *WorkspaceAdminCreateParams, opts ...RequestOption) (*Workspace, error)
```
Get a workspace by ID

```go
client.WorkspaceAdmin().Retrieve(ctx context.Context, params *WorkspaceAdminRetrieveParams, opts ...RequestOption) (*Workspace, error)
```
Archive a workspace

```go
client.WorkspaceAdmin().Archive(ctx context.Context, params *WorkspaceAdminArchiveParams, opts ...RequestOption) error
```
Update a workspace

```go
client.WorkspaceAdmin().Update(ctx context.Context, params *WorkspaceAdminUpdateParams, opts ...RequestOption) (*Workspace, error)
```
List workspace members

```go
client.WorkspaceAdmin().ListMembers(ctx context.Context, params *WorkspaceAdminListMembersParams, opts ...RequestOption) (*Page[WorkspaceMember], error)
```
Add a member to a workspace

```go
client.WorkspaceAdmin().AddMember(ctx context.Context, params *WorkspaceAdminAddMemberParams, opts ...RequestOption) (*WorkspaceMember, error)
```
Remove a member from a workspace

```go
client.WorkspaceAdmin().RemoveMember(ctx context.Context, profileID string, params *WorkspaceAdminRemoveMemberParams, opts ...RequestOption) error
```

## profiles

Retrieves the profile for the credentials accessing the API

```go
client.Profiles().Whoami(ctx context.Context, opts ...RequestOption) (*Profile, error)
```

## workspaces

List workspaces

```go
client.Workspaces().List(ctx context.Context, params *WorkspaceListParams, opts ...RequestOption) (*Page[Workspace], error)
```

## agents

List agents

```go
client.Agents().List(ctx context.Context, params *AgentListParams, opts ...RequestOption) (*Page[Agent], error)
```
Create a new agent

```go
client.Agents().Create(ctx context.Context, params *AgentCreateParams, opts ...RequestOption) (*Agent, error)
```
List feedback for an agent

```go
client.Agents().ListFeedback(ctx context.Context, agentID string, params *AgentListFeedbackParams, opts ...RequestOption) (*Page[ObjectiveFeedback], error)
```
List webhook deliveries

```go
client.Agents().ListWebhookDeliveries(ctx context.Context, agentID string, params *AgentListWebhookDeliveriesParams, opts ...RequestOption) (*Page[WebhookDelivery], error)
```
Get an agent by ID

```go
client.Agents().Retrieve(ctx context.Context, id string, params *AgentRetrieveParams, opts ...RequestOption) (*Agent, error)
```
Delete an agent

```go
client.Agents().Delete(ctx context.Context, id string, params *AgentDeleteParams, opts ...RequestOption) error
```
Update an agent

```go
client.Agents().Update(ctx context.Context, id string, params *AgentUpdateParams, opts ...RequestOption) (*Agent, error)
```
Archive an agent

```go
client.Agents().Archive(ctx context.Context, id string, params *AgentArchiveParams, opts ...RequestOption) (*Agent, error)
```
Publish an agent

```go
client.Agents().Publish(ctx context.Context, id string, params *AgentPublishParams, opts ...RequestOption) (*Agent, error)
```
Unarchive an agent

```go
client.Agents().Unarchive(ctx context.Context, id string, params *AgentUnarchiveParams, opts ...RequestOption) (*Agent, error)
```
Unpublish an agent

```go
client.Agents().Unpublish(ctx context.Context, id string, params *AgentUnpublishParams, opts ...RequestOption) (*Agent, error)
```

## agents.schedules

List schedules

```go
client.Agents().Schedules().List(ctx context.Context, agentID string, params *AgentScheduleListParams, opts ...RequestOption) (*Page[AgentSchedule], error)
```
Create a new schedule

```go
client.Agents().Schedules().Create(ctx context.Context, agentID string, params *AgentScheduleCreateParams, opts ...RequestOption) (*AgentSchedule, error)
```
Get a schedule by ID

```go
client.Agents().Schedules().Retrieve(ctx context.Context, agentID string, id string, params *AgentScheduleRetrieveParams, opts ...RequestOption) (*AgentSchedule, error)
```
Delete a schedule

```go
client.Agents().Schedules().Delete(ctx context.Context, agentID string, id string, params *AgentScheduleDeleteParams, opts ...RequestOption) error
```
Update a schedule

```go
client.Agents().Schedules().Update(ctx context.Context, agentID string, id string, params *AgentScheduleUpdateParams, opts ...RequestOption) (*AgentSchedule, error)
```
Archive a schedule

```go
client.Agents().Schedules().Archive(ctx context.Context, agentID string, id string, params *AgentScheduleArchiveParams, opts ...RequestOption) (*AgentSchedule, error)
```
Pause a schedule

```go
client.Agents().Schedules().Pause(ctx context.Context, agentID string, id string, params *AgentSchedulePauseParams, opts ...RequestOption) (*AgentSchedule, error)
```
Resume a schedule

```go
client.Agents().Schedules().Resume(ctx context.Context, agentID string, id string, params *AgentScheduleResumeParams, opts ...RequestOption) (*AgentSchedule, error)
```

## agents.variations

List variations

```go
client.Agents().Variations().List(ctx context.Context, agentID string, params *AgentVariationListParams, opts ...RequestOption) (*Page[AgentVariation], error)
```
Create a new variation

```go
client.Agents().Variations().Create(ctx context.Context, agentID string, params *AgentVariationCreateParams, opts ...RequestOption) (*AgentVariation, error)
```
Get a variation by ID

```go
client.Agents().Variations().Retrieve(ctx context.Context, agentID string, id string, params *AgentVariationRetrieveParams, opts ...RequestOption) (*AgentVariation, error)
```
Delete a variation

```go
client.Agents().Variations().Delete(ctx context.Context, agentID string, id string, params *AgentVariationDeleteParams, opts ...RequestOption) error
```
Update a variation

```go
client.Agents().Variations().Update(ctx context.Context, agentID string, id string, params *AgentVariationUpdateParams, opts ...RequestOption) (*AgentVariation, error)
```
Add an assignment to a variation

```go
client.Agents().Variations().AddAssignment(ctx context.Context, agentID string, variationID string, params *AgentVariationAddAssignmentParams, opts ...RequestOption) (*VariationAssignment, error)
```
Remove an assignment from a variation

```go
client.Agents().Variations().RemoveAssignment(ctx context.Context, agentID string, variationID string, id string, params *AgentVariationRemoveAssignmentParams, opts ...RequestOption) error
```
Attach a memory layer to a variation

```go
client.Agents().Variations().AddMemoryLayer(ctx context.Context, agentID string, variationID string, params *AgentVariationAddMemoryLayerParams, opts ...RequestOption) (*VariationMemoryLayerAssignment, error)
```
Remove a memory layer assignment from a variation

```go
client.Agents().Variations().RemoveMemoryLayer(ctx context.Context, agentID string, variationID string, id string, params *AgentVariationRemoveMemoryLayerParams, opts ...RequestOption) error
```
Update a variation's memory layer assignment

```go
client.Agents().Variations().UpdateMemoryLayer(ctx context.Context, agentID string, variationID string, id string, params *AgentVariationUpdateMemoryLayerParams, opts ...RequestOption) (*VariationMemoryLayerAssignment, error)
```

## ai_provider_keys

List AI provider keys

```go
client.AIProviderKeys().List(ctx context.Context, params *AIProviderKeyListParams, opts ...RequestOption) (*Page[AIProviderKey], error)
```
Create a new AI provider key

```go
client.AIProviderKeys().Create(ctx context.Context, params *AIProviderKeyCreateParams, opts ...RequestOption) (*AIProviderKey, error)
```
Get an AI provider key by ID

```go
client.AIProviderKeys().Retrieve(ctx context.Context, id string, params *AIProviderKeyRetrieveParams, opts ...RequestOption) (*AIProviderKey, error)
```
Delete an AI provider key

```go
client.AIProviderKeys().Delete(ctx context.Context, id string, params *AIProviderKeyDeleteParams, opts ...RequestOption) error
```
Update an AI provider key

```go
client.AIProviderKeys().Update(ctx context.Context, id string, params *AIProviderKeyUpdateParams, opts ...RequestOption) (*AIProviderKey, error)
```

## memory_layers

List memory layers

```go
client.MemoryLayers().List(ctx context.Context, params *MemoryLayerListParams, opts ...RequestOption) (*Page[MemoryLayer], error)
```
Create a new memory layer

```go
client.MemoryLayers().Create(ctx context.Context, params *MemoryLayerCreateParams, opts ...RequestOption) (*MemoryLayer, error)
```
Get a memory layer by ID

```go
client.MemoryLayers().Retrieve(ctx context.Context, id string, params *MemoryLayerRetrieveParams, opts ...RequestOption) (*MemoryLayer, error)
```
Delete a memory layer

```go
client.MemoryLayers().Delete(ctx context.Context, id string, params *MemoryLayerDeleteParams, opts ...RequestOption) error
```
Update a memory layer

```go
client.MemoryLayers().Update(ctx context.Context, id string, params *MemoryLayerUpdateParams, opts ...RequestOption) (*MemoryLayer, error)
```

## memory_layers.entries

List memory entries

```go
client.MemoryLayers().Entries().List(ctx context.Context, memoryLayerID string, params *MemoryEntryListParams, opts ...RequestOption) (*Page[MemoryEntry], error)
```
Create a new memory entry

```go
client.MemoryLayers().Entries().Create(ctx context.Context, memoryLayerID string, params *MemoryEntryCreateParams, opts ...RequestOption) (*MemoryEntryDetail, error)
```
Get a memory entry by ID

```go
client.MemoryLayers().Entries().Retrieve(ctx context.Context, memoryLayerID string, id string, params *MemoryEntryRetrieveParams, opts ...RequestOption) (*MemoryEntryDetail, error)
```
Delete a memory entry

```go
client.MemoryLayers().Entries().Delete(ctx context.Context, memoryLayerID string, id string, params *MemoryEntryDeleteParams, opts ...RequestOption) error
```
Update a memory entry

```go
client.MemoryLayers().Entries().Update(ctx context.Context, memoryLayerID string, id string, params *MemoryEntryUpdateParams, opts ...RequestOption) (*MemoryEntryDetail, error)
```

## models

List models

```go
client.Models().List(ctx context.Context, params *ModelListParams, opts ...RequestOption) (*Page[Model], error)
```
Get a model by ID

```go
client.Models().Retrieve(ctx context.Context, id string, params *ModelRetrieveParams, opts ...RequestOption) (*Model, error)
```
Disable a model

```go
client.Models().Disable(ctx context.Context, id string, params *ModelDisableParams, opts ...RequestOption) (*Model, error)
```
Enable a model

```go
client.Models().Enable(ctx context.Context, id string, params *ModelEnableParams, opts ...RequestOption) (*Model, error)
```
Swap models on agent variations

```go
client.Models().SwapOnVariations(ctx context.Context, params *ModelSwapOnVariationsParams, opts ...RequestOption) error
```

## objectives

List objectives

```go
client.Objectives().List(ctx context.Context, params *ObjectiveListParams, opts ...RequestOption) (*Page[Objective], error)
```
Create a new objective

```go
client.Objectives().Create(ctx context.Context, params *ObjectiveCreateParams, opts ...RequestOption) (*Objective, error)
```
Get an objective by ID

```go
client.Objectives().Retrieve(ctx context.Context, id string, params *ObjectiveRetrieveParams, opts ...RequestOption) (*Objective, error)
```
List objective context windows

```go
client.Objectives().ListContextWindows(ctx context.Context, objectiveID string, params *ObjectiveListContextWindowsParams, opts ...RequestOption) (*Page[ObjectiveContextWindow], error)
```
Get objective context usage

```go
client.Objectives().RetrieveDiagnostics(ctx context.Context, objectiveID string, params *ObjectiveRetrieveDiagnosticsParams, opts ...RequestOption) (*GetObjectiveDiagnosticsResponse, error)
```
List objective events

```go
client.Objectives().ListEvents(ctx context.Context, objectiveID string, params *ObjectiveListEventsParams, opts ...RequestOption) (*Page[ObjectiveEvent], error)
```
Stream objective events

```go
client.Objectives().StreamEvents(ctx context.Context, objectiveID string, params *ObjectiveStreamEventsParams, opts ...RequestOption) (*Stream[ObjectiveEvent], error)
```
List feedback for an objective

```go
client.Objectives().ListFeedback(ctx context.Context, objectiveID string, params *ObjectiveListFeedbackParams, opts ...RequestOption) (*Page[ObjectiveFeedback], error)
```
Submit feedback for an objective

```go
client.Objectives().CreateFeedback(ctx context.Context, objectiveID string, params *ObjectiveCreateFeedbackParams, opts ...RequestOption) (*ObjectiveFeedback, error)
```
List objective tasks

```go
client.Objectives().ListTasks(ctx context.Context, objectiveID string, params *ObjectiveListTasksParams, opts ...RequestOption) (*Page[ObjectiveTask], error)
```
Get an objective task by ID

```go
client.Objectives().RetrieveTask(ctx context.Context, objectiveID string, id string, params *ObjectiveRetrieveTaskParams, opts ...RequestOption) (*ObjectiveTask, error)
```
List objective tool calls

```go
client.Objectives().ListToolCalls(ctx context.Context, objectiveID string, params *ObjectiveListToolCallsParams, opts ...RequestOption) (*Page[ObjectiveToolCall], error)
```
Get an objective tool call by ID

```go
client.Objectives().RetrieveToolCall(ctx context.Context, objectiveID string, toolCallID string, params *ObjectiveRetrieveToolCallParams, opts ...RequestOption) (*ObjectiveToolCallWithResult, error)
```
Approve a tool call

```go
client.Objectives().ApproveToolCall(ctx context.Context, objectiveID string, toolCallID string, params *ObjectiveApproveToolCallParams, opts ...RequestOption) (*ObjectiveToolCall, error)
```
Deny a tool call

```go
client.Objectives().DenyToolCall(ctx context.Context, objectiveID string, toolCallID string, params *ObjectiveDenyToolCallParams, opts ...RequestOption) (*ObjectiveToolCall, error)
```
Set a bare tool call's content

```go
client.Objectives().SetToolCallContent(ctx context.Context, objectiveID string, toolCallID string, params *ObjectiveSetToolCallContentParams, opts ...RequestOption) (*ObjectiveToolCall, error)
```
List objective tools

```go
client.Objectives().ListTools(ctx context.Context, objectiveID string, params *ObjectiveListToolsParams, opts ...RequestOption) (*Page[ObjectiveTool], error)
```
Cancel an objective

```go
client.Objectives().Cancel(ctx context.Context, objectiveID string, params *ObjectiveCancelParams, opts ...RequestOption) (*Objective, error)
```
Compact an objective

```go
client.Objectives().Compact(ctx context.Context, objectiveID string, params *ObjectiveCompactParams, opts ...RequestOption) (*CompactObjectiveResponse, error)
```
Continue an objective

```go
client.Objectives().Continue(ctx context.Context, objectiveID string, params *ObjectiveContinueParams, opts ...RequestOption) (*ObjectiveEvent, error)
```

## tool_search

Search for tools or tool sets

```go
client.ToolSearch().SearchOrSets(ctx context.Context, params *ToolSearchSearchOrSetsParams, opts ...RequestOption) (*SearchToolsOrToolSetsResponse, error)
```

## tenants

List tenants

```go
client.Tenants().List(ctx context.Context, params *TenantListParams, opts ...RequestOption) (*Page[Tenant], error)
```
Get a tenant by ID

```go
client.Tenants().Retrieve(ctx context.Context, id string, params *TenantRetrieveParams, opts ...RequestOption) (*Tenant, error)
```
Erase a tenant

```go
client.Tenants().Delete(ctx context.Context, id string, params *TenantDeleteParams, opts ...RequestOption) (*Tenant, error)
```
List a tenant's subjects

```go
client.Tenants().ListSubjects(ctx context.Context, tenantID string, params *TenantListSubjectsParams, opts ...RequestOption) (*Page[Subject], error)
```

## tool_sets

List tool sets

```go
client.ToolSets().List(ctx context.Context, params *ToolSetListParams, opts ...RequestOption) (*Page[ToolSet], error)
```
Create a new tool set

```go
client.ToolSets().Create(ctx context.Context, params *ToolSetCreateParams, opts ...RequestOption) (*ToolSet, error)
```
Get a tool set by ID

```go
client.ToolSets().Retrieve(ctx context.Context, id string, params *ToolSetRetrieveParams, opts ...RequestOption) (*ToolSet, error)
```
Delete a tool set

```go
client.ToolSets().Delete(ctx context.Context, id string, params *ToolSetDeleteParams, opts ...RequestOption) error
```
Update a tool set

```go
client.ToolSets().Update(ctx context.Context, id string, params *ToolSetUpdateParams, opts ...RequestOption) (*ToolSet, error)
```
Archive a tool set

```go
client.ToolSets().Archive(ctx context.Context, id string, params *ToolSetArchiveParams, opts ...RequestOption) (*ToolSet, error)
```
Unarchive a tool set

```go
client.ToolSets().Unarchive(ctx context.Context, id string, params *ToolSetUnarchiveParams, opts ...RequestOption) (*ToolSet, error)
```
List tool set events

```go
client.ToolSets().ListEvents(ctx context.Context, toolSetID string, params *ToolSetListEventsParams, opts ...RequestOption) (*Page[ToolSetEvent], error)
```
Get consumed OpenAPI spec

```go
client.ToolSets().RetrieveOpenAPISpec(ctx context.Context, toolSetID string, params *ToolSetRetrieveOpenAPISpecParams, opts ...RequestOption) (*GetToolSetOpenAPISpecResponse, error)
```
List tool set usage

```go
client.ToolSets().ListUsage(ctx context.Context, toolSetID string, params *ToolSetListUsageParams, opts ...RequestOption) (*Page[ToolSetUsage], error)
```

## tool_sets.secrets

List tool set secrets

```go
client.ToolSets().Secrets().List(ctx context.Context, toolSetID string, params *ToolSetSecretListParams, opts ...RequestOption) (*Page[ToolSetSecret], error)
```
Create a new tool set secret

```go
client.ToolSets().Secrets().Create(ctx context.Context, toolSetID string, params *ToolSetSecretCreateParams, opts ...RequestOption) (*ToolSetSecret, error)
```
Get a tool set secret by ID

```go
client.ToolSets().Secrets().Retrieve(ctx context.Context, toolSetID string, id string, params *ToolSetSecretRetrieveParams, opts ...RequestOption) (*ToolSetSecret, error)
```
Delete a tool set secret

```go
client.ToolSets().Secrets().Delete(ctx context.Context, toolSetID string, id string, params *ToolSetSecretDeleteParams, opts ...RequestOption) error
```
Update a tool set secret

```go
client.ToolSets().Secrets().Update(ctx context.Context, toolSetID string, id string, params *ToolSetSecretUpdateParams, opts ...RequestOption) (*ToolSetSecret, error)
```

## tool_sets.tools

List tools

```go
client.ToolSets().Tools().List(ctx context.Context, toolSetID string, params *ToolListParams, opts ...RequestOption) (*Page[Tool], error)
```
Create a new tool

```go
client.ToolSets().Tools().Create(ctx context.Context, toolSetID string, params *ToolCreateParams, opts ...RequestOption) (*Tool, error)
```
Get a tool by ID

```go
client.ToolSets().Tools().Retrieve(ctx context.Context, toolSetID string, id string, params *ToolRetrieveParams, opts ...RequestOption) (*Tool, error)
```
Delete a tool

```go
client.ToolSets().Tools().Delete(ctx context.Context, toolSetID string, id string, params *ToolDeleteParams, opts ...RequestOption) error
```
Update a tool

```go
client.ToolSets().Tools().Update(ctx context.Context, toolSetID string, id string, params *ToolUpdateParams, opts ...RequestOption) (*Tool, error)
```
Omit a tool

```go
client.ToolSets().Tools().Omit(ctx context.Context, toolSetID string, id string, params *ToolOmitParams, opts ...RequestOption) (*Tool, error)
```
Restore a tool

```go
client.ToolSets().Tools().Restore(ctx context.Context, toolSetID string, id string, params *ToolRestoreParams, opts ...RequestOption) (*Tool, error)
```

## uploads

Create an upload

```go
client.Uploads().Create(ctx context.Context, params *UploadCreateParams, opts ...RequestOption) (*Upload, error)
```
Get an upload by ID

```go
client.Uploads().Retrieve(ctx context.Context, id string, params *UploadRetrieveParams, opts ...RequestOption) (*Upload, error)
```

## widget_sessions

List widget sessions

```go
client.WidgetSessions().List(ctx context.Context, params *WidgetSessionListParams, opts ...RequestOption) (*Page[WidgetSession], error)
```
Create a widget session

```go
client.WidgetSessions().Create(ctx context.Context, params *WidgetSessionCreateParams, opts ...RequestOption) (*WidgetSession, error)
```
Delete all of a tenant's widget sessions

```go
client.WidgetSessions().DeleteTenant(ctx context.Context, params *WidgetSessionDeleteTenantParams, opts ...RequestOption) (*DeleteTenantWidgetSessionsResponse, error)
```
Get a widget session by ID

```go
client.WidgetSessions().Retrieve(ctx context.Context, id string, params *WidgetSessionRetrieveParams, opts ...RequestOption) (*WidgetSession, error)
```
Delete a widget session

```go
client.WidgetSessions().Delete(ctx context.Context, id string, params *WidgetSessionDeleteParams, opts ...RequestOption) error
```
Revoke a widget session

```go
client.WidgetSessions().Revoke(ctx context.Context, id string, params *WidgetSessionRevokeParams, opts ...RequestOption) (*WidgetSession, error)
```

## widgets

List widgets

```go
client.Widgets().List(ctx context.Context, params *WidgetListParams, opts ...RequestOption) (*Page[Widget], error)
```
Create a new widget

```go
client.Widgets().Create(ctx context.Context, params *WidgetCreateParams, opts ...RequestOption) (*Widget, error)
```
Get a widget by ID

```go
client.Widgets().Retrieve(ctx context.Context, id string, params *WidgetRetrieveParams, opts ...RequestOption) (*Widget, error)
```
Delete a widget

```go
client.Widgets().Delete(ctx context.Context, id string, params *WidgetDeleteParams, opts ...RequestOption) error
```
Update a widget

```go
client.Widgets().Update(ctx context.Context, id string, params *WidgetUpdateParams, opts ...RequestOption) (*Widget, error)
```
Archive a widget

```go
client.Widgets().Archive(ctx context.Context, id string, params *WidgetArchiveParams, opts ...RequestOption) (*Widget, error)
```
Unarchive a widget

```go
client.Widgets().Unarchive(ctx context.Context, id string, params *WidgetUnarchiveParams, opts ...RequestOption) (*Widget, error)
```

## workspace_secrets

List workspace secrets

```go
client.WorkspaceSecrets().List(ctx context.Context, params *WorkspaceSecretListParams, opts ...RequestOption) (*Page[WorkspaceSecret], error)
```
Create a new workspace secret

```go
client.WorkspaceSecrets().Create(ctx context.Context, params *WorkspaceSecretCreateParams, opts ...RequestOption) (*WorkspaceSecret, error)
```
Get a workspace secret by ID

```go
client.WorkspaceSecrets().Retrieve(ctx context.Context, id string, params *WorkspaceSecretRetrieveParams, opts ...RequestOption) (*WorkspaceSecret, error)
```
Delete a workspace secret

```go
client.WorkspaceSecrets().Delete(ctx context.Context, id string, params *WorkspaceSecretDeleteParams, opts ...RequestOption) error
```
Update a workspace secret

```go
client.WorkspaceSecrets().Update(ctx context.Context, id string, params *WorkspaceSecretUpdateParams, opts ...RequestOption) (*WorkspaceSecret, error)
```

## webhooks

```go
VerifyWebhook(secret string, payload []byte, headers http.Header) error
UnwrapWebhook(secret string, payload []byte, headers http.Header) (*ObjectiveEventWebhookData, error)
client.Webhooks().Unwrap(payload []byte, headers http.Header) (*ObjectiveEventWebhookData, error)
```
