package azure

import (
	"fmt"
	"strings"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/onecloud/pkg/cloudprovider"
	"yunion.io/x/onecloud/pkg/compute/models"
)

type SStorage struct {
	zone *SZone

	storageType string
}

func (self *SStorage) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (self *SStorage) GetId() string {
	return fmt.Sprintf("%s/%s", self.zone.GetGlobalId(), strings.ToLower(self.storageType))
}

func (self *SStorage) GetName() string {
	return fmt.Sprintf("%s-%s", self.zone.region.client.providerName, strings.ToLower(self.storageType))
}

func (self *SStorage) GetGlobalId() string {
	return fmt.Sprintf("%s/%s/%s", self.zone.region.GetGlobalId(), self.zone.region.client.subscriptionId, strings.ToLower(self.storageType))
}

func (self *SStorage) IsEmulated() bool {
	return true
}

func (self *SStorage) GetIZone() cloudprovider.ICloudZone {
	return self.zone
}

func (self *SStorage) GetEnabled() bool {
	return true
}

func (self *SStorage) GetCapacityMB() int {
	return 0 // unlimited
}

func (self *SStorage) CreateIDisk(name string, sizeGb int, desc string) (cloudprovider.ICloudDisk, error) {
	if diskId, err := self.zone.region.createDisk(self.storageType, name, int32(sizeGb), desc, ""); err != nil {
		return nil, err
	} else if disk, err := self.zone.region.GetDisk(diskId); err != nil {
		return nil, err
	} else {
		disk.storage = self
		return disk, nil
	}
}

func (self *SStorage) GetIDisk(diskId string) (cloudprovider.ICloudDisk, error) {
	if disk, err := self.zone.region.GetDisk(diskId); err != nil {
		return nil, err
	} else {
		disk.storage = self
		return disk, nil
	}
	return nil, cloudprovider.ErrNotImplemented
}

func (self *SStorage) GetIDisks() ([]cloudprovider.ICloudDisk, error) {
	if disks, err := self.zone.region.GetDisks(); err != nil {
		return nil, err
	} else {
		idisks := make([]cloudprovider.ICloudDisk, 0)
		for i := 0; i < len(disks); i += 1 {
			storageType := strings.ToLower(string(disks[i].Sku.Name))
			if disks[i].Location == self.zone.region.Name && storageType == self.storageType {
				disks[i].storage = self
				idisks = append(idisks, &disks[i])
				log.Debugf("find disk %s for storage %s", disks[i].GetName(), self.GetName())
			}
		}
		return idisks, nil
	}
}

func (self *SStorage) GetIStoragecache() cloudprovider.ICloudStoragecache {
	return self.zone.region.getStoragecache()
}

func (self *SStorage) GetManagerId() string {
	return self.zone.region.client.providerId
}

func (self *SStorage) GetMediumType() string {
	if strings.HasPrefix(self.storageType, "premium") {
		return models.DISK_TYPE_SSD
	}
	return models.DISK_TYPE_ROTATE
}

func (self *SStorage) GetStorageConf() jsonutils.JSONObject {
	conf := jsonutils.NewDict()
	return conf
}

func (self *SStorage) GetStatus() string {
	return models.STORAGE_ONLINE
}

func (self *SStorage) GetStorageType() string {
	return strings.ToLower(self.storageType)
}

func (self *SStorage) Refresh() error {
	// do nothing
	return nil
}
