package main

import (
	"log"
	"os"

	"github.com/RobinHung/ubuntu-esm-scraper/pkg"
	"github.com/urfave/cli"
)

type ubuntuEsmPatches struct {
	Title     string
	Date      string
	CVEs      []string
	CVEScores []cve
}

type cve struct {
	ID                  string
	BaseScore           string
	Severity            string
	ImpactScore         string
	ExploitabilityScore string
}

func newApp(appname string) *cli.App {
	cli.AppHelpTemplate = `NAME:
	{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}
USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}
VERSION:
	{{.Version}}{{end}}{{end}}{{if .Description}}
DESCRIPTION:
	{{.Description}}{{end}}{{if len .Authors}}
AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
	{{range $index, $author := .Authors}}{{if $index}}
	{{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}
OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}
`
	app := cli.NewApp()
	app.Name = "USNTracker"
	app.Usage = "Ubuntu Security Notices Tracker"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "target, t",
			Value: "",
			Usage: "ubuntu version to track",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output file name, only supports json or csv file format",
		},
		cli.BoolFlag{
			Name:  "display, d",
			Usage: "display the result",
		},
	}
	// app.Commands = []cli.Command{}
	app.Action = pkg.Run
	return app
}

func main() {
	if err := newApp("USNTracker").Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("usn.ubuntu.com", "people.canonical.com"),
	// 	// colly.MaxDepth(3),
	// )

	// c2 := colly.NewCollector(
	// 	colly.AllowedDomains("nvd.nist.gov"),
	// 	colly.AllowURLRevisit(),
	// )

	// var cves [][]string
	// patches := []ubuntuEsmPatches{}
	// var dates []string
	// var titles []string
	// pageCount := 0

	// vulns := []cve{}
	// vv := [][]cve{}

	// c.OnHTML("em", func(e *colly.HTMLElement) {
	// 	dates = append(dates, e.Text)
	// })

	// c.OnHTML("body.home", func(e *colly.HTMLElement) {
	// 	var sub []string
	// 	vv = append(vv, vulns)

	// 	e.ForEach(".p-heading--four", func(_ int, el *colly.HTMLElement) {
	// 		titles = append(titles, el.Text)

	// 		el.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
	// 			link := elem.Attr("href")
	// 			if !strings.HasPrefix(link, "https://usn.ubuntu.com/") {
	// 				return
	// 			}
	// 			err := c.Visit(elem.Request.AbsoluteURL(link))
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		})
	// 	})

	// 	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
	// 		el.ForEach("a[href]", func(_ int, elem *colly.HTMLElement) {
	// 			link := elem.Attr("href")

	// 			if strings.HasPrefix(link, "https://people.canonical.com/~ubuntu-security/cve/") || strings.HasPrefix(elem.Text, "CVE-") {
	// 				sub = append(sub, elem.Text)

	// 				cveScoringLink := "https://nvd.nist.gov/vuln/detail/" + elem.Text
	// 				err := c2.Visit(cveScoringLink)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 			}
	// 		})
	// 	})

	// 	cves = append(cves, sub)
	// 	pageCount++
	// })

	// c2.OnHTML("div#p_lt_WebPartZone1_zoneCenter_pageplaceholder_p_lt_WebPartZone1_zoneCenter_VulnerabilityDetail_VulnDetailFormPanel", func(e *colly.HTMLElement) {
	// 	vuln := cve{}
	// 	e.ForEach("span", func(_ int, el *colly.HTMLElement) {
	// 		attr := el.Attr("data-testid")
	// 		if attr == "page-header-vuln-id" {
	// 			vuln.ID = strings.TrimSpace(el.Text)
	// 		}
	// 		if attr == "vuln-cvssv3-base-score" {
	// 			vuln.BaseScore = strings.TrimSpace(el.Text)
	// 		}
	// 		if attr == "vuln-cvssv3-base-score-severity" {
	// 			vuln.Severity = strings.TrimSpace(el.Text)
	// 		}
	// 		if attr == "vuln-cvssv3-impact-score" {
	// 			vuln.ImpactScore = strings.TrimSpace(el.Text)
	// 		}
	// 		if attr == "vuln-cvssv3-exploitability-score" {
	// 			vuln.ExploitabilityScore = strings.TrimSpace(el.Text)
	// 		}
	// 	})

	// 	vv[pageCount] = append(vv[pageCount], vuln)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })
	// c2.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// c.Visit("https://usn.ubuntu.com/releases/ubuntu-14.04-esm/")

	// for i, t := range titles {
	// 	patch := ubuntuEsmPatches{}
	// 	patch.Title = t
	// 	patch.Date = dates[i]
	// 	patch.CVEs = cves[i]
	// 	patch.CVEScores = vv[i]
	// 	patches = append(patches, patch)
	// }

	// // Print out json data
	// // enc := json.NewEncoder(os.Stdout)
	// // enc.Encode(patches)

	// jsonData, err := json.Marshal(patches)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = ioutil.WriteFile("trusty-esm-patches.json", jsonData, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}