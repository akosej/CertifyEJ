# Solución de Problemas y Preguntas Frecuentes

## 🐛 Errores Comunes y Soluciones

### 1. Error: "could not import github.com/jung-kurt/gofpdf"

**Problema**: Las dependencias no están instaladas.

**Solución**:
```bash
go mod tidy
```

o

```bash
go mod download
```

---

### 2. Error: "Error leyendo JSON 'certificados.json'"

**Problema**: El archivo JSON no existe o la ruta es incorrecta.

**Solución**:
- Verifica que el archivo exista en el directorio actual
- Usa la ruta completa: `-json "C:\ruta\completa\datos.json"`
- Verifica que el nombre del archivo esté correcto

---

### 3. Error: "Error parseando JSON"

**Problema**: El JSON tiene errores de sintaxis.

**Solución**:
- Verifica que todas las comas estén en su lugar
- Verifica que todas las llaves `{}` y corchetes `[]` estén balanceados
- Usa un validador JSON online (jsonlint.com)
- Revisa que los campos RGB tengan valores numéricos

**Ejemplo de JSON inválido**:
```json
{
  "event": {
    "name": "Mi Evento",  // ❌ Sin coma al final
  }
}
```

**JSON válido**:
```json
{
  "event": {
    "name": "Mi Evento"
  }
}
```

---

### 4. Los certificados salen en blanco o sin texto

**Problema**: Colores blancos sobre fondo blanco.

**Solución**:
- Cambia los colores a valores oscuros:
```json
"font_color": { "r": 0, "g": 0, "b": 0 }
```

---

### 5. La imagen de fondo no aparece

**Problema**: La ruta a la imagen es incorrecta o el archivo no existe.

**Solución**:
- Verifica que el archivo exista: `Test-Path background.png`
- Usa rutas absolutas: `"C:\\imagenes\\fondo.png"`
- Usa barras dobles en Windows: `\\` o barras simples hacia adelante: `/`
- El programa continuará sin fondo si no encuentra la imagen

---

### 6. Error: "error creando directorio de salida"

**Problema**: No hay permisos para crear carpetas.

**Solución**:
- Ejecuta como administrador
- Usa una ruta donde tengas permisos de escritura
- Verifica el espacio en disco

---

### 7. El texto se sale del certificado o se ve mal

**Problema**: Textos muy largos.

**Solución**:
- Acorta los títulos de trabajos
- Divide nombres largos de eventos
- Edita la función `generateCertificate()` en `main.go` para ajustar tamaños

---

### 8. Caracteres especiales se ven mal (á, é, ñ, etc.)

**Problema**: La biblioteca gofpdf usa fuentes básicas.

**Solución actual**:
- Las fuentes Arial estándar de gofpdf tienen soporte limitado para acentos
- Los caracteres españoles básicos (á, é, í, ó, ú, ñ) deberían funcionar

**Solución avanzada** (requiere modificar código):
- Usar fuentes TTF personalizadas con soporte UTF-8
- Ver documentación de gofpdf sobre AddUTF8Font

---

## ❓ Preguntas Frecuentes

### ¿Puedo generar certificados sin imagen de fondo?

Sí, simplemente deja el campo vacío o pon una ruta inexistente:
```json
"background_image": ""
```

Se usará un fondo de color sólido.

---

### ¿Puedo usar el programa en Linux o Mac?

Sí, es multiplataforma. Compila así:

**Linux/Mac**:
```bash
go build -o generador-certificados main.go
./generador-certificados
```

---

### ¿Cómo cambio el tamaño del certificado?

Edita `main.go`, línea donde se crea el PDF:
```go
pdf := gofpdf.New("L", "mm", "A4", "")
```

Cambiar `"A4"` por `"Letter"` para tamaño carta.

---

### ¿Puedo añadir un logo al certificado?

Sí, edita la función `generateCertificate()` y añade:
```go
pdf.Image("logo.png", 10, 10, 30, 0, false, "", 0, "")
```

Parámetros: (archivo, x, y, ancho, alto, ...)

---

### ¿Cómo cambio el tamaño de las fuentes?

Edita `main.go`, busca las líneas con `SetFont()`:
```go
pdf.SetFont("Arial", "B", 36)  // 36 es el tamaño
```

---

### ¿Puedo generar certificados de diferentes eventos en una sola ejecución?

No directamente. Ejecuta el programa varias veces:
```bash
.\generador-certificados.exe -json evento1.json -output evento1
.\generador-certificados.exe -json evento2.json -output evento2
```

---

### ¿Cómo importo participantes desde Excel?

1. Exporta tu Excel a CSV
2. Usa un conversor CSV a JSON online o un script Python:

```python
import csv
import json

participants = []
with open('participantes.csv', 'r', encoding='utf-8') as f:
    reader = csv.DictReader(f)
    for row in reader:
        participants.append({
            "name": row['nombre'],
            "affiliation": row['institucion'],
            "participation_type": row['tipo'],
            "work_title": row['trabajo']
        })

data = {
    "event": { /* ... */ },
    "participants": participants
}

with open('certificados.json', 'w', encoding='utf-8') as f:
    json.dump(data, f, indent=2, ensure_ascii=False)
```

---

### ¿Los PDFs tienen seguridad o marcas de agua?

No por defecto. Puedes añadirlas editando el código:

```go
// Marca de agua
pdf.SetTextColor(200, 200, 200)
pdf.SetFont("Arial", "", 60)
pdf.RotatedText(80, 100, "MUESTRA", 45)
```

---

### ¿Puedo añadir códigos QR de verificación?

Sí, pero requiere instalar una biblioteca adicional:
```bash
go get github.com/boombuler/barcode
go get github.com/boombuler/barcode/qr
```

Luego editar el código para generar y añadir QR codes.

---

### ¿Cómo personalizo completamente el diseño?

Edita la función `generateCertificate()` en `main.go`. Es donde se define:
- Posiciones (X, Y)
- Tamaños de fuente
- Colores
- Textos

Usa la documentación de gofpdf: https://pkg.go.dev/github.com/jung-kurt/gofpdf

---

### ¿Hay límite de participantes?

No hay límite técnico. Se ha probado con miles de participantes sin problemas.

---

### ¿Puedo firmar digitalmente los PDFs?

La biblioteca gofpdf no soporta firmas digitales nativas. Necesitarías:
- Usar otra biblioteca como `unidoc/unipdf` (de pago)
- O firmar los PDFs después con herramientas externas

---

## 🔍 Debug y Diagnóstico

### Ver versión de Go
```bash
go version
```

### Verificar dependencias
```bash
go mod verify
```

### Ver módulos instalados
```bash
go list -m all
```

### Compilar con información de debug
```bash
go build -v -x -o generador-certificados.exe
```

### Ejecutar con más información
```bash
$ErrorActionPreference = "Continue"
.\generador-certificados.exe
```

---

## 📧 Reportar problemas

Si encuentras un problema no listado aquí:

1. Verifica que tienes la última versión de Go
2. Ejecuta `go mod tidy`
3. Intenta con el archivo `ejemplo-simple.json`
4. Documenta el error exacto y los pasos para reproducirlo
5. Incluye tu versión de Go: `go version`

---

## 💡 Tips y Trucos

### Generación rápida para testing
```bash
# Crea un JSON con solo 1 participante para pruebas rápidas
.\generador-certificados.exe -json ejemplo-simple.json -output test
```

### Backup automático
```powershell
# Copia de seguridad antes de generar
Copy-Item certificados.json certificados_backup.json
```

### Generación por lotes
```powershell
# Script para múltiples eventos
$eventos = @("evento1.json", "evento2.json", "evento3.json")
foreach ($evento in $eventos) {
    $nombre = [System.IO.Path]::GetFileNameWithoutExtension($evento)
    .\generador-certificados.exe -json $evento -output "output_$nombre"
}
```

### Verificar JSON antes de generar
```powershell
# Validar JSON en PowerShell
Get-Content certificados.json | ConvertFrom-Json
```
