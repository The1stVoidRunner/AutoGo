package main

type obj struct {
	Bot 			string
	Admin 			string
	Status 			string
	BotMaster		bool
	BotCommander	string
	CmdsRun			int
	Prefix 		string
	GreetMsg	string
	ByeMsg		string
	RoleSys		string
	Name		string
	AntiLink	bool
	Action		string
	Silent		bool
	Active		string
	HelpCmd		string
}


type role struct {
	ID	string
}


type invite struct {
	MaxAge		int
	MaxUses		int
	Temporary	bool
	XkcdPass	bool
}


type commands struct {
	Greet			string
	Bye				string
	Prefix			string
	Kick			string
	Ban 			string
	Autorole		string
	SetPunish		string
	AllowLinks		string
	DenyLinks		string
	AddMaster		string
	DelMaster		string
	Invites			string
	Meme			string
	Joke			string
	Give			string
	Take			string
	Giveme			string
	Mute			string
	Unmute			string
}
