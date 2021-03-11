/*
   Copyright 2020 Docker Compose CLI authors

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

package context

import (
	gocontext "context"

	"golang.org/x/net/context"

	cliflags "github.com/docker/cli/cli/flags"
)

type currentContextKey struct{}
type cliOptionsKey struct{}

// WithCurrentContext sets the name of the current docker context
func WithCurrentContext(ctx gocontext.Context, contextName string) context.Context {
	return context.WithValue(ctx, currentContextKey{}, contextName)
}

// CurrentContext returns the current context name
func CurrentContext(ctx context.Context) string {
	cc, _ := ctx.Value(currentContextKey{}).(string)
	return cc
}

// WithCliOptions sets CLI options
func WithCliOptions(ctx gocontext.Context, options cliflags.CommonOptions) context.Context {
	return context.WithValue(ctx, cliOptionsKey{}, options)
}

// CliOptions returns common cli options
func CliOptions(ctx context.Context) cliflags.CommonOptions {
	cc, _ := ctx.Value(cliOptionsKey{}).(cliflags.CommonOptions)
	return cc
}