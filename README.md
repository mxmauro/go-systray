# go-systray

A cross-platform Go library to place an icon and menu in the notification area.

This repository is a fork of [energye/systray](https://github.com/energye/systray) which,
in turn, is a fork of [getlantern/systray](https://github.com/getlantern/systray).

## Features

* Supported on Windows, macOS, Linux and many BSD systems
* Menu items can be checked and/or disabled
* Methods may be called from any Goroutine
* tray icon supports mouse click, double click, and right click

## API

```go
package main

import (
    "fmt"

    "github.com/mxmauro/go-systray"
    "github.com/mxmauro/go-systray/assets"
)

func main() {
    systray.Run(onReady, onExit)
}

func onReady() {
    systray.SetIcon(assets.Wifi3Image)
    systray.SetTitle("Awesome App")
    systray.SetTooltip("Pretty awesome超级棒")
    systray.SetOnClick(func (menu systray.IMenu) {
        fmt.Println("SetOnClick")
    })
    systray.SetOnDClick(func (menu systray.IMenu) {
        fmt.Println("SetOnDClick")
    })
    systray.SetOnRClick(func (menu systray.IMenu) {
        menu.ShowMenu()
        fmt.Println("SetOnRClick")
    })
    mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

    // Sets the icon of a menu item.
    mQuit.SetIcon(assets.SadDizzyImage)
}

func onExit() {
    // clean up here
}
```

### Run in another toolkit

Most graphical toolkits will grab the main loop so the `Run` code above is not possible.
For this reason there is another entry point `RunWithExternalLoop`.
This function of the library returns a start and end function that should be called
when the application has started and will end, to loop in appropriate features.

Note: this package requires cgo, so make sure you set `CGO_ENABLED=1` before building.

## Platform notes

### Linux/BSD

This implementation uses DBus to communicate through the SystemNotifier/AppIndicator spec,
older tray implementations may not load the icon.

If you are running an older desktop environment, or system tray provider, you may require
a proxy app which can convert the new DBus calls to the old format.
The recommended tool for Gnome based trays is [snixembed](https://git.sr.ht/~steef/snixembed),
others are available. Search for "StatusNotifierItems XEmbedded" in your package manager.

### Windows

```sh
go build -ldflags -H=windowsgui
```

### macOS

On macOS, you will need to create an application bundle to wrap the binary; simply use add
folders with the following minimal structure and assets:

```
SystrayApp.app/
  Contents/
    Info.plist
    MacOS/
      go-executable
    Resources/
      SystrayApp.icns
```

If bundling manually, you may want to add one or both of the following to your Info.plist:

```xml
	<!-- avoid having a blurry icon and text -->
	<key>NSHighResolutionCapable</key>
	<string>True</string>

	<!-- avoid showing the app on the Dock -->
	<key>LSUIElement</key>
	<string>1</string>
```

Consult the [Official Apple Documentation here](https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/BundleTypes/BundleTypes.html#//apple_ref/doc/uid/10000123i-CH101-SW1).

## Credits

- https://github.com/energye/systray
- https://github.com/getlantern/systray
- https://github.com/xilp/systray
- https://github.com/cratonica/trayhost
- https://www.iconfinder.com/unicons-iconscout
