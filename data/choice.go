package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ChoiceModel is the main subelement of the page
type ChoiceModel struct {
	ID                 uint   `json:"id" gorm:"primary_key"`
	Name               string `json:"name"`
	Text               string `json:"text"`
	ParentStoryID      uint   `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	DefaultNextStoryID uint   `json:"nextStoryID" 		gorm:"column:default_next_story_id"`
}

// DTO methods

func (choice ChoiceModel) ToDto() gin.H {
	ret := make(gin.H)
	ret["id"] = choice.ID
	ret["name"] = choice.Name
	ret["text"] = choice.Text
	ret["parentStoryID"] = choice.ParentStoryID
	ret["nextStoryID"] = choice.DefaultNextStoryID

	return ret
}

// Useful methods

func (choice *ChoiceModel) GetNextStory() StoryModel {
	var story StoryModel
	DB.First(&story, choice.DefaultNextStoryID)
	return story
}

func (choice *ChoiceModel) GetAttributeStacks() []AttributeStack {
	var stacks []AttributeStack
	DB.Where("owner_id = ? and type = ?", choice.ID, "choice").Find(&stacks)
	return stacks
}

// CRUD methods

func (choice *ChoiceModel) Save() error {
	DB.Save(choice)
	return nil
}

func (choice *ChoiceModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(choice, id)
	if choice.ID != id {
		return errors.New("No choice found with the given ID")
	}

	return nil
}

func (choice *ChoiceModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(choice).Update(fieldName, fieldValue)

	return nil
}

func (choice *ChoiceModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(choice).Update(name, value)
	}

	return nil
}

func (choice *ChoiceModel) Delete() error {
	DB.Delete(choice)
	return nil
}
