package command

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"
	"context"

	"github.com/SUT-technology/judgino/pkg/slogger"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gopkg.in/yaml.v3"
)


func Run() error {
	var username, password string

	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/development.yaml", "Configuration File")
	flag.StringVar(&username, "username", "", "Admin's username")
	flag.StringVar(&password, "password", "", "Admin's password")
	flag.Parse()

	if username == "" || password == "" {
		return fmt.Errorf("username and password must be provided")
	}
	c, err := Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	logger := slogger.NewJSONLogger(c.Logger.Level, os.Stdout)
	slog.SetDefault(logger)




	db, err := NewPool(c.DB)
	if err != nil {
		fmt.Printf("Error initializing database connection: %v\n", err)
	}

	ctx := context.Background()
	err = db.Query(ctx, func(r *Repo) error {
		user, err := r.Users.CreateAdmin(ctx, username, password)
		if err != nil {
			return fmt.Errorf("creating admin: %w", err)
		}
		slog.Debug("created admin", slog.Int("id", int(user.ID)))
		return nil
	})
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	slog.Debug("initial db")


	return nil
}

type Config struct {
	DB     DB     `yaml:"db"`
	Logger Logger `yaml:"logger"`
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
}

type RateLimiter struct {
	Enabled bool          `yaml:"enabled"`
	Rate    int           `yaml:"rate"`
	Burst   int           `yaml:"burst"`
	Expires time.Duration `yaml:"expires"`
}
type Server struct {
	Port        string      `yaml:port`
	SecretKey   string      `yaml:"secret_key" validate:"required"`
	Logger      bool        `yaml:"logger" validate:"required"`
	RateLimiter RateLimiter `yaml:"rate-limiter"`
	Addr        string      `yaml:"addr"`
}

type DB struct {
	Port     string `yaml:"port" validate:"required"`
	DBName   string `yaml:"db_name" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Username string `yaml:"username" validate:"required"`
}

type Logger struct {
	Level string `yaml:"level" validate:"required,oneof=trace debug info warn error fatal"`
}

func (c Config) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

func Load(path string) (Config, error) {
	fmt.Println("Loading config file:", path)
	f, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading file: %w", err)
	}

	c, err := Parse(f)
	if err != nil {
		return Config{}, fmt.Errorf("parsing configs: %w", err)
	}

	if err = c.Validate(); err != nil {
		return Config{}, fmt.Errorf("validating configs: %w", err)
	}

	return c, nil
}

// Parse reads the yaml data into a Config struct. It does not perform any validations on the configurations themselves.
func Parse(data []byte) (Config, error) {
	c := Config{}
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return Config{}, fmt.Errorf("parsing yaml file: %w", err)
	}
	return c, nil
}

type GormQuerier struct {
	DB *gorm.DB
}

func (g *GormQuerier) Exec(query string, args ...interface{}) error {
	return g.DB.Exec(query, args...).Error
}

func (g *GormQuerier) Find(dest interface{}, query string, args ...interface{}) error {
	return g.DB.Raw(query, args...).Scan(dest).Error
}

func (g *GormQuerier) First(dest interface{}, query string, args ...interface{}) error {
	return g.DB.Raw(query, args...).First(dest).Error
}

func NewDB(db *gorm.DB) Tables {
	return Tables{
		Users: newUsersTable(db),
	}
}




type QueryFunc = func(r *Repo) error

type Querier interface {
	Exec(query string, args ...interface{}) error
	Find(dest interface{}, query string, args ...interface{}) error
	First(dest interface{}, query string, args ...interface{}) error
}


type Tables struct {
	Users UserRepository
}

type Repo struct {
	Tables
	Querier
}


type Pool struct {
	db *gorm.DB
}

// New opens a connection to a PostgreSQL database using GORM
func NewPool(cfg DB) (*Pool, error) {
	// PostgreSQL DSN (update placeholders with your actual database details)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open(%s): %w", dsn, err)
	}

	// err = db.AutoMigrate(&entity.User{}, &entity.Question{}, &entity.Submission{})
	// if err != nil {
	// 	log.Fatalf("Error migrating schema: %v", err)
	// }
	// fmt.Println("Connected to PostgreSQL using GORM!")
	return &Pool{db: db}, nil
}

// Query starts a transaction and executes the given function
func (p *Pool) Query(ctx context.Context, f QueryFunc) error {
	// Begin a transaction
	tx := p.db.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	r := &Repo{
		Querier: &GormQuerier{DB: tx},
		Tables:  NewDB(tx),
	}

	// Execute the function with the transaction
	err := f(r)
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			return fmt.Errorf("rollback: %w query: %w", rollbackErr, err)
		}
		return fmt.Errorf("query: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// Close is a placeholder for GORM cleanup (not strictly needed)
func (p *Pool) Close() error {
	fmt.Println("Closing database connection.")
	return nil
}

type UserRepository interface {
	CreateAdmin(ctx context.Context, username string, password string) (user User, error error)
}


type usersTable struct {
	db *gorm.DB
}

func newUsersTable(db *gorm.DB) usersTable {
	return usersTable{db: db}
}

type User struct {
    ID                   uint   `gorm:"primaryKey"`
    FirstName            string `gorm:"size:255;not null"`
    Email                *string `gorm:"size:255;unique"`
    Phone                *string `gorm:"size:11;unique"`
    Username             string  `gorm:"not null;unique"`
    Password             string `gorm:"size:255"`
    Role                 string `gorm:"size:255;not null"`
    CreatedQuestionsCount int64 `gorm:"not null"`
    SolvedQuestionsCount  int64 `gorm:"not null"`
}

func (c usersTable) CreateAdmin(ctx context.Context, username string, password string) (User, error) {
    var user User

    if err := c.db.Where("username = ?", username).First(&user).Error; err == nil {
        user.Role = "admin"
        if err := c.db.Save(&user).Error; err != nil {
            return User{}, fmt.Errorf("updating user role: %w", err)
        }
        return user, nil
    }

    newUser := User{
        Username: username,
        Password: password, // Make sure to hash the password before saving in production
        Role:     "admin",
		Email: nil,
		Phone: nil,
    }

    if err := c.db.Create(&newUser).Error; err != nil {
        return User{}, fmt.Errorf("creating new user: %w", err)
    }

    return newUser, nil
}