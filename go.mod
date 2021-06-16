module main

go 1.16

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/joho/godotenv v1.3.0
	commands v0.0.0
	cmd_handler v0.0.0
)

replace commands v0.0.0 => ./src/commands

replace cmd_handler v0.0.0 => ./src/cmd_handler
