package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/dalecosta1/sinaloa-api/entity"
	"github.com/dalecosta1/sinaloa-api/repository"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	videoRepository := repository.NewVideoRepository()
	defer videoRepository.CloseDB()

	videoService := New(videoRepository)

	videoService.Save(getVideo())

	videos := videoService.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
