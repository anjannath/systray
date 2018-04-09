Package systray is a cross platfrom Go library to place an icon and menu in the notification area.
Tested on Windows XP,7,10 and Linux Mint 18.3.
## Usage
```go
func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app", systray.ItemDefault)
}

func onExit() {
	// clean up here
}
```
Menu item can be checked and / or disabled. Methods except `Run()` can be invoked from any goroutine. See demo code under `example` folder.

## Platform specific concerns

### Linux

```sh
sudo apt-get install libgtk-3-dev libappindicator3-dev
```
Checked menu item not implemented on Linux yet.

Warning: GtkStatusIcon has been deprecated in GTK 3.14!


## Prepare icons

You can use `filetobyte` utility at `cmd` folder. Two ways:
1. Do
```sh
go get github.com/amkulikov/systray/cmd/filetobytes
```
and after that use compiled binary directly, or use code generation in your project
```
//go:generate filetobytes -package=tray -dest=tray/data.go path/to/iconfile/or/iconsfolder
```

2. Use `go run` to execute util.
```
//go:generate go run github.com/amkulikov/systray/cmd/filetobytes/main.go -package=tray -dest=tray/data.go example/icon
```

It will generate file `data.go` in `tray` folder that will contain
```
var filesByteData = map[string][]byte{...}
```
where key is a path to sourcefile and value is a binary representation of resource.


## Credits

- https://github.com/xilp/systray
- https://github.com/cratonica/trayhost
- https://github.com/getlantern/systray
