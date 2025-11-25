package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
)

func GetStudentsDbHandler(students []models.Student, r *http.Request) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving Data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class FROM students WHERE 1=1 "
	var args []any

	query, args = utils.AddFilters(r, query, args)
	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving Data")
	}
	defer rows.Close()

	//teacherList := make([]models.Teacher, 0)
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			//http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error retrieving Data")
		}
		students = append(students, student)
	}

	return students, nil

}

func AddStudentsDBHandler(newStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		//http.Error(w, "Error connction to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error retrieving Data")

	}
	defer db.Close()

	//stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject)VALUES (?,?,?,?,?)")
	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		//http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error retrieving Data")
	}
	defer stmt.Close()

	addedStudents := make([]models.Student, len(newStudents))
	for i, newStudent := range newStudents {
		//res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)

		values := utils.GetStructValues(newStudent)
		res, err := stmt.Exec(values...)

		if err != nil {
			//http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error updating Data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			//http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error updating Data")
		}
		newStudent.ID = int(lastID)
		addedStudents[i] = newStudent

	}

	return addedStudents, nil
}

func GetStudentByID(id int) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	var teacher models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM student WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class)

	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data")
	} else if err != nil {
		fmt.Println(err)
		//http.Error(w, "Database query error", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "errr retrieving data")
	}
	return teacher, nil
}

func UpdateStudent(id int, updatedStudent models.Student) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)

		return models.Student{}, utils.ErrorHandler(err, "errr Updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(
		&existingStudent.ID,
		&existingStudent.FirstName,
		&existingStudent.LastName,
		&existingStudent.Email,
		&existingStudent.Class,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Student{}, utils.ErrorHandler(err, "errr retrieving data")

		}
		log.Println(err)
		//	http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "errr retrieving data")

	}

	updatedStudent.ID = existingStudent.ID

	// Fixed: underscore for result, fixed column name, removed extra comma
	_, err = db.Exec("UPDATE student SET first_name = ?, last_name = ?, email = ?, class = ?  WHERE id = ?",
		updatedStudent.FirstName,
		updatedStudent.LastName,
		updatedStudent.Email,
		updatedStudent.Class,

		updatedStudent.ID,
	)
	if err != nil {
		log.Println(err)
		//	http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "errr retrieving data")

	}

	return updatedStudent, nil

}

func PatchStudents(updates []map[string]interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "invalid Id")
	}

	for _, update := range updates {
		id, ok := update["id"].(float64) // JSON numbers decode as float64
		if !ok {
			tx.Rollback()
			return utils.ErrorHandler(err, "invalid Id")
		}

		var studentFromDb models.Student
		err := db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", int(id)).
			Scan(&studentFromDb.ID, &studentFromDb.FirstName, &studentFromDb.LastName, &studentFromDb.Email, &studentFromDb.Class)
		if err != nil {
			log.Println("ID:", id)
			log.Printf("Type: %T", id)
			log.Println(err)
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Teacher not found")
			}
			//http.Error(w, "Error retrieving teacher", http.StatusInternalServerError)
			return utils.ErrorHandler(err, "Teacher not found")
		}

		// Apply updates using reflection
		studentVal := reflect.ValueOf(&studentFromDb).Elem()
		studentType := studentVal.Type()

		for k, v := range update {
			if k == "id" {
				// skip updating id field
				continue
			}

			// FIX: Use teacherType.NumField() or teacherVal.NumField()
			for i := 0; i < studentVal.NumField(); i++ {
				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("cannot convert %v to %v", val.Type(), fieldVal.Type())
							//http.Error(w, "Type conversion error", http.StatusBadRequest)
							return utils.ErrorHandler(err, "error updating data")
						}
					}
					break
				}
			}
		}

		// TODO: Add UPDATE query here to save changes to database
		_, err = tx.Exec("UPDATE students SET first_name=?, last_name=?, email=?, class=? WHERE id=?",
			studentFromDb.FirstName, studentFromDb.LastName, studentFromDb.Email, studentFromDb.Class, studentFromDb.ID)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "error updating data")
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Println(err)
		//http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "error updating data")
	}

	return nil

}

func PatchOneStudent(id int, updates map[string]interface{}) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		//http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "error updatiing data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(
		&existingStudent.ID,
		&existingStudent.FirstName,
		&existingStudent.LastName,
		&existingStudent.Email,
		&existingStudent.Class,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Student{}, utils.ErrorHandler(err, "Teacher not found")
		}
		log.Println(err)
		//http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "teacher not found")
	}

	//Introducing the Reflect package
	teacherVal := reflect.ValueOf(&existingStudent).Elem()
	teacherType := teacherVal.Type()

	for k, v := range updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			fmt.Println("k from reflect mechanism", k)
			field := teacherType.Field(i)
			fmt.Println(field.Tag.Get("json"))
			if field.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {
					fieldVal := teacherVal.Field(i)
					fmt.Println("fieldVal", fieldVal)
					fmt.Println("teacherVal.Field(i).Type():", teacherVal.Field(i).Type())
					fmt.Println("reflect.ValueOf(v): ", reflect.ValueOf(v))
					fieldVal.Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
				}
			}
		}
	}

	// Fixed: underscore for result, fixed column name, removed extra comma
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?",
		existingStudent.FirstName,
		existingStudent.LastName,
		existingStudent.Email,
		existingStudent.Class,
		existingStudent,
		existingStudent.ID,
	)
	if err != nil {
		log.Println(err)
		//http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "error upadting data")
	}

	return existingStudent, nil

}

func DeleteOneStudent(id int) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	result, err := db.Exec("DELETE  FROM teachers WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "Teacher not found")
	}

	return nil
}

func DeleteStudents(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error updating data")
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		return nil, utils.ErrorHandler(err, "error updating data")
	}

	defer stmt.Close()

	deleteIds := []int{}
	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "error updating data")
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "error updating data")
		}
		// if teacher was deleted then add ID to the deletedID slice
		if rowsAffected > 0 {
			deleteIds = append(deleteIds, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d nt found", id))
		}
	}

	//Commit
	err = tx.Commit()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error updating data")
	}

	if len(deleteIds) < 1 {
		return nil, utils.ErrorHandler(err, "IDs do not exist")
	}
	return deleteIds, nil
}
