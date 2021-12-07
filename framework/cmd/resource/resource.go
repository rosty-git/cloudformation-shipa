package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/rostislavgit/cloudformation-shipa/framework/internal/shipa"
)

func defaultPoolConfig(name string) *shipa.PoolConfig {
	return &shipa.PoolConfig{
		Name: name,
		Resources: &shipa.PoolResources{
			General: &shipa.PoolGeneral{
				Setup: &shipa.PoolSetup{
					Provisioner: "kubernetes",
				},
				Router: "traefik",
				AppQuota: &shipa.PoolAppQuota{
					Limit: "10",
				},
			},
		},
	}
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := shipa.NewClient(*currentModel.ShipaHost, *currentModel.ShipaToken)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	framework := defaultPoolConfig(*currentModel.Name)
	raw, _ := json.Marshal(framework)
	err = client.CreatePoolConfig(context.Background(), framework)
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

	_, err = client.GetPoolConfig(context.Background(), *currentModel.Name)
	if err != nil {
		return handler.ProgressEvent{}, err
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

	err = client.UpdatePoolConfig(context.Background(), defaultPoolConfig(*currentModel.Name))
	if err != nil {
		return handler.ProgressEvent{}, err
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
	err = client.DeletePool(context.Background(), *currentModel.Name)
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
