package sql

import (
	log "azura-test/src/business/usecases/log"
	"database/sql"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

// A SQLHandler belong to the infrastructure layer.
type SQLHandler struct {
	Conn *sqlx.DB
}

type Config struct {
	Driver string
	// Leder  ConnConfig
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

// type ConnConfig struct {
// }

type Interface interface {
	Leader() Command
}

type sqlDB struct {
	endOnce *sync.Once
	leader  Command
	cfg     Config
	log     log.Logger
}

func Init(cfg Config, log log.Logger) Interface {
	sql := &sqlDB{
		endOnce: &sync.Once{},
		cfg:     cfg,
		log:     log,
	}
	sql.initDB()
	return sql
}

func (s *sqlDB) Leader() Command {
	return s.leader
}

func (s *sqlDB) initDB() {
	db, err := s.connect(true)

	if err != nil {
		s.log.LogError("[FATAL] can not connect to db %s leader: %s on port %d, with error %s", s.cfg.DB, s.cfg.Host, s.cfg.Port, err)
	}

	s.log.LogAccess("SQL: [LEADER] driver=%s db=%s @%s:%v", s.cfg.Driver, s.cfg.DB, s.cfg.Host, s.cfg.Port)
	s.leader = initCommand(db)
}

func (s *sqlDB) connect(toLeader bool) (*sqlx.DB, error) {
	conf := s.cfg

	uri, err := s.getURI(conf)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(s.cfg.Driver, uri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	sqlxDB := sqlx.NewDb(db, s.cfg.Driver)
	return sqlxDB, nil
}

func (s *sqlDB) getURI(conf Config) (string, error) {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DB), nil
}
