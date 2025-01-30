package systray_test

import (
	"testing"
	"time"

	"github.com/mxmauro/go-systray"
	"github.com/mxmauro/go-systray/assets"
)

// -----------------------------------------------------------------------------

func TestSystray(t *testing.T) {
	systray.Run(func() {
		onReady(t)
	}, func() {
		now := time.Now()
		t.Log("Exit at " + now.String())
	})
}

func onReady(t *testing.T) {
	t.Log("onReady called")

	systray.SetTitle("SysTray Test")
	systray.SetTooltip("Hello! I'm a tooltip")
	systray.SetOnClick(func(menu systray.IMenu) {
		if menu != nil { // menu for linux nil
			menu.ShowMenu()
		}
		t.Log("SetOnClick called")
	})
	systray.SetOnDClick(func(menu systray.IMenu) {
		if menu != nil { // menu for linux nil
			menu.ShowMenu()
		}
		t.Log("SetOnDClick called")
	})
	// OnRClick linux not impl
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
		t.Log("SetOnRClick called")
	})

	systray.CreateMenu()
	addQuitItem(t)
	mChange := systray.AddMenuItem("Change Me", "Change Me")
	mChecked := systray.AddMenuItemCheckbox("Checked", "Check Me", true)
	mEnabled := systray.AddMenuItem("Enabled", "Enabled")
	// Sets the icon of a menu item. Only available on Mac.
	mEnabled.SetIcon(assets.CheckCircleImage)

	systray.AddMenuItem("Ignored", "Ignored")

	subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
	subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
	subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")
	subMenuBottom2.SetIcon(assets.SadDizzyImage)
	systray.AddSeparator()
	mToggle := systray.AddMenuItem("Toggle", "Toggle some menu items")
	shown := true
	toggle := func() {
		if shown {
			subMenuBottom.Check()
			subMenuBottom2.Hide()
			mEnabled.Hide()
			shown = false
			mEnabled.Disable()
		} else {
			subMenuBottom.Uncheck()
			subMenuBottom2.Show()
			mEnabled.Show()
			mEnabled.Enable()
			shown = true
		}
	}
	mReset := systray.AddMenuItem("Reset", "Reset all items")

	mChange.Click(func() {
		mChange.SetTitle("I've changed")
	})
	mChecked.Click(func() {
		if mChecked.Checked() {
			mChecked.Uncheck()
			mChecked.SetTitle("Unchecked")
		} else {
			mChecked.Check()
			mChecked.SetTitle("Checked")
		}
	})
	mEnabled.Click(func() {
		mEnabled.SetTitle("Disabled")
		t.Log("mEnabled.Click() called", mEnabled.Disabled())
		mEnabled.Disable()
	})
	subMenuBottom2.Click(func() {
		panic("panic button pressed")
	})
	subMenuBottom.Click(func() {
		toggle()
	})
	mReset.Click(func() {
		systray.ResetMenu()
		addQuitItem(t)
	})
	mToggle.Click(func() {
		toggle()
	})

	// tray icon switch
	go func() {
		idx := 0

		for {
			idx += 1
			if idx > 4 {
				idx = 1
			}
			switch idx {
			case 1:
				systray.SetIcon(assets.Wifi1Image)
			case 2:
				systray.SetIcon(assets.Wifi2Image)
			case 3:
				systray.SetIcon(assets.Wifi3Image)
			case 4:
				systray.SetIcon(assets.Wifi2Image)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func addQuitItem(t *testing.T) {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	mQuit.Click(func() {
		t.Log("mQuit.Click() called")
		systray.Quit()
		//systray.Quit()// macos error
		//end() // macos error
		t.Log("Finished quitting")
	})
}
