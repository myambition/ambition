// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: d5b3153b9f
// Version Date: Thu Jul 27 18:20:46 UTC 2017

package handlers

import (
	pb "github.com/myambition/ambition/services/rello/rello-service"
)

// CheckListWebhook implements Service.
func CheckListWebhook(ModelCheckListWebhook pb.Model, ActionCheckListWebhook pb.Action) (*pb.ChecklistUpdate, error) {
	request := pb.ChecklistUpdate{
		Model:  &ModelCheckListWebhook,
		Action: &ActionCheckListWebhook,
	}
	return &request, nil
}