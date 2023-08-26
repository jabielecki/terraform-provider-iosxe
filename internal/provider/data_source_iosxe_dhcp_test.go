// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeDHCP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_dhcp.test", "relay_information_trust_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_dhcp.test", "relay_information_option_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_dhcp.test", "relay_information_option_vpn", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_dhcp.test", "snooping", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_dhcp.test", "snooping_vlans.0.vlan_id", "3-4"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeDHCPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeDHCPConfig() string {
	config := `resource "iosxe_dhcp" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	relay_information_trust_all = false` + "\n"
	config += `	relay_information_option_default = false` + "\n"
	config += `	relay_information_option_vpn = true` + "\n"
	config += `	snooping = true` + "\n"
	config += `	snooping_vlans = [{` + "\n"
	config += `		vlan_id = "3-4"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_dhcp" "test" {
			depends_on = [iosxe_dhcp.test]
		}
	`
	return config
}
