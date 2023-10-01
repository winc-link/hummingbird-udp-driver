/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package driver

import (
	"context"
	"fmt"
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/model"
	"github.com/winc-link/hummingbird-sdk-go/service"
	"github.com/winc-link/hummingbird-udp-driver/internal/device"
	"github.com/winc-link/hummingbird-udp-driver/internal/server"
)

type UdpProtocolDriver struct {
	sd *service.DriverService
}

// CloudPluginNotify 云插件启动/停止通知
func (dr UdpProtocolDriver) CloudPluginNotify(ctx context.Context, t commons.CloudPluginNotifyType, name string) error {
	//TODO implement me
	panic("implement me")
}

// DeviceNotify 设备添加/修改/删除通知
func (dr UdpProtocolDriver) DeviceNotify(ctx context.Context, t commons.DeviceNotifyType, deviceId string, device model.Device) error {
	//TODO implement me
	panic("implement me")
}

// ProductNotify 产品添加/修改/删除通知
func (dr UdpProtocolDriver) ProductNotify(ctx context.Context, t commons.ProductNotifyType, productId string, product model.Product) error {
	//TODO implement me
	panic("implement me")
}

// Stop 蜂鸟物联网平台通知
func (dr UdpProtocolDriver) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

// HandlePropertySet 设备属性设置
func (dr UdpProtocolDriver) HandlePropertySet(ctx context.Context, deviceId string, data model.PropertySet) error {
	_, ok := dr.sd.GetDeviceById(deviceId)
	if !ok {
		return fmt.Errorf(" device [%s] not found", deviceId)
	}
	return nil
}

// HandlePropertyGet 设备属性查询
func (dr UdpProtocolDriver) HandlePropertyGet(ctx context.Context, deviceId string, data model.PropertyGet) error {
	//TODO implement me
	panic("implement me")
}

// HandleServiceExecute 设备服务调用
func (dr UdpProtocolDriver) HandleServiceExecute(ctx context.Context, deviceId string, data model.ServiceExecuteRequest) error {
	//TODO implement me
	panic("implement me")
}

// NewUdpProtocolDriver Udp协议驱动
func NewUdpProtocolDriver(ctx context.Context, sd *service.DriverService) *UdpProtocolDriver {
	loadDevices(sd)
	go server.GetUdpServer().Start(sd)
	go cancel(sd, ctx)
	return &UdpProtocolDriver{
		sd: sd,
	}
}

// loadDevices 获取所有已经创建成功的设备，保存在内存中。
func loadDevices(sd *service.DriverService) {
	for _, dev := range sd.GetDeviceList() {
		device.NewDevice(dev.Id, dev.DeviceSn, dev.ProductId, dev.Status == commons.DeviceOnline)
	}
}

// cancel 监听驱动退出，如果驱动退出则把此驱动关联的设备设置成离线
func cancel(sd *service.DriverService, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			for _, dev := range device.GetAllDevice() {
				sd.Offline(dev.GetDeviceId())
			}
		}
	}
}
