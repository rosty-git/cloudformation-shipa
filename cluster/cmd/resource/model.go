// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Name       *string           `json:",omitempty"`
	Endpoint   *ClusterEndpoint  `json:",omitempty"`
	Resources  *ClusterResources `json:",omitempty"`
	ShipaHost  *string           `json:",omitempty"`
	ShipaToken *string           `json:",omitempty"`
}

// ClusterEndpoint is autogenerated from the json schema
type ClusterEndpoint struct {
	Addresses         []string `json:",omitempty"`
	Certificate       *string  `json:",omitempty"`
	ClientCertificate *string  `json:",omitempty"`
	ClientKey         *string  `json:",omitempty"`
	Token             *string  `json:",omitempty"`
	Username          *string  `json:",omitempty"`
	Password          *string  `json:",omitempty"`
}

// ClusterResources is autogenerated from the json schema
type ClusterResources struct {
	Frameworks         []Framework         `json:",omitempty"`
	IngressControllers []IngressController `json:",omitempty"`
}

// Framework is autogenerated from the json schema
type Framework struct {
	Name *string `json:",omitempty"`
}

// IngressController is autogenerated from the json schema
type IngressController struct {
	IngressIP     *string `json:",omitempty"`
	ServiceType   *string `json:",omitempty"`
	Type          *string `json:",omitempty"`
	HTTPPort      *int    `json:",omitempty"`
	HTTPSPort     *int    `json:",omitempty"`
	ProtectedPort *int    `json:",omitempty"`
	Debug         *bool   `json:",omitempty"`
	AcmeEmail     *string `json:",omitempty"`
	AcmeServer    *string `json:",omitempty"`
}