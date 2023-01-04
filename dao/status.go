package dao

import "github.com/SXUOJ/backend/models"

func InsertStatus(result models.Result) error {
	return db.Create(&models.ResultSql{Result: result}).Error
}

// 通过问题id和用户id获取Status列表
// qid: 问题id
// id: 	用户id
// amount: 单页展示数量
// page:页号
func GetStatusListByQid(qid string, uid string, amount int, page int) ([]*models.Result, error) {
	var (
		offset     = (page - 1) * amount
		resultSqls []models.ResultSql
		results    []*models.Result
	)

	db.Limit(amount).Offset(offset).Find(&resultSqls)

	for _, v := range resultSqls {
		results = append(results, &v.Result)
	}

	return results, nil
}

// 通过提交id获取status详细
// submit_id: 提交id
func GetStatusByAdmitId(submit_id string) (models.Result, error) {
	var resultSql models.ResultSql
	tx := db.Where("submit_id = ?", submit_id).First(&resultSql)
	//返回

	return resultSql.Result, tx.Error
}
