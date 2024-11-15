// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podman

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/UCF/terraform-provider-podman/internal/podman"
	"github.com/stretchr/testify/assert"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"scaffolding": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]func() *schema.Provider{
			"podman": func() *schema.Provider {
				return podman.Provider()
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `provider "podman" {
					registry_url = "https://example.com"
					username = "testuser"
					password = "testpassword"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("provider.podman", "registry_url", "https://example.com"),
					resource.TestCheckResourceAttr("provider.podman", "username", "testuser"),
				),
			},
		},
	})
}

