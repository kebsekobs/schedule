package db

import (
	"database/sql"
	"log"
	"slices"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kebsekobs/schedule/backend/internal/generation"
)

func InsertClasses(db *sql.DB, classes map[int]*generation.Class) error {
	query := "INSERT INTO classes (id, name, groupid, teacherid, hours, streamid) VALUES "
	for range classes {
		query += "(?, ?, ?, ?, ?, ?),"
	}
	query = query[:len(query)-1] // Remove the trailing comma
	query += " ON DUPLICATE KEY UPDATE name=VALUES(name), groupid=VALUES(groupid), teacherid=VALUES(teacherid), hours=VALUES(hours), streamid=VALUES(streamid)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}
	for classid, class := range classes {
		values = append(values, classid, class.Name, class.Group.ID, class.Teacher.ID, class.Hours, class.ID)
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

	rows, err := db.Query("SELECT id, groupid, teacherid, name, hours, streamid FROM classes")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupID, name string
		var id, hours, teacherID, streamID int
		err = rows.Scan(&id, &groupID, &teacherID, &name, &hours, &streamID)
		if err != nil {
			log.Println(err)
			continue
		}

		teacher := getTeacherByID(teachers, teacherID)
		if teacher == nil {
			log.Println("No teacher")
			continue
		}
		group := getGroupByID(groups, groupID)
		if group == nil {
			log.Println("No group")
			continue
		}

		if streamID == 0 {
			classes = append(
				classes,
				generation.Class{
					ID:      id,
					Teacher: teacher,
					Group:   group,
					Name:    name,
					Hours:   hours,
				})
		} else {
			class := getClassByID(commonClasses, streamID)
			if class != nil {
				if !slices.Contains(class.Groups, group) {
					class.Groups = append(class.Groups, group)
				}
				if class.Hours == 0 && hours != 0 {
					class.Hours = hours
				}
			} else {
				commonClasses = append(
					commonClasses,
					&generation.CommonClass{
						ID:      streamID,
						Teacher: teacher,
						Groups:  []*generation.Group{group},
						Name:    name,
						Hours:   hours,
					})
			}
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
