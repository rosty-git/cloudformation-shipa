package resource

import "testing"

func assertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func Test_convertModel(t *testing.T) {
	app := "my-aws-app"
	restartApp := false
	policyMode := "allow-all"

	currentModel := &Model{
		App:        &app,
		RestartApp: &restartApp,
		Egress: &NetworkPolicyConfig{
			PolicyMode: &policyMode,
		},
	}

	networkPolicy := convertModel(currentModel)
	assertEqual(t, networkPolicy.App, app)
	assertEqual(t, networkPolicy.RestartApp, restartApp)
	assertEqual(t, networkPolicy.Egress.PolicyMode, policyMode)
}
