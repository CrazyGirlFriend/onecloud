// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package monitor

import (
	time "time"

	"yunion.io/x/onecloud/pkg/apis"
)

// SAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlert.
type SAlert struct {
	apis.SVirtualResourceBase
	apis.SEnabledResourceBase
	// Frequency is evaluate period
	Frequency int64       `json:"frequency"`
	Settings  interface{} `json:"settings"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	UsedBy    string      `json:"used_by"`
	// Silenced       bool
	ExecutionError string `json:"execution_error"`
	// If an alert rule has a configured `For` and the query violates the configured threshold
	// it will first go from `OK` to `Pending`. Going from `OK` to `Pending` will not send any
	// notifications. Once the alert rule has been firing for more than `For` duration, it will
	// change to `Alerting` and send alert notifications.
	For                 int64       `json:"for"`
	EvalData            interface{} `json:"eval_data"`
	State               string      `json:"state"`
	NoDataState         string      `json:"no_data_state"`
	ExecutionErrorState string      `json:"execution_error_state"`
	LastStateChange     time.Time   `json:"last_state_change"`
	StateChanges        int         `json:"state_changes"`
}

// SAlertJointsBase is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertJointsBase.
type SAlertJointsBase struct {
	apis.SVirtualJointResourceBase
	AlertId string `json:"alert_id"`
}

// SAlertnotification is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertnotification.
type SAlertnotification struct {
	SAlertJointsBase
	NotificationId string      `json:"notification_id"`
	State          string      `json:"state"`
	Index          byte        `json:"index"`
	UsedBy         string      `json:"used_by"`
	Params         interface{} `json:"params"`
}

// SDataSource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SDataSource.
type SDataSource struct {
	apis.SStandaloneResourceBase
	Type      string `json:"type"`
	Url       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	IsDefault *bool  `json:"is_default,omitempty"`
}

// SNotification is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SNotification.
type SNotification struct {
	apis.SVirtualResourceBase
	Type                  string      `json:"type"`
	IsDefault             bool        `json:"is_default"`
	SendReminder          bool        `json:"send_reminder"`
	DisableResolveMessage bool        `json:"disable_resolve_message"`
	Frequency             int64       `json:"frequency"`
	Settings              interface{} `json:"settings"`
}
