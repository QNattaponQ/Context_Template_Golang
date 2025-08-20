# Examples

ตัวอย่างโค้ดที่ใช้เป็น "แพทเทิร์น" ให้ผู้ช่วยโค้ด (AI) เลียนแบบเวลาสร้างฟีเจอร์ใหม่

## ไฟล์

- **server.go**
  - ตัวอย่างการสร้าง HTTP server ด้วย [chi router](https://github.com/go-chi/chi)
  - มี health check endpoint (`/health`)

- **repository.go**
  - ตัวอย่าง repository pattern สำหรับ `User`
  - ใช้ [GORM](https://gorm.io/) + `context.Context`
  - แสดงวิธีแยก interface + struct implementation

- **tests/test_example.go**
  - ตัวอย่าง unit test สไตล์ table-driven test
  - ใช้ SQLite in-memory DB + [testify](https://github.com/stretchr/testify)
  - ครอบคลุม use case: create user + find user by ID