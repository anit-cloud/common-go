package entity

type User struct {
	UserName string
	Name     string
	Role     string
}

func CreateUserFromSlide(raw interface{}) *User {
	result := &User{
		UserName: raw.(map[string]interface{})["userName"].(string),
		Name:     raw.(map[string]interface{})["name"].(string),
		Role:     raw.(map[string]interface{})["role"].(string),
	}
	return result
}
