package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/rostislavgit/cloudformation-shipa/app-cname/internal/shipa"
)

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

	appCname := &shipa.AppCname{
		App:    *currentModel.App,
		Cname:  *currentModel.Cname,
		Scheme: *currentModel.Scheme,
	}

	raw, _ := json.Marshal(appCname)
	err = client.CreateAppCname(context.Background(), appCname)
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

	app, err := client.GetApp(context.Background(), *currentModel.App)
	// app not exists
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if len(app.Entrypoints) > 0 {
		last := app.Entrypoints[len(app.Entrypoints)-1]
		currentModel.Cname = &last.Cname
		currentModel.Scheme = &last.Scheme
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

	appCname := &shipa.AppCname{
		App:    *currentModel.App,
		Cname:  *currentModel.Cname,
		Scheme: *currentModel.Scheme,
	}

	raw, _ := json.Marshal(appCname)
	err = client.UpdateAppCname(context.Background(), appCname)
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

	appCname := &shipa.AppCname{
		App:    *currentModel.App,
		Cname:  *currentModel.Cname,
		Scheme: *currentModel.Scheme,
	}

	err = client.DeleteAppCname(context.Background(), appCname)
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
