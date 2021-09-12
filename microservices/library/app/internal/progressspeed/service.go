package progressspeed

import (
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alexandr-io/backend/library/data"
)

// Serv instance of book progress service
var Serv *Service

// Service is the struct containing database repository needed for book progress methods of the interface
type Service struct {
	repo Repository
}

// NewService create and set instance of Service
func NewService(repo Repository) *Service {
	Serv = &Service{repo: repo}
	return Serv
}

// UpsertProgressSpeed upsert new reading speed blocks
func (s *Service) UpsertProgressSpeed(userID primitive.ObjectID, language string, wordNumber int) error {
	now := time.Now()

	currentProgressSpeed, err := s.repo.Read(userID, language)
	if err != nil && err.(*fiber.Error).Code == fiber.StatusNotFound {
		return s.repo.Upsert(&data.ProgressSpeed{
			UserID:     userID,
			Language:   language,
			LastUpdate: now,
		})
	}

	var progressBlock []data.ProgressHistory
	readingTime10s := now.Sub(currentProgressSpeed.LastUpdate).Seconds() / math.Floor(float64(wordNumber)/10.0)
	for i := 0; i < wordNumber/10; i++ {
		progressBlock = append(progressBlock, data.ProgressHistory{
			WordNumber: 10,
			Time:       readingTime10s,
		})
	}

	progressSpeed := data.ProgressSpeed{
		UserID:     userID,
		Language:   language,
		LastUpdate: now,
		History:    append(currentProgressSpeed.History, progressBlock...),
	}

	if err = s.repo.Upsert(&progressSpeed); err != nil {
		return err
	}

	return nil
}

// ReadReadingSpeed read a reading speed for a number of word in a specific language
func (s *Service) ReadReadingSpeed(userID primitive.ObjectID, language string, wordNumber int) (*data.ReadingSpeed, error) {
	currentProgressSpeed, err := s.repo.Read(userID, language)
	if err != nil {
		return nil, err
	}

	var i = 0
	var total float64 = 0
	for _, block := range currentProgressSpeed.History {
		total += block.Time
		i++
	}

	return &data.ReadingSpeed{Speed: total / float64(i) / 10 * float64(wordNumber)}, nil
}
