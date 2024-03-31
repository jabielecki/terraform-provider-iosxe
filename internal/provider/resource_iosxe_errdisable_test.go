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

func TestAccIosxeErrdisable(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_all", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_arp_inspection", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_dhcp_rate_limit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_dtp_flap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_l2ptguard", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_link_flap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_pppoe_ia_rate_limit", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_security_violation_shutdown_vlan", "true"))
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1713") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "detect_cause_loopdetect", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_dtp_flap_max_flaps", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_dtp_flap_time", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_link_flap_max_flaps", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_link_flap_time", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_pagp_flap_max_flaps", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "flap_setting_cause_pagp_flap_time", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_interval", "855"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_arp_inspection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_bpduguard", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_channel_misconfig", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_dhcp_rate_limit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_dtp_flap", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_gbic_invalid", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_inline_power", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_l2ptguard", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_link_flap", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_link_monitor_failure", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_loopback", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_mac_limit", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_pagp_flap", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_port_mode_failure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_pppoe_ia_rate_limit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_psp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_psecure_violation", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_security_violation", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_sfp_config_mismatch", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_storm_control", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_udld", "true"))
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1713") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_errdisable.test", "recovery_cause_loopdetect", "true"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeErrdisableConfig_minimum(),
			},
			{
				Config: testAccIosxeErrdisableConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_errdisable.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/errdisable",
			},
		},
	})
}

func testAccIosxeErrdisableConfig_minimum() string {
	config := `resource "iosxe_errdisable" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeErrdisableConfig_all() string {
	config := `resource "iosxe_errdisable" "test" {` + "\n"
	config += `	detect_cause_all = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	detect_cause_arp_inspection = true` + "\n"
	}
	config += `	detect_cause_dhcp_rate_limit = true` + "\n"
	config += `	detect_cause_dtp_flap = true` + "\n"
	config += `	detect_cause_l2ptguard = true` + "\n"
	config += `	detect_cause_link_flap = true` + "\n"
	config += `	detect_cause_pppoe_ia_rate_limit = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	detect_cause_security_violation_shutdown_vlan = true` + "\n"
	}
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1713") != "" {
		config += `	detect_cause_loopdetect = true` + "\n"
	}
	config += `	flap_setting_cause_dtp_flap_max_flaps = 80` + "\n"
	config += `	flap_setting_cause_dtp_flap_time = 90` + "\n"
	config += `	flap_setting_cause_link_flap_max_flaps = 80` + "\n"
	config += `	flap_setting_cause_link_flap_time = 90` + "\n"
	config += `	flap_setting_cause_pagp_flap_max_flaps = 80` + "\n"
	config += `	flap_setting_cause_pagp_flap_time = 90` + "\n"
	config += `	recovery_interval = 855` + "\n"
	config += `	recovery_cause_all = true` + "\n"
	config += `	recovery_cause_arp_inspection = true` + "\n"
	config += `	recovery_cause_bpduguard = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_channel_misconfig = true` + "\n"
	}
	config += `	recovery_cause_dhcp_rate_limit = true` + "\n"
	config += `	recovery_cause_dtp_flap = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_gbic_invalid = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_inline_power = true` + "\n"
	}
	config += `	recovery_cause_l2ptguard = true` + "\n"
	config += `	recovery_cause_link_flap = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_link_monitor_failure = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_loopback = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_mac_limit = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_pagp_flap = true` + "\n"
	}
	config += `	recovery_cause_port_mode_failure = true` + "\n"
	config += `	recovery_cause_pppoe_ia_rate_limit = true` + "\n"
	config += `	recovery_cause_psp = true` + "\n"
	config += `	recovery_cause_psecure_violation = true` + "\n"
	config += `	recovery_cause_security_violation = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_sfp_config_mismatch = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	recovery_cause_storm_control = true` + "\n"
	}
	config += `	recovery_cause_udld = true` + "\n"
	if os.Getenv("IOSXE179") != "" || os.Getenv("IOSXE1713") != "" {
		config += `	recovery_cause_loopdetect = true` + "\n"
	}
	config += `}` + "\n"
	return config
}
