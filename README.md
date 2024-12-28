# DWMBAR

## How to build?

> make build

* required dwm-status2d patch
* nerd fonts package(media-fonts/nerd-fonts-symbols on gentoo)

## How to run on dwm start

Add `~/.dwm/bin/dwmbar` to x init file

## Run example

Show only time:
```shell
# Russian
dwmbar --noBrightness -noCpu -noLang -noMemory -noNetworkStats -noNetworkState -noPowerState -noTemp -noVolume -noNotificationsState -enableNotificationsStateBgBlinking --lang=ru

# English
dwmbar --noBrightness -noCpu -noLang -noMemory -noNetworkStats -noNetworkState -noPowerState -noTemp -noVolume -noNotificationsState -enableNotificationsStateBgBlinking --lang=en
```

## Screenshot:

![Demo](screenshots/demo.png)

## Structure:

- [drawer](drawer) - draw final bar
- [drawer_templates](drawer_templates) - templates for drawer use `fmt.Sprintf`
- [drawer_theme](drawer_theme) - utility function for create new drawer theme
- [metrics_collector](metrics_collector) - here we collect whole information using [state_providers](state_providers)
- [snapshot](snapshot) - just for "current" collected infos(collects new data every time)
- [state_providers](state_providers) - here infrastructure level getters, use specific commands or libraries
- [themes](themes) - predefined, created by me themes
- [util](util) - utility functions, like shared layer

## Flow

We have 2 separate modules: [metrics_collector](metrics_collector) and [drawer](drawer)

* [metrics_collector](metrics_collector) -> [state_providers](state_providers) -> OS(Getting information)
* [drawer](drawer) -> [drawer_templates](drawer_templates) + [drawer_theme](drawer_theme) -> OS(Write bar)

## Specific

Here described all library, commands and platform choices of this tool

### State providers

* [battery_state](state_providers/battery_state) - use library `github.com/distatus/battery`
* [brightness_state](state_providers/brightness_state) - use command `brightnessctl`
* [cpu_stat](state_providers/cpu_stat) - use library `github.com/mackerelio/go-osstat/cpu`
* [cpu_temp](state_providers/cpu_temp) - use library `github.com/ssimunic/gosensors`
* [keyboard_layout](state_providers/keyboard_layout) - use command `xkblayout-state`
* [network_connection_state](state_providers/network_connection_state) - use directory `/sys/class/net/` lookup and parsing for wired connections and `iwctl` for wireless connections
* [network_stat](state_providers/network_stat) use library `github.com/mackerelio/go-osstat/network`
* [notifications_state](state_providers/notifications_state) - use command `dunstctl`
* [volume_state](state_providers/volume_state) - use library `github.com/itchyny/volume-go`

### Drawer

* `xsetroot` - for write result string in dwm bar section

## Platform specific

* `linux`, `x11`, `dwm` - just for formal list
* `iwd` - wireless network daemon
* `dunst` - notifications daemon
* `media-fonts/nerd-fonts-symbols` on gentoo
* `dwm-status2d` patch
