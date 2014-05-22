package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var java *Java = &Java{js.Global.Get("Java")}

type Java struct {
	js.Object
}

func (java *Java) Type(typeName string) js.Object {
	return java.Call("type", typeName)
}

type Stage struct {
	js.Object
}

func (stage *Stage) SetScene(scene *Scene) {
	stage.Call("setScene", stage)
}

func (stage *Stage) Show() {
	stage.Call("show")
}

type Pane struct {
	js.Object
}

func (pane *Pane) GetChildren() *NodeList {
	return &NodeList{pane.Call("getChildren")}
}

type StackPane struct {
	*Pane
}

func NewStackPane() *StackPane {
	return &StackPane{&Pane{java.Type("javafx.scene.layout.StackPane").New()}}
}

type NodeList struct {
	js.Object
}

func (nodeList *NodeList) Add(node *Node) {
	nodeList.Call("add", node)
}

type Node struct {
	js.Object
}

type Label struct {
	*Node
}

func NewLabel(text string) *Label {
	return &Label{&Node{java.Type("javafx.scene.control.Label").New(text)}}
}

type Scene struct {
	js.Object
}

func NewScene(pane *Pane, width, height int) *Scene {
	return &Scene{java.Type("javafx.scene.Scene").New(pane, width, height)}
}

func main() {
}

func start(stage *Stage) {
	stage.Set("title", "Hello World!")
	label := NewLabel("Hello world!")
	root := NewStackPane()
	root.GetChildren().Add(label.Node)

	scene := NewScene(root.Pane, 300, 200)
	stage.SetScene(scene)

	stage.Show()
}

/*
// Java の import 文の代わりにこんな感じに書ける
var StackPane = javafx.scene.layout.StackPane;
var Scene     = javafx.scene.Scene;
var Label     = javafx.scene.control.Label;
var Screen    = javafx.stage.Screen;

function start(stage) {
    stage.title = "Hello World!";

    var label = new Label("Hello world!");
    var root  = new StackPane();
    root.getChildren().add(label);

    var scene = new Scene(root, 300, 200);
    stage.setScene(scene);

    stage.show();
}
*/
