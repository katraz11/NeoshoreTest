package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Name struct{
	Common string
}

type Currency struct{
	Name string  `json:"name"`
	Symbol string`json:"symbol"`
}

type CurrencyName struct{
	name string
}

type CurrencySymbo struct{
	name string
}


type CountriesFetched struct{
	Name Name `json:"name"`
	Cca2 string  `json:"cca2"`// country code on 2 digits
	Cca3 string  `json:"cca3"`// country code on 3 digits
	Region string  `json:"region"`
	 Subregion string `json:"subregion"`
	Currency map[string] Currency  `json:"currencies"`
	
	
  };

  type Countries struct{
	name string  // official name in english
	cca2 string  // country code on 2 digits
	cca3 string  // country code on 3 digits
	region string
	subregion string
	currency string 
	currencySymbo string
  }

func main() {

	response,err:=http.Get("https://restcountries.com/v3.1/all")

	if err!=nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData,err:= ioutil.ReadAll(response.Body)

	if err !=nil{
		log.Fatal(err)
	}



		
	  var countryesFetched [] CountriesFetched
	  var countryes[] Countries
	  var EuropeCountries[] Countries
	  var NorAmericaCountryes[] Countries
	 
	  var count int;
		err2 := json.Unmarshal([]byte(responseData),&countryesFetched)

		if err2 !=nil{
			fmt.Println("error unmarshalling json %v",err2)
		}

		for _,value:=range countryesFetched{
		
			var country Countries
			country.cca2=value.Cca2
			country.cca3=value.Cca3
			country.name=value.Name.Common
			country.region=value.Region
			country.subregion=value.Subregion
			;
			for key,values:= range value.Currency{
				country.currency=key
				country.currencySymbo=values.Symbol
			}
			if country.region=="Europe" {
				EuropeCountries=append(EuropeCountries,country)
			}else if country.subregion=="North America" {
				NorAmericaCountryes=append(NorAmericaCountryes, country)
			}
			countryes=append(countryes,country)
			count++
		}


		fmt.Printf("European countries: %v \n",EuropeCountries)
		fmt.Printf("Nourth America countries: %v",NorAmericaCountryes)
		

}
