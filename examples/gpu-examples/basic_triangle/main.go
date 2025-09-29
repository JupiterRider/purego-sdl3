package main

import (
	"github.com/jupiterrider/purego-sdl3/examples/gpu-examples/internal/common"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

var (
	useWireframeMode bool            = false
	useSmallViewport bool            = false
	useScissorRect   bool            = false
	smallViewport    sdl.GPUViewport = sdl.GPUViewport{
		X:        160,
		Y:        120,
		W:        320,
		H:        240,
		MinDepth: 0.1,
		MaxDepth: 1,
	}
	scissorRect sdl.Rect = sdl.Rect{
		X: 320,
		Y: 240,
		W: 320,
		H: 240,
	}
)

func main() {
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	device := sdl.CreateGPUDevice(sdl.GPUShaderFormatSPIRV|sdl.GPUShaderFormatDXIL|sdl.GPUShaderFormatMSL, false, "")
	defer sdl.DestroyGPUDevice(device)

	window := sdl.CreateWindow("gpu basic triangle example", 640, 480, sdl.WindowResizable)
	if window == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)

	if !sdl.ClaimWindowForGPUDevice(device, window) {
		panic(sdl.GetError())
	}
	defer sdl.ReleaseWindowFromGPUDevice(device, window)

	// Create the shaders
	vertexShader := common.LoadShader(device, "RawTriangle.vert", 0, 0, 0, 0)
	if vertexShader == nil {
		sdl.Log("Failed to create vertex shader!")
		return
	}

	fragmentShader := common.LoadShader(device, "SolidColor.frag", 0, 0, 0, 0)
	if fragmentShader == nil {
		sdl.Log("Failed to create fragment shader!")
		return
	}

	colorTargetDesc := []sdl.GPUColorTargetDescription{
		{
			Format: sdl.GetGPUSwapchainTextureFormat(device, window),
		},
	}

	// Create the pipelines
	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
		PrimitiveType:  sdl.GPUPrimitiveTypeTrianglelist,
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			ColorTargetDescriptions: &colorTargetDesc[0],
			NumColorTargets:         1,
		},
	}

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPUFillModeFill
	fillPipeline := sdl.CreateGPUGraphicsPipeline(device, &pipelineCreateInfo)

	if fillPipeline == nil {
		sdl.Log("Failed to create fill pipeline!")
		return
	}
	defer sdl.ReleaseGPUGraphicsPipeline(device, fillPipeline)

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPUFillModeLine
	linePipeline := sdl.CreateGPUGraphicsPipeline(device, &pipelineCreateInfo)
	if linePipeline == nil {
		sdl.Log("Failed to create line pipeline!")
		return
	}
	defer sdl.ReleaseGPUGraphicsPipeline(device, linePipeline)

	// Clean up shader resources
	sdl.ReleaseGPUShader(device, vertexShader)
	sdl.ReleaseGPUShader(device, fragmentShader)

	sdl.Log("Press Left to toggle wireframe mode")
	sdl.Log("Press Down to toggle small viewport")
	sdl.Log("Press Right to toggle scissor rect")

Outer:
	for {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				break Outer
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					break Outer
				}

				if event.Key().Key == sdl.KeycodeLeft {
					useWireframeMode = !useWireframeMode
				}

				if event.Key().Key == sdl.KeycodeDown {
					useSmallViewport = !useSmallViewport
				}

				if event.Key().Key == sdl.KeycodeRight {
					useScissorRect = !useScissorRect
				}
			}
		}

		commandBuffer := sdl.AcquireGPUCommandBuffer(device)
		if commandBuffer == nil {
			sdl.Log("AcquireGPUCommandBuffer failed: %s", sdl.GetError())
			return
		}

		var swapchainTexture *sdl.GPUTexture
		if !sdl.WaitAndAcquireGPUSwapchainTexture(commandBuffer, window, &swapchainTexture, nil, nil) {
			sdl.Log("WaitAndAcquireGPUSwapchainTexture failed: %s", sdl.GetError())
			return
		}

		if swapchainTexture != nil {
			colorTargetInfo := sdl.GPUColorTargetInfo{}
			colorTargetInfo.Texture = swapchainTexture
			colorTargetInfo.ClearColor = sdl.FColor{R: 0, G: 0, B: 0, A: 1}
			colorTargetInfo.LoadOp = sdl.GPULoadOpClear
			colorTargetInfo.StoreOp = sdl.GPUStoreOpStore

			renderPass := sdl.BeginGPURenderPass(commandBuffer, []sdl.GPUColorTargetInfo{colorTargetInfo}, nil)
			if useWireframeMode {
				sdl.BindGPUGraphicsPipeline(renderPass, linePipeline)
			} else {
				sdl.BindGPUGraphicsPipeline(renderPass, fillPipeline)
			}

			if useSmallViewport {
				sdl.SetGPUViewport(renderPass, &smallViewport)
			}

			if useScissorRect {
				sdl.SetGPUScissor(renderPass, &scissorRect)
			}
			sdl.DrawGPUPrimitives(renderPass, 3, 1, 0, 0)
			sdl.EndGPURenderPass(renderPass)
		}

		if !sdl.SubmitGPUCommandBuffer(commandBuffer) {
			panic(sdl.GetError())
		}
	}
}
