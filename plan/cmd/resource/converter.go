package resource

import "github.com/rostislavgit/cloudformation-shipa/plan/internal/shipa"

func convertModel(currentModel *Model) *shipa.CreatePlanRequest {
	return &shipa.CreatePlanRequest{
		Name:     optionalString(currentModel.Name),
		Memory:   optionalString(currentModel.Memory),
		Swap:     optionalString(currentModel.Swap),
		CPUShare: int64(optionalInt(currentModel.CPUShare)),
		Default:  optionalBool(currentModel.Default),
		Public:   optionalBool(currentModel.Public),
		Org:      optionalString(currentModel.Org),
		Teams:    currentModel.Teams,
	}
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
