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

package server

import (
	"github.com/winc-link/hummingbird-sdk-go/service"
	"net"
	"sync"
)

var GlobalDriverService *service.DriverService

type UdpServer struct {
	ClientCons map[string]*Connect
	Lock       sync.RWMutex
}

type Connect struct {
	Client     *net.UDPAddr
	DeviceInfo interface{}
}

var udpServer = &UdpServer{}

func init() {
	udpServer = &UdpServer{
		ClientCons: map[string]*Connect{},
	}
}

func GetUdpServer() *UdpServer {
	return udpServer
}

func (c *UdpServer) Start(sd *service.DriverService) {
	GlobalDriverService = sd

	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9091,
	})
	if err != nil {
		GlobalDriverService.GetLogger().Error(err)
	}
	defer listener.Close()
	// 2.循环读取数据
	for {
		var data [1024]byte
		_, client, err := listener.ReadFromUDP(data[:])
		if err != nil {
			GlobalDriverService.GetLogger().Error(err)
			break
		}
		//deviceSn := strconv.Itoa(int(data[0]))
		//if udpServer.ClientCons[deviceSn] == nil {
		//	udpServer.ClientCons[deviceSn] = &Connect{
		//		Client: client,
		//	}
		//}

		// 3.回复消息
		listener.WriteToUDP([]byte("recevied success!"), client)
	}

}
