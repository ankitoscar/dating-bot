module example.com/dating-bot

go 1.18

require example.com/chat v0.0.0

require github.com/joho/godotenv v1.5.1 // indirect

replace example.com/chat => ./chat
