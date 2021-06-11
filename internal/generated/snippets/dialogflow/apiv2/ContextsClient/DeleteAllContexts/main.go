// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START dialogflow_v2_generated_Contexts_DeleteAllContexts_sync]

package main

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func main() {
	ctx := context.Background()
	c, err := dialogflow.NewContextsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.DeleteAllContextsRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAllContexts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END dialogflow_v2_generated_Contexts_DeleteAllContexts_sync]