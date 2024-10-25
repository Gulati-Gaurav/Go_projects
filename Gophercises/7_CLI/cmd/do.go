package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var done = &cobra.Command{
	Use:   "done",
	Short: "Run this if u wanna mark a todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		dbInstance := createConnCreateBucket()
		defer dbInstance.Close()

		key, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		err2 := dbInstance.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasksTable"))
			return b.Delete(itob(key))
		})
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("deleted")
	},
}

func createConnCreateBucket() *bolt.DB {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	dbInstance, err := bolt.Open(homeDir+"/tasks.db", 0600, &bolt.Options{Timeout: 4 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	dbInstance.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasksTable"))
		return err
	})
	return dbInstance
}
