package orm

var (
	dbHOST     = "127.0.0.1"
	dbUSER     = "root"
	dbPORT     = "3306"
	dbNAME     = "great_db"
	dbPASSWORD = "123"
	DbUrl      = dbUSER + ":" + dbPASSWORD + "@tcp(" + dbHOST + ":" + dbPORT + ")/" + dbNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
)
