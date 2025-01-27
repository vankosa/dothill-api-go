/*
 * Copyright (c) 2021 Enix, SAS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 * Authors:
 * Paul Laffitte <paul.laffitte@enix.fr>
 * Arthur Chaloin <arthur.chaloin@enix.fr>
 * Alexandre Buisine <alexandre.buisine@enix.fr>
 * Joe Skazinski <joseph.skazinski@seagate.com>
 */

package dothill

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Login : Called automatically, may be called manually if credentials changed
func (client *Client) Login() error {
	userpass := fmt.Sprintf("%s_%s", client.Username, client.Password)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(userpass)))
	res, _, err := client.FormattedRequest("/login/%s", hash)

	if err != nil {
		return err
	}

	client.sessionKey = res.ObjectsMap["status"].PropertiesMap["response"].Data
	return nil
}

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/create/volume/pool/\"%s\"/size/%s/tier-affinity/no-affinity/\"%s\"", pool, size, name)
}

// CreateHost : creates a host
func (client *Client) CreateHost(name, iqn string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/create/host/initiators/\"%s\"/\"%s\"", iqn, name)
}

// MapVolume : map a volume to host + LUN
func (client *Client) MapVolume(name, host, access string, lun int) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/map/volume/access/%s/lun/%d/initiator/%s/\"%s\"", access, lun, host, name)
}

// ShowVolumes : get informations about volumes
func (client *Client) ShowVolumes(volumes ...string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/show/volumes/\"%s\"", strings.Join(volumes, ","))
}

// UnmapVolume : unmap a volume from host
func (client *Client) UnmapVolume(name, host string) (*Response, *ResponseStatus, error) {
	if host == "" {
		return client.FormattedRequest("/unmap/volume/\"%s\"", name)
	}

	return client.FormattedRequest("/unmap/volume/initiator/\"%s\"/\"%s\"", host, name)
}

// ExpandVolume : extend a volume if there is enough space on the vdisk
func (client *Client) ExpandVolume(name, size string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/expand/volume/size/\"%s\"/\"%s\"", size, name)
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/volumes/\"%s\"", name)
}

// DeleteHost : deletes a host by its ID or nickname
func (client *Client) DeleteHost(name string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/host/\"%s\"", name)
}

// ShowHostMaps : list the volume mappings for given host
// If host is an empty string, mapping for all hosts is shown
func (client *Client) ShowHostMaps(host string) ([]Volume, *ResponseStatus, error) {
	if len(host) > 0 {
		host = fmt.Sprintf("\"%s", host)
	}

	smallhost := strings.Split(host, ":")
	hostids := fmt.Sprintf(smallhost[1] + ".*")
	hostname := fmt.Sprintf("\"%s\"", hostids)

	res, status, err := client.FormattedRequest("/show/maps/%s", hostname)
	if err != nil {
		return nil, status, err
	}

	mappings := make([]Volume, 0)
	for _, rootObj := range res.Objects {
		if rootObj.Name != "hosts-view" {
			continue
		}

		for _, object := range rootObj.Objects {
			if object.Name == "volume-view" {
				vol := Volume{}
				vol.fillFromObject(&object)
				mappings = append(mappings, vol)
			}
		}
	}

	return mappings, status, err
}

// ShowSnapshots : list snapshots
func (client *Client) ShowSnapshots(names ...string) (*Response, *ResponseStatus, error) {
	if len(names) == 0 {
		return client.FormattedRequest("/show/snapshots")
	}
	return client.FormattedRequest("/show/snapshots/%q", strings.Join(names, ","))
}

// CreateSnapshot : create a snapshot in a snap pool and the snap pool if it doesn't exsits
func (client *Client) CreateSnapshot(name string, snapshotName string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/create/snapshots/volumes/%q/%q", name, snapshotName)
}

// DeleteSnapshot : delete a snapshot
func (client *Client) DeleteSnapshot(names ...string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/snapshot/%q", strings.Join(names, ","))
}

// CopyVolume : create an new volume by copying another one or a snapshot
func (client *Client) CopyVolume(sourceName string, destinationName string, pool string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/copy/volume/destination-pool/%q/name/%q/%q", pool, destinationName, sourceName)
}
