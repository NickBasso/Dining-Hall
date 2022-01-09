
# Dining-Hall

Dining Hall API which takes care of supplying kitchen with requests and receiving deliveries afterwards

## Setup & Run:

### Standard:
- change directory to root of DinningHall service
- run "go run src/main.go"

### Docker:
- change directory to root of DinningHall service
- run "docker build -t dininghall ."
- run "docker run -it -p 4005:4005 dininghall"
- check out port 4005