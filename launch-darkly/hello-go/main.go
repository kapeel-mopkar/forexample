package main

import (
  "fmt"
  "time"
  
  ld "gopkg.in/launchdarkly/go-client.v4"
)

func main() {
	client, _ := ld.MakeClient("sdk-ba54d7ce-dd5c-40b1-89a6-7626d1ddc98a", 5*time.Second)

	key := "bob@example.com"
	first := "Bob"
	last := "Loblaw"
	custom := map[string]interface{}{"groups": "beta_testers"}

	user := ld.User{
	Key: &key,
	FirstName: &first,
	LastName:  &last,
	Custom:    &custom,
	}

	showFeature, _ := client.BoolVariation("new-search-bar", user, false)

	if showFeature {
		// application code to show the feature
		fmt.Println("Showing your feature to " + *user.Key)
	} else {
		// the code to run if the feature is off
		fmt.Println("Not showing your feature to " + *user.Key)
	}

	client.Close()

}