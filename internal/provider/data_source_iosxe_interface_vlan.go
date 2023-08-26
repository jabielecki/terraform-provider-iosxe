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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &InterfaceVLANDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceVLANDataSource{}
)

func NewInterfaceVLANDataSource() datasource.DataSource {
	return &InterfaceVLANDataSource{}
}

type InterfaceVLANDataSource struct {
	clients map[string]*restconf.Client
}

func (d *InterfaceVLANDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_vlan"
}

func (d *InterfaceVLANDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface VLAN configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.Int64Attribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"autostate": schema.BoolAttribute{
				MarkdownDescription: "Enable auto-state determination for VLAN",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface specific description",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown the selected interface",
				Computed:            true,
			},
			"ip_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy ARP",
				Computed:            true,
			},
			"ip_redirects": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Redirect messages",
				Computed:            true,
			},
			"ip_unreachables": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Unreachable messages",
				Computed:            true,
			},
			"vrf_forwarding": schema.StringAttribute{
				MarkdownDescription: "Configure forwarding table",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_address_mask": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: "Enable IP processing without an explicit address",
				Computed:            true,
			},
			"ip_dhcp_relay_source_interface": schema.StringAttribute{
				MarkdownDescription: "Set source interface for relayed messages",
				Computed:            true,
			},
			"ip_access_group_in": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_in_enable": schema.BoolAttribute{
				MarkdownDescription: "inbound packets",
				Computed:            true,
			},
			"ip_access_group_out": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_out_enable": schema.BoolAttribute{
				MarkdownDescription: "outbound packets",
				Computed:            true,
			},
			"helper_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a destination address for UDP broadcasts",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"global": schema.BoolAttribute{
							MarkdownDescription: "Helper-address is global",
							Computed:            true,
						},
						"vrf": schema.StringAttribute{
							MarkdownDescription: "VRF name for helper-address (if different from interface VRF)",
							Computed:            true,
						},
					},
				},
			},
			"bfd_template": schema.StringAttribute{
				MarkdownDescription: "BFD template",
				Computed:            true,
			},
			"bfd_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD under the interface",
				Computed:            true,
			},
			"bfd_local_address": schema.StringAttribute{
				MarkdownDescription: "The Source IP address to be used for BFD sessions over this interface.",
				Computed:            true,
			},
			"bfd_interval": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"bfd_interval_min_rx": schema.Int64Attribute{
				MarkdownDescription: "Minimum receive interval capability",
				Computed:            true,
			},
			"bfd_interval_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Multiplier value used to compute holddown",
				Computed:            true,
			},
			"bfd_echo": schema.BoolAttribute{
				MarkdownDescription: "Use echo adjunct as bfd detection mechanism",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable IPv6 on interface",
				Computed:            true,
			},
			"ipv6_mtu": schema.Int64Attribute{
				MarkdownDescription: "Set IPv6 Maximum Transmission Unit",
				Computed:            true,
			},
			"ipv6_nd_ra_suppress_all": schema.BoolAttribute{
				MarkdownDescription: "Suppress all IPv6 RA",
				Computed:            true,
			},
			"ipv6_address_autoconfig_default": schema.BoolAttribute{
				MarkdownDescription: "Insert default route",
				Computed:            true,
			},
			"ipv6_address_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Obtain IPv6 address from DHCP server",
				Computed:            true,
			},
			"ipv6_link_local_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"link_local": schema.BoolAttribute{
							MarkdownDescription: "Use link-local address",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"eui_64": schema.BoolAttribute{
							MarkdownDescription: "Use eui-64 interface identifier",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *InterfaceVLANDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *InterfaceVLANDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceVLANData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = InterfaceVLANData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
