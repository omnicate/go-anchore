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

// A service status record
type Service struct {

	// The unique id of the host on which the service is executing
	HostId string `json:"host_id,omitempty"`

	// Registered service name
	ServiceName string `json:"service_name,omitempty"`

	// The url to reach the service, including port as needed
	BaseUrl string `json:"base_url,omitempty"`

	// A state indicating the condition of the service. Normal operation is 'registered'
	StatusMessage string `json:"status_message,omitempty"`

	// The version of the service as reported by the service implementation on registration
	Version string `json:"version,omitempty"`
}
