package db

import{
	"gopkg.in/mgo.v2"
	"config"
	"log"
}

func Connect() *Connection {
	
	session, err := mgo.Dial(Config.DbConnect)

	if err != nil {
    	log.Fatal("Db connection error: ", err)
  	}

	return session.DB(Config.DbCollection)

}

func Close(s *Session) {
	
	s.Close()
}