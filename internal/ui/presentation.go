package ui

// PresentationLabels contains localized labels used by lightweight UI entrypoints.
// Native graphical frontends may use the same labels until richer view resources exist.
type PresentationLabels struct {
	Status    string
	Available string
	Gated     string
	Warning   string
}

// LabelsForLocale returns deterministic EN/pt-PT presentation copy.
func LabelsForLocale(locale string) PresentationLabels {
	if NormalizeLocale(locale) == LocalePortuguese {
		return PresentationLabels{
			Status:    "Estado da interface:",
			Available: "disponível",
			Gated:     "bloqueado",
			Warning:   "aviso",
		}
	}
	return PresentationLabels{
		Status:    "UI status:",
		Available: "available",
		Gated:     "gated",
		Warning:   "warning",
	}
}
