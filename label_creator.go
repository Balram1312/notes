package main

import (
	"fmt"
	"log"
	"time"

	"github.com/phpdave11/gofpdf"
)

// generateShippingLabel generates the shipping label in PDF format as per the specifications
func generateShippingLabel() {
	startTime := time.Now()

	// Create a new PDF document with dimensions 10x10 cm
	pdf := gofpdf.New("P", "cm", "A4", "")
	pdf.SetMargins(0, 0, 0)
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 10, Ht: 10})

	// Set font for the text (scaled down 10px, from 6px to 5px)
	pdf.SetFont("Arial", "", 5)

	// 1. Add Contract ID (scaled down)
	pdf.CellFormat(0, 0.7, "Contract ID: 369613", "", 1, "L", false, 0, "")

	// Logo at 18x18 px (scaled down from 20x20 px)
	pdf.Image("logo.png", 6.1, pdf.GetY()-0.45, 1.8, 1.8, false, "", 0, "")

	// Move to next line after logo
	pdf.Ln(0)

	// 3. Add "SENDER ADDRESS" (scaled down)
	pdf.MultiCell(6.1, 0.3, "SENDER ADDRESS\nBalram\nViale Italia 537,\n19125 La Spezia (SP)", "LBT", "L", false)

	// 6. Add a table with two columns for QR Code and Barcode (scaled down)
	pdf.SetDrawColor(0, 0, 0) // Set color for table borders

	// Define the overall table width for QR Code and Barcode table
	pageTableWidth := 8.8 // This will define the total width of the table

	// Define proportional column widths for QR Code and Barcode table
	qrColWidth := [2]float64{
		pageTableWidth * 0.5, // 50% of the total table width for QR Code column
		pageTableWidth * 0.5, // 50% for Barcode column
	}

	// Table headers for QR Code and Barcode with borders (scaled down)
	pdf.CellFormat(qrColWidth[0], 1.8, "QR CODE", "1", 0, "C", false, 0, "") // Left cell for QR Code
	pdf.CellFormat(qrColWidth[1], 1.8, "BARCODE", "1", 1, "C", false, 0, "") // Right cell for Barcode
	pdf.Ln(0)

	// 7. Add Delivery Address section
	// Delivery address will use the same pageTableWidth proportionally as well
	pageTableWidthDelivery := pageTableWidth // Can adjust independently for the address if needed

	// Delivery Address content
	pdf.MultiCell(pageTableWidthDelivery, 0.35, "DELIVERY ADDRESS\n\nPiyush Sonawane, Via Giotto Ciardi\n18,57121 Livorno (LI)", "1", "L", false)
	pdf.Ln(0)

	// Define column widths for the data table (the main data table, not QR or Barcode table)
	// Proportional widths of the columns
	colWidth := [5]float64{
		pageTableWidth * 0.2, // 20% for column 1
		pageTableWidth * 0.2, // 20% for column 2
		pageTableWidth * 0.2, // 20% for column 3
		pageTableWidth * 0.2, // 20% for column 4
		pageTableWidth * 0.2, // 20% for column 5
	}

	lnBreakHeight := 0.7

	// Row 1 with a "rowspan" effect on the first column
	// First column: "Lv02" with a larger height (scaled down)
	pdf.CellFormat(colWidth[0], 1.4, "Lv02", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[1], 0.7, "Row 1 Col 2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[2], 0.7, "Row 1 Col 3", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[3], 0.7, "Row 1 Col 4", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[4], 0.7, "Row 1 Col 5", "1", 0, "L", false, 0, "")
	pdf.Ln(lnBreakHeight)

	// Row 2 with a "rowspan" effect on the first column (skip first column)
	pdf.CellFormat(colWidth[0], 0.7, "", "2", 0, "L", false, 0, "") // Empty cell for first column (simulating rowspan)
	pdf.CellFormat(colWidth[1], 0.7, "Row 2 Col 2", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[2], 0.7, "Row 2 Col 3", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[3], 0.7, "Row 2 Col 4", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidth[4], 0.7, "Row 2 Col 5", "1", 0, "L", false, 0, "")
	pdf.Ln(lnBreakHeight)

	// Bottom row with a "colspan" effect (scaled down)
	pdf.CellFormat(pageTableWidth, 0.7, "Bottom Row Colspan", "1", 0, "L", false, 0, "") // Colspan across all columns
	pdf.Ln(lnBreakHeight)

	// Save the PDF file
	err := pdf.OutputFileAndClose("shipping_label.pdf")
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	fmt.Println("Shipping label generated successfully!")
	log.Println("time taken to generate label: ", time.Since(startTime))
}

func main() {
	// Generate the shipping label
	generateShippingLabel()
}
