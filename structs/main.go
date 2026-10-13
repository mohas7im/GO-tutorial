package main

import "fmt"

type Patient struct {
	Id             int
	Name           string
	Age            int
	Email          string
	Specialization string
	Experience     string
}

func main() {
	singlePatient := Patient{
		Id:    1,
		Name:  "Hashim",
		Age:   25,
		Email: "mohammedhashim530@gmail.com",
	}
	fmt.Println("Single Patient:", singlePatient)

	patients := []Patient{
		{
			Id:             1,
			Name:           "Hashim",
			Specialization: "doctor",
			Experience:     "5 years",
		},
		{
			Id:             2,
			Name:           "rahul",
			Specialization: "doctor pro",
			Experience:     "6 years",
		},
	}

	fmt.Println("\nList of Patients:")
	for _, patient := range patients {
		fmt.Println(patient.Id, patient.Name, patient.Specialization, patient.Experience)
	}
}
