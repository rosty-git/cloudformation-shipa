package resource

import "github.com/rostislavgit/cloudformation-shipa/shipa"

func convertModel(currentModel *Model) *shipa.CreateAppRequest {
	return &shipa.CreateAppRequest{
		Name:      optionalString(currentModel.Name),
		TeamOwner: optionalString(currentModel.Teamowner),
		Pool:      optionalString(currentModel.Framework),
		Plan:      optionalString(currentModel.Plan),
		Tags:      currentModel.Tags,
	}
}

func optionalString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
