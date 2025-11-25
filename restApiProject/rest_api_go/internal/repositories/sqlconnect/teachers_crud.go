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

func GetTeachersDbHandler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving Data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1 "
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
		var teacher models.Teacher
		err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			//http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error retrieving Data")
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil

}

func AddTeachersDBHandler(newTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		//http.Error(w, "Error connction to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error retrieving Data")

	}
	defer db.Close()

	//stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject)VALUES (?,?,?,?,?)")
	stmt, err := db.Prepare(utils.GenerateInsertQuery("teachers", models.Teacher{}))
	if err != nil {
		//http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error retrieving Data")
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		//res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)

		values := utils.GetStructValues(newTeacher)
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
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher

	}

	return addedTeachers, nil
}

func GetTeacherByID(id int) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "error retrieving data")
	} else if err != nil {
		fmt.Println(err)
		//http.Error(w, "Database query error", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "errr retrieving data")
	}
	return teacher, nil
}

func UpdateTeacher(id int, updatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)

		return models.Teacher{}, utils.ErrorHandler(err, "errr Updating data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(
		&existingTeacher.ID,
		&existingTeacher.FirstName,
		&existingTeacher.LastName,
		&existingTeacher.Email,
		&existingTeacher.Class,
		&existingTeacher.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, utils.ErrorHandler(err, "errr retrieving data")

		}
		log.Println(err)
		//	http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "errr retrieving data")

	}

	updatedTeacher.ID = existingTeacher.ID

	// Fixed: underscore for result, fixed column name, removed extra comma
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
		updatedTeacher.FirstName,
		updatedTeacher.LastName,
		updatedTeacher.Email,
		updatedTeacher.Class,
		updatedTeacher.Subject,
		updatedTeacher.ID,
	)
	if err != nil {
		log.Println(err)
		//	http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "errr retrieving data")

	}

	return updatedTeacher, nil

}

func PatchTeachers(updates []map[string]interface{}) error {
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

		var teacherFromDb models.Teacher
		err := db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", int(id)).
			Scan(&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject)
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
		teacherVal := reflect.ValueOf(&teacherFromDb).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				// skip updating id field
				continue
			}

			// FIX: Use teacherType.NumField() or teacherVal.NumField()
			for i := 0; i < teacherType.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
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
		_, err = tx.Exec("UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id=?",
			teacherFromDb.FirstName, teacherFromDb.LastName, teacherFromDb.Email, teacherFromDb.Class, teacherFromDb.Subject, teacherFromDb.ID)
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

func PatchOneTeacher(id int, updates map[string]interface{}) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		//http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "error updatiing data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(
		&existingTeacher.ID,
		&existingTeacher.FirstName,
		&existingTeacher.LastName,
		&existingTeacher.Email,
		&existingTeacher.Class,
		&existingTeacher.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
		}
		log.Println(err)
		//http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "teacher not found")
	}

	//Introducing the Reflect package
	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
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
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
		existingTeacher.FirstName,
		existingTeacher.LastName,
		existingTeacher.Email,
		existingTeacher.Class,
		existingTeacher.Subject,
		existingTeacher.ID,
	)
	if err != nil {
		log.Println(err)
		//http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Teacher{}, utils.ErrorHandler(err, "error upadting data")
	}

	return existingTeacher, nil

}

func DeleteOneTeacher(id int) error {
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

func DeleteTeachers(ids []int) ([]int, error) {
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
