package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-andiamo/chioas"
)

type UserCreate struct {
	FirstName string       `json:"first_name" oas:"description:The first name of the user,required,example"`
	LastName  string       `json:"last_name" oas:"description:The last name of the user,required,example"`
	Category  UserCategory `json:"category" oas:"$ref:Category,required"`
}

type UserUpdate struct {
	FirstName string       `json:"first_name" oas:"description:The first name of the user,example"`
	LastName  string       `json:"last_name" oas:"description:The last name of the user,example"`
	Category  UserCategory `json:"category" oas:"$ref:Category"`
}
type User struct {
	Id        int          `json:"id" oas:"description:The id of the user,required,example"`
	FirstName string       `json:"first_name" oas:"description:The first name of the user,required,example"`
	LastName  string       `json:"last_name" oas:"description:The last name of the user,required,example"`
	Category  UserCategory `json:"category" oas:"$ref:Category"`
}

type UserCategory struct {
	Id   int    `json:"id" oas:"description:The id of the category,example"`
	Name string `json:"name" oas:"description: The name of the category,example,# this is a comment in the spec"`
}

type UserList struct {
	Cursor string `json:"cursor" oas:"required,example"`
	Items  []User `json:"items" oas:"$ref:'User',required,example"`
}

var userPaths = chioas.Path{
	Tag: "user",
	Methods: map[string]chioas.Method{
		http.MethodGet: {
			Comment:     "Some comment about getting users",
			Handler:     (*api).GetUsers,
			Description: "Search the collection of all users",
			OperationId: "search",
			QueryParams: []chioas.QueryParam{
				{
					Name:        "status",
					Description: "some query param description",
					SchemaRef:   "Status",
				},
			},
			Responses: map[int]chioas.Response{
				http.StatusOK: {
					IsArray:   true,
					SchemaRef: "UserList",
				},
			},
		},
		http.MethodPost: {
			Comment:     "SOme comment about creating a User",
			Handler:     (*api).GetUser,
			Description: "Create a new User",
			OperationId: "create",
			Request: &chioas.Request{
				Description: "User to be added to the store",
				Schema:      UserCreate{},
			},
			Responses: map[int]chioas.Response{
				http.StatusCreated: {
					SchemaRef: "User",
				},
			},
		},
	},
	Paths: map[string]chioas.Path{
		"/{userId}": {
			PathParams: map[string]chioas.PathParam{
				"userId": {
					Description: "id of the User",
				},
			},
			Methods: chioas.Methods{
				http.MethodGet: {
					Comment:     "some path method comment for retrieving a User",
					Handler:     (*api).GetUser,
					Summary:     "retrieve",
					Description: "some description path method for retrieving a User",
					OperationId: "retrieve",
					Responses: chioas.Responses{
						http.StatusOK: {
							SchemaRef: "User",
						},
					},
				},
				http.MethodPatch: {
					Comment:     "some path method comment for updating a User",
					Handler:     (*api).UpdateUser,
					Summary:     "update",
					Description: "some description path method for updating a User",
					OperationId: "update",
					Request: &chioas.Request{
						Description: "Update an existent user in the store",
						SchemaRef:   "UserUpdate",
					},
					Responses: chioas.Responses{
						http.StatusOK: {
							SchemaRef: "User",
						},
					},
				},
			},
		},
	},
}

func (d *api) GetUsers(writer http.ResponseWriter, request *http.Request) {
	res := map[string]any{
		"hello": "you listed/queried users",
	}
	enc := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	_ = enc.Encode(res)
}

func (d *api) PostUsers(writer http.ResponseWriter, request *http.Request) {
	res := map[string]any{
		"hello": "you added a user",
	}
	enc := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	_ = enc.Encode(res)
}

func (d *api) GetUser(writer http.ResponseWriter, request *http.Request) {
	res := map[string]any{
		"hello": "you retrieved a user",
	}
	enc := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	_ = enc.Encode(res)
}

func (d *api) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	res := map[string]any{
		"hello": "you updated a user",
	}
	enc := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	_ = enc.Encode(res)
}
