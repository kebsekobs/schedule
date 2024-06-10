package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	api "github.com/kebsekobs/schedule/backend/internal/apientity"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

// CRUD методы для таблицы "classes"

func CreateClass(db *sql.DB, class api.Discipline) error {
	err := deleteClassGroupsLinks(db, class.ID)
	if err != nil {
		return err
	}
	err = insertClassesGroupsLinks(db, class.ID, class.RelatedGroupsId)
	if err != nil {
		return err
	}
	query := "INSERT INTO classes (name, teacherid, hours) VALUES (?,?,?)"
	_, err = db.Exec(query, class.Name, class.Teachers, class.Hours)
	if err != nil {
		return err
	}
	return nil
}

func GetClasses(db *sql.DB) ([]api.Discipline, error) {
	query := "SELECT id, name, teacherid, hours FROM classes"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var classes []api.Discipline
	for rows.Next() {
		var class api.Discipline
		err := rows.Scan(&class.ID, &class.Name, &class.Teachers, &class.Hours)
		if err != nil {
			return nil, err
		}
		class.RelatedGroupsId, err = getClassesGroupsLinks(db, class.ID)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func UpdateClass(db *sql.DB, class api.Discipline) error {
	err := deleteClassGroupsLinks(db, class.ID)
	if err != nil {
		return err
	}
	err = insertClassesGroupsLinks(db, class.ID, class.RelatedGroupsId)
	if err != nil {
		return err
	}
	query := "UPDATE classes SET name = ?, teacherid = ?, hours = ? WHERE id = ?"
	_, err = db.Exec(query, class.Name, class.Teachers, class.Hours, class.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteClass(db *sql.DB, id int) error {
	query := "DELETE FROM classes WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	err = deleteClassGroupsLinks(db, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertClasses(db *sql.DB, classes map[int]*generation.Class) error {
	query := "INSERT INTO classes (id, name, teacherid, hours) VALUES "
	for range classes {
		query += "(?, ?, ?, ?),"
	}
	query = query[:len(query)-1] // Remove the trailing comma
	query += " ON DUPLICATE KEY UPDATE name=VALUES(name), teacherid=VALUES(teacherid), hours=VALUES(hours)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}
	for classid, class := range classes {
		if class.ID == 0 {
			values = append(values, classid, class.Name, class.Teacher.ID, class.Hours)
			err = insertClassesGroupsLinks(db, classid, []string{class.Group.ID})
			if err != nil {
				return err
			}
		} else {
			values = append(values, class.ID, class.Name, class.Teacher.ID, class.Hours)
			err = insertClassesGroupsLinks(db, class.ID, []string{class.Group.ID})
			if err != nil {
				return err
			}
		}

	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}

func SelectClasses(db *sql.DB, groups []*generation.Group, teachers []*generation.Teacher) ([]generation.Class, []*generation.CommonClass, error) {
	var classes []generation.Class

	var commonClasses []*generation.CommonClass

	rows, err := db.Query("SELECT id, teacherid, name, hours FROM classes")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id, hours, teacherID int
		var groupIds []*generation.Group
		err = rows.Scan(&id, &teacherID, &name, &hours)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdsRows, err := db.Query("SELECT groupid FROM classes_groups WHERE classid = ?")
		if err != nil {
			log.Println(err)
			continue
		}
		defer groupIdsRows.Close()
		for groupIdsRows.Next() {
			var groupID string
			err = rows.Scan(&groupID)
			if err != nil {
				log.Println(err)
				continue
			}
			group := getGroupByID(groups, groupID)
			if group == nil {
				log.Println("No group")
				continue
			}
			groupIds = append(groupIds, group)
		}
		teacher := getTeacherByID(teachers, teacherID)
		if teacher == nil {
			log.Println("No teacher")
			continue
		}

		if len(groupIds) == 1 {
			classes = append(
				classes,
				generation.Class{
					ID:      id,
					Teacher: teacher,
					Group:   groupIds[0],
					Name:    name,
					Hours:   hours,
				})
		} else {
			commonClasses = append(
				commonClasses,
				&generation.CommonClass{
					ID:      id,
					Teacher: teacher,
					Groups:  groupIds,
					Name:    name,
					Hours:   hours,
				})
		}
	}

	commonClassesNoPointers := make([]generation.CommonClass, len(commonClasses))
	for i, cc := range commonClasses {
		commonClassesNoPointers[i] = *cc
	}

	return classes, commonClasses, nil
}

func getTeacherByID(teachers []*generation.Teacher, teacherID int) *generation.Teacher {
	for _, teacher := range teachers {
		if teacher.ID == teacherID {
			return teacher
		}
	}
	return nil
}

func getGroupByID(groups []*generation.Group, groupID string) *generation.Group {
	for _, group := range groups {
		if group.ID == groupID {
			return group
		}
	}
	return nil
}

func getClassByID(classes []*generation.CommonClass, id int) *generation.CommonClass {
	for _, class := range classes {
		if class.ID == id {
			return class
		}
	}
	return nil
}

func getClassesGroupsLinks(db *sql.DB, classID int) ([]string, error) {
	query := "SELECT groupid FROM classes_groups WHERE classid = ?"
	var groupIds []string
	log.Println(1111)
	rows, err := db.Query(query, classID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(2222)

	defer rows.Close()
	for rows.Next() {
		var groupId string
		err = rows.Scan(&groupId)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIds = append(groupIds, groupId)
	}

	return groupIds, nil
}

func insertClassesGroupsLinks(db *sql.DB, classID int, groupIds []string) error {
	query := "INSERT IGNORE INTO classes_groups (classid, groupid) VALUES "
	for range groupIds {
		query += "(?, ?),"
	}
	query = query[:len(query)-1] // Remove the trailing comma

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}
	for _, groupId := range groupIds {
		values = append(values, classID, groupId)
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}

func deleteClassGroupsLinks(db *sql.DB, id int) error {
	query := "DELETE FROM classes_groups WHERE classid = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
