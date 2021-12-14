package resource

import "testing"

func assertEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func Test_convertModel(t *testing.T) {
	name := "my-aws-app-3"
	teamowner := "shipa-team"
	framework := "oke-dev-fw"

	currentModel := &Model{
		Name:      &name,
		Teamowner: &teamowner,
		Framework: &framework,
	}

	app := convertModel(currentModel)
	assertEqual(t, app.Name, name)
	assertEqual(t, app.TeamOwner, teamowner)
	assertEqual(t, app.Pool, framework)
}
