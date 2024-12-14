package controllers

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title" form:"title"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
}

type ListMovies []Movie

var Data = ListMovies{
	{
		Id:          1,
		Title:       "Lovely Runner",
		Image:       "example/lovelyrunner.com",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Id:          2,
		Title:       "Love Next Door",
		Image:       "example/lovenextdoor.com",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Id:          3,
		Title:       "Uncontrollably Fond",
		Image:       "example/uncontrollablyfond.com",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Id:          4,
		Title:       "Hometown Cha-cha-cha",
		Image:       "example/hometown.com",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Id:          5,
		Title:       "Start Up",
		Image:       "example/startup.com",
		Description: "Lorem ipsum dolor sit amet",
	},
	{Id: 6, Title: "The Silent Journey", Image: "https://example.com/images/silent-journey.jpg", Description: "A gripping adventure about a lone traveler exploring an ancient world."},
	{Id: 7, Title: "Mystery of the Depths", Image: "https://example.com/images/mystery-depths.jpg", Description: "A thrilling underwater mystery set in the depths of the ocean."},
	{Id: 8, Title: "Lost City of Z", Image: "https://example.com/images/lost-city-z.jpg", Description: "An epic story of explorers seeking the lost city in the Amazon jungle."},
	{Id: 9, Title: "The Last Stand", Image: "https://example.com/images/last-stand.jpg", Description: "A group of soldiers make their final stand against overwhelming forces."},
	{Id: 10, Title: "Firestorm", Image: "https://example.com/images/firestorm.jpg", Description: "A dangerous wildfire threatens a small town and the people inside it."},
	{Id: 11, Title: "Into the Wild", Image: "https://example.com/images/into-wild.jpg", Description: "A young man’s journey into the wild to escape the chaos of society."},
	{Id: 12, Title: "Edge of Tomorrow", Image: "https://example.com/images/edge-tomorrow.jpg", Description: "A soldier gets caught in a time loop, constantly reliving his last battle."},
	{Id: 13, Title: "Kingdom of Shadows", Image: "https://example.com/images/kingdom-shadows.jpg", Description: "A medieval kingdom caught in a power struggle as darkness rises."},
	{Id: 14, Title: "The Forgotten Kingdom", Image: "https://example.com/images/forgotten-kingdom.jpg", Description: "An ancient civilization’s secrets are uncovered after centuries of silence."},
	{Id: 15, Title: "Winds of War", Image: "https://example.com/images/winds-war.jpg", Description: "A thrilling war drama about the rise of a powerful empire in the 1940s."},
	{Id: 16, Title: "In the Heat of the Night", Image: "https://example.com/images/heat-night.jpg", Description: "A crime detective works to solve a case amidst racial tensions in the South."},
	{Id: 17, Title: "The Midnight Escape", Image: "https://example.com/images/midnight-escape.jpg", Description: "A group of prisoners escape during the dead of night, with deadly consequences."},
	{Id: 18, Title: "Rogue’s Gallery", Image: "https://example.com/images/rogues-gallery.jpg", Description: "A band of outlaws come together for one last heist in a post-apocalyptic world."},
	{Id: 19, Title: "Nightfall", Image: "https://example.com/images/nightfall.jpg", Description: "A chilling tale of survival after an unexplained event plunges the world into darkness."},
	{Id: 20, Title: "Dawn of the Machines", Image: "https://example.com/images/dawn-machines.jpg", Description: "The first signs of artificial intelligence becoming self-aware lead to a global crisis."},
	{Id: 21, Title: "Rise of the Phoenix", Image: "https://example.com/images/rise-phoenix.jpg", Description: "A young hero discovers their destiny while the world stands on the brink of destruction."},
	{Id: 22, Title: "The Final Countdown", Image: "https://example.com/images/final-countdown.jpg", Description: "A time-traveling adventure where a team tries to prevent a catastrophic event."},
	{Id: 23, Title: "Shadows of the Past", Image: "https://example.com/images/shadows-past.jpg", Description: "A young woman uncovers dark family secrets that lead her on a path of danger."},
	{Id: 24, Title: "Beyond the Horizon", Image: "https://example.com/images/beyond-horizon.jpg", Description: "A visually stunning journey of discovery, traveling beyond the edge of the world."},
	{Id: 25, Title: "Into the Abyss", Image: "https://example.com/images/into-abyss.jpg", Description: "A psychological thriller set deep in the mountains, where nothing is as it seems."},
	{Id: 26, Title: "The Last Breath", Image: "https://example.com/images/last-breath.jpg", Description: "A heart-wrenching story of love and survival in a post-apocalyptic world."},
	{Id: 27, Title: "The Golden Key", Image: "https://example.com/images/golden-key.jpg", Description: "A young man discovers an ancient key that unlocks a hidden world of secrets."},
	{Id: 28, Title: "Under the Red Sky", Image: "https://example.com/images/under-red-sky.jpg", Description: "An emotionally charged drama set in a city on the brink of collapse."},
	{Id: 29, Title: "Eclipse", Image: "https://example.com/images/eclipse.jpg", Description: "A supernatural thriller where a series of murders coincides with a rare solar eclipse."},
	{Id: 30, Title: "Echoes of the Past", Image: "https://example.com/images/echoes-past.jpg", Description: "A historical epic following the untold stories of a family throughout the centuries."},
}
