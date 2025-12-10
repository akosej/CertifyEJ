# Plantilla de JSON para Certificados

Este es un archivo de plantilla que puedes copiar y modificar según tus necesidades.

## Plantilla completa con todos los campos

```json
{
  "background_image": "./bg.jpg",
  "output_directory": "certificados_generados",
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
    "name": "Nombre Completo del Evento",
    "certificate_title": "CERTIFICADO DE PARTICIPACIÓN",
    "location": "Ciudad, País",
    "start_date": "DD de mes de AAAA",
    "end_date": "DD de mes de AAAA",
    "isbn": "ISBN-XXX-XXX-XXXX-X",
    "organizers": "Lista de organizadores",
    "signature_name": "Nombre Completo del Firmante",
    "signature_role": "Cargo del Firmante",
    "signature_image": "ruta/a/firma.png",
    "signature_width": 40
  },
  "participants": [
    {
      "name": "Nombre Completo del Participante",
      "affiliation": "Institución/Universidad",
      "participation_type": "ponente",
      "work_title": "Título del trabajo presentado"
    }
  ]
}
```

## Campos opcionales

Puedes dejar estos campos vacíos (`""`) si no los necesitas:

- `background_image`: Se usará un fondo de color sólido
- `signature_image`: No se mostrará imagen de firma
- `isbn`: No se mostrará ISBN
- `organizers`: No se mostrarán organizadores
- `affiliation`: No se mostrará afiliación
- `work_title`: No se mostrará título del trabajo
- `participation_type`: Se usará "participante" por defecto

## Tipos de certificado comunes

### 1. Certificado de Ponencia

```json
{
  "certificate_title": "CERTIFICADO DE PONENCIA",
  "participation_type": "ponente",
  "work_title": "Título de la ponencia"
}
```

### 2. Certificado de Asistencia

```json
{
  "certificate_title": "CERTIFICADO DE ASISTENCIA",
  "participation_type": "asistente",
  "work_title": ""
}
```

### 3. Certificado de Autoría

```json
{
  "certificate_title": "CERTIFICADO DE AUTOR",
  "participation_type": "autor",
  "work_title": "Título del artículo o trabajo"
}
```

### 4. Certificado de Moderación

```json
{
  "certificate_title": "CERTIFICADO DE MODERACIÓN",
  "participation_type": "moderador",
  "work_title": "Nombre de la mesa o panel"
}
```

## Colores comunes para títulos

### Azul corporativo
```json
"title_color": { "r": 0, "g": 51, "b": 102 }
```

### Rojo elegante
```json
"title_color": { "r": 139, "g": 0, "b": 0 }
```

### Verde oscuro
```json
"title_color": { "r": 0, "g": 100, "b": 0 }
```

### Dorado
```json
"title_color": { "r": 184, "g": 134, "b": 11 }
```

### Negro
```json
"title_color": { "r": 0, "g": 0, "b": 0 }
```

## Importación masiva desde CSV

Si tienes muchos participantes, puedes convertir un CSV a JSON usando herramientas online o scripts de Python.

Ejemplo de CSV:
```csv
nombre,afiliacion,tipo,trabajo
Juan Pérez,Universidad A,ponente,Título del trabajo 1
María González,Universidad B,autor,Título del trabajo 2
```

## Notas importantes

1. Las fechas pueden tener cualquier formato de texto
2. Los nombres de archivo se generan automáticamente reemplazando espacios por guiones bajos
3. El directorio de salida se crea automáticamente si no existe
4. Las rutas de imágenes pueden ser relativas o absolutas
5. Si una imagen no se encuentra, el programa continuará sin ella
