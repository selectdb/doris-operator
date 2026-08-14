#!/usr/bin/env bash
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

set -euo pipefail
# shellcheck source=testlib.sh
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/testlib.sh"
# shellcheck source=../release.env
source "${TOOLS_ROOT}/release.env"
# shellcheck source=../lib/release-common.sh
source "${TOOLS_ROOT}/lib/release-common.sh"

[[ -n "$VERSION" ]] || fail "release.env does not set VERSION"
[[ -n "$GIT_REMOTE" ]] || fail "release.env does not set GIT_REMOTE"
assert_eq "$VERSION" "$TAG"
assert_eq "apache-doris-operator-${VERSION}-src" "$PKG_BASE"
assert_eq "${PKG_BASE}/" "$ARCHIVE_PREFIX"
assert_eq "https://dist.apache.org/repos/dist/dev/doris/doris-operator/${VERSION}" "$DEV_SVN_DIR"
assert_eq "https://dist.apache.org/repos/dist/release/doris/doris-operator/${VERSION}" "$RELEASE_SVN_DIR"
assert_eq "https://downloads.apache.org/doris/KEYS" "$KEYS_URL"
assert_eq "https://dist.apache.org/repos/dist/dev/doris" "$DEV_KEYS_SVN_BASE"
assert_eq "https://dist.apache.org/repos/dist/release/doris" "$RELEASE_KEYS_SVN_BASE"
assert_eq "apache/doris-operator" "$GITHUB_REPO"
assert_eq "https://github.com/${GITHUB_REPO}/releases/tag/${TAG}" "$GITHUB_TAG_URL"
assert_eq "apache/doris:operator-${VERSION}" "$DOCKER_IMAGE"
assert_eq "https://hub.docker.com/r/apache/doris/tags?name=operator-${VERSION}" "$DOCKER_IMAGE_URL"

validate_release_config

error_file="$(mktemp)"
trap 'rm -f "$error_file"' EXIT
if (TAG="not-${VERSION}"; validate_release_config) 2>"$error_file"; then
  fail "configuration validation accepted a tag that differs from VERSION"
fi
assert_file_contains "$error_file" "TAG must equal VERSION"

if (GITHUB_REPO="doris-operator"; validate_release_config) 2>"$error_file"; then
  fail "configuration validation accepted a GITHUB_REPO without an owner"
fi
assert_file_contains "$error_file" "GITHUB_REPO must be <owner>/<repo>"

if (DOCKER_IMAGE="apache/doris:operator-0.0.0"; validate_release_config) 2>"$error_file"; then
  fail "configuration validation accepted a DOCKER_IMAGE from another version"
fi
assert_file_contains "$error_file" "DOCKER_IMAGE must be"

if grep -Eq '^[[:space:]]*ASF_(USERNAME|PASSWORD)=' "${TOOLS_ROOT}/release.env"; then
  fail "release.env stores SVN credentials"
fi

pass
