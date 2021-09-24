// Copyright 2018 StreamSets Inc.
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
package destinations

import (
	_ "github.com/streamsets/datacollector-edge/stages/destinations/azure/eventhubs"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/azure/iothub"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/coap"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/firehose"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/http"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/influxdb"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/kafka"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/kinesis"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/mqtt"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/s3"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/toerror"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/toevent"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/trash"
	_ "github.com/streamsets/datacollector-edge/stages/destinations/websocket"
)
