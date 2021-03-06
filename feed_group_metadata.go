/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.5
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package anchore_engine

import (
	"time"
)

type FeedGroupMetadata struct {
	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	LastSync time.Time `json:"last_sync,omitempty"`

	RecordCount int32 `json:"record_count,omitempty"`
}
