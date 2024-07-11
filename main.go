package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    var licenseType, name, projectPath string
    flag.StringVar(&licenseType, "type", "MIT", "Tipo de licencia (MIT, Apache, GPL)")
    flag.StringVar(&name, "name", "Your Name", "Nombre del propietario de la licencia")
    flag.StringVar(&projectPath, "path", ".", "Ruta del proyecto donde se creará el archivo LICENSE")
    flag.Parse()

    licenseContent := generateLicense(licenseType, name)

    licenseFilePath := filepath.Join(projectPath, "LICENSE")
    err := os.WriteFile(licenseFilePath, []byte(licenseContent), 0644)
    if err != nil {
        fmt.Println("Error creando el archivo LICENSE:", err)
        return
    }

    fmt.Println("Archivo LICENSE creado en:", licenseFilePath)
}

func generateLicense(licenseType, name string) string {
    switch licenseType {
    case "MIT", "mit":
        return fmt.Sprintf(`MIT License

    Copyright (c) %s

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.`, name)
    case "Apache", "apache":
        return fmt.Sprintf(`Apache License 2.0

    Copyright (c) %s

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.`, name)
    case "GPL", "gpl":
        return fmt.Sprintf(`GNU GENERAL PUBLIC LICENSE
    Version 3, 29 June 2007

    Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>

    Everyone is permitted to copy and distribute verbatim copies
    of this license document, but changing it is not allowed.

    %s`, name)
    default:
        return "Tipo de licencia no soportado."
    }
}
