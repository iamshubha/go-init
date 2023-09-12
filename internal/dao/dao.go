package dao

type DBImpl interface {
}

type daoService struct {
	// db *gorm.DB
}

func NewDB(
// db *gorm.DB
) *daoService {
	return &daoService{
		// db: db,
	}
}

func NewSqlDB() *daoService {
	// var db *gorm.DB
	// func() { db = initDB() }()
	// db.AutoMigrate(&daoService{})
	return &daoService{
		// db: db,
	}
}

// func initDB() *gorm.DB {

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {

// 		panic(err)
// 	}

// 	return db
// }
