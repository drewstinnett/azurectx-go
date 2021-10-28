package subscription

import (
	"encoding/json"
	"errors"
	"os/exec"
	"sort"
)

type Subscription struct {
	EnvironmentName  string `json:"environmentName,omitempty"`
	HomeTenantId     string `json:"homeTenantId,omitempty"`
	ID               string `json:"id,omitempty"`
	IsDefault        bool   `json:"isDefault,omitempty"`
	ManagedByTenants []struct {
		TenantId string `json:"tenantId,omitempty"`
	} `json:"managedByTenants,omitempty"`
	Name     string `json:"name,omitempty"`
	State    string `json:"state,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	User     struct {
		Name string `json:"name,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"user,omitempty"`
}

func GetSubscriptions() ([]Subscription, error) {
	// profileFile, err := homedir.Expand("~/.azure/azureProfile.json")
	out, err := exec.Command("az", "account", "list", "--output=json").Output()
	if err != nil {
		return nil, err
	}
	var s []Subscription

	err = json.Unmarshal(out, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func GetSubscriptionId(subname string) (string, error) {
	subs, err := GetSubscriptions()
	if err != nil {
		return "", err
	}
	for _, s := range subs {
		if s.Name == subname {
			return s.ID, nil
		}
	}
	return "", errors.New("No subscription found")
}

func GetSubscriptionNames() ([]string, error) {
	subs, err := GetSubscriptions()
	if err != nil {
		return nil, err
	}
	var names []string
	for _, s := range subs {
		names = append(names, s.Name)
	}
	sort.Strings(names)
	return names, nil
}

func GetCurrentSubscriptionName() (string, error) {
	subs, err := GetSubscriptions()
	if err != nil {
		return "", err
	}
	for _, s := range subs {
		if s.IsDefault {
			return s.Name, nil
		}
	}
	return "", errors.New("No current subscription found")
}

func SetCurrentSubscriptionName(subname string) error {
	_, err := exec.Command("az", "account", "set", "-s", subname).Output()
	if err != nil {
		return err
	}
	return nil
}
