package metadata

import (
	"github.com/avislash/nftstamper/lib/metadata"
)

type HoundMetadataFetcher struct {
	metadata.Fetcher[HoundMetadata]
}

func NewHoundMetadataFetcher(baseURL string) *HoundMetadataFetcher {
	return &HoundMetadataFetcher{metadata.NewJSONMetadataFetcher[HoundMetadata](baseURL)}
}
