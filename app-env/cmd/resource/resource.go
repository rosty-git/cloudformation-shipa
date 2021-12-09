package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/rostislavgit/cloudformation-shipa/app-env/internal/shipa"
)

func convertModel(currentModel *Model) *shipa.CreateAppEnv {
	var envs []*shipa.AppEnv
	for _, env := range currentModel.Envs {
		envs = append(envs, &shipa.AppEnv{
			Name:  *env.Name,
			Value: *env.Value,
		})
	}

	return &shipa.CreateAppEnv{
		App:       *currentModel.App,
		Envs:      envs,
		NoRestart: *currentModel.Norestart,
		Private:   *currentModel.Private,
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

	appEnv := convertModel(currentModel)
	raw, _ := json.Marshal(appEnv)
	err = client.CreateAppEnvs(context.Background(), appEnv)
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

	appEnv, err := client.GetAppEnvs(context.Background(), *currentModel.App)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	for _, env := range appEnv {
		currentModel.Envs = append(currentModel.Envs, Env{
			Name:  &env.Name,
			Value: &env.Value,
		})
	}

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

	appEnv := convertModel(currentModel)
	raw, _ := json.Marshal(appEnv)
	// TODO: update
	err = client.CreateAppEnvs(context.Background(), appEnv)
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

	appEnv := convertModel(currentModel)

	err = client.DeleteAppEnvs(context.Background(), appEnv)
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
