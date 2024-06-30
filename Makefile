make:

# argument is not valid
argtest:
	go run .
	go run . "0...." "....." "....." "....." "....."
	go run . "....." ".A..." "....." "....." "....."
	go run . "....." "....." "....." "....." "...."
	go run . "....." ".55..." "....." "....." "....."
	go run . "....." "....." "....." "....." "...."
	go run . "....." "....." "....." "....." "....." "....."
	go run . "*...." "....." "....." "....." "....."
	go run . "....." "..1.." "2...." "...3." ".2..."

	go run . "....." "....." "....." "....."
	go run . "....." "....." "....."
	go run . "....." "....."
	go run . "....."
	go run .

	go run . "....." "....." "....." "....." "....." "....."
	go run . "....." "....." "....." "....." "...."
	go run . "....." "....." "....." "....." "..."
	go run . "....." "....." "....." "....." ".."
	go run . "....." "....." "....." "....." "."

	go run . "....." "....." "....." "....." "......"
	go run . "....." "....." "....." "......" "....."
	go run . "....." "....." "......" "....." "....."
	go run . "....." "......" "....." "....." "....."
	go run . "......" "....." "....." "....." "....."
	go run . "......" "......" "......" "......" "......"

	go run . "....." "....." "....." "....." "...１."
	go run . "....." "....." "....." "....." "....あ"
	go run . "....." "....." "....." "....." "....a"
	go run . "....." "....." "....." "....." "....1"
	go run . "....." "....." "....1" "....." "....."
	go run . "....." "..1.." "....." "....." "....."
	go run . "...1." "....." "....." "....." "....."

# solution exists
passtest:
	go run . "....." "....." "....." "....." "....."
	go run . "..3.." ".5..." "....." "3...." "..2.."
#	go run . "2...." "...4." "....." ".1..." "....3"
#	go run . "...4." ".1..." "....." "....2" "5...."
#	go run . "5...." "...1." "....." "....3" "4...."
#	go run . "..2.." "3...." "...4." "....." ".1..."
#	go run . "....." "..4.." ".2..." "....." "3...2"
#	go run . "4...." "....." "..2.." "...1." "....."
#	go run . ".3..." "..2.." "....." "4...." "...1."
	go run . "....." "..3.." "2...." "....." "...4."
#	go run . "..1.." ".4..." "....." "...2." "....3"
#	go run . "....." "..2.." ".3..." "4...." "....1"
#	go run . "....." "..4.." ".1..." "....." "...3."
	go run . "...2." "....." "....3" "....." "2...."
	go run . "....." "2...." "....." "....3" "....."
#	go run . "....3" "..2.." "....." "....." "..1.."
#	go run . "..4.." "....." ".3..." "....." "1...."

# solution does not exist
failtest:
	go run . "....." "..5.." ".2..." "...4." "..1.."
	go run . "....." ".1..." "..3.." "...2." "....."
	go run . ".1..." "..2.." "...3." "..1.." "....."
	go run . "..1.." "....." "...3." "2...." "..4.."
	go run . "..2.." "...1." ".3..." "....." "....4"
	go run . "2...." "...4." "....." "..1.." ".3..."
	go run . "....." "....2" "..1.." "....." "....3"
	go run . "2...." "....." "...1." "....4" "....."
	go run . "....2" "...3." "....." "1...." "..4.."
	go run . "....." "...2." "....3" "..1.." "4...."
	go run . "3...." "..1.." "....." "....4" "..2.."
	go run . "....." "..1.." "....." "...3." "2...."
	go run . "....3" "....." "..2.." "....." "1...."
# division exists
divtest:
	go run . ".4567" "4.456" "54.45" "654.4" "7654."
	go run . "66.66" "66.66" "66.66" "66.66" "66.66"
	go run . "33.66" "33.66" "..577" "66799" "66799"
	go run . "543.5" "43.56" "3.567" ".5678" "56789"
	go run . "5.345" "65.34" "765.3" "8765." "98765"
	go run . "123.5" "234.6" "345.7" "456.8" "567.9"
	go run . "56789" ".5678" "..567" "...56" "....5"
	go run . "12.45" "23.56" "34.67" "45.78" "56.89"
	go run . "5.432" "65.43" "765.4" "8765." "98765"
	go run . "123.5" "234.6" "345.7" "456.8" "567.9"
	go run . "56789" ".5678" "..567" "...56" "....5"
	go run . "12.45" "23.56" "34.67" "45.78" "56.89"