# Testing Guide

## ✅ Completed Tasks

### 1. Dependencies Installed
- ✅ Go dependencies installed (`go mod download`)
- ✅ Node.js backend dependencies installed (`npm install`)
- ✅ Compiled Go binary created (`go-pdf-service`)

### 2. Demo PDFs Generated
Successfully generated **3 sample PDF reports** without requiring database:

| File | Size | Status |
|------|------|--------|
| `demo_student_1_report.pdf` | 2.4 KB | ✅ Generated |
| `demo_student_2_report.pdf` | 2.4 KB | ✅ Generated |
| `demo_student_3_report.pdf` | 2.4 KB | ✅ Generated |

These PDFs demonstrate the complete PDF generation functionality with:
- Professional header and formatting
- Student personal information
- Academic details (class, section, roll number)
- Parent/guardian information
- Address details
- Color-coded sections
- Timestamp of generation

## 🎯 How to View the Generated PDFs

Open the PDFs with:

```bash
# On macOS
open demo_student_1_report.pdf
open demo_student_2_report.pdf
open demo_student_3_report.pdf

# Or use Preview
open -a Preview demo_student_*.pdf
```

## 🔧 Full Integration Testing (with PostgreSQL)

Once PostgreSQL is installed and configured, follow these steps:

### Step 1: Install PostgreSQL (if not installed)

**macOS:**
```bash
brew install postgresql@14
brew services start postgresql@14
```

**Alternative - Download from:**
https://www.postgresql.org/download/

### Step 2: Set Up Database

```bash
# Create database
createdb school_mgmt

# Run schema
cd /Users/srinivasulureddymunagala/teamreach-clone/diemission
psql -d school_mgmt -f seed_db/tables.sql
psql -d school_mgmt -f seed_db/seed-db.sql
```

### Step 3: Start Backend Service

```bash
# Terminal 1: Start Node.js backend
cd /Users/srinivasulureddymunagala/teamreach-clone/diemission/backend
npm start
```

The backend should start on `http://localhost:5007`

### Step 4: Start Go Service

```bash
# Terminal 2: Start Go PDF service
cd /Users/srinivasulureddymunagala/teamreach-clone/go_pdf_report_generator
go run main.go

# Or use the compiled binary
./go-pdf-service
```

The Go service should start on `http://localhost:8080`

### Step 5: Generate Real PDFs from Backend Data

```bash
# Terminal 3: Generate reports
# This will fetch real data from the backend API and generate PDFs

# Test with different student IDs
curl http://localhost:8080/api/v1/students/1/report -o student_1.pdf
curl http://localhost:8080/api/v1/students/2/report -o student_2.pdf
curl http://localhost:8080/api/v1/students/3/report -o student_3.pdf

# Run the test script
./test.sh
```

## 📊 Testing Results Summary

### What's Working ✅

1. **Go Service Compilation**: Binary builds successfully
2. **PDF Generation Logic**: Generates valid PDF documents
3. **Data Models**: Student data structure properly defined
4. **Service Architecture**: Clean separation of concerns
   - Handlers (HTTP layer)
   - Services (Business logic)
   - Models (Data structures)
5. **Demo Mode**: Can generate PDFs with mock data
6. **Error Handling**: Comprehensive error handling implemented

### Next Steps for Full Testing 🔄

1. Install and configure PostgreSQL
2. Run database migrations
3. Start both services (Node.js + Go)
4. Test API integration with real data
5. Verify PDF content matches database records

## 🎨 Sample Student Data in Demo PDFs

### Student 1 - John Doe
- **ID**: 1
- **Class**: Grade 10, Section A
- **Roll**: 15
- **Email**: john.doe@school.com
- **Phone**: +1-234-567-8901
- **Parents**: Robert Doe (Father), Mary Doe (Mother)
- **Address**: 123 Main Street, Springfield, IL 62701

### Student 2 - Jane Smith
- **ID**: 2
- **Class**: Grade 9, Section B
- **Roll**: 23
- **Email**: jane.smith@school.com
- **Phone**: +1-234-567-8911
- **Parents**: Michael Smith (Father), Sarah Smith (Mother)
- **Address**: 456 Oak Avenue, Springfield, IL 62702

### Student 3 - Alex Johnson
- **ID**: 3
- **Class**: Grade 11, Section A
- **Roll**: 8
- **Email**: alex.johnson@school.com
- **Phone**: +1-234-567-8921
- **Parents**: David Johnson (Father), Lisa Johnson (Mother)
- **Address**: 789 Pine Road, Springfield, IL 62703

## 🐛 Troubleshooting

### PostgreSQL Not Installed
**Symptom**: `command not found: createdb` or `command not found: psql`

**Solution**: Install PostgreSQL using Homebrew or download from postgresql.org

### Backend Won't Start
**Symptom**: Backend crashes or won't connect to database

**Solution**:
1. Verify PostgreSQL is running: `pg_isready`
2. Check `.env` file in backend directory
3. Ensure database is created and seeded
4. Check logs for specific errors

### Go Service Connection Error
**Symptom**: `Failed to fetch student data: connection refused`

**Solution**:
1. Ensure Node.js backend is running on port 5007
2. Check `BACKEND_API_URL` in `.env`
3. Verify backend health: `curl http://localhost:5007/api/v1/auth/login`

## 📁 Files Created

```
go_pdf_report_generator/
├── demo_student_1_report.pdf    ← Demo PDF #1
├── demo_student_2_report.pdf    ← Demo PDF #2
├── demo_student_3_report.pdf    ← Demo PDF #3
├── demo.go                      ← Standalone demo
├── test.sh                      ← Integration test script
├── main.go                      ← Main service
├── go-pdf-service              ← Compiled binary
└── [other source files]
```

---

**Status**: Demo completed successfully! ✅
**PDFs Generated**: 3
**Next**: Install PostgreSQL for full integration testing
