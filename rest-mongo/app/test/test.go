package main

import(
    "fmt"
    "gopkg.in/mgo.v2/bson"
    "github.com/forexample/rest-mongo/app/models"
)

func main(){
    movie := models.Movie{
        ID: bson.NewObjectId(),
        Name:"Avengers Endgame",
        CoverImage:"https://www.hindustantimes.com/rf/image_size_960x540/HT/p2/2018/12/24/Pictures/_eff2e394-0771-11e9-8b39-01e96223c804.jpeg",
        Description:"Tony Stark will unite all six Infinity Stones and die, says Avengers Endgame theory",
    }

    fmt.Println(movie)
}