package service

import (
	"goTut/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}
	//Setup Expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)
	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)
	post := entity.Post{Title: "A", Text: "B"}
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "Title is empty", err.Error())
}
