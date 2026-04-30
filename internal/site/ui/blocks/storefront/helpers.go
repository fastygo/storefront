package storefront

import "strings"

func labelOrFallback(value string, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func hasAction(action Action) bool {
	return strings.TrimSpace(action.Label) != "" && strings.TrimSpace(action.Href) != ""
}

func actionLabel(action Action) string {
	if strings.TrimSpace(action.AriaLabel) != "" {
		return action.AriaLabel
	}
	return action.Label
}

func imageSrc(image ImageRef, fallback string) string {
	if strings.TrimSpace(image.Src) != "" {
		return image.Src
	}
	return fallback
}
