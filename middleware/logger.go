/*
Copyright 2015 The ContainerOps Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package middleware

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/macaron.v1"

	"github.com/containerops/dockyard/setting"
)

func logger() macaron.Handler {
	return func(ctx *macaron.Context) {
		if setting.RunMode == "dev" {
			log.Info("------------------------------------------------------------------------------")
			log.Info(time.Now().String())
		}

		log.WithFields(log.Fields{
			"Method": ctx.Req.Method,
			"URL":    ctx.Req.RequestURI,
		}).Info(ctx.Req.Header)

	}
}
