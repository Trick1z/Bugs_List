package services

import (
	"bug_list/internal/models/query" // ✅ ใช้ / ถูกต้อง
)

type UserService struct {
	query *query.Query
}

func NewUserService(q *query.Query) *UserService {
	return &UserService{query: q}
}

// ตัวอย่าง: Get all users
func (s *UserService) GetAllUsers() (interface{}, error) {
	users, err := s.query.User.Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// ตัวอย่าง: Get user by ID
func (s *UserService) GetUserByID(id int) (interface{}, error) {
	user, err := s.query.User.Where(s.query.User.UserID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// เพิ่ม methods อื่นๆ ตามต้องการ...
