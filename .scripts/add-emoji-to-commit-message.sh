declare -A type_emojis
while IFS="=" read -r type emoji; do
  type_emojis["$type"]="$emoji"
done < <(jq -r '.[] | "\(.type)=\(.emoji)"' ./.config/commit-types.json)

commit_msg_file="$1"
commit_msg=$(cat "$commit_msg_file")

type=$(echo "$commit_msg" | head -n1 | grep -oP '^(\w+)(?=(\(.*\))?!?:)')

if [[ -n "$type" && -n "${type_emojis[$type]}" ]]; then
	emoji=${type_emojis[$type]}
	new_msg="$emoji $commit_msg"
	echo "$new_msg" > "$commit_msg_file"
fi
