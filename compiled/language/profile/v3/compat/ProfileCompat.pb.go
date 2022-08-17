//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// profile/ProfileCompat.proto is a deprecated file.

package compat

import (
	v31 "github.com/CodapeWild/dktrace-skywalking-agent/compiled/common/v3"
	v3 "github.com/CodapeWild/dktrace-skywalking-agent/compiled/language/profile/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_profile_ProfileCompat_proto protoreflect.FileDescriptor

var file_profile_ProfileCompat_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x90, 0x02, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x5b, 0x0a, 0x16, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x55, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x26, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x42, 0xbe, 0x01, 0x0a,
	0x3c, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x50, 0x01, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x64, 0x61,
	0x70, 0x65, 0x57, 0x69, 0x6c, 0x64, 0x2f, 0x64, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2d, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0xb8, 0x01, 0x01, 0xaa, 0x02, 0x24, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_profile_ProfileCompat_proto_goTypes = []interface{}{
	(*v3.ProfileTaskCommandQuery)(nil), // 0: skywalking.v3.ProfileTaskCommandQuery
	(*v3.ThreadSnapshot)(nil),          // 1: skywalking.v3.ThreadSnapshot
	(*v3.ProfileTaskFinishReport)(nil), // 2: skywalking.v3.ProfileTaskFinishReport
	(*v31.Commands)(nil),               // 3: skywalking.v3.Commands
}
var file_profile_ProfileCompat_proto_depIdxs = []int32{
	0, // 0: ProfileTask.getProfileTaskCommands:input_type -> skywalking.v3.ProfileTaskCommandQuery
	1, // 1: ProfileTask.collectSnapshot:input_type -> skywalking.v3.ThreadSnapshot
	2, // 2: ProfileTask.reportTaskFinish:input_type -> skywalking.v3.ProfileTaskFinishReport
	3, // 3: ProfileTask.getProfileTaskCommands:output_type -> skywalking.v3.Commands
	3, // 4: ProfileTask.collectSnapshot:output_type -> skywalking.v3.Commands
	3, // 5: ProfileTask.reportTaskFinish:output_type -> skywalking.v3.Commands
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_profile_ProfileCompat_proto_init() }
func file_profile_ProfileCompat_proto_init() {
	if File_profile_ProfileCompat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_ProfileCompat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_ProfileCompat_proto_goTypes,
		DependencyIndexes: file_profile_ProfileCompat_proto_depIdxs,
	}.Build()
	File_profile_ProfileCompat_proto = out.File
	file_profile_ProfileCompat_proto_rawDesc = nil
	file_profile_ProfileCompat_proto_goTypes = nil
	file_profile_ProfileCompat_proto_depIdxs = nil
}
