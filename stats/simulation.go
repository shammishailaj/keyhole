// Copyright 2018 Kuei-chun Chen. All rights reserved.

package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CollectionName -
var CollectionName = "keyhole"

// Schema -
type Schema struct {
	ID            string `json:"_id" bson:"_id"`
	FavoriteBook  string `json:"favoriteBook" bson:"favoriteBook"`
	FavoriteCity  string `json:"favoriteCity" bson:"favoriteCity"`
	FavoriteMovie string `json:"favoriteMovie" bson:"favoriteMovie"`
	FavoriteMusic string `json:"favoriteMusic" bson:"favoriteMusic"`
	FavoriteSport string `json:"favoriteSport" bson:"favoriteSport"`
}

// Favorites -
type Favorites struct {
	Sports []string
	Music  []string
	Cities []string
	Books  []string
	Movies []string
}

var favorites = Favorites{
	Sports: []string{"Baseball", "Boxing", "Dodgeball", "Figure skating", "Football", "Horse racing", "Mountaineering", "Skateboard", "Ski", "Soccer"},
	Music:  []string{"Blues", "Classical", "Country", "Easy Listening", "Electronic", "Hip Pop", "Jazz", "Opera", "Soul", "Rock"},
	Cities: []string{"Atlanta", "Bangkok", "Beijing", "London", "Paris", "Singapore", "New York", "Istanbul", "Dubai", "Taipei"},
	Books:  []string{"Journey to the West", "The Lord of the Rings", "In Search of Lost Time", "Don Quixote", "Ulysses", "The Great Gatsby", "Moby Dick", "Hamlet", "War and Peace", "The Odyssey"},
	Movies: []string{"The Godfather", "The Shawshank Redemption", "Schindler's List", "Raging Bull", "Casablanca", "Citizen Kane", "Gone with the Wind", "The Wizard of Oz", "One Flew Over the Cuckoo's Nest", "Lawrence of Arabia"},
}

// MongoConn -
type MongoConn struct {
	uri      string
	ssl      bool
	sslCA    string
	dbName   string
	tps      int
	filename string
	verbose  bool
	cleannUp bool
}

var simDocs []bson.M

// initialize an array of documents for simulation test.  If a template is available
// read the sample json and replace them with random values.  Otherwise, use the demo
// example.
func (m MongoConn) initSimDocs() {
	rand.Seed(time.Now().Unix())
	total := 512
	if m.filename == "" {
		for len(simDocs) < total {
			simDocs = append(simDocs, GetRandomDoc())
		}
		return
	}

	sdoc := GetDocByTemplate(m.filename, true)
	bytes, _ := json.MarshalIndent(sdoc, "", "   ")
	if m.verbose {
		fmt.Println(string(bytes))
	}
	doc := make(map[string]interface{})
	json.Unmarshal(bytes, &doc)

	for len(simDocs) < total {
		ndoc := make(map[string]interface{})
		traverseDocument(&ndoc, doc, false)
		delete(ndoc, "_id")
		simDocs = append(simDocs, ndoc)
	}
}

// New - Constructor
func New(uri string, ssl bool, sslCA string, dbName string, tps int, filename string, verbose bool, cleanUp bool) MongoConn {
	m := MongoConn{uri, ssl, sslCA, dbName, tps, filename, verbose, cleanUp}
	m.initSimDocs()
	return m
}

func unique(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us[:3]
}

// GetRandomDoc -
func GetRandomDoc() bson.M {
	var filler1, filler2 bytes.Buffer
	for i := 0; i < 80/len("simagix."); i++ {
		filler1.WriteString("simagix.")
		filler2.WriteString("mongodb.")
	}
	var n = len(favorites.Sports)
	favoriteSports := []string{favorites.Sports[rand.Intn(n)], favorites.Sports[rand.Intn(n)], favorites.Sports[rand.Intn(n)]}
	favoriteSports = unique(append(favoriteSports, favorites.Sports[0:3]...))
	n = len(favorites.Music)
	favoriteMusic := []string{favorites.Music[rand.Intn(n)], favorites.Music[rand.Intn(n)], favorites.Music[rand.Intn(n)]}
	favoriteMusic = unique(append(favoriteMusic, favorites.Music[0:3]...))
	n = len(favorites.Cities)
	favoriteCities := []string{favorites.Cities[rand.Intn(n)], favorites.Cities[rand.Intn(n)], favorites.Cities[rand.Intn(n)]}
	favoriteCities = unique(append(favoriteCities, favorites.Cities[0:3]...))
	n = len(favorites.Books)
	favoriteBooks := []string{favorites.Books[rand.Intn(n)], favorites.Books[rand.Intn(n)], favorites.Books[rand.Intn(n)]}
	favoriteBooks = unique(append(favoriteBooks, favorites.Books[0:3]...))
	n = len(favorites.Movies)
	favoriteMovies := []string{favorites.Movies[rand.Intn(n)], favorites.Movies[rand.Intn(n)], favorites.Movies[rand.Intn(n)]}
	favoriteMovies = unique(append(favoriteMovies, favorites.Movies[0:3]...))
	favoritesList := []bson.M{
		bson.M{"sport": favoriteSports[0], "music": favoriteMusic[0], "city": favoriteCities[0], "book": favoriteBooks[0], "movie": favoriteMovies[0]},
		bson.M{"sport": favoriteSports[1], "music": favoriteMusic[1], "city": favoriteCities[1], "book": favoriteBooks[1], "movie": favoriteMovies[1]},
		bson.M{"sport": favoriteSports[2], "music": favoriteMusic[2], "city": favoriteCities[2], "book": favoriteBooks[2], "movie": favoriteMovies[2]},
	}

	doc := bson.M{
		"favorites": bson.M{
			"sports": favoriteSports, "sport": favoriteSports[0],
			"musics": favoriteMusic, "music": favoriteMusic[0],
			"cities": favoriteCities, "city": favoriteCities[0],
			"books": favoriteBooks, "book": favoriteBooks[0],
			"movies": favoriteMovies, "movie": favoriteMovies[0],
		},
		"favoritesList":  favoritesList,
		"favoriteSports": favoriteSports,
		"favoriteMusics": favoriteMusic,
		"favoriteCities": favoriteCities,
		"favoriteBooks":  favoriteBooks,
		"favoriteMovies": favoriteMovies,
		"favoriteSport":  favoriteSports[0],
		"favoriteMusic":  favoriteMusic[0],
		"favoriteCity":   favoriteCities[0],
		"favoriteBook":   favoriteBooks[0],
		"favoriteMovie":  favoriteMovies[0],
		"filler1":        filler1.String(),
		"filler2":        filler2.String(),
		"number":         rand.Intn(1000),
		"numbers":        []int{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)},
		"ts":             time.Now(),
	}
	return doc
}

// PopulateData - Insert docs to evaluate performance/bandwidth
// {
//	favorites: {
//		sports: []
//		cities: []
//	}
//	favoriteSports: []
//	favoriteSports1
//	favoriteSports2
//	favoriteSports3
// }
func (m MongoConn) PopulateData() {
	s := 0
	batchSize := 20
	if m.tps < batchSize {
		batchSize = m.tps
	}
	for s < 57 { // 3 seconds less of a minute
		s++
		session, err := GetSession(m.uri, m.ssl, m.sslCA)
		if err == nil {
			session.SetMode(mgo.Primary, true)
			c := session.DB(m.dbName).C(CollectionName)
			bt := time.Now()
			bulk := c.Bulk()

			for i := 0; i < m.tps; i += batchSize {
				var contentArray []interface{}
				for n := 0; n < batchSize; n++ {
					contentArray = append(contentArray, simDocs[i%len(simDocs)])
				}
				bulk.Insert(contentArray...)
				_, err := bulk.Run()
				if err != nil {
					log.Println(err)
					session.Close()
					break
				}
			}

			t := time.Now()
			elapsed := t.Sub(bt)
			if elapsed.Seconds() > time.Second.Seconds() {
				x := math.Floor(elapsed.Seconds())
				s += int(x)
				elapsed = time.Duration(elapsed.Seconds() - x)
			}
			et := time.Second.Seconds() - elapsed.Seconds()
			if et > 0 {
				time.Sleep(time.Millisecond * time.Duration(int(1000*et)))
			} else if m.verbose {
				fmt.Println("Populate", "TPS Overflows", et)
			}
			session.Close()
		} else {
			time.Sleep(time.Second)
		}
	}
}

// Simulate simulates CRUD for load tests
func (m MongoConn) Simulate(duration int) {
	var schema = Schema{}
	results := []bson.M{}
	change := bson.M{"$set": bson.M{"ts": time.Now()}}
	burstBegin := time.NewTimer(1 * time.Minute)
	waitms := 2
	totalTPS := m.tps / 2
	isTeardown := false

	go func() {
		<-burstBegin.C
		if m.verbose {
			println("burst begins")
		}
		totalTPS = m.tps
	}()
	burstEnd := time.NewTimer(time.Duration(duration-1) * time.Minute)
	go func() {
		<-burstEnd.C
		if m.verbose {
			println("burst ends")
		}
		totalTPS = m.tps / 2
		isTeardown = true
	}()

	for {
		session, err := GetSession(m.uri, m.ssl, m.sslCA)
		session.SetMode(mgo.Primary, true)
		if err == nil {
			c := session.DB(m.dbName).C(CollectionName)
			var book string
			var city string
			var movie string
			beginTime := time.Now()
			for i := 0; i < totalTPS; i++ {
				doc := simDocs[i%len(simDocs)]
				bytes, _ := json.Marshal(doc)
				json.Unmarshal(bytes, &schema)
				city = schema.FavoriteCity
				book = schema.FavoriteBook
				movie = schema.FavoriteMovie

				if isTeardown {
					c.RemoveAll(bson.M{"_search": strconv.FormatInt(int64(i)*100, 16)})
					time.Sleep(time.Millisecond * time.Duration(waitms))
				} else {
					_id := bson.NewObjectIdWithTime(time.Now())
					doc["_search"] = strconv.FormatInt(int64(i)*100, 16)
					c.Upsert(_id, doc)
					time.Sleep(time.Millisecond * time.Duration(waitms))
					if m.filename == "" {
						// c.Find(bson.M{"favoriteCity": city}).Sort("favoriteCity").Limit(512).All(&results)
						c.Find(bson.M{"favoriteCity": city}).Limit(20).All(&results)
						c.Find(bson.M{"favoriteCity": city, "favoriteBook": book}).One(&results)
						c.Update(bson.M{"_id": _id}, change)
						// c.Find(bson.M{"favoriteCity": city, "favoriteBook": book, "FavoriteMovie": movie}).One(&results)
						c.Find(bson.M{"favoritesList": bson.M{"$elemMatch": bson.M{"movie": movie}}}).One(&results)
						// c.Find(bson.M{"favoritesList": bson.M{"$elemMatch": bson.M{"book": book}}}).Limit(100).All(&results)
					} else {
						if i == 20 {
							c.Find(bson.M{"COLLSCAN": doc["_search"]}).One(&results) // simulate COLLSCAN
						} else if (i % 21) == 20 {
							c.Find(bson.M{"_search": doc["_search"]}).Sort("_search").Limit(10).All(&results)
						} else {
							c.Find(bson.M{"_id": _id}).One(&results)
							time.Sleep(time.Millisecond * time.Duration(waitms))
							if (i % 2) == 0 {
								c.Update(bson.M{"_id": _id}, change)
							} else {
								c.Remove(bson.M{"_id": _id})
							}
						}
					}
				}
				if (i % 21) == 20 {
					seconds := 1 - time.Now().Sub(beginTime).Seconds()
					if seconds < 0 {
						if m.verbose {
							fmt.Println("Simulate", "TPS overflows and break out of loop", seconds)
						}
						break
					}
				}
			}
			seconds := 1 - time.Now().Sub(beginTime).Seconds()
			if seconds > 0 {
				time.Sleep(time.Millisecond * time.Duration(int(1000*seconds)))
			} else if m.verbose {
				fmt.Println("Simulate", "TPS overflows", seconds)
			}
			if m.filename == "" {
				c.Find(bson.M{"favoritesList": bson.M{"$elemMatch": bson.M{"book": book}}}).Sort("favoriteCity").Limit(20).All(&results)
			}
			session.Close()
		}
	}
}

// Cleanup drops the temp database
func (m MongoConn) Cleanup() {
	if m.cleannUp == false {
		return
	}
	log.Println("cleanup", m.uri)
	session, err := GetSession(m.uri, m.ssl, m.sslCA)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	time.Sleep(2 * time.Second)
	log.Println("dropping collection", m.dbName, CollectionName)
	session.DB(m.dbName).C(CollectionName).DropCollection()
	log.Println("dropping database", m.dbName)
	session.DB(m.dbName).DropDatabase()
}

// CreateIndexes creates indexes
func (m MongoConn) CreateIndexes() {
	session, _ := GetSession(m.uri, m.ssl, m.sslCA)
	defer session.Close()
	c := session.DB(m.dbName).C(CollectionName)

	if m.filename == "" {
		c.EnsureIndexKey("favoriteCity")
	}
	c.EnsureIndexKey("_search")
}
