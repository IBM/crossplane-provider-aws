/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type Distribution string

const (
	Distribution_Random      Distribution = "Random"
	Distribution_ByLogStream Distribution = "ByLogStream"
)

type ExportTaskStatusCode string

const (
	ExportTaskStatusCode_CANCELLED      ExportTaskStatusCode = "CANCELLED"
	ExportTaskStatusCode_COMPLETED      ExportTaskStatusCode = "COMPLETED"
	ExportTaskStatusCode_FAILED         ExportTaskStatusCode = "FAILED"
	ExportTaskStatusCode_PENDING        ExportTaskStatusCode = "PENDING"
	ExportTaskStatusCode_PENDING_CANCEL ExportTaskStatusCode = "PENDING_CANCEL"
	ExportTaskStatusCode_RUNNING        ExportTaskStatusCode = "RUNNING"
)

type OrderBy string

const (
	OrderBy_LogStreamName OrderBy = "LogStreamName"
	OrderBy_LastEventTime OrderBy = "LastEventTime"
)

type QueryStatus string

const (
	QueryStatus_Scheduled QueryStatus = "Scheduled"
	QueryStatus_Running   QueryStatus = "Running"
	QueryStatus_Complete  QueryStatus = "Complete"
	QueryStatus_Failed    QueryStatus = "Failed"
	QueryStatus_Cancelled QueryStatus = "Cancelled"
)
