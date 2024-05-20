package autolabel

import "fmt"

func labelInstance_update(logentry *AuditLogEntry, paths []string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
