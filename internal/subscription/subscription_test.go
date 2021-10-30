package subscription_test

import (
	"errors"
	"testing"

	"github.com/drewstinnett/azurectx-go/internal/commander"
	"github.com/drewstinnett/azurectx-go/internal/subscription"
	"github.com/stretchr/testify/require"
)

type TestGoodCommander struct{}

func (r TestGoodCommander) Output(command string, args ...string) ([]byte, error) {
	out := ` [ { "cloudName": "AzureCloud", "id": "abc", "isDefault": false, "name": "Test Subscription 1" }, { "cloudName": "AzureCloud", "id": "def", "isDefault": true, "name": "Test Subscription 2" } ] `
	return []byte(out), nil
}

type TestFailCommander struct{}

func (r TestFailCommander) Output(command string, args ...string) ([]byte, error) {
	out := ``
	return []byte(out), errors.New("Fail Output")
}

func TestRefreshSubscriptions(t *testing.T) {
	var cmd commander.Commander = TestGoodCommander{}
	c, err := subscription.NewClient(&cmd)
	require.NoError(t, err)
	err = c.RefreshSubscriptions()
	require.NoError(t, err)

	// Make sure we got some subs back
	require.Greater(t, len(c.Subscriptions), 0)
}

func TestGetSubscriptionNames(t *testing.T) {
	var cmd commander.Commander = TestGoodCommander{}
	c, err := subscription.NewClient(&cmd)
	require.NoError(t, err)
	names, err := c.GetSubscriptionNames()
	require.NoError(t, err)

	// Make sure we got some subs back
	require.Equal(t, []string{"Test Subscription 1", "Test Subscription 2"}, names)
}

func TestDefaultSubscriptions(t *testing.T) {
	var cmd commander.Commander = TestGoodCommander{}
	c, err := subscription.NewClient(&cmd)
	require.NoError(t, err)

	def, err := c.GetCurrentSubscriptionName()
	require.NoError(t, err)

	require.Equal(t, "Test Subscription 2", def)
}

func TestErrorListing(t *testing.T) {
	var cmd commander.Commander = TestFailCommander{}
	_, err := subscription.NewClient(&cmd)
	require.Error(t, err)

	/*
		def, err := c.GetCurrentSubscriptionName()
		require.NoError(t, err)

		require.Equal(t, "Test Subscription 2", def)
	*/
}
