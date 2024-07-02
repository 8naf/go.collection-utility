#BLACK="\e[30m"
#RED="\e[31m"
#GREEN="\e[32m"
#YELLOW="\e[33m"
#BLUE="\e[34m"
#PURPLE="\e[35m"
#CYAN="\e[36m"
#WHITE="\e[37m"
#RESET="\e[0m"

export TERM=xterm-256color

BLACK=$(tput setaf 0)
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
BLUE=$(tput setaf 4)
PURPLE=$(tput setaf 5)
CYAN=$(tput setaf 6)
WHITE=$(tput setaf 7)
RESET=$(tput sgr0)

hline() {
	color=${1:-$RESET}

	box_drawings_light_horizontal="\u2500"
	start=1
	end=$(tput cols)
	printf "${color}$box_drawings_light_horizontal%.0s${RESET}" $(seq $start $end);
	echo
}
