# Shared Params Types

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#AccountResourceMetadataParam">AccountResourceMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#BareMetadataParam">BareMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#CreateOperationMetadataParam">CreateOperationMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#CreateResourceMetadataParam">CreateResourceMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#ResourceMetadataParam">ResourceMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#UpdateResourceMetadataParam">UpdateResourceMetadataParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#AccountResourceMetadata">AccountResourceMetadata</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#BareMetadata">BareMetadata</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#OperationMetadata">OperationMetadata</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/shared#ResourceMetadata">ResourceMetadata</a>

# Account

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ProfileParam">ProfileParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ProfileSpecParam">ProfileSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AccountSpec">AccountSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Profile">Profile</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ProfileSpec">ProfileSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#RotateWebhookSigningKeyResponse">RotateWebhookSigningKeyResponse</a>

Methods:

- <code title="get /v1/account">client.Account.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/account/rotate_webhook_signing_key">client.Account.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AccountService.RotateWebhookSigningKey">RotateWebhookSigningKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#RotateWebhookSigningKeyResponse">RotateWebhookSigningKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentParam">AgentParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentInfoParam">AgentInfoParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSpecParam">AgentSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Agent">Agent</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentInfo">AgentInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSpec">AgentSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Page">Page</a>

Methods:

- <code title="post /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentNewParams">AgentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentUpdateParams">AgentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentListParams">AgentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Agent">Agent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{id}">client.Agents.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## WebhookDeliveries

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDelivery">WebhookDelivery</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDeliveryData">WebhookDeliveryData</a>

Methods:

- <code title="get /v1/agents/{agentId}/webhook_deliveries">client.Agents.WebhookDeliveries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentWebhookDeliveryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentWebhookDeliveryListParams">AgentWebhookDeliveryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDelivery">WebhookDelivery</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AgentVariations

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationParam">AgentVariationParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationInfoParam">AgentVariationInfoParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecParam">AgentVariationSpecParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecCompactionConfigParam">AgentVariationSpecCompactionConfigParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecConstraintsParam">AgentVariationSpecConstraintsParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecModelConfigParam">AgentVariationSpecModelConfigParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecToolSelectionParam">AgentVariationSpecToolSelectionParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigSummarizationStrategyParam">CompactionConfigSummarizationStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigToolResultClearingStrategyParam">CompactionConfigToolResultClearingStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSelectionAssignedToolsParam">ToolSelectionAssignedToolsParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSelectionAutoDiscoveryParam">ToolSelectionAutoDiscoveryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignmentParam">VariationAssignmentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationInfo">AgentVariationInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpec">AgentVariationSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecCompactionConfig">AgentVariationSpecCompactionConfig</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecConstraints">AgentVariationSpecConstraints</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecModelConfig">AgentVariationSpecModelConfig</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecToolSelection">AgentVariationSpecToolSelection</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigSummarizationStrategy">CompactionConfigSummarizationStrategy</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigToolResultClearingStrategy">CompactionConfigToolResultClearingStrategy</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSelectionAssignedTools">ToolSelectionAssignedTools</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSelectionAutoDiscovery">ToolSelectionAutoDiscovery</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignment">VariationAssignment</a>

Methods:

- <code title="post /v1/agents/{agentId}/variations">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationNewParams">AgentVariationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/variations/{id}">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/agents/{agentId}/variations/{id}">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationUpdateParams">AgentVariationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/variations">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationListParams">AgentVariationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agentId}/variations/{id}">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/agent_variations/{agentVariationId}/assignments">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.AddAssignment">AddAssignment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentVariationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationAddAssignmentParams">AgentVariationAddAssignmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignment">VariationAssignment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agent_variations/{agentVariationId}/assignments/{id}">client.AgentVariations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.RemoveAssignment">RemoveAssignment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentVariationID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Objectives

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveDataParam">ObjectiveDataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveDataSecretParam">ObjectiveDataSecretParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AssistantMessage">AssistantMessage</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AssistantToolCall">AssistantToolCall</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CallableTool">CallableTool</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ContextWindowCompacted">ContextWindowCompacted</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Objective">Objective</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContextWindow">ObjectiveContextWindow</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContextWindowData">ObjectiveContextWindowData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveData">ObjectiveData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveDataSecret">ObjectiveDataSecret</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveError">ObjectiveError</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveEventData">ObjectiveEventData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveEventInfo">ObjectiveEventInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveInfo">ObjectiveInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveStatus">ObjectiveStatus</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SubObjectiveCreated">SubObjectiveCreated</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolApprovalRequested">ToolApprovalRequested</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolApproved">ToolApproved</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolCalled">ToolCalled</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolDenied">ToolDenied</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolError">ToolError</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolResult">ToolResult</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UserMessage">UserMessage</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveCompactResponse">ObjectiveCompactResponse</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContinueResponse">ObjectiveContinueResponse</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveListEventsResponse">ObjectiveListEventsResponse</a>

Methods:

- <code title="post /v1/objectives">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveNewParams">ObjectiveNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Objective">Objective</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives/{id}">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Objective">Objective</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveListParams">ObjectiveListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Objective">Objective</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/objectives/{objectiveId}/cancel">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveCancelParams">ObjectiveCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Objective">Objective</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/objectives/{objectiveId}/compact">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.Compact">Compact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveCompactParams">ObjectiveCompactParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveCompactResponse">ObjectiveCompactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/objectives/{objectiveId}/continue">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.Continue">Continue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContinueParams">ObjectiveContinueParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContinueResponse">ObjectiveContinueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives/{objectiveId}/context_windows">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.ListContextWindows">ListContextWindows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveListContextWindowsParams">ObjectiveListContextWindowsParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveContextWindow">ObjectiveContextWindow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives/{objectiveId}/events">client.Objectives.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveListEventsParams">ObjectiveListEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveListEventsResponse">ObjectiveListEventsResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTool">ObjectiveTool</a>

Methods:

- <code title="get /v1/objectives/{objectiveId}/tools">client.Objectives.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolListParams">ObjectiveToolListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTool">ObjectiveTool</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ToolCalls

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCall">ObjectiveToolCall</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallData">ObjectiveToolCallData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallInfo">ObjectiveToolCallInfo</a>

Methods:

- <code title="get /v1/objectives/{objectiveId}/tool_calls">client.Objectives.ToolCalls.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallListParams">ObjectiveToolCallListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCall">ObjectiveToolCall</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/objectives/{objectiveId}/tool_calls/{toolCallId}/approve">client.Objectives.ToolCalls.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, toolCallID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallApproveParams">ObjectiveToolCallApproveParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCall">ObjectiveToolCall</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/objectives/{objectiveId}/tool_calls/{toolCallId}/deny">client.Objectives.ToolCalls.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallService.Deny">Deny</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, toolCallID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCallDenyParams">ObjectiveToolCallDenyParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveToolCall">ObjectiveToolCall</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tasks

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTask">ObjectiveTask</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTaskData">ObjectiveTaskData</a>

Methods:

- <code title="get /v1/objectives/{objectiveId}/tasks/{id}">client.Objectives.Tasks.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTask">ObjectiveTask</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives/{objectiveId}/tasks">client.Objectives.Tasks.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTaskService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTaskListParams">ObjectiveTaskListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveTask">ObjectiveTask</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Feedback

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackDataParam">ObjectiveFeedbackDataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedback">ObjectiveFeedback</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackData">ObjectiveFeedbackData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackInfo">ObjectiveFeedbackInfo</a>

Methods:

- <code title="post /v1/objectives/{objectiveId}/feedback">client.Objectives.Feedback.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackNewParams">ObjectiveFeedbackNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedback">ObjectiveFeedback</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/objectives/{objectiveId}/feedback">client.Objectives.Feedback.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectiveID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedbackListParams">ObjectiveFeedbackListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedback">ObjectiveFeedback</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelSpec">ModelSpec</a>

Methods:

- <code title="get /v1/models/{id}">client.Models.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelListParams">ModelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Model">Model</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/models/{id}/status">client.Models.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelService.SetStatus">SetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ModelSetStatusParams">ModelSetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SearchSearchToolsOrToolSetsResponse">SearchSearchToolsOrToolSetsResponse</a>

Methods:

- <code title="get /v1/search/tools_or_tool_sets">client.Search.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SearchService.SearchToolsOrToolSets">SearchToolsOrToolSets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SearchSearchToolsOrToolSetsParams">SearchSearchToolsOrToolSetsParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SearchSearchToolsOrToolSetsResponse">SearchSearchToolsOrToolSetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ToolSets

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#McpToolFilterParam">McpToolFilterParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapterParam">ToolSetAdapterParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapterHTTPParam">ToolSetAdapterHTTPParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapterMcpParam">ToolSetAdapterMcpParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetSpecParam">ToolSetSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#McpToolFilter">McpToolFilter</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SyncCompleted">SyncCompleted</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SyncFailed">SyncFailed</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#SyncStarted">SyncStarted</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSet">ToolSet</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapter">ToolSetAdapter</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapterHTTP">ToolSetAdapterHTTP</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetAdapterMcp">ToolSetAdapterMcp</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetEvent">ToolSetEvent</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetEventData">ToolSetEventData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetInfo">ToolSetInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetSpec">ToolSetSpec</a>

Methods:

- <code title="post /v1/tool_sets">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetNewParams">ToolSetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSet">ToolSet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool_sets/{id}">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSet">ToolSet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/tool_sets/{id}">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetUpdateParams">ToolSetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSet">ToolSet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool_sets">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetListParams">ToolSetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSet">ToolSet</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/tool_sets/{id}">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/tool_sets/{toolSetId}/events">client.ToolSets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetListEventsParams">ToolSetListEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetEvent">ToolSetEvent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ConfigHTTPParam">ConfigHTTPParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ConfigMcpParam">ConfigMcpParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSpecParam">ToolSpecParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSpecConfigParam">ToolSpecConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ConfigHTTP">ConfigHTTP</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ConfigMcp">ConfigMcp</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Tool">Tool</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolInfo">ToolInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSpec">ToolSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSpecConfig">ToolSpecConfig</a>

Methods:

- <code title="post /v1/tool_sets/{toolSetId}/tools">client.ToolSets.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolNewParams">ToolSetToolNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool_sets/{toolSetId}/tools/{id}">client.ToolSets.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/tool_sets/{toolSetId}/tools/{id}">client.ToolSets.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolUpdateParams">ToolSetToolUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool_sets/{toolSetId}/tools">client.ToolSets.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolListParams">ToolSetToolListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Tool">Tool</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/tool_sets/{toolSetId}/tools/{id}">client.ToolSets.Tools.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetToolService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolSetID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# APIKeys

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeySpecParam">APIKeySpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyInfo">APIKeyInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeySpec">APIKeySpec</a>

Methods:

- <code title="post /v1/api_keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyNewParams">APIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api_keys/{id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/api_keys/{id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyUpdateParams">APIKeyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api_keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyListParams">APIKeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api_keys/{id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /v1/api_keys/{id}/rotate">client.APIKeys.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WorkspaceSecrets

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretSpecParam">WorkspaceSecretSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecret">WorkspaceSecret</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretInfo">WorkspaceSecretInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretSpec">WorkspaceSecretSpec</a>

Methods:

- <code title="post /v1/workspace_secrets">client.WorkspaceSecrets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretNewParams">WorkspaceSecretNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecret">WorkspaceSecret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workspace_secrets/{id}">client.WorkspaceSecrets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecret">WorkspaceSecret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/workspace_secrets/{id}">client.WorkspaceSecrets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretUpdateParams">WorkspaceSecretUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecret">WorkspaceSecret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workspace_secrets">client.WorkspaceSecrets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretListParams">WorkspaceSecretListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecret">WorkspaceSecret</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/workspace_secrets/{id}">client.WorkspaceSecrets.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSecretService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Workspaces

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Workspace">Workspace</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceSpec">WorkspaceSpec</a>

Methods:

- <code title="get /v1/workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceListParams">WorkspaceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Workspace">Workspace</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workspaces/current">client.Workspaces.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WorkspaceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UnsafeUnwrapWebhookEvent">UnsafeUnwrapWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UnwrapWebhookEvent">UnwrapWebhookEvent</a>
