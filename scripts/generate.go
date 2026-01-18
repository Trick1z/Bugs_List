package main

import (
	"bug_list/config"
	"bug_list/internal/database"
	"log"

	"gorm.io/gen"
)

func main() {
	// โหลด config
	cfg := config.Load()

	// เชื่อมต่อ database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	db := database.GetDB()

	// Configure generator
	g := gen.NewGenerator(gen.Config{
		OutPath:        "./internal/models/query",
		Mode:           gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:  true,
		FieldCoverable: false,
	})

	g.UseDB(db)

	// Generate models จาก ALL tables
	g.ApplyBasic(g.GenerateAllTable()...)

	// หรือ generate เฉพาะบาง tables
	// g.ApplyBasic(
	//     g.GenerateModel("users"),
	//     g.GenerateModel("products"),
	//     g.GenerateModel("orders"),
	// )

	// Execute generation
	g.Execute()

	log.Println("🎉 Models generated successfully in ./internal/models/query/")
}
