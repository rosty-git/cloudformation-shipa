package resource

import "github.com/rostislavgit/cloudformation-shipa/shipa"

func optionalBool(val *bool) bool {
	if val != nil {
		return *val
	}

	return false
}

func optionalString(val *string) string {
	if val != nil {
		return *val
	}

	return ""
}

func optionalInt(val *int) int64 {
	if val != nil {
		return int64(*val)
	}

	return 0
}

func convertModel(currentModel *Model) *shipa.AppDeploy {
	return &shipa.AppDeploy{
		App:            *currentModel.App,
		Image:          *currentModel.Image,
		AppConfig:      convertAppConfig(currentModel.AppConfig),
		CanarySettings: convertCanarySettings(currentModel.CanarySettings),
		PodAutoScaler:  convertPodAutoScaler(currentModel.PodAutoScaler),
		Port:           convertPort(currentModel.Port),
		Registry:       convertRegistry(currentModel.Registry),
		Volumes:        convertVolumes(currentModel.Volumes),
	}
}

func convertAppConfig(config *AppConfig) *shipa.AppDeployConfig {
	if config == nil {
		return nil
	}

	return &shipa.AppDeployConfig{
		Team:        optionalString(config.Team),
		Framework:   optionalString(config.Framework),
		Description: optionalString(config.Description),
		Env:         config.Env,
		Plan:        optionalString(config.Plan),
		Router:      optionalString(config.Router),
		Tags:        config.Tags,
	}
}

func convertCanarySettings(settings *CanarySettings) *shipa.AppDeployCanarySettings {
	if settings == nil {
		return nil
	}

	return &shipa.AppDeployCanarySettings{
		StepInterval: optionalInt(settings.StepInterval),
		StepWeight:   optionalInt(settings.StepWeight),
		Steps:        optionalInt(settings.Steps),
	}
}

func convertPodAutoScaler(scaler *PodAutoScaler) *shipa.AppDeployPodAutoScaler {
	if scaler == nil {
		return nil
	}

	return &shipa.AppDeployPodAutoScaler{
		MaxReplicas:                    optionalInt(scaler.MaxReplicas),
		MinReplicas:                    optionalInt(scaler.MinReplicas),
		TargetCPUUtilizationPercentage: optionalInt(scaler.TargetCPUUtilizationPercentage),
	}
}

func convertPort(port *Port) *shipa.AppDeployPort {
	if port == nil {
		return nil
	}

	return &shipa.AppDeployPort{
		Number:   optionalInt(port.Number),
		Protocol: optionalString(port.Protocol),
	}
}

func convertRegistry(registry *Registry) *shipa.AppDeployRegistry {
	if registry == nil {
		return nil
	}

	return &shipa.AppDeployRegistry{
		User:   optionalString(registry.User),
		Secret: optionalString(registry.Secret),
	}
}

func convertVolumes(volumes []Volume) (out []*shipa.AppDeployVolume) {
	for _, v := range volumes {
		out = append(out, convertVolume(v))
	}

	return
}

func convertVolume(v Volume) *shipa.AppDeployVolume {
	if v.Name == nil && v.Path == nil {
		return nil
	}

	return &shipa.AppDeployVolume{
		Name:    optionalString(v.Name),
		Path:    optionalString(v.Path),
		Options: convertVolumeOptions(v.Options),
	}
}

func convertVolumeOptions(options *VolumeOptions) *shipa.VolumeOptions {
	if options == nil || (options.Prop1 == nil && options.Prop2 == nil && options.Prop3 == nil) {
		return nil
	}

	return &shipa.VolumeOptions{
		Prop1: optionalString(options.Prop1),
		Prop2: optionalString(options.Prop2),
		Prop3: optionalString(options.Prop3),
	}
}
