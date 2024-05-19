package autolabel

import "fmt"

func labelInstance_bulkInsert(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
