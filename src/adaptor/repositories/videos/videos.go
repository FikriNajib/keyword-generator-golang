package videos

import (
	"context"
	"errors"
	"go.elastic.co/apm/v2"
	"gorm.io/gorm"
	ar "keyword-generator/src/application/repositories"
	"keyword-generator/src/infrastructure"
	"time"
)

type Videos struct {
	ID                      int       `gorm:"column:id;autoIncrement"`
	ContestantID            int       `gorm:"column:contestant_id"`
	CompetitionID           int       `gorm:"column:competition_id"`
	CommentForVideoID       int       `gorm:"column:comment_for_video_id"`
	CategoryID              int       `gorm:"column:category_id"`
	VideoID                 string    `gorm:"column:video_id"`
	VideoTemplateID         int       `gorm:"column:video_template_id"`
	SongTemplateID          int       `gorm:"column:song_template_id"`
	ChallengeID             int       `gorm:"column:challenge_id"`
	VideoTitle              string    `gorm:"column:video_title"`
	Thumbnail               string    `gorm:"column:thumbnail"`
	VideoSource             string    `gorm:"column:video_source"`
	OriginalSource          string    `gorm:"column:original_source"`
	Duration                int       `gorm:"column:duration"`
	Platform                string    `gorm:"column:platform"`
	UploadWith              string    `gorm:"column:upload_with"`
	SummaryVideo            string    `gorm:"column:summary_video"`
	Gender                  string    `gorm:"column:gender"`
	FormatGroup             string    `gorm:"column:format_group"`
	KualitasVideo           string    `gorm:"column:kualitas_video"`
	KeteranganKualitasVideo string    `gorm:"column:keterangan_kualitas_video"`
	GenreVideoID            int       `gorm:"column:genre_video_id"`
	SubgenreVideoID         int       `gorm:"column:subgenre_video_id"`
	UserVote                int       `gorm:"column:user_vote"`
	JudgeVote               int       `gorm:"column:judge_vote"`
	VotePerVideo            int       `gorm:"column:vote_per_video"`
	Shared                  int       `gorm:"column:shared"`
	Downloaded              int       `gorm:"column:downloaded"`
	Views                   int       `gorm:"column:views"`
	Pinned                  int       `gorm:"column:pinned"`
	SortingPinned           int       `gorm:"column:sorting_pinned"`
	NotesPinned             string    `gorm:"column:notes_pinned"`
	IsComment               bool      `gorm:"column:is_comment"`
	VideoStatus             string    `gorm:"column:video_status"`
	Status                  string    `gorm:"column:status"`
	FailedMessage           string    `gorm:"column:failed_message"`
	MaxRetryUpload          int       `gorm:"column:max_retry_upload"`
	CreateAt                time.Time `gorm:"column:create_at"`
	UpdateAt                time.Time `gorm:"column:update_at"`
	TusdFilename            string    `gorm:"tusd_filename"`
	UpdatedBy               string    `gorm:"column:updated_by"`
}

type repository struct {
	DbRepository *infrastructure.DBRepository
}

func NewVideosRepository(db *infrastructure.DBRepository) ar.VideosRepository {
	return &repository{DbRepository: db}
}

func (r *repository) GetKeyword(ctx context.Context, contentID []int) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, "src/adaptor/repositories/videos/videos.go", "GetKeyword")
	defer span.End()
	var hashtagList []string
	sql := `select distinct k.name FROM videos_has_keyword vhk inner join keyword k on vhk.keyword_id = k.id where vhk.video_id in ?`

	err := r.DbRepository.FindAll(r.DbRepository.DBMETUBE, ctx, &hashtagList, sql, contentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hashtagList, nil
		}
		apm.CaptureError(ctx, err).Send()
		return hashtagList, err
	}

	return hashtagList, err
}
