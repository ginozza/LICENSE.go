# CLI de Generación de Licencias

Este es un CLI simple desarrollado en Go para generar archivos de licencia para proyectos. Permite a los usuarios especificar el tipo de licencia (MIT, Apache, GPL), el nombre del propietario de la licencia y la ubicación del proyecto donde se creará el archivo LICENSE.

### Uso

A continuación se muestra cómo usar el CLI desde la línea de comandos:

```bash
# Instalar dependencias (si es necesario)
go mod tidy

# Compilar el CLI
go build -o license-cli

# Ejecutar el CLI para generar una licencia MIT
./license-cli -type MIT -name "Tu Nombre" -path /ruta/a/tu/proyecto
```

Por defecto el tipo de licenia es MIT, y la ruta por defecto es la carpeta donde abres la terminal. Puedes cambiar el tipo de licencia, tu nombre, y ruta por defecto, para más comodidad desde el archivo `main.go`.

### Opciones de licencia soportadas
- MIT: Licencia estándar que permite el uso, copia, modificación, fusión, publicación, distribución, sublicencia y venta del software.
- Apache: Licencia que permite el uso, modificación, distribución y sublicencia del software bajo términos específicos.
- GPL: Licencia que asegura que el software sea libre para todos sus usuarios.

### Contribución
Contribución

Si quieres contribuir a este repositorio con nuevas funcionalidades, o mejoras, por favor sigue estos pasos:

1. Haz un fork del repositorio y clónalo localmente.
2. Crea una nueva rama (git checkout -b feature/mejora).
3. Haz tus cambios y haz commit (git commit -am 'Añadir una mejora').
4. Sube la rama (git push origin feature/mejora).
5. Abre un Pull Request.