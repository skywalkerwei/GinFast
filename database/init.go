package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// Departments 部门表
type Departments struct {
	gorm.Model
	ParentID int    `gorm:"column:parent_id;type:int(11);not null;default:0" json:"parentId"` // 上级部门
	Title    string `gorm:"column:title;type:varchar(512);not null;default:''" json:"title"`  // 部门名称
}

// DepartmentsColumns get sql column name.获取数据库列名
var DepartmentsColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	ParentID  string
	Title     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	ParentID:  "parent_id",
	Title:     "title",
}

// EvaluationParts 调查问卷块表
type EvaluationParts struct {
	gorm.Model
	EvaluationID int    `gorm:"column:evaluation_id;type:int(11);not null" json:"evaluationId"`  // 问卷id
	Title        string `gorm:"column:title;type:varchar(512);not null;default:''" json:"title"` // 块名称
}

// EvaluationPartsColumns get sql column name.获取数据库列名
var EvaluationPartsColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	EvaluationID string
	Title        string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	EvaluationID: "evaluation_id",
	Title:        "title",
}

// EvaluationQuestions 调查问卷块表
type EvaluationQuestions struct {
	gorm.Model
	PartID  int            `gorm:"column:part_id;type:int(10);not null" json:"partId"`                  // 块id
	Title   string         `gorm:"column:title;type:varchar(512);not null;default:''" json:"title"`     // 问题名称
	Score   float64        `gorm:"column:score;type:decimal(8,2);not null;default:0.00" json:"score"`   // 分数
	Weight  float64        `gorm:"column:weight;type:decimal(8,2);not null;default:0.00" json:"weight"` // 权重
	Sort    int            `gorm:"column:sort;type:int(10);not null;default:0" json:"sort"`             // 排序
	Content datatypes.JSON `gorm:"column:content;type:json;not null" json:"content"`                    // 绩效考评标准
}

// EvaluationQuestionsColumns get sql column name.获取数据库列名
var EvaluationQuestionsColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	PartID    string
	Title     string
	Score     string
	Weight    string
	Sort      string
	Content   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	PartID:    "part_id",
	Title:     "title",
	Score:     "score",
	Weight:    "weight",
	Sort:      "sort",
	Content:   "content",
}

// Evaluations 调查问卷表
type Evaluations struct {
	gorm.Model
	Title string `gorm:"column:title;type:varchar(512);not null;default:''" json:"title"` // 问卷名称
}

// EvaluationsColumns get sql column name.获取数据库列名
var EvaluationsColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Title     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Title:     "title",
}

// PositionEvaluation 职位-问卷绑定关系表
type PositionEvaluation struct {
	gorm.Model
	PositionID   int `gorm:"column:position_id;type:int(11);not null" json:"positionId"`     // 职位id
	EvaluationID int `gorm:"column:evaluation_id;type:int(11);not null" json:"evaluationId"` // 问卷id
	Weight       int `gorm:"column:weight;type:int(11);not null" json:"weight"`              // 权重
}

// PositionEvaluationColumns get sql column name.获取数据库列名
var PositionEvaluationColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	PositionID   string
	EvaluationID string
	Weight       string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	PositionID:   "position_id",
	EvaluationID: "evaluation_id",
	Weight:       "weight",
}

// PositionUser 员工职位表
type PositionUser struct {
	gorm.Model
	UId        int `gorm:"column:uid;type:int(11);not null" json:"uid"`                // 用户id
	PositionID int `gorm:"column:position_id;type:int(11);not null" json:"positionId"` // 职位id
}

// PositionUserColumns get sql column name.获取数据库列名
var PositionUserColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	UId        string
	PositionID string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	UId:        "uid",
	PositionID: "position_id",
}

// Positions 部门职位表
type Positions struct {
	gorm.Model
	DepartmentID int    `gorm:"column:department_id;type:int(11);not null" json:"departmentId"`   // 部门
	ParentID     int    `gorm:"column:parent_id;type:int(11);not null;default:0" json:"parentId"` // 职位上级
	Title        string `gorm:"column:title;type:varchar(512);not null;default:''" json:"title"`  // 职位名称
}

// PositionsColumns get sql column name.获取数据库列名
var PositionsColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	DepartmentID string
	ParentID     string
	Title        string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	DepartmentID: "department_id",
	ParentID:     "parent_id",
	Title:        "title",
}

// Respondent 调查
type Respondent struct {
	gorm.Model
	FormUId      int            `gorm:"column:form_uid;type:int(11);not null;default:0" json:"formUid"`           // 调查人id
	EvaluationID int            `gorm:"column:evaluation_id;type:int(11);not null;default:0" json:"evaluationId"` // 问卷id
	AnswerUId    int            `gorm:"column:answer_uid;type:int(11);not null;default:0" json:"answerUid"`       // 填写问卷人id
	Score        float64        `gorm:"column:score;type:decimal(8,2);not null" json:"score"`                     // 分数
	Content      datatypes.JSON `gorm:"column:content;type:json;not null" json:"content"`                         // 填写内容
	Status       bool           `gorm:"column:status;type:tinyint(1);not null;default:0" json:"status"`           // 完成状态
}

// RespondentColumns get sql column name.获取数据库列名
var RespondentColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	FormUId      string
	EvaluationID string
	AnswerUId    string
	Score        string
	Content      string
	Status       string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	FormUId:      "form_uid",
	EvaluationID: "evaluation_id",
	AnswerUId:    "answer_uid",
	Score:        "score",
	Content:      "content",
	Status:       "status",
}

// UserEvaluation 用户-问卷绑定关系表
type UserEvaluation struct {
	gorm.Model
	UId          int     `gorm:"column:uid;type:int(11);not null" json:"uid"`                    // 用户id
	EvaluationID int     `gorm:"column:evaluation_id;type:int(11);not null" json:"evaluationId"` // 问卷id
	Score        float64 `gorm:"column:score;type:decimal(8,2);not null" json:"score"`           // 获取重分数
	Weight       int     `gorm:"column:weight;type:int(11);not null" json:"weight"`              // 权重
}

// UserEvaluationColumns get sql column name.获取数据库列名
var UserEvaluationColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	UId          string
	EvaluationID string
	Score        string
	Weight       string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	UId:          "uid",
	EvaluationID: "evaluation_id",
	Score:        "score",
	Weight:       "weight",
}

// Users 用户表
type Users struct {
	gorm.Model
	Mobile  string    `gorm:"column:mobile;type:char(11);not null;default:''" json:"mobile"`                      // 手机号
	Name    string    `gorm:"column:name;type:varchar(64);not null;default:''" json:"name"`                       // 名字
	Avatar  string    `gorm:"column:avatar;type:varchar(512);not null;default:''" json:"avatar"`                  // 头像
	Sex     bool      `gorm:"column:sex;type:tinyint(1);not null;default:0" json:"sex"`                           // sex1 男 2女 0未知
	LoginAt time.Time `gorm:"column:login_at;type:timestamp;not null;default:2022-01-01 00:00:00" json:"loginAt"` // 登录时间
	LoginIP string    `gorm:"column:login_ip;type:varchar(32);not null;default:''" json:"loginIp"`                // 登录IP
	Status  bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`                     // 状态1正常 0禁用
	Openid  string    `gorm:"column:openid;type:varchar(64);not null;default:''" json:"openid"`                   // openid
}

// UsersColumns get sql column name.获取数据库列名
var UsersColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Mobile    string
	Name      string
	Avatar    string
	Sex       string
	LoginAt   string
	LoginIP   string
	Status    string
	Openid    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Mobile:    "mobile",
	Name:      "name",
	Avatar:    "avatar",
	Sex:       "sex",
	LoginAt:   "login_at",
	LoginIP:   "login_ip",
	Status:    "status",
	Openid:    "openid",
}
