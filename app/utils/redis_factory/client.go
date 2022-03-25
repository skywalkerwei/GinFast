package redis_factory

import (
	"encoding/json"
	"errors"
	"ginfast/app/core/event_manage"
	"ginfast/app/global/my_errors"
	"ginfast/app/global/variable"
	"ginfast/app/utils/yml_config"
	"ginfast/app/utils/yml_config/ymlconfig_interf"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var redisPool *redis.Pool
var configYml ymlconfig_interf.YmlConfigInterf

// 处于程序底层的包，init 初始化的代码段的执行会优先于上层代码，因此这里读取配置项不能使用全局配置项变量
func init() {
	configYml = yml_config.CreateYamlFactory()
	redisPool = initRedisClientPool()
}
func initRedisClientPool() *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     configYml.GetInt("Redis.MaxIdle"),                        //最大空闲数
		MaxActive:   configYml.GetInt("Redis.MaxActive"),                      //最大活跃数
		IdleTimeout: configYml.GetDuration("Redis.IdleTimeout") * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", configYml.GetString("Redis.Host")+":"+configYml.GetString("Redis.Port"))
			if err != nil {
				variable.ZapLog.Error(my_errors.ErrorsRedisInitConnFail + err.Error())
				return nil, err
			}
			auth := configYml.GetString("Redis.Auth") //通过配置项设置redis密码
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
					variable.ZapLog.Error(my_errors.ErrorsRedisAuthFail + err.Error())
				}
			}
			_, _ = conn.Do("select", configYml.GetInt("Redis.IndexDb"))
			return conn, err
		},
	}
	// 将redis的关闭事件，注册在全局事件统一管理器，由程序退出时统一销毁
	event_manage.CreateEventManageFactory().Set(variable.EventDestroyPrefix+"Redis", func(args ...interface{}) {
		_ = redisPool.Close()
	})
	return redisPool
}

//  从连接池获取一个redis连接
func GetOneRedisClient() *RedisClient {
	maxRetryTimes := configYml.GetInt("Redis.ConnFailRetryTimes")
	var oneConn redis.Conn
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		if oneConn.Err() != nil {
			//variable.ZapLog.Error("Redis：网络中断,开始重连进行中..." , zap.Error(oneConn.Err()))
			if i == maxRetryTimes {
				variable.ZapLog.Error(my_errors.ErrorsRedisGetConnFail, zap.Error(oneConn.Err()))
				return nil
			}
			//如果出现网络短暂的抖动，短暂休眠后，支持自动重连
			time.Sleep(time.Second * configYml.GetDuration("Redis.ReConnectInterval"))
		} else {
			break
		}
	}
	return &RedisClient{oneConn}
}

// 定义一个redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// 为redis-go 客户端封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args...)
}

// 释放连接到连接池
func (r *RedisClient) ReleaseOneRedisClient() {
	_ = r.client.Close()
}

//  封装几个数据类型转换的函数

//bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

//string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

//string map 类型转换
func (r *RedisClient) StringMap(reply interface{}, err error) (map[string]string, error) {
	return redis.StringMap(reply, err)
}

//strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

//Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

//int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

//int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

//int map 类型转换
func (r *RedisClient) IntMap(reply interface{}, err error) (map[string]int, error) {
	return redis.IntMap(reply, err)
}

//Int64Map 类型转换
func (r *RedisClient) Int64Map(reply interface{}, err error) (map[string]int64, error) {
	return redis.Int64Map(reply, err)
}

//int64s 类型转换
func (r *RedisClient) Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

//uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

//Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

// 以上封装了很多最常见类型转换函数，其他您可以参考以上格式自行封装
func (r *RedisClient) Set(key string, val interface{}, expire int64) error {
	value, err := r.encode(val)
	if err != nil {
		return err
	}
	if expire > 0 {
		_, err := r.Execute("SETEX", r.getKey(key), expire, value)
		return err
	}
	_, err = r.Execute("SET", r.getKey(key), value)
	return err
}

// Exists 检查键是否存在
func (r *RedisClient) Exists(key string) (bool, error) {
	isExists, errExists := r.Execute("EXISTS", r.getKey(key))
	if errExists != nil {
		return false, errExists
	}
	if isExists.(int64) > 0 {
		return true, nil
	} else {
		return false, errors.New("key 不存在")
	}
}

//Del 删除键
func (r *RedisClient) Del(key string) error {
	_, err := r.Execute("DEL", r.getKey(key))
	return err
}

// Flush 清空当前数据库中的所有 key，慎用！
func (r *RedisClient) Flush() error {
	_, err := r.Execute("FLUSHDB")
	return err
}

// TTL 以秒为单位。当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1
func (r *RedisClient) TTL(key string) (ttl int64, err error) {
	return r.Int64(r.Execute("TTL", r.getKey(key)))
}

// Expire 设置键过期时间，expire的单位为秒
func (r *RedisClient) Expire(key string, expire int64) error {
	_, err := r.Bool(r.Execute("EXPIRE", r.getKey(key), expire))
	return err
}

// Incr 将 key 中储存的数字值增一
func (r *RedisClient) Incr(key string) (val int64, err error) {
	return r.Int64(r.Execute("INCR", r.getKey(key)))
}

// IncrBy 将 key 所储存的值加上给定的增量值（increment）。
func (r *RedisClient) IncrBy(key string, amount int64) (val int64, err error) {
	return r.Int64(r.Execute("INCRBY", r.getKey(key), amount))
}

// Decr 将 key 中储存的数字值减一。
func (r *RedisClient) Decr(key string) (val int64, err error) {
	return r.Int64(r.Execute("DECR", r.getKey(key)))
}

// DecrBy key 所储存的值减去给定的减量值（decrement）。
func (r *RedisClient) DecrBy(key string, amount int64) (val int64, err error) {
	return r.Int64(r.Execute("DECRBY", r.getKey(key), amount))
}

// Get 获取键值。一般不直接使用该值，而是配合下面的工具类方法获取具体类型的值，或者直接使用github.com/gomodule/redigo/redis包的工具方法。
func (r *RedisClient) Get(key string) (interface{}, error) {
	return r.Execute("GET", r.getKey(key))
}

// GetString 获取string类型的键值
func (r *RedisClient) GetString(key string) (string, error) {
	return r.String(r.Get(key))
}

// GetInt 获取int类型的键值
func (r *RedisClient) GetInt(key string) (int, error) {
	return r.Int(r.Get(key))
}

// GetInt64 获取int64类型的键值
func (r *RedisClient) GetInt64(key string) (int64, error) {
	return r.Int64(r.Get(key))
}

// GetBool 获取bool类型的键值
func (r *RedisClient) GetBool(key string) (bool, error) {
	return r.Bool(r.Get(key))
}

// GetObject 获取非基本类型stuct的键值。在实现上，使用json的Marshal和Unmarshal做序列化存取。
func (r *RedisClient) GetObject(key string, val interface{}) error {
	reply, err := r.Get(key)
	return r.decode(reply, err, val)
}

// GetMap 返回map
func (r *RedisClient) GetMap(key string) (data map[string]interface{}, err error) {
	reply, err := r.Get(key)
	err = r.decode(reply, err, &data)
	return data, err
}

// getKey 将健名加上指定的前缀。
func (r *RedisClient) getKey(key string) string {
	return configYml.GetString("Redis.Prefix") + key
}

// encode 序列化要保存的值
func (r *RedisClient) encode(val interface{}) (interface{}, error) {
	var value interface{}
	switch v := val.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		value = v
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		value = string(b)
	}
	return value, nil
}

// decode 反序列化保存的struct对象
func (r *RedisClient) decode(reply interface{}, err error, val interface{}) error {
	str, err := r.String(reply, err)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), val)
}
