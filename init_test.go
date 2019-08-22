package ultralight

import "testing"

func TestInit(t *testing.T) {
	if funcMap == nil {
		t.Error("viewChangeTitle map not made!")
	}
	if viewChangeTitle == nil {
		t.Error("viewChangeTitle map not made!")
	}
	if viewChangeURL == nil {
		t.Error("viewChangeURL map not made!")
	}
	if viewChangeTooltip == nil {
		t.Error("viewChangeTooltip map not made!")
	}
	if viewChangeCursor == nil {
		t.Error("viewChangeCursor map not made!")
	}
	if viewAddConsoleMessage == nil {
		t.Error("viewAddConsoleMessage map not made!")
	}
	if viewBeginLoading == nil {
		t.Error("viewBeginLoading map not made!")
	}
	if viewFinishLoading == nil {
		t.Error("viewFinishLoading map not made!")
	}
	if viewUpdateHistory == nil {
		t.Error("viewUpdateHistory map not made!")
	}
	if viewDOMReady == nil {
		t.Error("viewDOMReady map not made!")
	}
}
