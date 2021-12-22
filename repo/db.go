package repo

import (
	"fmt"

	"github.com/epa-datos/first-api-training/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

var tables = []interface{}{
	&entity.Post{},
}

func Connect() {
	//dsn := fmt.Sprintf("root:pass@tcp(%s:%s)/ejemplo?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQLDB_HOST"), os.Getenv("MYSQLDB_PORT"))
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate()
	fmt.Println("Ya se conecto a la db")
}

func migrate() {
	for _, t := range tables {
		db.AutoMigrate(t)
	}
}

func NewPost(post *entity.Post) error {
	return db.Create(post).Error
}

func GetPost(id int) (*entity.Post, error) {
	post := &entity.Post{}
	return post, db.First(post, id).Error
}

func GetPosts() ([]entity.Post, error) {
	posts := []entity.Post{}
	return posts, db.Find(&posts).Error
}
