package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {

	fmt.Println("---Grade Result---")
	reader := bufio.NewReader(os.Stdin)
	var names []string
	var scores []float64
	mapGradeSCore := map[string]float64 {}
	//mapGradeAvg := map[string]float64 {}
	for{
		fmt.Println("1. Add Student Score")
		fmt.Println("2. Update Existing Score")
		fmt.Println("3. View Class Average")
		fmt.Println("4. Show Top 3 scores") 
		fmt.Println("5. Generate Full Report")
		fmt.Println("6. Show Grade Distribution")
		fmt.Println("7. Exit")
		choice := getInput(reader)

		switch choice{
		case "1" :
			fmt.Print("name:")
			name :=getInput(reader)
			fmt.Print("score:")
			scoreStr := getInput(reader)
			score,_ := strconv.ParseFloat(scoreStr,64)
			names = append(names,name)
			scores =append(scores,score)
			mapGradeSCore[name] =score
			fmt.Println("name and score successfully added")	
		
		case "2" :	
			fmt.Print("Enter the name of the student whose score should be updated:")
			nameupdate := getInput(reader)
			fmt.Print("Enter New Score:")
			newscoreStr := getInput(reader)
			newScore,_ := strconv.ParseFloat(newscoreStr,64)
			if _, exists := mapGradeSCore[nameupdate]; exists {
			mapGradeSCore[nameupdate] = newScore
			for i, name := range names {
				if name == nameupdate {
					scores[i] = newScore
					break
				}
			}
			fmt.Println("Successfully Updated")
			} else {
			fmt.Println("Name not found")
			}
			
		case "3":
			if len(scores) == 0 {
				fmt.Println("No scores available to calculate average.")
				break
			}
			var Avgscore float64 
			for _,score := range scores{
				Avgscore += score 
			}
			Avgscore = Avgscore / float64(len(scores))
			fmt.Println("The average score is ",Avgscore)

		
		case "4":
			if len(scores) == 0 {
				fmt.Println("No scores available.")
				break
			}
			sortedScores := make([]float64, len(scores))
			copy(sortedScores, scores)
			sort.Float64s(sortedScores)
			fmt.Println("Highest scores")
			if len(sortedScores)<3{
				for i := len(sortedScores)-1;i>=0;i--{
					fmt.Println(sortedScores[i])
				}
			} else {
				for i := len(sortedScores)-1;i>=len(sortedScores)-3;i--{
					fmt.Println(sortedScores[i])
				}
			}
		case "5":
			if len(mapGradeSCore) == 0 {
				fmt.Println("No data to display.")
				break
			}

			// Sort names
			sortedNames := make([]string, len(names))
			copy(sortedNames, names)
			sort.Strings(sortedNames)
			
			for _,name := range sortedNames {
				score := mapGradeSCore[name]
				grade := ""
				if score >= 91 {
					grade = "A"
				} else if score >= 81 {
					grade = "B"
				} else if score >= 71 {
					grade = "C"
				} else if score >= 61 {
					grade = "D"
				} else if score >= 51 {
					grade = "E"
				} else {
					grade = "F"
}
			
				fmt.Printf("%v: %v (%v)\n",name,score,grade)
			}
				var Avgscore float64 
				for _,score := range scores{
					Avgscore += score 
				}
				Avgscore1 := Avgscore/float64(len(scores))
				fmt.Println("Class Average ",Avgscore1)
				sortedScores := make([]float64, len(scores))
				copy(sortedScores, scores)
				sort.Float64s(sortedScores)
				fmt.Println("highest score:",sortedScores[len(sortedScores)-1])
				fmt.Println("Lowest score:",sortedScores[0])
				passcount := 0
				for _,passScore := range scores{
					if passScore >=40{
						passcount+=1
					}
				}
				fmt.Println("Pass COunt:",passcount)
			case "6":
				gradeA := 0
				gradeB := 0
				gradeC := 0
				gradeD := 0
				gradeE := 0
				gradeF := 0
				for _,score := range mapGradeSCore {
				if score >= 91 {
					gradeA+=1
				} else if score >= 81 {
					gradeB+=1
				} else if score >= 71 {
					gradeC+=1
				} else if score >= 61 {
					gradeD+=1
				} else if score >= 51 {
					gradeE+=1
				} else {
					gradeF+=1
				}
				}
				fmt.Print("Grade Distribution: A:",gradeA)
				fmt.Print(" B:",gradeB)
				fmt.Print(" C:",gradeC)
				fmt.Print(" D:",gradeD)
				fmt.Print(" E:",gradeE)
				fmt.Println(" F:",gradeF)
			case "7":
				fmt.Println("Exiting program...")
    			return
				
		}
	}
}

func getInput(reader *bufio.Reader) string{
	input,_ :=reader.ReadString('\n')
	inputStr :=strings.TrimSpace(input)
	return inputStr
}