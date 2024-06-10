package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"

	api "github.com/kebsekobs/schedule/backend/internal/apientity"
	"github.com/kebsekobs/schedule/backend/internal/config"
	"github.com/kebsekobs/schedule/backend/internal/db"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

var (
	_db = dbConnect()
)

// конектимся к бд
func dbConnect() *sql.DB {
	fmt.Println(config.DBConfig.GetString("user"))
	localdb, err := sql.Open("mysql", config.DBConfig.GetString("user")+":"+config.DBConfig.GetString(
		"password")+"@/"+config.DBConfig.GetString("name"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected successfully.")
	}
	// настройка пула соединений
	localdb.SetConnMaxLifetime(time.Minute * 60)
	localdb.SetMaxOpenConns(1)
	localdb.SetMaxIdleConns(1)
	return localdb
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Это домашняя страница 🏡")
}

func parseData(w http.ResponseWriter, r *http.Request) {
	var resp api.Response
	w.Header().Set("Content-Type", "application/json")

	// парсинг аудиторий из rooms.xlsx
	rooms, err := generation.ParseRooms()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Rooms parsing error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// вставка аудиторий в таблицу
	err = db.InsertRooms(_db, rooms)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Rooms insert error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// парсинг занятий из classes.xlsx
	classes, groups, teachers, err := generation.ParseClasses()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Classes parsing error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// вставка аудиторий в таблицу
	err = db.InsertRooms(_db, rooms)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Rooms insert error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// вставка групп в таблицу
	err = db.InsertGroups(_db, groups)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Groups insert error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// вставка учителей в таблицу
	teachers, err = db.InsertTeachers(_db, teachers)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Teachers insert error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"message":"Error marshal data"}`, http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	for _, class := range classes {
		class.Teacher = teachers[class.Teacher.Name]
	}

	// вставка занятий в таблицу
	err = db.InsertClasses(_db, classes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Classes insert error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Message = "OK"
	response, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error processing data", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func generate(w http.ResponseWriter, r *http.Request) {
	var resp api.Response

	rooms, err := db.SelectRooms(_db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Rooms select error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	groups, err := db.SelectGroups(_db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Groups select error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	teachers, err := db.SelectTeachers(_db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Teachers select error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	classes, commonClasses, err := db.SelectClasses(_db, groups, teachers)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Rooms select error"
		response, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}

	// classesWithHours := make([]generation.ClassWithHours, len(classes))
	// for i, class := range classes {
	// 	classesWithHours[i] = class
	// }

	// commonClassesWithHours := make([]generation.ClassWithHours, len(commonClasses))
	// for i, commonClass := range commonClasses {
	// 	commonClassesWithHours[i] = commonClass
	// }

	// filteredClasses := generation.FilterClassesByHours(classesWithHours)
	// for i, class := range filteredClasses {
	// 	classes[i] = class.(*generation.Class)
	// }
	// filteredCommonClasses := generation.FilterClassesByHours(commonClassesWithHours)
	// for i, commonClass := range filteredCommonClasses {
	// 	commonClasses[i] = commonClass.(*generation.CommonClass)
	// }

	// for _, c := range commonClasses {
	// 	for _, v := range c.Groups {
	// 		if v.ID == "БOЮ35-ИАФ2001" {
	// 			log.Println(c.ID)
	// 		}
	// 	}
	// }
	filteredCommonClasses := generation.FilterCommonClassesByHours(commonClasses)
	filteredClasses := generation.FilterClassesByHours(classes)
	sort.Sort(generation.ByGroupQuantitySum(filteredCommonClasses))

	generation := generation.Generation{
		Groups:        groups,
		CommonClasses: filteredCommonClasses,
		Classes:       filteredClasses,
		Rooms:         rooms,
		Hours:         5,
		Days:          12,
	}

	generation.StartGeneration()
}

// ГРУППЫ

func getGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Implement group get logic here
	groups, err := db.GetGroups(_db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(groups)
	w.WriteHeader(http.StatusOK)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// groupID := r.URL.Path[len("/groups/"):]
	var newGroup api.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Implement group update logic here
	err = db.UpdateGroup(_db, newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	// groupID := r.URL.Path[len("/groups/"):]
	var newGroup api.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Implement group deletion logic here
	err = db.DeleteGroup(_db, newGroup.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func addGroup(w http.ResponseWriter, r *http.Request) {
	var newGroup api.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement group add logic here
	err = db.CreateGroup(_db, newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ПРЕПОДАВАТЕЛИ

func getTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	teachers, err := db.GetTeachers(_db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(teachers)
	w.WriteHeader(http.StatusOK)
}

func getTeacherByID(w http.ResponseWriter, r *http.Request) {
	// teacherID := r.URL.Path[len("/teachers/"):]
	var selectedTeacher api.Teacher
	err := json.NewDecoder(r.Body).Decode(&selectedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	teacher, err := db.GetTeacherByID(_db, selectedTeacher.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(teacher)
	w.WriteHeader(http.StatusOK)
}

func updateTeacher(w http.ResponseWriter, r *http.Request) {
	// teacherID := r.URL.Path[len("/teachers/"):]
	var updatedTeacher api.Teacher
	err := json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the teacher
	err = db.UpdateTeacher(_db, updatedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func addTeacher(w http.ResponseWriter, r *http.Request) {
	var newTeacher api.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new teacher
	err = db.CreateTeacher(_db, newTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	var newTeacher api.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new teacher
	err = db.DeleteTeacher(_db, newTeacher.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// АУДИТОРИИ

func getClassrooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rooms, err := db.GetRooms(_db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(rooms)
	w.WriteHeader(http.StatusOK)
}

func updateClassroom(w http.ResponseWriter, r *http.Request) {
	var updatedRoom api.Classroom
	err := json.NewDecoder(r.Body).Decode(&updatedRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the room
	err = db.UpdateRoom(_db, updatedRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteClassroom(w http.ResponseWriter, r *http.Request) {
	var newRoom api.Classroom
	err := json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new room
	err = db.DeleteRoom(_db, newRoom.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func addClassroom(w http.ResponseWriter, r *http.Request) {
	var newRoom api.Classroom
	err := json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new room
	err = db.CreateRoom(_db, newRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ПАРЫ

func getDisciplines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	classes, err := db.GetClasses(_db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(classes)
	w.WriteHeader(http.StatusOK)
}

func updateDiscipline(w http.ResponseWriter, r *http.Request) {
	var updatedClass api.Discipline
	err := json.NewDecoder(r.Body).Decode(&updatedClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the class
	err = db.UpdateClass(_db, updatedClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteDiscipline(w http.ResponseWriter, r *http.Request) {
	var delClass api.Discipline
	err := json.NewDecoder(r.Body).Decode(&delClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new room
	err = db.DeleteClass(_db, delClass.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func addDiscipline(w http.ResponseWriter, r *http.Request) {
	var newClass api.Discipline
	err := json.NewDecoder(r.Body).Decode(&newClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new class
	err = db.CreateClass(_db, newClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RunServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/parsedata", parseData)
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getGroups(w, r)
		case http.MethodPost:
			addGroup(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/groups/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateGroup(w, r)
		case http.MethodDelete:
			deleteGroup(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTeachers(w, r)
		case http.MethodPost:
			addTeacher(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/teachers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTeacherByID(w, r)
		case http.MethodPut:
			updateTeacher(w, r)
		case http.MethodDelete:
			deleteTeacher(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/classrooms", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getClassrooms(w, r)
		case http.MethodPost:
			addClassroom(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/classrooms/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateClassroom(w, r)
		case http.MethodDelete:
			deleteClassroom(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/disciplines", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getDisciplines(w, r)
		case http.MethodPost:
			addDiscipline(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/disciplines/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateDiscipline(w, r)
		case http.MethodDelete:
			deleteDiscipline(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
