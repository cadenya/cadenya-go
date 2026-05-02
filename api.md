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

## Feedback

Methods:

- <code title="get /v1/agents/{agentId}/feedback">client.Agents.Feedback.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentFeedbackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentFeedbackListParams">AgentFeedbackListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveFeedback">ObjectiveFeedback</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## WebhookDeliveries

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDelivery">WebhookDelivery</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDeliveryData">WebhookDeliveryData</a>

Methods:

- <code title="get /v1/agents/{agentId}/webhook_deliveries">client.Agents.WebhookDeliveries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentWebhookDeliveryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentWebhookDeliveryListParams">AgentWebhookDeliveryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#WebhookDelivery">WebhookDelivery</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Variations

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationParam">AgentVariationParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationInfoParam">AgentVariationInfoParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecParam">AgentVariationSpecParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecCompactionConfigParam">AgentVariationSpecCompactionConfigParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecConstraintsParam">AgentVariationSpecConstraintsParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecModelConfigParam">AgentVariationSpecModelConfigParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecProgressiveDiscoveryParam">AgentVariationSpecProgressiveDiscoveryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigSummarizationStrategyParam">CompactionConfigSummarizationStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigToolResultClearingStrategyParam">CompactionConfigToolResultClearingStrategyParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignmentParam">VariationAssignmentParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerAssignmentParam">VariationMemoryLayerAssignmentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationInfo">AgentVariationInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpec">AgentVariationSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecCompactionConfig">AgentVariationSpecCompactionConfig</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecConstraints">AgentVariationSpecConstraints</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecModelConfig">AgentVariationSpecModelConfig</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationSpecProgressiveDiscovery">AgentVariationSpecProgressiveDiscovery</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigSummarizationStrategy">CompactionConfigSummarizationStrategy</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CompactionConfigToolResultClearingStrategy">CompactionConfigToolResultClearingStrategy</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignment">VariationAssignment</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerAssignment">VariationMemoryLayerAssignment</a>

Methods:

- <code title="post /v1/agents/{agentId}/variations">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationNewParams">AgentVariationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/variations/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/agents/{agentId}/variations/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationUpdateParams">AgentVariationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/variations">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationListParams">AgentVariationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariation">AgentVariation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agentId}/variations/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/agents/{agentId}/variations/{variationId}/assignments">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.AddAssignment">AddAssignment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, variationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationAddAssignmentParams">AgentVariationAddAssignmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignment">VariationAssignment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/{agentId}/variations/{variationId}/memory_layer_assignments">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.AddMemoryLayer">AddMemoryLayer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, variationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationAddMemoryLayerParams">AgentVariationAddMemoryLayerParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerAssignment">VariationMemoryLayerAssignment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agentId}/variations/{variationId}/assignments/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.RemoveAssignment">RemoveAssignment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, variationID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v1/agents/{agentId}/variations/{variationId}/memory_layer_assignments/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.RemoveMemoryLayer">RemoveMemoryLayer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, variationID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /v1/agents/{agentId}/variations/{variationId}/memory_layer_assignments/{id}">client.Agents.Variations.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationService.UpdateMemoryLayer">UpdateMemoryLayer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, variationID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationUpdateMemoryLayerParams">AgentVariationUpdateMemoryLayerParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerAssignment">VariationMemoryLayerAssignment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Schedules

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleSpecParam">AgentScheduleSpecParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleSpecScheduleParam">AgentScheduleSpecScheduleParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleCalendarParam">ScheduleCalendarParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleIntervalParam">ScheduleIntervalParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleRangeParam">ScheduleRangeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSchedule">AgentSchedule</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleInfo">AgentScheduleInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleSpec">AgentScheduleSpec</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleSpecSchedule">AgentScheduleSpecSchedule</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleCalendar">ScheduleCalendar</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleInterval">ScheduleInterval</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ScheduleRange">ScheduleRange</a>

Methods:

- <code title="post /v1/agents/{agentId}/schedules">client.Agents.Schedules.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleNewParams">AgentScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSchedule">AgentSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/schedules/{id}">client.Agents.Schedules.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSchedule">AgentSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/agents/{agentId}/schedules/{id}">client.Agents.Schedules.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleUpdateParams">AgentScheduleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSchedule">AgentSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agentId}/schedules">client.Agents.Schedules.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleListParams">AgentScheduleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentSchedule">AgentSchedule</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agentId}/schedules/{id}">client.Agents.Schedules.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Objectives

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryReferenceParam">MemoryReferenceParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveDataParam">ObjectiveDataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ObjectiveDataSecretParam">ObjectiveDataSecretParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AssistantMessage">AssistantMessage</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AssistantToolCall">AssistantToolCall</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#CallableTool">CallableTool</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ContextWindowCompacted">ContextWindowCompacted</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryRead">MemoryRead</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryReference">MemoryReference</a>
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

# MemoryLayers

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerSpecParam">MemoryLayerSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayer">MemoryLayer</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerInfo">MemoryLayerInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerSpec">MemoryLayerSpec</a>

Methods:

- <code title="post /v1/memory_layers">client.MemoryLayers.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerNewParams">MemoryLayerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayer">MemoryLayer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/memory_layers/{id}">client.MemoryLayers.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayer">MemoryLayer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/memory_layers/{id}">client.MemoryLayers.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerUpdateParams">MemoryLayerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayer">MemoryLayer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/memory_layers">client.MemoryLayers.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerListParams">MemoryLayerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayer">MemoryLayer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/memory_layers/{id}">client.MemoryLayers.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Entries

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryCreateSpecParam">MemoryEntryCreateSpecParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryUpdateSpecParam">MemoryEntryUpdateSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntry">MemoryEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryDetail">MemoryEntryDetail</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryInfo">MemoryEntryInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntrySpec">MemoryEntrySpec</a>

Methods:

- <code title="post /v1/memory_layers/{memoryLayerId}/entries">client.MemoryLayers.Entries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memoryLayerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryNewParams">MemoryLayerEntryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryDetail">MemoryEntryDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/memory_layers/{memoryLayerId}/entries/{id}">client.MemoryLayers.Entries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memoryLayerID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryDetail">MemoryEntryDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/memory_layers/{memoryLayerId}/entries/{id}">client.MemoryLayers.Entries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memoryLayerID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryUpdateParams">MemoryLayerEntryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryDetail">MemoryEntryDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/memory_layers/{memoryLayerId}/entries">client.MemoryLayers.Entries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memoryLayerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryListParams">MemoryLayerEntryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntry">MemoryEntry</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/memory_layers/{memoryLayerId}/entries/{id}">client.MemoryLayers.Entries.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memoryLayerID <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Uploads

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadSpecParam">UploadSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Upload">Upload</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadInfo">UploadInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadSpec">UploadSpec</a>

Methods:

- <code title="post /v1/uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadNewParams">UploadNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/uploads/{id}">client.Uploads.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#UploadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# BulkWorkspaceResources

Params Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentEntryParam">AgentEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleEntryParam">AgentScheduleEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationEntryParam">AgentVariationEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyDataParam">BulkWorkspaceApplyDataParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryItemParam">MemoryEntryItemParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntryParam">MemoryLayerEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolEntryParam">ToolEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetEntryParam">ToolSetEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignmentEntryParam">VariationAssignmentEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerEntryParam">VariationMemoryLayerEntryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentEntry">AgentEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentScheduleEntry">AgentScheduleEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#AgentVariationEntry">AgentVariationEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApply">BulkWorkspaceApply</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyData">BulkWorkspaceApplyData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyInfo">BulkWorkspaceApplyInfo</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyStatus">BulkWorkspaceApplyStatus</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryEntryItem">MemoryEntryItem</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#MemoryLayerEntry">MemoryLayerEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolEntry">ToolEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#ToolSetEntry">ToolSetEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationAssignmentEntry">VariationAssignmentEntry</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#VariationMemoryLayerEntry">VariationMemoryLayerEntry</a>

Methods:

- <code title="get /v1/bulk_workspace_applies/{id}">client.BulkWorkspaceResources.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApply">BulkWorkspaceApply</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bulk_workspace_applies">client.BulkWorkspaceResources.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceListParams">BulkWorkspaceResourceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApply">BulkWorkspaceApply</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/bulk_workspace_applies">client.BulkWorkspaceResources.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceService.Apply">Apply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceApplyParams">BulkWorkspaceResourceApplyParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApply">BulkWorkspaceApply</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Results

Response Types:

- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResult">BulkWorkspaceApplyResult</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultData">BulkWorkspaceApplyResultData</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataAgentOutcome">BulkWorkspaceApplyResultDataAgentOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataAgentScheduleOutcome">BulkWorkspaceApplyResultDataAgentScheduleOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataAgentVariationOutcome">BulkWorkspaceApplyResultDataAgentVariationOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataMemoryEntryOutcome">BulkWorkspaceApplyResultDataMemoryEntryOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataMemoryLayerOutcome">BulkWorkspaceApplyResultDataMemoryLayerOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataToolOutcome">BulkWorkspaceApplyResultDataToolOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataToolSetOutcome">BulkWorkspaceApplyResultDataToolSetOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataVariationAssignmentOutcome">BulkWorkspaceApplyResultDataVariationAssignmentOutcome</a>
- <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome">BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome</a>

Methods:

- <code title="get /v1/bulk_workspace_applies/{bulkWorkspaceApplyId}/results">client.BulkWorkspaceResources.Results.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bulkWorkspaceApplyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceResourceResultListParams">BulkWorkspaceResourceResultListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go/packages/pagination#CursorPagination">CursorPagination</a>[<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go">cadenya</a>.<a href="https://pkg.go.dev/github.com/cadenya/cadenya-go#BulkWorkspaceApplyResult">BulkWorkspaceApplyResult</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
