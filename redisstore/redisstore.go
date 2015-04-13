package redisstore

import (
	"log"
	"github.com/garyburd/redigo/redis"
)

// TODO: Store can use twice less space as even (expect 2) numbers are not prime. Needs wrapper?
// TODO: Store (not Store type) must be interface

type Store struct {
	key string
	conn redis.Conn
}

func (s *Store) Close() {
	s.conn.Close()
}

func logErrors(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func New() *Store {
	c, err := redis.Dial("tcp", ":6379")
	logErrors("error connection to database, %v", err)

	store := &Store{
		key: "fe:primes",
		conn: c,
	}

	return store
}

func (rs *Store) Reset() {
	err := rs.conn.Send("DEL", rs.key)
	logErrors("Reset(), %v", err)
}

func (rs *Store) Set(pos int, val int) {
	err := rs.conn.Send("SETBIT", rs.key, pos, val)
	logErrors("Set(), %v", err)
}

func (rs *Store) Get(pos int) int {
	val, err := redis.Int(rs.conn.Do("GETBIT", rs.key, pos))
	logErrors("Get(), %v", err)

	return val
}

func (rs *Store) Add(pos int) {
	rs.Set(pos, 1)
}

func (rs *Store) Invert(pos int) {
	val := rs.Get(pos)
	rs.Set(pos, val ^ 1)
}

func (rs *Store) FindFrom(offset int) int {
	// BITPOS [start] is byte not bit - http://redis.io/commands/bitpos
	var startByte int = offset / 8
	val, err := redis.Int(rs.conn.Do("BITPOS", rs.key, 1, startByte))
	logErrors("Find(), %v", err)
	if(val < offset && -1 != val) {
		nextOffset := (startByte + 1)*8
		val = rs.firstInRange_Slow(offset, nextOffset - 1)
		if(-1 == val) {
			val = rs.FindFrom(nextOffset)
		}
	}

	return val
}

// Makes calls to Redis for every position in the range. Consider to use FindFrom() for more optimal search.
func (rs *Store) firstInRange_Slow(start int, end int) int {
	for at := start; at<= end; at++ {
		if rs.Get(at) == 1  {
			return at
		}
	}
	return -1
}
