package main

import (
	"fmt"
	"github.com/aybabtme/benchkit"
	"github.com/aybabtme/benchkit/benchplot"
	"github.com/aybabtme/goamz/s3"
	"github.com/aybabtme/parajson"
	"github.com/aybabtme/trie"
	"github.com/aybabtme/uniplot/spark"
	"github.com/codegangsta/cli"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func memplotCommand() ([]cli.Flag, func(*cli.Context)) {

	fileFlag := cli.StringFlag{Name: "file", Usage: "file containing the keys to read from"}
	typeFlag := cli.StringFlag{Name: "type", Usage: "type of set to benchmark"}
	widthFlag := cli.Float64Flag{Name: "width", Value: 8.0, Usage: "width of the plot to render"}
	heightFlag := cli.Float64Flag{Name: "height", Value: 6.0, Usage: "height of the plot to render"}

	flags := []cli.Flag{fileFlag, typeFlag, widthFlag, heightFlag}

	return flags, func(c *cli.Context) {
		var (
			settype  = c.String(typeFlag.Name)
			filename = c.String(fileFlag.Name)
			width    = c.Float64(widthFlag.Name)
			height   = c.Float64(heightFlag.Name)
		)

		hadError := true
		switch {
		case settype == "":
			log.Println("Missing value for", typeFlag.Name)
		case filename == "":
			log.Println("Missing value for", fileFlag.Name)
		default:
			hadError = false
		}
		if hadError {
			cli.ShowAppHelp(c)
			return
		}

		set, setName, err := setImpl(settype)
		if err != nil {
			log.Printf("Bad type flag: %v", err)
			return
		}

		file, err := os.Open(filename)
		if err != nil {
			log.Printf("opening=%q\terror=%v", filename, err)
			return
		}
		defer func() { _ = file.Close() }()

		keys, err := decodeKeys(spark.Reader(file))
		if err != nil {
			log.Printf("decoding=%q\terror=%v", filename, err)
			return
		}
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		log.Printf("baseline-in-use=%s", humanize.Bytes(mem.HeapAlloc-mem.HeapReleased))

		log.Printf("key-count=%d", len(keys))

		keyfactor := 100

		log.Printf("benchmarking...")
		start := time.Now()
		n := len(keys)/keyfactor + 2
		results := benchkit.Bench(benchkit.Memory(n)).Each(func(each benchkit.BenchEach) {
			lastj := 0
			each.Before(0)
			for j, key := range keys {

				if j%keyfactor == 0 {
					each.After(lastj)
					each.Before(lastj + 1)
					lastj++
				}
				set.Add(key.Key)

			}
			each.After(lastj)
		}).(*benchkit.MemResult)
		log.Printf("done in %v", time.Since(start))

		plottitle := fmt.Sprintf("Insertion of %d keys in a %s", len(keys), setName)
		log.Printf("plotting %q", plottitle)
		p, err := benchplot.PlotMemory(plottitle, fmt.Sprintf("times %d keys", keyfactor), results)
		if err != nil {
			log.Printf("plotting=%q\terror=%v", plottitle, err)
			return
		}

		base := filepath.Base(filename)
		ext := filepath.Ext(base)
		cleanname := base[:len(base)-len(ext)]

		plotname := fmt.Sprintf("%s_%s.svg", cleanname, setName)
		log.Printf("saving to %q (%gx%g)", plotname, width, height)
		if err := p.Save(width, height, plotname); err != nil {
			log.Printf("saving=%q\terror=%v", plotname, err)
		}
	}
}

func setImpl(settype string) (trie.Set, string, error) {
	switch settype {
	// case "trie":
	// 	return trie.NewTrieSet('#', 96), "TrieSet", nil
	case "gomap":
		return trie.NewGoMapSet(8096), "GoMapSet", nil
	case "ternary":
		return trie.NewTernarySet(), "TernarySet", nil
	}
	return nil, "", fmt.Errorf("%q is not a valid set type: valids are gomap and ternary", settype)
}

func decodeKeys(r io.Reader) ([]s3.Key, error) {
	var keys []s3.Key
	keyc, errc := parajson.Decode(r, runtime.NumCPU(), func() interface{} {
		return &s3.Key{}
	})

	for proto := range keyc {
		key := proto.(*s3.Key)
		keys = append(keys, *key)
	}

	var err error
	for e := range errc {
		err = e
	}
	return keys, err
}
