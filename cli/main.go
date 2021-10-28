package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/drewstinnett/azurectx-go/internal/fzf"
	"github.com/drewstinnett/azurectx-go/internal/subscription"
	flag "github.com/spf13/pflag"
)

func main() {
	current := flag.BoolP("current", "c", false, "Show current subscription")
	list := flag.BoolP("list", "l", false, "List all subscriptions")

	flag.Parse()

	if *current {
		// Show current subscription
		s, err := subscription.GetCurrentSubscriptionName()
		CheckErr(err)
		fmt.Println(s)
	} else if *list {
		// List all subscriptions
		names, err := subscription.GetSubscriptionNames()
		CheckErr(err)
		for _, name := range names {
			fmt.Println(name)
		}
	} else if len(flag.Args()) == 0 {
		// Set a subscription from the picker
		s, err := fzf.PickSubscription()
		CheckErr(err)
		err = subscription.SetCurrentSubscriptionName(s)
		CheckErr(err)
	} else if len(flag.Args()) > 0 {
		// Set subscription to the argument
		var subName string
		if len(flag.Args()) > 1 {
			subName = strings.Join(flag.Args(), " ")
		} else {
			subName = flag.Args()[0]
		}
		err := subscription.SetCurrentSubscriptionName(subName)
		CheckErr(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
