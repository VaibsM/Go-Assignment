package main
import "fmt"

var m map[string]int;

func add(key string, val int) {
    m[key] = val;
}

func retrieve(key string) int {
    val, ok := m[key];
    if !ok {
        return -1;
    } else {
        return val;
    }
}

func length() int {
    return len(m);
}

func remove(key string) {
    delete(m, key);
}

func printAll() {
    for key, value := range m {
        fmt.Printf("%q: %d\n", key, value)
    }
}

func main() {
    var temp string;
    var val int;
    var loop bool;
    var a int;
    loop = true;
    m = make(map[string]int);
    for loop {
		fmt.Print("\n1. Add a Person\n2. Check a Person's age\n3. Change a Person's age\n4. Check number of people added\n5. Delete a record\n6. View Entries\n7. Exit\nEnter your choice: ");
		fmt.Scan(&a);
		switch a {
		case 1:
		    fmt.Print("Enter Person's name: ");
		    fmt.Scanln(&temp);
		    fmt.Print("Enter Person's age: ");
		    fmt.Scanln(&val);
		    add(temp, val);
		case 2:
		    fmt.Print("Enter Person's name: ");
		    fmt.Scanln(&temp);
		    t := retrieve(temp);
		    if t != -1 {
		        fmt.Println("Person's age is: ", t);
		    } else {
		        fmt.Println("Person not found");
		    }
		case 3:
		    fmt.Print("Enter Person's name: ");
		    fmt.Scanln(&temp);
		    t := m[temp];
		    if t != 0 {
		        fmt.Print("Enter Person's age: ");
		        fmt.Scanln(&val);
		        add(temp, val);
		    } else {
		        fmt.Println("Person not found");
		    }
		case 4:
			fmt.Println("Number of people added: ", length());
		case 5:
		    fmt.Print("Enter Person's name: ");
		    fmt.Scanln(&temp);
		    t := m[temp];
		    if t != 0 {
		        remove(temp);
		    } else {
		        fmt.Println("Person not found");
		    }
		case 6:
		    printAll();
		case 7:
		    fmt.Println();
			loop = false;
		default:
			fmt.Println("\n\nInvalid Input\n");
		}
	}
}
