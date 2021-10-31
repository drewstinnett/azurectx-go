package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/drewstinnett/azurectx-go/internal/commander"
	"github.com/drewstinnett/azurectx-go/internal/subscription"
	flag "github.com/spf13/pflag"
)

func main() {
	current := flag.BoolP("current", "c", false, "Show current subscription")
	list := flag.BoolP("list", "l", false, "List all subscriptions")

	flag.Parse()

	var cmdr commander.Commander = commander.RealCommander{}
	c, err := subscription.NewClient(&cmdr)
	if err != nil {
		log.Fatal(err)
	}

	if len(c.Subscriptions) == 0 {
		log.Fatal("No subscriptions found, make sure you are logged in")
	}

	if *current {
		// Show current subscription
		s, err := c.GetCurrentSubscriptionName()
		CheckErr(err)
		fmt.Println(s)
	} else if *list {
		// List all subscriptions
		names, err := c.GetSubscriptionNames()
		CheckErr(err)
		for _, name := range names {
			fmt.Println(name)
		}
	} else if len(flag.Args()) == 0 {
		// Set a subscription from the picker
		if c.FZFInstalled {
			s, err := c.PickSubscription()
			CheckErr(err)
			err = c.SetCurrentSubscriptionName(s)
			CheckErr(err)
			fmt.Printf("Switched to '%v'", s)
		} else {
			// But if no fzf, just list them
			names, err := c.GetSubscriptionNames()
			CheckErr(err)
			for _, name := range names {
				fmt.Println(name)
			}
		}
	} else if len(flag.Args()) > 0 {
		// Set subscription to the argument
		var subName string
		if len(flag.Args()) > 1 {
			subName = strings.Join(flag.Args(), " ")
		} else {
			subName = flag.Args()[0]
		}
		err := c.SetCurrentSubscriptionName(subName)
		CheckErr(err)
		fmt.Printf("Switched to '%v'", subName)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
