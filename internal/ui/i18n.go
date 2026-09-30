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
		"screen.desktop.description": "Native desktop customization capabilities",
		"screen.desktop.description.kde": "Reviewed KDE color-scheme customization is available; broader desktop customization remains limited",
		"screen.desktop.reason": "No reviewed native desktop mutation adapter is available in this runtime",
		"screen.themes.title": "Themes",
		"screen.themes.description": "Theme creation/import",
		"screen.themes.reason": "The theme pipeline is not integrated yet",
		"warning.no_adapters": "No mutation adapters are available.",
		"warning.safe_adapters": "%d reviewed adapter(s) currently available; only capabilities exposed by those adapters are enabled.",
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
		"screen.desktop.description": "Capacidades nativas de personalização do ambiente de trabalho",
		"screen.desktop.description.kde": "A alteração revista do esquema de cores KDE está disponível; a personalização mais ampla do ambiente de trabalho continua limitada",
		"screen.desktop.reason": "Não existe neste runtime um adaptador nativo de alteração revisto",
		"screen.themes.title": "Temas",
		"screen.themes.description": "Criação/importação de temas",
		"screen.themes.reason": "O sistema de temas ainda não está integrado",
		"warning.no_adapters": "Não existem adaptadores de alteração disponíveis.",
		"warning.safe_adapters": "%d adaptador(es) revisto(s) disponível(eis); apenas as capacidades expostas por esses adaptadores são ativadas.",
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
