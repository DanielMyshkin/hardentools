/*
    Hardentools
    Copyright (C) 2017  Claudio Guarnieri, Mariano Graziano

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "fmt"
    "golang.org/x/sys/windows/registry"
)

var adobe_versions = []string{
    "DC", // Acrobat Reader DC
    "XI", // Acrobat Reader XI - To test
}

/*
bEnableJS possible values:
0 - Disable AcroJS
1 - Enable AcroJS
*/

func trigger_pdf_js(enable bool) {
    var value uint32

    if enable {
        // Enable 
        fmt.Println("[*] Enabling Acrobat Reader JavaScript")
        value = 1
    } else {
        // Disable Packager
        fmt.Println("[*] Disabling Acrobat Reader JavaScript")
        value = 0
    }

    for _, adobe_version := range adobe_versions {
        path := fmt.Sprintf("SOFTWARE\\Adobe\\Acrobat Reader\\%s\\JSPrefs", adobe_version)
        key, err := registry.OpenKey(registry.CURRENT_USER, path, registry.WRITE)
        // Check
        if err == nil {
            key1, _, _ := registry.CreateKey(registry.CURRENT_USER, path, registry.WRITE)
            key = key1
        }
        key.SetDWordValue("bEnableJS", value)
        key.Close()
    }
}

/*
bAllowOpenFile set to 0 and 
bSecureOpenFile set to 1 to disable
the opening of non-PDF documents
*/

func trigger_pdf_objects(enable bool) {
    var allow_value uint32
    var secure_value uint32

    if enable {
        // Enable 
        fmt.Println("[*] Enabling the opening of objects embedded in PDF documents")
        allow_value = 1
        secure_value = 0
    } else {
        // Disable
        fmt.Println("[*] Disabling the opening of objects embedded in PDF documents")
        allow_value = 0
        secure_value = 1
    }

    for _, adobe_version := range adobe_versions {
        path := fmt.Sprintf("SOFTWARE\\Adobe\\Acrobat Reader\\%s\\Originals", adobe_version)
        key, err := registry.OpenKey(registry.CURRENT_USER, path, registry.WRITE)

        // Check
        if err == nil {
            key1, _, _ := registry.CreateKey(registry.CURRENT_USER, path, registry.WRITE)
            key = key1
        }
        key.SetDWordValue("bAllowOpenFile", allow_value)
        key.SetDWordValue("bSecureOpenFile", secure_value)
        key.Close()
    }
}
