package storage

import "github.com/wesovilabs/beyond-examples/smartcache/model"

var countries = []*model.Country{
	{
		ID:   "ESP",
		Name: "Spain",
	},
	{
		ID:   "IR",
		Name: "Ireland",
	},
}

var cities = []*model.City{
	{
		ID:   "MD",
		Name: "Madrid",
		Country: "ESP",
	},
	{
		ID:   "ML",
		Name: "Málaga",
		Country: "ESP",
	},
	{
		ID:   "VL",
		Name: "Valencia",
		Country: "ESP",
	},
	{
		ID:   "DUB",
		Name: "Dublin",
		Country: "IR",
	},
	{
		ID:   "GAL",
		Name: "Galway",
		Country: "IR",
	},
}
