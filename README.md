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
dwmbar --noBrightness -noCpu -noLang -noMemory -noNetworkStats -noNetworkState -noPowerState -noTemp -noVolume
```

## Screenshot:

![Demo](screenshots/demo.png)