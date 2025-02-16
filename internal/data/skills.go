package data

import "github.com/ropfoo/gothulhu/internal/model"

var SkillDefinitions = []model.SkillDefinition{
	// Interpersonal Skills
	{
		Id:          "chm",
		Name:        "Charm",
		Description: "The ability to impress and influence others through personality and social grace.",
	},
	{
		Id:          "itm",
		Name:        "Intimidate",
		Description: "The ability to influence others through threats, fear, and shows of force.",
	},
	{
		Id:          "prs",
		Name:        "Persuade",
		Description: "The ability to convince others through logic, fast-talking, or debate.",
	},
	{
		Id:          "fst",
		Name:        "Fast Talk",
		Description: "The ability to mislead or confuse others through quick thinking and rapid speech.",
	},

	// Physical Skills
	{
		Id:          "clm",
		Name:        "Climb",
		Description: "The ability to scale surfaces and navigate vertical challenges.",
	},
	{
		Id:          "drv",
		Name:        "Drive Auto",
		Description: "The ability to operate and control automobiles and other motor vehicles.",
	},
	{
		Id:          "fig",
		Name:        "Fighting (Brawl)",
		Description: "The ability to engage in unarmed combat and close-quarters fighting.",
	},
	{
		Id:          "frm",
		Name:        "Firearms",
		Description: "The ability to use and maintain firearms effectively.",
	},
	{
		Id:          "fad",
		Name:        "First Aid",
		Description: "The ability to provide emergency medical treatment and stabilize injuries.",
	},
	{
		Id:          "jmp",
		Name:        "Jump",
		Description: "The ability to leap across distances and heights effectively.",
	},
	{
		Id:          "swm",
		Name:        "Swim",
		Description: "The ability to stay afloat and move through water efficiently.",
	},
	{
		Id:          "thw",
		Name:        "Throw",
		Description: "The ability to accurately throw objects and weapons.",
	},

	// Investigative Skills
	{
		Id:          "art",
		Name:        "Art/Craft",
		Description: "Knowledge and skill in various artistic and craft disciplines.",
	},
	{
		Id:          "dsg",
		Name:        "Disguise",
		Description: "The ability to alter one's appearance and impersonate others.",
	},
	{
		Id:          "spt",
		Name:        "Spot Hidden",
		Description: "The ability to notice hidden objects, clues, and details in the environment.",
	},
	{
		Id:          "lsn",
		Name:        "Listen",
		Description: "The ability to detect and interpret sounds and conversations.",
	},
	{
		Id:          "lck",
		Name:        "Locksmith",
		Description: "The ability to understand, open, and modify locks and security systems.",
	},
	{
		Id:          "lib",
		Name:        "Library Use",
		Description: "The ability to find and extract information from books and archives.",
	},
	{
		Id:          "trc",
		Name:        "Track",
		Description: "The ability to follow trails and track movements of creatures or people.",
	},

	// Knowledge Skills
	{
		Id:          "acc",
		Name:        "Accounting",
		Description: "Knowledge of financial records, transactions, and business practices.",
	},
	{
		Id:          "ant",
		Name:        "Anthropology",
		Description: "Knowledge of human cultures, societies, and their development.",
	},
	{
		Id:          "arc",
		Name:        "Archaeology",
		Description: "Knowledge of historical artifacts, sites, and excavation methods.",
	},
	{
		Id:          "his",
		Name:        "History",
		Description: "Knowledge of historical events, periods, and their significance.",
	},
	{
		Id:          "law",
		Name:        "Law",
		Description: "Knowledge of legal systems, regulations, and procedures.",
	},
	{
		Id:          "med",
		Name:        "Medicine",
		Description: "Knowledge of medical procedures, diseases, and treatments.",
	},
	{
		Id:          "occ",
		Name:        "Occult",
		Description: "Knowledge of mystical traditions, supernatural beliefs, and rituals.",
	},
	{
		Id:          "sci",
		Name:        "Science",
		Description: "Knowledge of scientific principles, methods, and discoveries.",
	},

	// Language Skills
	{
		Id:          "lng",
		Name:        "Language (Other)",
		Description: "The ability to speak, read, and understand foreign languages.",
	},
	{
		Id:          "own",
		Name:        "Own Language",
		Description: "Proficiency in one's native language and literacy.",
	},

	// Technical Skills
	{
		Id:          "ele",
		Name:        "Electrical Repair",
		Description: "The ability to repair and maintain electrical equipment.",
	},
	{
		Id:          "mch",
		Name:        "Mechanical Repair",
		Description: "The ability to repair and maintain mechanical devices.",
	},
	{
		Id:          "opr",
		Name:        "Operate Heavy Machinery",
		Description: "The ability to operate and control large industrial equipment.",
	},
	{
		Id:          "pht",
		Name:        "Photography",
		Description: "The ability to take, develop, and analyze photographs.",
	},

	// Specialized Skills
	{
		Id:          "psy",
		Name:        "Psychoanalysis",
		Description: "The ability to understand and treat mental disorders.",
	},
	{
		Id:          "rid",
		Name:        "Ride",
		Description: "The ability to mount and control horses and other animals.",
	},
	{
		Id:          "plt",
		Name:        "Pilot",
		Description: "The ability to operate aircraft and other flying vehicles.",
	},
	{
		Id:          "sur",
		Name:        "Survival",
		Description: "The ability to stay alive in hostile environments.",
	},
}
