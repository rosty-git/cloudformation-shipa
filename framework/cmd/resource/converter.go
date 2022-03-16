package resource

import "github.com/rostislavgit/cloudformation-shipa/shipa"

func convertModel(currentModel *Model) *shipa.PoolConfig {
	return &shipa.PoolConfig{
		Name:      optionalString(currentModel.Name),
		Resources: convertResources(currentModel.Resources),
	}
}

func convertResources(resources *PoolResources) *shipa.PoolResources {
	if resources == nil {
		return nil
	}

	return &shipa.PoolResources{
		General: convertGeneral(resources.General),
		Node:    convertNode(resources.Node),
	}
}

func convertGeneral(general *PoolGeneral) *shipa.PoolGeneral {
	if general == nil {
		return nil
	}
	return &shipa.PoolGeneral{
		Setup:            convertSetup(general.Setup),
		Plan:             convertPlan(general.Plan),
		Security:         convertSecurity(general.Security),
		Access:           convertAccess(general.Access),
		Services:         convertAccess(general.Services),
		Router:           optionalString(general.Router),
		Volumes:          general.Volumes,
		ContainerPolicy:  convertContainerPolicy(general.ContainerPolicy),
		NodeSelectors:    convertNodeSelectors(general.NodeSelectors),
		PodAutoScaler:    convertPodAutoScaler(general.PodAutoScaler),
		DomainPolicy:     convertDomainPolicy(general.DomainPolicy),
		AppAutoDiscovery: convertAppAutoDiscovery(general.AppAutoDiscovery),
		NetworkPolicy:    convertNetworkPolicy(general.NetworkPolicy),
	}
}

func convertAppAutoDiscovery(discovery *AppAutoDiscovery) *shipa.AppAutoDiscovery {
	if discovery == nil {
		return nil
	}

	return &shipa.AppAutoDiscovery{
		AppSelector: convertAppSelectorLabels(discovery.AppSelector),
		Suffix:      optionalString(discovery.Suffix),
	}
}

func convertAppSelectorLabels(selector []AppSelectorLabels) (out []*shipa.AppSelectorLabels) {
	for _, s := range selector {
		out = append(out, &shipa.AppSelectorLabels{
			Label: optionalString(s.Label),
		})
	}
	return
}

func convertDomainPolicy(policy *DomainPolicy) *shipa.DomainPolicy {
	if policy == nil {
		return nil
	}

	return &shipa.DomainPolicy{
		AllowedCnames: policy.AllowedCnames,
	}
}

func convertPodAutoScaler(scaler *PodAutoScaler) *shipa.PodAutoScaler {
	if scaler == nil {
		return nil
	}

	return &shipa.PodAutoScaler{
		MinReplicas:                    optionalInt(scaler.MinReplicas),
		MaxReplicas:                    optionalInt(scaler.MaxReplicas),
		TargetCPUUtilizationPercentage: optionalInt(scaler.TargetCPUUtilizationPercentage),
		DisableAppOverride:             optionalBool(scaler.DisableAppOverride),
	}
}

func convertNodeSelectors(selectors *NodeSelectors) *shipa.NodeSelectors {
	if selectors == nil {
		return nil
	}
	return &shipa.NodeSelectors{
		Terms:  convertNodeSelectorsTerms(selectors.Terms),
		Strict: optionalBool(selectors.Strict),
	}
}

func convertNodeSelectorsTerms(terms *NodeSelectorsTerms) *shipa.NodeSelectorsTerms {
	if terms == nil {
		return nil
	}

	return &shipa.NodeSelectorsTerms{
		Environment: optionalString(terms.Environment),
		OS:          optionalString(terms.OS),
	}
}

func convertNetworkPolicy(policy *PoolNetworkPolicy) *shipa.PoolNetworkPolicy {
	if policy == nil {
		return nil
	}
	return &shipa.PoolNetworkPolicy{
		Ingress:            convertNetworkPolicyConfig(policy.Ingress),
		Egress:             convertNetworkPolicyConfig(policy.Egress),
		DisableAppPolicies: optionalBool(policy.DisableAppPolicies),
	}
}

func convertNetworkPolicyConfig(config *NetworkPolicyConfig) *shipa.NetworkPolicyConfig {
	if config == nil {
		return nil
	}
	return &shipa.NetworkPolicyConfig{
		PolicyMode:        optionalString(config.PolicyMode),
		CustomRules:       convertNetworkPolicyRule(config.CustomRules),
		ShipaRules:        convertNetworkPolicyRule(config.ShipaRules),
		ShipaRulesEnabled: config.ShipaRulesEnabled,
	}
}

func convertNetworkPolicyRule(rules []NetworkPolicyRule) (out []*shipa.NetworkPolicyRule) {
	for _, r := range rules {
		out = append(out, &shipa.NetworkPolicyRule{
			ID:           optionalString(r.ID),
			Enabled:      optionalBool(r.Enabled),
			Description:  optionalString(r.Description),
			Ports:        convertNetworkPorts(r.Ports),
			Peers:        convertNetworkPeers(r.Peers),
			AllowedApps:  r.AllowedApps,
			AllowedPools: r.AllowedPools,
		})
	}
	return
}

func convertNetworkPeers(peers []NetworkPeer) (out []*shipa.NetworkPeer) {
	for _, p := range peers {
		out = append(out, &shipa.NetworkPeer{
			PodSelector:       convertNetworkPeerSelector(p.PodSelector),
			NamespaceSelector: convertNetworkPeerSelector(p.NamespaceSelector),
			IPBlock:           p.IPBlock,
		})
	}
	return
}

func convertNetworkPeerSelector(selector *NetworkPeerSelector) *shipa.NetworkPeerSelector {
	if selector == nil {
		return nil
	}
	return &shipa.NetworkPeerSelector{
		MatchLabels:      convertMatchLabels(selector.MatchLabels),
		MatchExpressions: convertMatchExpressions(selector.MatchExpressions),
	}
}

func convertMatchExpressions(expressions []SelectorExpression) (out []*shipa.SelectorExpression) {
	for _, exp := range expressions {
		out = append(out, &shipa.SelectorExpression{
			Key:      optionalString(exp.Key),
			Operator: optionalString(exp.Operator),
			Values:   exp.Values,
		})
	}
	return out
}

func convertMatchLabels(labels []Pair) map[string]string {
	var out map[string]string
	if len(labels) == 0 {
		return out
	}
	out = make(map[string]string)

	for _, label := range labels {
		key := optionalString(label.Key)
		value := optionalString(label.Value)
		out[key] = value
	}
	return out
}

func convertNetworkPorts(ports []NetworkPort) (out []*shipa.NetworkPort) {
	for _, p := range ports {
		out = append(out, &shipa.NetworkPort{
			Protocol: optionalString(p.Protocol),
			Port:     optionalInt(p.Port),
		})
	}
	return
}

func convertContainerPolicy(policy *PoolContainerPolicy) *shipa.PoolContainerPolicy {
	if policy == nil {
		return nil
	}
	return &shipa.PoolContainerPolicy{
		AllowedHosts: policy.AllowedHosts,
	}
}

func convertAccess(access *PoolServiceAccess) *shipa.PoolServiceAccess {
	if access == nil {
		return nil
	}
	return &shipa.PoolServiceAccess{
		Append:    access.Append,
		Blacklist: access.Blacklist,
	}
}

func convertSecurity(security *PoolSecurity) *shipa.PoolSecurity {
	if security == nil {
		return nil
	}
	return &shipa.PoolSecurity{
		DisableScan:        optionalBool(security.DisableScan),
		ScanPlatformLayers: optionalBool(security.ScanPlatformLayers),
		IgnoreComponents:   security.IgnoreComponents,
		IgnoreCVES:         security.IgnoreCVES,
	}
}

func convertPlan(plan *PoolPlan) *shipa.PoolPlan {
	if plan == nil {
		return nil
	}
	return &shipa.PoolPlan{
		Name: optionalString(plan.Name),
	}
}

func convertSetup(setup *PoolSetup) *shipa.PoolSetup {
	if setup == nil {
		return nil
	}
	return &shipa.PoolSetup{
		Default:             optionalBool(setup.Default),
		Public:              optionalBool(setup.Public),
		Provisioner:         optionalString(setup.Provisioner),
		KubernetesNamespace: optionalString(setup.KubernetesNamespace),
	}
}

func convertNode(node *PoolNode) *shipa.PoolNode {
	if node == nil {
		return nil
	}
	return &shipa.PoolNode{
		Drivers:   node.Drivers,
		AutoScale: convertAutoScale(node.AutoScale),
	}
}

func convertAutoScale(scale *PoolAutoScale) *shipa.PoolAutoScale {
	if scale == nil {
		return nil
	}
	return &shipa.PoolAutoScale{
		MaxContainer: optionalInt(scale.MaxContainer),
		MaxMemory:    optionalInt(scale.MaxMemory),
		ScaleDown:    optionalFloat(scale.ScaleDown),
		Rebalance:    optionalBool(scale.Rebalance),
	}
}

func optionalFloat(val *float64) float64 {
	if val == nil {
		return 0
	}
	return *val
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
