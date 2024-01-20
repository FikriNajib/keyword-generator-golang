package story

import (
	"context"
	"errors"
	"go.elastic.co/apm/v2"
	"gorm.io/gorm"
	ar "keyword-generator/src/application/repositories"
	"keyword-generator/src/infrastructure"
	"time"
)

type Story struct {
	ID             int       `gorm:"column:id"`
	Type           string    `gorm:"column:type"`
	ProgramID      int       `gorm:"column:program_id"`
	Title          string    `gorm:"column:title"`
	Description    string    `gorm:"column:description"`
	BGType         string    `gorm:"column:bg_type"`
	ImageType      string    `gorm:"column:image_type"`
	Image          string    `gorm:"column:image"`
	Image1         string    `gorm:"column:image1"`
	ReleaseDate    time.Time `gorm:"column:release_date"`
	ExpiredDate    time.Time `gorm:"column:expired_date"`
	Status         string    `gorm:"column:status"`
	LinkVideo      string    `gorm:"column:link_video"`
	Link           string    `gorm:"column:link"`
	ColorCode      string    `gorm:"column:color_code"`
	Sorting        int       `gorm:"column:sorting"`
	SwipeType      string    `gorm:"column:swipe_type"`
	Permalink      string    `gorm:"column:permalink"`
	Deeplink       string    `gorm:"column:deeplink"`
	SwipeValue     string    `gorm:"column:swipe_value"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	CreatedBy      int       `gorm:"column:created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	UpdatedBy      int       `gorm:"column:updated_by"`
	CatchupDate    time.Time `gorm:"column:catchup_date"`
	CatchupChannel string    `gorm:"column:catchup_channel"`
	MandatoryLogin bool      `gorm:"column:mandatory_login"`
	Keywords       string    `gorm:"column:keywords"`
	Duration       int       `gorm:"column:duration"`
	Thumbnail      string    `gorm:"column:thumbnail"`
}

type repository struct {
	DbRepository *infrastructure.DBRepository
}

func NewStoryRepository(db *infrastructure.DBRepository) ar.StoryRepository {
	return &repository{DbRepository: db}
}

func (r *repository) GetHashtag(ctx context.Context, contentID []int) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, "src/adaptor/repositories/story/story.go", "GetHashtag")
	defer span.End()
	var keywords []string
	sql := `select distinct t.name FROM story_has_tags sht inner join tag t on sht.tag_id  = t.id where sht.story_id  in ?`

	err := r.DbRepository.FindAll(r.DbRepository.DBFTA, ctx, &keywords, sql, contentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return keywords, nil
		}
		apm.CaptureError(ctx, err).Send()
		return keywords, err
	}

	return keywords, err
}
