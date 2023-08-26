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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeInterfaceTunnel(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "name", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_mtu", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_nd_ra_suppress_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_address_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_link_local_addresses.0.address", "fe80::9656:d028:8652:66ba"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_link_local_addresses.0.link_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_addresses.0.prefix", "2005:DB8::/32"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv6_addresses.0.eui_64", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "tunnel_destination_ipv4", "2.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "crypto_ipsec_df_bit", "clear"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv4_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ipv4_address_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_dhcp_relay_source_interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "helper_addresses.0.global", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "helper_addresses.0.vrf", "VRF1"))
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "tunnel_mode_ipsec_ipv4", "true"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "bfd_template", "bfd_template1"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "bfd_enable", "true"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_tunnel.test", "bfd_local_address", "1.2.3.4"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceTunnelPrerequisitesConfig + testAccIosxeInterfaceTunnelConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceTunnelPrerequisitesConfig + testAccIosxeInterfaceTunnelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_tunnel.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Tunnel=90",
			},
		},
	})
}

const testAccIosxeInterfaceTunnelPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
		"address-family/ipv6" = ""
	}
}

`

func testAccIosxeInterfaceTunnelConfig_minimum() string {
	config := `resource "iosxe_interface_tunnel" "test" {` + "\n"
	config += `	name = 90` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceTunnelConfig_all() string {
	config := `resource "iosxe_interface_tunnel" "test" {` + "\n"
	config += `	name = 90` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	ip_unreachables = false` + "\n"
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_mtu = 1300` + "\n"
	config += `	ipv6_nd_ra_suppress_all = true` + "\n"
	config += `	ipv6_address_dhcp = true` + "\n"
	config += `	ipv6_link_local_addresses = [{` + "\n"
	config += `		address = "fe80::9656:d028:8652:66ba"` + "\n"
	config += `		link_local = true` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		prefix = "2005:DB8::/32"` + "\n"
	config += `		eui_64 = true` + "\n"
	config += `	}]` + "\n"
	config += `	tunnel_destination_ipv4 = "2.2.2.2"` + "\n"
	config += `	crypto_ipsec_df_bit = "clear"` + "\n"
	config += `	ipv4_address = "10.1.1.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	config += `	ip_dhcp_relay_source_interface = "Loopback100"` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `	}]` + "\n"
	if os.Getenv("C8000V") != "" {
		config += `	tunnel_mode_ipsec_ipv4 = true` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_template = "bfd_template1"` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_enable = true` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1710") != "" {
		config += `	bfd_local_address = "1.2.3.4"` + "\n"
	}
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
