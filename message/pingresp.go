// Copyright (c) 2014 Dataence, LLC. All rights reserved.
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

package message

// A PINGRESP Packet is sent by the Server to the Client in response to a PINGREQ
// Packet. It indicates that the Server is alive.
type PingrespMessage struct {
	header
}

var _ Message = (*PingrespMessage)(nil)

// NewPingrespMessage creates a new PINGRESP message.
func NewPingrespMessage() *PingrespMessage {
	msg := &PingrespMessage{}
	msg.SetType(PINGRESP)

	return msg
}

func (this *PingrespMessage) Decode(src []byte) (int, error) {
	return this.header.decode(src)
}

func (this *PingrespMessage) Encode(dst []byte) (int, error) {
	return this.header.encode(dst)
}
