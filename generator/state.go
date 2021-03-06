// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"bytes"
	"text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const stateTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import * as ProtocTypes from './protoc_types_pb';

export interface ProtocState { {{range $i, $entity := .}}
  {{$entity.FieldName}}: {
    isLoading: boolean;
    error: { code: string; message: string; };
    {{if $entity.Repeated}}value: ProtocTypes.{{$entity.FullTypeName}}.AsObject[];
    {{else}}value: ProtocTypes.{{$entity.FullTypeName}}.AsObject | null;{{end}}
  },
  {{end}}
}

export const initialProtocState : ProtocState = { {{range $i, $entity := .}}
  {{$entity.FieldName}}: {
    isLoading: false,
    error: null,
    {{if $entity.Repeated}}value: [],
    {{else}}value: null,{{end}}
  },
  {{end}}
}
`

type StateEntity struct {
	FieldName    string
	FullTypeName string
	Repeated     bool
}

func CreateStateFile(stateFields []*gp.FieldDescriptorProto, debug bool) (*File, error) {
	stateEntities := []*StateEntity{}

	// transform stateFields into our StateEntity implementation so template can read values
	for _, entity := range stateFields {
		stateEntities = append(stateEntities, &StateEntity{
			FieldName:    entity.GetJsonName(),
			FullTypeName: CreatePackageAndTypeString(entity.GetTypeName()),
			Repeated:     entity.GetLabel() == 3,
		})
	}

	tmpl := template.Must(template.New("state").Parse(stateTemplate))

	var output bytes.Buffer
	tmpl.Execute(&output, stateEntities)

	return &File{
		Name:    "state_pb.ts",
		Content: output.String(),
	}, nil
}
