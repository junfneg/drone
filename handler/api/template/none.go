// Copyright 2019 Drone IO, Inc.
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

//go:build oss
// +build oss

package template

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)
}

func HandleCreate(store core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleUpdate(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleDelete(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleFind(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleList(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleListAll(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}

func HandleAll(core.TemplateStore) http.HandlerFunc {
	return notImplemented
}
