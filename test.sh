start_time="$(date -u +%s.%N)"
echo "Starting tests..."
# Reverse test
go run . --reverse example00.txt

go run . --reverse=example00.txt

go run . --reverse=example01.txt

go run . --reverse=example02.txt

go run . --reverse=example03.txt

go run . --reverse=example04.txt

go run . --reverse=example05.txt

go run . --reverse=example06.txt

go run . --reverse=example07.txt

go run . --reverse=example08.txt

go run . --reverse=example09.txt

go run . --reverse=example10.txt

go run . --reverse=example11.txt

go run . --reverse=example12.txt

go run . --reverse=example13.txt

# Output test
go run . --output=banner.txt "hello" standard

go run . --output=banner.txt "Hello There\!" shadow

go run . --output test.txt banana standard 

go run . --output=test00.txt "First\nTest" standard

go run . --output=test01.txt "hello" standard

go run . --output=test02.txt "123 -> #$%" standard

go run . --output=test03.txt "432 -> #$%&@" shadow

go run .  --output=test04.txt "There" shadow

go run . --output=test05.txt "123 -> \"#$%@" tinkertoy

go run . --output=test06.txt "2 you" tinkertoy

go run . --output=test07.txt "Testing long output\!" standard

go run . --output=test08.txt "abcDEF" standard

go run . --output=test010.txt "#%$" standard

go run . --output=test011.txt "abcDEF 123" standard

go run . --output=test08.txt "abcDEF" randomFont

# go run . --output=test010.txt "#%$" randomFont

# go run . --output=test011.txt "abcDEF 123" randomFont


end_time="$(date -u +%s.%N)"
echo "Tests finished!"
elapsed="$(awk "BEGIN { print $end_time - $start_time }")"
echo "Total of $elapsed seconds elapsed for processing"