# GotrueBot
This is a telegram bot written in golang to fetch mobile number details from truecaller.

## Installation

```bash
git clone https://github.com/Linuxinet/GotrueBot

cd GotrueBot
```

Create a `.env` file with `BOT_TOKEN` and `TRUECALLER_TOKEN` variables and run / build.

```bash
go run .
```

`TRUECALLER_TOKEN` was your truecaller authentication token 
(**Eg:**   `Bearer xxxxx--xxxxxx-v-xxxxxxxxx-xxxxx-xxxxxxxxxxxxxxxxxxxxxx`)

`BOT_TOKEN` is a telegram bot token. you can get it from the [@Botfather Bot](https://t.me/BotFather) in telegram.

## TODO

+ Modify it to host in heroku or other platforms.

_**Pull requests are welcome**_