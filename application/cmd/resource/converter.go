package resource

import "github.com/rostislavgit/cloudformation-shipa/application/internal/shipa"

func convertModel(currentModel *Model) *shipa.App {
    return &shipa.App{
        Name:        optionalString(currentModel.Name),
        Description: optionalString(currentModel.Description),
        TeamOwner:   optionalString(currentModel.Teamowner),
        Pool:        optionalString(currentModel.Framework),
        Plan:        convertPlan(currentModel.Plan),
        Units:       convertUnits(currentModel.Units),
        Cname:       currentModel.Cname,
        IP:          optionalString(currentModel.IP),
        Org:         optionalString(currentModel.Org),
        Entrypoints: convertEntrypoints(currentModel.Entrypoints),
        Routers:     convertRouters(currentModel.Routers),
        Lock:        convertLock(currentModel.Lock),
        Tags:        currentModel.Tags,
        Platform:    optionalString(currentModel.Platform),
        Status:      optionalString(currentModel.Status),
        Error:       optionalString(currentModel.Error),
    }
}

func optionalString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func convertUnits(units []Unit) (out []*shipa.Unit) {
	for _, u := range units {
		out = append(out, convertUnit(u))
	}

	return
}

func convertUnit(u Unit) *shipa.Unit {
	return &shipa.Unit{
		ID:          optionalString(u.ID),
		Name:        optionalString(u.Name),
		AppName:     optionalString(u.AppName),
		ProcessName: optionalString(u.ProcessName),
		Type:        optionalString(u.Type),
		IP:          optionalString(u.IP),
		Status:      optionalString(u.Status),
		Version:     optionalString(u.Version),
		Org:         optionalString(u.Org),
		HostAddr:    optionalString(u.HostAddr),
		HostPort:    optionalString(u.HostPort),
		Address:     convertAddress(u.Address),
	}
}

func convertAddress(a *Address) *shipa.Address {
	return &shipa.Address{
		Scheme:      optionalString(a.Scheme),
		Host:        optionalString(a.Host),
		Opaque:      optionalString(a.Opaque),
		User:        optionalString(a.User),
		Path:        optionalString(a.Path),
		RawPath:     optionalString(a.RawPath),
		ForceQuery:  optionalBool(a.ForceQuery),
		RawQuery:    optionalString(a.RawQuery),
		Fragment:    optionalString(a.Fragment),
		RawFragment: optionalString(a.RawFragment),
	}
}

func convertLock(lock *Lock) *shipa.Lock {
	if lock == nil {
		return nil
	}

	return &shipa.Lock{
		Locked:      optionalBool(lock.Locked),
		Reason:      optionalString(lock.Reason),
		Owner:       optionalString(lock.Owner),
		AcquireDate: optionalString(lock.AcquireDate),
	}
}

func convertRouters(routers []Router) (out []*shipa.Router) {
	for _, r := range routers {
		out = append(out, convertRouter(r))
	}

	return
}

func convertRouter(r Router) *shipa.Router {
	return &shipa.Router{
		Name:    optionalString(r.Name),
		Type:    optionalString(r.Type),
		Address: optionalString(r.Address),
		Default: optionalBool(r.Default),
	}
}

func convertEntrypoints(entrypoints []Entrypoint) (out []*shipa.Entrypoint) {
	for _, entry := range entrypoints {
		out = append(out, convertEntrypoint(entry))
	}

	return
}

func convertEntrypoint(entry Entrypoint) *shipa.Entrypoint {
	return &shipa.Entrypoint{
		Cname:  optionalString(entry.Cname),
		Scheme: optionalString(entry.Scheme),
	}
}

func convertPlan(plan *Plan) *shipa.Plan {
	if plan == nil {
		return nil
	}

	return &shipa.Plan{
		Name:     optionalString(plan.Name),
		Memory:   optionalIntToInt64(plan.Memory),
		Swap:     optionalIntToInt64(plan.Swap),
		CPUShare: optionalIntToInt64(plan.CPUShare),
		Default:  optionalBool(plan.Default),
		Public:   optionalBool(plan.Public),
		Org:      optionalString(plan.Org),
		Teams:    plan.Teams,
	}
}

func optionalBool(val *bool) bool {
	if val == nil {
		return false
	}

	return *val
}

func optionalIntToInt64(val *int) int64 {
	if val == nil {
		return 0
	}

	return int64(*val)
}
