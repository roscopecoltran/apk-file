package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"github.com/agrison/go-tablib"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/jessfraz/apk-file/version"
)

const (
	// BANNER is what is printed for help/info output
	BANNER = `             _          __ _ _
  __ _ _ __ | | __     / _(_) | ___
 / _` + "`" + ` | '_ \| |/ /____| |_| | |/ _ \
| (_| | |_) |   <_____|  _| | |  __/
 \__,_| .__/|_|\_\    |_| |_|_|\___|
      |_|

 Search apk package contents via the command line.
 Version: %s

`
	alpineContentsSearchURI = "https://pkgs.alpinelinux.org/contents"
)

type fileInfo struct {
	path, pkg, branch, repo, arch string
}

var (
	arch string
	repo string
	output string
	result string
	prefixPath string 
	filename string

	debug bool
	vrsn  bool
	save bool

	validOutput = []string{"markdown", "csv", "yaml", "json", "xlsx", "xml", "tsv", "mysql", "postgres", "html", "ascii"}
	validArches = []string{"x86", "x86_64", "armhf"}
	validRepos  = []string{"main", "community", "testing"}
)

func init() {
	// Parse flags
	flag.StringVar(&arch, "arch", "", "arch to search for ("+strings.Join(validArches, ", ")+")")
	flag.StringVar(&repo, "repo", "", "repository to search in ("+strings.Join(validRepos, ", ")+")")
	flag.StringVar(&output, "output", "", "output results with  ("+strings.Join(validOutput, ", ")+") format.")
	flag.StringVar(&prefixPath, "./output", "results", "output results to prefix_path (default: ./output).")
	flag.StringVar(&filename, "filename", "results", "output results to filename: (default: ./results.[FORMAT]).")

	flag.BoolVar(&save, "save", true, "save output results to the output_file.[FORMAT].")
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("apk-file version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	// Set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if arch != "" && !stringInSlice(arch, validArches) {
		logrus.Fatalf("%s is not a valid arch", arch)
	}

	if repo != "" && !stringInSlice(repo, validRepos) {
		logrus.Fatalf("%s is not a valid repo", repo)
	}
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	if flag.NArg() < 1 {
		logrus.Fatal("must pass a file to search for.")
	}

	f, p := getFileAndPath(flag.Arg(0))

	query := url.Values{
		"file":   {f},
		"path":   {p},
		"branch": {""},
		"repo":   {repo},
		"arch":   {arch},
	}

	uri := fmt.Sprintf("%s?%s", alpineContentsSearchURI, query.Encode())
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		logrus.Fatalf("requesting %s failed: %v", uri, err)
	}

	files := getFilesInfo(doc)

	ds 		:= tablib.NewDataset([]string{"file", "package", "branch", "repository", "architecture"})
	
	for _, f := range files {
		// https://github.com/agrison/go-tablib
		ds.AppendValues(f.path, f.pkg, f.branch, f.repo, f.arch)
	}

	// "markdown", "csv", "yaml", "json", "xlsx", "xml", "tsv", "mysql", "postgres", "html", "ascii"

	switch output {

		case "csv":
			result, _ := ds.CSV()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		case "tsv":
			result, _ := ds.TSV()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		case "yaml":
			result, _ := ds.YAML()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		case "json":
			result, _ := ds.JSON()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		case "xlsx":
			result, _ := ds.XLSX()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		case "xml":
			result, _ := ds.XML()
			if save == true {
				if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
				    fmt.Println(err)
				}
			}
			fmt.Println(result)
		/*
		case "mysql":
			result, _ := ds.MySQL()
			fmt.Println(result)
			if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
			    fmt.Println(err)
			}
			fmt.Println(result)
		case "postgres":
			result := ds.Postgres()
			fmt.Println(result)
			if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
			    fmt.Println(err)
			}
			fmt.Println(result)
		*/
		case "html":
			result, _ := ds.XLSX()
			fmt.Println(result)
			if result.WriteFile(prefixPath+"/"+filename+"."+output, 0644) != nil {
			    fmt.Println(err)
			}
			fmt.Println(result)
		case "ascii":
		default:
			ascii := ds.Tabular("grid" /* tablib.TabularGrid */)	
			fmt.Println(ascii)
	}


}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func getFilesInfo(d *goquery.Document) []fileInfo {
	files := []fileInfo{}
	d.Find(".table tr:not(:first-child)").Each(func(j int, l *goquery.Selection) {
		f := fileInfo{}
		rows := l.Find("td")
		rows.Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				f.path = s.Text()
			case 1:
				f.pkg = s.Text()
			case 2:
				f.branch = s.Text()
			case 3:
				f.repo = s.Text()
			case 4:
				f.arch = s.Text()
			default:
				logrus.Warn("Unmapped value for column %d with value %s", i, s.Text())
			}
		})
		files = append(files, f)
	})
	return files
}

func getFileAndPath(arg string) (file string, dir string) {
	file = "*" + path.Base(arg) + "*"
	dir = path.Dir(arg)
	if dir != "" && dir != "." {
		dir = "*" + dir
		file = strings.TrimPrefix(file, "*")
	} else {
		dir = ""
	}
	return file, dir
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
