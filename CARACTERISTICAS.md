# Resumen del Proyecto: Generador de Certificados

## 📋 Descripción
Programa en Go para generar automáticamente certificados de participación en eventos a partir de datos en formato JSON.

## ✨ Características principales

### Funcionalidades
- ✅ Generación masiva de certificados en PDF (formato A4 horizontal)
- ✅ Personalización completa mediante archivos JSON
- ✅ Soporte para imágenes de fondo personalizadas
- ✅ Soporte para imágenes de firma digital
- ✅ Colores personalizables (RGB) para títulos y texto
- ✅ Múltiples tipos de participación (ponente, autor, asistente, etc.)
- ✅ Campos opcionales que se omiten si están vacíos
- ✅ Parámetros de línea de comandos para mayor flexibilidad
- ✅ Sanitización automática de nombres de archivo
- ✅ Creación automática de directorios de salida
- ✅ Manejo de errores robusto
- ✅ Contador de éxitos y errores

### Datos que soporta
**Del evento:**
- Nombre del evento
- Título del certificado personalizable
- Ubicación
- Fechas (inicio y fin)
- ISBN (opcional)
- Organizadores (opcional)
- Datos del firmante (nombre, cargo, imagen)

**De los participantes:**
- Nombre completo
- Afiliación institucional (opcional)
- Tipo de participación (ponente, autor, asistente, etc.)
- Título del trabajo presentado (opcional)

## 📁 Estructura del proyecto

```
certificate/
├── main.go                           # Código principal
├── go.mod                            # Dependencias Go
├── go.sum                            # Checksums de dependencias
├── generador-certificados.exe       # Ejecutable compilado
│
├── certificados.json                # Configuración principal
├── ejemplo-simple.json              # Ejemplo simplificado
│
├── README.md                        # Documentación completa
├── QUICK_START.md                   # Guía rápida de inicio
├── PLANTILLA.md                     # Plantillas y ejemplos
│
├── generar.bat                      # Script de menú (Windows)
├── test.ps1                         # Script de prueba (PowerShell)
│
└── certificados/                    # Carpeta de salida (generada)
```

## 🚀 Uso rápido

### Opción 1: Ejecutar directamente
```bash
go run main.go
```

### Opción 2: Compilar y ejecutar
```bash
go build -o generador-certificados.exe
.\generador-certificados.exe
```

### Opción 3: Con parámetros personalizados
```bash
.\generador-certificados.exe -json datos.json -bg fondo.png -output salida
```

### Opción 4: Script de menú (Windows)
```bash
.\generar.bat
```

### Opción 5: Test rápido (PowerShell)
```powershell
.\test.ps1
```

## 📊 Parámetros de línea de comandos

| Parámetro | Descripción | Valor por defecto |
|-----------|-------------|-------------------|
| `-json` | Archivo JSON con los datos | `certificados.json` |
| `-bg` | Imagen de fondo (sobrescribe JSON) | Del JSON o `background.png` |
| `-output` | Directorio de salida (sobrescribe JSON) | Del JSON o `certificados` |

## 🎨 Personalización

### Colores predeterminados
- **Texto**: RGB(50, 50, 50) - Gris oscuro
- **Títulos**: RGB(0, 51, 102) - Azul corporativo

### Fuentes
- **Título principal**: Arial Bold, 36pt
- **Nombre del participante**: Arial Bold, 28pt
- **Nombre del evento**: Arial Bold, 16pt
- **Texto general**: Arial Regular, 11-14pt

### Formato de salida
- **Formato**: PDF
- **Tamaño**: A4 (297 × 210 mm)
- **Orientación**: Horizontal (landscape)
- **Nombre de archivo**: `certificado_<nombre_participante>.pdf`

## 🔧 Requisitos técnicos

### Software necesario
- Go 1.20 o superior
- Sistema operativo: Windows, Linux o macOS

### Dependencias
- `github.com/jung-kurt/gofpdf` v1.16.2 (instalación automática con `go mod tidy`)

## 📝 Campos del JSON

### Obligatorios
- `event.name`
- `event.signature_name`
- `event.signature_role`
- `participants[].name`

### Opcionales
- Todo lo demás (se omite si está vacío)

## 🎯 Casos de uso

1. **Congresos académicos**: Certificados de ponencias y asistencias
2. **Talleres**: Certificados de participación
3. **Eventos corporativos**: Certificados de capacitación
4. **Conferencias virtuales**: Certificados de asistencia online
5. **Publicaciones**: Certificados de autoría de artículos

## 🔒 Validaciones y seguridad

- ✅ Sanitización de nombres de archivo
- ✅ Verificación de existencia de imágenes
- ✅ Manejo de errores por participante (no detiene el proceso)
- ✅ Creación segura de directorios
- ✅ Validación de formato JSON

## 📈 Rendimiento

- Genera ~1 certificado por segundo (depende del hardware)
- Sin límite de participantes
- Uso de memoria eficiente
- Procesamiento secuencial confiable

## 🛠️ Mejoras futuras posibles

- [ ] Soporte para múltiples fuentes personalizadas
- [ ] Plantillas de diseño predefinidas
- [ ] Generación de códigos QR de verificación
- [ ] Exportación a otros formatos (PNG, JPG)
- [ ] Interfaz gráfica (GUI)
- [ ] API REST para integración
- [ ] Importación desde CSV/Excel
- [ ] Envío automático por email
- [ ] Firma digital criptográfica

## 📄 Licencia

Código abierto - Libre para uso y modificación

## 👤 Soporte

Para problemas o sugerencias:
1. Revisa el README.md completo
2. Verifica que go.mod esté actualizado: `go mod tidy`
3. Asegúrate de que las rutas a las imágenes sean correctas
4. Verifica que el JSON tenga formato válido

## 🎓 Ejemplo de flujo de trabajo

1. **Preparación**: Editar `certificados.json` con datos del evento
2. **Opcional**: Preparar imagen de fondo y firma
3. **Ejecución**: `go run main.go`
4. **Resultado**: Certificados en carpeta `certificados/`
5. **Distribución**: Enviar PDFs a participantes

## ⏱️ Tiempo estimado

- Configuración inicial: 10-15 minutos
- Por cada evento adicional: 5 minutos
- Generación de 100 certificados: ~2 minutos
