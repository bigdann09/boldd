package services

import (
	"log"

	"github.com/boldd/internal/config"
	"github.com/boldd/internal/infrastructure/auth/jwt"
	"github.com/boldd/internal/infrastructure/mail"
	"github.com/boldd/internal/infrastructure/persistence"
	"github.com/boldd/internal/infrastructure/persistence/redis"
	seeder "github.com/boldd/internal/infrastructure/persistence/seeders"
	"github.com/boldd/internal/infrastructure/storage"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/boldd/pkgs/logger"
	goredis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	DB      *gorm.DB
	Logger  *zap.Logger
	Redis   *goredis.Client
	Token   jwt.ITokenService
	Storage storage.ICloudinary
	Mail    mail.IMail
}

func NewServices(cfg *config.Config) *Service {
	// register database
	db, err := persistence.NewDB(&cfg.DatabaseConfig)
	if err != nil {
		log.Println("could not connect to database")
		panic(err)
	}

	// register custom validators
	validator := validator.NewValidator(db)
	validator.RegisterValidators()

	// register logger
	logger := logger.NewLogger(cfg.Environment)
	defer logger.Sync()

	// register redis
	redis := redis.NewRedisClient(&cfg.RedisConfig)

	// register token service
	jwt := jwt.NewTokenService(&cfg.JWTConfig)

	// register cloudinary
	cloudinary, err := storage.NewCloudinary(&cfg.CloudinaryConfig)
	if err != nil {
		log.Println("could not initialize cloudinary")
		panic(err)
	}

	// register seeders
	seeder := seeder.NewSeeder(db, logger)
	seeder.Run()

	// register mail
	mail := mail.NewMail(
		cfg.MailConfig.From,
		cfg.MailConfig.Username,
		cfg.MailConfig.Password,
		cfg.MailConfig.Host,
		cfg.MailConfig.Port,
	)

	return &Service{
		DB:      db,
		Logger:  logger,
		Redis:   redis,
		Token:   jwt,
		Mail:    mail,
		Storage: cloudinary,
	}
}
