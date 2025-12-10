# Generador de Certificados para Eventos

Programa en Go para generar certificados de participación en eventos de forma automática a partir de datos en formato JSON.

## Características

- ✅ Generación masiva de certificados en formato PDF
- ✅ Personalización completa mediante JSON
- ✅ Soporte para imagen de fondo personalizada
- ✅ Múltiples tipos de participación (ponente, autor, asistente)
- ✅ Colores personalizables para títulos y texto
- ✅ Inclusión de datos completos (ISBN, organizadores, fechas, afiliación)
- ✅ Firma digital opcional con imagen
- ✅ Parámetros por línea de comandos

## Requisitos

- Go 1.20 o superior
- Biblioteca `gofpdf` (se instala automáticamente)

## Instalación

1. Clona o descarga este repositorio
2. Instala las dependencias:

```bash
go mod download
```

o simplemente:

```bash
go mod tidy
```

## Estructura del JSON

El archivo `certificados.json` debe contener:

```json
{
  "background_image": "background.png",
  "output_directory": "certificados",
  "font_color": {
    "r": 50,
    "g": 50,
    "b": 50
  },
  "title_color": {
    "r": 0,
    "g": 51,
    "b": 102
  },
  "event": {
    "name": "Nombre del Evento",
    "certificate_title": "CERTIFICADO DE PARTICIPACIÓN",
    "location": "Ciudad, País",
    "start_date": "14 de junio de 2025",
    "end_date": "17 de junio de 2025",
    "isbn": "ISBN-978-959-16-1234-5",
    "organizers": "Organizadores del evento",
    "signature_name": "Nombre del Firmante",
    "signature_role": "Cargo del Firmante",
    "signature_image": "firma.png",
    "signature_width": 40,
    "logo_image": "./logo.png",
    "logo_width": 20,
    "logo_x": 20,
    "logo_y": 15
  },
  "participants": [
    {
      "name": "Nombre Completo",
      "affiliation": "Universidad o Institución",
      "participation_type": "ponente",
      "work_title": "Título del trabajo presentado"
    }
  ]
}
```

### Campos del JSON

#### Configuración general:
- **background_image**: Ruta a la imagen de fondo (PNG o JPG). Opcional.
- **output_directory**: Carpeta donde se guardarán los certificados generados.
- **font_color**: Color RGB del texto principal (valores 0-255).
- **title_color**: Color RGB de los títulos (valores 0-255).

#### Datos del evento:
- **name**: Nombre completo del evento.
- **certificate_title**: Título que aparecerá en el certificado (ej: "CERTIFICADO DE PARTICIPACIÓN").
- **location**: Lugar donde se celebró el evento.
- **start_date** / **end_date**: Fechas del evento (formato texto libre).
- **isbn**: Código ISBN del evento (opcional).
- **organizers**: Entidades organizadoras (opcional).
- **signature_name**: Nombre de quien firma.
- **signature_role**: Cargo de quien firma.
- **signature_image**: Ruta a imagen de la firma (opcional).

#### Datos de participantes:
- **name**: Nombre completo del participante.
- **affiliation**: Institución o afiliación (opcional).
- **participation_type**: Tipo de participación (ej: "ponente", "autor", "asistente").
- **work_title**: Título del trabajo presentado (opcional, se omite si está vacío).

## Uso

### Uso básico:

```bash
go run main.go
```

Esto leerá `certificados.json` y generará los certificados en la carpeta `certificados/`.

### Con parámetros personalizados:

```bash
go run main.go -json datos.json -bg fondo.png -output mis_certificados
```

**Parámetros disponibles:**
- `-json`: Archivo JSON con los datos (por defecto: `certificados.json`)
- `-bg`: Imagen de fondo (sobrescribe el valor del JSON)
- `-output`: Directorio de salida (sobrescribe el valor del JSON)

### Compilar el programa:

```bash
go build -o generador-certificados.exe
```

Luego puedes ejecutarlo directamente:

```bash
.\generador-certificados.exe
```

## Ejemplos

### Generar certificados con imagen de fondo:

1. Coloca tu imagen de fondo como `background.png` en el directorio del programa
2. Edita `certificados.json` con los datos de tu evento
3. Ejecuta:

```bash
go run main.go
```

### Generar sin imagen de fondo:

Si no proporcionas una imagen de fondo o la ruta no existe, se usará un fondo de color sólido.

```bash
go run main.go -bg ""
```

### Generar con múltiples eventos:

Puedes crear diferentes archivos JSON para diferentes eventos:

```bash
go run main.go -json evento1.json -output evento1_certs
go run main.go -json evento2.json -output evento2_certs
```

## Formato de salida

Los certificados se generan como archivos PDF en formato A4 horizontal (landscape) con el nombre:

```
certificado_<nombre_del_participante>.pdf
```

Los espacios y caracteres especiales en los nombres se reemplazan por guiones bajos.

## Personalización

### Colores:

Puedes cambiar los colores del texto y títulos editando los valores RGB en el JSON:

```json
"font_color": {
  "r": 50,
  "g": 50,
  "b": 50
},
"title_color": {
  "r": 0,
  "g": 51,
  "b": 102
}
```

### Diseño:

Para modificar el diseño del certificado (tamaños de fuente, posiciones, etc.), edita la función `generateCertificate()` en `main.go`.

## Solución de problemas

### Error: "no such file or directory"
- Verifica que las rutas a las imágenes (background_image, signature_image) sean correctas
- Usa rutas relativas desde el directorio donde ejecutas el programa

### Error: "could not import github.com/jung-kurt/gofpdf"
- Ejecuta: `go mod tidy`

### Los certificados salen en blanco
- Verifica que la imagen de fondo no esté corrupta
- Verifica que los colores no sean blancos (255, 255, 255) sobre fondo blanco

## Licencia

Este programa es de código abierto y puede ser modificado según tus necesidades.

## Autor

Creado para facilitar la generación masiva de certificados en eventos académicos y profesionales.
