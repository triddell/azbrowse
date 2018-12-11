package armclient

type SubResponse struct {
	Subs []struct {
		ID                   string `json:"id"`
		SubscriptionID       string `json:"subscriptionId"`
		DisplayName          string `json:"displayName"`
		State                string `json:"state"`
		SubscriptionPolicies struct {
			LocationPlacementID string `json:"locationPlacementId"`
			QuotaID             string `json:"quotaId"`
			SpendingLimit       string `json:"spendingLimit"`
		} `json:"subscriptionPolicies"`
	} `json:"value"`
}

type ResourceGroupResponse struct {
	Groups []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Location   string `json:"location"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
		} `json:"properties"`
	} `json:"value"`
}

type ResourceReseponse struct {
	Resources []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Sku  struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku"`
		Kind     string `json:"kind"`
		Location string `json:"location"`
		Tags     struct {
			MsResourceUsage string `json:"ms-resource-usage"`
		} `json:"tags"`
	} `json:"value"`
}

type ProvidersResponse struct {
	Providers []struct {
		ID            string `json:"id"`
		Namespace     string `json:"namespace"`
		Authorization struct {
			ApplicationID    string `json:"applicationId"`
			RoleDefinitionID string `json:"roleDefinitionId"`
		} `json:"authorization,omitempty"`
		ResourceTypes []struct {
			ResourceType string        `json:"resourceType"`
			Locations    []interface{} `json:"locations"`
			APIVersions  []string      `json:"apiVersions"`
		} `json:"resourceTypes"`
		RegistrationState string `json:"registrationState"`
		Authorizations    []struct {
			ApplicationID    string `json:"applicationId"`
			RoleDefinitionID string `json:"roleDefinitionId"`
		} `json:"authorizations,omitempty"`
	} `json:"value"`
}
