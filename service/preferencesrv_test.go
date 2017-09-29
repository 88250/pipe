package service

import (
	"testing"

	"github.com/b3log/solo.go/model"
)

func TestGetPreference(t *testing.T) {
	setting := Preference.GetPreference(model.SettingNamePreferenceBlogTitle, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "Solo.go 示例" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "Solo.go 示例", setting.Value)
	}
}

func TestGetPreferences(t *testing.T) {
	settings := Preference.GetPreferences(1, model.SettingNamePreferenceBlogTitle, model.SettingNamePreferenceBlogSubtitle)
	if nil == settings {
		t.Errorf("settings is nil")

		return
	}
	if 1 > len(settings) {
		t.Errorf("settings is empty")

		return
	}

	if "Solo.go 示例" != settings[model.SettingNamePreferenceBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "Solo.go 示例", settings[model.SettingNamePreferenceBlogTitle].Value)
	}
}
