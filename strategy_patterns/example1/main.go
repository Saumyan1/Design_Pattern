package main

import (
	"fmt"
	"strings"
)


//TEXT_FORMATTER - CONTEXT

type TextEditor struct{
	Format Formatter;

}

//Constructor
func NewTextEditor(f Formatter) *TextEditor{
	return &TextEditor{
		Format: f,
	}
}

func (t *TextEditor)SetFormatter(f Formatter) {
	t.Format = f;
}

func(t* TextEditor) Printt(st string){
	fmt.Println(t.Format.Format(st))
}



//Formatter - Strategy INTERFACE
type Formatter interface{
	Format(string)string

}



//CONCRETE STRATEGIES
type UpperCase struct{

}

type LowerCase struct{

}

func (uc *UpperCase) Format(st string)string{
	return strings.ToUpper(st)
}

func(lc *LowerCase) Format(st string) string{
	return strings.ToLower(st)
}


//CLIENT
func main(){
	fmt.Println("This is design pattern")
	upper := &UpperCase{}
	lower := &LowerCase{}
	editor := NewTextEditor(upper)
	editor.Printt("hello world")
	editor.SetFormatter(lower)
	editor.Printt("HELLO WORLD")
}