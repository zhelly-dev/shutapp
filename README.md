# ShutApp
> A straightforward productivity tool that temporarily restricts access to selected apps for a set duration.
Block your time consuming applications and boost your productivity

## Building
1. Clone the repository
   `git clone https://github.com/zhelly-dev/shutapp.git`
2. Navigate to the project directory:
   `cd shutapp`
3. Install dependencies:
   `go mod tidy`
4. Build
   `go build main.go`

## Usage
`shutapp.exe -time=[time in minutes] -name=[process.exe]`

## TODO
1. Add banlist.txt in format `[process.exe] - [time]`
2. Add GUI
