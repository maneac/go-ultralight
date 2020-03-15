package main

import (
	"fmt"
	"strconv"

	"github.com/maneac/go-ultralight"
)

var gUI *ui

const uiHeight = 79

type ui struct {
	window       *ultralight.Window
	overlay      *ultralight.Overlay
	uiHeight     int
	tabHeight    int
	scale        float32
	tabs         map[uint]*tab
	activeTabID  uint
	tabIDCounter uint
	cursor       ultralight.Cursor
}

func createUI(window *ultralight.Window) *ui {
	u := ui{}
	u.window = window
	u.overlay = ultralight.CreateOverlay(window, window.GetWidth(), uiHeight, 0, 0)
	gUI = &u
	view := u.overlay.GetView()
	u.tabs = make(map[uint]*tab)
	u.activeTabID = 0
	u.tabIDCounter = 0

	window.SetResizeCallback(func(_, _ uint) {
		tabHeight := window.GetHeight() - uiHeight
		if tabHeight < 1 {
			tabHeight = 1
		}
		u.overlay.Resize(window.GetWidth(), uiHeight)
		for _, tab := range u.tabs {
			tab.resize(window.GetWidth(), tabHeight)
		}
	})

	view.SetDOMReadyCallback(func() {
		bindCallbacks(&u, view)
	})

	view.LoadURL("file://assets/ui.html")

	return &u
}

func bindCallbacks(u *ui, view *ultralight.View) {
	view.BindJSCallback("OnBack", func(v *ultralight.View, params []string) {
		t := u.activeTab()
		if t != nil {
			t.view().GoBack()
		}
	})
	view.BindJSCallback("OnForward", func(v *ultralight.View, params []string) {
		t := u.activeTab()
		if t != nil {
			t.view().GoForward()
		}
	})
	view.BindJSCallback("OnRefresh", func(v *ultralight.View, params []string) {
		t := u.activeTab()
		if t != nil {
			t.view().Reload()
		}
	})
	view.BindJSCallback("OnStop", func(v *ultralight.View, params []string) {
		t := u.activeTab()
		if t != nil {
			t.view().Stop()
		}
	})
	view.BindJSCallback("OnRequestNewTab", func(v *ultralight.View, params []string) {
		u.createNewTab()
	})
	view.BindJSCallback("OnRequestTabClose", func(v *ultralight.View, params []string) {
		if len(params) == 1 {
			intID, _ := strconv.Atoi(params[0])
			id := uint(intID)
			tab := u.tabs[id]
			if tab == nil {
				return
			}
			if len(u.tabs) == 1 {
				globalBrowser.app.Quit()
			}
			u.view().EvaluateScript(fmt.Sprintf("closeTab(%d)", id))
			if id != u.activeTabID {
				delete(u.tabs, id)
			} else {
				tab.setReadyToClose(true)
			}
		}
	})
	view.BindJSCallback("OnActiveTabChange", func(v *ultralight.View, params []string) {
		if len(params) == 1 {
			intID, _ := strconv.Atoi(params[0])
			id := uint(intID)
			tab := u.tabs[id]
			if tab == nil {
				return
			}
			u.tabs[u.activeTabID].hide()

			if u.tabs[u.activeTabID].readyToClose() {
				delete(u.tabs, u.activeTabID)
			}

			u.activeTabID = id
			tab.show()

			tabView := tab.view()
			u.setLoading(tabView.IsLoading())
			u.setCanGoBack(tabView.CanGoBack())
			u.setCanGoForward(tabView.CanGoForward())
			u.setURL(tabView.GetURL())
		}
	})
	view.BindJSCallback("OnRequestChangeURL", func(v *ultralight.View, params []string) {
		if len(params) == 1 {
			if len(u.tabs) > 0 {
				u.tabs[u.activeTabID].view().LoadURL(params[0])
			}
		}
	})

	u.createNewTab()
}

func (u *ui) createNewTab() {
	id := u.tabIDCounter
	u.tabIDCounter++
	window := globalBrowser.window
	tabHeight := window.GetHeight() - uiHeight
	if tabHeight < 1 {
		tabHeight = 1
	}
	u.tabs[id] = u.createTab(id, window.GetWidth(), tabHeight, 0, uiHeight)
	u.tabs[id].view().LoadURL("file://assets/new_tab_page.html")
	u.view().EvaluateScript(fmt.Sprintf("addTab(%d,\"%s\",\"%s\")", id, "New Tab", ""))
}

func (u *ui) updateTabTitle(id uint, title string) {
	u.view().EvaluateScript(fmt.Sprintf("updateTab(%d,%q,%q)", id, title, ""))
}

func (u *ui) updateTabURL(id uint, url string) {
	if id == u.activeTabID && len(u.tabs) > 0 {
		u.setURL(url)
	}
}

func (u *ui) updateTabNavigation(id uint, isLoading, canGoBack, canGoForward bool) {
	if id == u.activeTabID && len(u.tabs) > 0 {
		u.setLoading(isLoading)
		u.setCanGoBack(canGoBack)
		u.setCanGoForward(canGoForward)
	}
}

func (u *ui) setLoading(isLoading bool) {
	u.view().EvaluateScript(fmt.Sprintf("updateLoading(%t)", isLoading))
}

func (u *ui) setCanGoBack(canGoBack bool) {
	u.view().EvaluateScript(fmt.Sprintf("updateBack(%t)", canGoBack))
}

func (u *ui) setCanGoForward(canGoForward bool) {
	u.view().EvaluateScript(fmt.Sprintf("updateForward(%t)", canGoForward))
}

func (u *ui) setURL(url string) {
	u.view().EvaluateScript(fmt.Sprintf("updateURL(%q)", url))
}

func (u *ui) setCursor(cursor ultralight.Cursor) {
	globalBrowser.window.SetCursor(cursor)
}

func (u *ui) activeTab() *tab {
	return u.tabs[u.activeTabID]
}

func (u *ui) view() *ultralight.View {
	return u.overlay.GetView()
}
