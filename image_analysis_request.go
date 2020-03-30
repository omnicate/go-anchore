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

// A request to add an image to be watched and analyzed by the engine. Optionally include the dockerfile content. Either digest or tag must be present
type ImageAnalysisRequest struct {

	// Content of the dockerfile for the image, if available
	Dockerfile string `json:"dockerfile,omitempty"`

	// A full pullable digest reference for an image. e.g. docker.io/nginx@sha256:abc123
	Digest string `json:"digest,omitempty"`

	// Full pullable tag reference for image. e.g. docker.io/nginx:latest
	Tag string `json:"tag,omitempty"`

	// The type of image this is adding, defaults to \"docker\"
	ImageType string `json:"image_type,omitempty"`

	// Annotations to be associated with the added image in key/value form
	Annotations *interface{} `json:"annotations,omitempty"`
}
