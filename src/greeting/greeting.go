package greeting

import "fmt"

type Salutation struct  {
	Name string
	Greeting string
}

type Renamable interface {
	Rename(newName string)
}


func (salutation *Salutation) Rename(newName string){
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

type Saluations []Salutation

type Printer func(string)()

func (salutations Saluations) Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		}else {
			do(alternate)
		}
	}
}
func GetPrefix(name string)(prefix string){

	prefixMap := map[string]string{
		"Bob": "Mr ",
		"Joe": "Dr ",
		"Amy": "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr Dr "

	delete(prefixMap, "Mary")
	return prefixMap[name]

}

func CreateMessage(name, greeting string) (message string, alternate string){
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}
func Print(s string){
	fmt.Print(s)
}
func PrintLine(s string){
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer{
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func PrintCustom(s string, custom string){
	fmt.Println(s + custom)
}
