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
HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=release.env
source "${HERE}/release.env"
# shellcheck source=lib/release-common.sh
source "${HERE}/lib/release-common.sh"

usage() {
  printf 'Usage: %s [--mail-only | --github-only]\n' "$0"
  printf '  --mail-only    regenerate announcement drafts without packaging or SVN\n'
  printf '  --github-only  publish or refresh the GitHub release only\n'
  printf '\n'
  printf 'A full run is idempotent: it detects the release files already present\n'
  printf 'in release SVN and a GitHub release that already carries these notes,\n'
  printf 'and skips those steps instead of failing.\n'
}

mail_only=0
github_only=0
while [[ "$#" -gt 0 ]]; do
  case "$1" in
    --mail-only) mail_only=1 ;;
    --github-only) github_only=1 ;;
    -h|--help) usage; exit 0 ;;
    *) usage >&2; die "unknown argument: $1" ;;
  esac
  shift
done

if [[ "$mail_only" -eq 1 && "$github_only" -eq 1 ]]; then
  usage >&2
  die "--mail-only and --github-only are mutually exclusive"
fi

write_announce_email() {
  local subject body_file eml_file body
  mkdir -p "$WORK_DIR"
  subject="[ANNOUNCE] Apache Doris Operator ${VERSION} release"
  body_file="${WORK_DIR}/announce-email.txt"
  eml_file="${WORK_DIR}/announce-email.eml"

  body="$(
    cat <<EOF
Hi all,

We are pleased to announce the release of Apache Doris Operator ${VERSION}.

Apache Doris Operator automates the deployment and management of Apache Doris clusters on Kubernetes.

Release downloads:
${DOWNLOAD_PAGE_URL}

Formal source artifact:
${RELEASE_SVN_DIR}/${PKG_BASE}.tar.gz

Release notes:
${RELEASE_NOTES_URL}

Thank you to everyone who contributed to this release.

Best regards,
${SIGNER_NAME}
EOF
  )"

  {
    printf 'Subject: %s\n' "$subject"
    printf '\n%s\n' "$body"
  } > "$body_file"
  {
    printf 'To: %s\n' "$ANNOUNCE_TO"
    printf 'Subject: %s\n' "$subject"
    printf 'Content-Type: text/plain; charset=UTF-8\n'
    printf '\n%s\n' "$body"
  } > "$eml_file"

  ok "announcement draft: ${body_file}"
  ok "mail draft: ${eml_file}"
  printf 'Subject: %s\n' "$subject"
  printf '%s\n' '----------------------------------------------------------------'
  printf '%s\n' "$body"
  printf '%s\n' '----------------------------------------------------------------'
  printf 'Review and send the message manually to %s. No email was sent.\n' "$ANNOUNCE_TO"
}

# Writes the GitHub release notes and prints only the file path.
write_github_release_notes() {
  local notes_file="${WORK_DIR}/github-release-notes.md"
  mkdir -p "$WORK_DIR"

  cat > "$notes_file" <<EOF
Apache Doris Operator ${VERSION} is released.

## Release note

${RELEASE_NOTES_URL}

## Docker image

\`${DOCKER_IMAGE}\`

\`\`\`shell
docker pull ${DOCKER_IMAGE}
\`\`\`

${DOCKER_IMAGE_URL}

## Source download

The official Apache source release of ${VERSION}:

- ${RELEASE_SVN_DIR}/${PKG_BASE}.tar.gz
- ${RELEASE_SVN_DIR}/${PKG_BASE}.tar.gz.asc
- ${RELEASE_SVN_DIR}/${PKG_BASE}.tar.gz.sha512

KEYS: ${KEYS_URL}

How to verify: ${VERIFY_GUIDE_URL}
EOF

  printf '%s\n' "$notes_file"
}

github_release_exists() {
  gh release view "$TAG" --repo "$GITHUB_REPO" >/dev/null 2>&1
}

github_release_body() {
  gh release view "$TAG" --repo "$GITHUB_REPO" --json body --jq '.body' 2>/dev/null |
    tr -d '\r'
}

# Turns the existing Git tag into a GitHub release. Never creates a tag, and
# never rewrites notes that already match the generated ones.
publish_github_release() {
  local notes_file title current desired

  gh auth status >/dev/null 2>&1 ||
    die "gh is not authenticated; run 'gh auth login' or export GH_TOKEN"

  notes_file="$(write_github_release_notes)"
  title="Apache Doris Operator ${VERSION}"
  ok "GitHub release notes: ${notes_file}"
  printf '%s\n' '----------------------------------------------------------------'
  cat "$notes_file"
  printf '%s\n' '----------------------------------------------------------------'

  if github_release_exists; then
    current="$(github_release_body)"
    desired="$(cat "$notes_file")"
    if [[ "$current" == "$desired" ]]; then
      ok "GitHub release ${TAG} is already published: ${GITHUB_TAG_URL}"
      return 0
    fi

    warn "GitHub release ${TAG} exists with different notes"
    if ! confirm "Replace the notes of GitHub release ${TAG}?"; then
      warn "left the existing GitHub release untouched: ${GITHUB_TAG_URL}"
      return 0
    fi
    gh release edit "$TAG" --repo "$GITHUB_REPO" \
      --title "$title" --notes-file "$notes_file"
    ok "updated GitHub release: ${GITHUB_TAG_URL}"
    return 0
  fi

  printf 'GitHub release target: %s (from existing tag %s)\n' "$GITHUB_REPO" "$TAG"
  if ! confirm "Create the GitHub release for tag ${TAG}?"; then
    warn "stopped before creating the GitHub release"
    return 0
  fi
  gh release create "$TAG" --repo "$GITHUB_REPO" --verify-tag \
    --title "$title" --notes-file "$notes_file"
  ok "created GitHub release: ${GITHUB_TAG_URL}"
}

if [[ "$mail_only" -eq 1 ]]; then
  validate_release_config mail
  ok "mail-only mode: skipping tag, package, signing, SVN, and GitHub operations"
  write_announce_email
  exit 0
fi

validate_release_config

if [[ "$github_only" -eq 1 ]]; then
  require_tools git gh || die "install the missing release prerequisites"
  ok "github-only mode: skipping package, signing, SVN, and mail operations"
  verify_tag_consistency
  publish_github_release
  exit 0
fi

require_tools git gpg svn sha512sum gzip gh || die "install the missing release prerequisites"
export GPG_TTY="$(tty 2>/dev/null || true)"

verify_tag_consistency

release_published=0
release_state="$(
  svn_version_dir_state "$RELEASE_SVN_DIR" \
    "${PKG_BASE}.tar.gz" "${PKG_BASE}.tar.gz.asc" "${PKG_BASE}.tar.gz.sha512"
)"

case "$release_state" in
  complete)
    ok "release files already published, skipping packaging and upload: ${RELEASE_SVN_DIR}/"
    release_published=1
    ;;
  partial)
    die "incomplete release directory: ${RELEASE_SVN_DIR}/; inspect and repair it manually"
    ;;
  *)
    SIGNER="$(resolve_signing_key)"
    ok "signer: ${SIGNER}"
    prepare_source_artifacts "$SIGNER"

    stage_and_commit_version_dir \
      "$RELEASE_SVN_BASE" \
      "$RELEASE_SVN_DIR" \
      "release-svn" \
      "Release Apache Doris Operator ${VERSION}" \
      "${SOURCE_ARTIFACTS[@]}"
    release_published="$SVN_COMMITTED"
    ;;
esac

if [[ "$release_published" -eq 1 ]]; then
  publish_github_release
  write_announce_email
fi
