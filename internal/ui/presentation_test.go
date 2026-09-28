package ui

import "testing"

func TestLabelsForLocalePortuguese(t *testing.T) {
	labels := LabelsForLocale("pt_PT.UTF-8")
	if labels.Status != "Estado da interface:" || labels.Available != "disponível" || labels.Gated != "bloqueado" || labels.Warning != "aviso" {
		t.Fatalf("unexpected Portuguese labels: %#v", labels)
	}
}

func TestLabelsForLocaleFallsBackToEnglish(t *testing.T) {
	labels := LabelsForLocale("de-DE")
	if labels.Status != "UI status:" || labels.Available != "available" || labels.Gated != "gated" || labels.Warning != "warning" {
		t.Fatalf("unexpected fallback labels: %#v", labels)
	}
}
