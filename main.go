package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg" // For JPEG support in image.Decode
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Datos del evento
type Event struct {
	Name             string  `json:"name"`
	Location         string  `json:"location"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	ISBN             string  `json:"isbn"`       // Corregido de ICBN
	Organizers       string  `json:"organizers"` // Organizadores del evento
	SignatureName    string  `json:"signature_name"`
	SignatureRole    string  `json:"signature_role"`
	SignatureImage   string  `json:"signature_image"`   // Opcional: imagen de firma
	SignatureWidth   float64 `json:"signature_width"`   // Ancho de la firma en mm (por defecto 40mm)
	CertificateTitle string  `json:"certificate_title"` // Título del certificado
	LogoImage        string  `json:"logo_image"`        // Opcional: imagen de logo
	LogoWidth        float64 `json:"logo_width"`        // Ancho del logo en mm (por defecto 30mm)
	LogoX            float64 `json:"logo_x"`            // Posición X del logo (por defecto 10mm)
	LogoY            float64 `json:"logo_y"`            // Posición Y del logo (por defecto 10mm)
}

// Participantes
type Participant struct {
	Name               string `json:"name"`
	WorkTitle          string `json:"work_title"`
	ParticipationType  string `json:"participation_type"` // Ej: "Ponente", "Asistente", "Autor"
	AffiliationAddress string `json:"affiliation"`        // Institución o afiliación
}

// Estructura general
type CertificateData struct {
	Event           Event         `json:"event"`
	Participants    []Participant `json:"participants"`
	BackgroundImage string        `json:"background_image"` // Ruta al fondo
	OutputDirectory string        `json:"output_directory"` // Directorio de salida
	FontColor       RGB           `json:"font_color"`       // Color del texto
	TitleColor      RGB           `json:"title_color"`      // Color del título
}

// RGB para colores personalizados
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// ---------------------------------------------

// convertPNGToNonInterlaced reads a PNG file and converts it to non-interlaced format
// Returns a temporary file path that should be cleaned up by the caller
func convertPNGToNonInterlaced(pngPath string) (string, bool, error) {
	// Read the original file
	file, err := os.Open(pngPath)
	if err != nil {
		return "", false, err
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return "", false, err
	}

	// If it's not a PNG, return the original path
	if format != "png" {
		return pngPath, false, nil
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "cert_*.png")
	if err != nil {
		return "", false, err
	}
	tmpPath := tmpFile.Name()

	// Encode as non-interlaced PNG
	encoder := png.Encoder{
		CompressionLevel: png.DefaultCompression,
	}
	err = encoder.Encode(tmpFile, img)
	tmpFile.Close()

	if err != nil {
		os.Remove(tmpPath)
		return "", false, err
	}

	return tmpPath, true, nil
}

// safeImagePath converts PNG images to non-interlaced format if needed
// Returns the path to use (original or temporary) and whether cleanup is needed
func safeImagePath(imagePath string) (string, bool, error) {
	ext := strings.ToLower(filepath.Ext(imagePath))
	if ext == ".png" {
		return convertPNGToNonInterlaced(imagePath)
	}
	return imagePath, false, nil
}

func generateCertificate(bgImage string, event Event, p Participant, fontColor, titleColor RGB, outputDir string) error {
	pdf := gofpdf.New("L", "mm", "A4", "")

	// Configurar codificación UTF-8 para caracteres especiales (tildes, ñ, etc.)
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.AddPage()

	// Track temporary files for cleanup
	var tempFiles []string
	defer func() {
		for _, tmpFile := range tempFiles {
			os.Remove(tmpFile)
		}
	}()

	// Verificar si existe la imagen de fondo
	if _, err := os.Stat(bgImage); err == nil {
		// Convert PNG to non-interlaced if needed
		safePath, isTemp, err := safeImagePath(bgImage)
		if err != nil {
			return fmt.Errorf("error procesando imagen de fondo: %v", err)
		}
		if isTemp {
			tempFiles = append(tempFiles, safePath)
		}
		// Fondo completo
		pdf.Image(safePath, 0, 0, 297, 210, false, "", 0, "")
	} else {
		// Si no hay fondo, usar color de fondo predeterminado
		pdf.SetFillColor(245, 245, 250)
		pdf.Rect(0, 0, 297, 210, "F")
	}

	// Logo (si existe)
	if event.LogoImage != "" {
		if _, err := os.Stat(event.LogoImage); err == nil {
			logoWidth := event.LogoWidth
			if logoWidth <= 0 {
				logoWidth = 30 // Valor por defecto
			}
			// Convert PNG to non-interlaced if needed
			safePath, isTemp, err := safeImagePath(event.LogoImage)
			if err != nil {
				return fmt.Errorf("error procesando logo: %v", err)
			}
			if isTemp {
				tempFiles = append(tempFiles, safePath)
			}
			pdf.Image(safePath, event.LogoX, event.LogoY, logoWidth, 0, false, "", 0, "")
		}
	}

	// Título del certificado
	pdf.SetY(25)
	pdf.SetFont("Arial", "B", 32)
	pdf.SetTextColor(titleColor.R, titleColor.G, titleColor.B)
	certificateTitle := event.CertificateTitle
	if certificateTitle == "" {
		certificateTitle = "CERTIFICADO DE PARTICIPACIÓN"
	}
	pdf.CellFormat(0, 15, tr(certificateTitle), "", 1, "C", false, 0, "")

	// Texto introductorio
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 13)
	pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)
	pdf.CellFormat(0, 8, tr("Se otorga el presente certificado a:"), "", 1, "C", false, 0, "")

	// Nombre del participante
	pdf.Ln(3)
	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(titleColor.R, titleColor.G, titleColor.B)
	pdf.CellFormat(0, 12, tr(p.Name), "", 1, "C", false, 0, "")

	// Afiliación (si existe)
	if p.AffiliationAddress != "" {
		pdf.Ln(1)
		pdf.SetFont("Arial", "I", 11)
		pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)
		pdf.CellFormat(0, 6, tr(p.AffiliationAddress), "", 1, "C", false, 0, "")
	}

	// Tipo de participación y trabajo
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 13)
	pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)

	participationType := p.ParticipationType
	if participationType == "" {
		participationType = "participante"
	}

	text := fmt.Sprintf("Por su participación como %s en el evento", participationType)
	pdf.CellFormat(0, 7, tr(text), "", 1, "C", false, 0, "")

	// Nombre del evento
	pdf.Ln(1)
	pdf.SetFont("Arial", "B", 15)
	pdf.SetTextColor(titleColor.R, titleColor.G, titleColor.B)
	pdf.MultiCell(0, 7, tr(fmt.Sprintf("\"%s\"", event.Name)), "", "C", false)

	// Título del trabajo (si existe)
	if p.WorkTitle != "" {
		pdf.Ln(2)
		pdf.SetFont("Arial", "I", 12)
		pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)
		pdf.CellFormat(0, 6, tr("con el trabajo:"), "", 1, "C", false, 0, "")
		pdf.Ln(0.5)
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, 6, tr(fmt.Sprintf("\"%s\"", p.WorkTitle)), "", "C", false)
	}

	// Información del evento (fechas, ubicación, ISBN)
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)

	eventInfo := fmt.Sprintf("Celebrado del %s al %s en %s",
		event.StartDate, event.EndDate, event.Location)
	pdf.CellFormat(0, 5, tr(eventInfo), "", 1, "C", false, 0, "")

	// ISBN si existe
	if event.ISBN != "" {
		pdf.Ln(1)
		pdf.CellFormat(0, 5, tr(fmt.Sprintf("ISBN: %s", event.ISBN)), "", 1, "C", false, 0, "")
	}

	// Organizadores si existen
	if event.Organizers != "" {
		pdf.Ln(1)
		pdf.SetFont("Arial", "I", 9)
		pdf.CellFormat(0, 5, tr(fmt.Sprintf("Organizado por: %s", event.Organizers)), "", 1, "C", false, 0, "")

	}
	// Firma
	pdf.Ln(8)

	// Si hay imagen de firma, colocarla primero (arriba del nombre)
	if event.SignatureImage != "" {
		if _, err := os.Stat(event.SignatureImage); err == nil {
			// Ancho de firma configurable (por defecto 40mm)
			signatureWidth := event.SignatureWidth
			if signatureWidth <= 0 {
				signatureWidth = 40 // Valor por defecto
			}
			// Convert PNG to non-interlaced if needed
			safePath, isTemp, err := safeImagePath(event.SignatureImage)
			if err != nil {
				return fmt.Errorf("error procesando firma: %v", err)
			}
			if isTemp {
				tempFiles = append(tempFiles, safePath)
			}
			// Centrar la imagen de firma
			xPos := (297 - signatureWidth) / 2 // Centrar en A4 horizontal (297mm)
			pdf.Image(safePath, xPos, pdf.GetY(), signatureWidth, 0, false, "", 0, "")
			pdf.Ln(signatureWidth*0.4 + 5) // Espacio suficiente para que la imagen no se sobreponга
		}
	}

	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(fontColor.R, fontColor.G, fontColor.B)
	pdf.CellFormat(0, 5, tr(event.SignatureName), "", 1, "C", false, 0, "")

	if event.SignatureRole != "" {
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 5, tr(event.SignatureRole), "", 1, "C", false, 0, "")
	}

	// Crear directorio de salida si no existe
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creando directorio de salida: %v", err)
		}
	}

	// Guardar archivo con nombre sanitizado
	safeName := sanitizeFilename(p.Name)
	filename := filepath.Join(outputDir, fmt.Sprintf("certificado_%s.pdf", safeName))
	return pdf.OutputFileAndClose(filename)
}

// Función para sanitizar nombres de archivo
func sanitizeFilename(name string) string {
	// Reemplazar caracteres no válidos en nombres de archivo
	replacer := strings.NewReplacer(
		" ", "_",
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	return replacer.Replace(name)
}

// ---------------------------------------------

func main() {
	// Flags para línea de comandos
	jsonFileFlag := flag.String("json", "certificados.json", "Archivo JSON con los datos")
	bgImageFlag := flag.String("bg", "", "Imagen de fondo (sobrescribe el JSON)")
	outputDirFlag := flag.String("output", "", "Directorio de salida (sobrescribe el JSON)")
	flag.Parse()

	// Archivo JSON
	jsonFile, err := os.ReadFile(*jsonFileFlag)
	if err != nil {
		log.Fatalf("Error leyendo JSON '%s': %v", *jsonFileFlag, err)
	}

	var data CertificateData
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Fatalf("Error parseando JSON: %v", err)
	}

	// Valores predeterminados
	if data.FontColor.R == 0 && data.FontColor.G == 0 && data.FontColor.B == 0 {
		data.FontColor = RGB{R: 50, G: 50, B: 50} // Gris oscuro
	}
	if data.TitleColor.R == 0 && data.TitleColor.G == 0 && data.TitleColor.B == 0 {
		data.TitleColor = RGB{R: 0, G: 51, B: 102} // Azul oscuro
	}

	// Sobrescribir con flags si se proporcionan
	bgImage := data.BackgroundImage
	if *bgImageFlag != "" {
		bgImage = *bgImageFlag
	}
	if bgImage == "" {
		bgImage = "background.png" // Valor por defecto
	}

	outputDir := data.OutputDirectory
	if *outputDirFlag != "" {
		outputDir = *outputDirFlag
	}
	if outputDir == "" {
		outputDir = "certificados" // Valor por defecto
	}

	fmt.Printf("Generando certificados para %d participantes...\n", len(data.Participants))
	fmt.Printf("Imagen de fondo: %s\n", bgImage)
	fmt.Printf("Directorio de salida: %s\n\n", outputDir)

	successCount := 0
	errorCount := 0

	for i, participant := range data.Participants {
		fmt.Printf("[%d/%d] Generando certificado para: %s... ",
			i+1, len(data.Participants), participant.Name)

		err := generateCertificate(bgImage, data.Event, participant,
			data.FontColor, data.TitleColor, outputDir)
		if err != nil {
			fmt.Printf("❌ ERROR\n")
			log.Printf("   Error: %v\n", err)
			errorCount++
		} else {
			fmt.Printf("✔ OK\n")
			successCount++
		}
	}

	fmt.Printf("\n========================================\n")
	fmt.Printf("Proceso completado.\n")
	fmt.Printf("✔ Exitosos: %d\n", successCount)
	if errorCount > 0 {
		fmt.Printf("❌ Errores: %d\n", errorCount)
	}
	fmt.Printf("========================================\n")
}
