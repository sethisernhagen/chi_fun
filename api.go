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
		Title:       "Chi Fun - OpenAPI 3.0",
		Description: `this is some definition description`,
		Version:     "0.1.0",
	},
	Tags: chioas.Tags{
		{
			Name:        "user",
			Description: "definition user tag description",
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
		}).Must(UserResponse{
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
	},
}
