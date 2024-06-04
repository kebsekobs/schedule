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
	var resp Response
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
	var resp Response

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

	json.NewEncoder(w).Encode(groups)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groupID := r.URL.Path[len("/groups/"):]
	var newGroup Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Implement group update logic here
	w.WriteHeader(http.StatusOK)
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	groupID := r.URL.Path[len("/groups/"):]

	// Implement group deletion logic here
	w.WriteHeader(http.StatusOK)
}

func addGroup(w http.ResponseWriter, r *http.Request) {
	var newGroup Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement group add logic here
	w.WriteHeader(http.StatusOK)
}

// ПРЕПОДАВАТЕЛИ

func getTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

func getTeacherByID(w http.ResponseWriter, r *http.Request) {
	teacherID := r.URL.Path[len("/teachers/"):]

	// Implement teacher get logic here
	http.Error(w, "Teacher not found", http.StatusNotFound)
}

func updateTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := r.URL.Path[len("/teachers/"):]
	var updatedTeacher Teacher
	err := json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the teacher

	http.Error(w, "Teacher not found", http.StatusNotFound)
}

func addTeacher(w http.ResponseWriter, r *http.Request) {
	var newTeacher Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new teacher
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTeacher)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := r.URL.Path[len("/teachers/"):]
	// Remove the teacher

	http.Error(w, "Teacher not found", http.StatusNotFound)
}

// АУДИТОРИИ

func getClassrooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}

func updateClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.URL.Path[len("/classrooms/"):]
	var updatedClassroom Classroom
	err := json.NewDecoder(r.Body).Decode(&updatedClassroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the classroom

	http.Error(w, "Classroom not found", http.StatusNotFound)
}

func deleteClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID := r.URL.Path[len("/classrooms/"):]
	// Remove the classroom

	http.Error(w, "Classroom not found", http.StatusNotFound)
}

func addClassroom(w http.ResponseWriter, r *http.Request) {
	var newClassroom Classroom
	err := json.NewDecoder(r.Body).Decode(&newClassroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClassroom)
}

// ПАРЫ

func getDisciplines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(disciplines)
}

func updateDiscipline(w http.ResponseWriter, r *http.Request) {
	disciplineID := r.URL.Path[len("/disciplines/"):]
	var updatedDiscipline Discipline
	err := json.NewDecoder(r.Body).Decode(&updatedDiscipline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, "Discipline not found", http.StatusNotFound)
}

func deleteDiscipline(w http.ResponseWriter, r *http.Request) {
	disciplineID := r.URL.Path[len("/disciplines/"):]

	http.Error(w, "Discipline not found", http.StatusNotFound)
}

func addDiscipline(w http.ResponseWriter, r *http.Request) {
	var newDiscipline Discipline
	err := json.NewDecoder(r.Body).Decode(&newDiscipline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDiscipline)
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
