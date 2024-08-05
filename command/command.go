// command/command.go
package command

// Command represents a command structure for messages
type Command struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}
