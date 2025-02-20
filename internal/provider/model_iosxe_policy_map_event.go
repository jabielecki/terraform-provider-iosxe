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
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PolicyMapEvent struct {
	Device       types.String                 `tfsdk:"device"`
	Id           types.String                 `tfsdk:"id"`
	Name         types.String                 `tfsdk:"name"`
	EventType    types.String                 `tfsdk:"event_type"`
	MatchType    types.String                 `tfsdk:"match_type"`
	ClassNumbers []PolicyMapEventClassNumbers `tfsdk:"class_numbers"`
}

type PolicyMapEventData struct {
	Device       types.String                 `tfsdk:"device"`
	Id           types.String                 `tfsdk:"id"`
	Name         types.String                 `tfsdk:"name"`
	EventType    types.String                 `tfsdk:"event_type"`
	MatchType    types.String                 `tfsdk:"match_type"`
	ClassNumbers []PolicyMapEventClassNumbers `tfsdk:"class_numbers"`
}
type PolicyMapEventClassNumbers struct {
	Number        types.Int64                               `tfsdk:"number"`
	Class         types.String                              `tfsdk:"class"`
	ExecutionType types.String                              `tfsdk:"execution_type"`
	ActionNumbers []PolicyMapEventClassNumbersActionNumbers `tfsdk:"action_numbers"`
}
type PolicyMapEventClassNumbersActionNumbers struct {
	Number                                       types.Int64  `tfsdk:"number"`
	PauseReauthentication                        types.Bool   `tfsdk:"pause_reauthentication"`
	Authorize                                    types.Bool   `tfsdk:"authorize"`
	TerminateConfig                              types.String `tfsdk:"terminate_config"`
	ActivateServiceTemplateConfigServiceTemplate types.String `tfsdk:"activate_service_template_config_service_template"`
	AuthenticateUsingMethod                      types.String `tfsdk:"authenticate_using_method"`
	AuthenticateUsingRetries                     types.Int64  `tfsdk:"authenticate_using_retries"`
	AuthenticateUsingRetryTime                   types.Int64  `tfsdk:"authenticate_using_retry_time"`
	AuthenticateUsingPriority                    types.Int64  `tfsdk:"authenticate_using_priority"`
}

func (data PolicyMapEvent) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=%s/event=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.EventType.ValueString())))
}

func (data PolicyMapEventData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=%s/event=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.EventType.ValueString())))
}

// if last path element has a key -> remove it
func (data PolicyMapEvent) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data PolicyMapEvent) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.EventType.IsNull() && !data.EventType.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"event-type", data.EventType.ValueString())
	}
	if !data.MatchType.IsNull() && !data.MatchType.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match-type", data.MatchType.ValueString())
	}
	if len(data.ClassNumbers) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number", []interface{}{})
		for index, item := range data.ClassNumbers {
			if !item.Number.IsNull() && !item.Number.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"number", strconv.FormatInt(item.Number.ValueInt64(), 10))
			}
			if !item.Class.IsNull() && !item.Class.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"class", item.Class.ValueString())
			}
			if !item.ExecutionType.IsNull() && !item.ExecutionType.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"execution-type", item.ExecutionType.ValueString())
			}
			if len(item.ActionNumbers) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number", []interface{}{})
				for cindex, citem := range item.ActionNumbers {
					if !citem.Number.IsNull() && !citem.Number.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"number", strconv.FormatInt(citem.Number.ValueInt64(), 10))
					}
					if !citem.PauseReauthentication.IsNull() && !citem.PauseReauthentication.IsUnknown() {
						if citem.PauseReauthentication.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"pause.reauthentication", map[string]string{})
						}
					}
					if !citem.Authorize.IsNull() && !citem.Authorize.IsUnknown() {
						if citem.Authorize.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"authorize", map[string]string{})
						}
					}
					if !citem.TerminateConfig.IsNull() && !citem.TerminateConfig.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"terminate-config", citem.TerminateConfig.ValueString())
					}
					if !citem.ActivateServiceTemplateConfigServiceTemplate.IsNull() && !citem.ActivateServiceTemplateConfigServiceTemplate.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"activate.service-template-config.service-template", citem.ActivateServiceTemplateConfigServiceTemplate.ValueString())
					}
					if !citem.AuthenticateUsingMethod.IsNull() && !citem.AuthenticateUsingMethod.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"authenticate.using.method", citem.AuthenticateUsingMethod.ValueString())
					}
					if !citem.AuthenticateUsingRetries.IsNull() && !citem.AuthenticateUsingRetries.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"authenticate.using.retries", strconv.FormatInt(citem.AuthenticateUsingRetries.ValueInt64(), 10))
					}
					if !citem.AuthenticateUsingRetryTime.IsNull() && !citem.AuthenticateUsingRetryTime.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"authenticate.using.retry-time", strconv.FormatInt(citem.AuthenticateUsingRetryTime.ValueInt64(), 10))
					}
					if !citem.AuthenticateUsingPriority.IsNull() && !citem.AuthenticateUsingPriority.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class-number"+"."+strconv.Itoa(index)+"."+"action-number"+"."+strconv.Itoa(cindex)+"."+"authenticate.using.priority", strconv.FormatInt(citem.AuthenticateUsingPriority.ValueInt64(), 10))
					}
				}
			}
		}
	}
	return body
}

func (data *PolicyMapEvent) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "event-type"); value.Exists() && !data.EventType.IsNull() {
		data.EventType = types.StringValue(value.String())
	} else {
		data.EventType = types.StringNull()
	}
	if value := res.Get(prefix + "match-type"); value.Exists() && !data.MatchType.IsNull() {
		data.MatchType = types.StringValue(value.String())
	} else {
		data.MatchType = types.StringNull()
	}
	for i := range data.ClassNumbers {
		keys := [...]string{"number"}
		keyValues := [...]string{strconv.FormatInt(data.ClassNumbers[i].Number.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "class-number").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("number"); value.Exists() && !data.ClassNumbers[i].Number.IsNull() {
			data.ClassNumbers[i].Number = types.Int64Value(value.Int())
		} else {
			data.ClassNumbers[i].Number = types.Int64Null()
		}
		if value := r.Get("class"); value.Exists() && !data.ClassNumbers[i].Class.IsNull() {
			data.ClassNumbers[i].Class = types.StringValue(value.String())
		} else {
			data.ClassNumbers[i].Class = types.StringNull()
		}
		if value := r.Get("execution-type"); value.Exists() && !data.ClassNumbers[i].ExecutionType.IsNull() {
			data.ClassNumbers[i].ExecutionType = types.StringValue(value.String())
		} else {
			data.ClassNumbers[i].ExecutionType = types.StringNull()
		}
		for ci := range data.ClassNumbers[i].ActionNumbers {
			keys := [...]string{"number"}
			keyValues := [...]string{strconv.FormatInt(data.ClassNumbers[i].ActionNumbers[ci].Number.ValueInt64(), 10)}

			var cr gjson.Result
			r.Get("action-number").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("number"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].Number.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].Number = types.Int64Value(value.Int())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].Number = types.Int64Null()
			}
			if value := cr.Get("pause.reauthentication"); !data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication.IsNull() {
				if value.Exists() {
					data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication = types.BoolValue(true)
				} else {
					data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication = types.BoolValue(false)
				}
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication = types.BoolNull()
			}
			if value := cr.Get("authorize"); !data.ClassNumbers[i].ActionNumbers[ci].Authorize.IsNull() {
				if value.Exists() {
					data.ClassNumbers[i].ActionNumbers[ci].Authorize = types.BoolValue(true)
				} else {
					data.ClassNumbers[i].ActionNumbers[ci].Authorize = types.BoolValue(false)
				}
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].Authorize = types.BoolNull()
			}
			if value := cr.Get("terminate-config"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].TerminateConfig.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].TerminateConfig = types.StringValue(value.String())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].TerminateConfig = types.StringNull()
			}
			if value := cr.Get("activate.service-template-config.service-template"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].ActivateServiceTemplateConfigServiceTemplate.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].ActivateServiceTemplateConfigServiceTemplate = types.StringValue(value.String())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].ActivateServiceTemplateConfigServiceTemplate = types.StringNull()
			}
			if value := cr.Get("authenticate.using.method"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingMethod.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingMethod = types.StringValue(value.String())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingMethod = types.StringNull()
			}
			if value := cr.Get("authenticate.using.retries"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetries.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetries = types.Int64Value(value.Int())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetries = types.Int64Null()
			}
			if value := cr.Get("authenticate.using.retry-time"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetryTime.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetryTime = types.Int64Value(value.Int())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetryTime = types.Int64Null()
			}
			if value := cr.Get("authenticate.using.priority"); value.Exists() && !data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingPriority.IsNull() {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingPriority = types.Int64Value(value.Int())
			} else {
				data.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingPriority = types.Int64Null()
			}
		}
	}
}

func (data *PolicyMapEventData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "match-type"); value.Exists() {
		data.MatchType = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "class-number"); value.Exists() {
		data.ClassNumbers = make([]PolicyMapEventClassNumbers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyMapEventClassNumbers{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Number = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("class"); cValue.Exists() {
				item.Class = types.StringValue(cValue.String())
			}
			if cValue := v.Get("execution-type"); cValue.Exists() {
				item.ExecutionType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("action-number"); cValue.Exists() {
				item.ActionNumbers = make([]PolicyMapEventClassNumbersActionNumbers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyMapEventClassNumbersActionNumbers{}
					if ccValue := cv.Get("number"); ccValue.Exists() {
						cItem.Number = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("pause.reauthentication"); ccValue.Exists() {
						cItem.PauseReauthentication = types.BoolValue(true)
					} else {
						cItem.PauseReauthentication = types.BoolValue(false)
					}
					if ccValue := cv.Get("authorize"); ccValue.Exists() {
						cItem.Authorize = types.BoolValue(true)
					} else {
						cItem.Authorize = types.BoolValue(false)
					}
					if ccValue := cv.Get("terminate-config"); ccValue.Exists() {
						cItem.TerminateConfig = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("activate.service-template-config.service-template"); ccValue.Exists() {
						cItem.ActivateServiceTemplateConfigServiceTemplate = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("authenticate.using.method"); ccValue.Exists() {
						cItem.AuthenticateUsingMethod = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("authenticate.using.retries"); ccValue.Exists() {
						cItem.AuthenticateUsingRetries = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("authenticate.using.retry-time"); ccValue.Exists() {
						cItem.AuthenticateUsingRetryTime = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("authenticate.using.priority"); ccValue.Exists() {
						cItem.AuthenticateUsingPriority = types.Int64Value(ccValue.Int())
					}
					item.ActionNumbers = append(item.ActionNumbers, cItem)
					return true
				})
			}
			data.ClassNumbers = append(data.ClassNumbers, item)
			return true
		})
	}
}

func (data *PolicyMapEvent) getDeletedItems(ctx context.Context, state PolicyMapEvent) []string {
	deletedItems := make([]string, 0)
	if !state.MatchType.IsNull() && data.MatchType.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/match-type", state.getPath()))
	}
	for i := range state.ClassNumbers {
		stateKeyValues := [...]string{strconv.FormatInt(state.ClassNumbers[i].Number.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.ClassNumbers[i].Number.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.ClassNumbers {
			found = true
			if state.ClassNumbers[i].Number.ValueInt64() != data.ClassNumbers[j].Number.ValueInt64() {
				found = false
			}
			if found {
				if !state.ClassNumbers[i].Class.IsNull() && data.ClassNumbers[j].Class.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/class", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.ClassNumbers[i].ExecutionType.IsNull() && data.ClassNumbers[j].ExecutionType.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/execution-type", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				for ci := range state.ClassNumbers[i].ActionNumbers {
					cstateKeyValues := [...]string{strconv.FormatInt(state.ClassNumbers[i].ActionNumbers[ci].Number.ValueInt64(), 10)}

					cemptyKeys := true
					if !reflect.ValueOf(state.ClassNumbers[i].ActionNumbers[ci].Number.ValueInt64()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.ClassNumbers[j].ActionNumbers {
						found = true
						if state.ClassNumbers[i].ActionNumbers[ci].Number.ValueInt64() != data.ClassNumbers[j].ActionNumbers[cj].Number.ValueInt64() {
							found = false
						}
						if found {
							if !state.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].PauseReauthentication.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/pause/reauthentication", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].Authorize.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].Authorize.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/authorize", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].TerminateConfig.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].TerminateConfig.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/terminate-config", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].ActivateServiceTemplateConfigServiceTemplate.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].ActivateServiceTemplateConfigServiceTemplate.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/activate/service-template-config/service-template", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingMethod.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].AuthenticateUsingMethod.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/authenticate/using/method", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetries.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].AuthenticateUsingRetries.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/authenticate/using/retries", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingRetryTime.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].AuthenticateUsingRetryTime.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/authenticate/using/retry-time", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.ClassNumbers[i].ActionNumbers[ci].AuthenticateUsingPriority.IsNull() && data.ClassNumbers[j].ActionNumbers[cj].AuthenticateUsingPriority.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v/authenticate/using/priority", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							break
						}
					}
					if !found {
						deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v/action-number=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/class-number=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *PolicyMapEvent) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.ClassNumbers {
		keyValues := [...]string{strconv.FormatInt(data.ClassNumbers[i].Number.ValueInt64(), 10)}

		for ci := range data.ClassNumbers[i].ActionNumbers {
			ckeyValues := [...]string{strconv.FormatInt(data.ClassNumbers[i].ActionNumbers[ci].Number.ValueInt64(), 10)}
			if !data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication.IsNull() && !data.ClassNumbers[i].ActionNumbers[ci].PauseReauthentication.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class-number=%v/action-number=%v/pause/reauthentication", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
			if !data.ClassNumbers[i].ActionNumbers[ci].Authorize.IsNull() && !data.ClassNumbers[i].ActionNumbers[ci].Authorize.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class-number=%v/action-number=%v/authorize", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
		}
	}
	return emptyLeafsDelete
}

func (data *PolicyMapEvent) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.MatchType.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match-type", data.getPath()))
	}
	for i := range data.ClassNumbers {
		keyValues := [...]string{strconv.FormatInt(data.ClassNumbers[i].Number.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/class-number=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
