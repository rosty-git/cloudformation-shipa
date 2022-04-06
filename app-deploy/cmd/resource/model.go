// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	App            *string         `json:",omitempty"`
	Image          *string         `json:",omitempty"`
	AppConfig      *AppConfig      `json:",omitempty"`
	CanarySettings *CanarySettings `json:",omitempty"`
	PodAutoScaler  *PodAutoScaler  `json:",omitempty"`
	Port           *Port           `json:",omitempty"`
	Registry       *Registry       `json:",omitempty"`
	Volumes        []Volume        `json:",omitempty"`
	ShipaHost      *string         `json:",omitempty"`
	ShipaToken     *string         `json:",omitempty"`
}

// AppConfig is autogenerated from the json schema
type AppConfig struct {
	Team        *string  `json:",omitempty"`
	Framework   *string  `json:",omitempty"`
	Description *string  `json:",omitempty"`
	Env         []string `json:",omitempty"`
	Plan        *string  `json:",omitempty"`
	Router      *string  `json:",omitempty"`
	Tags        []string `json:",omitempty"`
}

// CanarySettings is autogenerated from the json schema
type CanarySettings struct {
	StepInterval *int `json:",omitempty"`
	StepWeight   *int `json:",omitempty"`
	Steps        *int `json:",omitempty"`
}

// PodAutoScaler is autogenerated from the json schema
type PodAutoScaler struct {
	MaxReplicas                    *int `json:",omitempty"`
	MinReplicas                    *int `json:",omitempty"`
	TargetCPUUtilizationPercentage *int `json:",omitempty"`
}

// Port is autogenerated from the json schema
type Port struct {
	Number   *int    `json:",omitempty"`
	Protocol *string `json:",omitempty"`
}

// Registry is autogenerated from the json schema
type Registry struct {
	User   *string `json:",omitempty"`
	Secret *string `json:",omitempty"`
}

// Volume is autogenerated from the json schema
type Volume struct {
	Name         *string        `json:",omitempty"`
	MountPath    *string        `json:",omitempty"`
	MountOptions *VolumeOptions `json:",omitempty"`
}

// VolumeOptions is autogenerated from the json schema
type VolumeOptions struct {
	AdditionalProp1 *string `json:",omitempty"`
	AdditionalProp2 *string `json:",omitempty"`
	AdditionalProp3 *string `json:",omitempty"`
}
