// Copyright 2025 EMQ Technologies Co., Ltd.
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

package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	IOCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kuiper",
		Subsystem: "io",
		Name:      "counter",
		Help:      "counter of IO",
	}, []string{LblType, LblIOType, LblStatusType, LblRuleIDType, LblOpIDType})

	IODurationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "kuiper",
		Subsystem: "io",
		Name:      "duration_hist",
		Help:      "Historgram Duration of IO",
		Buckets:   prometheus.ExponentialBuckets(10, 2, 20), // 10us ~ 5s
	}, []string{LblType, LblIOType, LblRuleIDType, LblOpIDType})
)

func init() {
	prometheus.MustRegister(IOCounter)
	prometheus.MustRegister(IODurationHist)
}
