// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
This module defines an standard table of storage driver. The default storage
driver is sample driver used for testing. If you want to use other storage
plugin, just modify Init() method.

*/

package drivers

import (
	_ "github.com/opensds/opensds/contrib/drivers/ceph"
	"github.com/opensds/opensds/contrib/drivers/sample"
	pb "github.com/opensds/opensds/pkg/dock/proto"
)

type VolumeDriver interface {
	//Any initialization the volume driver does while starting.
	Setup() error
	//Any operation the volume driver does while stoping.
	Unset() error

	CreateVolume(opt *pb.CreateVolumeOpts) (*pb.Volume, error)

	PullVolume(volIdentifier string) (*pb.Volume, error)

	DeleteVolume(volIdentifier string) error

	InitializeConnection(opt *pb.CreateAttachmentOpts) (*pb.VolumeConnection, error)

	CreateSnapshot(opt *pb.CreateVolumeSnapshotOpts) (*pb.VolumeSnapshot, error)

	PullSnapshot(snapIdentifier string) (*pb.VolumeSnapshot, error)

	DeleteSnapshot(opt *pb.DeleteVolumeSnapshotOpts) error

	ListPools() (*[]pb.StoragePool, error)
}

func Init(resourceType string) VolumeDriver {
	switch resourceType {
	// case "ceph":
	// 	return &ceph.Driver{}
	default:
		return &sample.Driver{}
	}
}
