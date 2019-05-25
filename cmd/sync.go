package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/DefinitelyNotAGoat/go-explore/mongodb"
	"github.com/DefinitelyNotAGoat/go-explore/sync"
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"github.com/spf13/cobra"
)

func newSync() *cobra.Command {
	var URL string
	var mongo string
	var db string
	var refresh int

	var sync = &cobra.Command{
		Use:   "sync",
		Short: "sync is the main driver for syncing the explorer into mongodb",
		Run: func(cmd *cobra.Command, args []string) {
			mg, err := mongodb.NewMongoService(mongo, db)
			if err != nil {
				fmt.Printf("could not connect to mogno: %v", err)
				os.Exit(1)
			}

			gt, err := goTezos.NewGoTezos(URL)
			if err != nil {
				fmt.Printf("could not connect to tezos node: %v", err)
				os.Exit(1)
			}
			syncer := sync.NewSync(time.Second*time.Duration(refresh), gt, mg)
			errch := make(chan error, 100)
			quit := make(chan int)
			syncer.Sync(quit, errch)
			mg.StateCheck(errch)

			for err := range errch {
				fmt.Println(err)
			}
		},
	}
	sync.PersistentFlags().IntVarP(&refresh, "refresh", "r", 10, "refresh interval in seconds for sync")
	sync.PersistentFlags().StringVarP(&db, "database", "d", "tezos", "databse name")
	sync.PersistentFlags().StringVarP(&mongo, "mongo", "n", "mongodb://localhost:27017", "mondo dial address")
	sync.PersistentFlags().StringVarP(&URL, "node", "u", "http://127.0.0.1:8732", "address to the node to query (default http://127.0.0.1:8732)(e.g. https://mainnet-node.tzscan.io:443)")
	return sync
}
