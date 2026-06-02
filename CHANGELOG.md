# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

Targeted for the 0.3.0 release.


### Added
- `mail batch` subcommand: execute multiple label/unlabel/archive/move/flag/delete
  operations from a JSON array over a single IMAP session, with per-op validation and
  an optional `--stop-on-error` flag. Reported by @Juan-de-Costa-Rica (#10).
- Server-side filtering for `mail list` (`--unread` via IMAP `SEARCH UNSEEN`, new
  `--flagged` via `SEARCH FLAGGED`), additional envelope fields (`from_address`, `to`,
  `message_id`, `in_reply_to`), and `--fields` selection. Reported by @Juan-de-Costa-Rica (#9).
- `PM_CLI_BRIDGE_PASSWORD` environment variable as a fallback for the Bridge password,
  for headless/no-keyring environments. Reported by @Juan-de-Costa-Rica (#8).
