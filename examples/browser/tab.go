package main

import (
	"github.com/maneac/go-ultralight"
)

type tab struct {
	ui                              *ui
	overlay                         *ultralight.Overlay
	id                              uint
	isReadyToClose                  bool
	containerWidth, containerHeight uint
}

func (u *ui) createTab(id uint, width, height uint, x, y int) *tab {
	t := tab{}
	t.ui = u
	t.id = id
	t.containerWidth = width
	t.containerHeight = height
	t.overlay = ultralight.CreateOverlay(u.window, width, height, x, y)

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
}

func (t *tab) hide() {
	t.overlay.Hide()
	t.overlay.Unfocus()
}

func (t *tab) resize(width, height uint) {
	t.containerHeight = height
	t.containerWidth = width
	t.overlay.Resize(width, height)
}
