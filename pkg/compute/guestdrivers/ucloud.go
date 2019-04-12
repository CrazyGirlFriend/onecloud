// Copyright 2019 Yunion
//
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

package guestdrivers

import "yunion.io/x/onecloud/pkg/compute/models"

type SUCloudGuestDriver struct {
	SManagedVirtualizedGuestDriver
}

func (self *SUCloudGuestDriver) GetHypervisor() string {
	return models.HYPERVISOR_UCLOUD
}

func (self *SUCloudGuestDriver) GetDefaultSysDiskBackend() string {
	return models.STORAGE_UCLOUD_CLOUD_SSD
}

func (self *SUCloudGuestDriver) GetDeployStatus() ([]string, error) {
	return []string{models.VM_READY}, nil
}

func (self *SUCloudGuestDriver) GetMinimalSysDiskSizeGb() int {
	return 10
}

func (self *SUCloudGuestDriver) GetGuestInitialStateAfterCreate() string {
	return models.VM_RUNNING
}

func (self *SUCloudGuestDriver) GetDetachDiskStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func (self *SUCloudGuestDriver) GetAttachDiskStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func init() {
	driver := SUCloudGuestDriver{}
	models.RegisterGuestDriver(&driver)
}
