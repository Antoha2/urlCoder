package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (r *repositoryImplDB) LongUrl(url *RepUrl) error {
	query := "INSERT INTO urlList (long_url, created_at) VALUES ($1, $2) RETURNING id"
	result := r.rep.Table("urlList").Raw(query, url.Long_url, url.CreateAt).Scan(&url.Id)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка сознания задачи")
	}

	log.Println("создана запись - ", url)
	return nil
}

/*
func (r *repositoryImplDB) Create(task *RepTask) error {

	var count int
	stmtCount, err := r.rep.DB.Query("SELECT count(task_id) FROM todolist WHERE user_id = $1", task.UserId)
	if err != nil {
		panic(err)
	}
	for stmtCount.Next() {

		err := stmtCount.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	if count < countTask {

		var id int
		query := "INSERT INTO todolist (user_id, text, isdone) VALUES ($1, $2, $3) RETURNING task_id"
		row := r.rep.DB.QueryRow(query, task.UserId, task.Text, task.IsDone)
		if err := row.Scan(&id); err != nil {
			panic(err)
		}
		task.Id = id
		fmt.Println("создана запись - ", task)
		return nil
	}

	errStr := fmt.Sprintf("не больше %d записей на пользователя", countTask)
	return errors.New(errStr)

}*/
