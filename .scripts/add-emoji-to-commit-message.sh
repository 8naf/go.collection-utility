source ./.scripts/_common.sh

declare -A type_emojis
while IFS="=" read -r type emoji; do
  type_emojis["$type"]="$emoji"
done < <(jq -r '.[] | "\(.type)=\(.emoji)"' ./.config/commit-types.json)

printf "adding emoji to commit message\n"

declare commit_msg_file="$1"
declare commit_msg=$(cat "$commit_msg_file")
declare commit_msg_header=$(echo "$commit_msg" | head -n1)

declare emoji=":\w+: "
declare type="\w+"
declare scope="\(.*\)"
declare subject=".*"
declare pattern="^($emoji)?($type)($scope)?!?: ($subject)$"

IFS="-"
read -r emoji type <<< $(echo "$commit_msg_header" | sed -En "s/${pattern}/\1-\2/p")

if [[ -n "$emoji" ]]; then
	printf "found emoji: ${CYAN}$emoji${RESET}\n"
	printf "skipped: the commit messsage already added an emoji\n"
	exit 0
fi

if [[ -z "$type" ]]; then
	printf "skipped: the commit messsage does not have the type\n"
	exit 1
fi

if [[ -z "${type_emojis[$type]}" ]]; then
	printf "found type: ${CYAN}$type${RESET}\n"
	printf "skipped: the type is not one of: "
	printf '%s, ' "${!type_emojis[@]}" | sed 's/, $/\n/'
	exit 1
fi

emoji=${type_emojis[$type]}
declare new_msg="$emoji $commit_msg"
echo "$new_msg" > "$commit_msg_file"
printf "added: ${CYAN}$emoji${RESET} emoji for ${CYAN}$type${RESET} type\n"
