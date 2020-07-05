package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	l "github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

type musicScore struct {
	Id			int    `json:"id"`
	ServiceName string `json:"serviceName"`
	MusicName   string `json:"musicName"`
	Composer    string `json:"composer"`
	Price       int    `json:"price"`
	Url         string `json:"url"`
	Instrument  string `json:"instrument"`
	ServiceId   string `json:"serviceId"`
	Difficulty  string `json:"difficulty"`
	CreatedAt   string `json:"createdAt"`

}

type requestParam struct {
	Instrument string `json:"instrument"`
	MusicName  string `json:"musicName"`
}

type hoge struct {
	Str string `json:"str"`
}


func main() {
	l.Start(handler)
}


func handler(params requestParam) (hoge, error) {
	fmt.Println(params)
	db := connectDb()
	defer db.Close()
	name := "%" + params.MusicName + "%"
	rows, err := db.Query("SELECT * FROM music_scores WHERE instrument = ? AND musicName like ? ", params.Instrument, name)
	_ = mapping(rows)

	hoges := hoge{"hogehoge"}
	return hoges, err
}

func connectDb() *sql.DB {
	host:=os.Getenv("DB_HOST")
	port:="3306"
	pwd:=os.Getenv("DB_PASSWORD")
	user:= os.Getenv("DB_USER")
	dbName:= os.Getenv("DB_NAME")
	dsn:= user + ":" + pwd + "@tcp(" + host  + ":" +port + ")" + "/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		log.Fatal(err)
	}
	return db
}

func mapping(rows *sql.Rows) []musicScore {
	var  musicScores []musicScore
	for rows.Next() {
		var mc musicScore
		err := rows.Scan(&mc.Id, &mc.ServiceName, &mc.MusicName, &mc.Composer, &mc.Price, &mc.Url, &mc.Instrument, &mc.ServiceId, &mc.Difficulty, &mc.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		musicScores = append(musicScores,mc)
	}
	return musicScores
}
