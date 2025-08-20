## FEATURE
เพิ่ม REST API endpoint `/users` สำหรับ:
- POST `/users` เพื่อสร้าง user ใหม่ (รับ JSON: name, email, age)
- GET `/users/{id}` เพื่อดึงข้อมูล user ตาม ID

ข้อมูลต้องเก็บลง PostgreSQL โดยใช้ GORM เป็น ORM

## EXAMPLES
อ้างอิง patterns จากโค้ดตัวอย่าง:
- `examples/server.go` → โครงสร้างการตั้งค่า HTTP server ด้วย chi router
- `examples/repository.go` → รูปแบบ repository สำหรับติดต่อฐานข้อมูล
- `examples/tests/test_example.go` → สไตล์การเขียน unit test และ table-driven test

## DOCUMENTATION
- [GORM Documentation](https://gorm.io/docs/)
- [Chi Router](https://github.com/go-chi/chi)
- Database schema:
  ```sql
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
  );