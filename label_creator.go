package main

import (
	"fmt"
	"log"

	"github.com/phpdave11/gofpdf"
)

// generateShippingLabel generates the shipping label in PDF format as per the specifications
func generateShippingLabel() {
	// Create a new PDF document with dimensions 10x10 cm
	pdf := gofpdf.New("P", "cm", "A4", "")
	pdf.SetMargins(0, 0, 0)
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 10, Ht: 10})

	// Set font for the text
	pdf.SetFont("Arial", "", 8)

	// 1. Add Contract ID
	pdf.CellFormat(0, 0.8, "Contract ID: 369613", "", 1, "L", false, 0, "")

	// 2. Add Horizontal Line and Company Logo
	pdf.SetDrawColor(0, 0, 0)                                          // black color for line
	pdf.Line(0, pdf.GetY(), 6, pdf.GetY())                             // horizontal line at 60% width
	pdf.Image("logo.png", 6.5, pdf.GetY()-0.5, 2, 2, false, "", 0, "") // Logo at 20x20 px

	// Move to next line after logo and line
	pdf.Ln(1.5)

	// 3. Add "SENDER ADDRESS"
	pdf.CellFormat(0, 0.8, "SENDER ADDRESS", "", 1, "L", false, 0, "")

	// 4. Add sender's name
	pdf.CellFormat(0, 0.8, "Balram", "", 1, "L", false, 0, "")

	// 5. Add sender's address (2 lines)
	pdf.CellFormat(0, 0.8, "Viale Italia 537,", "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 0.8, "19125 La Spezia (SP)", "", 1, "L", false, 0, "")
	pdf.Ln(0.5)

	// 6. Add a table with two columns for QR Code and Barcode
	pdf.SetDrawColor(0, 0, 0) // Set color for table borders

	// Table headers for QR Code and Barcode with borders
	pdf.CellFormat(4.5, 1.5, "QR CODE", "1", 0, "C", false, 0, "") // Left cell for QR Code
	pdf.CellFormat(4.5, 1.5, "BARCODE", "1", 1, "C", false, 0, "") // Right cell for Barcode
	pdf.Ln(0.5)

	// 7. Add Destination Code and Delivery Address
	// First Cell: Destination Code with rotated text
	pdf.Rect(0, pdf.GetY(), 2, 4, "D") // Rectangle for Destination Code
	pdf.SetTextColor(255, 0, 0)        // Red color for rotated text

	// Rotate text for vertical alignment using TransformBegin, TransformRotate, and TransformEnd
	pdf.TransformBegin()
	pdf.TransformRotate(90, 1, pdf.GetY()+2) // Rotate text by 90 degrees around a point
	pdf.Text(1, pdf.GetY()+1, "LV01")        // Positioned for vertical center
	pdf.TransformEnd()

	pdf.SetTextColor(0, 0, 0) // Reset text color to black

	// Second Cell: Delivery Address
	pdf.Rect(3, pdf.GetY(), 7, 4, "D") // Rectangle for Delivery Address
	pdf.SetXY(3.2, pdf.GetY()+0.5)
	pdf.CellFormat(0, 0.8, "DELIVERY ADDRESS", "", 1, "L", false, 0, "")
	pdf.SetX(3.2)
	pdf.CellFormat(0, 0.8, "Piyush Sonawane", "", 1, "L", false, 0, "")
	pdf.SetX(3.2)
	pdf.CellFormat(0, 0.8, "Via Giotto Ciardi 18,", "", 1, "L", false, 0, "")
	pdf.SetX(3.2)
	pdf.CellFormat(0, 0.8, "57121 Livorno (LI)", "", 1, "L", false, 0, "")
	pdf.Ln(1.5)

	// 8. Add 5x3 Table with rowspan in the first column and colspan in the bottom row
	pdf.SetFont("Arial", "", 7)

	// Row 1 with rowspan in the first cell
	pdf.CellFormat(2.2, 1.6, "Row 1 Col 1 (Rowspan)", "1", 0, "L", false, 0, "") // Rowspan for first column
	pdf.CellFormat(3.3, 0.8, "Row 1 Col 2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(3.3, 0.8, "Row 1 Col 3", "1", 1, "L", false, 0, "")

	// Row 2 with rowspan in the first column
	pdf.CellFormat(3.3, 0.8, "Row 2 Col 2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(3.3, 0.8, "Row 2 Col 3", "1", 1, "L", false, 0, "")

	// Bottom row with colspan
	pdf.CellFormat(9.8, 0.8, "Bottom Row Colspan", "1", 1, "L", false, 0, "") // Colspan across all columns

	// Save the PDF file
	err := pdf.OutputFileAndClose("shipping_label.pdf")
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	fmt.Println("Shipping label generated successfully!")
}

func main() {
	// Generate the shipping label
	generateShippingLabel()
}
