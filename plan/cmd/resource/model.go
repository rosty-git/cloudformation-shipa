// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Name       *string  `json:",omitempty"`
	Memory     *string  `json:",omitempty"`
	Swap       *string  `json:",omitempty"`
	CPUShare   *int     `json:",omitempty"`
	Default    *bool    `json:",omitempty"`
	Public     *bool    `json:",omitempty"`
	Org        *string  `json:",omitempty"`
	Teams      []string `json:",omitempty"`
	ShipaHost  *string  `json:",omitempty"`
	ShipaToken *string  `json:",omitempty"`
}