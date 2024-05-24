package logstruct

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// AuditLogEntry represents a LogEntry as described at
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
type AuditLogEntry struct {
	ProtoPayload *AuditLogProtoPayload `json:"protoPayload"`
	Resource     *AuditResource        `json:"resource"`
}

// AuditLogProtoPayload represents AuditLog within the LogEntry.protoPayload
// See https://cloud.google.com/logging/docs/reference/audit/auditlog/rest/Shared.Types/AuditLog
type AuditLogProtoPayload struct {
	MethodName         string                   `json:"methodName"`
	Request            *AuditRequest            `json:"request"`
	AuthenticationInfo *AuditAuthenticationInfo `json:"authenticationInfo"`
	ResourceName       string                   `json:"resourceName"`
	Response           *AuditResponse           `json:"response"`
}
type AuditRequest struct {
	MachineType string `json:"machineType"`
	Name        string `json:"name"`
}
type AuditResponse struct {
	InstanceId    string `json:"targetId"`
	User          string `json:"user"`
	OperationType string `json:"operationType"`
}
type AuditResource struct {
	Labels *AuditResourceLabels `json:"labels"`
	Type   string               `json:"type"`
}
type AuditResourceLabels struct {
	ResourceId string `json:"instance_id"`
	ProjectId  string `json:"project_id"`
	Zone       string `json:"zone"`
}
type AuditAuthenticationInfo struct {
	PrincipalEmail string `json:"principalEmail"`
}
