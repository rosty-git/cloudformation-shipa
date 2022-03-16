module github.com/rostislavgit/cloudformation-shipa/network-policy

go 1.16

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.0.3
	github.com/rostislavgit/cloudformation-shipa/shipa v0.0.0-00010101000000-000000000000
)

replace github.com/rostislavgit/cloudformation-shipa/shipa => ../shipa
