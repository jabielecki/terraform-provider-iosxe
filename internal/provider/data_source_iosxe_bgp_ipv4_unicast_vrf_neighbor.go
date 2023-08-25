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
	_ datasource.DataSource              = &BGPIPv4UnicastVRFNeighborDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPIPv4UnicastVRFNeighborDataSource{}
)

func NewBGPIPv4UnicastVRFNeighborDataSource() datasource.DataSource {
	return &BGPIPv4UnicastVRFNeighborDataSource{}
}

type BGPIPv4UnicastVRFNeighborDataSource struct {
	clients map[string]*restconf.Client
}

func (d *BGPIPv4UnicastVRFNeighborDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_ipv4_unicast_vrf_neighbor"
}

func (d *BGPIPv4UnicastVRFNeighborDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP IPv4 Unicast VRF Neighbor configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"asn": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"remote_as": schema.StringAttribute{
				MarkdownDescription: "Specify a BGP peer-group remote-as",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Neighbor specific description",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administratively shut down this neighbor",
				Computed:            true,
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"log_neighbor_changes_disable": schema.BoolAttribute{
				MarkdownDescription: "disable",
				Computed:            true,
			},
			"password_enctype": schema.Int64Attribute{
				MarkdownDescription: "Encryption type (0 to disable encryption, 7 for proprietary)",
				Computed:            true,
			},
			"password_text": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"timers_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"timers_holdtime": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"timers_minimum_neighbor_hold": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "Set the BGP version to match a neighbor",
				Computed:            true,
			},
			"fall_over_default_route_map": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"fall_over_bfd_multi_hop": schema.BoolAttribute{
				MarkdownDescription: "Force BFD multi-hop to detect failure",
				Computed:            true,
			},
			"fall_over_bfd_single_hop": schema.BoolAttribute{
				MarkdownDescription: "Force BFD single-hop to detect failure",
				Computed:            true,
			},
			"fall_over_bfd_check_control_plane_failure": schema.BoolAttribute{
				MarkdownDescription: "Retrieve control plane dependent failure info from BFD for BGP GR/NSR operation",
				Computed:            true,
			},
			"fall_over_bfd_strict_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD strict-mode",
				Computed:            true,
			},
			"fall_over_maximum_metric_route_map": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"disable_connected_check": schema.BoolAttribute{
				MarkdownDescription: "one-hop away EBGP peer using loopback address",
				Computed:            true,
			},
			"ttl_security_hops": schema.Int64Attribute{
				MarkdownDescription: "IP hops",
				Computed:            true,
			},
			"local_as_as_no": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"local_as_no_prepend": schema.BoolAttribute{
				MarkdownDescription: "Do not prepend local-as to updates from ebgp peers",
				Computed:            true,
			},
			"local_as_replace_as": schema.BoolAttribute{
				MarkdownDescription: "Replace real AS with local AS in the EBGP updates",
				Computed:            true,
			},
			"local_as_dual_as": schema.BoolAttribute{
				MarkdownDescription: "Accept either real AS or local AS from the ebgp peer",
				Computed:            true,
			},
			"update_source_loopback": schema.StringAttribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"activate": schema.BoolAttribute{
				MarkdownDescription: "Enable the address family for this neighbor",
				Computed:            true,
			},
			"send_community": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"route_reflector_client": schema.BoolAttribute{
				MarkdownDescription: "Configure a neighbor as Route Reflector client",
				Computed:            true,
			},
			"route_maps": schema.ListNestedAttribute{
				MarkdownDescription: "Apply route map to neighbor",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"in_out": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"route_map_name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"ebgp_multihop": schema.BoolAttribute{
				MarkdownDescription: "Allow EBGP neighbors not on directly connected networks. For single-hop ebgp peers, delete ebgp-multihop directly.",
				Computed:            true,
			},
			"ebgp_multihop_max_hop": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
	}
}

func (d *BGPIPv4UnicastVRFNeighborDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *BGPIPv4UnicastVRFNeighborDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BGPIPv4UnicastVRFNeighborData

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
		config = BGPIPv4UnicastVRFNeighborData{Device: config.Device}
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
