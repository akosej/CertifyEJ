# ⚡ Inicio Rápido en 3 Pasos

## 🎯 Para empezar YA

### Paso 1️⃣: Edita tus datos
Abre `certificados.json` y cambia:
- Nombre del evento
- Fechas
- Datos del firmante
- Lista de participantes

### Paso 2️⃣: Ejecuta
```bash
go run main.go
```

o si ya compilaste:
```bash
.\generador-certificados.exe
```

### Paso 3️⃣: Listo!
Revisa la carpeta `certificados/` 🎉

---

## 📚 Documentación disponible

| Archivo | ¿Para qué sirve? |
|---------|------------------|
| **QUICK_START.md** | Guía de inicio rápido detallada |
| **README.md** | Documentación completa del programa |
| **PLANTILLA.md** | Plantillas y ejemplos de JSON |
| **CARACTERISTICAS.md** | Resumen de características del proyecto |
| **TROUBLESHOOTING.md** | Solución de problemas comunes |

---

## 🎮 Ejemplos de uso

### Usar archivo diferente
```bash
go run main.go -json mi-evento.json
```

### Cambiar imagen de fondo
```bash
go run main.go -bg mi-fondo.png
```

### Cambiar carpeta de salida
```bash
go run main.go -output mis-certificados
```

### Todo junto
```bash
go run main.go -json evento.json -bg fondo.jpg -output salida
```

---

## 🎨 Personalización rápida

### Cambiar colores del título
En `certificados.json`:
```json
"title_color": { "r": 139, "g": 0, "b": 0 }
```

### Cambiar tipo de certificado
```json
"certificate_title": "CERTIFICADO DE ASISTENCIA"
```

### Quitar campos opcionales
Déjalos vacíos:
```json
"work_title": "",
"isbn": "",
"affiliation": ""
```

---

## 🚀 Comandos útiles

### Primera vez
```bash
go mod tidy
go run main.go
```

### Compilar
```bash
go build -o certifEJ.exe
```

### Test rápido
```bash
.\test.ps1
```

### Menú interactivo (Windows)
```bash
.\generar.bat
```

---

## ✅ Checklist antes de generar

- [ ] Go 1.20+ instalado
- [ ] Ejecutaste `go mod tidy`
- [ ] Editaste `certificados.json` con tus datos
- [ ] (Opcional) Tienes `background.png` preparado
- [ ] (Opcional) Tienes imagen de firma preparada

---

## 🎁 Archivos de ejemplo incluidos

- **certificados.json**: Ejemplo completo con 4 participantes
- **ejemplo-simple.json**: Ejemplo minimalista con 1 participante

Usa cualquiera como punto de partida!

---

## 💡 Tip profesional

Para generar rápido certificados de prueba:
```bash
go run main.go -json ejemplo-simple.json -output test
```

Revisa el resultado, y cuando estés satisfecho, usa tu archivo real.

---

**¿Tienes problemas?** → Lee `TROUBLESHOOTING.md`

**¿Quieres personalizar más?** → Lee `README.md` completo

**¡Disfruta generando certificados!** 🎓✨
