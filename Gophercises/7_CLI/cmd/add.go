package cmd

import (
	"encoding/binary"
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Run this if u wanna add a todo",
	Run: func(cmd *cobra.Command, args []string) {
		data := strings.Join(args, " ")

		dbInstance := createConnCreateBucket()
		defer dbInstance.Close()

		err2 := dbInstance.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasksTable"))
			id, _ := b.NextSequence()
			err := b.Put(itob(int(id)), []byte(data))

			return err
		})

		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("Added")
	},
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func init() {
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(done)
	rootCmd.AddCommand(list)
}
