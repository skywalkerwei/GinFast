package model

import (
	"ginfast/app/global/variable"
	"ginfast/app/utils/md5_encrypt"
	"go.uber.org/zap"
	"time"
)

// 创建 userFactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）

func CreateAdminUserModel(sqlType string) *AdminUsersModel {
	return &AdminUsersModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type AdminUsersModel struct {
	BaseModel   `json:"-"`
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Pass        string `json:"-"`
	Phone       string `json:"phone"`
	RealName    string `gorm:"column:real_name" json:"real_name"`
	Status      int    `json:"status"`
	Token       string `json:"token"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
}

// 表名
func (u *AdminUsersModel) TableName() string {
	return "tb_users"
}

// 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *AdminUsersModel) Register(userName, pass, userIp string) bool {
	sql := "INSERT  INTO tb_users(user_name,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  user_name=?)"
	result := u.Exec(sql, userName, pass, userIp, userName)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// 用户登录,
func (u *AdminUsersModel) Login(userName string, pass string) *AdminUsersModel {
	sql := "select id, user_name,real_name,pass,phone  from tb_users where  user_name=?  limit 1"
	result := u.Raw(sql, userName).First(u)
	if result.Error == nil {
		// 账号密码验证成功
		if len(u.Pass) > 0 && (u.Pass == md5_encrypt.Base64Md5(pass)) {
			return u
		}
	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}

//记录用户登陆（login）生成的token，每次登陆记录一次token
func (u *AdminUsersModel) OauthLoginToken(userId int64, token string, expiresAt int64, clientIp string) bool {
	sql := "INSERT   INTO  `tb_oauth_access_tokens`(fr_user_id,`action_name`,token,expires_at,client_ip) " +
		"SELECT  ?,'login',? ,?,? FROM DUAL    WHERE   NOT   EXISTS(SELECT  1  FROM  `tb_oauth_access_tokens` a WHERE  a.fr_user_id=?  AND a.action_name='login' AND a.token=?)"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if u.Exec(sql, userId, token, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, token).Error == nil {
		return true
	}
	return false
}

//用户刷新token,条件检查: 相关token在过期的时间之内，就符合刷新条件
func (u *AdminUsersModel) OauthRefreshConditionCheck(userId int64, oldToken string) bool {
	// 首先判断旧token在本系统自带的数据库已经存在，才允许继续执行刷新逻辑
	var oldTokenIsExists int
	sql := "SELECT count(*)  as  counts FROM tb_oauth_access_tokens  WHERE fr_user_id =? and token=? and NOW()<DATE_ADD(expires_at,INTERVAL  ? SECOND)"
	if u.Raw(sql, userId, oldToken, variable.ConfigYml.GetInt64("Token.JwtTokenRefreshAllowSec")).First(&oldTokenIsExists).Error == nil && oldTokenIsExists == 1 {
		return true
	}
	return false
}

//用户刷新token
func (u *AdminUsersModel) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   tb_oauth_access_tokens   SET  token=? ,expires_at=?,client_ip=?,updated_at=NOW(),action_name='refresh'  WHERE   fr_user_id=? AND token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, oldToken).Error == nil {
		go u.UpdateUserloginInfo(clientIp, userId)
		return true
	}
	return false
}

// 更新用户登陆次数、最近一次登录ip、最近一次登录时间
func (u *AdminUsersModel) UpdateUserloginInfo(last_login_ip string, userId int64) {
	sql := "UPDATE  tb_users   SET  login_times=IFNULL(login_times,0)+1,last_login_ip=?,last_login_time=?  WHERE   id=?  "
	_ = u.Exec(sql, last_login_ip, time.Now().Format(variable.DateFormat), userId)
}

//当用户更改密码后，所有的token都失效，必须重新登录
func (u *AdminUsersModel) OauthResetToken(userId int, newPass, clientIp string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	userItem, err := u.ShowOneItem(userId)
	if userItem != nil && err == nil && userItem.Pass == newPass {
		return true
	} else if userItem != nil {
		sql := "UPDATE  tb_oauth_access_tokens  SET  revoked=1,updated_at=NOW(),action_name='ResetPass',client_ip=?  WHERE  fr_user_id=?  "
		if u.Exec(sql, clientIp, userId).Error == nil {
			return true
		}
	}
	return false
}

//当tb_users 删除数据，相关的token同步删除
func (u *AdminUsersModel) OauthDestroyToken(userId int) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	sql := "DELETE FROM  tb_oauth_access_tokens WHERE  fr_user_id=?  "
	//判断>=0, 有些没有登录过的用户没有相关token，此语句执行影响行数为0，但是仍然是执行成功
	if u.Exec(sql, userId).Error == nil {
		return true
	}
	return false
}

// 判断用户token是否在数据库存在+状态OK
func (u *AdminUsersModel) OauthCheckTokenIsOk(userId int64, token string) bool {
	sql := "SELECT   token  FROM  `tb_oauth_access_tokens`  WHERE   fr_user_id=?  AND  revoked=0  AND  expires_at>NOW() ORDER  BY  expires_at  DESC , updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, userId, maxOnlineUsers).Rows()
	defer func() {
		//  凡是查询类记得释放记录集
		_ = rows.Close()
	}()

	if err == nil && rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					return true
				}
			}
		}
	}
	return false
}

// 禁用一个用户的: 1.tb_users表的 status 设置为 0，tb_oauth_access_tokens 表的所有token删除
// 禁用一个用户的token请求（本质上就是把tb_users表的 status 字段设置为 0 即可）
func (u *AdminUsersModel) SetTokenInvalid(userId int) bool {
	sql := "delete from  `tb_oauth_access_tokens`  where  `fr_user_id`=?  "
	if u.Exec(sql, userId).Error == nil {
		if u.Exec("update  tb_users  set  status=0 where   id=?", userId).Error == nil {
			return true
		}
	}
	return false
}

//根据用户ID查询一条信息
func (u *AdminUsersModel) ShowOneItem(userId int) (*AdminUsersModel, error) {
	sql := "SELECT  `id`, `user_name`,`pass`, `real_name`, `phone`, `status` FROM  `tb_users`  WHERE `status`=1 and   id=? LIMIT 1"
	result := u.Raw(sql, userId).First(u)
	if result.Error == nil {
		return u, nil
	} else {
		return nil, result.Error
	}
}

// 查询数据之前统计条数
func (u *AdminUsersModel) counts(userName string) (counts int64) {
	sql := "SELECT  count(*) as counts  FROM  tb_users  WHERE status=1 and   user_name like ?"
	if res := u.Raw(sql, "%"+userName+"%").First(&counts); res.Error != nil {
		variable.ZapLog.Error("UsersModel - counts 查询数据条数出错", zap.Error(res.Error))
	}
	return counts
}

// 查询（根据关键词模糊查询）
func (u *AdminUsersModel) Show(userName string, limitStart, limitItems int) (counts int64, temp []AdminUsersModel) {
	if counts = u.counts(userName); counts > 0 {
		sql := "SELECT  `id`, `user_name`, `real_name`, `phone`, `status`  FROM  `tb_users`  WHERE `status`=1 and   user_name like ? LIMIT ?,?"
		if res := u.Raw(sql, "%"+userName+"%", limitStart, limitItems).Find(&temp); res.RowsAffected > 0 {
			return counts, temp
		}
	}
	return 0, nil
}

//新增
func (u *AdminUsersModel) Store(userName string, pass string, realName string, phone string, remark string) bool {
	sql := "INSERT  INTO tb_users(user_name,pass,real_name,phone,remark) SELECT ?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  user_name=?)"
	if u.Exec(sql, userName, pass, realName, phone, remark, userName).RowsAffected > 0 {
		return true
	}
	return false
}

//更新
func (u *AdminUsersModel) Update(id int, userName string, pass string, realName string, phone string, remark string, clientIp string) bool {
	sql := "update tb_users set user_name=?,pass=?,real_name=?,phone=?,remark=?  WHERE status=1 AND id=?"
	if u.Exec(sql, userName, pass, realName, phone, remark, id).RowsAffected >= 0 {
		if u.OauthResetToken(id, pass, clientIp) {
			return true
		}
	}
	return false
}

//删除用户以及关联的token记录
func (u *AdminUsersModel) Destroy(id int) bool {
	if u.Delete(u, id).Error == nil {
		if u.OauthDestroyToken(id) {
			return true
		}
	}
	return false
}
