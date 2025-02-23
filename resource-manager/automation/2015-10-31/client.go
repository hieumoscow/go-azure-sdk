package v2015_10_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/activity"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/agentregistrationinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/connection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/connectiontype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/credential"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/dsccompilationjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/dscconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/dscnode"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/dscnodeconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/jobschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/jobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/linkedworkspace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/module"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/nodereports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/objectdatatypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/runbook"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/runbookdraft"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/testjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/testjobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/typefields"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/variable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/watcher"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/webhook"
)

type Client struct {
	Activity                     *activity.ActivityClient
	AgentRegistrationInformation *agentregistrationinformation.AgentRegistrationInformationClient
	AutomationAccount            *automationaccount.AutomationAccountClient
	Certificate                  *certificate.CertificateClient
	Connection                   *connection.ConnectionClient
	ConnectionType               *connectiontype.ConnectionTypeClient
	Credential                   *credential.CredentialClient
	DscCompilationJob            *dsccompilationjob.DscCompilationJobClient
	DscConfiguration             *dscconfiguration.DscConfigurationClient
	DscNode                      *dscnode.DscNodeClient
	DscNodeConfiguration         *dscnodeconfiguration.DscNodeConfigurationClient
	HybridRunbookWorkerGroup     *hybridrunbookworkergroup.HybridRunbookWorkerGroupClient
	Job                          *job.JobClient
	JobSchedule                  *jobschedule.JobScheduleClient
	JobStream                    *jobstream.JobStreamClient
	LinkedWorkspace              *linkedworkspace.LinkedWorkspaceClient
	ListKeys                     *listkeys.ListKeysClient
	Module                       *module.ModuleClient
	NodeReports                  *nodereports.NodeReportsClient
	ObjectDataTypes              *objectdatatypes.ObjectDataTypesClient
	Runbook                      *runbook.RunbookClient
	RunbookDraft                 *runbookdraft.RunbookDraftClient
	Schedule                     *schedule.ScheduleClient
	Statistics                   *statistics.StatisticsClient
	TestJob                      *testjob.TestJobClient
	TestJobStream                *testjobstream.TestJobStreamClient
	TypeFields                   *typefields.TypeFieldsClient
	Usages                       *usages.UsagesClient
	Variable                     *variable.VariableClient
	Watcher                      *watcher.WatcherClient
	Webhook                      *webhook.WebhookClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityClient := activity.NewActivityClientWithBaseURI(endpoint)
	configureAuthFunc(&activityClient.Client)

	agentRegistrationInformationClient := agentregistrationinformation.NewAgentRegistrationInformationClientWithBaseURI(endpoint)
	configureAuthFunc(&agentRegistrationInformationClient.Client)

	automationAccountClient := automationaccount.NewAutomationAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&automationAccountClient.Client)

	certificateClient := certificate.NewCertificateClientWithBaseURI(endpoint)
	configureAuthFunc(&certificateClient.Client)

	connectionClient := connection.NewConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionClient.Client)

	connectionTypeClient := connectiontype.NewConnectionTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionTypeClient.Client)

	credentialClient := credential.NewCredentialClientWithBaseURI(endpoint)
	configureAuthFunc(&credentialClient.Client)

	dscCompilationJobClient := dsccompilationjob.NewDscCompilationJobClientWithBaseURI(endpoint)
	configureAuthFunc(&dscCompilationJobClient.Client)

	dscConfigurationClient := dscconfiguration.NewDscConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscConfigurationClient.Client)

	dscNodeClient := dscnode.NewDscNodeClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeClient.Client)

	dscNodeConfigurationClient := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeConfigurationClient.Client)

	hybridRunbookWorkerGroupClient := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerGroupClient.Client)

	jobClient := job.NewJobClientWithBaseURI(endpoint)
	configureAuthFunc(&jobClient.Client)

	jobScheduleClient := jobschedule.NewJobScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&jobScheduleClient.Client)

	jobStreamClient := jobstream.NewJobStreamClientWithBaseURI(endpoint)
	configureAuthFunc(&jobStreamClient.Client)

	linkedWorkspaceClient := linkedworkspace.NewLinkedWorkspaceClientWithBaseURI(endpoint)
	configureAuthFunc(&linkedWorkspaceClient.Client)

	listKeysClient := listkeys.NewListKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&listKeysClient.Client)

	moduleClient := module.NewModuleClientWithBaseURI(endpoint)
	configureAuthFunc(&moduleClient.Client)

	nodeReportsClient := nodereports.NewNodeReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&nodeReportsClient.Client)

	objectDataTypesClient := objectdatatypes.NewObjectDataTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&objectDataTypesClient.Client)

	runbookClient := runbook.NewRunbookClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookClient.Client)

	runbookDraftClient := runbookdraft.NewRunbookDraftClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookDraftClient.Client)

	scheduleClient := schedule.NewScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduleClient.Client)

	statisticsClient := statistics.NewStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&statisticsClient.Client)

	testJobClient := testjob.NewTestJobClientWithBaseURI(endpoint)
	configureAuthFunc(&testJobClient.Client)

	testJobStreamClient := testjobstream.NewTestJobStreamClientWithBaseURI(endpoint)
	configureAuthFunc(&testJobStreamClient.Client)

	typeFieldsClient := typefields.NewTypeFieldsClientWithBaseURI(endpoint)
	configureAuthFunc(&typeFieldsClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	variableClient := variable.NewVariableClientWithBaseURI(endpoint)
	configureAuthFunc(&variableClient.Client)

	watcherClient := watcher.NewWatcherClientWithBaseURI(endpoint)
	configureAuthFunc(&watcherClient.Client)

	webhookClient := webhook.NewWebhookClientWithBaseURI(endpoint)
	configureAuthFunc(&webhookClient.Client)

	return Client{
		Activity:                     &activityClient,
		AgentRegistrationInformation: &agentRegistrationInformationClient,
		AutomationAccount:            &automationAccountClient,
		Certificate:                  &certificateClient,
		Connection:                   &connectionClient,
		ConnectionType:               &connectionTypeClient,
		Credential:                   &credentialClient,
		DscCompilationJob:            &dscCompilationJobClient,
		DscConfiguration:             &dscConfigurationClient,
		DscNode:                      &dscNodeClient,
		DscNodeConfiguration:         &dscNodeConfigurationClient,
		HybridRunbookWorkerGroup:     &hybridRunbookWorkerGroupClient,
		Job:                          &jobClient,
		JobSchedule:                  &jobScheduleClient,
		JobStream:                    &jobStreamClient,
		LinkedWorkspace:              &linkedWorkspaceClient,
		ListKeys:                     &listKeysClient,
		Module:                       &moduleClient,
		NodeReports:                  &nodeReportsClient,
		ObjectDataTypes:              &objectDataTypesClient,
		Runbook:                      &runbookClient,
		RunbookDraft:                 &runbookDraftClient,
		Schedule:                     &scheduleClient,
		Statistics:                   &statisticsClient,
		TestJob:                      &testJobClient,
		TestJobStream:                &testJobStreamClient,
		TypeFields:                   &typeFieldsClient,
		Usages:                       &usagesClient,
		Variable:                     &variableClient,
		Watcher:                      &watcherClient,
		Webhook:                      &webhookClient,
	}
}
