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

	token := new(RepToken)

	//подсчет общего количества токенов в базе
	var countTokens int64
	if err := r.rep.Table("tokenlist").Model(url).Count(&countTokens).Error; err != nil {
		log.Println(err)
		return err
	}
	if countTokens == 0 {
		return errors.New("в базе нет свободных токенов")
	}

	//проверка на совпадения
	sqlConditionAUrl := fmt.Sprintf(" long_url = '%s'", url.Long_url)
	readUrl := new(RepLongUrl)
	r.rep.Table("urllist").Where(sqlConditionAUrl).Find(&readUrl).Scan(&readUrl)
	if readUrl.Id != 0 {
		url.Token_id = readUrl.Token_id
		token.Id = readUrl.Token_id
		url.Id = readUrl.Id
		if err := r.rep.Table("tokenlist").Find(&token, "token_id = ?", token.Id).Scan(&token).Error; err != nil {
			log.Println(err)
			return err
		}
		url.Token = token.Token
		log.Println("такой url уже есть в базе")
		return nil //errors.New("такой url уже есть в базе")
	}

	//генерация случайного id токена
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	token.Id = r1.Intn(int(countTokens))

	if err := r.rep.Table("tokenlist").Find(&token, "token_id = ?", token.Id).Scan(&token).Error; err != nil {
		log.Println(err)
		return err
	}

	url.Token = token.Token
	url.Token_id = token.Id

	result := r.rep.Table("urllist").Select("id", "token_id", "long_url", "created_at", "expiry_at").Create(&url)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка создания записи")
	}

	return nil
}

//генерация новых токенов
func (r *repositoryImplDB) RepGenTokens() error {

	var count int64
	var tok RepToken
	if err := r.rep.Table("tokenlist").Model(tok).Where("token_id > 0").Count(&count).Error; err != nil {
		log.Println(err)
		return err
	}
	if count > 0 {
		err := fmt.Sprintf("в базе уже есть %v токенов", count)
		log.Println(err)
		return errors.New(err)
	}

	log.Println("начало генерации токенов")
	countTokens := 2 * 62 * 62 //количество токенов

	t := make([]RepToken, countTokens)

	id := 0
	s1 := 97
	for i1 := 0; i1 < 2; i1++ {
		if i1 == 26 {
			s1 = 65
		}
		if i1 == 52 {
			s1 = 48
		}
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

				token := fmt.Sprintf("%s%s%s", string(s1), string(s2), string(s3))
				if id <= countTokens {
					for {
						r1 := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(countTokens) + 1
						if t[r1-1].Id == 0 {
							t[r1-1].Token = token
							t[r1-1].Id = r1
							break
						}
					}
				}
				s3++
				id++

			}
			s2++
		}
		s1++
	}
	if err := r.rep.Table("tokenlist").Create(t).Error; err != nil {
		log.Println(err)
		return err
	}
	log.Println("создание токенов завершено")

	return nil
}
