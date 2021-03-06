// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proto

import (
	"github.com/hazelcast/hazelcast-go-client/serialization"
)

func mapProjectWithPredicateCalculateSize(name string, projection serialization.Data, predicate serialization.Data) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	dataSize += dataCalculateSize(projection)
	dataSize += dataCalculateSize(predicate)
	return dataSize
}

// MapProjectWithPredicateEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func MapProjectWithPredicateEncodeRequest(name string, projection serialization.Data, predicate serialization.Data) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, mapProjectWithPredicateCalculateSize(name, projection, predicate))
	clientMessage.SetMessageType(mapProjectWithPredicate)
	clientMessage.IsRetryable = true
	clientMessage.AppendString(name)
	clientMessage.AppendData(projection)
	clientMessage.AppendData(predicate)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// MapProjectWithPredicateDecodeResponse decodes the given client message.
// It returns a function which returns the response parameters.
func MapProjectWithPredicateDecodeResponse(clientMessage *ClientMessage) func() (response []serialization.Data) {
	// Decode response from client message
	return func() (response []serialization.Data) {
		if clientMessage.IsComplete() {
			return
		}
		responseSize := clientMessage.ReadInt32()
		response = make([]serialization.Data, responseSize)
		for responseIndex := 0; responseIndex < int(responseSize); responseIndex++ {
			if !clientMessage.ReadBool() {
				responseItem := clientMessage.ReadData()
				response[responseIndex] = responseItem
			} else {
				response[responseIndex] = nil
			}
		}
		return
	}
}
