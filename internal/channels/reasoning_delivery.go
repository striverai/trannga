package channels

import (
	"os"
	"strings"
)

const (
	ReasoningDeliveryOff           = "off"
	ReasoningDeliveryStreamingOnly = "streaming_only"
	ReasoningDeliveryAlwaysBubbles = "always_bubbles"
)

type ResolvedReasoningDelivery struct {
	Mode                string
	ShowInChannel       bool
	ForceProviderStream bool
	BubbleDelivery      bool
}

func ResolveReasoningDelivery(mode string, legacyReasoningStream *bool) ResolvedReasoningDelivery {
	if mode == "" {
		if env := os.Getenv("GOCLAW_REASONING_DELIVERY"); env != "" {
			mode = env
		}
	}

	switch NormalizeReasoningDeliveryMode(mode) {
	case ReasoningDeliveryOff:
		return ResolvedReasoningDelivery{Mode: ReasoningDeliveryOff}
	case ReasoningDeliveryAlwaysBubbles:
		return ResolvedReasoningDelivery{
			Mode:                ReasoningDeliveryAlwaysBubbles,
			ShowInChannel:       true,
			ForceProviderStream: true,
			BubbleDelivery:      true,
		}
	case ReasoningDeliveryStreamingOnly:
		return ResolvedReasoningDelivery{
			Mode:          ReasoningDeliveryStreamingOnly,
			ShowInChannel: true,
		}
	}

	if legacyReasoningStream != nil && *legacyReasoningStream {
		return ResolvedReasoningDelivery{
			Mode:          ReasoningDeliveryStreamingOnly,
			ShowInChannel: true,
		}
	}
	return ResolvedReasoningDelivery{
		Mode:          ReasoningDeliveryOff,
		ShowInChannel: false,
	}
}

func NormalizeReasoningDeliveryMode(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case ReasoningDeliveryOff:
		return ReasoningDeliveryOff
	case ReasoningDeliveryStreamingOnly:
		return ReasoningDeliveryStreamingOnly
	case ReasoningDeliveryAlwaysBubbles:
		return ReasoningDeliveryAlwaysBubbles
	default:
		return ""
	}
}

func ShouldStreamProviderForDelivery(channelStreaming bool, delivery ResolvedReasoningDelivery) bool {
	return channelStreaming || delivery.ForceProviderStream
}
