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

func TestAccIosxeBGPIPv4UnicastVRFNeighbor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "remote_as", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "description", "BGP Neighbor Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "shutdown", "false"))
	if os.Getenv("IOSXE176") != "" || os.Getenv("IOSXE179") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "cluster_id", "2.2.2.2"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "log_neighbor_changes_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "password_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "password", "LINE"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "timers_keepalive_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "timers_holdtime", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "timers_minimum_neighbor_hold", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "version", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_default_route_map", "RMAP"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_bfd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_bfd_single_hop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_bfd_check_control_plane_failure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_bfd_strict_mode", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "fall_over_maximum_metric_route_map", "ROUTEMAP"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "update_source_loopback", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "activate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "send_community", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_reflector_client", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "soft_reconfiguration", "inbound"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "default_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "default_originate_route_map", "RM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.in_out", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "route_maps.0.route_map_name", "RM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "ha_mode_graceful_restart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "next_hop_self_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_ipv4_unicast_vrf_neighbor.test", "advertisement_interval", "300"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig + testAccIosxeBGPIPv4UnicastVRFNeighborConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_bgp_ipv4_unicast_vrf_neighbor.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast/vrf=VRF1/ipv4-unicast/neighbor=3.3.3.3",
			},
		},
	})
}

const testAccIosxeBGPIPv4UnicastVRFNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"rd" = "1:1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/with-vrf/ipv4=unicast"
	attributes = {
		"af-name" = "unicast"
	}
	lists = [
		{
			name = "vrf"
			key = "name"
			items = [
				{
					"name" = "VRF1"
				},
			] 
		},
	]
	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
}

resource "iosxe_restconf" "PreReq3" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
	}
}

`

func testAccIosxeBGPIPv4UnicastVRFNeighborConfig_minimum() string {
	config := `resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPIPv4UnicastVRFNeighborConfig_all() string {
	config := `resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	remote_as = "65000"` + "\n"
	config += `	description = "BGP Neighbor Description"` + "\n"
	config += `	shutdown = false` + "\n"
	if os.Getenv("IOSXE176") != "" || os.Getenv("IOSXE179") != "" {
		config += `	cluster_id = "2.2.2.2"` + "\n"
	}
	config += `	log_neighbor_changes_disable = true` + "\n"
	config += `	password_type = 1` + "\n"
	config += `	password = "LINE"` + "\n"
	config += `	timers_keepalive_interval = 30` + "\n"
	config += `	timers_holdtime = 40` + "\n"
	config += `	timers_minimum_neighbor_hold = 30` + "\n"
	config += `	version = 4` + "\n"
	config += `	fall_over_default_route_map = "RMAP"` + "\n"
	config += `	fall_over_bfd = true` + "\n"
	config += `	fall_over_bfd_single_hop = true` + "\n"
	config += `	fall_over_bfd_check_control_plane_failure = true` + "\n"
	config += `	fall_over_bfd_strict_mode = true` + "\n"
	config += `	fall_over_maximum_metric_route_map = "ROUTEMAP"` + "\n"
	config += `	update_source_loopback = "100"` + "\n"
	config += `	activate = true` + "\n"
	config += `	send_community = "both"` + "\n"
	config += `	route_reflector_client = false` + "\n"
	config += `	soft_reconfiguration = "inbound"` + "\n"
	config += `	default_originate = true` + "\n"
	config += `	default_originate_route_map = "RM1"` + "\n"
	config += `	route_maps = [{` + "\n"
	config += `		in_out = "in"` + "\n"
	config += `		route_map_name = "RM1"` + "\n"
	config += `	}]` + "\n"
	config += `	ha_mode_graceful_restart = true` + "\n"
	config += `	next_hop_self = true` + "\n"
	config += `	next_hop_self_all = true` + "\n"
	config += `	advertisement_interval = 300` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"
	return config
}
