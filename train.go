package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// cred           = "mongodb://localhost:27017"
	// dbname         = "test"
	// collectionName = "trainers"
	limit = 1
)

var (
	wg sync.WaitGroup
	ch = make(chan int, limit)
)

type Data struct {
	TrainNo   string `bson:"trainNo   string"`
	TrainName string `bson:"trainName string"`
	SEQ       string `bson:"seq       string"`
	Code      string `bson:"code      string"`
	StName    string `bson:"stName    string"`
	ATime     string `bson:"atime     string"`
	DTime     string `bson:"dtime     string"`
	Distance  string `bson:"distance  string"`
	SS        string `bson:"ss        string"`
	SSname    string `bson:"ssname    string"`
	Ds        string `bson:"ds        string"`
	DsName    string `bson:"dsName    string"`
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename) // Open CSV file
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll() // Read File into a Variable
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func dbConn() (*mongo.Collection, *mongo.Client) {
	cred := goDotEnvVariable("CRED")
	dbname := goDotEnvVariable("DBNAME")
	collectionName := goDotEnvVariable("COLLECTION")
	clientOptions := options.Client().ApplyURI(cred)
	client, err := mongo.Connect(context.TODO(), clientOptions) // Connect to MongoDB
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil) // Check the connection
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbname).Collection(collectionName)
	fmt.Println("Connected to MongoDB!")
	return collection, client
}

func getallTrains(w http.ResponseWriter, r *http.Request) {
	collection, client := dbConn()
	defer client.Disconnect(context.TODO())
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var trains []Data
	if err = cursor.All(context.TODO(), &trains); err != nil {
		log.Fatal(err)
	}
	bytedata, err := json.MarshalIndent(trains, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func insert(read *bool) {
	collection, client := dbConn()
	defer client.Disconnect(context.TODO())
	if *read {
		csvData, err := ReadCsv("Indian_railway1.csv")
		if err != nil {
			panic(err)
		}
		csvData = csvData[1:31]
		for _, line := range csvData {
			ch <- 1     //	wg.Add(1)
			go func() { //defer wg.Done()
				data := Data{
					TrainNo:   line[0],
					TrainName: line[1],
					SEQ:       line[2],
					Code:      line[3],
					StName:    line[4],
					ATime:     line[5],
					DTime:     line[6],
					Distance:  line[7],
					SS:        line[8],
					SSname:    line[9],
					Ds:        line[10],
					DsName:    line[11],
				} //time.Sleep(time.Second)
				_, err := collection.InsertOne(context.TODO(), data)
				if err != nil {
					panic(err)
				}
				<-ch
			}() //	wg.Wait()
		}
		for i := 0; i < limit; i++ {
			ch <- 1
		}
	} else {
		fmt.Println("No file Read")
	}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load( /*"../main.env"*/ "main.env") // load .env file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	start, _ := strconv.Atoi(params["id"])
	start--
	skip := int64(start * 10)
	if skip <= 0 {
		skip = 0
	}
	collection, client := dbConn()
	defer client.Disconnect(context.TODO())
	options := options.Find()
	options.SetLimit(5)
	options.SetSkip(skip)
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, options)
	if err != nil {
		log.Fatal(err)
	}
	var trains []Data
	if err = cursor.All(context.TODO(), &trains); err != nil {
		log.Fatal(err)
	}
	bytedata, err := json.MarshalIndent(trains, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func main() {
	run := flag.Bool("fork", false, "a bool")
	read := flag.Bool("now", false, "a bool")
	flag.Parse()
	start := time.Now()
	if *run {
		insert(read)
	} else {
		fmt.Println("No file found")
	} //insert()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	router := mux.NewRouter()
	router.HandleFunc("/Trains/{id}", getPost).Methods(("GET"))
	router.HandleFunc("/Trains", getallTrains).Methods(("GET"))
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.HandleFunc("/Trains", getallTrains)
	log.Fatal(http.ListenAndServe(":8080", router))
	// http.ListenAndServe(":8080", nil)
}
