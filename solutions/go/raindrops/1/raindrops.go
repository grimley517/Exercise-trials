package raindrops
import (
    "strconv"
)

func getDivisors () map[int]string {
    divisors := map[int]string{
        3: "Pling",
        5: "Plang",
        7: "Plong"}
    return divisors
    }

func Convert (number int) string {
	result := ""
    for key, value := range getDivisors(){
        if (number % key==0){
            result = result + value
        }
    }
    if result == "" {
        result = strconv.Itoa(number)
    }
    return result
}
