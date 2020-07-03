/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dsl

import (
	"context"

	"mosn.io/api"
	"mosn.io/mosn/pkg/cel/attribute"
	"mosn.io/pkg/buffer"
)

type DSLFilter struct {
	context               context.Context
	receiverFilterHandler api.StreamReceiverFilterHandler
	senderHandler         api.StreamSenderFilterHandler

	dsl *DSL
}

// NewDSLFilter used to create new dsl filter
func NewDSLFilter(ctx context.Context, dsl *DSL) *DSLFilter {
	filter := &DSLFilter{
		context: ctx,
		dsl:     dsl,
	}
	return filter
}

func (f *DSLFilter) OnReceive(ctx context.Context, headers api.HeaderMap, buf buffer.IoBuffer, trailers api.HeaderMap) api.StreamFilterStatus {
	println("r00000000000", f.dsl.LogDSL)
	println("r1111111111", ctx == nil)
	bag := attribute.NewMutableBag(nil)
	bag.Set("ctx", ctx)
	out, _ := f.dsl.LogDSL.Evaluate(bag)
	println("r0000000000000", out)

	return api.StreamFilterContinue
}

func (f *DSLFilter) Append(ctx context.Context, headers api.HeaderMap, buf buffer.IoBuffer, trailers api.HeaderMap) api.StreamFilterStatus {

	return api.StreamFilterContinue
}

func (f *DSLFilter) SetReceiveFilterHandler(handler api.StreamReceiverFilterHandler) {
	f.receiverFilterHandler = handler
}

func (f *DSLFilter) SetSenderFilterHandler(handler api.StreamSenderFilterHandler) {
	f.senderHandler = handler
}

func (f *DSLFilter) OnDestroy() {}

func (f *DSLFilter) Log(ctx context.Context, reqHeaders api.HeaderMap, respHeaders api.HeaderMap, requestInfo api.RequestInfo) {
	//if reqHeaders == nil || respHeaders == nil || requestInfo == nil || f.metrics == nil || len(f.metrics.definitions) == 0 {
	//	return
	//}

	//attributes := ExtractAttributes(reqHeaders, respHeaders, requestInfo, f.requestTotalSize, time.Now())
	//stats, err := f.metrics.Stat(attributes)
	//if err != nil {
	//	log.DefaultLogger.Errorf("stats error: %s", err.Error())
	//	return
	//}
	println("00000000000", f.dsl.LogDSL)
	println("1111111111", ctx == nil)
	bag := attribute.NewMutableBag(nil)
	bag.Set("ctx", ctx)
	out, _ := f.dsl.LogDSL.Evaluate(bag)
	println("0000000000000", out.(bool))
}
