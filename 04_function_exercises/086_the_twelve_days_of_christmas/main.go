package main

import "fmt"

func main() {

	for i := 1; i <= 12; i++ {
		fmt.Println(displayLyrics(i))
	}

}

func displayLyrics(numOfDay int) string {

	present := func(numOfDay int) string {
		present := ""
		if numOfDay > 1 {
			for i := numOfDay; i != 0; i-- {
				switch i {
				case 1:
					present += "And a partridge in a pear tree \n"
				case 2:
					present += "Two turtledoves \n"
				case 3:
					present += "Three French hens \n"
				case 4:
					present += "Four calling birds \n"
				case 5:
					present += "Five gold rings (five golden rings) \n"
				case 6:
					present += "Six geese a-laying \n"
				case 7:
					present += "Seven swans a-swimming \n"
				case 8:
					present += "Eight maids a-milking \n"
				case 9:
					present += "Nine ladies dancing \n"
				case 10:
					present += "Ten lords a-leaping \n"
				case 11:
					present += "I sent eleven pipers piping \n"
				case 12:
					present += "Twelve drummers drumming \n"
				}
			} 
		} else {
			present += "A partridge in a pear tree \n"
		}
		return present
	}

	lyrics := "On the " + convertIntegerToOrdinal(numOfDay) + " day of Christmas, my true love sent to me: \n" + present(numOfDay)

	return lyrics

}

func convertIntegerToOrdinal(num int) string {

	ordNum := ""

	switch num {
	case 1:
		ordNum = "First"
	case 2:
		ordNum = "Second"
	case 3:
		ordNum = "Third"
	case 4:
		ordNum = "Fourth"
	case 5:
		ordNum = "Fifth"
	case 6:
		ordNum = "Sixth"
	case 7:
		ordNum = "Seventh"
	case 8:
		ordNum = "Eighth"
	case 9:
		ordNum = "Ninth"
	case 10:
		ordNum = "Tenth"
	case 11:
		ordNum = "Eleventh"
	case 12:
		ordNum = "Twelfth"
	}
	
 return ordNum

}