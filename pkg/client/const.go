// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package client

const (
	// This is used to retrieve cluster details (such as cluster CA details + apiserver hostname) from the environment.
	kubernetesExecInfoVarName = "KUBERNETES_EXEC_INFO"
	// The URL used to connect to IMDS.
	baseImdsURL = "http://169.254.169.254"
	// The default scope used to request AAD tokens from IMDS.
	defaultAKSAADServerScope = "6dae42f8-4368-4678-94ff-3960e28e3630/.default"
	// The default azure config path on Linux systems, used to infer identity-related details.
	defaultLinuxAzureJSONPath = "/etc/kubernetes/azure.json"
	// The default azure config path on Windows systems.
	defaultWindowsAzureJSONPath = "c:\\k\\azure.json"
	// The clientId used to denote Managed Service Identities (MSI).
	managedServiceIdentity = "msi"
	// The template used to specify the AAD login authority in public cloud environments.
	microsoftLoginAuthorityTemplate = "https://login.microsoftonline.com/%s"
)
