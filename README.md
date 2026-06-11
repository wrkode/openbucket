# OpenBucket

[![license](https://img.shields.io/badge/license-AGPL%20V3-blue)](https://github.com/wrkode/openbucket/blob/main/LICENSE)

OpenBucket is a community-maintained fork of [MinIO](https://github.com/minio/minio), the high-performance, S3-compatible object storage server.

In late 2025 the upstream MinIO project stopped publishing community binary releases, removed the management features from the community web console, and was subsequently placed in maintenance mode and archived. OpenBucket exists to keep a fully usable, fully open community edition alive:

- **Binary releases** — pre-built binaries and container images, built transparently with GitHub Actions
- **Full web UI** — the complete management console (users, policies, replication, monitoring), not just an object browser
- **S3 API compatibility** — unchanged from upstream MinIO
- **AGPLv3** — same license as upstream, always

OpenBucket is not affiliated with or endorsed by MinIO, Inc. "MinIO" is a trademark of MinIO, Inc.

## Migrating from MinIO

OpenBucket is a drop-in replacement for a MinIO server deployment:

- The on-disk format, S3 API, and admin API are unchanged.
- Configuration uses `OPENBUCKET_*` environment variables (e.g. `OPENBUCKET_ROOT_USER`), but all legacy `MINIO_*` variables continue to work as a fallback. `_MINIO_*` debug variables map to `_OPENBUCKET_*` likewise.
- Point the `openbucket` binary at your existing MinIO data directories and start it the same way you started `minio`.

## Quickstart

### Build from source

Requires Go 1.25 or later.

```sh
go install github.com/wrkode/openbucket@latest
openbucket server /data --console-address :9001
```

The server starts with default root credentials `minioadmin:minioadmin` unless `OPENBUCKET_ROOT_USER` / `OPENBUCKET_ROOT_PASSWORD` are set. Open <http://127.0.0.1:9001> for the web console, or connect any S3-compatible client to port 9000:

```sh
mc alias set local http://localhost:9000 minioadmin minioadmin
mc admin info local
```

### Container image

```sh
docker run -p 9000:9000 -p 9001:9001 \
  -e OPENBUCKET_ROOT_USER=admin -e OPENBUCKET_ROOT_PASSWORD=changeme123 \
  ghcr.io/wrkode/openbucket:latest server /data --console-address :9001
```

### Makefile

```sh
make build    # builds ./openbucket
make test     # builds and runs the test suite
make docker   # builds the container image
```

## Project status

OpenBucket forked from the final upstream commit of `minio/minio` (master, December 2025). Current goals, roughly in order:

1. ~~Rebranded, building server with `OPENBUCKET_*`/`MINIO_*` dual env support~~ ✅
2. Restore the full management console UI (based on the last full-featured `minio/console`)
3. Automated release pipeline: cross-compiled binaries, checksums, signed container images
4. Dependency and security updates

Contributions are welcome — issues and pull requests are open.

## License

OpenBucket is licensed under the [GNU AGPLv3](LICENSE), the same license as the upstream MinIO project. Original code copyright MinIO, Inc.; modifications copyright OpenBucket contributors. See [COMPLIANCE.md](COMPLIANCE.md) for license compliance notes.
