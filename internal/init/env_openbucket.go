// Copyright (c) 2026 OpenBucket Contributors.
//
// This file is part of OpenBucket Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package init

import (
	"os"
	"strings"
)

// OpenBucket is configured with OPENBUCKET_* (and _OPENBUCKET_*) environment
// variables. Internally the configuration subsystem still keys off the
// historical MINIO_* names, so the canonical OPENBUCKET_* variables are
// mirrored onto their MINIO_* equivalents before any package reads the
// environment. OPENBUCKET_* wins when both are set; bare MINIO_* variables
// keep working as a compatibility fallback for existing deployments.
func init() {
	for _, kv := range os.Environ() {
		k, v, ok := strings.Cut(kv, "=")
		if !ok {
			continue
		}
		var legacy string
		switch {
		case strings.HasPrefix(k, "OPENBUCKET_"):
			legacy = "MINIO_" + strings.TrimPrefix(k, "OPENBUCKET_")
		case strings.HasPrefix(k, "_OPENBUCKET_"):
			legacy = "_MINIO_" + strings.TrimPrefix(k, "_OPENBUCKET_")
		default:
			continue
		}
		os.Setenv(legacy, v)
	}
}
