package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lngramos/three"
)

var (
	scene    *three.Scene
	camera   three.PerspectiveCamera
	renderer three.WebGLRenderer
	mesh     *three.Mesh
)

func main() {
	document := js.Global.Get("document")
	windowWidth := js.Global.Get("innerWidth").Float()
	windowHeight := js.Global.Get("innerHeight").Float()
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	camera.Position.Set(0, 0, 400)

	scene = three.NewScene()

	light := three.NewDirectionalLight(three.NewColor(126, 255, 255), 0.5)
	light.Position.Set(256, 256, 256).Normalize()
	scene.Add(light)

	renderer = three.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	document.Get("body").Call("appendChild", renderer.Get("domElement"))

	// Create cube
	geometry := three.NewBoxGeometry(&three.BoxGeometryParameters{
		Width:  128,
		Height: 128,
		Depth:  128,
	})

	// geometry2 := three.NewCircleGeometry(three.CircleGeometryParameters{
	// 	Radius:      50,
	// 	Segments:    20,
	// 	ThetaStart:  0,
	// 	ThetaLength: 2,
	// })

	materialParams := three.NewMaterialParameters()
	materialParams.Color = three.NewColor(0, 123, 211)
	materialParams.Shading = three.SmoothShading
	materialParams.Side = three.FrontSide
	material := three.NewMeshBasicMaterial(materialParams)
	// material := three.NewMeshLambertMaterial(materialParams)
	// material := three.NewMeshPhongMaterial(materialParams)
	mesh = three.NewMesh(geometry, material)

	scene.Add(mesh)

	animate()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	pos := mesh.Object.Get("rotation")
	pos.Set("x", pos.Get("x").Float()+float64(0.01))
	pos.Set("y", pos.Get("y").Float()+float64(0.01))

	renderer.Render(scene, camera)
}
