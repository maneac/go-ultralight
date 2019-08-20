package main

import (
	"github.com/maneac/go-ultralight"
)

const inspectorHeight = 300

type tab struct {
	ui                              *ui
	overlay                         *ultralight.Overlay
	inspectorOverlay                *ultralight.Overlay
	id                              uint
	isReadyToClose                  bool
	containerWidth, containerHeight int
}

func (u *ui) createTab(id uint, width, height int, x, y int) *tab {
	t := tab{}
	t.ui = u
	t.id = id
	t.containerWidth = width
	t.containerHeight = height
	t.overlay = u.window.CreateOverlay(width, height, x, y)

	t.view().SetChangeTitleCallback(func(title string) {
		t.ui.updateTabTitle(t.id, title)
	})
	t.view().SetChangeURLCallback(func(url string) {
		t.ui.updateTabURL(t.id, url)
	})
	t.view().SetChangeCursorCallback(func(cursor ultralight.Cursor) {
		if t.id == t.ui.activeTabID {
			t.ui.setCursor(cursor)
		}
	})
	t.view().SetBeginLoadingCallback(func() {
		t.ui.updateTabNavigation(id, t.view().IsLoading(), t.view().CanGoBack(), t.view().CanGoForward())
	})
	t.view().SetFinishLoadingCallback(func() {
		t.ui.updateTabNavigation(id, t.view().IsLoading(), t.view().CanGoBack(), t.view().CanGoForward())
	})
	t.view().SetUpdateHistoryCallback(func() {
		t.ui.updateTabNavigation(id, t.view().IsLoading(), t.view().CanGoBack(), t.view().CanGoForward())
	})
	return &t
}

func (t *tab) setReadyToClose(ready bool) {
	t.isReadyToClose = ready
}

func (t *tab) readyToClose() bool {
	return t.isReadyToClose
}

func (t *tab) view() *ultralight.View {
	return t.overlay.GetView()
}

func (t *tab) show() {
	t.overlay.Show()
	t.overlay.Focus()
	if t.inspectorOverlay != nil {
		t.inspectorOverlay.Show()
	}
}

func (t *tab) hide() {
	t.overlay.Hide()
	t.overlay.Unfocus()
	if t.inspectorOverlay != nil {
		t.inspectorOverlay.Hide()
	}
}

func (t *tab) toggleInspector() {
	// if t.inspectorOverlay == nil {
	// 	overlay := t.ui.window.CreateOverlay(t.overlay.GetView().Inspector(), 0, 0)
	// } else {
	// 	if t.inspectorOverlay.IsHidden() {
	// 		t.inspectorOverlay.Show()
	// 	} else {
	// 		t.inspectorOverlay.Hide()
	// 	}
	// }
	// t.resize(t.containerWidth, t.containerHeight)
}

func (t *tab) resize(width, height int) {
	t.containerWidth = width
	t.containerHeight = height
	contentHeight := t.containerHeight
	if t.inspectorOverlay != nil && !t.inspectorOverlay.IsHidden() {
		t.inspectorOverlay.Resize(width, height)
		contentHeight -= inspectorHeight
	}
	if contentHeight < 1 {
		contentHeight = 1
	}
	t.overlay.Resize(t.containerWidth, t.containerHeight)
	if t.inspectorOverlay != nil && !t.inspectorOverlay.IsHidden() {
		t.inspectorOverlay.MoveTo(0, t.overlay.GetY()+int(t.overlay.GetHeight()))
	}
}
