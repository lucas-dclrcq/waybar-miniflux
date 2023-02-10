# Waybar Miniflux Applet

This small tool fetch unread counter from miniflux and format it for [Waybar](https://github.com/Alexays/Waybar)

## How to

1. Install cli : `go install git@github.com:lucas-dclrcq/waybar-miniflux.git`
2. Add the following to your waybar configuration
```json
  "custom/miniflux": {
    "exec": "~/go/bin/miniflux --url MINIFLUX_URL --token MINIFLUX_API_TOKEN",
    "format": "{} \uf09e",
    "tooltip": true,
    "return-type": "json",
    "interval": 300
  },
```
3. Reload waybar