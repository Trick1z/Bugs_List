package main

import (
	"bug_list/config"
	"bug_list/internal/database"
	"bug_list/internal/models/query"
	"log"
)

func main() {
	// โหลด configuration
	cfg := config.Load()

	// เชื่อมต่อ database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	// Initialize query
	query.SetDefault(database.GetDB())
	q := query.Use(database.GetDB())

	// ===== ทดสอบดึงข้อมูล =====

	// // 1. ดึง Users ทั้งหมด
	// users, err := q.User.Find()
	// if err != nil {
	// 	log.Println("❌ Error fetching users:", err)
	// } else {
	// 	log.Printf("✅ Found %d users:", len(users))
	// 	for i, user := range users {
	// 		log.Printf("  [%d] User: %+v", i+1, user)
	// 	}
	// }

	// // 2. ดึง User คนแรก
	// firstUser, err := q.User.First()
	// if err != nil {
	// 	log.Println("❌ Error fetching first user:", err)
	// } else {
	// 	log.Printf("✅ First User: %+v", firstUser)
	// }

	// 3. ดึง User ตาม ID (เปลี่ยน 1 เป็น ID ที่มีจริง)
	userByID, err := q.User.Where(q.User.UserID.Eq(1)).First()
	if err != nil {
		log.Println("❌ Error fetching user by ID:", err)
	} else {
		log.Printf("%v", userByID)
		// Output: {1 John john@example.com}

		log.Printf("%+v", userByID)
		// Output: {ID:1 userByID:John Email:john@example.com}  ← เห็นชื่อ field ด้วย

		log.Printf("%#v", userByID)
		// Output: main.User{ID:1, Name:"John", Email:"john@example.com"}
	}

	log.Println("🎉 Test completed!")
}
