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
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-restconf"
)

const (
	YangPatch = false
)

// IosxeProvider defines the provider implementation.
type IosxeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// IosxeProviderModel describes the provider data model.
type IosxeProviderModel struct {
	Username types.String               `tfsdk:"username"`
	Password types.String               `tfsdk:"password"`
	URL      types.String               `tfsdk:"url"`
	Insecure types.Bool                 `tfsdk:"insecure"`
	Retries  types.Int64                `tfsdk:"retries"`
	Devices  []IosxeProviderModelDevice `tfsdk:"devices"`
}

type IosxeProviderModelDevice struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

func (p *IosxeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "iosxe"
	resp.Version = p.version
}

func (p *IosxeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the IOS-XE device. This can also be set as the IOSXE_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the IOS-XE device. This can also be set as the IOSXE_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco IOS-XE device. Optionally a port can be added with `:12345`. The default port is `443`. This can also be set as the IOSXE_URL environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the IOSXE_INSECURE environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the IOSXE_RETRIES environment variable. Defaults to `10`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 99),
				},
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Device name.",
							Required:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL of the Cisco IOS-XE device.",
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (p *IosxeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config IosxeProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("IOSXE_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("IOSXE_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("IOSXE_URL")
		if url == "" && len(config.Devices) > 0 {
			url = config.Devices[0].URL.ValueString()
		}
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("IOSXE_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("IOSXE_RETRIES")
		if retriesStr == "" {
			retries = 10
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	clients := make(map[string]*restconf.Client)
	c, err := restconf.NewClient(url, username, password, insecure, restconf.MaxRetries(int(retries)), restconf.SkipDiscovery("/restconf", true))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create restconf client:\n\n"+err.Error(),
		)
		return
	}
	clients[""] = c

	for _, device := range config.Devices {
		c, err := restconf.NewClient(device.URL.ValueString(), username, password, insecure, restconf.MaxRetries(int(retries)), restconf.SkipDiscovery("/restconf", true))
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create restconf client:\n\n"+err.Error(),
			)
			return
		}
		clients[device.Name.ValueString()] = c
	}

	resp.DataSourceData = clients
	resp.ResourceData = clients
}

func (p *IosxeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRestconfResource,
		NewAAAResource,
		NewAAAAccountingResource,
		NewAAAAuthenticationResource,
		NewAAAAuthorizationResource,
		NewAccessListExtendedResource,
		NewAccessListStandardResource,
		NewBannerResource,
		NewBFDTemplateSingleHopResource,
		NewBGPResource,
		NewBGPAddressFamilyIPv4Resource,
		NewBGPAddressFamilyIPv4VRFResource,
		NewBGPAddressFamilyIPv6Resource,
		NewBGPAddressFamilyIPv6VRFResource,
		NewBGPAddressFamilyL2VPNResource,
		NewBGPIPv4UnicastNeighborResource,
		NewBGPIPv4UnicastVRFNeighborResource,
		NewBGPIPv6UnicastNeighborResource,
		NewBGPL2VPNEVPNNeighborResource,
		NewBGPNeighborResource,
		NewCDPResource,
		NewClockResource,
		NewCryptoIKEv2Resource,
		NewCryptoIKEv2KeyringResource,
		NewCryptoIKEv2PolicyResource,
		NewCryptoIKEv2ProfileResource,
		NewCryptoIKEv2ProposalResource,
		NewCryptoIPSecProfileResource,
		NewCryptoIPSecTransformSetResource,
		NewCTSResource,
		NewDHCPResource,
		NewEVPNResource,
		NewEVPNInstanceResource,
		NewInterfaceEthernetResource,
		NewInterfaceLoopbackResource,
		NewInterfaceMPLSResource,
		NewInterfaceNVEResource,
		NewInterfaceOSPFResource,
		NewInterfaceOSPFProcessResource,
		NewInterfaceOSPFv3Resource,
		NewInterfacePIMResource,
		NewInterfacePortChannelResource,
		NewInterfacePortChannelSubinterfaceResource,
		NewInterfaceSwitchportResource,
		NewInterfaceTunnelResource,
		NewInterfaceVLANResource,
		NewLoggingResource,
		NewLoggingIPv4HostTransportResource,
		NewLoggingIPv4HostVRFTransportResource,
		NewLoggingIPv6HostTransportResource,
		NewLoggingIPv6HostVRFTransportResource,
		NewMDTSubscriptionResource,
		NewMSDPResource,
		NewMSDPVRFResource,
		NewNTPResource,
		NewOSPFResource,
		NewOSPFVRFResource,
		NewPIMResource,
		NewPIMVRFResource,
		NewPrefixListResource,
		NewRadiusResource,
		NewRadiusServerResource,
		NewRouteMapResource,
		NewServiceResource,
		NewSNMPServerResource,
		NewSNMPServerGroupResource,
		NewSNMPServerUserResource,
		NewStaticRouteResource,
		NewStaticRouteVRFResource,
		NewSystemResource,
		NewTemplateResource,
		NewUsernameResource,
		NewVLANResource,
		NewVLANConfigurationResource,
		NewVRFResource,
	}
}

func (p *IosxeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRestconfDataSource,
		NewAAADataSource,
		NewAAAAccountingDataSource,
		NewAAAAuthenticationDataSource,
		NewAAAAuthorizationDataSource,
		NewAccessListExtendedDataSource,
		NewAccessListStandardDataSource,
		NewBannerDataSource,
		NewBFDTemplateSingleHopDataSource,
		NewBGPDataSource,
		NewBGPAddressFamilyIPv4DataSource,
		NewBGPAddressFamilyIPv4VRFDataSource,
		NewBGPAddressFamilyIPv6DataSource,
		NewBGPAddressFamilyIPv6VRFDataSource,
		NewBGPAddressFamilyL2VPNDataSource,
		NewBGPIPv4UnicastNeighborDataSource,
		NewBGPIPv4UnicastVRFNeighborDataSource,
		NewBGPIPv6UnicastNeighborDataSource,
		NewBGPL2VPNEVPNNeighborDataSource,
		NewBGPNeighborDataSource,
		NewCDPDataSource,
		NewClockDataSource,
		NewCryptoIKEv2DataSource,
		NewCryptoIKEv2KeyringDataSource,
		NewCryptoIKEv2PolicyDataSource,
		NewCryptoIKEv2ProfileDataSource,
		NewCryptoIKEv2ProposalDataSource,
		NewCryptoIPSecProfileDataSource,
		NewCryptoIPSecTransformSetDataSource,
		NewCTSDataSource,
		NewDHCPDataSource,
		NewEVPNDataSource,
		NewEVPNInstanceDataSource,
		NewInterfaceEthernetDataSource,
		NewInterfaceLoopbackDataSource,
		NewInterfaceMPLSDataSource,
		NewInterfaceNVEDataSource,
		NewInterfaceOSPFDataSource,
		NewInterfaceOSPFProcessDataSource,
		NewInterfaceOSPFv3DataSource,
		NewInterfacePIMDataSource,
		NewInterfacePortChannelDataSource,
		NewInterfacePortChannelSubinterfaceDataSource,
		NewInterfaceSwitchportDataSource,
		NewInterfaceTunnelDataSource,
		NewInterfaceVLANDataSource,
		NewLoggingDataSource,
		NewLoggingIPv4HostTransportDataSource,
		NewLoggingIPv4HostVRFTransportDataSource,
		NewLoggingIPv6HostTransportDataSource,
		NewLoggingIPv6HostVRFTransportDataSource,
		NewMDTSubscriptionDataSource,
		NewMSDPDataSource,
		NewMSDPVRFDataSource,
		NewNTPDataSource,
		NewOSPFDataSource,
		NewOSPFVRFDataSource,
		NewPIMDataSource,
		NewPIMVRFDataSource,
		NewPrefixListDataSource,
		NewRadiusDataSource,
		NewRadiusServerDataSource,
		NewRouteMapDataSource,
		NewServiceDataSource,
		NewSNMPServerDataSource,
		NewSNMPServerGroupDataSource,
		NewSNMPServerUserDataSource,
		NewStaticRouteDataSource,
		NewStaticRouteVRFDataSource,
		NewSystemDataSource,
		NewTemplateDataSource,
		NewUsernameDataSource,
		NewVLANDataSource,
		NewVLANConfigurationDataSource,
		NewVRFDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &IosxeProvider{
			version: version,
		}
	}
}
