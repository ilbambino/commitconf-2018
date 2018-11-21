package main

import (
	"math/rand"
	"time"

	ui "github.com/gizak/termui"
)

const greedy = 80

func main() {

	rand.Seed(time.Now().UnixNano())

	err := ui.Init()
	if err != nil {
		panic(err)
	}

	ui.Handle("q", "<Insert>", func(ui.Event) {
		ui.StopLoop()
	})

	defer ui.Close()

	impressionsBC := ui.NewBarChart()
	impressionsBC.BorderLabel = "Resources Impressions"
	impressionsBC.Width = 66
	impressionsBC.Height = 20
	impressionsBC.TextColor = ui.ColorGreen
	impressionsBC.BarColor = ui.ColorRed
	impressionsBC.NumColor = ui.ColorYellow
	impressionsBC.BarGap = 3
	impressionsBC.BarWidth = 10

	consumptionsBC := ui.NewBarChart()
	consumptionsBC.BorderLabel = "Resources Consumptions"
	consumptionsBC.Width = 66
	consumptionsBC.Height = 20
	consumptionsBC.TextColor = ui.ColorGreen
	consumptionsBC.BarColor = ui.ColorBlue
	consumptionsBC.NumColor = ui.ColorYellow
	consumptionsBC.BarGap = 3
	consumptionsBC.BarWidth = 10

	valueBC := ui.NewBarChart()
	valueBC.BorderLabel = "Resources Value * 100"
	valueBC.Width = 66
	valueBC.Height = 20
	valueBC.TextColor = ui.ColorGreen
	valueBC.BarColor = ui.ColorBlue
	valueBC.NumColor = ui.ColorYellow
	valueBC.BarGap = 3
	valueBC.BarWidth = 10

	resources := initialResources()

	impressionsBC.DataLabels = getNames(resources)
	consumptionsBC.DataLabels = getNames(resources)
	valueBC.DataLabels = getNames(resources)
	valueBC.Data = getValues(resources)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, impressionsBC),
			ui.NewCol(6, 0, consumptionsBC)),
		ui.NewRow(
			ui.NewCol(6, 0, valueBC)),
	)

	// calculate layout
	ui.Body.Align()

	ui.Handle("a", "<Insert>", func(ui.Event) {
		resources = append(resources, randomResource())
		impressionsBC.DataLabels = getNames(resources)
		consumptionsBC.DataLabels = getNames(resources)
		valueBC.DataLabels = getNames(resources)
		valueBC.Data = getValues(resources)
	})

	ticker := time.NewTicker(time.Millisecond * 30)
	go func() {
		for {
			user := randomUser()
			if rand.Int31n(100) < 100-greedy {
				// choose randomly
				chosen := &resources[rand.Intn(len(resources))]
				chosen.Use()
			} else {
				// allocate
				maxScore := 0.0
				maxResource := &resources[0]

				for i, resource := range resources {
					score := scoreCalculation(resource, user)
					if score > maxScore {
						maxScore = score
						maxResource = &resources[i]
					}
				}
				maxResource.Use()
			}

			impressionsBC.Data = getImpressions(resources)
			consumptionsBC.Data = getConsumptions(resources)
			ui.Render(impressionsBC, consumptionsBC, valueBC)
			<-ticker.C
		}
	}()

	ui.Loop()
}

// helpers ********************

func getImpressions(resources []Resource) []int {

	impressions := make([]int, len(resources))
	for i, resource := range resources {
		impressions[i] = resource.Impressions
	}
	return impressions
}

func getConsumptions(resources []Resource) []int {

	impressions := make([]int, len(resources))
	for i, resource := range resources {
		impressions[i] = resource.Consumptions
	}
	return impressions
}

func getNames(resources []Resource) []string {

	impressions := make([]string, len(resources))
	for i, resource := range resources {
		impressions[i] = resource.Name
	}
	return impressions
}

func getValues(resources []Resource) []int {

	impressions := make([]int, len(resources))
	for i, resource := range resources {
		impressions[i] = int(resource.Value * 100)
	}
	return impressions
}
