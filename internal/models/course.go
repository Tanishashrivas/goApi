package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	CourseId    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName  string             `json:"coursename"`
	CoursePrice int                `json:"price"`
	Author      *Author            `json:"author"`
}

// fake db
// var Courses = []Course{
// 	{
// 		CourseId:    "C101",
// 		CourseName:  "Go Programming for Beginners",
// 		CoursePrice: 499,
// 		Author: &Author{
// 			FullName: "John Doe",
// 			Website:  "https://johndoe.dev",
// 		},
// 	},
// 	{
// 		CourseId:    "C102",
// 		CourseName:  "Advanced Golang Techniques",
// 		CoursePrice: 799,
// 		Author: &Author{
// 			FullName: "Jane Smith",
// 			Website:  "https://janesmith.com",
// 		},
// 	},
// 	{
// 		CourseId:    "C103",
// 		CourseName:  "Building REST APIs with Go",
// 		CoursePrice: 599,
// 		Author: &Author{
// 			FullName: "Alice Johnson",
// 			Website:  "https://alicejohnson.io",
// 		},
// 	},
// 	{
// 		CourseId:    "C104",
// 		CourseName:  "Concurrency in Go",
// 		CoursePrice: 899,
// 		Author: &Author{
// 			FullName: "Bob Williams",
// 			Website:  "https://bobwilliams.tech",
// 		},
// 	},
// 	{
// 		CourseId:    "C105",
// 		CourseName:  "Mastering Go Routines",
// 		CoursePrice: 699,
// 		Author: &Author{
// 			FullName: "Charlie Brown",
// 			Website:  "https://charliebrown.dev",
// 		},
// 	},
// 	{
// 		CourseId:    "C106",
// 		CourseName:  "Go for Web Development",
// 		CoursePrice: 649,
// 		Author: &Author{
// 			FullName: "David Wilson",
// 			Website:  "https://davidwilson.com",
// 		},
// 	},
// 	{
// 		CourseId:    "C107",
// 		CourseName:  "Data Structures and Algorithms in Go",
// 		CoursePrice: 749,
// 		Author: &Author{
// 			FullName: "Eve Adams",
// 			Website:  "https://eveadams.io",
// 		},
// 	},
// 	{
// 		CourseId:    "C108",
// 		CourseName:  "Go Microservices with Kubernetes",
// 		CoursePrice: 999,
// 		Author: &Author{
// 			FullName: "Frank Martin",
// 			Website:  "https://frankmartin.tech",
// 		},
// 	},
// 	{
// 		CourseId:    "C109",
// 		CourseName:  "Unit Testing in Go",
// 		CoursePrice: 550,
// 		Author: &Author{
// 			FullName: "Grace Lee",
// 			Website:  "https://gracelee.dev",
// 		},
// 	},
// 	{
// 		CourseId:    "C110",
// 		CourseName:  "Building Scalable Systems with Go",
// 		CoursePrice: 1100,
// 		Author: &Author{
// 			FullName: "Henry Ford",
// 			Website:  "https://henryford.io",
// 		},
// 	},
// }
