package env

import "os"


// Always true
var BUILDKITE = os.Getenv("BUILDKITE")

// The agent session token for the job. The variable is read by the agent artifact and meta-data commands.
// 
// Example:
// 83d544ccc223c157d2bf80d3f2a32982c32c3c0db8e3674820da5064783fb091
var BUILDKITE_AGENT_ACCESS_TOKEN = os.Getenv("BUILDKITE_AGENT_ACCESS_TOKEN")

// The value of the debug agent configuration option.
var BUILDKITE_AGENT_DEBUG = os.Getenv("BUILDKITE_AGENT_DEBUG")

// The value of the disconnect-after-job agent configuration option.
var BUILDKITE_AGENT_DISCONNECT_AFTER_JOB = os.Getenv("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB")

// The value of the disconnect-after-idle-timeout agent configuration option.
// 
// Example:
// 10
var BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT = os.Getenv("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT")

// The value of the endpoint agent configuration option. This is set as an environment variable by the bootstrap and then read by most of the buildkite-agent commands.
var BUILDKITE_AGENT_ENDPOINT = os.Getenv("BUILDKITE_AGENT_ENDPOINT")

// A list of the experimental agent features that are currently enabled. The value can be set using the --experiment flag on the buildkite-agent start command or in your agent configuration file.
// 
// Example:
// experiment1,experiment2
var BUILDKITE_AGENT_EXPERIMENT = os.Getenv("BUILDKITE_AGENT_EXPERIMENT")

// The value of the health-check-addr agent configuration option.
// 
// Example:
// localhost:8080
var BUILDKITE_AGENT_HEALTH_CHECK_ADDR = os.Getenv("BUILDKITE_AGENT_HEALTH_CHECK_ADDR")

// The UUID of the agent.
// 
// Example:
// 1a222222-e999-3636-8ddd-802222222222
var BUILDKITE_AGENT_ID = os.Getenv("BUILDKITE_AGENT_ID")

// The name of the agent that ran the job.
// 
// Example:
// elastic-builders-088264dc4f9
var BUILDKITE_AGENT_NAME = os.Getenv("BUILDKITE_AGENT_NAME")

// The process ID of the agent.
// 
// Example:
// 6
var BUILDKITE_AGENT_PID = os.Getenv("BUILDKITE_AGENT_PID")

// The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// tmp/capybara/**/*;coverage/**/*
var BUILDKITE_ARTIFACT_PATHS = os.Getenv("BUILDKITE_ARTIFACT_PATHS")

// The path where artifacts will be uploaded. This variable is read by the buildkite-agent artifact upload command, and during the artifact upload phase of command steps. It can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.
// 
// Example:
// s3://name-of-your-s3-bucket/$BUILDKITE_PIPELINE_ID/$BUILDKITE_BUILD_ID/$BUILDKITE_JOB_ID
var BUILDKITE_ARTIFACT_UPLOAD_DESTINATION = os.Getenv("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION")

// The path to the directory containing the buildkite-agent binary.
// 
// Example:
// /usr/local/bin
var BUILDKITE_BIN_PATH = os.Getenv("BUILDKITE_BIN_PATH")

// The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by BUILDKITE_COMMIT.
// 
// Example:
// main
var BUILDKITE_BRANCH = os.Getenv("BUILDKITE_BRANCH")

// The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// /var/lib/buildkite-agent/builds/agent-1/pipeline-2
var BUILDKITE_BUILD_CHECKOUT_PATH = os.Getenv("BUILDKITE_BUILD_CHECKOUT_PATH")

// The name of the user who authored the commit being built. May be unverified. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
// 
// Example:
// Carol Danvers
var BUILDKITE_BUILD_AUTHOR = os.Getenv("BUILDKITE_BUILD_AUTHOR")

// The notification email of the user who authored the commit being built. May be unverified. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
// 
// Example:
// cdanvers@kree-net.com
var BUILDKITE_BUILD_AUTHOR_EMAIL = os.Getenv("BUILDKITE_BUILD_AUTHOR_EMAIL")

// The name of the user who created the build. May be unverified.
// 
// Example:
// Carol Danvers
var BUILDKITE_BUILD_CREATOR = os.Getenv("BUILDKITE_BUILD_CREATOR")

// The notification email of the user who created the build.
// 
// Example:
// cdanvers@kree-net.com
var BUILDKITE_BUILD_CREATOR_EMAIL = os.Getenv("BUILDKITE_BUILD_CREATOR_EMAIL")

// A colon separated list of unverified non-private team slugs that the build creator belongs to.
// 
// Example:
// everyone:platform
var BUILDKITE_BUILD_CREATOR_TEAMS = os.Getenv("BUILDKITE_BUILD_CREATOR_TEAMS")

// The UUID of the build.
// 
// Example:
// 4735ba57-80d0-46e2-8fa0-b28223a86586
var BUILDKITE_BUILD_ID = os.Getenv("BUILDKITE_BUILD_ID")

// The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.
// 
// Example:
// 1514
var BUILDKITE_BUILD_NUMBER = os.Getenv("BUILDKITE_BUILD_NUMBER")

// The value of the build-path agent configuration option.
// 
// Example:
// /var/lib/buildkite-agent/builds/
var BUILDKITE_BUILD_PATH = os.Getenv("BUILDKITE_BUILD_PATH")

// The url for this build on Buildkite.
// 
// Example:
// https://buildkite.com/acme-inc/my-project/builds/1514
var BUILDKITE_BUILD_URL = os.Getenv("BUILDKITE_BUILD_URL")

// The value of the cancel-grace-period agent configuration option in seconds.
var BUILDKITE_CANCEL_GRACE_PERIOD = os.Getenv("BUILDKITE_CANCEL_GRACE_PERIOD")

// The value of the cancel-signal agent configuration option. The value can be modified by exporting the environment variable in the environment or pre-checkout hooks.
var BUILDKITE_CANCEL_SIGNAL = os.Getenv("BUILDKITE_CANCEL_SIGNAL")

// Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in environment or pre-checkout hooks.
var BUILDKITE_CLEAN_CHECKOUT = os.Getenv("BUILDKITE_CLEAN_CHECKOUT")

// The command that will be run for the job.
// 
// Example:
// script/buildkite/specs
var BUILDKITE_COMMAND = os.Getenv("BUILDKITE_COMMAND")

// The opposite of the value of the no-command-eval agent configuration option.
var BUILDKITE_COMMAND_EVAL = os.Getenv("BUILDKITE_COMMAND_EVAL")

// The exit code from the last command run in the command hook.
// 
// Example:
// -1
var BUILDKITE_COMMAND_EXIT_STATUS = os.Getenv("BUILDKITE_COMMAND_EXIT_STATUS")

// The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like HEAD.
// 
// Example:
// 83a20ec058e2fb00e7fa4558c4c6e81e2dcf253d
var BUILDKITE_COMMIT = os.Getenv("BUILDKITE_COMMIT")

// The path to the agent config file.
var BUILDKITE_CONFIG_PATH = os.Getenv("BUILDKITE_CONFIG_PATH")

// The path to the file containing the job's environment variables.
// 
// Example:
// /tmp/job-env-36711a2a-711a-484e-b180-e1b3711a80cf51b18711a
var BUILDKITE_ENV_FILE = os.Getenv("BUILDKITE_ENV_FILE")

// The value of the git-clean-flags agent configuration option. The value can be modified by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// -ffxdq
var BUILDKITE_GIT_CLEAN_FLAGS = os.Getenv("BUILDKITE_GIT_CLEAN_FLAGS")

// The value of the git-clone-flags agent configuration option. The value can be modified by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// -v
var BUILDKITE_GIT_CLONE_FLAGS = os.Getenv("BUILDKITE_GIT_CLONE_FLAGS")

// The opposite of the value of the no-git-submodules agent configuration option.
var BUILDKITE_GIT_SUBMODULES = os.Getenv("BUILDKITE_GIT_SUBMODULES")

// The GitHub deployment ID. Only available on builds triggered by a GitHub Deployment.
// 
// Example:
// 87972451
var BUILDKITE_GITHUB_DEPLOYMENT_ID = os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_ID")

// The name of the GitHub deployment environment. Only available on builds triggered by a GitHub Deployment.
// 
// Example:
// production
var BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT = os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT")

// The name of the GitHub deployment task. Only available on builds triggered by a GitHub Deployment.
// 
// Example:
// deploy
var BUILDKITE_GITHUB_DEPLOYMENT_TASK = os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_TASK")

// The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a GitHub Deployment.
// 
// Example:
// production
var BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD = os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD")

// The value of the hooks-path agent configuration option.
// 
// Example:
// /etc/buildkite-agent/hooks/
var BUILDKITE_HOOKS_PATH = os.Getenv("BUILDKITE_HOOKS_PATH")

// A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.
// 
// Example:
// BUILDKITE_GIT_CLEAN_FLAGS
var BUILDKITE_IGNORED_ENV = os.Getenv("BUILDKITE_IGNORED_ENV")

// The internal UUID Buildkite uses for this job.
// 
// Example:
// e44f9784-e20e-4b93-a21d-f41fd5869db9
var BUILDKITE_JOB_ID = os.Getenv("BUILDKITE_JOB_ID")

// The path to a temporary file containing the logs for this job. Requires enabling the enable-job-log-tmpfile agent configuration option.
// 
// Example:
// /tmp/buildkite_job_log1931317484
var BUILDKITE_JOB_LOG_TMPFILE = os.Getenv("BUILDKITE_JOB_LOG_TMPFILE")

// The label/name of the current job.
// 
// Example:
// "🔨 Specs"
var BUILDKITE_LABEL = os.Getenv("BUILDKITE_LABEL")

// The exit code of the last hook that ran, used internally by the hooks.
// 
// Example:
// -1
var BUILDKITE_LAST_HOOK_EXIT_STATUS = os.Getenv("BUILDKITE_LAST_HOOK_EXIT_STATUS")

// The opposite of the value of the no-local-hooks agent configuration option.
var BUILDKITE_LOCAL_HOOKS_ENABLED = os.Getenv("BUILDKITE_LOCAL_HOOKS_ENABLED")

// The message associated with the build, usually the commit message.
// 
// Example:
// Added a great new feature
var BUILDKITE_MESSAGE = os.Getenv("BUILDKITE_MESSAGE")

// The organization name on Buildkite as used in URLs.
// 
// Example:
// acme-inc
var BUILDKITE_ORGANIZATION_SLUG = os.Getenv("BUILDKITE_ORGANIZATION_SLUG")

// The index of each parallel job created from a parallel build step, starting from 0. For a build step with parallelism: 5, the value would be 0, 1, 2, 3, and 4 respectively.
// 
// Example:
// 0
var BUILDKITE_PARALLEL_JOB = os.Getenv("BUILDKITE_PARALLEL_JOB")

// The total number of parallel jobs created from a parallel build step. For a build step with parallelism: 5, the value is 5.
// 
// Example:
// 5
var BUILDKITE_PARALLEL_JOB_COUNT = os.Getenv("BUILDKITE_PARALLEL_JOB_COUNT")

// The default branch for this pipeline.
// 
// Example:
// main
var BUILDKITE_PIPELINE_DEFAULT_BRANCH = os.Getenv("BUILDKITE_PIPELINE_DEFAULT_BRANCH")

// The displayed pipeline name on Buildkite.
// 
// Example:
// my_project
var BUILDKITE_PIPELINE_NAME = os.Getenv("BUILDKITE_PIPELINE_NAME")

// The ID of the source code provider for the pipeline's repository.
// 
// Example:
// github
var BUILDKITE_PIPELINE_PROVIDER = os.Getenv("BUILDKITE_PIPELINE_PROVIDER")

// The pipeline slug on Buildkite as used in URLs.
// 
// Example:
// my-project
var BUILDKITE_PIPELINE_SLUG = os.Getenv("BUILDKITE_PIPELINE_SLUG")

// A JSON object containing a list plugins used in the step, and their configuration.
// 
// Example:
// [{"github.com/buildkite-plugins/docker-buildkite-plugin#v3.7.0":{"image":"node:lts-alpine3.14"}}]
var BUILDKITE_PLUGINS = os.Getenv("BUILDKITE_PLUGINS")

// The opposite of the value of the no-plugins agent configuration option.
var BUILDKITE_PLUGINS_ENABLED = os.Getenv("BUILDKITE_PLUGINS_ENABLED")

// The value of the plugins-path agent configuration option.
// 
// Example:
// /etc/buildkite-agent/plugins/
var BUILDKITE_PLUGINS_PATH = os.Getenv("BUILDKITE_PLUGINS_PATH")

// Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the environment or pre-checkout hooks, or in a pipeline.yml file. It can also be enabled using the no-plugin-validation agent configuration option.
var BUILDKITE_PLUGIN_VALIDATION = os.Getenv("BUILDKITE_PLUGIN_VALIDATION")

// The number of the pull request, if this branch is a pull request.
// 
// Example:
// "123" for pull request #123, or "false" if not a pull request
var BUILDKITE_PULL_REQUEST = os.Getenv("BUILDKITE_PULL_REQUEST")

// The base branch that the pull request is targeting.
// 
// Example:
// "main", or "" if not a pull request
var BUILDKITE_PULL_REQUEST_BASE_BRANCH = os.Getenv("BUILDKITE_PULL_REQUEST_BASE_BRANCH")

// Set to true when the pull request is a draft. This variable is only available if a build contains a draft pull request.
// 
// Example:
// true
var BUILDKITE_PULL_REQUEST_DRAFT = os.Getenv("BUILDKITE_PULL_REQUEST_DRAFT")

// The repository URL of the pull request.
// 
// Example:
// "git://github.com/acme-inc/my-project.git", or "``" if not a pull request
var BUILDKITE_PULL_REQUEST_REPO = os.Getenv("BUILDKITE_PULL_REQUEST_REPO")

// The UUID of the original build this was rebuilt from.
// 
// Example:
// "4735ba57-80d0-46e2-8fa0-b28223a86586", or "" if not a rebuild
var BUILDKITE_REBUILT_FROM_BUILD_ID = os.Getenv("BUILDKITE_REBUILT_FROM_BUILD_ID")

// The UUID of the original build this was rebuilt from.
// 
// Example:
// "1514", or "" if not a rebuild
var BUILDKITE_REBUILT_FROM_BUILD_NUMBER = os.Getenv("BUILDKITE_REBUILT_FROM_BUILD_NUMBER")

// A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// +refs/weird/123abc:refs/local/weird/456
var BUILDKITE_REFSPEC = os.Getenv("BUILDKITE_REFSPEC")

// The repository of your pipeline. This variable can be set by exporting the environment variable in the environment or pre-checkout hooks.
// 
// Example:
// git@github.com:acme-inc/my-project.git
var BUILDKITE_REPO = os.Getenv("BUILDKITE_REPO")

// Only available if the experimental feature git-mirrors is enabled. The path to the shared git mirror.
// 
// Example:
// /tmp/buildkite-git-mirrors
var BUILDKITE_REPO_MIRROR = os.Getenv("BUILDKITE_REPO_MIRROR")

// How many times this job has been retried.
// 
// Example:
// 0
var BUILDKITE_RETRY_COUNT = os.Getenv("BUILDKITE_RETRY_COUNT")

// The access key ID for your S3 IAM user, for use with private S3 buckets. The variable is read by the buildkite-agent artifact upload command, and during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.
// 
// Example:
// AKIAIOSFODNN7EXAMPLE
var BUILDKITE_S3_ACCESS_KEY_ID = os.Getenv("BUILDKITE_S3_ACCESS_KEY_ID")

// The access URL for your private S3 bucket, if you are using a proxy. The variable is read by the buildkite-agent artifact upload command, as well as during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.
// 
// Example:
// https://buildkite-artifacts.example.com/
var BUILDKITE_S3_ACCESS_URL = os.Getenv("BUILDKITE_S3_ACCESS_URL")

// The Access Control List to be set on artifacts being uploaded to your private S3 bucket. The variable is read by the buildkite-agent artifact upload command, as well as during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.Must be one of the following values which map to S3 Canned ACL grants.
var BUILDKITE_S3_ACL = os.Getenv("BUILDKITE_S3_ACL")

// The region of your private S3 bucket. The variable is read by the buildkite-agent artifact upload command, as well as during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.
var BUILDKITE_S3_DEFAULT_REGION = os.Getenv("BUILDKITE_S3_DEFAULT_REGION")

// The secret access key for your S3 IAM user, for use with private S3 buckets. The variable is read by the buildkite-agent artifact upload command, as well as during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks. Do not print or export this variable anywhere except your agent hooks.
// 
// Example:
// wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
var BUILDKITE_S3_SECRET_ACCESS_KEY = os.Getenv("BUILDKITE_S3_SECRET_ACCESS_KEY")

// Whether to enable encryption for the artifacts in your private S3 bucket. The variable is read by the buildkite-agent artifact upload command, as well as during the artifact upload phase of command steps. The value can only be set by exporting the environment variable in the environment, pre-checkout or pre-command hooks.
var BUILDKITE_S3_SSE_ENABLED = os.Getenv("BUILDKITE_S3_SSE_ENABLED")

// The value of the shell agent configuration option.
// 
// Example:
// "/bin/bash -e -c"
var BUILDKITE_SHELL = os.Getenv("BUILDKITE_SHELL")

// The source of the event that created the build.
var BUILDKITE_SOURCE = os.Getenv("BUILDKITE_SOURCE")

// The opposite of the value of the no-ssh-keyscan agent configuration option.
var BUILDKITE_SSH_KEYSCAN = os.Getenv("BUILDKITE_SSH_KEYSCAN")

// A unique string that identifies a step.
// 
// Example:
// 080b7d73-986d-4a39-a510-b34f9faf4710
var BUILDKITE_STEP_ID = os.Getenv("BUILDKITE_STEP_ID")

// The value of the key command step attribute, a unique string set by you to identify a step.
// 
// Example:
// tests-06
var BUILDKITE_STEP_KEY = os.Getenv("BUILDKITE_STEP_KEY")

// The name of the tag being built, if this build was triggered from a tag.
// 
// Example:
// v1.2.3
var BUILDKITE_TAG = os.Getenv("BUILDKITE_TAG")

// The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified. Jobs that time out with an exit status of 0 are marked as "passed".
// 
// Example:
// "15" for 15 minutes, or "false" if no timeout is set
var BUILDKITE_TIMEOUT = os.Getenv("BUILDKITE_TIMEOUT")

// Set to "datadog" to send metrics to the Datadog APM using localhost:8126, or DD_AGENT_HOST:DD_AGENT_APM_PORT.Also available as a buildkite agent configuration option.
// 
// Example:
// datadog
var BUILDKITE_TRACING_BACKEND = os.Getenv("BUILDKITE_TRACING_BACKEND")

// The UUID of the build that triggered this build.
// 
// Example:
// "5aa7c894-c8c0-435b-bc17-13923b90f163", or "" if the build was not triggered from another build
var BUILDKITE_TRIGGERED_FROM_BUILD_ID = os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_ID")

// The number of the build that triggered this build.
// 
// Example:
// "1264", or "" if the build was not triggered from another build
var BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER = os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER")

// The slug of the pipeline that was used to trigger this build.
// 
// Example:
// "build-and-test", or "" if the build was not triggered from another build
var BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG = os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG")

// The name of the user who unblocked the build.
// 
// Example:
// Carol Danvers
var BUILDKITE_UNBLOCKER = os.Getenv("BUILDKITE_UNBLOCKER")

// The notification email of the user who unblocked the build.
// 
// Example:
// carol@nasa.gov
var BUILDKITE_UNBLOCKER_EMAIL = os.Getenv("BUILDKITE_UNBLOCKER_EMAIL")

// The UUID of the user who unblocked the build.
// 
// Example:
// 4735ba57-80d0-46e2-8fa0-b28223a86586
var BUILDKITE_UNBLOCKER_ID = os.Getenv("BUILDKITE_UNBLOCKER_ID")

// A colon separated list of non-private team slugs that the user who unblocked the build belongs to.
// 
// Example:
// everyone:platform
var BUILDKITE_UNBLOCKER_TEAMS = os.Getenv("BUILDKITE_UNBLOCKER_TEAMS")

// Always `true".
var CI = os.Getenv("CI")
