package main

import (
	"github.com/go-andiamo/chioas"
	"github.com/go-chi/chi/v5"
)

type api struct {
	chioas.Definition
}

func (a *api) SetupRoutes(r chi.Router) error {
	return a.Definition.SetupRoutes(r, a)
}

var userStoreApi = &api{
	Definition: apiDef,
}

var apiDef = chioas.Definition{
	DocOptions: chioas.DocOptions{
		ServeDocs: true,
		CheckRefs: true,
		AsJson:    true,
	},
	Info: chioas.Info{
		Title:   "Chi Fun",
		Version: "0.1.0",
	},
	Tags: chioas.Tags{
		{
			Name: "user",
		},
	},
	Paths: chioas.Paths{
		"/users": userPaths,
	},
	Components: &components,
}

var components = chioas.Components{
	Schemas: chioas.Schemas{
		(&chioas.Schema{
			Name:        "User",
			Description: "User schema description",
			Comment:     chioas.SourceComment(),
		}).Must(User{
			Id:        4,
			FirstName: "Larry",
			LastName:  "Smith",
		}),
		(&chioas.Schema{
			Name:        "Category",
			Description: "Category schema description",
			Comment:     chioas.SourceComment(),
		}).Must(UserCategory{
			Id:   9,
			Name: "Student",
		}),
		{
			Name:    "Status",
			Type:    "string",
			Default: "available",
			Enum:    []any{"available", "pending", "sold"},
		},
		(&chioas.Schema{
			Name: "UserUpdate",
		}).Must(UserUpdate{
			FirstName: "Larry",
			LastName:  "Smith",
		}),
		(&chioas.Schema{
			Name: "UserList",
		}).Must(UserList{
			Cursor: "some cursor",
			Items:  []User{{Id: 6, FirstName: "Larrs", LastName: "Smith"}},
		}),
	},
}
