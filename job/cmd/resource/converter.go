package resource

import "github.com/rostislavgit/cloudformation-shipa/shipa"

func convertModel(currentModel *Model) *shipa.JobCreateRequest {
	return &shipa.JobCreateRequest{
		Name:         optionalString(currentModel.Name),
		Framework:    optionalString(currentModel.Framework),
		Containers:   convertContainers(currentModel.Containers),
		Policy:       convertPolicy(currentModel.Policy),
		BackoffLimit: int64(optionalInt(currentModel.BackoffLimit)),
		Completions:  int64(optionalInt(currentModel.Completions)),
		Parallelism:  int64(optionalInt(currentModel.Parallelism)),
		Suspend:      optionalBool(currentModel.Suspend),
		Description:  optionalString(currentModel.Description),
		Team:         optionalString(currentModel.Team),
		Type:         optionalString(currentModel.Type),
		Version:      optionalString(currentModel.Version),
	}
}

func convertPolicy(policy *Policy) *shipa.JobPolicy {
	if policy == nil {
		return nil
	}

	return &shipa.JobPolicy{
		RestartPolicy: optionalString(policy.RestartPolicy),
	}
}

func convertContainers(containers []JobContainer) (out []*shipa.JobContainer) {
	for _, container := range containers {
		out = append(out, &shipa.JobContainer{
			Command: container.Command,
			Image:   optionalString(container.Image),
			Name:    optionalString(container.Name),
		})
	}

	return
}

func optionalString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func optionalBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func optionalInt(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}
