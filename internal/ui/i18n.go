package ui

import "strings"

const (
	LocaleEnglish    = "en"
	LocalePortuguese = "pt-PT"
)

type messages map[string]string

var catalog = map[string]messages{
	LocaleEnglish: {
		"screen.home.title": "Home",
		"screen.home.description": "Repository and runtime status",
		"screen.preview.title": "Preview",
		"screen.preview.description": "Review declarative changes before apply",
		"screen.history.title": "History & Rollback",
		"screen.history.description": "Recover previously applied transactions",
		"screen.desktop.title": "Desktop",
		"screen.desktop.description": "Native desktop customization adapters",
		"screen.desktop.reason": "KDE/GNOME/XFCE mutation adapters are not integrated yet",
		"screen.themes.title": "Themes",
		"screen.themes.description": "Theme creation/import",
		"screen.themes.reason": "theme pipeline is not integrated yet",
		"warning.no_adapters": "No mutation adapters are available.",
		"warning.safe_adapters": "%d safe adapter(s) currently available; desktop-specific mutation remains gated.",
		"splash.tagline": "Created by one. Improved by many. Available to all.",
		"a11y.high_contrast": "High contrast",
		"a11y.keyboard": "Keyboard navigation",
	},
	LocalePortuguese: {
		"screen.home.title": "Início",
		"screen.home.description": "Estado do repositório e da execução",
		"screen.preview.title": "Pré-visualização",
		"screen.preview.description": "Rever alterações declarativas antes de aplicar",
		"screen.history.title": "Histórico e reversão",
		"screen.history.description": "Recuperar transações aplicadas anteriormente",
		"screen.desktop.title": "Ambiente de trabalho",
		"screen.desktop.description": "Adaptadores nativos de personalização do ambiente de trabalho",
		"screen.desktop.reason": "Os adaptadores de alteração KDE/GNOME/XFCE ainda não estão integrados",
		"screen.themes.title": "Temas",
		"screen.themes.description": "Criação/importação de temas",
		"screen.themes.reason": "O sistema de temas ainda não está integrado",
		"warning.no_adapters": "Não existem adaptadores de alteração disponíveis.",
		"warning.safe_adapters": "%d adaptador(es) seguro(s) disponível(eis); as alterações específicas do ambiente de trabalho continuam bloqueadas.",
		"splash.tagline": "Criado por um. Melhorado por muitos. Disponível para todos.",
		"a11y.high_contrast": "Alto contraste",
		"a11y.keyboard": "Navegação por teclado",
	},
}

func NormalizeLocale(locale string) string {
	normalized := strings.ToLower(strings.TrimSpace(strings.ReplaceAll(locale, "_", "-")))
	if normalized == "pt" || strings.HasPrefix(normalized, "pt-pt") {
		return LocalePortuguese
	}
	return LocaleEnglish
}

func message(locale, key string) string {
	locale = NormalizeLocale(locale)
	if value, ok := catalog[locale][key]; ok {
		return value
	}
	return catalog[LocaleEnglish][key]
}
