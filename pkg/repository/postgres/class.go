package postgres

import (
	"fmt"

	"github.com/Murolando/hakaton_geo/ent"
	"github.com/jmoiron/sqlx"
)

type ClassPostgres struct {
	db *sqlx.DB
}

func NewClassPostgres(db *sqlx.DB) *ClassPostgres {
	return &ClassPostgres{
		db: db,
	}
}

func (r *ClassPostgres) DashboardClass(userId int64) ([]*ent.ClassProgressResponce, error) {

	list := make([]*ent.ClassProgressResponce, 0)
	query := fmt.Sprintf(`
	SELECT class.id,name
	FROM "%s"
	INNER JOIN "%s" on "%s".user_id = $1
	`, classTable, userClassTable, userClassTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var class ent.ClassProgressResponce
		if err := rows.Scan(&class.Id, &class.Name); err != nil {
			return nil, err
		}
		queryMaxProgress := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row := r.db.QueryRow(queryMaxProgress, class.Id)
		if err := row.Scan(&class.MaxProgressBar); err != nil {
			return nil, err
		}
		queryCurrentProgress := fmt.Sprintf(`
		SELECT COUNT(id) - (COUNT(id)/2)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row = r.db.QueryRow(queryCurrentProgress, class.Id)
		if err := row.Scan(&class.ProgressBar); err != nil {
			return nil, err
		}
		list = append(list, &class)
	}
	return list, nil
}

func (r *ClassPostgres) CommonProgressInfo(userId int64) (*ent.ChildDashClassResponce, error) {
	var responce ent.ChildDashClassResponce
	query := fmt.Sprintf(`
	SELECT class.id,name
	FROM "%s"
	INNER JOIN "%s" on "%s".user_id = $1

	`, classTable, userClassTable, userClassTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	var cnt int
	mxExSum := 0
	mxThSum := 0
	exSum := 0
	thSum := 0
	for rows.Next() {
		var class ent.ClassProgressResponce
		if err := rows.Scan(&class.Id, &class.Name); err != nil {
			return nil, err
		}
		queryMaxProgress := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row := r.db.QueryRow(queryMaxProgress, class.Id)
		if err := row.Scan(&cnt); err != nil {
			return nil, err
		}
		mxExSum += cnt

		queryTheoryProgress := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s"
		WHERE lesson_type_id = 1 AND class_id = $1`, lessonTable)
		row = r.db.QueryRow(queryTheoryProgress, class.Id)
		if err := row.Scan(&cnt); err != nil {
			return nil, err
		}
		mxThSum += cnt

		queryCurrentExProgress := fmt.Sprintf(`
		SELECT COUNT(id) - (COUNT(id)/2)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row = r.db.QueryRow(queryCurrentExProgress, class.Id)
		if err := row.Scan(&cnt); err != nil {
			return nil, err
		}
		exSum += cnt

		queryCurrentThProgress := fmt.Sprintf(`
		SELECT COUNT(id) - (COUNT(id)/2)
		FROM "%s"
		WHERE lesson_type_id = 1 AND class_id = $1`, lessonTable)
		row = r.db.QueryRow(queryCurrentThProgress, class.Id)
		if err := row.Scan(&cnt); err != nil {
			return nil, err
		}
		thSum += cnt
	}
	responce.MaxExProgressBar = mxExSum
	responce.MaxTheoryProgressBar = mxThSum
	responce.ExProgressBar = exSum
	responce.TheoryProgressBar = thSum
	return &responce, nil
}

func (r *ClassPostgres) MyClass(userId int64) ([]*ent.ChildMyClassResponce, error) {
	list := make([]*ent.ChildMyClassResponce, 0)
	query := fmt.Sprintf(`
	SELECT class.id,name
	FROM "%s"
	INNER JOIN "%s" on "%s".user_id = $1
	`, classTable, userClassTable, userClassTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var class ent.ChildMyClassResponce
		if err := rows.Scan(&class.Id, &class.Name); err != nil {
			return nil, err
		}
		queryMaxProgress := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row := r.db.QueryRow(queryMaxProgress, class.Id)
		if err := row.Scan(&class.MaxProgressBar); err != nil {
			return nil, err
		}
		queryCurrentProgress := fmt.Sprintf(`
		SELECT COUNT(id) - (COUNT(id)/2)
		FROM "%s"
		WHERE lesson_type_id != 1 AND class_id = $1`, lessonTable)
		row = r.db.QueryRow(queryCurrentProgress, class.Id)
		if err := row.Scan(&class.ProgressBar); err != nil {
			return nil, err
		}

		queryChildCount := fmt.Sprintf(`
		SELECT COUNT("user".id)
		FROM "%s"
		JOIN "%s" on "user".id = user_id
		WHERE class_id = $1 and role_id = 1`, userClassTable, userTable)
		row = r.db.QueryRow(queryChildCount, class.Id)
		if err := row.Scan(&class.ChildCount); err != nil {
			return nil, err
		}

		// TeacherName
		queryTeacherName := fmt.Sprintf(`
		SELECT "user".name
		FROM "%s"
		JOIN "%s" on role_id = 2
		WHERE class_id = $1`, userClassTable, userTable)
		row = r.db.QueryRow(queryTeacherName, class.Id)
		if err := row.Scan(&class.TeacherName); err != nil {
			return nil, err
		}
		list = append(list, &class)

	}
	return list, nil
}

func (r *ClassPostgres) IsClassMember(userId int64, classId int) (bool, error) {
	var num int
	query := fmt.Sprintf(`
	SELECT COUNT(id)
	FROM "%s"
	WHERE user_id = $1 AND class_id  = $2
	`, userClassTable)
	row := r.db.QueryRow(query, userId, classId)
	if err := row.Scan(&num); err != nil {
		return false, err
	}
	if num >0 {
		return true,nil
	}
	return false, nil
}

func (r *ClassPostgres) OneClass(classId int) (*ent.OneClassInfoResponce, error) {
	// fields
	var class ent.OneClassInfoResponce
	query := fmt.Sprintf(`
	SELECT id,name
	FROM "%s"
	WHERE id = $1
	`, classTable)
	row := r.db.QueryRow(query, classId)
	if err := row.Scan(&class.Id,&class.Name); err != nil {
		return nil, err
	}

	// users
	queryUsers := fmt.Sprintf(`
	SELECT u.name,u.role_id,r.name
	FROM "%s" 
	INNER JOIN "%s" u on u.id = "user_class".user_id
	INNER JOIN "%s" r on r.id = u.role_id
	WHERE "user_class".class_id = $1
	`, userClassTable, userTable, roleTable)
	rows, err := r.db.Query(queryUsers, classId)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var user ent.ClassUserResponce
		if err := rows.Scan(&user.Name, &user.RoleId,&user.RoleName); err != nil {
			return nil, err
		}
		class.User = append(class.User, &user)
	}
	// lessons
	queryLessons := fmt.Sprintf(`
	SELECT "lesson".name, video, created_at,expired_at,lesson_access,t.name
	FROM "%s" 
	JOIN "%s" t on t.id = lesson_type_id
	WHERE class_id = $1
	`,lessonTable,lessonTypeTable)
	rows, err = r.db.Query(queryLessons, classId)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var lesson ent.LessonInfoResponce
		if err := rows.Scan(&lesson.Name, &lesson.Video,&lesson.CreatedAt,&lesson.ExpiredAt,
			&lesson.LessonAccess,&lesson.LessonTypeName); err != nil {
			return nil, err
		}
		class.Lesson = append(class.Lesson, &lesson)
	}
	return &class, nil
}
