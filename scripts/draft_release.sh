#!/usr/bin/env bash

# Parse arguments
while [[ $# -gt 0 ]]; do
  case $1 in
  --release-type)
    release_type="$2"
    shift 2
    ;;
  --dry-run)
    dry_run="$2"
    shift 2
    ;;
  *)
    echo "Unknown option: $1"
    exit 1
    ;;
  esac
done
if [[ -z "$release_type" ]]; then
  release_type="minor"
fi

# Fail if not on main branch
echo 'Checking current branch...'
if [[ $(git rev-parse --abbrev-ref HEAD) != "main" ]]; then
  echo "You must be on the main branch to draft a release."
  exit 1
fi
# Fail if there are uncommitted changes
echo 'Checking for uncommitted changes...'
if [[ -n $(git status --porcelain) ]]; then
  echo "You have uncommitted changes. Please commit or stash them before drafting a release."
  exit 1
fi
# Get the latest tag
echo 'Fetching latest tags...'
git fetch --tags
latest_tag=$(git describe --tags --abbrev=0)
# Compute a new semantic tag based on the latest tag and provided argument
echo 'Computing new tag...'
if [[ "$release_type" == "major" ]]; then
  new_tag=$(echo "$latest_tag" | sed 's/v//' | awk -F. '{print $1+1".0.0"}')
elif [[ "$release_type" == "minor" ]]; then
  new_tag=$(echo "$latest_tag" | sed 's/v//' | awk -F. '{print $1"."$2+1".0"}')
elif [[ "$release_type" == "patch" ]]; then
  new_tag=$(echo "$latest_tag" | sed 's/v//' | awk -F. '{print $1"."$2"."$3+1}')
else
  echo "Invalid release type: $release_type"
  exit 1
fi
new_tag="v$new_tag"
# Exit if dry run is enabled
echo "New tag to be created: $new_tag"
if [[ "$dry_run" == "true" ]]; then
  echo "Dry run mode enabled. No changes will be made."
  exit 0
fi
# Create the new tag
echo "Creating new tag: $new_tag"
git tag -a "$new_tag" -m "Release $new_tag"
# Push the new tag to the remote
echo "Pushing new tag to remote..."
git push origin "$new_tag"
# Create a draft release on GitHub and generate release-notes
echo "Creating draft release on GitHub..."
gh release create "$new_tag" --draft --generate-notes --title "Release $new_tag" --notes "Release for $new_tag"
