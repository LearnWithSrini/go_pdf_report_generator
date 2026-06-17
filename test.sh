#!/bin/bash

# Test script for Go PDF Report Generator

echo "🧪 Testing Go PDF Report Generator Service"
echo "=========================================="
echo ""

# Check if service is running
echo "📡 Checking if service is running..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Service is running"
else
    echo "❌ Service is not running. Please start it with: go run main.go"
    exit 1
fi

echo ""
echo "📄 Generating PDF report for student ID 1..."
curl -s http://localhost:8080/api/v1/students/1/report -o student_1_report.pdf

if [ -f "student_1_report.pdf" ]; then
    SIZE=$(stat -f%z student_1_report.pdf 2>/dev/null || stat -c%s student_1_report.pdf 2>/dev/null)
    if [ "$SIZE" -gt 1000 ]; then
        echo "✅ PDF generated successfully! (Size: $SIZE bytes)"
        echo "📁 Report saved as: student_1_report.pdf"
        echo ""
        echo "🎉 Test completed successfully!"
    else
        echo "❌ PDF file is too small or empty"
        cat student_1_report.pdf
        exit 1
    fi
else
    echo "❌ Failed to generate PDF"
    exit 1
fi
