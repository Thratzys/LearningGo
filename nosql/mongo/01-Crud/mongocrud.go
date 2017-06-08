package main

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type Movie struct {
	Title		string 	`json: "title"`
	Poster	string	`json: "poster"`
	Genre		string	`json: "genre"`
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}

func GetMovies() []Movie {
	session := GetSession()
	defer session.Close()

	collection := session.DB("movies").C("movies")

	var movies []Movie
	err := collection.Find(bson.M{}).All(&movies)
	if err != nil {
		panic(err)
	}

	return movies
}

func ShowMovie(title string) Movie {
	session := GetSession()
	defer session.Close()

	collection := session.DB("movies").C("movies")
	
	var movie = Movie{}
	err := collection.Find(bson.M{"title": title}).One(&movie)
	if err != nil {
		panic(err)
	}

	return movie
}

func CreateMovie(movie *Movie) bool {
	session := GetSession()
	defer session.Close()

	collection := session.DB("movies").C("movies")

	_, err := collection.UpsertId(movie.Title, movie)
	if err != nil {
		return false
	}
	return true
}

func DeleteMovie(id string) bool {
	session := GetSession()
	defer session.Close()

	collection := session.DB("movies").C("movies")
	err := collection.RemoveId(id)
	if err != nil {
		panic(err)
	}

	return true
}

func main() {
	switch os.Args[1] {
		case "all":
			fmt.Println(GetMovies())
		case "show":
			fmt.Println(ShowMovie(os.Args[2]))
		case "create":
			var movie = Movie {
				Title:  os.Args[2],
				Poster: os.Args[3],
				Genre:	os.Args[4],
			}
			CreateMovie(&movie)
		case "delete":
			fmt.Println(DeleteMovie(os.Args[2]))
	}
}