// Simple mock backend server to test the Go PDF service
const http = require('http');

const students = {
  '2': {
    success: true,
    data: {
      id: 2,
      name: 'John Smith',
      email: 'john.smith@student.com',
      phone: '+1-234-567-8901',
      dob: '2005-03-15',
      father_name: 'Robert Smith',
      father_phone: '+1-234-567-8900',
      mother_name: 'Mary Smith',
      mother_phone: '+1-234-567-8902',
      address: '123 Main St, Springfield, IL',
      class_name: 'Grade 10',
      section_name: 'A',
      roll: 15,
      status: 'active',
      created_at: '2024-12-01T10:00:00Z',
      updated_at: '2025-06-17T10:00:00Z'
    }
  },
  '3': {
    success: true,
    data: {
      id: 3,
      name: 'Jane Doe',
      email: 'jane.doe@student.com',
      phone: '+1-234-567-8911',
      dob: '2006-07-22',
      father_name: 'Michael Doe',
      father_phone: '+1-234-567-8910',
      mother_name: 'Sarah Doe',
      mother_phone: '+1-234-567-8912',
      address: '456 Oak Ave, Springfield, IL',
      class_name: 'Grade 9',
      section_name: 'B',
      roll: 23,
      status: 'active',
      created_at: '2024-10-01T10:00:00Z',
      updated_at: '2025-06-17T10:00:00Z'
    }
  },
  '4': {
    success: true,
    data: {
      id: 4,
      name: 'Alex Johnson',
      email: 'alex.johnson@student.com',
      phone: '+1-234-567-8921',
      dob: '2004-11-30',
      father_name: 'David Johnson',
      father_phone: '+1-234-567-8920',
      mother_name: 'Lisa Johnson',
      mother_phone: '+1-234-567-8922',
      address: '789 Pine Rd, Springfield, IL',
      class_name: 'Grade 11',
      section_name: 'A',
      roll: 8,
      status: 'active',
      created_at: '2024-08-01T10:00:00Z',
      updated_at: '2025-06-17T10:00:00Z'
    }
  }
};

const server = http.createServer((req, res) => {
  // Enable CORS
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Content-Type', 'application/json');

  // Parse URL
  const url = new URL(req.url, `http://${req.headers.host}`);

  // Health check
  if (url.pathname === '/health') {
    res.writeHead(200);
    res.end('OK');
    return;
  }

  // Match /api/v1/students/:id
  const match = url.pathname.match(/^\/api\/v1\/students\/(\d+)$/);

  if (match && req.method === 'GET') {
    const studentId = match[1];

    if (students[studentId]) {
      console.log(`✅ Serving student ${studentId}: ${students[studentId].data.name}`);
      res.writeHead(200);
      res.end(JSON.stringify(students[studentId]));
    } else {
      console.log(`❌ Student ${studentId} not found`);
      res.writeHead(404);
      res.end(JSON.stringify({ success: false, message: 'Student not found' }));
    }
  } else {
    res.writeHead(404);
    res.end(JSON.stringify({ success: false, message: 'Not found' }));
  }
});

const PORT = 5007;
server.listen(PORT, () => {
  console.log(`🚀 Mock Backend Server running on http://localhost:${PORT}`);
  console.log(`📚 Available students: 2, 3, 4`);
  console.log(`🔗 Example: http://localhost:${PORT}/api/v1/students/2`);
});
