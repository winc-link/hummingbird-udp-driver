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

package client

import (
	"github.com/winc-link/hummingbird-sdk-go/service"
	"log"
	"net"
)

type UdpClient struct {
	sd *service.DriverService
}

func (t *UdpClient) Start() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9091,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 2.发送消息
	conn.Write([]byte(""))
	// 3.接收消息
	data := make([]byte, 1024)
	conn.ReadFromUDP(data)
}
