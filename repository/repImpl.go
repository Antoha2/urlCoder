package repository

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func (r *repositoryImplDB) RepAddLongUrl(url *RepLongUrl) error {
	//проверка на наличие свободных токенов
	var countTokens int64
	if err := r.rep.Table("urllist").Model(url).Count(&countTokens).Error; err != nil {
		log.Println(err)
		return err
	}
	if countTokens == 0 {
		return errors.New("в базе нет токенов")
	}
	var countUrls int64
	if err := r.rep.Table("urllist").Model(url).Where("long_url is null").Count(&countUrls).Error; err != nil {
		log.Println(err)
		return err
	}
	if countTokens != 0 && countUrls == 0 {
		return errors.New("нет свободных токенов")

	}

	//проверка на совпадения
	sqlConditionAUrl := fmt.Sprintf(" long_url = '%s'", url.Long_url)
	readUrl := new(RepLongUrl)
	r.rep.Table("urllist").Where(sqlConditionAUrl).Find(&readUrl).Scan(&readUrl)
	if readUrl.Id != 0 {
		url.Token = readUrl.Token
		errStr := "такой url уже есть в базе"
		log.Println(errStr)
		return errors.New(errStr)
	}

	var c int
	var count_id int64
	query := "select id from urlList order by id desc LIMIT 1"
	r.rep.Table("urlList").Raw(query).Scan(&c)

	//генерация случайного id и проверка на повторение
	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		url.Id = r1.Intn(c + 1)
		//log.Println("rand_id - ", rand_id)
		if err := r.rep.Table("urllist").Model(url).Where("id = ? AND long_url is null", url.Id).Count(&count_id).Error; err != nil {
			log.Println(err)
			return err
		}
		if count_id != 0 {
			break
		}
	}

	query = "update urllist set long_url = $1, created_at = $2 where id = $3 RETURNING token"
	result := r.rep.Table("urlList").Raw(query, url.Long_url, url.CreateAt, url.Id).Scan(&url.Token)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка получения токена")
	}

	log.Printf("id - %v, получен токен - %v \n", url.Id, url.Token)
	return nil
}

//генерация новых токенов
func (r *repositoryImplDB) RepGenTokens(q int) error {

	var c int
	query := "select id from urlList order by id desc LIMIT 1"
	r.rep.Table("urlList").Raw(query).Scan(&c)
	log.Printf("сейчас в таблице токенов - %d , необходимо еще добавить - %d \n", c, q)
	if c >= 238328 {
		return errors.New("достигнуто максимальное коичество токенов")
	}

	var c1 int
	id := 1
	s1 := 97
	for i1 := 0; i1 < 62; i1++ {
		s2 := 97
		for i2 := 0; i2 < 62; i2++ {
			if i2 == 26 {
				s2 = 65
			}
			if i2 == 52 {
				s2 = 48
			}
			s3 := 97
			for i3 := 0; i3 < 62; i3++ {
				if i3 == 26 {
					s3 = 65
				}
				if i3 == 52 {
					s3 = 48
				}
				if id > c && id <= (c+q) {
					token := fmt.Sprintf("%s%s%s", string(s1), string(s2), string(s3))
					query := "INSERT INTO urlList (token) VALUES ($1) RETURNING id"
					result := r.rep.Table("urlList").Raw(query, token).Scan(&c1)
					if errors.Is(result.Error, gorm.ErrInvalidValue) {
						return errors.New("ошибка сознания задачи")
					}
				}
				s3++
				id++

				if id > (c + q) {
					log.Println("создано токенов -", c1, " из 238328 возможных")
					return nil
				}
			}
			s2++
		}
		s1++
	}
	log.Println("создано макимальное количество токенов - 238328")
	return errors.New("достигнуто максимальное коичество токенов")
}
