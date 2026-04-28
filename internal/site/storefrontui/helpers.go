package storefrontui

import (
	"strings"

	t "github.com/a-h/templ"
	"github.com/fastygo/ui8kit/utils"
)

func pageClass(props StorefrontPageProps) string {
	return utils.Cn("w-full bg-background text-foreground", props.Classes.Page)
}

func mainGridClass(props StorefrontPageProps) string {
	return utils.Cn("grid grid-cols-1 gap-6 xl:grid-cols-4", props.Classes.Main)
}

func contentClass(props StorefrontPageProps) string {
	return utils.Cn("min-w-0 xl:col-span-3", props.Classes.Content)
}

func railClass(props StorefrontPageProps) string {
	return utils.Cn("grid gap-6 xl:sticky xl:top-6 xl:self-start", props.Classes.Rail)
}

func softCardClass(extra ...string) string {
	return utils.Cn(append([]string{"overflow-hidden rounded-lg border border-border bg-card shadow-sm"}, extra...)...)
}

func sectionHeaderClass() string {
	return "flex items-center justify-between gap-4"
}

func iconCircleClass() string {
	return "inline-flex h-10 w-10 items-center justify-center rounded-full border border-border bg-background"
}

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

func headerNavLabel(value string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return "Main navigation"
}

func headerNavWrapClass(value string) string {
	return resolveClass(value, "ui-header-nav-wrap")
}

func headerNavClass(value string) string {
	return resolveClass(value, "ui-header-nav")
}

func headerNavLinkClass(value string) string {
	return resolveClass(value, "ui-header-nav-link")
}

func isExternalHref(href string) bool {
	return strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://")
}

func languageToggleID(value string) string {
	return resolveClass(value, "language-toggle")
}

func languageToggleClass(value string) string {
	return resolveClass(value, "ui-header-action-btn")
}

func languageToggleAttrs(props LanguageToggleProps) t.Attributes {
	attrs := t.Attributes{
		"data-default-locale":    props.DefaultLocale,
		"data-current-locale":    props.CurrentLocale,
		"data-next-locale":       props.NextLocale,
		"data-next-label":        props.NextLabel,
		"data-available-locales": strings.Join(props.AvailableLocales, ","),
	}
	if props.EnhanceWithJS {
		attrs["data-ui8kit-spa-lang"] = "1"
		attrs["data-spa-target"] = props.SPATarget
	}
	return attrs
}

func headerLanguageToggleClass(value string) string {
	return resolveClass(value, "ui-header-language")
}

func languageToggleLabel(label, locale string) string {
	if strings.TrimSpace(label) != "" {
		return label
	}
	return strings.ToUpper(strings.TrimSpace(locale))
}

func darkModeToggleID(value string) string {
	return resolveClass(value, "ui8kit-theme-toggle")
}

func darkModeToggleClass(value string) string {
	return resolveClass(value, "ui-header-theme-btn")
}

func darkModeToggleAttrs(props DarkModeToggleProps) t.Attributes {
	return t.Attributes{
		"data-switch-to-dark-label":  darkModeSwitchToDarkLabel(props.SwitchToDarkLabel),
		"data-switch-to-light-label": darkModeSwitchToLightLabel(props.SwitchToLightLabel),
		"title":                      darkModeToggleLabel(props.Label),
		"aria-pressed":               "false",
	}
}

func darkModeIconID(value string) string {
	return resolveClass(value, "theme-toggle-icon")
}

func darkModeIconClass(value string) string {
	return resolveClass(value, "ui-theme-icon latty latty-moon")
}

func darkModeToggleLabel(value string) string {
	return resolveClass(value, "Toggle theme")
}

func darkModeSwitchToDarkLabel(value string) string {
	return resolveClass(value, "Switch to dark mode")
}

func darkModeSwitchToLightLabel(value string) string {
	return resolveClass(value, "Switch to light mode")
}

func swatchAttrs(item SwatchItem, selectedValue string) t.Attributes {
	return t.Attributes{
		"aria-pressed": swatchSelected(item, selectedValue),
		"data-value":   item.Value,
	}
}

func swatchSelected(item SwatchItem, selectedValue string) bool {
	if item.Selected {
		return true
	}
	return strings.TrimSpace(selectedValue) != "" && item.Value == selectedValue
}

func swatchButtonClass(item SwatchItem, selectedValue string) string {
	selected := ""
	if swatchSelected(item, selectedValue) {
		selected = "border-primary"
	}
	return utils.Cn("inline-flex h-8 w-8 items-center justify-center rounded-full border border-border bg-background", selected)
}

func swatchDotClass(item SwatchItem) string {
	return utils.Cn("block h-5 w-5 rounded-full border border-border bg-muted", item.Class)
}

func quantityValue(value string) string {
	if strings.TrimSpace(value) == "" {
		return "1"
	}
	return value
}

func quantityMinusLabel(props QuantityProps) string {
	return resolveClass(props.MinusLabel, "Decrease quantity")
}

func quantityPlusLabel(props QuantityProps) string {
	return resolveClass(props.PlusLabel, "Increase quantity")
}

func resolveClass(value, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}
