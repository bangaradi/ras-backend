package student

import "github.com/gin-gonic/gin"

func FirstOrCreateStudent(ctx *gin.Context, student *Student) error {
	tx := db.WithContext(ctx).Where("iitk_email = ?", student.IITKEmail).FirstOrCreate(student)
	return tx.Error
}

func getStudentByID(ctx *gin.Context, student *Student, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(student)
	return tx.Error
}

func getStudentByEmail(ctx *gin.Context, student *Student, email string) error {
	tx := db.WithContext(ctx).Where("iitk_email =?", email).First(student)
	return tx.Error
}

func FetchStudents(ctx *gin.Context, students *[]Student, emails []string) error {
	tx := db.WithContext(ctx).Where("iitk_email IN ?", emails).Find(students)
	return tx.Error
}

func getAllStudents(ctx *gin.Context, students *[]Student) error {
	tx := db.WithContext(ctx).Find(students)
	return tx.Error
}

func updateStudentByID(ctx *gin.Context, student *Student) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", student.ID).Updates(student)
	return tx.RowsAffected > 0, tx.Error
}

func updateStudentByEmail(ctx *gin.Context, student *Student, email string) (bool, error) {
	tx := db.WithContext(ctx).Model(&Student{}).Where("iitk_email = ? AND is_editable = ?", email, false).Updates(student)
	return tx.RowsAffected > 0, tx.Error
}

func deleteStudent(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Student{})
	return tx.Error
}
