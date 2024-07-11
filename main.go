package main
import (
  "os/exec"
  "log"
  "fmt"
  "os"
  tea "github.com/charmbracelet/bubbletea"
)


// basically there are 3 things I need to build: the Model, the View, and the Update
// So the model, straight out of the example:
// this defines a new type called model, that has 3 attributes. 
type model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    selected map[int]struct{}   // which to-do items are selected
}


// this is a method that just instantiates the initial state of the model. There are different ways to do this but this seems simple enough...
func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices:  []string{"Run GO", "Run python", "Run Bash"},

		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}


// Init method is required. In our case we don't want it to do anything, but you could have it do some initial I/O or whatever. 
func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

// Update method is required. This method gets called whenever a message is recieved. Right now this returns a updated model, however 
// in a little bit this is where we are probably going to have it return a Cmd to go run one of the scripts...
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

// View method is required. This actually is what we end up using to render the UI. 
// This is the cool part of BT, this is just a string but BT just draws it for ya and it looks nice.
func (m model) View() string {
    // The header
    s := "What should we run?\n\n"

    // Iterate over our choices
    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
            // fmt.Printf("%s\n", choice)
            if choice == "Run python" { 
                runPython() 
            }
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
        
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}


func runPython() {
  cmd, err := exec.Command("python", "hello.py").Output()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(cmd))
  return 
}


func runBash() {
  cmd, err := exec.Command("bash","hello.sh").Output()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(cmd))
}


func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}


//func main () {
  //cmd, err := exec.Command("bash", "hello.sh").Output()
  //if err != nil {
   // log.Fatal(err)
  //}
  //fmt.Println(string(cmd))
  
 // cmd2, err2 := exec.Command("python", "hello.py").Output()
   // if err2 != nil {
     // log.Fatal(err2)
 //   }
 // fmt.Println(string(cmd2))
//}
