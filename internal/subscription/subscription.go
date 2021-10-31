package subscription

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sort"

	"github.com/drewstinnett/azurectx-go/internal/commander"
	"github.com/drewstinnett/azurectx-go/internal/fzf"
)

type Client struct {
	Cmd           commander.Commander
	Subscriptions []Subscription
	FZFInstalled  bool
}

func NewClient(cmdr *commander.Commander) (*Client, error) {
	c := &Client{}
	if cmdr != nil {
		c.Cmd = *cmdr
	}

	// Is fzf installed?
	_, err := exec.LookPath("fzf")
	if err == nil {
		c.FZFInstalled = true
	}

	// Is azure-cli installed?
	// If not, return an error, we really need that azure-cli
	_, err = exec.LookPath("az")
	if err != nil {
		return nil, errors.New("azure-cli is not installed")
	}

	err = c.RefreshSubscriptions()
	if err != nil {
		return nil, err
	}
	return c, nil
}

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

func (c *Client) RefreshSubscriptions() error {
	out, err := c.Cmd.Output("az", "account", "list", "--output=json")
	if err != nil {
		return err
	}
	var s []Subscription

	err = json.Unmarshal(out, &s)
	if err != nil {
		return err
	}
	c.Subscriptions = s
	return nil
}

func (c *Client) GetSubscriptions() ([]Subscription, error) {
	return c.Subscriptions, nil
}

func (c *Client) GetSubscriptionNames() ([]string, error) {
	var names []string
	for _, s := range c.Subscriptions {
		names = append(names, s.Name)
	}
	sort.Strings(names)
	return names, nil
}

func (c *Client) GetCurrentSubscriptionName() (string, error) {
	subs, err := c.GetSubscriptions()
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

func (c *Client) SetCurrentSubscriptionName(subname string) error {
	_, err := exec.Command("az", "account", "set", "-s", subname).Output()
	if err != nil {
		return errors.New("Could not set subscription, does it exist?")
	}
	return nil
}

func (c *Client) PickSubscription() (string, error) {
	subs, err := c.GetSubscriptionNames()
	if err != nil {
		return "", err
	}
	filtered := fzf.WithFilter("fzf", func(in io.WriteCloser) {
		for _, p := range subs {
			fmt.Fprintln(in, p)
		}
	})
	return filtered[0], nil
}
