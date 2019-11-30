package cache

import (
	"errors"
	"github.com/colefan/sailfish/log"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	CMD_EXISTS   = "EXISTS"
	CMD_SET      = "SET"
	CMD_HGET     = "HGET"
	CMD_GET      = "GET"
	CMD_EXPIRE   = "EXPIRE"
	CMD_HSET     = "HSET"
	CMD_HMSET    = "HMSET"
	CMD_INCRBY   = "INCRBY"
	CMD_SADD     = "SADD"
	CMD_SREM     = "SREM"
	CMD_HINCRBY  = "HINCRBY"
	CMD_DEL      = "DEL"
	CMD_HDEL     = "HDEL"
	CMD_SMEMBERS = "SMEMBERS"
)

var ErrorRedisConnIsNull = errors.New("redis conn is nil")

// RedisCache   redis cache
type RedisCache struct {
	client *redis.Pool
	conf   *RedisConfig
	// redisConn redis.Conn
	// connected bool
}

func (rc *RedisCache) Init() {
	rc.client = &redis.Pool{
		MaxIdle:     rc.conf.MaxIdle,
		MaxActive:   rc.conf.MaxActive,
		IdleTimeout: time.Duration(rc.conf.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				rc.conf.Address,
				redis.DialDatabase(rc.conf.DbNum),
				redis.DialPassword(rc.conf.Password))

			if err != nil {
				log.Errorf("redis pool dial failed,error = %v", err)
				return nil, err
			}
			return c, nil
		},
	}

	// rc.redisConn, err = redis.Dial("tcp",
	// 	conf.GetCacheConfInst().Address,
	// 	redis.DialDatabase(conf.GetCacheConfInst().DbNum),
	// 	redis.DialPassword(conf.GetCacheConfInst().Password))
	// if err == nil {
	// 	rc.connected = true
	// } else {
	// 	rc.connected = false
	// }

}

func (rc *RedisCache) getConn() redis.Conn {
	conn := rc.client.Get()
	if conn == nil {
		log.Error("redischache getConn is nil")
	}
	return conn
}

func (rc *RedisCache) Ping() error {
	conn := rc.getConn()
	if conn == nil {
		return ErrorRedisConnIsNull
	}
	_, err := conn.Do("PING")
	if err != nil {
		log.Error("reids ping error")
	}
	return err
}

func NewRedisCache(redisConf *RedisConfig) *RedisCache {
	c := &RedisCache{
		conf: redisConf,
	}
	return c
}

func (rc *RedisCache) KeyExists(key string) bool {
	conn := rc.getConn()
	if conn == nil {
		return false
	}
	defer conn.Close()
	is_key_exits, err := redis.Bool(conn.Do(CMD_EXISTS, key))
	if err != nil {
		log.Errorf("rediscache KeyExists error :%v", err)
		return false
	} else {
		return is_key_exits
	}
}

func (rc *RedisCache) KeyExpired(key string, expireSecond int) {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return
	}
	defer conn.Close()
	conn.Do(CMD_EXPIRE, key, expireSecond)
}

func (rc *RedisCache) Put(key string, jsonValue string, expireSecond int) {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return
	}
	defer conn.Close()
	_, err := conn.Do(CMD_SET, key, jsonValue)
	if err != nil {
		log.Errorf("Put failed,key = %v,val = %v,error = %v", key, jsonValue, err)
		return
	}
	if expireSecond > 0 {
		conn.Do(CMD_EXPIRE, key, expireSecond)
	}
}

func (rc *RedisCache) Get(key string) string {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return ""
	}
	defer conn.Close()
	r, err := redis.String(conn.Do(CMD_GET, key))
	if err != nil {
		log.Errorf("redis.String get error :%v", err)
		return ""
	} else {
		return r
	}
}

func (rc *RedisCache) HashGetField(key string, field string) string {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return ""
	}
	defer conn.Close()
	r, err := redis.String(conn.Do(CMD_HGET, key, field))
	if err != nil && err != redis.ErrNil {
		log.Errorf("HashGetField key = %v,field = %v error: %v", key, field, err)
		return ""
	}
	return r
}

func (rc *RedisCache) HashPutField(key string, field string, val string, expireSecond int) {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return
	}
	defer conn.Close()
	_, err := conn.Do(CMD_HSET, key, field, val)
	if err != nil {
		log.Errorf("HashPutField failed, key = %v,field = %v,val = %v,error = %v", key, field, val, err)
		return
	}
	if expireSecond > 0 {
		conn.Do(CMD_EXPIRE, key, expireSecond)
	}
}

func (rc *RedisCache) HashPut(key string, kv map[string]string, expireSecond int) {
	conn := rc.getConn()
	if conn == nil {
		log.Error("conn is nil")
		return
	}
	defer conn.Close()
	params := make([]interface{}, 0, len(kv)*2+1)
	params = append(params, key)
	for k, v := range kv {
		params = append(params, k, v)
	}
	_, err := conn.Do(CMD_HMSET, params...)
	if err != nil {
		log.Errorf("HashPut failed, params = %v,error = %v", params, err)
		return
	}
	if expireSecond > 0 {
		conn.Do(CMD_EXPIRE, key, expireSecond)
	}

}

func (rc *RedisCache) Del(key string) {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return
	}
	defer conn.Close()
	_, err := conn.Do(CMD_DEL, key)
	if err != nil {
		log.Errorf("key = %v, Del error :%v", key, err)
	}
}

func (rc *RedisCache) Increment(key string, incrAmount int) int {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return 0
	}
	defer conn.Close()
	reply, err := redis.Int(conn.Do(CMD_INCRBY, key, incrAmount))
	if err != nil {
		log.Errorf("key = %v, incrAmount = %d,Increment error :%v", key, incrAmount, err)
	}
	return reply
}

func (rc *RedisCache) SetAddMembers(key string, members []interface{}, expireSecond int) {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return
	}
	defer conn.Close()
	params := make([]interface{}, 0, 1+len(members))
	params = append(params, key)
	params = append(params, members...)
	_, err := conn.Do(CMD_SADD, params...)
	if err != nil {
		log.Errorf("SetAddMembers failed,params = %v,error = %v", params, err)
		return
	}
	if expireSecond > 0 {
		conn.Do(CMD_EXPIRE, key, expireSecond)
	}
}

func (rc *RedisCache) SetDeleteMembers(key string, members []interface{}) {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return
	}
	defer conn.Close()
	params := make([]interface{}, 0, 1+len(members))
	params = append(params, key)
	params = append(params, members...)
	_, err := conn.Do(CMD_SREM, params...)
	if err != nil {
		log.Errorf("SetDeleteMembers failed,params  = %v, error :%v", params, err)
	}
}

func (rc *RedisCache) SetMemebersStringSlice(key string) []string {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return nil
	}
	defer conn.Close()
	val, err := redis.Strings(conn.Do(CMD_SMEMBERS, key))
	if err != nil {
		log.Errorf("SetMemebers failed,key = %v,error = %v", key, err)
	}
	// glog.GetLogger().Info("value = %v. type =%t", val, val)
	return val
}
func (rc *RedisCache) HashIncrement(key string, field string, inc int) {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return
	}
	defer conn.Close()
	_, err := conn.Do(CMD_HINCRBY, key, field, inc)
	if err != nil {
		log.Errorf("HashIncrement error :%v", err)
	}
}

func (rc *RedisCache) HashDelField(key string, field string) {
	conn := rc.getConn()
	if conn == nil {
		log.Errorf("conn is nil ")
		return
	}
	defer conn.Close()
	_, err := conn.Do(CMD_HDEL, key, field)
	if err != nil {
		log.Errorf("Hash Del field failed ,key = %v,field = %v,error = %v", key, field, err)
	}
}