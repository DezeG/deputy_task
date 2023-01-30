package smallRepository

import "deputy/models"

var Roles = []models.Role{
	{
		1,
		"System Administrator",
		0,
	},
	{
		2,
		"Location Manager",
		1,
	},
	{
		3,
		"Supervisor",
		2,
	},
	{
		4,
		"Employee",
		3,
	},
	{
		5,
		"Trainer",
		3,
	},
}
