package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	l "github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

type MusicScore struct {
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

type Hoge struct {
	Val string `json:"val"`
}


func main() {
	l.Start(handler)
}



func handler(params requestParam) (Hoge, error) {
	fmt.Println(params)
	db := connectDb()
	defer db.Close()
	name := "%" + params.MusicName + "%"
	rows, err := db.Query("SELECT * FROM music_scores WHERE instrument = ? AND musicName like ? limit 4", params.Instrument, name)
	_ = mapping(rows)
	hoge := Hoge{"hogehoge"}
	return hoge, err
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

func mapping(rows *sql.Rows) MusicScore {
	//var  musicScores []musicScore
	var mcs MusicScore
	for rows.Next() {
		var mc MusicScore
		err := rows.Scan(&mc.Id, &mc.ServiceName, &mc.MusicName, &mc.Composer, &mc.Price, &mc.Url, &mc.Instrument, &mc.ServiceId, &mc.Difficulty, &mc.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		mcs = mc
		//musicScores = append(musicScores,mc)
	}
	return mcs
}
