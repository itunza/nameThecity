package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Country structure to store country and its capital.
type Country struct {
	name    string
	capital string
}

func main() {
	// List of some countries and their capitals.
	countries := []Country{
		{"Algeria", "Algiers"},
		{"Angola", "Luanda"},
		{"Benin", "Porto-Novo"},
		{"Botswana", "Gaborone"},
		{"Egypt", "Cairo"},
		{"Burkina Faso", "Ouagadougou"},
		{"Burundi", "Gitega"},
		{"Cape Verde", "Praia"},
		{"Cameroon", "Yaoundé"},
		{"Central African Republic", "Bangui"},
		{"Chad", "N'Djamena"},
		{"Comoros", "Moroni"},
		{"Democratic Republic of the Congo", "Kinshasa"},
		{"Republic of the Congo", "Brazzaville"},
		{"Djibouti", "Djibouti"},
		{"Egypt", "Cairo"},
		{"Equatorial Guinea", "Malabo"},
		{"Eritrea", "Asmara"},
		{"Eswatini (formerly Swaziland)", "Mbabane (administrative), Lobamba (royal and legislative)"},
		{"Ethiopia", "Addis Ababa"},
		{"Gabon", "Libreville"},
		{"The Gambia", "Banjul"},
		{"Ghana", "Accra"},
		{"Guinea", "Conakry"},
		{"Guinea-Bissau", "Bissau"},
		{"Ivory Coast (Côte d'Ivoire)", "Yamoussoukro"},
		{"Kenya", "Nairobi"},
		{"Lesotho", "Maseru"},
		{"Liberia", "Monrovia"},
		{"Libya", "Tripoli"},
		{"Madagascar", "Antananarivo"},
		{"Malawi", "Lilongwe"},
		{"Mali", "Bamako"},
		{"Mauritania", "Nouakchott"},
		{"Mauritius", "Port Louis"},
		{"Morocco", "Rabat"},
		{"Mozambique", "Maputo"},
		{"Namibia", "Windhoek"},
		{"Niger", "Niamey"},
		{"Nigeria", "Abuja"},
		{"Rwanda", "Kigali"},
		{"São Tomé and Príncipe", "São Tomé"},
		{"Senegal", "Dakar"},
		{"Seychelles", "Victoria"},
		{"Sierra Leone", "Freetown"},
		{"Somalia", "Mogadishu"},
		{"South Africa", "Pretoria (administrative), Bloemfontein (judicial), Cape Town (legislative)"},
		{"South Sudan", "Juba"},
		{"Sudan", "Khartoum"},
		{"Tanzania", "Dodoma"},
		{"Togo", "Lomé"},
		{"Tunisia", "Tunis"},
		{"Uganda", "Kampala"},
		{"Zambia", "Lusaka"},
		{"Zimbabwe", "Harare"},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for _, country := range countries {
		fmt.Printf("What is the capital of %s?\n", country.name)
		scanner.Scan()
		answer := scanner.Text()

		if strings.EqualFold(answer, country.capital) {
			fmt.Println("Correct! Well done!")
		} else {
			fmt.Printf("Oops! The correct answer is %s.\n", country.capital)
		}
	}
}
