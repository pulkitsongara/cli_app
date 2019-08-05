package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var email_id, password, result string
var input_val int32

func main() {
	loadMainMenu()
}
func loadMainMenu() {
	registration()
}

func registration() {
	fmt.Println("\nPress 1 to login.")
	fmt.Println("Press 2 to sign-up.")
	fmt.Println("Press 9 to exit from app.")
	fmt.Scanf("%d", &input_val)
	switch input_val {
	case 1:
		login()
	case 2:
		signup()
	case 9:
		os.Exit(500)
	default:
		fmt.Println("Not a valid choice.")
		registration()
	}
}

func userOperations() {
	fmt.Println("\nPress 1 to get the list of all journals.")
	fmt.Println("Press 2 to create new journal.")
	fmt.Println("Press 8 to go main menu.")
	fmt.Println("Press 9 to exit from app.")
	fmt.Scanf("%d", &input_val)
	switch input_val {
	case 1:
		getList()
	case 2:
		createJournal()
	case 8:
		loadMainMenu()
	case 9:
		os.Exit(500)
	default:
		fmt.Println("Not a valid choice.")
		registration()
	}
}

func signup() {
	fmt.Println("\nPlease enter email_id")
	fmt.Scanf("%s", &email_id)
	fmt.Println("Please enter Password")
	fmt.Scanf("%s", &password)
	file, err := os.OpenFile("user_credentials.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error while opening file")
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount > 10 {
		fmt.Println("Number of user limit exceeded")
	} else {
		file1, err := os.OpenFile("user_credentials.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("File is not opening")
		}
		scanner := bufio.NewScanner(file1)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Compare(line, email_id) == 0 {
				fmt.Println(email_id, password)

				fmt.Println("Email ID already exist.")
				registration()
				break
			} else {
				_, err = file.WriteString("email_id:" + email_id + ",password:" + password + "\n")
				if err != nil {
					fmt.Println("Error while writing file")
				}
				fmt.Println("User successfully registered.")
				userOperations()
				break
			}

		}
	}
	defer file.Close()
}

func login() {
	fmt.Println("\nPlease enter email_id")
	fmt.Scanf("%s", &email_id)
	fmt.Println("Please enter Password")
	fmt.Scanf("%s", &password)
	file, err := os.OpenFile("user_credentials.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error while opening file")
	}
	scanner := bufio.NewScanner(file)
	flag := false
	for scanner.Scan() {
		line := scanner.Text()
		half_part := strings.Split(line, ",")
		e_temp := []string{half_part[0]}
		e_half_line := strings.Join(e_temp, "")
		e_parts := strings.Split(e_half_line, ":")
		p_temp := []string{half_part[1]}
		p_half_line := strings.Join(p_temp, "")
		p_parts := strings.Split(p_half_line, ":")
		if e_parts[1] == email_id && p_parts[1] == password {
			flag = true
			userOperations()
			break
		}
	}
	if flag == false {
		fmt.Println("\nInvalid Email ID or Password")
		registration()
	}
	defer file.Close()
}

func getList() {
	file, err := os.OpenFile(email_id+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {

	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	userOperations()
}

func createJournal() {
	fmt.Print("Enter journal entry without pressing enter key:\n")
	file, err := os.OpenFile(email_id+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error while opening file")
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount < 50 {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		current_time := time.Now()
		formatted_time := fmt.Sprintf("%02d %s %d %02d:%02d:%02d",
			current_time.Day(), current_time.Month(), current_time.Year(),
			current_time.Hour(), current_time.Minute(), current_time.Second())
		_, err = file.WriteString(formatted_time + "-" + text)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		fmt.Println("\nJournal entry added successfully")
		userOperations()
	} else {
		file, err := os.OpenFile(email_id+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error while opening file")
		}
		temp_file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error while opening file")
		}
		scanner := bufio.NewScanner(file)
		count := 1
		for scanner.Scan() {
			line := scanner.Text()
			if count == 1 {
				count++
				continue
			}
			_, err = temp_file.WriteString(line + "\n")
			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			count++
		}
		new_reader := bufio.NewReader(os.Stdin)
		new_text, _ := new_reader.ReadString('\n')
		current_time := time.Now()
		formatted_time := fmt.Sprintf("%02d %s %d %02d:%02d:%02d",
			current_time.Day(), current_time.Month(), current_time.Year(),
			current_time.Hour(), current_time.Minute(), current_time.Second())
		_, err = temp_file.WriteString(formatted_time + "-" + new_text)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		err = os.Remove(email_id + ".txt")
		if err != nil {
			fmt.Println("Error while removing file.")
		}
		err = os.Rename("temp.txt", email_id+".txt")
		if err != nil {
			fmt.Println("Error while renaming file.")
		}
	}
	fmt.Println("\nJournal entry added successfully")
	userOperations()
	defer file.Close()
}

// jornal replace karna h 1st and 50th wala
