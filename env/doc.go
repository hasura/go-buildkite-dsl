// Package env provides access to the values of environment variables
// setup by buildkite during a build.
//
// These environment variables provide various build information
// which could be useful during a build. Example: [env.BUILDKITE_TAG] provides
// the value of "BUILDKITE_TAG“ environment variable which gets set to git tag that triggered a build.
//
// For more details on these pre-populated environment variables, please refer to this buildkite
// [docs page]
//
// [docs page]: https://buildkite.com/docs/pipelines/environment-variables.
package env
