package cmd

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "Run this if u wanna list all todos",
	Run: func(cmd *cobra.Command, args []string) {
		dbInstance := createConnCreateBucket()
		defer dbInstance.Close()

		err2 := dbInstance.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasksTable"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%d. %s\n", btoi(k), v)
			}
			return nil
		})
		if err2 != nil {
			log.Fatal(err2)
		}
	},
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
