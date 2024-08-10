/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources
import "gitlab.com/tokend/go/xdr"
import (
	"time"
)

type Link struct {
	// The timestamp of when the link was created
	CreatedAt *date-time `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	// The original URL
	OriginalUrl *string `json:"original_url,omitempty"`
	// The shortened URL
	ShortUrl *string `json:"short_url,omitempty"`
}
