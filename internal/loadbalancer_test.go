// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
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

package internal

import (
	"testing"

	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/internal/proto"
)

func TestRandomLoadBalancer_NextAddressWithNoMembers(t *testing.T) {
	cs := &clusterService{}
	cs.members.Store(make([]*proto.Member, 0)) // initialize with empty member slice
	lb := core.NewRandomLoadBalancer()
	lb.Init(cs)
	member := lb.Next()
	if member != nil {
		t.Errorf("RandomLoadBalancer should return nil when there are no members.")
	}
}
