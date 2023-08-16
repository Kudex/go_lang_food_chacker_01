package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
a:=app.New()
w:= a.NewWindow("GoCalcApp")
w.Resize(fyne.NewSize(300, 500))

title:=widget.NewLabel("Creating of the order")

name_label:=widget.NewLabel("Name:")
name:=widget.NewEntry()

food_label:=widget.NewLabel("Make a choice :")
food:=widget.NewCheckGroup([]string{"Cola","Burger","Pizza","Salat","Cheese"}, func(s [] string) {})

gender_label:=widget.NewLabel("Gender: ")
gender:=widget.NewRadioGroup([]string{"female","male"},func(s string){})

result:=widget.NewLabel("")

btn:=widget.NewButton("Create an order", func(){
	username:=name.Text
	food_arr:=food.Selected
	gender_val:=gender.Selected

	result.SetText(username+"\n"+gender_val+"\n")

	for _,i:=range food_arr{
		result.SetText(result.Text+i+"\n")
	}
})

w.SetContent(container.NewVBox(
	title,name_label,name,food_label,food,gender_label,gender,btn,result,
))
w.ShowAndRun()
}