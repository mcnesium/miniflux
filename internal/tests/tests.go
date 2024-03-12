// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package tests

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	miniflux "miniflux.app/v2/client"
)

const (
	testBaseURL           = "http://127.0.0.1:8080/"
	testAdminUsername     = "admin"
	testAdminPassword     = "test123"
	testStandardPassword  = "secret"
	testFeedURL           = "https://miniflux.app/feed.xml"
	testFeedTitle         = "Miniflux"
	testSubscriptionTitle = "Miniflux Releases"
	testWebsiteURL        = "https://miniflux.app"
)

func getRandomUsername() string {
	return fmt.Sprintf("user%10d", rand.Intn(math.MaxInt64))
}

func createClient(t *testing.T) *miniflux.Client {
	username := getRandomUsername()
	client := miniflux.New(testBaseURL, testAdminUsername, testAdminPassword)
	_, err := client.CreateUser(username, testStandardPassword, false)
	if err != nil {
		t.Fatal(err)
	}

	return miniflux.New(testBaseURL, username, testStandardPassword)
}

func createFeed(t *testing.T, client *miniflux.Client) (*miniflux.Feed, *miniflux.Category) {
	categories, err := client.Categories()
	if err != nil {
		t.Fatal(err)
	}

	feedID, err := client.CreateFeed(&miniflux.FeedCreationRequest{
		FeedURL:    testFeedURL,
		CategoryID: categories[0].ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if feedID == 0 {
		t.Fatalf(`Invalid feed ID, got %q`, feedID)
	}

	feed, err := client.Feed(feedID)
	if err != nil {
		t.Fatal(err)
	}

	return feed, categories[0]
}
