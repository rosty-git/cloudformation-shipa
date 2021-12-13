package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/rostislavgit/cloudformation-shipa/network-policy/internal/shipa"
)

func optionalString(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}

func optionalBool(val *bool) bool {
	if val == nil {
		return false
	}

	return *val
}

func optionalInt(val *int) int {
	if val == nil {
		return 0
	}

	return *val
}

func convertNetworkPolicyRules(rules []NetworkPolicyRule) (out []*shipa.NetworkPolicyRule) {
	for _, rule := range rules {
		out = append(out, &shipa.NetworkPolicyRule{
			ID:           optionalString(rule.ID),
			Enabled:      optionalBool(rule.Enabled),
			Description:  optionalString(rule.Description),
			Ports:        convertPorts(rule.Ports),
			Peers:        convertPeers(rule.Peers),
			AllowedApps:  rule.AllowedApps,
			AllowedPools: rule.AllowedFrameworks,
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
		MatchExpressions: convertSelectorExpressions(selector.MatchExpressions),
	}
}

func convertMatchLabels(pairs []Pair) (out map[string]string) {
	if len(out) > 0 {
		out = make(map[string]string)
	}

	for _, p := range pairs {
		key, val := optionalString(p.Key), optionalString(p.Value)
		out[key] = val
	}

	return
}

func convertSelectorExpressions(exp []SelectorExpression) (out []*shipa.SelectorExpression) {
	for _, e := range exp {
		out = append(out, convertSelectorExpression(e))
	}
	return
}

func convertSelectorExpression(exp SelectorExpression) *shipa.SelectorExpression {
	return &shipa.SelectorExpression{
		Key:      optionalString(exp.Key),
		Operator: optionalString(exp.Operator),
		Values:   exp.Values,
	}
}

func convertPeers(peers []NetworkPeer) (out []*shipa.NetworkPeer) {
	for _, p := range peers {
		out = append(out, convertPeer(p))
	}
	return
}

func convertPeer(peer NetworkPeer) *shipa.NetworkPeer {
	return &shipa.NetworkPeer{
		PodSelector:       convertNetworkPeerSelector(peer.PodSelector),
		NamespaceSelector: convertNetworkPeerSelector(peer.NamespaceSelector),
		IPBlock:           peer.IPBlock,
	}
}

func convertPorts(ports []NetworkPort) (out []*shipa.NetworkPort) {
	for _, p := range ports {
		out = append(out, convertNetworkPort(p))
	}
	return
}

func convertNetworkPort(networkPort NetworkPort) *shipa.NetworkPort {
	return &shipa.NetworkPort{
		Port:     optionalInt(networkPort.Port),
		Protocol: optionalString(networkPort.Protocol),
	}
}

func convertNetworkPolicyConfig(config *NetworkPolicyConfig) *shipa.NetworkPolicyConfig {
	if config == nil {
		return nil
	}

	return &shipa.NetworkPolicyConfig{
		PolicyMode:        optionalString(config.PolicyMode),
		CustomRules:       convertNetworkPolicyRules(config.CustomRules),
		ShipaRules:        convertNetworkPolicyRules(config.ShipaRules),
		ShipaRulesEnabled: config.ShipaRulesEnabled,
	}
}

func convertModel(currentModel *Model) *shipa.NetworkPolicy {
	return &shipa.NetworkPolicy{
		App:        *currentModel.App,
		Ingress:    convertNetworkPolicyConfig(currentModel.Ingress),
		Egress:     convertNetworkPolicyConfig(currentModel.Egress),
		RestartApp: *currentModel.RestartApp,
	}
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := shipa.NewClient(*currentModel.ShipaHost, *currentModel.ShipaToken)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	_, err = client.GetApp(context.Background(), *currentModel.App)
	// app not exists
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	payload := convertModel(currentModel)
	raw, _ := json.Marshal(payload)
	err = client.CreateOrUpdateNetworkPolicy(context.Background(), payload)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("ERR: %w. Payload: %s", err, string(raw))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := shipa.NewClient(*currentModel.ShipaHost, *currentModel.ShipaToken)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	networkPolicy, err := client.GetNetworkPolicy(context.Background(), *currentModel.App)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// TODO: convert shipa networkPolicy to model
	_ = networkPolicy

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := shipa.NewClient(*currentModel.ShipaHost, *currentModel.ShipaToken)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	_, err = client.GetApp(context.Background(), *currentModel.App)
	// app not exists
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	payload := convertModel(currentModel)
	raw, _ := json.Marshal(payload)
	err = client.CreateOrUpdateNetworkPolicy(context.Background(), payload)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("ERR: %w. Payload: %s", err, string(raw))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := shipa.NewClient(*currentModel.ShipaHost, *currentModel.ShipaToken)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	_, err = client.GetApp(context.Background(), *currentModel.App)
	// app not exists
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	err = client.DeleteNetworkPolicy(context.Background(), *currentModel.App)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "List complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}
