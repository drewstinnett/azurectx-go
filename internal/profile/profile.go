package profile

import (
	"github.com/drewstinnett/azurectx-go/internal/subscription"
)

type Profile struct {
	InstallationId string                      `json:"installationId,omitempty"`
	Subscriptions  []subscription.Subscription `json:"subscriptions,omitempty"`
}
