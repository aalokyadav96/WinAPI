package main

import (
    "golang.org/x/sys/windows/registry"
    "log"
)

func main() {
k, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE|registry.SET_VALUE)
if err != nil {
    log.Fatal(err)
}
if err := k.SetStringValue("xyz", "blahblah"); err != nil {
    log.Fatal(err)
}
if err := k.Close(); err != nil {
    log.Fatal(err)
}
}

//HKEY_CURRENT_USER\Control Panel\Desktop
//Wallpaper
//WallpaperOriginX
//WallpaperOriginX
//WallpaperStyle
